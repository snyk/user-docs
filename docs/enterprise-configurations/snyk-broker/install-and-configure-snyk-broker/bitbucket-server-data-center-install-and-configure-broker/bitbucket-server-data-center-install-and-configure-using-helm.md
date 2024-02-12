# Bitbucket Server/Data Center - install and configure using Helm

{% hint style="warning" %}
**Release status**

Snyk Code PR Checks are available only for Bitbucket DC/Server versions 7.0 and above.
{% endhint %}

For instructions on using the Snyk Broker Helm Chart, see [Install and configure Broker using Helm](../install-and-configure-broker-using-helm.md). The command to use follows. Refer to [BitBucket Server/Data Center - environment variables](bitbucket-server-data-center-environment-variables-for-snyk-broker.md) for Snyk Broker for definitions of the environment variables.

&#x20;in the Docker setup instructions.

Note: for `bitbucket` and `bitbucketApi` values do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=bitbucket-server \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set bitbucketUsername=<ENTER_USERNAME> \
             --set bitbucketPassword=<ENTER_PASSWORD> \
             --set bitbucket=<ENTER_BITBUCKET_URL> \
             --set bitbucketApi=<ENTER_BITBUCKET_API_URL> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set enabledAppRisk=true \
             -n snyk-broker --create-namespace
```

{% hint style="info" %}
Snyk AppRisk is set by default to **`false`**. Enable it by setting the flag to **`true`**.
{% endhint %}
