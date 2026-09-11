# Documentation health dashboard

Generates a health report for this repository — inventory, navigation integrity,
link health, and which pages nobody has edited — and renders it as a Confluence
page body.

The point of publishing to Confluence rather than building a site is that
nothing has to be hosted. The numbers come from the repository, the page lives
in the Docs space, and anyone in the organization can read it. There is no
service to keep running and nothing to pay for.

The published page is **Documentation health dashboard** in the `Docs` space.

## Running it

Needs Python 3.9 or later and nothing else. Run it from the repository root:

```bash
python3 tools/docs-dashboard/docs-dashboard.py --output dashboard.html
```

`--format json` gives the same figures as data, which is the easier form to
diff between runs or to chart elsewhere:

```bash
python3 tools/docs-dashboard/docs-dashboard.py --format json --output dashboard.json
```

`--bulk-threshold N` controls which commits count as edits. A commit touching
more than `N` files is treated as a mechanical sweep — a folder move, a
frontmatter pass — and ignored when working out when a page was last worked on.
The default of 200 keeps the three large restructuring commits from reporting
every page as freshly written. Lower it to be stricter about what counts as an
edit; raise it if a genuine bulk rewrite is being missed.

## Refreshing the published page

The HTML is the Confluence dialect that maps onto ADF, not storage format, so
it is accepted by the Confluence page API as a page body but will not render if
pasted into the editor as source.

To update the page, ask an agent with the Atlassian MCP connected to run the
script and overwrite the body of **Documentation health dashboard** in the
`Docs` space. Nothing else needs to change: the page ID stays the same, and
readers keep their link.

## What the numbers do and do not mean

Read the caveats at the foot of the published page before quoting a figure. The
two that matter most:

- Page history begins at the May 2026 restructure. The legacy `docs/` tree was
  deleted rather than moved, so git holds no link across it and no page can
  appear older than that here.
- An edit is a commit, so a typo fix counts the same as a rewrite.
