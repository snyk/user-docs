# View Infrastructure as Code issues

You can view configuration issues in your Infrastructure as Code configuration files, using [Snyk reports](https://docs.snyk.io/reports/reports).

{% hint style="info" %}
**Feature availability**\
This feature is available with all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

## Summary Page

Infrastructure as Code configuration issues appear in the summary statistics and graphs by default, showing all of the open issues across your projects and issue types:

![](../../../.gitbook/assets/image4.png)

To view IaC issues only, select **Configuration** from the **Summary filters** drop down:

![](../../../.gitbook/assets/screenshot\_2021-02-17\_at\_14.22.50.png)

For more details on the summary tab, see [Summary tab](https://docs.snyk.io/reports-1/reports/summary-tab) documentation.

## Issues page

Select the **Issues** page to see detailed information on open issues across all of your projects.

To view IaC issues only, select **Configuration** in the **Issue filters** drop down.

![](../../../.gitbook/assets/image3.png)

This shows the title of each issue and type, and the severity.

You can also view the issues ungrouped; this shows more information about the project file that the issue is found in, and details on when it was first introduced:

![](<../../../.gitbook/assets/image2-3- (1) (2) (2) (2) (3) (4) (4) (3) (1) (1) (10).png>)

For more details on the issues tab, see [Issues tab](https://docs.snyk.io/reports-1/reports/issues-tab) documentation.

## Export data

Issues can be exported as a CSV file in the same format as your vulnerabilities, using the **Export** button.

## API access

You can access the full list of issues via the API using the [latest issues endpoint](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues?console=1).

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
You can obtain your **public-org-id** from the Snyk UI **Settings** page when you view your target organization.
{% endhint %}

For the full list of parameters, see the [API Documentation](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues?console=1).
