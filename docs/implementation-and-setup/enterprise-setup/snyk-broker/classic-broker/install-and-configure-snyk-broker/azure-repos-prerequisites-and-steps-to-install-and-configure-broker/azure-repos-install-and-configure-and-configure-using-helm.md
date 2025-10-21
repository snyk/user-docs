# Azure Repos - install and configure and configure using Helm

{% hint style="info" %}
**Feature availability**

Snyk Azure Repos are available only for Azure DevOps/TFS 2020 or above.
{% endhint %}

Before installing, review the [prerequisites](./) and the general instructions for installation using [Helm](../install-and-configure-broker-using-helm.md).

To use this chart, you must first add the Snyk Broker Helm Chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`&#x20;

Then, run the following commands to install the Broker and customize the environment variables. Refer to [Azure Repos - environment variables for Snyk Broker](azure-repos-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

For the variable `azureReposHost` the value does not include `https://`. If you have more than one Azure organization, you must deploy a Broker for each one.

Snyk Essentials is set by default to `false`. Enable it by setting the flag to `true`.

{% hint style="info" %}
**Multi-tenant settings for regions**\
When installing, you must add a command in your script to set the `brokerServerUrl`. This is the URL of the Broker server for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
{% endhint %}

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=azure-repos \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set brokerServerUrl=<broker-region-url>
             --set azureReposToken=<ENTER_REPO_TOKEN> \
             --set azureReposOrg=<ENTER_REPO_ORG> \
             --set azureReposHost=<ENTER_REPO_HOST> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set enableEssentials=true \
             -n snyk-broker --create-namespace
```

You can pass any environment variable of your choice in the Helm command. For details, see [Custom additional options for Broker Helm Chart](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md). Follow the instructions for [Advanced configuration for Helm Chart installation](../advanced-configuration-for-helm-chart-installation/) to make configuration changes as needed.

You can verify that the Broker is running by looking at the settings for your brokered integration in the Snyk Web UI to see a confirmation message that you are connected. You can start importing Projects once you are connected.
