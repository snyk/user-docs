#!/usr/bin/env python3
"""
Fix broken relative links in migrated GitBook sections.

Two types of broken links exist in the migrated sections:
  Type A: Same-space content but ../  escapes the space root → remove the extra ../
  Type B: Cross-space content using relative paths → convert to GitBook absolute URL

Usage:
  python fix_links.py             # fix in-place
  python fix_links.py --dry-run   # print changes without writing
"""

import os
import re
import sys

REPO_ROOT = os.path.abspath(
    os.path.join(os.path.dirname(__file__), '..')
)
ORG_ID = '-M4tdxG8qotLgGZnLpFR'

SECTIONS = [
    'discover-snyk',
    'scan-fix-and-prevent',
    'developer-tools',
    'platform-administration',
    'snyk-data-and-governance',
]

SPACE_IDS = {
    'discover-snyk':         'L7HyJj9FsK1W4pNt8Gzl',
    'scan-fix-and-prevent':  'BJO0IZx7zB6bOkotxQP2',
    'developer-tools':       'IEEjSXQQu36y0vmFV8zf',
    'platform-administration': 'IgtgtomLQ2TUgSKOMSAm',
    'snyk-data-and-governance': 'ELvljsaLKPkSpffOkmsQ',
}

SECTION_ROOTS = {s: os.path.join(REPO_ROOT, s) for s in SECTIONS}

# ── path cache (built once, replaces per-link os.path.exists calls) ──────────

_PATH_CACHE: set[str] = set()

def build_cache():
    for section in SECTIONS:
        root = SECTION_ROOTS[section]
        if not os.path.isdir(root):
            continue
        for dirpath, dirnames, filenames in os.walk(root):
            _PATH_CACHE.add(dirpath)
            for f in filenames:
                _PATH_CACHE.add(os.path.join(dirpath, f))

# ── helpers ──────────────────────────────────────────────────────────────────

def get_source_section(filepath: str) -> str | None:
    rel = os.path.relpath(filepath, REPO_ROOT)
    top = rel.split(os.sep)[0]
    return top if top in SECTIONS else None


def path_exists(p: str) -> bool:
    p = p.rstrip(os.sep)
    return (p in _PATH_CACHE
            or p + '.md' in _PATH_CACHE
            or os.path.join(p, 'README.md') in _PATH_CACHE)


def find_target(resolved_abs: str, source_section: str):
    """
    Return (target_section, within_section_path) for a resolved absolute path.
    Returns (None, None) if we cannot determine the target.
    """
    # ── Case 1: resolved path is directly inside a known section directory ──
    for section in SECTIONS:
        root = SECTION_ROOTS[section]
        if resolved_abs.startswith(root + os.sep) or resolved_abs == root:
            within = os.path.relpath(resolved_abs, root)
            within = within.replace('\\', '/')
            # Verify the content actually lives there
            if path_exists(resolved_abs):
                return section, within
            # Path is within the section dir but doesn't exist → search others
            remainder = within  # e.g. "enterprise-setup/snyk-broker"
            for other in SECTIONS:
                if other == section:
                    continue
                candidate = os.path.join(SECTION_ROOTS[other], remainder)
                if path_exists(candidate):
                    return other, remainder
            # Still not found — return the best guess (section it resolves into)
            return None, None

    # ── Case 2: resolved path escaped all section roots (at repo root level) ──
    rel_to_root = os.path.relpath(resolved_abs, REPO_ROOT).replace('\\', '/')
    parts = rel_to_root.split('/')

    # If the first component is itself a section name (e.g. developer-tools/snyk-cli)
    if parts[0] in SECTIONS:
        section = parts[0]
        within = '/'.join(parts[1:]) if len(parts) > 1 else ''
        if within and path_exists(os.path.join(SECTION_ROOTS[section], within)):
            return section, within
        # Try other sections
        for other in SECTIONS:
            if other == section:
                continue
            candidate = os.path.join(SECTION_ROOTS[other], within)
            if within and path_exists(candidate):
                return other, within
        return None, None

    # Generic search: look for rel_to_root inside each section
    for section in SECTIONS:
        candidate = os.path.join(SECTION_ROOTS[section], rel_to_root)
        if path_exists(candidate):
            return section, rel_to_root

    return None, None


def to_gitbook_url(section: str, within: str, anchor: str = '') -> str:
    space_id = SPACE_IDS[section]
    # Normalise path separators
    path = within.replace('\\', '/')
    # Strip .md extension
    if path.endswith('.md'):
        path = path[:-3]
    # README → directory (drop the /README suffix)
    if path == 'README' or path.endswith('/README'):
        path = path[:-len('README')].rstrip('/')
    # Strip trailing slash
    path = path.rstrip('/')
    return f'https://app.gitbook.com/o/{ORG_ID}/s/{space_id}/{path}{anchor}'


