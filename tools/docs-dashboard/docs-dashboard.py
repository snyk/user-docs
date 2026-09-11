#!/usr/bin/env python3
"""Report the health of the documentation in this repository.

Measures four things a writer can act on:

  Inventory     how many pages and words each space holds
  Navigation    pages missing from SUMMARY.md, and SUMMARY.md entries with no page
  Links         absolute docs.snyk.io links, and relative links whose target is gone
  Attention     which pages have been edited lately, and which nobody has touched

Usage:
  docs-dashboard.py                      Confluence HTML body on stdout
  docs-dashboard.py --format json        the same figures as JSON
  docs-dashboard.py --output page.html   write instead of printing

The HTML is the Confluence flavour that maps to ADF, not storage format, so it
can be passed straight to the Confluence page API as a page body.

Generated pages are excluded from the navigation, link, and attention findings;
see is_generated below. They are still counted in the inventory, because they
are real pages on the site.

Reads the working tree and git history. Writes nothing but its own output.
"""

from __future__ import annotations

import argparse
import html
import json
import os
import re
import subprocess
import sys
from collections import Counter, defaultdict
from datetime import datetime, timedelta, timezone

# Paths whose contents no writer edits by hand. A finding in one of these is
# noise: the fix belongs in the generator or the upstream repository, not in the
# page. Kept in step with the generated-files table in AGENTS.md and with the
# exclusions in tools/alt-text-check/check-alt-text.sh.
GENERATED_PREFIXES = (
    "developer-tools/snyk-api/reference/",
    "developer-tools/snyk-cli/commands/",
)
GENERATED_FILES = (
    "developer-tools/snyk-api/changelog.md",
    "developer-tools/snyk-ide-plugins-and-extensions/compatibility-matrix.md",
    "scan-fix-and-prevent/scan-with-snyk/error-catalog.md",
)

# Commits that GitBook pushed to main on a writer's behalf, rather than commits
# that arrived through a pull request. AGENTS.md explains the two-way sync.
GITBOOK_SUBJECT = re.compile(r"^GITBOOK-\d+")
GITBOOK_COMMITTER_EMAIL = "ghost@gitbook.com"

# Accounts that commit as automation. Counting them among contributors would
# bury the humans.
BOT_AUTHORS = {
    "gitbook-bot",
    "github-actions[bot]",
    "dependabot[bot]",
    "snyk-bot",
    "snyk-docs-bot",
}

SUMMARY_ENTRY = re.compile(r"^(?P<indent>\s*)\*\s+\[(?P<title>[^\]]*)\]\((?P<target>[^)]*)\)")
MARKDOWN_LINK = re.compile(r"\[[^\]]*\]\(([^)\s]+)(?:\s+\"[^\"]*\")?\)")
CONTENT_REF = re.compile(r"\{%\s*content-ref\s+url=\"([^\"]+)\"")
ABSOLUTE_DOCS_LINK = re.compile(r"https://docs\.snyk\.io[^\s)\"'<>\]]*")
CODE_FENCE = re.compile(r"^\s*```")

# Anything of the form "scheme:" is an address for something other than a file
# in this repository — https:, mailto:, and the editor deep links the IDE
# quickstarts use, such as cursor: and antigravity:.
URI_SCHEME = re.compile(r"^[a-zA-Z][a-zA-Z0-9+.\-]*:")

# GitBook writes this marker in place of a link whose target it can no longer
# resolve, so it names a broken link precisely rather than by inference.
GITBOOK_BROKEN_MARKER = "/broken/"

# A commit touching more than this many files is a repository-wide mechanical
# change — a folder move, a frontmatter sweep — and says nothing about whether
# any single page was reviewed. Counting one as an edit would report the whole
# site as freshly written. See --bulk-threshold.
BULK_COMMIT_FILES = 200


def run_git(repo: str, *args: str) -> str:
    result = subprocess.run(
        ("git", "-C", repo) + args,
        check=True,
        capture_output=True,
        text=True,
        errors="replace",
    )
    return result.stdout


def is_generated(relpath: str) -> bool:
    return relpath.startswith(GENERATED_PREFIXES) or relpath in GENERATED_FILES


def discover_spaces(repo: str) -> list[str]:
    """A directory holding a SUMMARY.md is a GitBook space; nothing else is."""
    spaces = []
    for entry in sorted(os.listdir(repo)):
        if entry.startswith("."):
            continue
        if os.path.isfile(os.path.join(repo, entry, "SUMMARY.md")):
            spaces.append(entry)
    return spaces


def tracked_markdown(repo: str) -> list[str]:
    listing = run_git(repo, "ls-files", "*.md")
    return [line for line in listing.splitlines() if line]


def is_page(relpath: str, space: str) -> bool:
    """Pages are what the site renders. SUMMARY.md is navigation, and the files
    under .gitbook/ are includes and assets that are transcluded, not linked."""
    if relpath == f"{space}/SUMMARY.md":
        return False
    return not relpath.startswith(f"{space}/.gitbook/")


