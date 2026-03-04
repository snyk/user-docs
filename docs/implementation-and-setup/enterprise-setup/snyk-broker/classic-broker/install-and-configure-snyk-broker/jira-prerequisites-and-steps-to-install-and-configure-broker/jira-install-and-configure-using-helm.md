# Jira - install and configure using Helm

Before installing, review the [prerequisites](./) and the general instructions for installation using [Helm](../../../../../../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-helm.md).

{% hint style="info" %}
**Multi-tenant settings for regions**\
When installing, you must add a command in your script to set the `brokerServerUrl`. This is the URL of the Broker server for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
{% endhint %}

## Command to install the Jira integration

To use this chart, you must first add the Snyk Broker Helm Chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

Then, run the following commands to install the Broker and customize the environment variables. For definitions of the environment variables, see [Jira - environment variables for Snyk Broker](jira-environment-variables-for-snyk-broker.md).

Note: for `jiraHostname` value do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=jira \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set brokerServerUrl=<broker-region-url>
             --set jiraUsername=<ENTER_JIRA_USERNAME> \
             --set jiraPassword=<ENTER_JIRA_PASSWORD>  \
             --set jiraHostname=<ENTER_JIRA_HOSTNAME>  \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

You can pass any environment variable of your choice in the Helm command. For details, see [Custom additional options for Broker Helm Chart](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md). Follow the instructions for [Advanced configuration for Helm Chart installation](../advanced-configuration-for-helm-chart-installation/) to make configuration changes as needed.

You can verify that the Broker is running by looking at the settings for your brokered integration in the Snyk Web UI to see a confirmation message that you are connected. You can start importing Projects once you are connected.

## Jira PAT authentication for SSO-enabled JIRA

When SSO is enabled, JIRA usually prohibits the use of a username and password and requires the use of a personal access token (PAT).

When SSO is enabled, you must use a specific Jira version that will instead use the authorization header with the bearer token. To use this version, provide the following configuration:

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=jira-bearer-auth \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set jiraPat=<ENTER_JIRA_PAT> \
             --set jiraHostname=<ENTER_JIRA_HOSTNAME>  \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

{% hint style="info" %}
You must use the Helm chart version 2.2.0 or above.
{% endhint %}
