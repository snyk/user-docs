---
nav_context: agnostic
description: >-
  How to review Snyk reports at the Organization and Group level for insights
  into vulnerabilities across Projects
---

# Review reporting

{% include "../../../.gitbook/includes/pilot-guide-navigation.md" %}

The Snyk platform has a variety of [available reports](https://docs.snyk.io/scan-fix-and-prevent/prevent/analytics/reports-tab) at the Organization and Group level to provide deeper insights. This page covers a few of the key reports and capabilities.

## Issues Detail

The Issues Detail report displays all known issues in all of your Projects that Snyk is monitoring. Use it to review and prioritize issues across all of your repositories.

Apply filters such as fixability, exploit maturity, asset class, and severity to focus on the most important issues that affect your Organization.

## SLA Management

Use the SLA Management report to track the service level agreements (SLAs) your company must meet. You can adjust the SLA time for each severity and save it as a new view to monitor later.

## Analytics

You can access [Analytics](https://docs.snyk.io/scan-fix-and-prevent/prevent/analytics) at the Group or Organization level. This view breaks down the total issue count into baseline, preventable, and non-preventable issues:

* Baseline issues are the issues identified during the first import of the repository.
* Preventable issues are known issues that could have been identified earlier in the software development lifecycle (SDLC). In other words, a Snyk PR check could have caught the issue.
* Non-preventable issues are the result of an external factor, not of developers failing to shift left. Examples include a newly published vulnerability, such as a new zero-day vulnerability, or a newly created security rule.

These are only some of the available reports. To explore more Snyk reports at the Organization and Group levels, click **Change Report**.

{% hint style="info" %}
Additional resources

* [Docs: Analytics](https://docs.snyk.io/scan-fix-and-prevent/prevent/analytics)
* [Docs: Reports](https://docs.snyk.io/scan-fix-and-prevent/prevent/analytics/reports-tab)
* [Product Training: Reports](https://learn.snyk.io/catalog/?type=product-training\&topics=Reporting)
{% endhint %}
