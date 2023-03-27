# Install and configure Broker using Docker

## Snyk Broker Docker Installation

Installing the Snyk Broker through Docker commands is also straightforward. Follow the instructions for the different integrations:

* [GitHub](../../../snyk-broker-set-up-examples/broker-example-set-up-snyk-broker-with-github.md)
* [GitHub Enterprise](../../../snyk-broker-set-up-examples/setup-broker-with-github-enterprise.md)
* [Bitbucket Server/Data Centre](../../../snyk-broker-set-up-examples/data-center.md)
* [Gitlab](../../../snyk-broker-set-up-examples/setup-broker-with-gitlab.md)
* [Azure Repos](../../../snyk-broker-set-up-examples/setup-broker-with-azure-repos.md)
* [JFrog Artifactory](https://github.com/snyk/broker#artifactory)
* [Nexus](https://github.com/snyk/broker#nexus-3)
* [Jira](../../../snyk-broker-set-up-examples/setup-broker-with-jira.md)
* [Snyk Broker - Container Registry Agent](../../../snyk-broker-container-registry-agent/) (needed to connect to Container Registries)
* [Snyk Broker - Code Agent](../../../snyk-broker-code-agent/) (needed to enable SAST analysis)

Once the Broker is running, visit the settings for your brokered integration in [the Snyk Web UI](https://app.snyk.io) to see a message like the following and start importing projects:

<figure><img src="../../../../../.gitbook/assets/image (60) (2).png" alt="Brokered integration settings showing connected status"><figcaption><p>Brokered integration settings showing connected status</p></figcaption></figure>
