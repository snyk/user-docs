# Enable permissions for Snyk Broker from your third-party tool

**Assign permissions components**

Assign permissions based on your integration as follows:

* [**GitHub / GitHub Enterprise**](https://github.com/settings/tokens): see [GitHub integration](https://docs.snyk.io/integrations/git-repository-scm-integrations/github-integration).
* [**Bitbucket server**](https://confluence.atlassian.com/bitbucket/grant-repository-access-to-users-and-groups-221449716.html)**:** see [Bitbucket Data Center/Server integration](../git-repository-scm-integrations/bitbucket-data-center-server-integration.md).&#x20;
* [**GitLab**](https://docs.gitlab.com/ee/user/profile/personal\_access\_tokens.html): see [GitLab integration](https://docs.snyk.io/integrations/git-repository-scm-integrations/gitlab-integration).
* [**Azure Repos**](https://docs.microsoft.com/en-us/azure/devops/repos/): see [Azure Repos integration](https://docs.snyk.io/integrations/git-repository-scm-integrations/azure-repos-integration).
* [**Jira**](https://confluence.atlassian.com/cloud/api-tokens-938839638.html): Snyk needs user credentials with API access. See [Jira integration](https://docs.snyk.io/integrations/untitled-3/jira).&#x20;

Assign permissions to detect Infrastructure as Code files as follows:

* [**Terraform**](https://docs.snyk.io/snyk-infrastructure-as-code/scan-terraform-files/detecting-terraform-configuration-files-using-a-broker)
* [**CloudFormation**](https://docs.snyk.io/snyk-infrastructure-as-code/scan-cloudformation-files/detecting-cloudformation-configuration-files-using-a-broker)
* [**Kubernetes and other IaC**](https://docs.snyk.io/snyk-infrastructure-as-code/detecting-infrastructure-as-code-files-using-a-broker)

Add Snyk Container local Container Registry support through Snyk Broker

* ****[**Snyk Container Broker Agent**](https://docs.snyk.io/products/snyk-container/integrate-self-hosted-container-registries/snyk-integration-to-self-hosted-container-registries) setup
* ****[**Brokered ECR**](https://docs.snyk.io/products/snyk-container/integrate-self-hosted-container-registries/setting-up-the-container-registry-agent-for-a-brokered-ecr-integration) setup

Add Snyk Code local Git Repository support through Snyk Broker:

* ****[**Snyk Code Broker Agent**](https://docs.snyk.io/products/snyk-code/snyk-code-local-git-support) setup

## Generate credentials in the target application for Snyk Broker

After generating the credentials for the Broker's target application, configure the environment variables for launching the Broker.
