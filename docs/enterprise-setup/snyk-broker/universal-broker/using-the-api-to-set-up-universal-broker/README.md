# Using the API to set up Universal Broker

All flows available in the `snyk-broker-config` CLI tool are built on top of the public REST API. The workflows in the CLI tool implement particular flows to provide for ease of setup and use. The same workflows can be accomplished using the API, allowing for automation.

The `snyk-broker-config` tool provides non-workflow commands to simplify the automation of particular tasks through simple bash scripting, avoiding the use of a more complicated API-only path.

## Prerequisites

Before you begin, ensure you have:

* A user personal API key with Tenant admin access and at minimum Org admin access for the Organization on which you want to install the Snyk Broker App, allowing you to complete all steps in the deployment workflow.
* Access to a command-line tool like `curl` and `jq` for JSON parsing (recommended).
* Any necessary pre-configurations for your specific SCM integration.
* Snyk Broker App ID whenever you want to call the API. The Snyk Broker App ID differs for each [region](../../../../working-with-snyk/regional-hosting-and-data-residency.md#broker-client-urls).

{% hint style="info" %}
When you change an environment variable, you must restart your Broker.
{% endhint %}

## Understand the workflow

The [Universal Broker workflow diagrams](universal-broker-workflow-diagrams.md) illustrate the steps that are implemented in the `snyk broker config` tool when you use the commands to automate. The same workflows are implemented when you use the API.

## API configuration for Broker

You can either follow the general API steps to understand the underlying process or go directly to a specific SCM tutorial for detailed API commands. Use the instructions provided under the[Universal Broker API](../../../../snyk-api/reference/universal-broker.md) or specific instructions on how to [Set up an SCM connection using the API](using-the-api-to-set-up-a-github-connection.md).&#x20;

See the [Universal Broker APIs ](../../../../snyk-api/reference/universal-broker.md)in the API Reference for more details. An example is provided: [Using the API to set up a GitHub connection](using-the-api-to-set-up-a-github-connection.md).

{% hint style="info" %}
Use the Snyk Broker App ID whenever you want to call the API. The Snyk Broker App ID differs for each [region](../../../../working-with-snyk/regional-hosting-and-data-residency.md#broker-client-urls).
{% endhint %}

\
