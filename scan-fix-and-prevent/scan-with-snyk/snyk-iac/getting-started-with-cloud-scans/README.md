# Getting started with cloud scans

Use Snyk IaC cloud scans to find, view, and fix issues in deployed cloud resource configurations for AWS, Azure, and Google Cloud.

This page explains using cloud scans in the Snyk Web UI. For information about using cloud scans with the Snyk CLI, see [Test your IaC files](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/test-your-iac-files/).

## Prerequisites for cloud scans

To start using cloud scans you must have the following:

* A Snyk account. For details, see [Getting started](../../../discover-snyk/getting-started/#create-or-log-in-to-a-snyk-account).
* Snyk IaC on the enterprise plan.
* An existing Terraform, CloudFormation, or Azure Resource Manager environment to work in, or deployed AWS, Azure, or Google Cloud account to onboard.

## Import cloud environments

Navigate to your **Organization** **Settings** > **Cloud environments**.

To add a cloud environment, select the **Add environment** drop-down and select the cloud provider. Follow the steps in [AWS Integration: Web UI](../cloud-platform-integrations/aws-integration/aws-integration-web-ui/), [Google Cloud Integration: Web UI](../cloud-platform-integrations/google-cloud-integration/google-cloud-integration-web-ui/), or [Azure Integration: Web UI](../cloud-platform-integrations/azure-integration-for-cloud-configurations/azure-integration-web-ui/) to create the environment.&#x20;

You can also add an environment using the Snyk API:

* [AWS Integration: API](../cloud-platform-integrations/aws-integration/aws-integration-api/)
* [Google Cloud Integration: API](../cloud-platform-integrations/google-cloud-integration/google-cloud-integration-api/)
* [Azure Integration: API](../cloud-platform-integrations/azure-integration-for-cloud-configurations/azure-integration-api/)

## View cloud issues

To view your cloud issues in the Snyk Web UI, navigate to the Organization, and on the menu, select **Cloud**. For more information, see [View cloud issues in the Snyk Web UI](manage-cloud-issues/view-cloud-issues-in-the-snyk-web-ui.md).
