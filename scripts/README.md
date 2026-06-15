# Image cleanup — `rename_images.py`

A maintenance tool for the images under each section's `.gitbook/assets/` folder. It gives screenshots meaningful, reusable names and removes ones that nothing uses — safely, with a review step.

## The problem it solves

GitBook drops uploaded images into `.gitbook/assets/` with whatever name they arrived with. Over time that produces:

- **Non-descriptive names** — `1.png`, `2 (2).png`, `Screenshot 2025-04-24 at 14.52.22.png`, `unnamed.png`, `image (105).png`. You can't tell what an image is without opening it, and you can't reuse it across pages with confidence.
- **Orphans** — images that no page references anymore (left behind when a page was edited or deleted). They bloat the repo and make the assets folder hard to navigate.
- **Risky manual renames** — renaming an asset by hand means hunting down every `.gitbook/assets/…` reference in the Markdown and updating each one, or you silently break an image.

This script fixes all three: it proposes a descriptive name **derived from each image's own caption/alt text**, rewrites every reference for you, and deletes the orphans — and it never does any of that without a review step.

## What it does

1. **Scans** every section's `.gitbook/assets/` for generically-named images.
2. For each one, finds where it's used and reads the **caption / alt text** around each use.
3. **Proposes a descriptive kebab-case name** from the most common caption (e.g. a figure captioned *"Group Settings: SSO"* → `group-settings-sso.png`).
4. Writes a **review report** and a **mapping CSV** you can edit.
5. On `--apply`: renames the files, **rewrites every reference** to them, and (with `--delete-orphans`) deletes unreferenced images.

### Safety features

- **Dry run** (`--dry-run`) prints the full rename/delete plan without changing anything.
- **Live re-verification** — before deleting an "orphan", it re-scans the repo at apply time and refuses to delete anything that turns out to be referenced.
- **Section-scoped reference rewriting** — the same generic name (`1.png`) can exist in several sections with different meanings; rewrites are scoped per section so a reference is never pointed at the wrong section's file.
- **Robust reference parsing** — handles GitBook's angle-bracket link form `![alt](<…/assets/image (1).png>)` and references in non-Markdown files (`.yaml`, `.json`, `.html`).
- **Caption-relevance check** — flags any proposed name that doesn't actually reflect its captions, so you can correct it before applying.

## Requirements

- Python 3.8+ (standard library only).
- Optional: `pip install anthropic` + `ANTHROPIC_API_KEY` — only needed for `--vision` (see below).

## Usage

Run from the repo root.

```bash
# 1. Report only (safe — writes the report + CSV, changes nothing)
python3 scripts/rename_images.py

# 2. Review the outputs
#    image-rename-report.md     — per-image: proposed name, usage, captions, relevance
#    image-rename-mapping.csv   — edit proposed_name here if you disagree with any

# 3. See exactly what applying would do (no changes)
python3 scripts/rename_images.py --apply --delete-orphans --dry-run

# 4. Apply: rename + rewrite references + delete orphans
python3 scripts/rename_images.py --apply --delete-orphans
```

### Useful flags

| Flag | What it does |
|---|---|
| *(none)* | Report mode: writes `image-rename-mapping.csv` + `image-rename-report.md`. Changes nothing. |
| `--apply` | Execute the mapping CSV: rename files and rewrite references. |
| `--delete-orphans` | With `--apply`, also delete images that nothing references (re-verified live). |
| `--dry-run` | With `--apply`, print the plan and change nothing. |
| `--sections a,b` | Limit to these top-level section dirs (e.g. for a smaller, reviewable run). |
| `--exclude-sections a,b` | Skip these sections (applied after `--sections`). |
| `--all` | Include every image, not just generically-named ones. |
| `--csv PATH` | Use a custom mapping CSV path (keeps scoped runs from clobbering the canonical one). |
| `--vision` | Name context-less images (orphans / no caption) by having Claude look at them. Needs `anthropic` + `ANTHROPIC_API_KEY`. |

### Mapping CSV columns

| Column | Meaning |
|---|---|
| `repo_location` | Path to the image file. |
| `current_name` | Its current filename. |
| `proposed_name` | Suggested descriptive name (**edit this** to override). |
| `used_on_n_pages` | Number of distinct pages that reference it (`0` = orphan). |
| `source_captions` | The captions/alt text the name was derived from. |

To **keep** an image the report wants to rename or delete, edit its `proposed_name` (or delete its row) before `--apply`.

## Automated cleanup (GitHub Action)

[`.github/workflows/clean-up-images.yml`](../.github/workflows/clean-up-images.yml) runs this monthly and on manual dispatch, and **opens a PR for review** — it never pushes to `main`.

- **Excludes `docs` by default** (that section is managed separately). Override with the `exclude_sections` input on a manual run.
- Manual runs also accept a `sections` input (limit scope) and a `delete_orphans` toggle.
- The review report is attached to each run as an artifact.

To run the one-time cleanup: trigger the workflow via **Actions → Clean up doc images → Run workflow**.

## How this was built (summary)

The tool was developed iteratively, hardening it against problems found along the way:

1. **Catalog first** — scan the assets, derive names from captions, and emit a human-readable report + CSV (no changes yet).
2. **Add apply** — rename files and rewrite all references; default to a safe report-only mode.
3. **Add orphan deletion** — with a live re-verification so a stale CSV can never cause a wrong delete.
4. **Fix reference parsing** — GitBook's angle-bracket links were being mis-parsed, which had flagged real images as false orphans; fixed, and broadened to non-Markdown files.
5. **Fix section scoping** — a full-repo dry run revealed that rewriting references globally by filename broke links across the repo's parallel section trees; rewriting is now scoped per section.
6. **Add relevancy + dry run + section filters** — verify names match captions, preview the plan, and scope runs to a few folders.
7. **Wrap in a review-gated workflow** — monthly + manual, opening a PR rather than committing to `main`, with `docs` excluded by default.

Throughout, every destructive run was preceded by a dry run and a dangling-reference check (rename → confirm zero broken references → commit).
