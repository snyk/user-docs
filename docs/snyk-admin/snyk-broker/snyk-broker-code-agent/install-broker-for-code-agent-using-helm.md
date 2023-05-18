# Install Broker for Code Agent using Helm

To deploy the Snyk Code Agent, you must set the `enableCodeAgent` flag to `true`. See more information about the [Snyk Code Agent](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent). Ensure you have the proper entries in the accept.json file. Grab the example file for the appropriate SCM [HERE](https://github.com/snyk/broker/tree/master/client-templates). Ensure you have the additional entries as specified by the Snyk Code Agent documentation.

Here is an example command for GitLab:

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

Note: The `brokerClientUrl` is going to be the address of the Broker Container. The default port for the broker container is `8000`. See the values file for more information. Also, the accept.json must be in the same directory as the Helm Chart. For definitions of the environment variables see [Running the Code Agent container](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-4-setting-up-the-code-agent/step-4.2-running-the-code-agent-container#running-the-code-agent-container).
