# Nexus Repository - install and configure using Helm

{% hint style="info" %}
For information about brokered integration with Artifactory or Nexus Container Registry see [Snyk Broker -Container Registry Agent](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent).
{% endhint %}

## Nexus 3 Helm installation

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](../install-and-configure-broker-using-helm.md). The command to use follows. For definitions of the environment variables, see [Nexus Repository - environment variables for Snyk Broker](nexus-repository-environment-variables-for-snyk-broker.md).

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

## Nexus 2 Helm installation

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](../install-and-configure-broker-using-helm.md). The command to use follows. For definitions of the environment variables, see [Nexus Repository - environment variables for Snyk Broker](nexus-repository-environment-variables-for-snyk-broker.md).

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
