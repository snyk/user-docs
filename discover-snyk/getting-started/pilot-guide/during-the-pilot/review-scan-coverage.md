---
description: How to review Snyk scan coverage, including the Group-level SCM integration and the Inventory
nav_context: classic
---

{% include "../../../.gitbook/includes/new-navigation-banner.md" %}

# Review scan coverage

{% include "../../../.gitbook/includes/pilot-guide-navigation.md" %}

## Check ‌the Group-level SCM integration

Ensure that the Group-level SCM integration is configured. Navigate to the Group level, open the Integrations tab, and find your SCM in the Integration Hub. Follow the instructions for your SCM.

## Review the Inventory

The first page of the Inventory provides an overview of your most important repositories and identifies coverage gaps, showing which repos have been tested with Snyk and which have not. See the [Inventory menu](https://docs.snyk.io/scan-fix-and-prevent/fix/manage-assets) page for more details.

The **All Assets** page shows a complete list of all repositories, including the number of issues for each repo, which Snyk tests have run, tags ingested from the SCM, developers who contribute to this repo, and the repository freshness.

Click on the 'Not tested' section of the first pie chart on the overview page, or use coverage filters on the 'All Assets' page to view all repositories that the selected Snyk product has not tested.

## Classifying Assets

The Class column is available for each repository. This class is meant to reflect the business criticality of the asset from A (most critical) to D (least critical). Try setting a few of your most important repos manually to Class A. This attribute can be used in reporting to help focus on issues from your company’s most important repositories.

\\

While you can set the class manually, you can also create a policy to automatically classify the repository. For example, Snyk ingests metadata like GitHub Topics and Custom Properties that can be used in a policy.

For example, a policy can filter on all repositories where there is a compliance tag, which in this case comes from the GitHub Custom Property, and set the asset class to ‘A’.

The asset class is a great way to filter on the most critical repositories in your organization, and help take a risk-based approach to prioritization by accounting for business impact. The asset class can be used when reviewing scan coverage and in all available Snyk reports.

{% hint style="info" %}
Additional Resources

* [Manage assets](https://docs.snyk.io/scan-fix-and-prevent/fix/manage-assets)
* [Product Training: Snyk Essentials](https://learn.snyk.io/catalog/?type=product-training\&topics=Snyk+Essentials)
{% endhint %}
