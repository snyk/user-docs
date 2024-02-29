# GitHub Enterprise - install and configure using Helm

Before installing, **review the prerequisites and the general instructions** for installation using [Helm](../install-and-configure-broker-using-helm.md).

To use this chart, you must first add the Snyk Broker Helm Chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`&#x20;

Then, run the following commands to install the Broker and customize the environment variables. Refer to [GitHub Enterprise - environment variables for Snyk Broker](github-enterprise-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

Snyk AppRisk is set by default to `false`. Enable it by setting the flag to `true`.

For `github`, `githubApi` and `githubGraphQl` values do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-enterprise \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set github=<ENTER_GHE_ADDRESS> \
             --set githubApi=<ENTER_GHE_API_ADDRESS> \
             --set githubGraphQl=<ENTER_GRAPHQL_ADDRESS> \
             --set enabledAppRisk=true \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

You can verify that the Broker is running by looking at the settings for your brokered integration in the Snyk Web UI to see a confirmation message that you are connected. You can start importing Projects once you are connected.
