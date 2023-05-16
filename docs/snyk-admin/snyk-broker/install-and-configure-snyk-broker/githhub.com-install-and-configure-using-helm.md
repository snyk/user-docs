# githhub.com - install and configure using Helm

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](install-and-configure-broker-using-helm.md). For more details including definitions of the environment variables see [Set up Snyk Broker with GitHub](broker-example-set-up-snyk-broker-with-github.md).

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```
