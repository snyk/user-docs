# Install and configure Broker using Helm

Snyk Broker is an open-source tool that acts as a proxy between Snyk and special integrations, providing for access to your code by [snyk.io](http://snyk.io/) to scan it and return results to you. SCM integrations with Broker support Snyk Open Source, Snyk Code, Snyk Container (Dockerfile), and Snyk IaC. For details about Snyk Broker, including how it works, how to deploy it, commit signing, upgrading, and troubleshooting, see the [Snyk Broker user documentation](https://docs.snyk.io/snyk-admin/snyk-broker).

**If you are using Kubernetes**, Snyk recommends that you **install Snyk Broker with the** [**Broker Helm Chart**](https://github.com/snyk/snyk-broker-helm). For details, see the instructions for each integration listed in the section that follows, "Installation using the Snyk Broker Helm Chart".

For **all other environments**, you can install Snyk Broker using the [Docker images](https://github.com/snyk/broker) provided by Snyk. For details, see [Install and configure Broker using Docker](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-docker).

## Installation using the Snyk Broker Helm Chart

**Multi-tenant settings for EU and AU**\
When you set up Broker, Code Agent, or both for use in EU or AU Multi-tenant environments, additional environment variables with the specific URLs are required.\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`\
For the URLs, see [EU and AU account datacenter creation](https://docs.snyk.io/snyk-processes/data-residency-at-snyk#eu-and-au-datacenter-account-creation).

**NOTE**\
The Helm chart does not manage connectivity, and thus you will be responsible for managing [ingress](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/ingress-options-with-snyk-broker-helm-installation) in the Kubernetes cluster.

To use this chart, you must first add the Snyk Broker Helm Chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

Then run the commands for each SCM, registry, or Jira as explained in the user documentation:

* [GitHub](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/githhub.com-install-and-configure-using-helm) `scmType`: `github-com`
* [GitHub Enterprise](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/github-enterprise-install-and-configure-using-helm) `scmType`: `github-enterprise`
* [Bitbucket Server/Data Center](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-install-and-configure-using-helm) `scmType`: `bitbucket-server`
* [GitLab](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/gitlab-install-and-configure-using-helm) `scmType`: `gitlab`
* [Azure Repos](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/azure-repos-install-and-configure-and-configure-using-helm) `scmType`: `azure-repos`
* [JFrog Artifactory](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/artifactory-repository-install-and-configure-using-helm) `scmType`: `artifactory`
* [Nexus 3](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/nexus-repository-install-and-configure-using-helm) `scmType`: `nexus`
* [Nexus 2](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/nexus-repository-install-and-configure-using-helm) `scmType`: `nexus2`
* [Jira](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/jira-install-and-configure-using-helm) `scmType`: `jira`

Running the commands for each SCM, registry, or Jira creates a namespace called `snyk-broker`. To deploy into an existing namespace, adjust the `-n` parameter and delete the `--create-namespace` parameter. See also [Deploying multiple Brokers in the same namespace](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/advanced-setup-for-helm-chart-installation/deploying-multiple-brokers-in-the-same-namespace).

Additional commands are available to Install the Snyk Broker Code Agent and Container Registry Agent.

* [Snyk Broker Code Agent](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-code-agent/install-broker-for-code-agent-using-helm) (needed to enable SAST analysis)
* [Snyk Broker - Container Registry Agent](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent/install-broker-for-container-registry-agent-using-helm) (needed to connect to Container Registries; `scmType`: `container-registry-agent`\\)

## Set advanced parameters for Snyk Broker using the Helm Chart

When you set up Snyk Broker using Helm, you can set advanced parameters as explained in the following documentation:

* [Parameters for troubleshooting and providing your own certificate with Helm](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/parameters-for-troubleshooting-and-providing-your-own-certificate-with-helm)
* [Proxy settings for Broker Helm chart installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/proxy-settings-for-broker-helm-chart-installation)
* [Credential pooling with Docker and Helm](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/credential-pooling-with-docker-and-helm)
* [Adding custom accept.json for Helm installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/adding-custom-accept.json-for-helm-installation)
* [Ingress options with Snyk Broker Helm installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/ingress-options-with-snyk-broker-helm-installation)
* [Multi-tenant settings for Helm chart installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/multi-tenant-settings-for-helm-chart-installation)
* [Deploying multiple Brokers in the same namespace](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/deploying-multiple-brokers-in-the-same-namespace)
* [Custom additional options for Broker Helm Chart installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation)
* [Broker rules for Snyk Code](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation/broker-rules-for-snyk-code)
* [Kubernetes secrets and Helm Chart installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/advanced-setup-for-helm-chart-installation/kubernetes-secrets-and-helm-chart-installation)
* [Image repository, tab, and Image Pull Secret](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/advanced-setup-for-helm-chart-installation/image-repository-tab-and-image-pull-secret)
* [Service accounts for Helm Chart installation](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-snyk-broker/advanced-setup-for-helm-chart-installation/service-accounts-for-helm-chart-installation)
