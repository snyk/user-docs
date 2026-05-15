# Filter your issues

Snyk Issues operates at the Group level and provides a holistic view of all the issues within that Group. Those issues are also tied to specific Organizations. Use the top-level filter to choose which Organizations are relevant to you and see only the issues in those Organizations.

## Funnel view

The funnel view is a visual representation of your application's issues and risk factors. It allows you to filter the list of issues by specific risk factors or a combination of them. The **Deployed** and **Public facing** risk factors are clickable filters.

{% hint style="info" %}
The OS Condition risk factor is now available only in the table view of the Issues UI.
{% endhint %}

## Table view filters

By using the filters above the table view, you can filter your issues by the following criteria or any combination of them.

* **Issue status** - view only issues that are open, resolved, or ignored.
* **Risk factors** - view only issues with specific risk factors.
* **Asset name** - view only issues related to specific assets.
* **Asset class** - view only issues related to a specific asset class.
* **Source code** - view only issues related to specific source code.
* **Issue severity** - view only critical, high, medium, or low-severity issues.
* **"Fixed in" Available** - filter issues based on the "Fixed in" version availability.
* **Project name** - filter issues based on the name(s) of selected Project(s).
* **Snyk Product** - view only issues scanned by specific Snyk products.
* **Add filter** - filter issues based on additional options. Click the **Show all project filters** option to view all available filters, organized by General, Assets, Issues, Projects.

You can also add a variety of filters that you consider relevant for any particular scenario.

To view the issues of greatest concern during an initial triage, filter for both **Critical** and **High** severity. Snyk Open Source identifies critical vulnerabilities, and Snyk Code identifies vulnerabilities up to high severity.

Use the following filters to prioritize your results:

* **Initial triage:** Filter by **OS condition**, **Deployed**, and **Public facing**.
* **Compliance:** Use the **KEV** filter to narrow results to CVEs in the CISA KEV catalog.
* **Product:** Use the **Product** filter to separate Snyk Open Source and Snyk Code issues.

If you filter by **Asset class** and Snyk finds an issue in two repositories with different assigned classes, Snyk displays the class with the highest priority.
