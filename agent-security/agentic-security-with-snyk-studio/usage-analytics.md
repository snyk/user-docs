---
description: Report types that show how your teams use Snyk Studio
nav_context: agnostic
last_reviewed: 2026-09-16
last_reviewed_source: bulk-sync
---

# Usage analytics

Use Snyk analytics to understand adoption and scanning activity in developer workflows:

* [Overview tab](https://docs.snyk.io/scan-fix-and-prevent/prevent/analytics/overview-tab) for agentic scan and developer-adoption widgets
* [Developer IDE and CLI usage report](https://docs.snyk.io/scan-fix-and-prevent/prevent/analytics/reports-tab/prevention-reports#developer-ide-and-cli-usage-report) for local developer testing

Studio uses two execution paths: background CLI scans started by hooks and direct MCP tool calls from the assistant. These use the integration names `STUDIO` and `MCP`, respectively. Distinguish them when investigating Studio activity rather than treating direct MCP calls as the only Studio interaction.

{% hint style="info" %}
Granular agentic scan data requires Snyk CLI version 1.1298.1 or later. Visit the report documentation for its prerequisites. Analytics version requirements do not establish compatibility with the Studio installer.
{% endhint %}
