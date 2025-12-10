# Snyk Broker

{% hint style="info" %}
**Feature availability**

Snyk Broker is available only for Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Snyk Broker is an open-source tool that acts as a proxy between Snyk and source integrations, allowing for access by [snyk.io](http://snyk.io/) to scan your code and return results to you. The Broker allows Snyk to connect to remote resources in private repositories, leaving credentials inside the customer's network.

The diagram that follows illustrates the basic components.

<figure><img src="../../../.gitbook/assets/Snyk Broker diagram.png" alt="Snyk Broker WebSocket initiated by Client over HTTPS"><figcaption><p>Snyk Broker WebSocket initiated by Client over HTTPS</p></figcaption></figure>

## How Snyk Broker works

Snyk Broker includes a Server and a Client, the basic components that are the same across all integrations. The Broker Server runs on the Snyk SaaS backend and is provided by Snyk; no installation is required. You will install the Broker client and deploy it in your infrastructure.

The Broker client and server act together, sending requests by proxy from [snyk.io](http://snyk.io/) to a repository or Jira, fetching the files needed for scanning from repositories and fetching results using webhooks posted by the SCM service.

The Broker client runs within your internal network, keeping sensitive data such as SCM tokens within the network perimeter. The Broker connection allows for scanning using only requests on an approved data list. This narrows the access permissions to the absolute minimum required for Snyk to monitor a repository. For more information, see [Approved data list for Snyk Broker](broker-inbound-and-outbound-connections-and-allowed-requests.md#approved-data-list-for-snyk-broker).

Using Snyk Broker allows you to manage a fixed private IP for your integration that targets the Broker.

All data, both in transit and at rest, is encrypted. There is no need to open incoming ports; the communication is initiated outbound. After the connection is initiated, the secure WebSocket connection is bidirectional between the Client and the Server.

SCM integrations with Broker Support Snyk Code, Open Source, Container, IaC, and Essentials.

## Integrations supported by Snyk Broker

| Integration                        | Classic Broker support               | Universal Broker support             |
| ---------------------------------- | ------------------------------------ | ------------------------------------ |
| Artifactory Private Registry       | Yes                                  | Yes                                  |
| Azure Repository                   | Yes                                  | Yes                                  |
| Bitbucket Server                   | Yes                                  | Yes                                  |
| GitHub                             | Yes                                  | Yes                                  |
| GitHub Enterprise                  | Yes                                  | Yes                                  |
| GithHub Cloud app                  | No                                   | Yes                                  |
| GitHub Server app                  | No                                   | Yes                                  |
| GitLab                             | Yes                                  | Yes                                  |
| Jira                               | Yes                                  | Yes                                  |
| Nexus Private Registry             | Yes                                  | Yes                                  |
| Docker Hub                         | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |
| Elastic Container Registry         | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |
| Azure Container Registry           | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |
| Google Container Registry          | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |
| Artifactory Container Registry     | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |
| Harbor                             | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |
| Quay                               | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |
| GitHub Container Registry          | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |
| Nexus Container Registry           | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |
| DigitalOcean Container Registry    | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |
| GitLab Container Registry          | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |
| Google Artifact Container Registry | Yes (using Container Registry Agent) | Yes (using Container Registry Agent) |

## Using Universal Broker versus Classic Broker

The Universal Broker builds on the technology of the classic Snyk Broker to bring easier and more scalable configuration, enhanced security, and new capabilities. The aim is for the Universal Broker to replace the Classic Broker entirely.

The following compares the capabilities and features of the Classic Broker to those of the Universal Broker.

| Capabilities and features           | Classic Broker           | Universal Broker                      |
| ----------------------------------- | ------------------------ | ------------------------------------- |
| Deployment                          | Container and Helm chart | Container and Helm chart              |
| Connection parameters configuration | local                    | cloud                                 |
| Credentials                         | local                    | local                                 |
| Connection support                  | single (dedicated type)  | multiple (any type)                   |
| Configuration management            | none                     | tooling (`snyk-broker-config`) or API |
| Organization to connection mapping  | no                       | yes                                   |
| API                                 | none                     | yes                                   |