def count_words(text: str) -> int:
    """Prose words only: fenced code blocks and HTML tags are not writing."""
    words = 0
    in_fence = False
    for line in text.splitlines():
        if CODE_FENCE.match(line):
            in_fence = not in_fence
            continue
        if in_fence:
            continue
        stripped = re.sub(r"<[^>]+>", " ", line)
        words += len(stripped.split())
    return words


def parse_summary(repo: str, space: str) -> list[dict]:
    """Read one space's table of contents.

    Returns an entry per navigation line, with the target normalised to a path
    from the repository root. Indentation carries the nesting depth, two spaces
    per level, which is the only place the site's menu depth is recorded.
    """
    entries = []
    path = os.path.join(repo, space, "SUMMARY.md")
    with open(path, encoding="utf-8") as handle:
        for number, line in enumerate(handle, start=1):
            match = SUMMARY_ENTRY.match(line)
            if not match:
                continue
            target = match.group("target").strip()
            if not target or target.startswith(("http://", "https://", "mailto:", "#")):
                continue
            target = target.split("#", 1)[0]
            if not target:
                continue
            entries.append(
                {
                    "space": space,
                    "line": number,
                    "depth": len(match.group("indent")) // 2,
                    "title": match.group("title").strip(),
                    "target": os.path.normpath(os.path.join(space, target)),
                }
            )
    return entries


def resolve_link(source: str, target: str) -> str | None:
    """Turn a link written inside a page into a repository-root path.

    Returns None for anything that does not point at a file in this repository:
    external URLs, anchors, mail links, and the asset references that GitBook
    rewrites at build time.
    """
    target = target.strip()
    if not target or target.startswith(("#", "<")) or URI_SCHEME.match(target):
        return None
    target = target.split("#", 1)[0]
    if not target:
        return None
    target = target.replace("%20", " ")
    if target.startswith("/"):
        return os.path.normpath(target.lstrip("/"))
    return os.path.normpath(os.path.join(os.path.dirname(source), target))


def scan_pages(repo: str, spaces: list[str]) -> dict:
    """One pass over every page, collecting inventory and link findings."""
    pages: dict[str, dict] = {}
    absolute_links: list[dict] = []
    broken_links: list[dict] = []
    space_of = {}

    for relpath in tracked_markdown(repo):
        space = relpath.split("/", 1)[0]
        if space not in spaces or not is_page(relpath, space):
            continue
        space_of[relpath] = space

        with open(os.path.join(repo, relpath), encoding="utf-8", errors="replace") as handle:
            text = handle.read()

        pages[relpath] = {
            "space": space,
            "words": count_words(text),
            "generated": is_generated(relpath),
        }

        if is_generated(relpath):
            continue

        for number, line in enumerate(text.splitlines(), start=1):
            for hit in ABSOLUTE_DOCS_LINK.findall(line):
                absolute_links.append({"page": relpath, "space": space, "line": number, "url": hit})

        for target in MARKDOWN_LINK.findall(text) + CONTENT_REF.findall(text):
            resolved = resolve_link(relpath, target)
            if resolved is None:
                continue
            full = os.path.join(repo, resolved)
            if os.path.exists(full):
                continue
            # A link to a directory is valid when the directory has a README.
            if os.path.exists(os.path.join(full, "README.md")):
                continue
            broken_links.append(
                {
                    "page": relpath,
                    "space": space,
                    "target": target,
                    "resolved": resolved,
                    "gitbook_marker": target.startswith(GITBOOK_BROKEN_MARKER),
                }
            )

    return {
        "pages": pages,
        "space_of": space_of,
        "absolute_links": absolute_links,
        "broken_links": broken_links,
    }


def navigation_findings(repo: str, spaces: list[str], pages: dict) -> dict:
    """Compare each space's SUMMARY.md against the pages on disk.

    A page absent from SUMMARY.md does not exist on the site, and an entry whose
    file is missing is a dead menu item. Both are invisible without this check.
    """
    orphans: list[str] = []
    dangling: list[dict] = []
    by_space: dict[str, dict] = {}
    all_entries: list[dict] = []

    for space in spaces:
        entries = parse_summary(repo, space)
        all_entries.extend(entries)
        listed = {entry["target"] for entry in entries}

        space_pages = {path for path, meta in pages.items() if meta["space"] == space}
        space_orphans = sorted(
            path for path in space_pages - listed if not is_generated(path)
        )
        space_dangling = [
            entry
            for entry in entries
            if not os.path.exists(os.path.join(repo, entry["target"]))
        ]

        orphans.extend(space_orphans)
        dangling.extend(space_dangling)
        by_space[space] = {
            "entries": len(entries),
            "max_depth": max((entry["depth"] for entry in entries), default=0),
            "orphans": len(space_orphans),
            "dangling": len(space_dangling),
        }

    return {
        "by_space": by_space,
        "orphans": orphans,
        "dangling": dangling,
        "entries": len(all_entries),
    }


