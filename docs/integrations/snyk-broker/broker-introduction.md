# Snyk Broker introduction

## What is the purpose of Snyk Broker?

Snyk Broker is designed to connect Snyk products to self-hosted integrations that are not publicly accessible from the internet. Snyk Broker also allows you to do the following:

* Control Snyk access to your network, by limiting the files to which Snyk has access and the actions that Snyk can perform.
* Manage a fixed private IP for your integration, targeting the Broker.

## Snyk Broker components

Snyk Broker has a Server and a Client, components that are the same across all base integrations:

* **Broker Server** - running on Snyk SaaS backend\
  The Broker Server is provided by Snyk, and no installation is required on your part.
* **Broker Client** - a [Docker image](https://hub.docker.com/r/snyk/broker/) deployed in your infrastructure

<figure><img src="../../.gitbook/assets/Snyk Broker diagram.png" alt="Snyk Broker WebSocket initiated by Client over HTTPS"><figcaption><p>Snyk Broker WebSocket initiated by Client over HTTPS</p></figcaption></figure>

All data, both in transit and at rest, is encrypted. Communication between the Client and the Server takes place over a secure WebSocket connection. There is no need to open incoming ports since the communication is initiated outbound. Once the connection is initiated, the WebSocket connection is bi-directional.

## Using inbound and outbound connections with Snyk Broker

* There is no direct inbound connection from Snyk to the Broker Client. The Broker Client makes an outbound connection to [https://broker.snyk.io](https://broker.snyk.io), which establishes a WebSocket connection to allow communication with the Broker Server. Thus there is no need to allowlist a Snyk IP address; instead you can allow the Broker Client IP/port.
* The Broker Client initiates the outbound connection to establish the WebSocket. After the WebSocket connection initiated by the Broker Client is established, Snyk can send inbound requests to the Broker Client via the WebSocket. Thus you do not need to allow inbound connections to the Broker Client from Snyk-specific IP addresses or other external IP addresses.

## **Approved data list for Snyk Broker**

The Broker Client maintains an approved data list for inbound and outbound data requests. Only requests included in this approved list are allowed.

The default approved list limits requests as follows:

* **Inbound:** For Snyk Open Source, Snyk.io is allowed to fetch and view only dependency manifest files and the `.snyk` policy file. No other source code is viewed, extracted, or modified. Additional `.snyk` files may be checked in to support the Snyk patch mechanism and for any ignore instructions included in your vulnerability policy. Snyk Code and Snyk IaC need access to the entire repository. For more information see [How Snyk handles your data](../../snyk-processes/how-snyk-handles-your-data.md).
* **Outbound:** Git repo webhooks are set when you configure your Broker Client setup, to enable automatic Snyk scans triggered when new pull requests or merge events are submitted by your developers. Webhook notifications are delivered to Snyk via the Broker Client only for events relevant to Snyk actions--push to branch, pull request opened, and only when the event data also includes a dependency manifest file or a .`snyk` policy file.

Because of the limitations of the default approved list, if you are interested in scanning Infrastructure as Code files with the Snyk Broker, you must [add and configure an `accept.json`](snyk-broker-infrastructure-as-code-detection/) file in your Broker deployment.

To learn more about the approved data list and the `accept.json` file, see [Install and configure the Snyk Broker client](set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md), Custom approved listing filter.

## **Snyk Broker usage**

To use **Snyk Open Source** with Snyk Broker, you need only the Broker Server and the Broker Client components.

To use other Snyk products with Snyk Broker, you need to add an additional component or configurations, and to add parameters to the Broker Client setup:

* **Snyk Code** – add the [**Code Agent** ](snyk-broker-code-agent/)component to enable Snyk Code analysis of repositories in SCMs that are integrated through Snyk Broker.
* **Snyk Container** – add the [**Container Registry Agent**](snyk-broker-container-registry-agent/) to enable the connection to network-restricted container registries container registries and the analysis of container images.
* **Snyk Infrastructure as Code** – configure the [**`accept.json`** file with additional parameters](snyk-broker-infrastructure-as-code-detection/) to detect and analyze Terraform, CloudFormation, and Kubernetes configuration files through Snyk Broker.

## **Supported integrations with Snyk Broker**

Snyk Broker currently integrates with these Git repository systems:

* [GitHub](../git-repository-scm-integrations/github-integration.md) and [GitHub Enterprise](../git-repository-scm-integrations/github-enterprise-integration.md) (Cloud and On-prem)
* GitLab (Cloud and On-prem)
* [Bitbucket Server / Data Center](../git-repository-scm-integrations/bitbucket-data-center-server-integration.md) (On-prem)
* [Azure Repos](../git-repository-scm-integrations/azure-repos-integration.md) (Cloud and On-prem)

In addition, Snyk Broker integrates with [Jira Server/Jira Data Center](../notifications-ticketing-system-integrations/jira.md), [JFrog Artifactory](../private-registry-integrations/artifactory-registry-setup.md), and [Nexus Repository Manager](../private-registry-integrations/nexus-repo-manager-setup.md).

With the Container Registry Agent, Snyk Broker also connects to all [Snyk-supported container registries](snyk-broker-container-registry-agent/).

## Common questions about Snyk Broker

* How often is Snyk Broker updated?\
  Snyk Broker is updated any time there are new features available and when there are fixes.
* How often is Snyk Broker checked for vulnerabilities?\
  The Snyk Broker application and images are being tested on a daily basis for vulnerabilities.
* What is the SLA to fix vulnerabilities?\
  There is a 14-day SLA for fixing high vulnerabilities and five-day SLA for fixing critical vulnerabilities in public images.
