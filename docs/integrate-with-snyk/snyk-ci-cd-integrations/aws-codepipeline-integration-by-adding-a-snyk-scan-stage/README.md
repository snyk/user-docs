# AWS CodePipeline integration by adding a Snyk scan stage

Snyk integrates seamlessly with AWS CodePipeline to scan your application for open source security vulnerabilities and help you deliver secure applications with continuous delivery service. This integration allows CodePipeline users to make security an automated part of their build, test, and deploy phases.

{% hint style="info" %}
Snyk integration is currently available in AWS `sa-east-1` | `ca-central-1` | ap-`southeast-1` | `ap-southeast-2` | `ap-south-1` | `ap-northeast-2` | `ap-northeast-1` | `eu-west-3` | `eu-west-1` | `eu-north-1` | `us-east-1` | `us-west-2` | `eu-west-2` | `eu-central-1` regions. Snyk is actively working on expanding to additional regions.
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
