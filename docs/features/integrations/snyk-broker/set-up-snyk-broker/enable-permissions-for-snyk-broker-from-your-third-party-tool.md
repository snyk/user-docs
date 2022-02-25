# Enable permissions for Snyk Broker from your third-party tool

### **Assign permissions components**

Assign permissions based on your integration as follows:

* [**GitHub / GitHub Enterprise**](https://github.com/settings/tokens): see [github-integration.md](../../../../integrations/git-repository-scm-integrations/github-integration.md "mention").
* [**Bitbucket server**](https://confluence.atlassian.com/bitbucket/grant-repository-access-to-users-and-groups-221449716.html)**:** see [bitbucket-data-center-server-integration.md](../../../../integrations/git-repository-scm-integrations/bitbucket-data-center-server-integration.md "mention").
* [**GitLab**](https://docs.gitlab.com/ee/user/profile/personal\_access\_tokens.html): see [gitlab-integration.md](../../../../integrations/git-repository-scm-integrations/gitlab-integration.md "mention").
* [**Azure Repos**](https://docs.microsoft.com/en-us/azure/devops/repos/): see [azure-repos-integration.md](../../../../integrations/git-repository-scm-integrations/azure-repos-integration.md "mention").
* [**Jira**](https://confluence.atlassian.com/cloud/api-tokens-938839638.html): Snyk needs user credentials with API access. See [jira-integration](../../../../products/snyk-infrastructure-as-code/jira-integration/ "mention").

You can assign permissions to detect Infrastructure as Code files: see [snyk-broker-infrastructure-as-code-detection](../snyk-broker-infrastructure-as-code-detection/ "mention").

You can add Snyk Container local Container Registry support through Snyk Broker: see [snyk-broker-container-registry-agent](../snyk-broker-container-registry-agent/ "mention").

You can add Snyk Code local Git Repository support through Snyk Broker: see [snyk-broker-code-agent.md](../snyk-broker-code-agent.md "mention").

### Generate credentials in the target application for Snyk Broker

After generating the credentials for the Broker's target application, configure the environment variables for launching the Broker.
