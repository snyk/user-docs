# Nexus Repository - install and configure using Helm

{% hint style="info" %}
**Feature availability**

Integration with Nexus Repository Manager is available only for Enterprise plans.

For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

For information about non-brokered integration with Nexus Repository Manager including supported environments and versions and user permissions, see [Nexus Repository Manager setup](../../../../scan-using-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/). For information about brokered integration with Artifactory or Nexus Container Registry see [Snyk Broker -Container Registry Agent](../../snyk-broker-container-registry-agent/).

## Nexus 3 Helm installation

Before installing, review the [prerequisites](./) and the general instructions for installation using [Helm](../install-and-configure-broker-using-helm.md).

To use this chart, you must first add the Snyk Broker Helm Chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`&#x20;

Then, run the following commands to install the Broker and customize the environment variables. For definitions of the environment variables, see [Nexus Repository - environment variables for Snyk Broker](nexus-repository-environment-variables-for-snyk-broker.md).

Note: for `baseNexusUrl` and `nexusUrl` values include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=nexus \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set baseNexusUrl=<ENTER_BASE_NEXUS_URL> \
             --set nexusUrl=<ENTER_NEXUS_URL>
             --set brokerClientValidationUrl=<ENTER_BROKER_CLIENT_VALIDATION_URL> \
             -n snyk-broker --create-namespace
```

You can pass any environment variable of your choice in the Helm command. For details, see [Custom additional options for Broker Helm Chart](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md). Follow the instructions for [Advanced configuration for Helm Chart installation](../advanced-configuration-for-helm-chart-installation/) to make configuration changes as needed.

You can verify that the Broker is running by looking at the settings for your brokered integration in the Snyk Web UI to see a confirmation message that you are connected. You can start importing Projects once you are connected.

## Nexus 2 Helm installation

Before installing, review the [prerequisites](./) and the general instructions for installation using [Helm](../install-and-configure-broker-using-helm.md).

To use this chart, you must first add the Snyk Broker Helm Chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`&#x20;

Then run the following commands to customize the environment variables. For definitions of the environment variables, see [Nexus Repository - environment variables for Snyk Broker](nexus-repository-environment-variables-for-snyk-broker.md).

Note: for `baseNexusUrl` and `nexusUrl` values include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=nexus2 \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set baseNexusUrl=<ENTER_BASE_NEXUS_URL> \
             --set nexusUrl=<ENTER_NEXUS_URL>
             --set brokerClientValidationUrl=<ENTER_BROKER_CLIENT_VALIDATION_URL> \
             -n snyk-broker --create-namespace
```

You can pass any environment variable of your choice in the Helm command. For details, see [Custom additional options for Broker Helm Chart](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md). Follow the instructions for [Advanced configuration for Helm Chart installation](../advanced-configuration-for-helm-chart-installation/) to make configuration changes as needed.

You can verify that the Broker is running by looking at the settings for your brokered integration in the Snyk Web UI to see a confirmation message that you are connected. You can start importing Projects once you are connected.
