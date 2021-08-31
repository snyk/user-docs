# Broker introduction

Snyk Broker is an open-source tool, that can act as a proxy between Snyk and integrations including:

* Your Source Code Management \(SCM\) system on-premise platforms.
* Your publically-accessible Git-based repositories, allowing you to view and control Snyk activity in those repositories for increased data security.
* Your on-premise Jira installation.
* Artifactory for open source registry or container registry integrations.
* [Kubernetes configuration files](https://docs.snyk.io/snyk-infrastructure-as-code/scan-kubernetes-configuration-files/detecting-kubernetes-configuration-files-using-a-broker) using Snyk Infrastructure as Code \(IaC\).

Snyk Broker is an open-source project, hosted at [GitHub](https://github.com/snyk/broker), and published as a set of Docker images for specific integrations. See the [Github broker documentation](https://github.com/snyk/broker/blob/master/README.md).

{% hint style="info" %}
**Feature availability**  
Snyk Broker is available with Enterprise plans. See [Pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

### Broker benefits

Snyk Broker allows you to:

* Keep sensitive data such as your access tokens inside your private network, never sharing that information with Snyk.
* Provide controlled access to the network by Snyk, limiting the files Snyk can access, and the actions that Snyk can perform.

### Broker components

Snyk Broker has a client and a server component:

* Broker client: a [Docker image](https://hub.docker.com/r/snyk/broker/) deployed in your infrastructure.
* Broker server: running on Snyk SaaS backend

![](../../.gitbook/assets/image2-4-.png)

All data, both in transit and at rest, is encrypted. Communication between the client and server takes place over a secure WebSocket connection. On startup, it dials out and establishes a two way communication path. It does not require opening incoming ports.

### Using inbound and outbound connections

* There is no direct inbound connection from Snyk to the Broker client. The broker client makes an outbound connection to [https://broker.snyk.io](https://broker.snyk.io), which establishes a WebSocket connection to allow communication with the Broker server. This way, there is no need to grant permissions to specific IPs.
* The Broker client initiates the outbound connection to establish the WebSocket. After the websocket is established, that allows inbound requests from Snyk via the WebSocket, with no need to allow inbound connectivity to the world or to Snyk specific IP addresses.

### **Approved data list**

The Broker client maintains an approved data list for inbound and outbound data requests. Only requests included in this approved list are allowed.

The default approved list limits requests as follows:

* Inbound: Snyk.io is only allowed to fetch and view dependency manifest files and the Snyk policy file. No other source code is viewed, extracted, or modified. Additional files \(.snyk files\) may be checked in to support our patch mechanism and for any ignore instructions included in your vulnerability policy.
* Outbound: Git repo webhooks are set when you configure your Broker setup, to enable automatic Snyk scans triggered when new pull requests or merge events are submitted by your developers. Webhook notifications are delivered to Snyk via the Broker client for only events relevant to Snyk actions \(push to branch, pull request opened\), and only when the event data also includes a dependency manifest file or a Snyk policy file.

### **Supported integrations**

Snyk Broker currently integrates with these [Git Repository](https://support.snyk.io/hc/en-us/sections/360001138098-Git-repository-SCM-integrations) systems:

* [GitHub](https://docs.snyk.io/integrations/git-repository-scm-integrations/github-integration) and [GitHub Enterprise](https://support.snyk.io/hc/en-us/articles/360015951318-GitHub-Enterprise-Server-Integration)
* [GitLab](https://docs.snyk.io/integrations/git-repository-scm-integrations/gitlab-integration)
* [Bitbucket Server](https://support.snyk.io/hc/en-us/articles/360004002218-Bitbucket-Server-integration) and Bitbucket Data Center
* [Azure Repos](https://docs.snyk.io/integrations/git-repository-scm-integrations/azure-repos-integration)

In addition, Snyk Broker integrates with [Jira Server](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations/jira), Jira Data Center and [Artifactory](https://docs.snyk.io/integrations/private-registry-integrations/artifactory-registry-setup).

### **Supported manifest files**

Snyk.io fetches and views dependency manifest files to analyze and deliver vulnerability results. To get proper test results and to create Snyk projects, one or more supported manifest files must be present in the tested folder \(for integration with CLI\), or in the repository \(for integration with Git\).

See [Language support](https://support.snyk.io/hc/en-us/categories/360000456257-Language-support) for details of supported manifest files for different languages.

### **Broker usage**

When set up, developers can use Snyk Broker to enable standard Snyk product usage \(such as Snyk Open Source\), with the Broker validating all in / outbound requests, based on the approved list.

