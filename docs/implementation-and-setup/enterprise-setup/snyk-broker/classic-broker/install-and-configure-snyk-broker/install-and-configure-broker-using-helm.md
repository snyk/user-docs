# Install and configure Broker using Helm

{% hint style="info" %}
[Broker version 4.205.1](https://github.com/snyk/broker/blob/cb4f89e05eb42605f076321b952cdb7e57bf4111/config.default.json#L8) has been [released](https://updates.snyk.io). In this version, all `ACCEPT` rule flags will be enabled by default. This reduces needed configuration. If you do not want a specific `ACCEPT` rule flag to be enabled, you can opt out of the default `ACCEPT` all behavior by adding `ACCEPT_=false` to your Broker client configuration.
{% endhint %}

{% hint style="info" %}
**Multi-tenant settings for regions**\
When installing, you must add a command in your script to set the `brokerServerUrl`. This is the URL of the Broker server for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
{% endhint %}

Before starting installation, review the [Prerequisites](../prepare-snyk-broker-for-deployment/#prerequisites-for-snyk-broker) and other information on the page [Prepare Snyk Broker for deployment](../prepare-snyk-broker-for-deployment/).

**I**f you are using Kubernetes, Snyk recommends that you install Snyk Broker with the [Broker Helm Chart](https://github.com/snyk/snyk-broker-helm).

For all other environments, you can install Snyk Broker using the [Docker images](https://github.com/snyk/broker) provided by Snyk. For details, see [Install and configure Broker using Docker](install-and-configure-broker-using-docker.md).

{% hint style="info" %}
**Multi-tenant settings for regions other than the default**\
When you set up Snyk Broker for use in regions other than the default, additional environment variables with specific URLs are required. For the URLs and examples, see [Broker URLs](../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
{% endhint %}

## Install using the Snyk Broker Helm Chart

The Helm chart does not manage connectivity, and thus, you will be responsible for managing [ingress](advanced-configuration-for-helm-chart-installation/ingress-options-with-snyk-broker-helm-installation.md) in the Kubernetes cluster.

To use this chart, you must first add the Snyk Broker Helm Chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

Then run the commands to customize the environment variables for each SCM, registry, or Jira as explained on the following pages:

* [GitHub](github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-helm.md) `scmType`: `github-com`
* [GitHub Enterprise](github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-helm.md) `scmType`: `github-enterprise`
* [Bitbucket Server/Data Center](bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-helm.md) `scmType`: `bitbucket-server`
* [GitLab](gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-install-and-configure-using-helm.md) `scmType`: `gitlab`
* [Azure Repos](azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-install-and-configure-and-configure-using-helm.md) `scmType`: `azure-repos`
* [JFrog Artifactory](artifactory-repository-install-and-configure-broker/artifactory-repository-install-and-configure-using-helm.md) `scmType`: `artifactory`
* [Nexus 3](nexus-repository-prerequisites-and-steps-to-install-and-configure-broker/nexus-repository-install-and-configure-using-helm.md) `scmType`: `nexus`
* [Nexus 2](nexus-repository-prerequisites-and-steps-to-install-and-configure-broker/nexus-repository-install-and-configure-using-helm.md) `scmType`: `nexus2`
* [Jira](jira-prerequisites-and-steps-to-install-and-configure-broker/jira-install-and-configure-using-helm.md) `scmType`: `jira`

`scmType` designates the system type. For JFrog and Nexus, this is an artifact repository. For Jira, it is a ticket management system.

Running the commands for each SCM, registry, or Jira creates a namespace called `snyk-broker`. To deploy into an existing namespace, adjust the `-n` parameter and delete the `--create-namespace` parameter. See also [Deploying multiple Brokers in the same namespace](advanced-configuration-for-helm-chart-installation/deploying-multiple-brokers-in-the-same-namespace.md).

Beginning with version 2.0.0, all created objects have a suffix based on the release name, allowing for multiple Brokers in the same namespace. For backward compatibility, 2.1.0 introduces a `disableSuffixes` flag to revert to the 1.x.x behavior by adding `--set disableSuffixes=true`.

Additional commands are available to install [Snyk Broker - Container Registry Agent](../../snyk-broker-container-registry-agent/), needed to connect to Container Registries; `scmType`: `container-registry-agent`\\

You can verify that the Broker is running by looking at the settings for your brokered integration in [the Snyk Web UI](https://app.snyk.io) to see a confirmation message that you are connected. You can start importing Projects once you are connected.

## Advanced configuration using Helm

You can pass any environment variable of your choice in the Helm command. For details, see [Custom additional options for Broker Helm Chart](advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md).&#x20;

For example, if you want to pass BROKER\_CLIENT\_VALIDATION\_URL using the Helm chart, the additional parameters would be:

`--set env[0].name=BROKER_CLIENT_VALIDATION_URL \`\
`--set env[0].value=whatever_value_makes_sense`

`Additional parameters would be:`

`--set env[1].name=MY_OTHER_ENV_VAR \`\
`--set env[1].value="other env with spaces" \`\
`--set env[2].name=THIRD_ENV_VAR \`\
`--set env[2].value=and_so_on`

Follow the instructions for [Advanced configuration for Helm Chart installation](advanced-configuration-for-helm-chart-installation/) to make configuration changes as needed.
