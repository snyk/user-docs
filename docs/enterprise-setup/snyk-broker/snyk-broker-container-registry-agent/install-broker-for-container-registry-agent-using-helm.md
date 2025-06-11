# Install Broker for Container Registry Agent using Helm

Installing the Broker Container Registry Agent using Docker requires the parameter `CR_AGENT_URL`, but it is not required to install using Helm. The environment variables are defined for [installing with Docker ](./#set-up-the-remote-connection-for-container-registry-agent-using-docker)and also apply to installing with Helm.

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=container-registry-agent \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set crType=<ENTER_CR_TYPE> \
             --set crBase=<ENTER_CR_BASE_URL> \
             --set crUsername=<ENTER_CR_URSERNAME> \
             --set crPassword=<ENTER_CR_PASSWORD> \
             -n snyk-broker --create-namespace
```

Allowed values for `crType`:

```
acr
artifactory-cr
digitalocean-cr
docker-hub
ecr
gcr
google-artifact-cr
github-cr
gitlab-cr
harbor-cr
quay-cr
nexus-cr
```

Elastic Container Registry and Digital Ocean Container Registry require specific parameters as explained in the sections that follow.

## **Elastic Container Registry (ecr)**

* crRoleArn
* crRegion
* crExternalId

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=container-registry-agent \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set crType=ecr \
             --set crRoleArn=<ENTER_CR_ROLE_ARN> \
             --set crRegion=<ENTER_CR_REGION> \
             --set crExternalId=<ENTER_CR_EXTERNAL_ID> \
             -n snyk-broker --create-namespace
```

## **DigitalOcean Container Registry (digitalocean-cr)**

* crToken

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=container-registry-agent \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set crType=digitalocean-cr \
             --set crBase=<ENTER_CR_BASE_URL> \
             --set crToken=<ENTER_CR_TOKEN> \
             -n snyk-broker --create-namespace
```
