# Snyk CI/CD integrations

{% hint style="info" %}
Snyk recommends using the [Snyk CLI ](https://github.com/snyk/cli)for CI/CD integrations for the following reasons:

* You have the flexibility to test in-progress features of the CLI by using the [preview channel](../../snyk-cli/releases-and-channels-for-the-snyk-cli.md#preview).
* The CLI provides feature-rich [stable releases](../../snyk-cli/releases-and-channels-for-the-snyk-cli.md#stable) at a regular cadence.
* With the CLI, you have options to [extend use cases](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/) as you deploy Snyk at scale
{% endhint %}

When you [decide to use a CI/CD Integration,](../git-repository-and-ci-cd-integrations-comparisons.md) you typically adopt the integration in stages. You will select a deployment method and implement strategies for the code you are scanning. See [Snyk CI/CD Integration deployment and strategies](snyk-ci-cd-integration-deployment-and-strategies/).

For detailed information, you can refer to the pages for the integration you are using:

* [AWS CodePipeline integration](aws-codepipeline-integration-by-adding-a-snyk-scan-stage/)
* [Azure Pipelines integration](azure-pipelines-integration/)
* [Bitbucket Pipelines integration](bitbucket-pipelines-integration-using-a-snyk-pipe/)
* [CircleCI integration](circleci-integration-using-a-snyk-orb.md)
* [GitHub Actions integration](github-actions-for-snyk-setup-and-checking-for-vulnerabilities/)
* [Jenkins integration](jenkins-plugin-integration-with-snyk.md)
* [Maven integration](maven-plugin-integration-with-snyk.md)
* [TeamCity integration](teamcity-jetbrains-integration-using-the-snyk-security-plugin/)
* [Terraform Cloud integration for IaC](terraform-cloud-integration-for-snyk-iac-using-run-tasks/)
* [Terraform Enterprise integration for IaC](terraform-enterprise-integration-for-snyk-iac.md)

For integration with GitLab pipelines integration, see this [pipeline configuration](https://github.com/snyk-labs/snyk-cicd-integration-examples/blob/master/GitLabCICD/gitlab-npm.yml).

For additional examples of binary and npm integrations for CI/CD, see [GitHub CI/CD examples](https://github.com/snyk-labs/snyk-cicd-integration-examples).
