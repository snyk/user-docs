#!/usr/bin/env python3
"""
rename_images.py — find generically-named images in the user-docs repo, work out
what each one depicts (from its captions/alt text and surrounding docs), and
suggest a descriptive, reusable name for it.

By default the script only *reports*: it writes

    image-rename-report.md     human-readable review of every flagged image
    image-rename-mapping.csv   old_path,old_name,suggested_name,uses,locations

Review/edit the CSV, then re-run with --apply to actually rename the files and
rewrite every reference to them in the Markdown docs.

    python3 scripts/rename_images.py                 # report only (safe)
    python3 scripts/rename_images.py --all           # report ALL images, not just generic ones
    python3 scripts/rename_images.py --vision         # use Claude vision to name context-less images
    python3 scripts/rename_images.py --apply          # rename using image-rename-mapping.csv
    python3 scripts/rename_images.py --apply --csv my-mapping.csv

The base script has no third-party dependencies (Python 3.8+). The optional
--vision step additionally requires `pip install anthropic` and an
ANTHROPIC_API_KEY in the environment; it is imported lazily so the rest of the
script runs without it.
"""
from __future__ import annotations

import argparse
import base64
import csv
import json
import os
import re
import sys
from collections import Counter, defaultdict
from concurrent.futures import ThreadPoolExecutor, as_completed
from pathlib import Path

# ---------------------------------------------------------------------------
# Configuration
# ---------------------------------------------------------------------------

REPO_ROOT = Path(__file__).resolve().parent.parent

IMAGE_EXTS = {".png", ".jpg", ".jpeg", ".gif", ".svg", ".webp", ".bmp"}

REPORT_MD = REPO_ROOT / "image-rename-report.md"
MAPPING_CSV = REPO_ROOT / "image-rename-mapping.csv"

# A filename (stem only, no extension) is considered "generic / non-descriptive"
# if it matches any of these patterns.
GENERIC_PATTERNS = [
    re.compile(r"^\d+$"),                       # 1, 23, 100
    re.compile(r"^\d+[ _-]*\(\d+\)"),           # "1 (1)", "10 (2)"
    re.compile(r"\(\d+\)\s*\(\d+\)"),           # "1 (1) (1)"
    re.compile(r"^image[ _-]?\d*$", re.I),       # image, image1, image_2
    re.compile(r"^screenshot", re.I),            # Screenshot 2025-04-24 at ...
    re.compile(r"^screen[ _-]?shot", re.I),
    re.compile(r"^unnamed", re.I),               # unnamed, unnamed (1)
    re.compile(r"^download", re.I),              # download, download (3)
    re.compile(r"^paste", re.I),                 # pasted images
    re.compile(r"^untitled", re.I),
    re.compile(r"^[0-9a-f]{8,}$", re.I),         # hex / uuid-ish blobs
    re.compile(r"^img[ _-]?\d+$", re.I),         # IMG_1234
    re.compile(r"^capture", re.I),
]

# Words that carry no descriptive value when building a slug from a caption.
STOPWORDS = {
    "the", "a", "an", "of", "in", "on", "to", "for", "and", "or", "with", "at",
    "by", "is", "are", "this", "that", "your", "you", "from", "as", "it", "its",
    "page", "screen", "screenshot", "image", "view", "showing", "shows", "shown",
    "displaying", "displays", "example", "snyk", "here", "above", "below", "see",
    "click", "clicking", "select", "selecting", "where", "which", "when", "into",
    "can", "will", "be", "new", "all",
}

MAX_SLUG_WORDS = 6

# Vision: only raster formats Claude can actually look at (SVG is XML, PDF isn't
# an image block). Maps file extension -> API media_type.
VISION_MEDIA_TYPES = {
    ".png": "image/png",
    ".jpg": "image/jpeg",
    ".jpeg": "image/jpeg",
    ".gif": "image/gif",
    ".webp": "image/webp",
}
VISION_MODEL = "claude-opus-4-8"
VISION_MAX_BYTES = 5 * 1024 * 1024  # API per-image limit

# ---------------------------------------------------------------------------
# Reference + context extraction
# ---------------------------------------------------------------------------

