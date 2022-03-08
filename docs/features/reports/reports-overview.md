# Reports overview

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise and Business plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

The **Reports** area offers data and analytics across all of your projects, displaying historical and aggregated data about projects, issues, dependencies, and licenses. Data in each of the four tabs (seen below) is displayed based on the organization in which you are working, and you can filter this data with different parameters depending on the tab you're viewing.

Additionally, if your account is managed with groups, aggregated data for all of your organizations is displayed when you navigate to **Reports** from the **Group** level.

From the **Group** level, you can filter to view data for multiple organizations, similar to this example:

![](../../.gitbook/assets/mceclip0-28-.png)

You can also use the **Organization Filters** to save and store pre-filtered reports for selected groups of organizations. For more information about this, see our [Snyk groups overview](../user-and-group-management/managing-groups-and-organizations/snyk-groups-overview.md).

Additionally, at the Organization level see [General actions](general-actions.md) to filter for:

* project names
* project types
* vulnerability severity
* a specific period of time

The Reports area comprises these tabs:

* [Summary](summary-tab.md)—the main dashboard displays a birds-eye view of all of your issues (vulnerabilities and licenses), across all of your projects.
* [Issues](issues-tab.md)—all issues (vulnerabilities and licenses) across all of your projects, including their severity, any available fixes, and more.
* [Dependencies](dependencies-tab.md)—the package dependencies in your project and their health status.
* [Licenses](licenses-tab.md)—the licenses in all of your projects and their status.

Report data can also be generated and retrieved with our APIs. For more information about this, see our [API documentation](https://snyk.docs.apiary.io/#introduction).

{% hint style="warning" %}
There may be a delay from the time a project is tested and until that data appears in the Reports area. If you find that there is more than a 9-hour delay, [contact our Support team](https://support.snyk.io/hc/en-us/requests/new).
{% endhint %}

{% hint style="danger" %}
Read-only projects and their respective results will not appear in the Reports area.
{% endhint %}
