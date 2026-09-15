---
nav_context: agnostic
description: >-
  How to integrate Snyk into CI/CD pipelines, with the Snyk CLI recommended for
  flexibility
---

# Snyk CI/CDs

{% hint style="info" %}
Snyk recommends using the CLI for CI/CD integrations for the following reasons:

* You have the flexibility to test in-progress features of the CLI by using the preview channel.
* The CLI provides feature-rich stable releases at a regular cadence.
* With the CLI, you have options to extend use cases as you deploy Snyk at scale

For more information, see the [Snyk CLI repository](https://github.com/snyk/cli).
{% endhint %}

{% hint style="info" %}
**Snyk Secrets support**

Snyk Secrets does not support CI/CD plugin integrations. For CI/CD-time secret scanning, use [Secrets scanning in the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-secrets/secrets-scanning-in-the-snyk-cli) directly in your pipeline.
{% endhint %}

## Adopting a CI/CD integration

When you decide to use a CI/CD integration, you typically adopt the integration in stages, initially selecting a deployment method and the implementing strategies for the code you are scanning. For details, see [Snyk CI/CD Integration deployment and strategies](https://docs.snyk.io/developer-tools/integrations/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies).

For detailed information, see the pages for the integration you are using:

* [AWS CodePipeline integration with CodeBuild](aws-codepipeline-integration-by-adding-a-snyk-scan-stage.md)
* [Azure Pipelines integration](azure-pipelines-integration/)
* [Bitbucket Pipelines integration](bitbucket-pipelines-integration-using-a-snyk-pipe/bitbucket-pipelines-integration-how-it-works.md)
* [CircleCI integration](circleci-integration-using-a-snyk-orb.md)
* [GitHub Actions integration](github-actions-for-snyk-setup-and-checking-for-vulnerabilities/)
* [Jenkins integration](jenkins-plugin-integration-with-snyk.md)
* [Maven integration](github-actions-for-snyk-setup-and-checking-for-vulnerabilities/snyk-maven-action.md)
* [TeamCity integration](https://docs.snyk.io/developer-tools/integrations/snyk-ci-cd-integrations/teamcity-jetbrains-integration-using-the-snyk-security-plugin)
* [Terraform Cloud integration for IaC](terraform-cloud-integration-for-snyk-iac-using-run-tasks/set-up-the-terraform-cloud-integration-for-iac.md)
* [Terraform Enterprise integration for IaC](terraform-enterprise-integration-for-snyk-iac.md)

For integration with GitLab pipelines, see this [pipeline configuration](https://github.com/snyk-labs/snyk-cicd-integration-examples/blob/master/GitLabCICD/gitlab-npm.yml).

For additional examples of binary and npm integrations for CI/CD, see [GitHub CI/CD examples](https://github.com/snyk-labs/snyk-cicd-integration-examples).

## Support policy <a href="#support-policy" id="support-policy"></a>

Snyk supports the latest 12 months of CI/CD plugin versions, ensuring functionality and performance. Older versions are considered End-of-Support (EOS) and will not receive bug fixes or troubleshooting.

Snyk provides fixes only in new versions and cannot fix older versions. Customers must upgrade to benefit from improvements.

This policy fosters innovation while optimizing resources.

If you need help, submit a [request](https://support.snyk.io/) to Snyk Support.
