# Install and configure Broker using Helm

Snyk Broker is an open-source tool that acts as a proxy between Snyk and special integrations, providing for access to your code by [snyk.io](http://snyk.io/) to scan it and return results to you. SCM integrations with Broker support Snyk Open Source, Snyk Code, Snyk Container (Dockerfile), and Snyk IaC. For details about Snyk Broker, including how it works, how to deploy it, commit signing, upgrading, and troubleshooting, see the [Snyk Broker user documentation](../).

**If you are using Kubernetes**, Snyk recommends that you **install Snyk Broker with the** [**Broker Helm Chart**](https://github.com/snyk/snyk-broker-helm). For details, see the instructions for each integration listed in the section that follows, "Installation using the Snyk Broker Helm Chart".

For **all other environments**, you can install Snyk Broker using the [Docker images](https://github.com/snyk/broker) provided by Snyk. For details, see [Install and configure Broker using Docker](install-and-configure-broker-using-docker.md).

## Installation using the Snyk Broker Helm Chart

**Multi-tenant settings for EU and AU**\
When you set up Broker, Code Agent, or both for use in EU or AU Multi-tenant environments, additional environment variables with the specific URLs are required.\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`\
For the URLs, see [EU and AU account datacenter creation](https://docs.snyk.io/snyk-processes/data-residency-at-snyk#eu-and-au-datacenter-account-creation).

**NOTE**\
The Helm chart does not manage connectivity, and thus you will be responsible for managing [ingress](advanced-configuration-for-helm-chart-installation/ingress-options-with-snyk-broker-helm-installation.md) in the Kubernetes cluster.

To use this chart, you must first add the Snyk Broker Helm Chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

Then run the commands for each SCM, registry, or Jira as explained in the user documentation:

* [GitHub](github-install-and-configure-broker/githhub.com-install-and-configure-using-helm.md) `scmType`: `github-com`
* [GitHub Enterprise](github-enterprise-install-and-configure-broker/github-enterprise-install-and-configure-using-helm.md) `scmType`: `github-enterprise`
* [Bitbucket Server/Data Center](bitbucket-server-data-center-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-helm.md) `scmType`: `bitbucket-server`
* [GitLab](gitlab-install-and-configure-broker/gitlab-install-and-configure-using-helm.md) `scmType`: `gitlab`
* [Azure Repos](azure-repos-install-and-configure-broker/azure-repos-install-and-configure-and-configure-using-helm.md) `scmType`: `azure-repos`
* [JFrog Artifactory](artifactory-repository-install-and-configure-broker/artifactory-repository-install-and-configure-using-helm.md) `scmType`: `artifactory`
* [Nexus 3](nexus-repository-install-and-configure-broker/nexus-repository-install-and-configure-using-helm.md) `scmType`: `nexus`
* [Nexus 2](nexus-repository-install-and-configure-broker/nexus-repository-install-and-configure-using-helm.md) `scmType`: `nexus2`
* [Jira](jira-install-and-configure-broker/jira-install-and-configure-using-helm.md) `scmType`: `jira`

**NOTE**\
`scmType` designates the system type. For JFrog and Nexus, this is an artifact repository, For Jira, it is a ticket management system,

Running the commands for each SCM, registry, or Jira creates a namespace called `snyk-broker`. To deploy into an existing namespace, adjust the `-n` parameter and delete the `--create-namespace` parameter. See also Deploying multiple Brokers in the same namespace.

**NOTE**\
Beginning with version 2.0.0, all created objects have a suffix based on the release name, allowing for multiple brokers in the same namespace. For backward compatibility, 2.1.0 introduces a disableSuffixes flag to revert to the 1.x.x behavior by adding `--set disableSuffixes=true`

Additional commands are available to Install the Snyk Broker Code Agent and Container Registry Agent.

* [Snyk Broker Code Agent](../snyk-broker-code-agent/) (needed to enable SAST analysis)
* [Snyk Broker - Container Registry Agent](../snyk-broker-container-registry-agent/) (needed to connect to Container Registries; `scmType`: `container-registry-agent`\\)

You can verify that the Broker is running by looking at the settings for your brokered integration in [the Snyk Web UI](https://app.snyk.io) to see a confirmation message that you are connected. You can start importing Projects once you are connected.

## Set advanced parameters for Snyk Broker using the Helm Chart

When you set up Snyk Broker using Helm, you can set advanced parameters as explained on the following pages:

* [Parameters for troubleshooting and providing your own certificate with Helm](advanced-configuration-for-helm-chart-installation/parameters-for-troubleshooting-and-providing-your-own-certificate-with-helm.md)
* [Proxy settings for Broker Helm chart installation](advanced-configuration-for-helm-chart-installation/proxy-settings-for-broker-helm-chart-installation.md)
* [Credential pooling with Docker and Helm](advanced-configuration-for-snyk-broker-docker-installation/credential-pooling-with-docker-and-helm.md)
* [Adding custom accept.json for Helm installation](advanced-configuration-for-helm-chart-installation/adding-custom-accept.json-for-helm-installation.md)
* [Ingress options with Snyk Broker Helm installation](advanced-configuration-for-helm-chart-installation/ingress-options-with-snyk-broker-helm-installation.md)
* [Multi-tenant settings for Helm chart installation](advanced-configuration-for-helm-chart-installation/multi-tenant-settings-for-helm-chart-installation.md)
* [Deploying multiple Brokers in the same namespace](advanced-configuration-for-helm-chart-installation/deploying-multiple-brokers-in-the-same-namespace.md)
* [Custom additional options for Broker Helm Chart installation](advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md)
* [Broker rules for Snyk Code](advanced-configuration-for-helm-chart-installation/broker-rules-for-snyk-code.md)
* [High availability mode](../high-availability-mode.md)
* [Large manifest files rule addition](advanced-configuration-for-helm-chart-installation/snyk-open-source-scans-sca-of-large-manifest-files-helm-setup.md)
* [Insecure downstream mode](advanced-configuration-for-helm-chart-installation/insecure-downstream-mode.md)
