# Install Snyk Broker - Code Agent using Helm

To deploy the Snyk Broker Code Agent, you must set the `enableCodeAgent` flag to `true`. Ensure you have the proper entries in the `accept.json` file. Retrieve the example file for the SCM you are using from the  [Broker Client Templates repository](https://github.com/snyk/broker/tree/master/client-templates). Ensure you have the additional entries as specified by the Snyk Broker Code Agent documentation.

An example command for GitLab follows:

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=gitlab  \
             --set brokerToken=<ENTER_BROKER_TOKEN> \ 
             --set scmToken=<ENTER_SCM_TOKEN> \
             --set gitlab=<ENTER_GITLAB_URL>  \
            --set acceptJsonFile=accept.json \
            --set brokerClientUrl=http://<BROKER_CLIENT_URL> \ 
            --set enableCodeAgent=true \ 
            --set snykToken=<ENTER_SNYK_TOKEN> \
            -n snyk-broker --create-namespace
```

The `brokerClientUrl` is is the address of the Broker Container. The default port for the Broker container is `8000`. See the values file for more information.

The `accept.json` must be in the same directory as the Helm Chart.
