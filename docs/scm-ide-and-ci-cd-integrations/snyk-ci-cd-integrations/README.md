# Snyk CI/CDs

{% hint style="info" %}
Snyk recommends using the [CLI ](../../cli-ide-and-ci-cd-integrations/snyk-cli/)for CI/CD integrations for the following reasons:

* You have the flexibility to test in-progress features of the CLI by using the [preview channel](../../snyk-cli/releases-and-channels-for-the-snyk-cli.md#preview).
* The CLI provides feature-rich [stable releases](../../snyk-cli/releases-and-channels-for-the-snyk-cli.md#stable) at a regular cadence.
* With the CLI, you have options to [extend use cases](../../cli-ide-and-ci-cd-integrations/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/) as you deploy Snyk at scale

For more information, see the  [Snyk CLI repository](https://github.com/snyk/cli).
{% endhint %}

## Adopting a CI/CD integration

When you decide to use a CI/CD integration, you typically adopt the integration in stages, initially selecting a deployment method and the implementing strategies for the code you are scanning. For details, see [Snyk CI/CD Integration deployment and strategies](../../cli-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/).

For detailed information, see the pages for the integration you are using:

* [AWS CodePipeline integration with CodeBuild](../../cli-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/aws-codepipeline-integration-by-adding-a-snyk-scan-stage.md)
* [Azure Pipelines integration](../../cli-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/azure-pipelines-integration/)
* [Bitbucket Pipelines integration](../../cli-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/bitbucket-pipelines-integration-using-a-snyk-pipe/)
* [CircleCI integration](circleci-integration-using-a-snyk-orb.md)
* [GitHub Actions integration](github-actions-for-snyk-setup-and-checking-for-vulnerabilities/)
* [Jenkins integration](jenkins-plugin-integration-with-snyk.md)
* [Maven integration](maven-plugin-integration-with-snyk.md)
* [TeamCity integration](../../cli-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/teamcity-jetbrains-integration-using-the-snyk-security-plugin/)
* [Terraform Cloud integration for IaC](../../cli-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/terraform-cloud-integration-for-snyk-iac-using-run-tasks/)
* [Terraform Enterprise integration for IaC](../../cli-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/terraform-enterprise-integration-for-snyk-iac.md)

For integration with GitLab pipelines, see this [pipeline configuration](https://github.com/snyk-labs/snyk-cicd-integration-examples/blob/master/GitLabCICD/gitlab-npm.yml).

For additional examples of binary and npm integrations for CI/CD, see [GitHub CI/CD examples](https://github.com/snyk-labs/snyk-cicd-integration-examples).

## Support policy <a href="#support-policy" id="support-policy"></a>

Snyk supports the latest 12 months of CI/CD plugin versions, ensuring functionality and performance. Older versions are considered End-of-Support (EOS) and will not receive bug fixes or troubleshooting.

Snyk provides fixes only in new versions and cannot fix older versions. Customers must upgrade to benefit from improvements.

This policy fosters innovation while optimizing resources.

If you need help, submit a [request](https://support.snyk.io/) to Snyk Support.
