# Integrated Infrastructure as Code

{% hint style="info" %}
Integrated IaC is currently in Limited Availability with plans of General Availability following soon. If you are interested in Integrated IaC, please contact your Snyk account team.
{% endhint %}

## **What is Integrated IaC?**&#x20;

Integrated IaC is an evolution of our [Current IaC](../) product which now utilizes the same unified policy engine and data model as [Snyk Cloud](../../snyk-cloud/).

Integrated Infrastructure as Code (IaC) continues to help developers to write secure application configurations. Integrated IaC provides improved accuracy of issue detection with line numbers in the CLI and UI with detailed remediation advice, so you can make changes directly to your code, before applications reach production.

Integrated IaC helps developers write secure configurations for HashiCorp Terraform (variables and module support), AWS CloudFormation, Kubernetes, and Azure Resource Manager (ARM).&#x20;

## Unified Policy Engine/Ruleset

&#x20;Integrated IaC uses the same unified policy engine that powers Snyk Cloud, rather than the Current IaC policy engine. This means you can apply the same security rules to both IaC files and runtime cloud resources.&#x20;

## New Cloud UI

Integrated IaC has focused on providing visibility into an Organization's issues across all its configuration templates. As a result, traditional Projects are currently not supported in Integrated IaC. Instead, [issues](https://docs.snyk.io/products/snyk-cloud/snyk-cloud-issues) can be found using the new [Cloud tab](https://docs.snyk.io/products/snyk-cloud/snyk-cloud-issues/view-cloud-issues-in-the-snyk-web-ui) and [Cloud Issues API](https://docs.snyk.io/products/snyk-cloud/snyk-cloud-issues/view-cloud-issues-in-the-api) which will provide a snapshot into the resources that are misconfigured. This allows security teams to easily identify and investigate issues. Users can view all of their infrastructure as code resources in a single, centralized interface.&#x20;

## Powerful filters for triaging issues

You can filter issues based on parameters including rule severity and resource type, allowing you to target the most mission-critical resources. You can also group issues by rule or by resource, to more easily view relevant misconfigurations.&#x20;

## Resource configuration snapshots

Integrated IaC records the configuration attributes of every resource it scans for a user's configuration file compared to the Current IaC Projects view. This lets users review a resource’s configuration at a given moment in time and see the attributes that have caused a specific issue.

## Coming soon&#x20;

Snyk Integrated IaC is in Limited Availability, and our product and engineering teams are working hard to add additional features. Please reach out to your Snyk account team if you have questions on our near-term roadmap.

