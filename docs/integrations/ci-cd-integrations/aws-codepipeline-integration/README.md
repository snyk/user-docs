# AWS CodePipeline integration

Snyk integrates seamlessly with AWS CodePipeline to scan your application for open source security vulnerabilities and help you deliver secure applications with continuous delivery service. This integration allows CodePipeline users to make security an automated part of their build, test, and deploy phases.

{% hint style="info" %}
Snyk integration is currently available in AWS `sa-east-1` | `ca-central-1` | ap-`southeast-1` | `ap-southeast-2` | `ap-south-1` | `ap-northeast-2` | `ap-northeast-1` | `eu-west-3` | `eu-west-1` | `eu-north-1` | `us-east-1` | `us-west-2` | `eu-west-2` | `eu-central-1` regions. Snyk is actively working on expanding to additional regions.
{% endhint %}

##

## View scan results

You can view scan results in the AWS CodePipeline console, by clicking **Details** in the Scan stage:

![Details in the Scan stage](../../../.gitbook/assets/aws-cp-findings-report.png)

Click the **link to execution details** to view your detailed vulnerability report.

![Link to execution details](../../../.gitbook/assets/image4-2-.png)

## Test report details

Snyk analyzes the manifest file of the application and correlates the list of dependencies with the Snyk vulnerability database. [Snyk provides detailed reports for your open source code](../../../features/snyk-reports/reports-overview.md). By analyzing the manifest file, Snyk builds a full dependency tree, accurately identifying both direct and transitive dependencies (transitive dependencies account for 78% of the vulnerabilities detected by Snyk). This enables Snyk to show exactly how a vulnerability was introduced into the application.

![Snyk test report](../../../.gitbook/assets/prototype.png)

{% hint style="info" %}
Reports are stored for 14 days before they expire. Subsequent pipeline runs update the report and reset the retention period.
{% endhint %}
