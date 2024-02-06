# Introduction to Snyk Code custom rules

You can run custom rules as part of any scan run by Snyk Code, and the feature is available when you are using Snyk Code in any of the following:

* [Snyk Web UI](../../../getting-started/explore-snyk-through-the-web-ui.md)
* [Snyk CLI](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/)
* [IDE](../use-snyk-code-in-the-ide.md)

{% hint style="info" %}
**Prerequisite**\
Before using the feature, you must [enable Snyk Code custom rules in Snyk Preview](../../../snyk-admin/manage-settings/snyk-preview.md#enable-or-disable-a-feature).
{% endhint %}

## Using custom rules in the Snyk Web UI

You can create and test queries in a local environment, completely separate from your current Snyk scans. Perform one of the following actions in Snyk Web UI:

* [Run query on a repository](run-query.md#run-query-on-a-repository)
* [Run query on a code snippet](run-query.md#run-query-on-a-code-snippet)
* [Analyze query results](run-query.md#analyze-query-results)
* [Save custom rules](create-snyk-code-custom-rules.md)

## Using custom rules in the Snyk CLI

You can test your Code Projects using the Snyk CLI with regular commands and options as long as you have the [.snyk file](../../policies/the-.snyk-file.md) to hold any custom rules that you have created using Snyk Web UI.

For more information, see [Using Snyk Code via the CLI](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/).

## Using custom rules in the IDE

IDE integrations with Snyk support custom rules as long as you have the [.snyk file](../../policies/the-.snyk-file.md) to hold any custom rules that you have created using Snyk Web UI.&#x20;

For more information, see [Using Snyk Code via the IDE](../use-snyk-code-in-the-ide.md).

## Additional resource

To learn from an end-to-end example, see the Snyk Learn course [Snyk Code Custom Rules](https://learn.snyk.io/lesson/custom-rules-for-snyk-code/).
