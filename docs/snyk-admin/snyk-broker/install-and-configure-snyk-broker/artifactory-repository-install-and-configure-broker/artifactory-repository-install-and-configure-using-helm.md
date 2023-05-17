# Artifactory Repository - install and configure using Helm

{% hint style="info" %}
For information about brokered integration with Artifactory or Nexus Container Registry see [Snyk Broker -Container Registry Agent](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent).
{% endhint %}

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](../install-and-configure-broker-using-helm.md). For more details including definitions of the environment variables see [Set up Snyk Broker with Artifactory Repository](set-up-snyk-broker-with-artifactory-repository.md).

Note: for `artifactoryUrl` value do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=artifactory \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set artifactoryUrl=<ENTER_ARTIFACTORY_URL> \
             -n snyk-broker --create-namespace
```