# <figure> ... </figure> blocks (alt + optional figcaption)
RE_FIGURE = re.compile(r"<figure\b.*?</figure>", re.DOTALL | re.IGNORECASE)
RE_IMG_SRC = re.compile(r'<img\b[^>]*?\bsrc="([^"]+)"', re.IGNORECASE)
RE_IMG_ALT = re.compile(r'<img\b[^>]*?\balt="([^"]*)"', re.IGNORECASE)
RE_FIGCAPTION = re.compile(r"<figcaption\b[^>]*>(.*?)</figcaption>", re.DOTALL | re.IGNORECASE)
# markdown image: ![alt](path) OR GitBook's angle-bracket form ![alt](<path with spaces>)
RE_MD_IMG = re.compile(r"!\[([^\]]*)\]\((<[^>]*>|[^)]*)\)")
# any .gitbook/assets/<file> reference inside a non-markdown text file (yaml, html, …)
RE_ASSET_TOKEN = re.compile(
    r"\.gitbook/assets/(<[^>]+>|[^\s\"'`)>\],]+)", re.IGNORECASE)
RE_HEADING = re.compile(r"^#{1,6}\s+(.*)$", re.MULTILINE)
RE_TAGS = re.compile(r"<[^>]+>")

ASSET_RE = re.compile(r"\.gitbook/assets/", re.IGNORECASE)


def section_for(md_path: Path) -> Path:
    """Return the top-level section dir (the one whose .gitbook/assets this md uses)."""
    rel = md_path.relative_to(REPO_ROOT)
    return REPO_ROOT / rel.parts[0]


def asset_basename(src: str) -> str | None:
    """Given an <img src> or markdown target, return the asset basename, or None."""
    if not ASSET_RE.search(src):
        return None
    # strip query string / anchor, then take the part after assets/
    src = src.split("?", 1)[0].split("#", 1)[0]
    after = re.split(r"\.gitbook/assets/", src, flags=re.IGNORECASE)[-1]
    # GitBook wraps space/paren paths in angle brackets: ![alt](<...assets/x (1).png>)
    name = after.strip().strip("<>\"' ")
    return name or None


def clean_text(html: str) -> str:
    return re.sub(r"\s+", " ", RE_TAGS.sub(" ", html)).strip()


def nearest_heading(text: str, pos: int) -> str:
    last = ""
    for m in RE_HEADING.finditer(text):
        if m.start() > pos:
            break
        last = m.group(1).strip()
    return clean_text(last)


def extract_references(md_path: Path):
    """Yield (section, basename, context, md_path) for every image ref in a file."""
    section = section_for(md_path)
    try:
        text = md_path.read_text(encoding="utf-8")
    except (UnicodeDecodeError, OSError):
        return

    seen_spans = []

    # 1. <figure> blocks: richest context (figcaption beats alt)
    for fig in RE_FIGURE.finditer(text):
        block = fig.group(0)
        src_m = RE_IMG_SRC.search(block)
        if not src_m:
            continue
        name = asset_basename(src_m.group(1))
        if not name:
            continue
        cap = RE_FIGCAPTION.search(block)
        alt = RE_IMG_ALT.search(block)
        context = ""
        if cap and clean_text(cap.group(1)):
            context = clean_text(cap.group(1))
        elif alt and alt.group(1).strip():
            context = clean_text(alt.group(1))
        else:
            context = nearest_heading(text, fig.start())
        seen_spans.append(fig.span())
        yield section, name, context, md_path

    # 2. bare <img> not inside a figure we already handled
    for img in RE_IMG_SRC.finditer(text):
        if any(s <= img.start() < e for s, e in seen_spans):
            continue
        name = asset_basename(img.group(1))
        if not name:
            continue
        alt_m = RE_IMG_ALT.search(text[img.start():img.start() + 400])
        context = clean_text(alt_m.group(1)) if alt_m and alt_m.group(1).strip() else ""
        if not context:
            context = nearest_heading(text, img.start())
        yield section, name, context, md_path

    # 3. markdown images
    for md in RE_MD_IMG.finditer(text):
        alt, target = md.group(1), md.group(2)
        name = asset_basename(target)
        if not name:
            continue
        context = clean_text(alt) if alt.strip() else nearest_heading(text, md.start())
        yield section, name, context, md_path


# ---------------------------------------------------------------------------
# Naming
# ---------------------------------------------------------------------------

