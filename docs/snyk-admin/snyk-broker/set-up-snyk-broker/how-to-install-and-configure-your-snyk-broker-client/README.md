# Install, configure, and upgrade the Snyk Broker Client

{% hint style="info" %}
Snyk recommends using Helm as the simplest way to deploy Snyk Broker. You can also use Docker to run the Snyk Broker Client or run npm `install snyk-broker`. This documentation provides detailed instructions for using Helm and Docker.
{% endhint %}

## Snyk Broker Docker Installation

Installing the Snyk Broker through Docker commands is also straightforward. Follow the instructions for the different integrations:

* [GitHub](../../snyk-broker-set-up-examples/broker-example-set-up-snyk-broker-with-github.md)
* [GitHub Enterprise](../../snyk-broker-set-up-examples/setup-broker-with-github-enterprise.md)
* [Bitbucket Server/Data Centre](../../snyk-broker-set-up-examples/data-center.md)
* [Gitlab](../../snyk-broker-set-up-examples/setup-broker-with-gitlab.md)
* [Azure Repos](../../snyk-broker-set-up-examples/setup-broker-with-azure-repos.md)
* [JFrog Artifactory](https://github.com/snyk/broker#artifactory)&#x20;
* [Nexus](https://github.com/snyk/broker#nexus-3)
* [Jira](../../snyk-broker-set-up-examples/setup-broker-with-jira.md)
* [Snyk Broker - Container Registry Agent](../../snyk-broker-container-registry-agent/) (needed to connect to Container Registries)
* [Snyk Broker - Code Agent](../../snyk-broker-code-agent/) (needed to enable SAST analysis)

Once the Broker is running, visit the settings for your brokered integration in [the Snyk Web UI](https://app.snyk.io) to see a message like the following and start importing projects:

<figure><img src="../../../../.gitbook/assets/image (60) (1).png" alt="Brokered integration settings showing connected status"><figcaption><p>Brokered integration settings showing connected status</p></figcaption></figure>

## Upgrade your Snyk Broker Client

Snyk regularly updates the Broker Client in order to provide new features, bug fixes, and more. The full list of versions with release notes is available on [Snyk Broker GitHub](https://github.com/snyk/broker/releases). Snyk encourages you to [subscribe to the RSS feed](https://github.com/snyk/broker/releases.atom) for that page to receive information about versions as they are released.

When you upgrade your Broker there may be some new rules added that Snyk requires to function correctly. Therefore, you will need to re-initialize your API allow-list. If you added or removed any rules to [customize your allow-list](https://github.com/snyk/broker#custom-approved-listing-filter) (for example, to support files greater in size than 1Mb), you must  re-apply these changes to the new allow-list.
