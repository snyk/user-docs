# Snyk Broker introduction

### What is the purpose of the Snyk Broker?

The Snyk Broker is designed to connect Snyk products to self-hosted integrations that are not accessible from the internet. It also allows you to:

* Keep sensitive data such as your access tokens inside your private network, never sharing that information with Snyk.
* Control Snyk access to your network, by limiting the files that Snyk can access and the actions that Snyk can perform.
* Manage a fixed private IP for your integration (targeting the Broker).

### The Snyk Broker components

The Snyk Broker has a Server and a Client components that are the same across all base integrations:

* **Broker Server** - running on Snyk SaaS backend.\
  **Note:** The Broker Server is provided by Snyk, and no installation is required on your part.
* **Broker Client** - a [Docker image](https://hub.docker.com/r/snyk/broker/) deployed in your infrastructure.

![](<../../../.gitbook/assets/Snyk Broker diagram.png>)

All data, both in transit and at rest, is encrypted. Communication between the Client and the Server takes place over a secure WebSocket connection. It does not require opening incoming ports since the communication is initiated outbound, following which it establishes a two-way communication path over WebSocket.

### Using inbound and outbound connections

* There is no direct inbound connection from Snyk to the Broker Client. The Broker Client makes an outbound connection to [https://broker.snyk.io](https://broker.snyk.io), which establishes a WebSocket connection to allow communication with the Broker Server. This way, there is no need to grant permissions to specific IPs towards Snyk, as you can point to the Broker Client IP/port.
* The Broker Client initiates the outbound connection to establish the WebSocket. After the WebSocket is established, that allows inbound requests from Snyk via the WebSocket, with no need to allow inbound connectivity to the world or to Snyk-specific IP addresses.

### **Approved data list**

The Broker Client maintains an approved data list for inbound and outbound data requests. Only requests included in this approved list are allowed.

The default approved list limits requests as follows:

* **Inbound:** Snyk.io is only allowed to fetch and view dependency manifest files and the Snyk policy file. No other source code is viewed, extracted, or modified. Additional files (.snyk files) may be checked in to support our patch mechanism and for any ignore instructions included in your vulnerability policy.
* **Outbound:** Git repo webhooks are set when you configure your Broker Client setup, to enable automatic Snyk scans triggered when new pull requests or merge events are submitted by your developers. Webhook notifications are delivered to Snyk via the Broker Client for only events relevant to Snyk actions (push to branch, pull request opened), and only when the event data also includes a dependency manifest file or a Snyk policy file.

Because of the limitations of the default approved list, if you are interested in scanning Infrastructure as Code files with the Snyk Broker, you will need to [add and configure an `accept.json`](snyk-broker-infrastructure-as-code-detection/) file in your Broker deployment.

To learn more about the approved data list and the `accept.json` file, see [Custom approved listing filter](set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md).

### **The Snyk Broker usage**

To use **Snyk Open Source** with the Snyk Broker, all you need is the Broker Server and the Broker Client components.&#x20;

To use other Snyk products with the Snyk Broker, you need to add an additional component or configurations, and to add parameters to the Broker Client setup:

* **Snyk Code** – add the [**Code Agent** ](snyk-broker-code-agent.md)component to enable the Snyk Code analysis of repositories in SCMs that are integrated through the Snyk Broker.
* **Snyk Container** – add the [**Container Registry Agent**](snyk-broker-container-registry-agent/) to enable the connection to self-hosted/private container registries and the analysis of container images.
* **Snyk Infrastructure as Code** – configure the [**`accept.json`** file with additional parameters](snyk-broker-infrastructure-as-code-detection/) **** to detect and analyze Terraform, CloudFormation, and Kubernetes configuration files through the Snyk Broker.

### **Supported integrations**

Snyk Broker currently integrates with these Git Repository systems:

* [GitHub](https://docs.snyk.io/integrations/git-repository-scm-integrations/github-integration) and [GitHub Enterprise](https://docs.snyk.io/integrations/git-repository-scm-integrations/github-enterprise-integration) (Cloud and On-prem)
* [GitLab](https://docs.snyk.io/integrations/git-repository-scm-integrations/gitlab-integration) (Cloud and On-prem)
* [Bitbucket Server / Data Center](../git-repository-scm-integrations/bitbucket-data-center-server-integration.md) (On-prem)
* [Azure Repos](https://docs.snyk.io/integrations/git-repository-scm-integrations/azure-repos-integration) (Cloud and On-prem)

In addition, Snyk Broker integrates with [Jira Server/Jira Data Center](../notifications-ticketing-system-integrations/jira.md) and [JFrog Artifactory](../private-registry-integrations/artifactory-registry-setup.md).

With the Container Registry Agent, Snyk Broker also connects to all [Snyk supported Container Registries](snyk-broker-container-registry-agent/).

### Common questions

* How often is the Snyk Broker updated?\
  The Snyk Broker is updated anytime there are new features available and when there are fixes.
* How often is the Snyk Broker checked for vulnerabilities?\
  The Snyk Broker application and images are being tested on a daily basis for vulnerabilities.
* What is the SLA to fix vulnerabilities?\
  There is a 14 day SLA for fixing high vulnerabilities and 5 days SLA for fixing critical vulnerabilities in public images.
