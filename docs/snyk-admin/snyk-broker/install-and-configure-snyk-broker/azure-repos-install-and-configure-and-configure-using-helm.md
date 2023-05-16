# Azure Repos - install and configure and configure using Helm

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](install-and-configure-broker-using-helm.md). For more details including definitions of the environment variables see [Set up Snyk Broker with Azure Repos](setup-broker-with-azure-repos.md).

Note: for `azureReposHost` value do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=azure-repos \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set azureReposToken=<ENTER_REPO_TOKEN> \
             --set azureReposOrg=<ENTER_REPO_ORG> \
             --set azureReposHost=<ENTER_REPO_HOST> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```
