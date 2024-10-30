# AWS CodePipeline integration by adding a Snyk scan stage

{% hint style="warning" %}
**The Snyk integration for AWS CodePipeline will be discontinued.**

**Action required**

To safeguard the security of Snyk services and customers, Snyk has begun the deprecation of its integration with **AWS CodePipeline**. To minimize disruption, Snyk recommends that you transition to using **AWS CodeBuild** and the Snyk CLI as an alternative that will support the same use case and functionality.&#x20;

**Migration timeline**

Effective **October 30, 2024**, you will no longer be able to add or modify the Snyk plug-in for new or existing pipelines. Existing pipelines will continue to work as-is for six (6) months, though Snyk recommends migrating to the new process as soon as possible. To avoid disrupting your CI/CD workflows, you must transition to the Snyk CLI before **April 30, 2025**. Refer to the steps in this [migration guide ](migrating-to-aws-codebuild.md)to use Snyk CLI with AWS CodeBuild.\


Snyk is confident that AWS CodeBuild and the Snyk CLI will meet your requirements.
{% endhint %}

Snyk integrates seamlessly with AWS CodePipeline to scan your application for open source security vulnerabilities and help you deliver secure applications with continuous delivery service. This integration allows CodePipeline users to make security an automated part of their build, test, and deploy phases.

{% hint style="info" %}
Snyk integration is available in AWS `sa-east-1` | `ca-central-1` | `ap-southeast-1` | `ap-southeast-2` | `ap-south-1` | `ap-northeast-2` | `ap-northeast-1` | `eu-west-3` | `eu-west-1` | `eu-north-1` | `us-east-1` | `us-east-2` | `us-west-1` | `us-west-2` | `eu-west-2` | `eu-central-1` regions. Snyk is actively working on expanding to additional regions.
{% endhint %}

For setup and use details, see the following pages:

* [Language support for AWS CodePipeline](language-support-for-aws-codepipeline.md)
* [Setup requirements for AWS CodePipeline](setup-requirements-for-aws-codepipeline.md)
* [AWS CodePipeline CodeBuild step example](aws-code-pipeline-codebuild-step-example.md)
* [Setup steps for AWS CodePipeline](setup-steps-for-aws-codepipeline-integration.md)
* [View AWS CodePipeline scan results](view-aws-codepipeline-scan-results.md)
* [AWS CodePipeline test report results](aws-codepipeline-test-report-details.md)

{% hint style="info" %}
Snyk integration with AWS CodePipeline requires a UI-based authentication step as part of the setup. This is not compatible with automation done by non-interactive setup methods such as CloudFormation or Terraform.
{% endhint %}