def compute_relative(from_file: str, to_abs: str, had_trailing_slash: bool) -> str:
    from_dir = os.path.dirname(from_file)
    rel = os.path.relpath(to_abs, from_dir).replace('\\', '/')
    if had_trailing_slash and not rel.endswith('/'):
        rel += '/'
    return rel


# ── link fixing ───────────────────────────────────────────────────────────────

# Matches [text](url) and ![alt](url); captures url (group 2).
# Using a simple non-backtracking pattern for the bracket part.
LINK_RE = re.compile(
    r'(!?\[[^\]]*\])'          # [text] — no nested brackets (avoids catastrophic backtracking)
    r'\(([^\s)]+)'              # (url — up to whitespace or )
    r'(?:\s+"[^"]*")?'          # optional "title"
    r'\)'                       # )
)


def fix_link_url(url: str, current_file: str, source_section: str) -> str | None:
    """
    Return a corrected URL, or None if no change is needed.
    """
    # Only process relative links that start with ../
    if not url or not url.startswith('../'):
        return None

    # Split off anchor
    anchor = ''
    if '#' in url:
        idx = url.index('#')
        anchor = url[idx:]
        url_no_anchor = url[:idx]
    else:
        url_no_anchor = url

    had_trailing_slash = url_no_anchor.endswith('/')
    clean = url_no_anchor.rstrip('/')

    # Resolve absolute path
    current_dir = os.path.dirname(current_file)
    resolved = os.path.normpath(os.path.join(current_dir, clean))

    target_section, within = find_target(resolved, source_section)

    if target_section is None:
        return None  # Can't determine target — leave as-is

    if target_section == source_section:
        # ── Type A: same space ──
        target_abs = os.path.join(SECTION_ROOTS[source_section], within)
        new_rel = compute_relative(current_file, target_abs, had_trailing_slash)
        new_url = new_rel + anchor
        return new_url if new_url != url else None
    else:
        # ── Type B: different space → GitBook absolute URL ──
        new_url = to_gitbook_url(target_section, within, anchor)
        return new_url if new_url != url else None


def process_file(filepath: str, dry_run: bool = False) -> int:
    source_section = get_source_section(filepath)
    if source_section is None:
        return 0

    with open(filepath, 'r', encoding='utf-8', errors='replace') as f:
        content = f.read()

    # ── Skip code blocks ──────────────────────────────────────────────────────
    # We process the file line-by-line, replacing links only outside code blocks.
    in_fenced = False
    fence_char = ''
    result_lines = []
    total_fixes = 0

    for line in content.split('\n'):
        stripped = line.strip()
        # Detect fenced code block boundaries (``` or ~~~)
        if not in_fenced:
            if stripped.startswith('```') or stripped.startswith('~~~'):
                fence_char = stripped[:3]
                in_fenced = True
                result_lines.append(line)
                continue
        else:
            if stripped.startswith(fence_char):
                in_fenced = False
                result_lines.append(line)
                continue
            else:
                result_lines.append(line)
                continue

        # Process inline code: preserve backtick spans
        # We only fix links outside backtick spans on the same line.
        # Simple approach: replace links found by LINK_RE.
        def replace_link(m):
            nonlocal total_fixes
            bracket_part = m.group(1)
            url = m.group(2)
            fixed = fix_link_url(url, filepath, source_section)
            if fixed is not None:
                total_fixes += 1
                if dry_run:
                    rel_path = os.path.relpath(filepath, REPO_ROOT)
                    print(f'  {rel_path}: {url!r} → {fixed!r}')
                return f'{bracket_part}({fixed})'
            return m.group(0)

        new_line = LINK_RE.sub(replace_link, line)
        result_lines.append(new_line)

    if total_fixes == 0:
        return 0

    new_content = '\n'.join(result_lines)
    if not dry_run and new_content != content:
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(new_content)

    return total_fixes


def main():
    dry_run = '--dry-run' in sys.argv
    if dry_run:
        print('DRY RUN — no files will be modified\n')
    build_cache()

    total_files = 0
    total_fixes = 0

    for section in SECTIONS:
        section_root = SECTION_ROOTS[section]
        if not os.path.isdir(section_root):
            continue

        section_fixes = 0
        for root, _dirs, files in os.walk(section_root):
            for fname in sorted(files):
                if not fname.endswith('.md'):
                    continue
                filepath = os.path.join(root, fname)
                n = process_file(filepath, dry_run=dry_run)
                if n > 0:
                    total_files += 1
                    section_fixes += n
                    total_fixes += n
                    if not dry_run:
                        print(f'  fixed {n:3d} link(s)  {os.path.relpath(filepath, REPO_ROOT)}')

        if section_fixes:
            print(f'{section}: {section_fixes} fixes')

    print(f'\nTotal: {total_fixes} link(s) fixed across {total_files} file(s)')


if __name__ == '__main__':
    main()
