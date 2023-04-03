# Integrated Infrastructure as Code

{% hint style="info" %}
Integrated IaC is currently in Limited Availability with plans of General Availability following soon. If you are interested in Integrated IaC, please contact your Snyk account team.
{% endhint %}

## **What is Integrated IaC?**

Integrated IaC is an evolution of our [Current IaC](../) product, and it uses the same unified policy engine and data model as [Snyk Cloud](../../snyk-cloud/).

Integrated Infrastructure as Code (IaC) continues to help developers to write secure application configurations. Integrated IaC provides improved accuracy of issue detection with line numbers in the CLI and UI with detailed remediation advice, so you can make changes directly to your code, before applications reach production.

Integrated IaC helps developers write secure configurations for HashiCorp Terraform (variables and module support) and AWS CloudFormation, with support for Kubernetes and Azure Resource Manager (ARM) coming soon.

## Unified Policy Engine/Ruleset

Integrated IaC uses the same unified policy engine that powers Snyk Cloud, rather than the Current IaC policy engine. This means you can apply the same security rules to both IaC files and runtime cloud resources.

## New Cloud UI

Integrated IaC has focused on providing visibility into an Organization's issues across all its configuration templates. As a result, traditional Projects are currently not supported in Integrated IaC. Instead, [issues](https://docs.snyk.io/products/snyk-cloud/snyk-cloud-issues) can be found using the new [Cloud tab](https://docs.snyk.io/products/snyk-cloud/snyk-cloud-issues/view-cloud-issues-in-the-snyk-web-ui), which provides a snapshot into the resources that are misconfigured. This allows security teams to easily identify and investigate issues. Users can view all of their infrastructure as code resources in a single, centralized interface.

## Powerful filters for triaging issues

You can filter issues based on parameters including rule severity and resource type, allowing you to target the most mission-critical resources. You can also group issues by rule or by resource, to more easily view relevant misconfigurations.

## Resource configuration snapshots

Integrated IaC records the configuration attributes of every resource it scans for a user's configuration file compared to the Current IaC Projects view. This lets users review a resource’s configuration at a given moment in time and see the attributes that have caused a specific issue.

## Connect Cloud back to IaC

Snyk Cloud and integrated IaC enable you to fix Cloud issues \[link to new feature page] directly in the IaC source code that was used to deploy the misconfigured Cloud resources, by linking a Cloud issue to the underlying IaC template via a SCM source code link.

## Coming soon

Snyk Integrated IaC is in Limited Availability, and our product and engineering teams are working hard to add additional features. Please reach out to your Snyk account team if you have questions on our near-term roadmap.
