# Install and configure Broker using Helm

Installing the Snyk Broker with the [Broker Helm Chart](https://github.com/snyk/snyk-broker-helm) is the easiest way to deploy Snyk Broker if you are using Kubernetes.

**NOTE**\
The Helm chart does not manage connectivity, and thus you will be responsible for managing [ingress](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/ingress-options-with-snyk-broker-helm-installation) in the Kubernetes cluster.

**Multi-tenant settings for EU and AU**\
When you set up Broker, Code Agent, or both for use in EU or AU Multi-tenant environments, additional environment variables with the specific URLs are required.\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`\
For the URLs, see [EU and AU account datacenter creation](https://docs.snyk.io/snyk-processes/data-residency-at-snyk#eu-and-au-datacenter-account-creation).

The following are the allowed values for `scmType`:

Github.com: `github-com`\
Github Enterprise: `github-enterprise`\
Bitbucket: `bitbucket-server`\
Gitlab: `gitlab`\
Azure Repos: `azure-repos`\
Artifactory: `artifactory`\
Nexus: `nexus`\
Nexus2: `nexus2`\
Jira: `jira`\
Container Registry Agent: `container-registry-agent`\\

## Installation using the Snyk Broker Helm Chart

To use this chart, you must first add the Snyk Broker Helm Chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

Then run the commands for each SCM, registry, or Jira.

* GitHub
* GitHub Enterprise
* Bitbucket Server/Data Centre
* GitLab
* Azure Repos
* JFrog Artifactory
* Nexus 3
* Nexus 2
* Jira

Running the commands for each SCM, registry, or Jira creates a namespace called `snyk-broker`. To deploy into an existing namespace, adjust the `-n` parameter and delete the `--create-namespace` parameter. See also [Deploying multiple Brokers in the same namespace](advanced-setup-for-helm-chart-installation/deploying-multiple-brokers-in-the-same-namespace.md).

Additional commands are available to Install the Snyk Broker Code Agent and Container Registry Agent.

* [Snyk Broker Code Agent](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-code-agent/install-broker-for-code-agent-using-helm) (needed to enable SAST analysis)
* [Snyk Broker - Container Registry Agent](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent/install-broker-for-container-registry-agent-using-helm) (needed to connect to Container Registries)

## Set advanced parameters for Snyk Broker using the Helm Chart

When you set up Snyk Broker using Helm, you can set advanced parameters as explained on these pages:

* [Parameters for troubleshooting and providing your own certificate with Helm](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/parameters-for-troubleshooting-and-providing-your-own-certificate-with-helm)
* [Proxy settings for Broker Helm chart installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/proxy-settings-for-broker-helm-chart-installation)
* [Credential pooling with Docker and Helm](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/credential-pooling-with-docker-and-helm)
* [Adding custom accept.json for Helm installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/adding-custom-accept.json-for-helm-installation)
* [Ingress options with Snyk Broker Helm installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/ingress-options-with-snyk-broker-helm-installation)
* [Multi-tenant settings for Helm chart installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/multi-tenant-settings-for-helm-chart-installation)
* [Deploying multiple Brokers in the same namespace](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/deploying-multiple-brokers-in-the-same-namespace)
* [Custom additional options for Broker Helm Chart installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation)
* [Broker rules for Snyk Code](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/broker-rules-for-snyk-code)
* Kubernetes secrets and Helm Chart installation
* Image repository, tab, and Image Pull Secret
* Service accounts for Helm Chart installation
