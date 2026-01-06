# Filter your issues

Snyk Issues operates at the Group level and provides a holistic view of all the issues within that Group. Those issues are also tied to specific Organizations. Use the top-level filter to choose which Organizations are relevant to you and see only the issues in those Organizations.

## Funnel View

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

In an initial triage, you want to look at OS condition, Deployed, and Public facing, and choose both critical and high severity. Snyk Open Source finds critical vulnerabilities, while Snyk Code finds vulnerabilities up to high severity only. When you filter on the issues of greatest concern, filter for both critical and high severity.

When you filter based on the Asset class and an issue is found in two different repositories with two different classes assigned, the class with the highest priority is displayed in Snyk.

When you want to split out Open Source and Snyk Code issues, use a product filter.

<div align="left"><figure><img src="../../../.gitbook/assets/image (5).png" alt="Add filters option from the Issues page" width="303"><figcaption><p>Add filters option from the Issues page</p></figcaption></figure></div>
