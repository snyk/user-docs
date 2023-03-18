# Install Broker for Container Registry Agent using Helm

While the documentation for the [Snyk Broker](https://github.com/snyk/broker) requires the parameter `CR_AGENT_URL`, it is not required in this case.

Finally, you must include an `accept.json` file for this deployment. You must copy the new accept.json to the /snyk-broker folder

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=container-registry-agent \
             --set brokerClientUrl=http://container-registry-agent-broker-service:8000 \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set crType=<ENTER_CR_TYPE> \
             --set crBase=<ENTER_CR_BASE_URL> \
             --set crUsername=<ENTER_CR_URSERNAME> \
             --set crPassword=<ENTER_CR_PASSWORD> \
             --set acceptJsonFile=accept.json \
             -n snyk-broker --create-namespace
```

Allowed values for `crType`:

`artifactory-cr`\
`harbor-cr`\
`acr`\
`gcr`\
`docker-hub`\
`quay-cr`\
`nexus-cr`\
`github-cr`\
`ecr`\
`digitalocean-cr`

The following **Container Registry types (crType) require specific parameters**.

## **Elastic Container Registry (ecr)**

* crRoleArn
* crRegion
* crExternalId

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=container-registry-agent \
             --set brokerClientUrl=http://container-registry-agent-broker-service:8000 \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set crType=ecr \
             --set crRoleArn=<ENTER_CR_ROLE_ARN> \
             --set crRegion=<ENTER_CR_REGION> \
             --set crExternalId=<ENTER_CR_EXTERNAL_ID> \
             --set acceptJsonFile=accept.json \
             -n snyk-broker --create-namespace
```

## **DigitalOcean Container Registry (digitalocean-cr)**

* crToken

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=container-registry-agent \
             --set brokerClientUrl=http://container-registry-agent-broker-service:8000 \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set crType=digitalocean-cr \
             --set crBase=<ENTER_CR_BASE_URL> \
             --set crToken=<ENTER_CR_TOKEN> \
             --set acceptJsonFile=accept.json \
             -n snyk-broker --create-namespace
```
