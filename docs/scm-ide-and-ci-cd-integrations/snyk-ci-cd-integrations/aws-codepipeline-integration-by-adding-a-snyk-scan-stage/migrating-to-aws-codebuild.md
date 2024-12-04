# Migrating to AWS CodeBuild

{% hint style="warning" %}
**The Snyk integration for AWS CodePipeline will be discontinued**

\
**Action Required**

In order to safeguard the security of our services and our customers, Snyk has begun the deprecation of its integration with **AWS CodePipeline**. To minimize disruption, we recommend that you transition to using **AWS CodeBuild** and the Snyk CLI as an alternative which will support the same use case and functionality.&#x20;



**Migration Timeline**

Effective **Oct 30th, 2024**, you will no longer be able to add or modify the Snyk plug-in for new or existing pipelines. Existing pipelines will continue to work as-is for 6 months, though we recommend migrating to the new process as soon as possible. To avoid disrupting your CI/CD workflows, you must transition to the Snyk CLI before **April 30, 2025**. Please refer to the steps in the below migration guide to use Snyk CLI with AWS CodeBuild.
{% endhint %}

This guide outlines the steps for migrating your [Snyk Open Source](https://snyk.io/product/open-source-security-management/) security scanning workflow from the [Snyk and AWS CodePipeline integration](./) to [AWS CodeBuild](https://aws.amazon.com/codebuild/). By using the Snyk CLI and the built-in capabilities of CodeBuild, you can achieve a more streamlined and configurable solution for running Snyk software composition analysis (SCA) scans in your CI/CD pipeline.

## Migration goal

* **Current Setup**: Your workflow uses the Snyk CodePipeline plugin in a dedicated stage in CodePipeline.
* **Target Setup:** Snyk scanning is performed in a custom CodeBuild build step. This build step leverages the Snyk CLI directly to execute the scan and integrate the results into your pipeline.

## Prerequisites

* An active AWS account with CodeBuild and CodePipeline services enabled
* A Snyk account with the Snyk CLI configured
* Familiarity with CodeBuild project configuration and environment variables
* Understanding of your existing CodePipeline stages and their interaction with Snyk

## Migration steps

Follow the steps in these sections to migrate your [Snyk Open Source](https://snyk.io/product/open-source-security-management/) security scanning workflow from the [Snyk and AWS CodePipeline integration](./) to [AWS CodeBuild](https://aws.amazon.com/codebuild/)

### Set up CodeBuild

* Create a new CodeBuild project in your AWS account.
* Choose a compatible [base image](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html) for your project based on your programming language and dependencies.
* Review how to [authenticate the Snyk CLI with your account](../../../snyk-cli/authenticate-to-use-the-cli.md) and consider using an environment variable to store sensitive information such as your Snyk CLI token.

{% hint style="info" %}
The default **Service role** in AWS CodeBuild includes an IAM permission that allows the CodeBuild project to pull any secret from AWS Secrets Manager that starts with `/CodeBuild/` in the name. Refer to the [Troubleshooting](migrating-to-aws-codebuild.md#troubleshooting) section at the end of this guide for more information.
{% endhint %}

* Configure build commands:
  * [Install the Snyk CLI](../../../snyk-cli/install-or-update-the-snyk-cli/) using the commands appropriate for your operating system.
  * Define a build command that executes the Snyk scan using the CLI.
  * Define a build command that sends a snapshot of the project to Snyk for continuous monitoring (optional).
* Review the example `buildspec.yaml` that follows for more details:

```yaml
# buildspec.yaml
version: 0.2

phases:
  build:
    commands:
      # install the latest Snyk CLI from GitHub Releases
      - latest_version=$(curl -Is "https://github.com/snyk/cli/releases/latest" | grep "^location" | sed 's#.*tag/##g' | tr -d "\r")
      - snyk_cli_dl_linux="https://github.com/snyk/cli/releases/download/${latest_version}/snyk-linux"
      - curl -Lo /usr/local/bin/snyk $snyk_cli_dl_linux
      - chmod +x /usr/local/bin/snyk
      
      # authenticate the Snyk CLI
      - snyk auth $SNYK_TOKEN
      
      # perform a Snyk SCA scan; continue if vulnerabilities are found
      - snyk test || true
      
      # upload a snapshot of the project to Snyk for continuous monitoring
      - snyk monitor
```

### Set up CodePipeline

For some [Open Source](https://snyk.io/product/open-source-security-management/) Projects, you must build the Project before testing it with the Snyk CLI. Review the Snyk [documentation](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md) to determine whether Snyk requires your Project to be built before running an Open Source scan; then follow the instructions in the corresponding section below:

#### Snyk requires a built Project

* Edit your existing CodePipeline or create a new one.
* Create a new stage to build your Project, or edit the existing build stage.
* Add the commands from the example `buildspec.yaml` to your build stage so that the Snyk scan occurs immediately after the Project is built.

{% hint style="info" %}
The Snyk Open Source scan must be in the same CodeBuild action as the build process to ensure that Snyk has access to the full build workspace.
{% endhint %}

#### Snyk does not require a built Project

* Edit your existing CodePipeline or create a new one.
* Add a new build stage after your source code acquisition stage.
* Select your newly created CodeBuild Project for this stage.
* Select SourceArtifact under Input artifacts to allow Snyk to scan the source code directly.

### Result handling

The Snyk integration for CodePipeline only supports a limited set of [Snyk CLI](../../../snyk-cli/commands/) functionality and options. By using the Snyk CLI in CodeBuild, you have the ability to use any Snyk CLI features. However, if your goal is to replicate the behavior of the Snyk CodePipeline integration as closely as possible, you can follow these tips:

* The `snyk test` command produces a non-zero exit code when vulnerabilities are found. Consider adding `|| true` to the end of the command to circumvent this behavior.
* The [snyk-to-html](https://github.com/snyk/snyk-to-html) tool can be used to produce an HTML report of scan results by running a command similar to `snyk test --json | snyk-to-html -o snyk-results.html`
* Use the following [CLI options](https://docs.snyk.io/snyk-cli/commands) to reproduce behaviors you had configured in the AWS CodePipeline integration:
  * [--org=\<ORG\_ID>](https://docs.snyk.io/snyk-cli/commands/test#org-less-than-org_id-greater-than) - Specify the \<ORG\_ID> to run Snyk commands tied to a specific Snyk Organization.
  * [--severity-threshold=\<low|medium|high|critical>](https://docs.snyk.io/snyk-cli/commands/test#severity-threshold-less-than-low-or-medium-or-high-or-critical-greater-than) - Report only vulnerabilities at the specified level or higher.
  * [--all-projects](https://docs.snyk.io/snyk-cli/commands/test#all-projects) - Auto-detect all Projects in the working directory.
  * [--project-name=\<PROJECT\_NAME>](https://docs.snyk.io/snyk-cli/commands/monitor#project-name-less-than-project_name-greater-than) - Specify a custom Snyk Project name to the `snyk monitor` command.

### Test and validate

* Trigger a manual build in your CodePipeline to test the new CodeBuild integration.
* Verify that the Snyk scan executes successfully and outputs results as expected.
* Ensure your subsequent pipeline stages handle the scan output appropriately.

### Deployment

* When testing is complete, consider deploying the updated CodePipeline.
* Monitor your pipeline for successful Snyk scan execution and address any integration issues.

### Next Steps

Refer to the [Snyk CLI](https://docs.snyk.io/snyk-cli) documentation to incorporate additional security scans into your CI/CD pipeline.

## Conclusion

By following these steps and considerations, you can successfully migrate your security scanning workflow from the Snyk and AWS CodePipeline integration to a more streamlined and configurable solution using AWS CodeBuild.

## Troubleshooting

**How do I store the Snyk token in AWS Secrets Manager and use it in AWS CodeBuild?**

If you use an AWS Secrets Manager environment variable, store your token in AWS Secrets Manager as **plain text** and ensure that your CodeBuild service role has the `secretsmanager:GetSecretValue` permission in IAM. The **value** of the environment variable in AWS CodeBuild should be set to the **Secret name** in AWS Secrets Manager.
