# Deploying multiple Brokers in the same namespace

## Helm chart versions 2.x.x

Helm chart versions 2.x.x automatically adds a suffix to all object name based on the release name. Simply use different release name for each broker you want to add into the same namespace.

{% hint style="info" %}
For backward compatibility, 2.1.0 introduces a disableSuffixes flag to revert to the 1.x.x behavior listed below, by simply adding --set disableSuffixes=true or disableSuffixes=true in your values file.
{% endhint %}

## Helm chart versions 1.x.x

To deploy an additional Broker into the same namespace as an existing broker, follow these examples.

### Deploy with existing service account

```
helm install <ENTER_UNIQUE_CHART_NAME> snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set serviceAccount.create=false \
             --set serviceAccount.name=<EXISTING_SERVICE_ACCOUNT> \
             -n <EXISTING_NAMESPACE>
```

### Deploy with new service account

```
helm install <ENTER_UNIQUE_CHART_NAME> snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set serviceAccount.name=<NEW_SERVICE_ACCOUNT_TO_BE_CREATED> \
             -n <EXISTING_NAMESPACE>
```
