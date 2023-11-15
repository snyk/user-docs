# GitLab - install and configure using Helm

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](../install-and-configure-broker-using-helm.md). The command to use follows. Refer to [GitLab - environment variables for Snyk Broker](gitlab-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

{% hint style="info" %}
Notes for installation using Helm

For the `gitlab` value, do not include `https://.`

`brokerToken` is a Helm variable that feeds into and sets the `BROKER_TOKEN` environment variable. It is the Helm way to pass the token.

`ACCEPT_CODE` is [set to true by default](https://github.com/snyk/snyk-broker-helm/blob/465d4ef279755fa5c9507975a88348bab04c2264/charts/snyk-broker/templates/broker\_deployment.yaml#L383) in the chart, as is [ACCEPT\_IAC](https://github.com/snyk/snyk-broker-helm/blob/465d4ef279755fa5c9507975a88348bab04c2264/charts/snyk-broker/templates/broker\_deployment.yaml#L386C23-L386C43). You can disable them if needed using `disableAutoAcceptRules=true`, but otherwise, these are enabled.
{% endhint %}

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=gitlab \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set gitlab=<ENTER_GITLAB_URL> \
             --set scmToken=<ENTER_GITLAB_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```
