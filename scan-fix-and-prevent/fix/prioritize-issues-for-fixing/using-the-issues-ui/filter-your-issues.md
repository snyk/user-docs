---
description: How to filter your issues in the Snyk Issues UI
nav_context: classic
last_reviewed: 2026-11-09
last_reviewed_source: reviewed
---

{% include "../../../.gitbook/includes/new-navigation-banner.md" %}

# Filter your issues

Snyk Issues operates at the Group level and provides a holistic view of all the issues within that Group. Those issues are also tied to specific Organizations. Use the top-level filter to choose which Organizations are relevant to you and see only the issues in those Organizations.

## Table view filters

By using the filters above the table view, you can filter your issues by the following criteria or any combination of them.

* **Issue Status** - view only issues that are open, resolved, or ignored.
* **Asset Name** - view only issues related to specific assets.
* **Asset Class** - view only issues related to a specific asset class.
* **Source Code** - view only issues related to specific source code.
* **Issue Severity** - view only critical, high, medium, or low-severity issues.
* **"Fixed In" Available** - filter issues based on the "Fixed in" version availability.
* **Project Name** - filter issues based on the names of selected Projects.
* **Snyk Product** - view only issues scanned by specific Snyk products.
* **Add filter** - filter issues based on additional options, for example, **CISA KEV**, **CVE**, **CWE**, and **Computed Fixability**. To view all available filters, click **Show all issues filters** or **Show all projects filters**.

You can also add a variety of filters that you consider relevant for any particular scenario.

To view the issues of greatest concern during an initial triage, filter for both **Critical** and **High** severity. Snyk Open Source identifies critical vulnerabilities, and Snyk Code identifies vulnerabilities up to high severity.

Use the following filters to prioritize your results:

* Compliance: Use the **CISA KEV** filter to narrow results to CVEs in the CISA KEV catalog.
* Product: Use the **Snyk Product** filter to separate Snyk Open Source and Snyk Code issues.

If you filter by **Asset Class** and Snyk finds an issue in two repositories with different assigned classes, Snyk displays the class with the highest priority.
