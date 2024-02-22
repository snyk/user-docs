# Legacy reports overview

The legacy **Reports** area offers data and analytics about Projects, issues, dependencies, and licenses. Report information appears based on the Organization in which you are working, and you can filter this data with different parameters.

Additionally, if your account is managed with Groups, aggregated data for all of your organizations is displayed when you navigate to **Reports** from the **Group** level.

From the **Group** level, you can filter to view data for multiple Organizations, similar to this example:

<figure><img src="../../../.gitbook/assets/mceclip0-28-.png" alt="Group-level reporting"><figcaption><p>Group-level reporting</p></figcaption></figure>

You can also use the **Organization Filters** to save and store pre-filtered reports for selected groups of organizations.

Additionally, at the Organization level see [General actions](legacy-reports-general-actions.md) to filter for:

* Project names
* Project types
* vulnerability severity
* a specific period of time

The legacy **Reports tabs** are as follows:

* [Summary](legacy-reports-summary-tab.md): the main dashboard displays a birds-eye view of all of your issues (vulnerabilities and licenses), across all of your Projects.
* [Issues](legacy-reports-issues-tab.md): all issues (vulnerabilities and licenses) across all of your projects, including their severity, any available fixes, and more.
* [Dependencies](dependencies-tab.md): the package dependencies in your project and their health status.
* [Licenses](legacy-reports-licenses-tab.md): the licenses in all of your projects and their status.

Report data can also be generated and retrieved with Snkk APIs. For more information, see the [API v1 documentation](https://snyk.docs.apiary.io/#introduction).

{% hint style="info" %}
There may be a delay from the time a Project is tested when that data appears in the Reports area. If you find that there is more than a 9-hour delay, [contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new).
{% endhint %}

{% hint style="info" %}
Read-only and deactivated Projects and results do not appear in the Reports area.
{% endhint %}
