# GitHub Enterprise - install and configure using Helm

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](install-and-configure-broker-using-helm.md). For more details including definitions of the environment variables see [Set up Snyk Broker with GitHub Enterprise](setup-broker-with-github-enterprise.md).

Note: for `github`, `githubApi` and `githubGraphQl` values do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-enterprise \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set github=<ENTER_GHE_ADDRESS> \
             --set githubApi=<ENTER_GHE_API_ADDRESS> \
             --set githubGraphQl=<ENTER_GRAPHQL_ADDRESS> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```
