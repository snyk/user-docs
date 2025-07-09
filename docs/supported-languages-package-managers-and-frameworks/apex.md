# Apex



{% hint style="info" %}
Apex is supported only for Snyk Code.
{% endhint %}

## Applicability

The following functions are available for Apex:

* SCM import
* Test or monitor your app through CLI and IDE. For more information, see [Snyk CLI for Snyk Code](../cli-ide-and-ci-cd-integrations/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/).

## Package managers and supported file extensions

Snyk does not support any package managers for Apex, but supports the following file formats for Snyk Code: `.cls`, `.trigger`, `.tgr`

## Frameworks and libraries

Apex Standard Library is fully supported.

## Features

Snyk supports the following features for Apex:

* Support for Interfile analysis
* Custom rules
* Reports
* Interfile analysis

{% hint style="info" %}
The **Snyk FixPR** feature is not available for Apex. This means that you will not be notified if the PR checks fail when the following conditions are met:&#x20;

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available.**
* "**Fixed in" available** is set to **Yes.**
{% endhint %}
