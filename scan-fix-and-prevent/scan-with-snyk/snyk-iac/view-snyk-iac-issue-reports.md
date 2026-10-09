---
nav_context: classic
description: How to view Snyk IaC issue reports
---

# View Snyk IaC issue reports

{% include "../../.gitbook/includes/new-navigation-banner.md" %}

## View Snyk IaC issue reports

Set the **Issue Type** filter on [Snyk reports](../../prevent/analytics/reports-tab/#snyk-reporting-filters) to **Configuration** to view issues in your IaC configuration files.

<figure><img src="../../.gitbook/assets/Issue type filter - Configuration.png" alt="The Issue type filter set to Configuration."><figcaption><p>The <strong>Issue type</strong> filter set to <strong>Configuration</strong>.</p></figcaption></figure>

You can also use this filter on the **Issues** page to view your IaC issues across all of your Projects.

### API access to IaC issues

You can set the **product\_name** filter for the [Export](https://docs.snyk.io/developer-tools/snyk-api/reference/export) API to `Snyk IaC` to retrieve a list of IaC issues on your Org or Group.
