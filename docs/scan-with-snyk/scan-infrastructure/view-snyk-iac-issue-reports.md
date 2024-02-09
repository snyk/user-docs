# View Snyk IaC issue reports

Use [Snyk reports](../../manage-issues/reporting/legacy-reports/) to view issues in your IaC configuration files.

{% hint style="warning" %}
**Release status**\
IaC issue reports are available for all paid plans.&#x20;

See [Pricing plans](https://snyk.io/plans).
{% endhint %}

## Summary page

Infrastructure as Code configuration issues appear in the summary statistics and graphs by default, showing all of the open issues across your Projects and issue types:

<figure><img src="../../.gitbook/assets/image4.png" alt="Reports summary page"><figcaption><p>Reports summary page</p></figcaption></figure>

To view IaC issues only, select **Configuration** from the **Summary filters** drop down:

![](../../.gitbook/assets/screenshot\_2021-02-17\_at\_14.22.50.png)

For more details on the summary tab, see the [Reports Summary tab](../../manage-issues/reporting/legacy-reports/legacy-reports-summary-tab.md) documentation.

## Issues page

Select the **Issues** page to see detailed information on open issues across all of your Projects.

To view IaC issues only, select **Configuration** from the **Issue filters** drop-down.

<figure><img src="../../.gitbook/assets/image3 (2).png" alt="IaC issues page"><figcaption><p>IaC issues page</p></figcaption></figure>

This shows the title of each issue, the type, and the severity.

You can also view the issues ungrouped; this shows more information about the Project file where the issue is found, and details on when it was first introduced:

![](<../../.gitbook/assets/image2-3 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (5) (7).png>)

For more details on the issues tab, see [Issues tab](../../manage-issues/reporting/legacy-reports/legacy-reports-issues-tab.md) documentation.

### Export data

Use the Export button to export issues as a CSV file in the same format as your vulnerabilities.

### API access

You can access the full list of issues via the API using the API v1 [latest issues endpoint](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues?console=1).

To retrieve only your Infrastructure as Code issues, submit a body payload:

```
{
  "filters": {
    "orgs": ["my-public-org-id"],
    "types": [
      "configuration"
    ]
  }
}
```

{% hint style="info" %}
You can obtain your **public-org-id** from the Snyk UI **Settings** page when you view your target Organization.
{% endhint %}

For the full list of parameters, see the [Latest issues API documentation](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues?console=1).
