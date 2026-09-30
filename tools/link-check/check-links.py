#!/usr/bin/env python3
"""Report links that do not resolve to a docs.snyk.io page.

Each root folder that has a SUMMARY.md is published as its own GitBook space.
GitBook resolves relative links only inside one space, and only when the
target file exists. Any other relative link is rendered on docs.snyk.io as a
github.com URL instead of a docs page.

Errors (the link is broken on docs.snyk.io):
  cross-space   Relative link to a file in another section. Use an absolute
                https://docs.snyk.io/... URL.
  missing       Relative link to a file that does not exist.
  github        Hardcoded github.com/snyk/user-docs file URL. Use a relative
                link inside the section, or a docs.snyk.io URL across sections.

Warnings (the link works only while the target page stays where it is):
  gitbook-app   GitBook editor URL (app.gitbook.com). Readers without a GitBook
                account cannot open it if GitBook fails to resolve it. Use a
                docs.snyk.io URL.

Checks Markdown links, Markdown images, and HTML href and src attributes.
Content inside fenced code blocks is skipped.

Usage:
  check-links.py [file ...]   Check the named files
  check-links.py              Check every tracked Markdown file

Exits 1 when any error is found, 0 otherwise. Warnings do not change the exit
status.
"""

import re
import subprocess
import sys
from pathlib import Path
from urllib.parse import unquote

REPO_ROOT = Path(__file__).resolve().parents[2]

# ](target), ](<target>), or ](target "title"). Targets can contain spaces
# and one level of parentheses, as in GitBook asset names like image (60).png.
MD_LINK = re.compile(r"\]\(\s*<?((?:[^()<>]|\([^()]*\))+?)>?(?:\s+\"[^\"]*\")?\s*\)")
HTML_ATTR = re.compile(r"""\b(?:href|src)\s*=\s*["']([^"']+)["']""")
FENCE = re.compile(r"^\s*(```|~~~)")
GITHUB_FILE = re.compile(r"^https?://github\.com/snyk/user-docs/(?:blob|tree)/")
GITBOOK_APP = re.compile(r"^https?://app\.gitbook\.com/[so]/")


def sections():
    return {p.parent.name for p in REPO_ROOT.glob("*/SUMMARY.md")}


def section_of(path, known):
    """Return the root section a repo-relative path belongs to, or None."""
    parts = path.parts
    if parts and parts[0] in known:
        return parts[0]
    return None


def is_relative(target):
    return not (
        target.startswith(("#", "/", "mailto:", "tel:", "{{"))
        or re.match(r"^[a-zA-Z][a-zA-Z0-9+.-]*:", target)
    )


def check_target(rel_path, source_section, target, known):
    """Return (severity, kind, message) for a bad link, or None."""
    if GITHUB_FILE.match(target):
        return ("error", "github",
                "link points to GitHub instead of the docs site")
    if GITBOOK_APP.match(target):
        return ("warning", "gitbook-app",
                "link points to the GitBook editor instead of the docs site")
    if not is_relative(target):
        return None

    path_part = target.split("#", 1)[0].split("?", 1)[0]
    if not path_part:
        return None
    # GitBook escapes some characters in exported links, for example \&.
    path_part = unquote(path_part).replace("\\", "")
    resolved = (REPO_ROOT / rel_path.parent / path_part).resolve()
    try:
        resolved_rel = resolved.relative_to(REPO_ROOT)
    except ValueError:
        return ("error", "missing", "relative link points outside the repository")

    target_section = section_of(resolved_rel, known)
    if target_section and target_section != source_section:
        return ("error", "cross-space",
                f"relative link crosses into {target_section}/")
    if not resolved.exists():
        return ("error", "missing", "relative link target does not exist")
    return None


def check_file(rel_path, known):
    source_section = section_of(rel_path, known)
    if source_section is None:
        return []

    findings = []
    in_fence = False
    text = (REPO_ROOT / rel_path).read_text(encoding="utf-8", errors="replace")
    for lineno, line in enumerate(text.splitlines(), start=1):
        if FENCE.match(line):
            in_fence = not in_fence
            continue
        if in_fence:
            continue

        for match in list(MD_LINK.finditer(line)) + list(HTML_ATTR.finditer(line)):
            target = match.group(1).strip()
            result = check_target(rel_path, source_section, target, known)
            if result:
                severity, kind, message = result
                findings.append(
                    (severity, f"{rel_path}:{lineno}: {severity}: {kind}: {message}: {target}")
                )
    return findings


def main(argv):
    known = sections()
    if argv:
        files = [Path(f) for f in argv]
    else:
        out = subprocess.run(
            ["git", "ls-files", "*.md"],
            cwd=REPO_ROOT, capture_output=True, text=True, check=True,
        ).stdout
        files = [Path(f) for f in out.splitlines()]

    findings = []
    for f in files:
        if (REPO_ROOT / f).is_file():
            findings.extend(check_file(f, known))

    for _, finding in findings:
        print(finding)
    return 1 if any(severity == "error" for severity, _ in findings) else 0


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))
