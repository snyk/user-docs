# Azure Repos - install and configure and configure using Helm

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](../install-and-configure-broker-using-helm.md). The command to use follows. Refer to [Azure Repos - environment variables for Snyk Broker](azure-repos-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

Note: for `azureReposHost` value do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=azure-repos \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set azureReposToken=<ENTER_REPO_TOKEN> \
             --set azureReposOrg=<ENTER_REPO_ORG> \
             --set azureReposHost=<ENTER_REPO_HOST> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set enabledAppRisk=true \
             -n snyk-broker --create-namespace
```

{% hint style="info" %}
Snyk AppRisk is set by default to **`false`**. Enable it by setting the flag to **`true`**.
{% endhint %}
