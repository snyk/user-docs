# Snyk Broker - AppRisk

If your SCM instance is not publicly accessible, you need Snyk Broker. You can install and configure your Snyk Broker using Docker or Helm. The minimum supported Broker version for Snyk AppRisk is [4.171.0](https://github.com/snyk/broker/releases/tag/v4.171.0).

{% hint style="info" %}
Enable the Snyk AppRisk flag in your Snyk Broker deployment environment before running the commands.
{% endhint %}

* GitHub - install and configure Snyk Broker&#x20;
  * [using Docker](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/broker-example-set-up-snyk-broker-with-github.md#docker-run-command-to-set-up-a-broker-client-for-github)
  * [using Helm](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/githhub.com-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-environment-variables-for-snyk-broker.md)
* GitHub Enterprise - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-github-enterprise.md#docker-run-command-to-set-up-a-broker-client-for-github-enterprise)
  * [using Helm](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-environment-variables-for-snyk-broker.md)
* BitBucket - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md#docker-run-command-to-set-up-a-broker-client-for-bitbucket)
  * [using Helm](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker.md)
* GitLab - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-gitlab.md#docker-run-command-to-set-up-a-broker-client-for-gitlab)
  * [using Helm](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-environment-variables-for-snyk-broker.md)
* Azure - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md#docker-run-command-to-set-up-a-broker-client-for-azure-repos)
  * [using Helm](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-install-and-configure-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-environment-variables-for-snyk-broker.md)

You can find on [GitHub](https://github.com/snyk/broker/tree/565242baf003f06f445489dd96cc68c8386ede38/defaultFilters/apprisk) all the updated `.json` files that include the allowed list of accessible endpoints for the integrations.
