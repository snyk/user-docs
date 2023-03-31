# Deploying multiple Brokers in the same namespace

To deploy an additional Broker into the same namespace as an existing broker, follow these examples.

## Deploy with existing service account

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

## Deploy with new service account

```
helm install <ENTER_UNIQUE_CHART_NAME> snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set serviceAccount.name=<NEW_SERVICE_ACCOUNT_TO_BE_CREATED> \
             -n <EXISTING_NAMESPACE>
```
