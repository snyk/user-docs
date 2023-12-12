# Integrations for Snyk AppRisk

## Overview

The Integrations page shows all active integrations, including any data automatically synced from your existing Snyk Organizations, and provides access to the Integration Hub.

The supported Snyk organizations that are automatically synced are the following:

* Snyk Open Source
* Snyk Code
* Snyk IaC
* Snyk Container&#x20;

Each connected integration enables you to:

* Pause data syncing.
* Modify integration profiles and configurations.
* Delete the integration.
* Check when the integration was last synced and when the next sync is scheduled.&#x20;

## Using the Integration Hub

Use the Integration Hub page to onboard integrations and populate Snyk AppRisk with data from SCM tools.

You can add an integration by following these steps:

1. Open AppRisk and navigate to the Integrations page.
2. Click **Integration Hub**.
3. Click **Add** on the integration you want to connect.
4. Configure your connection and click **Done**.

Use the [Connect an SCM integration](connect-an-scm-integration.md) page for step-by-step details about how to setup an integration.

After the integration is validated, a card is displayed on the Integrations page, allowing you to enable or disable the connection, edit the settings, or remove it from your configuration.

<figure><img src="../../../.gitbook/assets/image (11) (4).png" alt="AppRisk - Integration status" width="375"><figcaption><p>Snyk AppRisk - Integration status</p></figcaption></figure>

## Using Snyk Broker

You can install and configure your integration by using the Snyk Broker with Docker, or Helm.

{% hint style="info" %}
Enable the Snyk AppRisk flag in your Snyk Broker deployment environment before running the commands.
{% endhint %}

* GitHub - install and configure Snyk Broker&#x20;
  * [using Doker](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-install-and-configure-broker/broker-example-set-up-snyk-broker-with-github.md#docker-run-command-to-set-up-a-broker-client-for-github)
  * [using Helm](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-install-and-configure-broker/githhub.com-install-and-configure-using-helm.md)
  * [environment variables](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-install-and-configure-broker/github-environment-variables-for-snyk-broker.md)
* GitHub Enterprise - install and configure Snyk Broker&#x20;
  * [using Docker](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-enterprise-install-and-configure-broker/setup-broker-with-github-enterprise.md#docker-run-command-to-set-up-a-broker-client-for-github-enterprise)
  * [using Helm](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-enterprise-install-and-configure-broker/github-enterprise-install-and-configure-using-helm.md)
  * [environment variables](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-enterprise-install-and-configure-broker/github-enterprise-environment-variables-for-snyk-broker.md)
* BitBucket - install and configure Snyk Broker&#x20;
  * [using Docker](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-install-and-configure-broker/data-center.md#docker-run-command-to-set-up-a-broker-client-for-bitbucket)
  * [using Helm](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-helm.md)
  * [environment variables](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker.md)
* GitLab - install and configure Snyk Broker&#x20;
  * [using Docker](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/gitlab-install-and-configure-broker/setup-broker-with-gitlab.md#docker-run-command-to-set-up-a-broker-client-for-gitlab)
  * [using Helm](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/gitlab-install-and-configure-broker/gitlab-install-and-configure-using-helm.md)
  * [environment variables](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/gitlab-install-and-configure-broker/gitlab-environment-variables-for-snyk-broker.md)
* Azure - install and configure Snyk Broker&#x20;
  * [using Docker](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/azure-repos-install-and-configure-broker/setup-broker-with-azure-repos.md#docker-run-command-to-set-up-a-broker-client-for-azure-repos)
  * [using Helm](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/azure-repos-install-and-configure-broker/azure-repos-install-and-configure-and-configure-using-helm.md)
  * [environment variables](../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/azure-repos-install-and-configure-broker/azure-repos-environment-variables-for-snyk-broker.md)

You can find on [GitHub](https://github.com/snyk/broker/tree/565242baf003f06f445489dd96cc68c8386ede38/defaultFilters/apprisk) all the updated .json files that include the allowed list of accessible endpoints by the integrations.

