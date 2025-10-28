# Apex

## Applicability and integration

{% hint style="info" %}
Apex is supported only for Snyk Code.
{% endhint %}

Available integrations:

* SCM import
* CLI and IDE: test or monitor your app. For more information, see [Snyk CLI for Snyk Code](../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/).

## Technical specifications

### Supported file formats

Apex Standard Library is fully supported. Snyk supports the following file formats for Snyk Code: `.cls`, `.trigger`, `.tgr`

### Supported features

Snyk supports the following features for Apex:

* Support for Interfile analysis
* Custom rules
* Reports
* Interfile analysis

{% hint style="info" %}
The **Snyk fix PR** feature is not available for Apex. This means that you will not be notified if the PR checks fail when the following conditions are met:

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available.**
* "**Fixed in" available** is set to **Yes.**
{% endhint %}
