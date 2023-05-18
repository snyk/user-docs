# githhub.com - install and configure using Helm

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](../install-and-configure-broker-using-helm.md). The command to use follows. Refer to [GitHub - environment variables for Snyk Broker](github-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```
