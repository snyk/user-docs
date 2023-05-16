# GitLab - install and configure using Helm

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](install-and-configure-broker-using-helm.md). For more details including permissions and definitions of the environment variables see [Set up Snyk Broker with GitLab](setup-broker-with-gitlab.md).

Note: for `gitlab` value do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=gitlab \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set gitlab=<ENTER_GITLAB_URL> \
             --set scmToken=<ENTER_GITLAB_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```
