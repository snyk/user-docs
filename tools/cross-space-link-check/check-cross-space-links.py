#!/usr/bin/env python3
"""Report relative links that point into a different docs section.

Each root folder that has a SUMMARY.md is published as its own GitBook space.
GitBook resolves relative links only inside one space. A relative link to a
file in another section, such as ../platform-administration/page.md, is
rendered on docs.snyk.io as a github.com URL instead of a docs page.

Use an absolute https://docs.snyk.io/... URL for links between sections.

Checks Markdown links, Markdown images, and HTML href and src attributes.
Content inside fenced code blocks is skipped.

Usage:
  check-cross-space-links.py [file ...]   Check the named files
  check-cross-space-links.py              Check every tracked Markdown file

Exits 0 when no violations are found, 1 otherwise.
"""

import re
import subprocess
import sys
from pathlib import Path

REPO_ROOT = Path(__file__).resolve().parents[2]

# ](target) or ](target "title"), and href="target" / src="target".
MD_LINK = re.compile(r"\]\(\s*<?([^)\s>]+)")
HTML_ATTR = re.compile(r"""\b(?:href|src)\s*=\s*["']([^"']+)["']""")
FENCE = re.compile(r"^\s*(```|~~~)")


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
            target = match.group(1)
            if not is_relative(target):
                continue
            path_part = target.split("#", 1)[0].split("?", 1)[0]
            if not path_part:
                continue
            resolved = (REPO_ROOT / rel_path.parent / path_part).resolve()
            try:
                resolved_rel = resolved.relative_to(REPO_ROOT)
            except ValueError:
                continue
            target_section = section_of(resolved_rel, known)
            if target_section and target_section != source_section:
                findings.append(
                    f"{rel_path}:{lineno}: relative link crosses into "
                    f"{target_section}/: {target}"
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

    for finding in findings:
        print(finding)
    return 1 if findings else 0


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))