def stray_directories(repo: str, spaces: list[str]) -> list[str]:
    """Top-level directories holding Markdown that belongs to no space.

    Content here is unreachable: with no SUMMARY.md there is nothing to render it.
    """
    stray = []
    for relpath in tracked_markdown(repo):
        top = relpath.split("/", 1)[0]
        if top in spaces or top == "tools" or "/" not in relpath:
            continue
        stray.append(relpath)
    return sorted(stray)


def git_history(repo: str, bulk_threshold: int = BULK_COMMIT_FILES) -> dict:
    """Date every path twice: when it was last edited, and when it last moved.

    The distinction matters because most commits in this repository are not
    edits. Two kinds are excluded from the edit date, since counting either
    would report the whole site as freshly written:

      a pure rename (100% similarity), which relocates a page without changing
      a word of it

      a commit touching more than bulk_threshold files, which is a mechanical
      sweep — a folder move, a frontmatter pass — and not a decision about any
      one page

    Renames are still followed, so a page keeps its history across a move.
    git log walks newest first, so the first sighting of a path is its latest.

    Returns the edit dates, the touch dates (any commit at all, so that pages
    only ever caught up in a sweep can still be dated), and how many commits
    were set aside as bulk.
    """
    listing = run_git(repo, "log", "--format=%x01%ct", "--name-status", "--find-renames")

    last_edit: dict[str, int] = {}
    last_touch: dict[str, int] = {}
    alias: dict[str, str] = {}
    bulk_skipped = 0
    timestamp = 0
    batch: list[tuple[str, str]] = []

    def canonical(path: str) -> str:
        """Follow a path forward through the renames seen so far."""
        seen = set()
        while path in alias and path not in seen:
            seen.add(path)
            path = alias[path]
        return path

    def flush() -> None:
        nonlocal bulk_skipped
        if not batch:
            return
        bulk = len(batch) > bulk_threshold
        if bulk:
            bulk_skipped += 1
        for status, path in batch:
            target = canonical(path)
            last_touch.setdefault(target, timestamp)
            if not bulk and status != "R100":
                last_edit.setdefault(target, timestamp)
        batch.clear()

    for line in listing.splitlines():
        if line.startswith("\x01"):
            flush()
            timestamp = int(line[1:])
            continue
        if not line:
            continue
        fields = line.split("\t")
        status = fields[0]
        if status.startswith("R") and len(fields) >= 3:
            old, new = fields[1], fields[2]
            batch.append((status, new))
            # Recorded now, while walking backwards, so that older commits
            # naming the old path are credited to the page's current path.
            alias[old] = canonical(new)
        elif len(fields) >= 2:
            batch.append((status, fields[1]))

    flush()
    return {"last_edit": last_edit, "last_touch": last_touch, "bulk_commits_ignored": bulk_skipped}


def tree_rebuilt_at(repo: str, spaces: list[str]) -> int:
    """When the current directory layout came into being.

    The site was reorganised into spaces in 2026, and the legacy tree was
    deleted rather than moved, so git carries no rename link across it. No page
    at a current path can appear older than this, and saying so is the
    difference between an honest freshness figure and a flattering one.
    """
    earliest = []
    for space in spaces:
        listing = run_git(repo, "log", "--reverse", "--format=%ct", "--", space)
        first = listing.split("\n", 1)[0].strip()
        if first:
            earliest.append(int(first))
    return min(earliest) if earliest else 0


def normalise_name(name: str) -> str:
    return re.sub(r"[^a-z0-9]", "", name.lower())


def email_handle(email: str) -> str:
    """The identifying part of a commit email.

    GitHub's private addresses carry an account number the person never chose
    (12345+asergaz@users.noreply.github.com); the handle after it is the part
    that identifies them.
    """
    local = email.split("@", 1)[0]
    local = re.sub(r"^\d+\+", "", local)
    return re.sub(r"[^a-z0-9]", "", local.lower())


def merge_identities(identities: Counter) -> list[tuple[str, int]]:
    """Collapse one person's several git identities into a single contributor.

    People commit under a display name and a username, and from a work address
    and a GitHub private address, so a raw author count splits regulars into
    two or three entries and understates every one of them. Identities are
    joined when they share either a normalised name or an email handle, then
    the most human-looking spelling is kept as the label.
    """
    parent: dict[str, str] = {}

    def find(key: str) -> str:
        parent.setdefault(key, key)
        while parent[key] != key:
            parent[key] = parent[parent[key]]
            key = parent[key]
        return key

    def union(left: str, right: str) -> None:
        left_root, right_root = find(left), find(right)
        if left_root != right_root:
            parent[right_root] = left_root

    for name, email in identities:
        keys = [f"n:{normalise_name(name)}"]
        handle = email_handle(email)
        if handle:
            keys.append(f"e:{handle}")
            # A username used as a display name is the same person as the
            # address it was sent from.
            if normalise_name(name) == handle:
                union(keys[0], keys[1])
        for key in keys[1:]:
            union(keys[0], key)

    groups: dict[str, Counter] = defaultdict(Counter)
    for (name, email), count in identities.items():
        groups[find(f"n:{normalise_name(name)}")][name] += count

    people = []
    for spellings in groups.values():
        total = sum(spellings.values())
        # Prefer "Mike Romard" over "mikeromard": a name with a space is the
        # one the person writes for a human reader.
        label = max(spellings, key=lambda name: (" " in name, spellings[name], len(name)))
        people.append((label, total))

    people.sort(key=lambda item: (-item[1], item[0]))
    return people


