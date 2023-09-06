# Getting started with custom rules



{% hint style="info" %}
**Feature availability**\
The Snyk Code custom rules feature is in [Open Beta](../../../more-info/snyk-feature-release-process.md#open-beta) and available only for Enterprise plans through Snyk Preview.&#x20;
{% endhint %}

{% hint style="warning" %}
Custom rules beta is not currently compatible with [SCM integrations via Broker](../../../enterprise-setup/snyk-broker/#integrations-with-snyk-broker).
{% endhint %}

You can run custom rules as part of any scan run by Snyk Code, and the feature is available when you are using Snyk Code in any of the following:

* [Snyk Web UI](../../../getting-started/exploring-the-snyk-web-ui.md)
* [Snyk CLI](../using-snyk-code-from-the-cli/)
* [IDE](../using-snyk-code-in-an-ide.md)

{% hint style="info" %}
**Prerequisite**\
Before using the feature, you must [enable Snyk Code custom rules in Snyk Preview](../../../snyk-admin/manage-settings/snyk-preview.md#enable-or-disable-a-feature).
{% endhint %}

## Using custom rules in the Snyk Web UI

You can create and test queries in a local environment, completely separate from your current Snyk scans. Perform one of the following actions in Snyk Web UI:

* [Run query on a repository](run-query.md#run-query-on-a-repository)
* [Run query on a code snippet](run-query.md#run-query-on-a-code-snippet)
* [Analyze query results](run-query.md#analyze-query-results)
* [Save custom rules](create-custom-rules.md)

## Using custom rules in the Snyk CLI

You can test your Code Projects using the Snyk CLI with regular commands and options as long as you have the [.snyk file](../../../manage-risk/policies/the-.snyk-file.md) to hold any custom rules that you have created using Snyk Web UI.

For more information, see [Using Snyk Code via the CLI](../using-snyk-code-from-the-cli/).

## Using custom rules in the IDE

IDE integrations with Snyk support custom rules as long as you have the [.snyk file](../../../manage-risk/policies/the-.snyk-file.md) to hold any custom rules that you have created using Snyk Web UI.&#x20;

For more information, see [Using Snyk Code via the IDE](../using-snyk-code-in-an-ide.md).

## Additional resource

To learn from an end-to-end example, see the Snyk Training course [Snyk Code Custom Rules](https://training.snyk.io/learn/course/snyk-code-custom-rules/main/snyk-code-custom-rules).
