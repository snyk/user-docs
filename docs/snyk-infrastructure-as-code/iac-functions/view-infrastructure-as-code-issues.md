# View Infrastructure as Code issues

You can view configuration issues in your Infrastructure as Code configuration files, using [Snyk reports](https://support.snyk.io/hc/en-us/sections/360001138198-Reports).

**Feature availability**  
This feature is available with all paid plans. See [Pricing plans](https://snyk.io/plans/) for more details.

Infrastructure as Code configuration issues appear in the summary statistics and graphs by default, showing all of the open issues across your projects and issue types:

![image4.png](https://support.snyk.io/hc/article_attachments/360016579057/image4.png)

To view IaC issues only, select **Configuration** from the **Summary filters** drop down:

![Screenshot\_2021-02-17\_at\_14.22.50.png](https://support.snyk.io/hc/article_attachments/360017012177/Screenshot_2021-02-17_at_14.22.50.png)

For more details on the summary tab, see [Summary tab](https://docs.snyk.io/reports-1/reports/summary-tab) documentation.

Select the **Issues** page to see detailed information on open issues across all of your projects.

To view IaC issues only, select **Configuration** in the **Issue filters** drop down.

![image3.png](https://support.snyk.io/hc/article_attachments/360016579157/image3.png)

This shows the title of each issue and type, and the severity.

You can also view the issues ungrouped; this shows more information about the project file that the issue is found in, and details on when it was first introduced:

![image2.png](https://support.snyk.io/hc/article_attachments/360016644818/image2.png)

For more details on the issues tab, see [Issues tab](https://docs.snyk.io/reports-1/reports/issues-tab) documentation.

Issues can be exported as a CSV file in the same format as your vulnerabilities, using the **Export** button.

You can access the full list of issues via the API using the [latest issues endpoint](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues?console=1).

To retrieve only your Infrastructure as Code issues, submit a body payload:

```text
{
  "filters": {
    "orgs": ["my-public-org-id"],
    "types": [
      "configuration"
    ]
  }
}
```

You can obtain your **public-org-id** from the Snyk UI **Settings** page when you view your target organization.

For the full list of parameters, see the [API Documentation](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues?console=1).