def git_activity(repo: str, months: int = 12) -> dict:
    """Commit volume by month, split by how the change reached main.

    GitBook writes straight to main when a writer saves in the editor; everything
    else arrives through a pull request. The ratio is the clearest picture of how
    the documentation is actually maintained.
    """
    since = (datetime.now(timezone.utc) - timedelta(days=months * 31)).strftime("%Y-%m-%d")
    listing = run_git(
        repo,
        "log",
        f"--since={since}",
        "--no-merges",
        "--format=%ct%x00%ce%x00%an%x00%ae%x00%s",
    )

    by_month: dict[str, Counter] = defaultdict(Counter)
    identities: Counter = Counter()
    total = Counter()

    for line in listing.splitlines():
        parts = line.split("\x00")
        if len(parts) != 5:
            continue
        timestamp, committer_email, author, author_email, subject = parts
        month = datetime.fromtimestamp(int(timestamp), timezone.utc).strftime("%Y-%m")
        gitbook = committer_email == GITBOOK_COMMITTER_EMAIL or bool(GITBOOK_SUBJECT.match(subject))
        kind = "gitbook" if gitbook else "pull_request"
        by_month[month][kind] += 1
        total[kind] += 1
        if author not in BOT_AUTHORS:
            identities[(author, author_email)] += 1

    ordered = sorted(by_month)[-months:]
    contributors = merge_identities(identities)
    return {
        "months": [
            {
                "month": month,
                "gitbook": by_month[month]["gitbook"],
                "pull_request": by_month[month]["pull_request"],
                "total": sum(by_month[month].values()),
            }
            for month in ordered
        ],
        "totals": {"gitbook": total["gitbook"], "pull_request": total["pull_request"]},
        "contributors": contributors,
    }


FRESHNESS_BUCKETS = (
    ("month", "Within a month", 30),
    ("quarter", "1 to 3 months", 90),
    ("older", "Over 3 months", None),
)


