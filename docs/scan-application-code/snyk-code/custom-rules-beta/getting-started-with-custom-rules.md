# Getting started with custom rules

## Enable Snyk Code custom rules in Snyk Preview

Enabling custom rules is only available for Enterprise plans. You can [enable Snyk Code custom rules in Snyk Preview](../../../snyk-admin/manage-settings/snyk-preview.md#enable-or-disable-a-feature).

{% hint style="warning" %}
Custom rules beta is not currently compatible with [SCM integrations via Broker](../../../enterprise-setup/snyk-broker/#integrations-with-snyk-broker).
{% endhint %}

## Getting started with custom rules across environments

You can run custom rules as part of any scan run by Snyk Code, and it is available across the following environments:

* [Snyk Web UI](../../../getting-started/getting-started-with-the-snyk-web-ui.md)
* [Snyk CLI](../cli-for-snyk-code/)
* [IDE](../using-snyk-code-via-ide.md)

### Snyk Web UI

You can create and test queries in a local environment, completely separate from your current Snyk scans. Perform one of the following actions in Snyk Web UI:

* [Run query on a repository](run-query.md#run-query-on-a-repository)
* [Run query on a code snippet](run-query.md#run-query-on-a-code-snippet)
* [Analyze query results](run-query.md#analyze-query-results)
* [Save custom rules](create-custom-rules.md)

### Snyk CLI

You can test your Code Projects using Snyk CLI with regular commands and options as long as you have the [.snyk file](../../../snyk-cli/test-for-vulnerabilities/the-.snyk-file.md) to hold any custom rules that you have created using Snyk Web UI.

:link: [Using Snyk Code via the CLI](../cli-for-snyk-code/)

### IDE

IDE integrations with Snyk support custom rules as long as you have the [.snyk file](../../../snyk-cli/test-for-vulnerabilities/the-.snyk-file.md) to hold any custom rules that you have created using Snyk Web UI.&#x20;

:link: [Using Snyk Code via the IDE](../using-snyk-code-via-ide.md)

## What's next?

To learn from an end-to-end example, see the following Snyk Training course: [Snyk Code Custom Rules](https://training.snyk.io/learn/course/snyk-code-custom-rules/main/snyk-code-custom-rules).
