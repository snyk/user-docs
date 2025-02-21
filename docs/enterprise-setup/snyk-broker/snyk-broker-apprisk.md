# Snyk Broker - Essentials

## Prerequisites

If your SCM instance is not publicly accessible, you need Snyk Broker. You can install and configure Snyk Broker using Docker or Helm. The minimum supported Broker version for Snyk Essentials is [4.171.0](https://github.com/snyk/broker/releases/tag/v4.171.0).

Ensure you have the Snyk Broker token for the Snyk Essentials integration. The Snyk support team can provide the needed token, or you can generate it yourself by following these instructions:

* Generate your Broker token by following the instructions from the [Obtain your Broker token for Snyk Broker ](snyk-broker-code-agent/install-snyk-broker-code-agent-using-docker/obtain-the-required-tokens-for-setup.md#obtain-your-broker-token-from-the-web-ui)page.&#x20;
* Copy and paste the Broker token on the integration setup menu from the Integration Hub.

## SCM integrations

* GitHub - install and configure Snyk Broker&#x20;
  * [using Docker](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-docker.md#docker-run-command-to-set-up-a-broker-client-for-github)
  * [using Helm](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-environment-variables-for-snyk-broker.md)
* GitHub Enterprise - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-docker.md#docker-run-command-to-set-up-a-broker-client-for-github-enterprise)
  * [using Helm](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-environment-variables-for-snyk-broker.md)
* BitBucket - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md#docker-run-command-to-set-up-a-broker-client-for-bitbucket)
  * [using Helm](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker-basic-auth.md)
* GitLab - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-gitlab.md#docker-run-command-to-set-up-a-broker-client-for-gitlab)
  * [using Helm](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-environment-variables-for-snyk-broker.md)
* Azure - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md#docker-run-command-to-set-up-a-broker-client-for-azure-repos)
  * [using Helm](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-install-and-configure-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-environment-variables-for-snyk-broker.md)

You can find on [GitHub](https://github.com/snyk/broker/tree/565242baf003f06f445489dd96cc68c8386ede38/defaultFilters/apprisk) all the updated `.json` files that include the allowed list of accessible endpoints for the integrations.

As the final step following a completed Broker setup, for any brokered integration you want to add to Snyk Essentials, you must obtain the Broker token. You can find the Broker token in your Organization integration general settings for your integration type, for example, GitHub, GitLab, and so on, as shown in the following image. For third-party integrations, see the next section.

<figure><img src="../../.gitbook/assets/Screenshot 2024-11-01 at 1.05.55 PM.png" alt="&#x22;&#x22;"><figcaption><p>Broker token in GitLab integration general settings</p></figcaption></figure>



