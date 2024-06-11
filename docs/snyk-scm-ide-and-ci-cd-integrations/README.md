# Snyk SCM, IDE, and CI/CD integrations

Snyk offers multiple integration types that you can utilize to enhance your current toolset.

When starting with Snyk, we recommend deciding on which primary integrations to implement:

* [SCM (Source Control Manager) integrations](git-repositories-scms-integrations-with-snyk/)
* [IDEs](use-snyk-in-your-ide/)
* [CI/CDs](snyk-ci-cd-integrations/)

As well as how to access Snyk functionality in your developer workflow:

* [Snyk Web UI](../getting-started/snyk-web-ui.md)
* [Snyk AppRisk](../manage-risk/snyk-apprisk/)
* [Snyk CLI](../snyk-cli/)
* [Snyk API](../snyk-api/)

{% hint style="warning" %}
Enterprise plan users have base-level access to all of the above functionality. Free and Team plan users do not have access to the Snyk API or Snyk AppRisk. See [Plans and pricing](https://snyk.io/plans/) for more information.
{% endhint %}

If you are an Enterprise customer, see [Choose rollout integrations](../implement-snyk/team-implementation-guide/phase-1-discovery-and-planning/choose-rollout-integrations.md) in the Enterprise implementation guide for tips and considerations on import strategies. We then suggest reading through our [Deployment recommendations](./#deployment-recommendations-for-git-integrations), so you have the smoothest rollout for your teams.

Snyk can automatically create pull requests (PRs) on your behalf to upgrade your dependencies based on scan results. This is compatible with a variety of Snyk integrations. For more information, see [View and understand Snyk upgrade pull requests for integrations](../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/introduction-to-git-repository-integrations/view-and-understand-snyk-upgrade-pull-requests.md).

## SCM integrations

Snyk offers the following Git integrations:

* [GitHub Cloud App](git-repositories-scms-integrations-with-snyk/snyk-github-cloud-app.md)
* [GitHub Enterprise](git-repositories-scms-integrations-with-snyk/snyk-github-enterprise-integration.md)
* [GitHub](git-repositories-scms-integrations-with-snyk/snyk-github-integration.md)
* [GitHub Read-only Projects](git-repositories-scms-integrations-with-snyk/snyk-github-read-only-projects.md)
* [GitLab](git-repositories-scms-integrations-with-snyk/snyk-gitlab-integration.md)
* [Bitbucket Cloud](git-repositories-scms-integrations-with-snyk/snyk-bitbucket-cloud-integration.md)
* [Bitbucket Cloud (Legacy)](git-repositories-scms-integrations-with-snyk/migrate-a-bitbucket-cloud-personal-access-token.md)
* [Bitbucket Cloud App](git-repositories-scms-integrations-with-snyk/snyk-bitbucket-cloud-app-integration.md)
* [Bitbucket Data Center/Server](git-repositories-scms-integrations-with-snyk/snyk-bitbucket-data-center-server-integration.md)
* [Azure Repositories (TFS)](git-repositories-scms-integrations-with-snyk/snyk-azure-repositories-tfs-integration.md)

Snyk requires certain permissions and access scopes for SCM integrations to ensure Snyk functionality works seamlessly. For more information, see [User permissions and access scopes for SCM integrations](git-repositories-scms-integrations-with-snyk/#user-permissions-and-access-scopes-for-scm-integrations).

## Snyk CI/CDs

Snyk offers the following CI/CDs:

* [AWS CodePipeline](snyk-ci-cd-integrations/aws-codepipeline-integration-by-adding-a-snyk-scan-stage/)
* [Azure Pipelines](snyk-ci-cd-integrations/azure-pipelines-integration/)
* [Bitbucket Pipelines](snyk-ci-cd-integrations/bitbucket-pipelines-integration-using-a-snyk-pipe/)
* [CircleCI](snyk-ci-cd-integrations/circleci-integration-using-a-snyk-orb.md)
* [GitHub Actions for Snyk](snyk-ci-cd-integrations/github-actions-for-snyk-setup-and-checking-for-vulnerabilities/)
* [Jenkins](snyk-ci-cd-integrations/jenkins-plugin-integration-with-snyk.md)
* [Maven](snyk-ci-cd-integrations/maven-plugin-integration-with-snyk.md)
* [TeamCity (JetBrains)](snyk-ci-cd-integrations/teamcity-jetbrains-integration-using-the-snyk-security-plugin/)
* [Terraform Cloud](snyk-ci-cd-integrations/terraform-cloud-integration-for-snyk-iac-using-run-tasks/)
* [Terraform Enterprise](snyk-ci-cd-integrations/terraform-enterprise-integration-for-snyk-iac.md)

## IDEs

Snyk offers the following IDEs:

* [Eclipse](use-snyk-in-your-ide/eclipse-plugin/)
* [JetBrains](use-snyk-in-your-ide/jetbrains-plugins/)
* [Visual Studio](use-snyk-in-your-ide/visual-studio-extension/)
* [Visual Studio Code](use-snyk-in-your-ide/visual-studio-code-extension/)
* [Snyk Language Server](use-snyk-in-your-ide/snyk-language-server.md)

## Deployment recommendations for SCM integrations

If you try to implement all the SCM integration features at the same time, you risk causing friction in your software development life cycle ([SDLC](https://snyk.io/learn/secure-sdlc/)), which in turn leads to a poor developer experience.

To ensure a smooth rollout of Snyk across your organization, Snyk provides a suggested deployment timeline consisting of deployment stages, configuration steps, and the desired outcome for each stage.

For detailed steps, see [Deployment recommendations for SCM integrations](./#deployment-recommendations-for-git-integrations).
