# View your Infrastructure as Code issues

* [ View your Infrastructure as Code issues](https://github.com/snyk/user-docs/tree/75cbddc84902693171786610d68edd1dc502bd55/hc/en-us/articles/360017495297-View-your-Infrastructure-as-Code-issues/README.md)

## View your Infrastructure as Code issues

You can view configuration issues in your Infrastructure as Code configuration files, using [Snyk reports](https://support.snyk.io/hc/en-us/sections/360001138198-Reports).

**Feature availability**  
This feature is available with all paid plans. See [Pricing plans](https://snyk.io/plans/) for more details.

Infrastructure as Code configuration issues appear in the summary statistics and graphs by default, showing all of the open issues across your projects and issue types:

To view IaC issues only, select **Configuration** from the **Summary filters** drop down:

For more details on the summary tab, see [Summary tab](https://support.snyk.io/hc/en-us/articles/360004002578-Summary-tab) documentation.

Select the **Issues** page to see detailed information on open issues across all of your projects.

To view IaC issues only, select **Configuration** in the **Issue filters** drop down.

This shows the title of each issue and type, and the severity.

You can also view the issues ungrouped; this shows more information about the project file that the issue is found in, and details on when it was first introduced:

For more details on the issues tab, see [Issues tab](https://support.snyk.io/hc/en-us/articles/360004002598-Issues-tab) documentation.

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