def is_generic(stem: str) -> bool:
    return any(p.search(stem) for p in GENERIC_PATTERNS)


def slugify(text: str, max_words: int = MAX_SLUG_WORDS) -> str:
    words = re.findall(r"[A-Za-z0-9]+", text.lower())
    kept = [w for w in words if w not in STOPWORDS]
    if not kept:                       # everything was a stopword; fall back
        kept = words
    return "-".join(kept[:max_words])


def suggest_name(contexts: list[str], ext: str, fallback: str) -> str:
    """Build a descriptive kebab-case name from the most common context string."""
    texts = [c for c in contexts if c]
    if texts:
        # most common non-empty caption wins; it best describes the depiction
        dominant = Counter(texts).most_common(1)[0][0]
        slug = slugify(dominant)
    else:
        slug = slugify(fallback)
    if not slug:
        slug = "image"
    return f"{slug}{ext.lower()}"


def caption_relevance(name: str, contexts: list[str]):
    """How well a proposed name reflects the image's captions.

    Returns (score, shared_words). score is the fraction of the name's
    meaningful words that appear in the captions (1.0 = fully caption-derived).
    Returns (None, []) when there is no caption to judge against (orphans /
    vision-named images) — relevancy only applies to referenced images.
    """
    caption_words = set()
    for c in contexts:
        if c:
            caption_words |= {w for w in re.findall(r"[A-Za-z0-9]+", c.lower())
                              if w not in STOPWORDS}
    if not caption_words:
        return None, []
    stem = os.path.splitext(name)[0]
    # ignore pure-digit tokens (e.g. the "-2" dedupe suffix) — not meaningful
    name_words = [w for w in stem.split("-") if w and not w.isdigit()]
    if not name_words:
        return 0.0, []
    shared = [w for w in name_words if w in caption_words]
    return len(shared) / len(name_words), shared


def dedupe(suggestions: dict) -> None:
    """Ensure suggested names are unique within a section (mutates in place)."""
    by_section_name: dict = defaultdict(list)
    for rec in suggestions.values():
        by_section_name[(rec["section"], rec["suggested"])].append(rec)
    for (_section, _name), recs in by_section_name.items():
        if len(recs) <= 1:
            continue
        for i, rec in enumerate(recs[1:], start=2):
            stem, ext = os.path.splitext(rec["suggested"])
            rec["suggested"] = f"{stem}-{i}{ext}"


# ---------------------------------------------------------------------------
# Scan
# ---------------------------------------------------------------------------

# Non-markdown text files that could also reference assets (configs, nav, html).
AUX_EXTS = {".yaml", ".yml", ".json", ".html", ".htm", ".txt", ".toml"}
# Directories we never treat as a reference source (self, generated output, vendored).
SKIP_DIRS = {".git", "node_modules", "scripts"}


def _in_section(path: Path, sections) -> bool:
    if sections is None:
        return True
    rel = path.relative_to(REPO_ROOT)
    return rel.parts and rel.parts[0] in sections


def find_md_files(sections=None):
    for p in REPO_ROOT.rglob("*.md"):
        if SKIP_DIRS & set(p.parts):
            continue
        if _in_section(p, sections):
            yield p


def find_aux_files(sections=None):
    """Non-.md text files (yaml/json/html/…) that might reference assets."""
    for p in REPO_ROOT.rglob("*"):
        if not p.is_file() or p.suffix.lower() not in AUX_EXTS:
            continue
        if SKIP_DIRS & set(p.parts) or ".gitbook" in p.parts:
            continue  # skip assets dir itself and vendored/self dirs
        if _in_section(p, sections):
            yield p


def find_image_files(sections=None):
    for assets in REPO_ROOT.rglob(".gitbook/assets"):
        if not assets.is_dir() or (SKIP_DIRS & set(assets.parts)):
            continue
        if not _in_section(assets, sections):
            continue
        for f in assets.iterdir():
            if f.is_file() and f.suffix.lower() in IMAGE_EXTS:
                yield f


def all_sections() -> set:
    """Top-level section dirs that own a .gitbook/assets folder."""
    out = set()
    for assets in REPO_ROOT.rglob(".gitbook/assets"):
        if assets.is_dir() and not (SKIP_DIRS & set(assets.parts)):
            out.add(assets.relative_to(REPO_ROOT).parts[0])
    return out


