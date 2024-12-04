# View AWS CodePipeline scan results

{% hint style="warning" %}
**The Snyk integration for AWS CodePipeline will be discontinued**

\
**Action Required**

In order to safeguard the security of our services and our customers, Snyk has begun the deprecation of its integration with **AWS CodePipeline**. To minimize disruption, we recommend that you transition to using **AWS CodeBuild** and the Snyk CLI as an alternative which will support the same use case and functionality.&#x20;



**Migration Timeline**

Effective **Oct 30th, 2024**, you will no longer be able to add or modify the Snyk plug-in for new or existing pipelines. Existing pipelines will continue to work as-is for 6 months, though we recommend migrating to the new process as soon as possible. To avoid disrupting your CI/CD workflows, you must transition to the Snyk CLI before **April 30, 2025**. Please refer to the steps in this [migration guide](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/aws-codepipeline-integration-by-adding-a-snyk-scan-stage/migrating-to-aws-codebuild) to use Snyk CLI with AWS CodeBuild.\


We are confident that AWS CodeBuild and the Snyk CLI will meet your requirements.&#x20;
{% endhint %}



You can view scan results in the AWS CodePipeline console, by clicking **Details** in the Scan stage:

![Details in the Scan stage](../../../.gitbook/assets/aws-cp-findings-report.png)

Click the **link to execution details** to view your detailed vulnerability report as shown in the following screenshot.

![Link to execution details](../../../.gitbook/assets/image4-2-.png)
