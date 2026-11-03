---
description: Snyk support for Apex with Snyk Code, including supported file formats, interfile analysis, and SCM, CLI, and IDE features
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Apex

{% hint style="info" %}
Snyk supports Apex only for Snyk Code.
{% endhint %}

## Supported file formats

Apex Standard Library is fully supported. Snyk supports the following file formats for Snyk Code: `.cls`, `.trigger`, `.tgr`

## Available features

For Apex, Snyk supports the following features:

* Reports
* Interfile analysis
* SCM import
* CLI and IDE: test or monitor your app. For more information, visit [Snyk CLI for Snyk Code](https://docs.snyk.io/developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code).

{% hint style="info" %}
The **Snyk Fix PR** feature is not available for Apex. This means that Snyk does not notify you if the PR checks fail when the following conditions are met:

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available**.
* **"Fixed in" available** is set to **Yes**.
{% endhint %}