def collect_references(sections=None) -> dict:
    """(section, basename) -> list of (context, source_path), across md + aux files."""
    refs: dict = defaultdict(list)
    for md in find_md_files(sections):
        for section, name, context, src in extract_references(md):
            refs[(section, name)].append((context, src))
    for aux in find_aux_files(sections):
        try:
            text = aux.read_text(encoding="utf-8")
        except (UnicodeDecodeError, OSError):
            continue
        section = section_for(aux)
        for m in RE_ASSET_TOKEN.finditer(text):
            name = m.group(1).strip("<>\"' ")
            if name:
                refs[(section, name)].append(("", aux))
    return refs


def scan(include_all: bool, sections=None):
    refs = collect_references(sections)

    records = {}
    for img in find_image_files(sections):
        section = section_for(img)
        stem = img.stem
        if not include_all and not is_generic(stem):
            continue
        key = (section, img.name)
        usages = refs.get(key, [])
        contexts = [c for c, _ in usages]
        locations = sorted({str(p.relative_to(REPO_ROOT)) for _, p in usages})
        page_fallback = " ".join(
            re.findall(r"[A-Za-z0-9]+", img.relative_to(section).as_posix())
        )
        records[str(img.relative_to(REPO_ROOT))] = {
            "section": str(section.relative_to(REPO_ROOT)),
            "path": img,
            "old_name": img.name,
            "suggested": suggest_name(contexts, img.suffix, page_fallback),
            "uses": len(usages),
            "locations": locations,
            "contexts": contexts,
            "vision_description": None,
        }
    return records


# ---------------------------------------------------------------------------
# Vision naming (optional) — name images that have no in-repo caption context
# ---------------------------------------------------------------------------

VISION_PROMPT = (
    "You are naming a screenshot/image from the Snyk product documentation so it "
    "can be reused across pages. Look at the image and return a concise, "
    "descriptive file name that captures what it depicts (the screen, feature, or "
    "action shown) rather than incidental details. Use lowercase kebab-case, 2-6 "
    "words, no file extension, no the/a/of filler, no the word 'snyk' or "
    "'screenshot'. Also return a one-sentence description of what the image shows."
)

VISION_SCHEMA = {
    "type": "object",
    "properties": {
        "name": {"type": "string", "description": "kebab-case file name, no extension"},
        "description": {"type": "string", "description": "one sentence on what is shown"},
    },
    "required": ["name", "description"],
    "additionalProperties": False,
}


def _needs_vision(rec: dict) -> bool:
    """True when an image has no usable in-repo caption/alt context to name from."""
    return not any(c for c in rec["contexts"])


def _name_one_with_vision(client, rec: dict, model: str) -> None:
    """Call Claude vision for a single record; mutate suggested/vision_description."""
    img: Path = rec["path"]
    ext = img.suffix.lower()
    media_type = VISION_MEDIA_TYPES.get(ext)
    if media_type is None:
        return  # SVG/PDF/etc. — nothing to look at
    if img.stat().st_size > VISION_MAX_BYTES:
        return
    data = base64.standard_b64encode(img.read_bytes()).decode("utf-8")

    response = client.messages.create(
        model=model,
        max_tokens=300,
        output_config={"format": {"type": "json_schema", "schema": VISION_SCHEMA}},
        messages=[{
            "role": "user",
            "content": [
                {"type": "image", "source": {"type": "base64", "media_type": media_type, "data": data}},
                {"type": "text", "text": VISION_PROMPT},
            ],
        }],
    )
    if response.stop_reason == "refusal":
        return
    text = next((b.text for b in response.content if b.type == "text"), None)
    if not text:
        return
    parsed = json.loads(text)
    slug = slugify(parsed.get("name", ""))
    if slug:
        rec["suggested"] = f"{slug}{ext}"
        rec["vision_description"] = parsed.get("description") or None


