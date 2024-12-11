# Install Snyk Broker - Code Agent using Helm

{% hint style="warning" %}
**Deprecated**

The Code Agent is deprecated and is no longer maintained.

The preferred method of running Snyk Code analysis using Snyk Broker is through [Brokered Code](../install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/git-clone-through-broker.md). The Code Agent is an alternative method without advantages. For details, contact your Snyk Integration Consultant or Technical Success Manager or contact [Snyk Support](https://support.snyk.io).

The automatic [PR Checks](../../../scan-with-snyk/pull-requests/pull-request-checks/) feature is not supported for Snyk Broker - Code Agent.
{% endhint %}

To deploy the Snyk Broker Code Agent, you must set the `enableCodeAgent` flag to `true`. Ensure you have the proper entries in the `accept.json` file. Retrieve the example file for the SCM you are using from the [Broker Client Templates repository](https://github.com/snyk/broker/tree/master/client-templates). Ensure you have the additional entries as specified by the Snyk Broker Code Agent documentation.

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
