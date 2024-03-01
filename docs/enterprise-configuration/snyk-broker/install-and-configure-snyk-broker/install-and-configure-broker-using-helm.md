# Install and configure Broker using Helm

Before starting installation, review the [Prerequisites](../prepare-snyk-broker-for-deployment.md#prerequisites-for-snyk-broker) and other information on the page [Prepare Snyk Broker for deployment](../prepare-snyk-broker-for-deployment.md).

**If you are using Kubernetes**, Snyk recommends that you **install Snyk Broker with the** [**Broker Helm Chart**](https://github.com/snyk/snyk-broker-helm).

For **all other environments**, you can install Snyk Broker using the [Docker images](https://github.com/snyk/broker) provided by Snyk. For details, see [Install and configure Broker using Docker](install-and-configure-broker-using-docker.md).

## Install using the Snyk Broker Helm Chart

The Helm chart does not manage connectivity, and thus you will be responsible for managing [ingress](advanced-configuration-for-helm-chart-installation/ingress-options-with-snyk-broker-helm-installation.md) in the Kubernetes cluster.

To use this chart, you must first add the Snyk Broker Helm Chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

Then run the commands to customize the environment variables for each SCM, registry, or Jira as explained on the following pages:

* [GitHub](github-prerequisites-and-steps-to-install-and-configure-broker/githhub.com-install-and-configure-using-helm.md) `scmType`: `github-com`
* [GitHub Enterprise](github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-helm.md) `scmType`: `github-enterprise`
* [Bitbucket Server/Data Center](bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-helm.md) `scmType`: `bitbucket-server`
* [GitLab](gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-install-and-configure-using-helm.md) `scmType`: `gitlab`
* [Azure Repos](azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-install-and-configure-and-configure-using-helm.md) `scmType`: `azure-repos`
* [JFrog Artifactory](artifactory-repository-install-and-configure-broker/artifactory-repository-install-and-configure-using-helm.md) `scmType`: `artifactory`
* [Nexus 3](nexus-repository-prerequisites-and-steps-to-install-and-configure-broker/nexus-repository-install-and-configure-using-helm.md) `scmType`: `nexus`
* [Nexus 2](nexus-repository-prerequisites-and-steps-to-install-and-configure-broker/nexus-repository-install-and-configure-using-helm.md) `scmType`: `nexus2`
* [Jira](jira-prerequisites-and-steps-to-install-and-configure-broker/jira-install-and-configure-using-helm.md) `scmType`: `jira`

`scmType` designates the system type. For JFrog and Nexus, this is an artifact repository, For Jira, it is a ticket management system.

Running the commands for each SCM, registry, or Jira creates a namespace called `snyk-broker`. To deploy into an existing namespace, adjust the `-n` parameter and delete the `--create-namespace` parameter. See also [Deploying multiple Brokers in the same namespace](advanced-configuration-for-helm-chart-installation/deploying-multiple-brokers-in-the-same-namespace.md).

Beginning with version 2.0.0, all created objects have a suffix based on the release name, allowing for multiple Brokers in the same namespace. For backward compatibility, 2.1.0 introduces a `disableSuffixes` flag to revert to the 1.x.x behavior by adding `--set disableSuffixes=true`.

Additional commands are available to Install the Snyk Broker Code Agent and Container Registry Agent.

* [Snyk Broker Code Agent](../snyk-broker-code-agent/): needed to enable SAST analysis
* [Snyk Broker - Container Registry Agent](../snyk-broker-container-registry-agent/): needed to connect to Container Registries; `scmType`: `container-registry-agent`\\

You can verify that the Broker is running by looking at the settings for your brokered integration in [the Snyk Web UI](https://app.snyk.io) to see a confirmation message that you are connected. You can start importing Projects once you are connected.

## Advanced configuration using Helm

You can pass any environment variable of your choice in the Helm command. For details, see [Custom additional options for Broker Helm Chart](advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md). Follow the instructions for [Advanced configuration for Helm Chart installation](advanced-configuration-for-helm-chart-installation/) to make configuration changes as needed.