def apply_vision(records: dict, model: str, workers: int) -> None:
    """Fill suggestions for context-less images by looking at them with Claude."""
    try:
        import anthropic  # lazy — only needed with --vision
    except ImportError:
        sys.exit("--vision needs the Anthropic SDK: pip install anthropic")
    if not (os.environ.get("ANTHROPIC_API_KEY") or os.environ.get("ANTHROPIC_AUTH_TOKEN")):
        sys.exit("--vision needs ANTHROPIC_API_KEY (or ANTHROPIC_AUTH_TOKEN) in the environment.")

    client = anthropic.Anthropic()
    targets = [r for r in records.values()
               if _needs_vision(r) and r["path"].suffix.lower() in VISION_MEDIA_TYPES]
    skipped = sum(1 for r in records.values()
                  if _needs_vision(r) and r["path"].suffix.lower() not in VISION_MEDIA_TYPES)
    if skipped:
        print(f"  vision: skipping {skipped} non-raster image(s) (SVG/PDF — can't be analyzed)")
    print(f"  vision: analyzing {len(targets)} context-less image(s) with {model}...")

    done = errors = 0
    with ThreadPoolExecutor(max_workers=workers) as pool:
        futures = {pool.submit(_name_one_with_vision, client, r, model): r for r in targets}
        for fut in as_completed(futures):
            done += 1
            try:
                fut.result()
            except Exception as e:  # keep going; leave the heuristic name in place
                errors += 1
                print(f"    ! {futures[fut]['old_name']}: {e}")
            if done % 25 == 0:
                print(f"    {done}/{len(targets)} analyzed")
    print(f"  vision: done ({done} analyzed, {errors} failed and left with fallback names)")


# ---------------------------------------------------------------------------
# Output
# ---------------------------------------------------------------------------

# Below this caption-overlap fraction, a referenced image's proposed name is
# flagged for manual review (it doesn't clearly reflect its captions).
RELEVANCE_THRESHOLD = 0.5


def write_report(records: dict, report_path: Path = REPORT_MD):
    rows = sorted(records.values(), key=lambda r: (-r["uses"], r["section"], r["old_name"]))
    # caption-relevance for every referenced image (None for orphans/vision)
    for r in rows:
        score, shared = caption_relevance(r["suggested"], r["contexts"])
        r["relevance"], r["relevance_shared"] = score, shared
    low = [r for r in rows
           if r["relevance"] is not None and r["relevance"] < RELEVANCE_THRESHOLD]

    with report_path.open("w", encoding="utf-8") as f:
        f.write("# Image rename report\n\n")
        f.write(f"Flagged **{len(rows)}** image(s) for renaming.\n\n")
        unused = [r for r in rows if r["uses"] == 0]
        f.write(f"- Referenced at least once: **{len(rows) - len(unused)}**\n")
        f.write(f"- Never referenced (orphans): **{len(unused)}**\n")
        f.write(f"- Names with weak caption relevance (review these): **{len(low)}**\n\n")
        f.write("Review the suggested names, edit `image-rename-mapping.csv`, "
                "then run `python3 scripts/rename_images.py --apply`.\n\n")

        if low:
            f.write("## ⚠ Proposed names to double-check (low caption overlap)\n\n")
            f.write("These names don't clearly draw from the image's captions — "
                    "edit `proposed_name` in the CSV if they're off.\n\n")
            for r in low:
                caps = "; ".join(dict.fromkeys(c for c in r["contexts"] if c)) or "—"
                f.write(f"- `{r['old_name']}` → `{r['suggested']}` "
                        f"(relevance {r['relevance']:.0%}) — captions: \"{caps}\"\n")
            f.write("\n")

        for r in rows:
            f.write(f"## `{r['old_name']}`  →  `{r['suggested']}`\n\n")
            f.write(f"- **Section:** {r['section']}\n")
            f.write(f"- **Used:** {r['uses']} time(s)\n")
            if r["relevance"] is not None:
                flag = " ⚠ review" if r["relevance"] < RELEVANCE_THRESHOLD else ""
                f.write(f"- **Caption relevance:** {r['relevance']:.0%}{flag}\n")
            if r["locations"]:
                f.write("- **Appears in:**\n")
                for loc in r["locations"]:
                    f.write(f"  - {loc}\n")
            else:
                f.write("- **Appears in:** _(not referenced — safe to rename or remove)_\n")
            ctx = [c for c in r["contexts"] if c]
            if ctx:
                f.write("- **Depicts (from captions / alt text):**\n")
                for c in dict.fromkeys(ctx):   # unique, order-preserving
                    f.write(f"  - \"{c}\"\n")
            elif r.get("vision_description"):
                f.write(f"- **Depicts (from image analysis):** \"{r['vision_description']}\"\n")
            f.write("\n")
    print(f"  report  -> {report_path}")
    return low