def freshness(pages: dict, history: dict, now: int) -> dict:
    """Sort every hand-written page by how recently someone worked on it.

    A page nobody has edited since a bulk sweep put it there is counted apart
    from the rest. It is not stale in the usual sense — it may be perfectly
    correct — but no one has looked at it deliberately, so it is the group worth
    knowing about.
    """
    last_edit = history["last_edit"]
    last_touch = history["last_touch"]

    buckets = {key: 0 for key, _, _ in FRESHNESS_BUCKETS}
    buckets["untouched"] = 0
    by_space: dict[str, Counter] = defaultdict(Counter)
    edited: list[dict] = []
    untouched: list[dict] = []

    for relpath, meta in pages.items():
        if meta["generated"]:
            continue
        timestamp = last_edit.get(relpath)
        if timestamp is None:
            buckets["untouched"] += 1
            by_space[meta["space"]]["untouched"] += 1
            since = last_touch.get(relpath)
            untouched.append(
                {
                    "page": relpath,
                    "space": meta["space"],
                    "days": int((now - since) // 86400) if since else None,
                    "words": meta["words"],
                }
            )
            continue

        days = int((now - timestamp) // 86400)
        bucket = FRESHNESS_BUCKETS[-1][0]
        for key, _, limit in FRESHNESS_BUCKETS:
            if limit is not None and days <= limit:
                bucket = key
                break
        buckets[bucket] += 1
        by_space[meta["space"]][bucket] += 1
        edited.append({"page": relpath, "space": meta["space"], "days": days, "words": meta["words"]})

    edited.sort(key=lambda item: item["days"], reverse=True)
    untouched.sort(key=lambda item: item["words"], reverse=True)
    return {
        "buckets": buckets,
        "by_space": {key: dict(value) for key, value in by_space.items()},
        "oldest": edited,
        "untouched": untouched,
    }


def build_report(repo: str, bulk_threshold: int = BULK_COMMIT_FILES) -> dict:
    spaces = discover_spaces(repo)
    scan = scan_pages(repo, spaces)
    pages = scan["pages"]
    navigation = navigation_findings(repo, spaces, pages)
    history = git_history(repo, bulk_threshold)
    rebuilt_at = tree_rebuilt_at(repo, spaces)
    now = int(datetime.now(timezone.utc).timestamp())

    inventory = {}
    for space in spaces:
        space_pages = [meta for meta in pages.values() if meta["space"] == space]
        assets_dir = os.path.join(repo, space, ".gitbook", "assets")
        inventory[space] = {
            "pages": len(space_pages),
            "generated_pages": sum(1 for meta in space_pages if meta["generated"]),
            "words": sum(meta["words"] for meta in space_pages),
            "assets": len(os.listdir(assets_dir)) if os.path.isdir(assets_dir) else 0,
        }

    head = run_git(repo, "rev-parse", "--short", "HEAD").strip()
    return {
        "generated_at": datetime.now(timezone.utc).strftime("%Y-%m-%d %H:%M UTC"),
        "commit": head,
        "spaces": spaces,
        "inventory": inventory,
        "totals": {
            "pages": len(pages),
            "generated_pages": sum(1 for meta in pages.values() if meta["generated"]),
            "words": sum(meta["words"] for meta in pages.values()),
            "assets": sum(item["assets"] for item in inventory.values()),
        },
        "navigation": navigation,
        "stray": stray_directories(repo, spaces),
        "links": {
            "absolute": scan["absolute_links"],
            "broken": scan["broken_links"],
            "gitbook_markers": sum(1 for item in scan["broken_links"] if item["gitbook_marker"]),
        },
        "freshness": freshness(pages, history, now),
        "activity": git_activity(repo),
        "method": {
            "bulk_threshold": bulk_threshold,
            "bulk_commits_ignored": history["bulk_commits_ignored"],
            "tree_rebuilt_at": datetime.fromtimestamp(rebuilt_at, timezone.utc).strftime("%Y-%m-%d")
            if rebuilt_at
            else None,
            "tree_age_days": int((now - rebuilt_at) // 86400) if rebuilt_at else None,
        },
    }


# --- Confluence rendering -------------------------------------------------
#
# The Confluence page API takes an HTML dialect that maps onto ADF. Only the
# constructs used below are safe: headings, paragraphs, tables, panels, status
# lozenges, and <details> expands. Storage-format macros are not accepted.

REPO_URL = "https://github.com/snyk/user-docs/blob/main"


def esc(text: object) -> str:
    return html.escape(str(text), quote=True)


def lozenge(text: str, colour: str) -> str:
    return f'<span data-type="status" data-color="{colour}">{esc(text)}</span>'


def verdict(count: int, warn: int, bad: int) -> str:
    if count >= bad:
        return lozenge(f"{count:,}", "red")
    if count >= warn:
        return lozenge(f"{count:,}", "yellow")
    return lozenge(f"{count:,}", "green")


def page_link(relpath: str) -> str:
    return f'<a href="{REPO_URL}/{esc(relpath)}">{esc(relpath)}</a>'


def table(headers: list[str], rows: list[list[str]], widths: list[int] | None = None) -> str:
    parts = ['<table data-layout="default"><thead><tr>']
    for index, header in enumerate(headers):
        width = f' data-colwidth="{widths[index]}"' if widths else ""
        parts.append(f"<th{width}><p><strong>{esc(header)}</strong></p></th>")
    parts.append("</tr></thead><tbody>")
    for row in rows:
        parts.append("<tr>")
        for index, cell in enumerate(row):
            width = f' data-colwidth="{widths[index]}"' if widths else ""
            parts.append(f"<td{width}><p>{cell}</p></td>")
        parts.append("</tr>")
    parts.append("</tbody></table>")
    return "".join(parts)


def expand(title: str, body: str) -> str:
    return f"<details><summary>{esc(title)}</summary>{body}</details>"


def bar(value: int, peak: int, width: int = 24) -> str:
    """A proportional bar drawn in block characters.

    Confluence has no chart element in this HTML dialect, and an image would
    need an attachment upload, so the trend is drawn in text.
    """
    if peak <= 0:
        return ""
    filled = max(1, round(value / peak * width)) if value else 0
    return "█" * filled


def render_html(report: dict) -> str:
    out: list[str] = []
    totals = report["totals"]
    navigation = report["navigation"]
    links = report["links"]
    fresh = report["freshness"]
    activity = report["activity"]

    hand_written = totals["pages"] - totals["generated_pages"]
    never_edited = fresh["buckets"]["untouched"]
    edited = hand_written - never_edited
    edited_pct = round(edited / hand_written * 100) if hand_written else 0

    out.append(
        '<div data-type="panel-info"><p>Health of the documentation source in '
        f'<a href="https://github.com/snyk/user-docs">snyk/user-docs</a>, the repository behind '
        f'<a href="https://docs.snyk.io">docs.snyk.io</a>. Generated from commit <code>{esc(report["commit"])}</code> '
        f'on {esc(report["generated_at"])}.</p>'
        "<p>Every figure is derived from the repository itself, so the page can be regenerated at any time "
        "and needs no service to stay up. Pages produced by a generator are counted in the inventory but "
        "excluded from the navigation, link, and freshness findings, because no writer can act on them.</p></div>"
    )

    out.append("<h2>At a glance</h2>")
    out.append(
        table(
            ["Measure", "Value", "What it means"],
            [
                [
                    "Published pages",
                    f"<strong>{totals['pages']:,}</strong>",
                    f"Across {len(report['spaces'])} spaces, of which {totals['generated_pages']:,} are generated "
                    f"and {hand_written:,} are hand-written",
                ],
                [
                    "Words of prose",
                    f"<strong>{totals['words']:,}</strong>",
                    "Excluding code blocks and HTML markup",
                ],
                [
                    "Edited since the restructure",
                    f"<strong>{edited_pct}%</strong>",
                    f"{edited:,} of {hand_written:,} hand-written pages; the remaining {never_edited:,} "
                    "were carried across and never individually edited",
                ],
                [
                    "Pages missing from navigation",
                    verdict(len(navigation["orphans"]), 1, 25),
                    "Present in the repository but absent from SUMMARY.md, so unreachable on the site",
                ],
                [
                    "Dead navigation entries",
                    verdict(len(navigation["dangling"]), 1, 10),
                    "Listed in SUMMARY.md with no page behind them",
                ],
                [
                    "Absolute docs.snyk.io links",
                    verdict(len(links["absolute"]), 1, 100),
                    "Should be relative paths; absolute links break whenever a page moves",
                ],
                [
                    "Broken relative links",
                    verdict(len(links["broken"]), 1, 25),
                    "Point at a file that is not in the repository",
                ],
            ],
            widths=[220, 110, 430],
        )
    )

    out.append("<h2>Inventory by space</h2>")
    out.append(
        "<p>One row per GitBook space. Depth is how many levels the space's navigation nests; "
        "a deep menu is harder to browse than a broad one.</p>"
    )
    rows = []
    for space in report["spaces"]:
        item = report["inventory"][space]
        nav = navigation["by_space"][space]
        rows.append(
            [
                f"<strong>{esc(space)}</strong>",
                f"{item['pages']:,}",
                f"{item['words']:,}",
                f"{item['assets']:,}",
                f"{nav['entries']:,}",
                str(nav["max_depth"] + 1),
            ]
        )
    rows.append(
        [
            "<strong>Total</strong>",
            f"<strong>{totals['pages']:,}</strong>",
            f"<strong>{totals['words']:,}</strong>",
            f"<strong>{totals['assets']:,}</strong>",
            f"<strong>{navigation['entries']:,}</strong>",
            "",
        ]
    )
    out.append(table(["Space", "Pages", "Words", "Assets", "Menu entries", "Depth"], rows))

    out.append("<h2>Navigation integrity</h2>")
    out.append(
        "<p>SUMMARY.md is the only thing that decides what appears on the site. A page that is not "
        "listed there is invisible however good it is, and an entry with no page behind it is a dead "
        "menu item. Neither shows up in a normal review.</p>"
    )
    rows = []
    for space in report["spaces"]:
        nav = navigation["by_space"][space]
        rows.append(
            [
                f"<strong>{esc(space)}</strong>",
                verdict(nav["orphans"], 1, 25),
                verdict(nav["dangling"], 1, 10),
            ]
        )
    out.append(table(["Space", "Missing from navigation", "Dead entries"], rows, widths=[300, 230, 230]))

    if navigation["orphans"]:
        body = "<ul>" + "".join(f"<li><p>{page_link(path)}</p></li>" for path in navigation["orphans"][:60]) + "</ul>"
        if len(navigation["orphans"]) > 60:
            body += f"<p>… and {len(navigation['orphans']) - 60:,} more.</p>"
        out.append(expand(f"{len(navigation['orphans']):,} pages not listed in any SUMMARY.md", body))

    if navigation["dangling"]:
        rows = [
            [esc(entry["space"]), esc(entry["title"]), f"<code>{esc(entry['target'])}</code>", str(entry["line"])]
            for entry in navigation["dangling"][:60]
        ]
        out.append(
            expand(
                f"{len(navigation['dangling']):,} navigation entries with no page",
                table(["Space", "Menu title", "Target", "SUMMARY.md line"], rows),
            )
        )

    if report["stray"]:
        body = "<ul>" + "".join(f"<li><p>{page_link(path)}</p></li>" for path in report["stray"][:40]) + "</ul>"
        out.append(
            '<div data-type="panel-warning"><p>Markdown sits outside every space, in a directory with no '
            "SUMMARY.md. GitBook cannot render it, so it is dead weight in the repository.</p></div>"
            + expand(f"{len(report['stray']):,} files outside any space", body)
        )

    out.append("<h2>Link health</h2>")
    out.append(
        "<p>Absolute <code>docs.snyk.io</code> URLs between pages do not survive a restructure: when the "
        "target moves, the link rots into a dead end instead of following the page. Relative paths do "
        "follow. Broken relative links point at a file that is no longer in the repository.</p>"
    )
    absolute_by_space = Counter(item["space"] for item in links["absolute"])
    broken_by_space = Counter(item["space"] for item in links["broken"])
    rows = []
    for space in report["spaces"]:
        rows.append(
            [
                f"<strong>{esc(space)}</strong>",
                verdict(absolute_by_space.get(space, 0), 1, 100),
                verdict(broken_by_space.get(space, 0), 1, 25),
            ]
        )
    rows.append(
        [
            "<strong>Total</strong>",
            verdict(len(links["absolute"]), 1, 100),
            verdict(len(links["broken"]), 1, 25),
        ]
    )
    out.append(table(["Space", "Absolute docs.snyk.io links", "Broken relative links"], rows, widths=[300, 230, 230]))

    if links["absolute"]:
        worst = Counter(item["page"] for item in links["absolute"]).most_common(25)
        rows = [[page_link(page), str(count)] for page, count in worst]
        out.append(
            expand(
                f"Pages holding the most absolute links ({len(links['absolute']):,} in total)",
                table(["Page", "Absolute links"], rows, widths=[600, 160]),
            )
        )

    if links["broken"]:
        if links["gitbook_markers"]:
            out.append(
                f"<p>Of the {len(links['broken']):,} broken links, <strong>{links['gitbook_markers']:,}</strong> "
                "are <code>/broken/</code> placeholders that GitBook wrote itself when it could no longer "
                "resolve a target. Those are not guesses: the editor has already recorded the link as dead, "
                "and it renders as a dead end on the site.</p>"
            )
        rows = [
            [
                page_link(item["page"]),
                f"<code>{esc(item['target'])}</code>",
                lozenge("GitBook marker", "red") if item["gitbook_marker"] else lozenge("Missing file", "yellow"),
            ]
            for item in links["broken"][:60]
        ]
        body = table(["Page", "Link target", "Kind"], rows, widths=[350, 280, 130])
        if len(links["broken"]) > 60:
            body += f"<p>… and {len(links['broken']) - 60:,} more.</p>"
        out.append(expand(f"All {len(links['broken']):,} broken links", body))

    out.append("<h2>Attention</h2>")
    method = report["method"]
    out.append(
        '<div data-type="panel-note"><p>Read the dates on this section with one thing in mind. The '
        f"documentation was reorganised into its current spaces on "
        f'<time datetime="{esc(method["tree_rebuilt_at"])}">{esc(method["tree_rebuilt_at"])}</time>, and the '
        "legacy tree was deleted rather than moved, so git holds no history for a page before that date. "
        f"No page can therefore appear older than {method['tree_age_days']:,} days here, and how long a page "
        "went untouched before the move is not something this dashboard can see.</p></div>"
    )
    out.append(
        "<p>Within that window, this is how recently someone deliberately worked on each page. "
        f"Repository-wide sweeps are not counted as edits — {method['bulk_commits_ignored']:,} commits "
        f"touching more than {method['bulk_threshold']} files each were set aside — because a folder move "
        "or a frontmatter pass touches every page without anyone reading one.</p>"
    )
    buckets = fresh["buckets"]
    rows = []
    labels = [
        ("month", "Within a month", "green"),
        ("quarter", "1 to 3 months", "green"),
        ("older", "Over 3 months", "yellow"),
        ("untouched", "Never individually edited", "red"),
    ]
    peak = max((buckets[key] for key, _, _ in labels), default=0)
    for key, label, colour in labels:
        count = buckets[key]
        share = round(count / hand_written * 100) if hand_written else 0
        rows.append([lozenge(label, colour), f"{count:,}", f"{share}%", f"<code>{bar(count, peak)}</code>"])
    out.append(table(["Last edited", "Pages", "Share", ""], rows, widths=[220, 100, 100, 340]))

    out.append(
        f"<p>The last row is the one to look at. Those <strong>{buckets['untouched']:,} pages</strong> have "
        "only ever been touched by bulk changes: they were carried into the new structure and no one has "
        "edited them since. That is not proof any of them is wrong, but it is the set nobody has checked.</p>"
    )

    rows = []
    for space in report["spaces"]:
        counts = fresh["by_space"].get(space, {})
        space_total = sum(counts.values())
        never = counts.get("untouched", 0)
        share = round(never / space_total * 100) if space_total else 0
        rows.append(
            [
                f"<strong>{esc(space)}</strong>",
                f"{space_total:,}",
                f"{counts.get('month', 0):,}",
                f"{counts.get('quarter', 0):,}",
                f"{counts.get('older', 0):,}",
                verdict(never, 1, max(10, space_total // 3)),
                f"{share}%",
            ]
        )
    out.append(
        table(
            [
                "Space",
                "Hand-written pages",
                "Within a month",
                "1 to 3 months",
                "Over 3 months",
                "Never edited",
                "Never edited share",
            ],
            rows,
        )
    )

    untouched = fresh["untouched"][:25]
    if untouched:
        rows = [[page_link(item["page"]), f"{item['words']:,}"] for item in untouched]
        out.append(
            expand(
                "The 25 largest pages nobody has edited since the restructure",
                table(["Page", "Words"], rows, widths=[620, 140]),
            )
        )

    oldest = fresh["oldest"][:25]
    if oldest:
        rows = [
            [page_link(item["page"]), f"{item['days']:,} days", f"{item['words']:,}"]
            for item in oldest
        ]
        out.append(
            expand(
                "The 25 edited pages left longest since",
                table(["Page", "Last edited", "Words"], rows, widths=[520, 140, 100]),
            )
        )

    out.append("<h2>Change activity</h2>")
    totals_activity = activity["totals"]
    combined = totals_activity["gitbook"] + totals_activity["pull_request"]
    gitbook_share = round(totals_activity["gitbook"] / combined * 100) if combined else 0
    out.append(
        f"<p>Commits over the last 12 months, split by route. <strong>{gitbook_share}%</strong> arrived "
        "straight from the GitBook editor, where a writer saves and the change is published without review; "
        "the rest came through a pull request. Both are legitimate, and the ratio shows how the documentation "
        "is actually maintained.</p>"
    )
    months = activity["months"]
    peak = max((month["total"] for month in months), default=0)
    rows = []
    for month in months:
        rows.append(
            [
                f"<strong>{esc(month['month'])}</strong>",
                f"{month['gitbook']:,}",
                f"{month['pull_request']:,}",
                f"{month['total']:,}",
                f"<code>{bar(month['total'], peak)}</code>",
            ]
        )
    rows.append(
        [
            "<strong>Total</strong>",
            f"<strong>{totals_activity['gitbook']:,}</strong>",
            f"<strong>{totals_activity['pull_request']:,}</strong>",
            f"<strong>{combined:,}</strong>",
            "",
        ]
    )
    out.append(table(["Month", "GitBook editor", "Pull request", "Total", ""], rows, widths=[110, 130, 120, 90, 310]))

    contributors = activity["contributors"][:15]
    if contributors:
        peak = contributors[0][1]
        rows = [
            [esc(name), f"{count:,}", f"<code>{bar(count, peak, 18)}</code>"]
            for name, count in contributors
        ]
        out.append(
            expand(
                f"Top contributors over the last 12 months ({len(activity['contributors']):,} people in total)",
                table(["Contributor", "Commits", ""], rows, widths=[260, 100, 400]),
            )
        )

    out.append("<h2>How this page is produced</h2>")
    out.append(
        "<p>A script reads the repository and its git history and writes this page through the Confluence "
        "API. There is no database, no service, and nothing to host: rerun it and the page updates in place. "
        "Anyone who can see this space sees the current figures.</p>"
    )
    out.append(
        '<pre><code class="language-bash">git clone https://github.com/snyk/user-docs\n'
        "cd user-docs\n"
        "python3 tools/docs-dashboard/docs-dashboard.py &gt; dashboard.html</code></pre>"
    )
    out.append(
        "<p>Caveats worth knowing before quoting a number. An edit is a commit, so a typo fix counts the "
        "same as a rewrite, and a page edited once looks the same as a page maintained continuously. Page "
        "history begins at the restructure, as the panel above explains. Word counts skip fenced code and "
        "HTML markup. Broken-link detection covers links between pages in this repository and does not "
        "follow external URLs or check anchors within a page. Contributors are counted by commit, which "
        "rewards small frequent commits over large ones, and several git identities belonging to one "
        "person are merged by name and address, so an unusual pair may still show up twice.</p>"
    )

    return "\n".join(out)


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    parser.add_argument("--repo", default=os.getcwd(), help="repository to measure (default: current directory)")
    parser.add_argument("--format", choices=("html", "json"), default="html", help="output format")
    parser.add_argument("--output", help="write to this file instead of stdout")
    parser.add_argument(
        "--bulk-threshold",
        type=int,
        default=BULK_COMMIT_FILES,
        metavar="N",
        help="treat a commit touching more than N files as a mechanical sweep and ignore it "
        f"when dating pages (default: {BULK_COMMIT_FILES})",
    )
    args = parser.parse_args()

    repo = os.path.abspath(args.repo)
    if not os.path.isdir(os.path.join(repo, ".git")):
        print(f"{repo} is not a git repository", file=sys.stderr)
        return 1

    report = build_report(repo, args.bulk_threshold)
    rendered = json.dumps(report, indent=2) if args.format == "json" else render_html(report)

    if args.output:
        with open(args.output, "w", encoding="utf-8") as handle:
            handle.write(rendered + "\n")
        print(f"Wrote {args.output}", file=sys.stderr)
    else:
        print(rendered)
    return 0


if __name__ == "__main__":
    sys.exit(main())
