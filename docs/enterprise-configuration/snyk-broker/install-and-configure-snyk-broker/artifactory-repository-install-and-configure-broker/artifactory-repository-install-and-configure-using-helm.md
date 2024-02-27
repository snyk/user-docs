# Artifactory Repository - install and configure using Helm

{% hint style="info" %}
**Feature availability**

Integration with Artifactory Repository is available only for Enterprise plans.

For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

For information about non-brokered integration with Artifactory Repository see [Artifactory Repository setup](../../../../integrate-with-snyk/package-repository-integrations/artifactory-package-repository-connection-setup/). For information about brokered integration with Artifactory Container Registry see [Snyk Broker -Container Registry Agent](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent).

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](../install-and-configure-broker-using-helm.md). For more details including definitions of the environment variables see [Artifactory Repository - environment variables for Snyk Broker](artifactory-repository-environment-variables-for-snyk-broker.md).

Note: for `artifactoryUrl` value do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=artifactory \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set artifactoryUrl=<ENTER_ARTIFACTORY_URL> \
             -n snyk-broker --create-namespace
```