def write_csv(records: dict, csv_path: Path = MAPPING_CSV):
    # n_pages = distinct docs that reference the image (not raw occurrence count)
    rows = sorted(records.values(),
                  key=lambda r: (-len(r["locations"]), r["section"], r["old_name"]))
    with csv_path.open("w", encoding="utf-8", newline="") as f:
        w = csv.writer(f)
        w.writerow([
            "repo_location", "current_name", "proposed_name",
            "used_on_n_pages", "source_captions",
        ])
        for r in rows:
            captions = list(dict.fromkeys(c for c in r["contexts"] if c))  # unique, ordered
            if not captions and r.get("vision_description"):
                captions = [f"(image analysis) {r['vision_description']}"]
            w.writerow([
                str(r["path"].relative_to(REPO_ROOT)),
                r["old_name"],
                r["suggested"],
                len(r["locations"]),
                " | ".join(captions),
            ])
    print(f"  mapping -> {csv_path}")


# ---------------------------------------------------------------------------
# Apply
# ---------------------------------------------------------------------------

def apply_mapping(csv_path: Path, delete_orphans: bool, dry_run: bool = False):
    if not csv_path.exists():
        sys.exit(f"Mapping CSV not found: {csv_path}. Run without --apply first to generate it.")

    # Live re-scan: never trust a possibly-stale CSV for a destructive op.
    live_refs = collect_references()

    renames = []   # (old_path, new_path, section, old_name, new_name)
    deletes = []   # old_path
    with csv_path.open(encoding="utf-8", newline="") as f:
        for row in csv.DictReader(f):
            old_path = REPO_ROOT / row["repo_location"]
            new_name = (row.get("proposed_name") or "").strip()
            old_name = row["current_name"].strip()
            try:
                n_pages = int(row.get("used_on_n_pages") or 0)
            except ValueError:
                n_pages = 0
            if not old_path.exists():
                print(f"  skip (missing file): {row['repo_location']}")
                continue

            section = str(section_for(old_path).relative_to(REPO_ROOT))
            live_count = len(live_refs.get((section, old_name), []))

            if n_pages == 0:
                if not delete_orphans:
                    continue  # orphan, but deletion not requested → leave it
                if live_count > 0:  # SAFETY: CSV said orphan, live scan disagrees
                    print(f"  KEEP (now referenced, not deleting): {row['repo_location']}")
                    continue
                deletes.append(old_path)
                continue

            # Referenced image → rename. Guard against a stale CSV claiming refs
            # that no longer exist (we'd rewrite nothing and orphan the new name).
            if not new_name or new_name == old_name:
                continue
            new_path = old_path.with_name(new_name)
            if new_path.exists():
                print(f"  skip (rename target exists): {new_path.relative_to(REPO_ROOT)}")
                continue
            renames.append((old_path, new_path, section, old_name, new_name))

    if not renames and not deletes:
        print("Nothing to apply.")
        return

    if dry_run:
        print(f"\nDRY RUN — no files changed. Would rename {len(renames)}, delete {len(deletes)}.")
        for old_path, new_path, _s, _o, _n in renames[:40]:
            print(f"  RENAME  {old_path.relative_to(REPO_ROOT)}  ->  {new_path.name}")
        if len(renames) > 40:
            print(f"  ... and {len(renames) - 40} more renames")
        for old_path in deletes[:40]:
            print(f"  DELETE  {old_path.relative_to(REPO_ROOT)}")
        if len(deletes) > 40:
            print(f"  ... and {len(deletes) - 40} more deletions")
        return

    # 1. Rewrite references first (so we never leave dangling links), then move.
    # IMPORTANT: assets are section-scoped. The same generic basename (e.g. 1.png)
    # can exist in several sections with DIFFERENT new names, so rewriting must be
    # scoped to each file's own section — a global by-basename rewrite corrupts
    # references in sibling sections (and parallel trees like docs/ vs the rest).
    renames_by_section: dict = defaultdict(dict)
    for _op, _np, section, old, new in renames:
        renames_by_section[section][old] = new

    changed_files = 0
    if renames_by_section:
        for src in (*find_md_files(), *find_aux_files()):
            section = str(section_for(src).relative_to(REPO_ROOT))
            section_map = renames_by_section.get(section)
            if not section_map:
                continue
            try:
                text = src.read_text(encoding="utf-8")
            except (UnicodeDecodeError, OSError):
                continue
            new_text = text
            for old, new in section_map.items():
                # match only inside an assets/ path, and require a real boundary
                # after the name so "1.png" can't match inside "image (1).png" etc.
                new_text = re.sub(
                    r"(\.gitbook/assets/)" + re.escape(old) + r"(?=[)>\"'\s\],]|$)",
                    lambda m, new=new: m.group(1) + new,
                    new_text,
                )
            if new_text != text:
                src.write_text(new_text, encoding="utf-8")
                changed_files += 1

    for old_path, new_path, _s, _o, _n in renames:
        os.rename(old_path, new_path)

    # 2. Delete confirmed-unreferenced orphans.
    for old_path in deletes:
        old_path.unlink()

    print(f"Renamed {len(renames)} file(s); updated references in {changed_files} doc(s); "
          f"deleted {len(deletes)} orphan(s).")


# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------

def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--apply", action="store_true",
                    help="rename files and rewrite references using the mapping CSV")
    ap.add_argument("--dry-run", action="store_true",
                    help="with --apply, print the rename/delete plan without changing anything")
    ap.add_argument("--all", action="store_true",
                    help="include every image, not only generically-named ones")
    ap.add_argument("--delete-orphans", action="store_true",
                    help="on --apply, delete generically-named images that nothing references "
                         "(re-verified live at apply time)")
    ap.add_argument("--sections", default=None,
                    help="comma-separated top-level dirs to limit the scan to (e.g. for a dry run)")
    ap.add_argument("--exclude-sections", default=None,
                    help="comma-separated top-level dirs to skip (applied after --sections)")
    ap.add_argument("--vision", action="store_true",
                    help="use Claude vision to name images that have no caption/alt context "
                         "(needs `pip install anthropic` + ANTHROPIC_API_KEY)")
    ap.add_argument("--vision-model", default=VISION_MODEL,
                    help=f"model for --vision (default: {VISION_MODEL})")
    ap.add_argument("--vision-workers", type=int, default=8,
                    help="concurrent vision requests (default: 8)")
    ap.add_argument("--csv", default=str(MAPPING_CSV),
                    help="mapping CSV path (default: image-rename-mapping.csv)")
    args = ap.parse_args()

    if args.apply:
        apply_mapping(Path(args.csv), delete_orphans=args.delete_orphans, dry_run=args.dry_run)
        return

    sections = None
    if args.sections:
        sections = {s.strip() for s in args.sections.split(",") if s.strip()}
    if args.exclude_sections:
        excluded = {s.strip() for s in args.exclude_sections.split(",") if s.strip()}
        base = sections if sections is not None else all_sections()
        sections = base - excluded
        print(f"Excluding sections: {', '.join(sorted(excluded))}")
    if sections is not None:
        print(f"Scanning sections: {', '.join(sorted(sections)) or '(none)'}")

    print("Scanning repository for images...")
    records = scan(include_all=args.all, sections=sections)
    n_orphans = sum(1 for r in records.values() if not r["locations"])
    print(f"Flagged {len(records)} image(s) ({n_orphans} orphan(s) with 0 references).")
    if args.vision:
        apply_vision(records, args.vision_model, args.vision_workers)
    dedupe(records)
    csv_path = Path(args.csv)
    # keep the default report name, but for a custom --csv put the report beside it
    report_path = REPORT_MD if csv_path == MAPPING_CSV else csv_path.with_name(csv_path.stem + "-report.md")
    low = write_report(records, report_path)
    write_csv(records, csv_path)
    if low:
        print(f"{len(low)} proposed name(s) have weak caption relevance — "
              f"see the '⚠ Proposed names to double-check' section of the report.")
    print("\nReview the report/CSV, then run:"
          "\n  python3 scripts/rename_images.py --apply --delete-orphans")


if __name__ == "__main__":
    main()
