# Integrated Infrastructure as Code

{% hint style="info" %}
**Feature availability**\
Integrated Infrastructure as Code (IaC) is a new version of Snyk IaC that - together with Snyk Cloud - secures cloud configurations across the entire SDLC, from code to deployed cloud environments. Integrated IaC is currently in **closed beta**. Please reach out to your account team if you would like access.
{% endhint %}

## **What is Integrated IaC?**

Integrated Infrastructure as Code (IaC) is a new version of Snyk IaC that secures cloud configurations across the entire SDLC, from code to deployed cloud environments. &#x20;

Integrated Infrastructure as Code (IaC) continues to help developers to write secure infrastructure configurations, with improvements including:

* Improved accuracy of issue detection with line numbers in the CLI.
* Support for HashiCorp Terraform modules and variables for SCM repositories and the CLI.
* Powerful code-to-cloud capabilities, including the ability to fix Cloud issues in IaC, and run IaC tests in the CLI with cloud context to reduce noise.

## Unified Policy Engine/Ruleset

Integrated IaC uses the same unified policy engine that powers Snyk Cloud, rather than the Current IaC policy engine. This means you can apply the same security rules to both IaC files and runtime cloud resources.

## New Issues UI

Integrated IaC has focused on providing visibility into an Organization's issues across all its configuration templates. As a result, traditional Projects are currently not supported in Integrated IaC. Instead, [issues](https://docs.snyk.io/products/snyk-cloud/snyk-cloud-issues) can be found using the new [Cloud tab](https://docs.snyk.io/products/snyk-cloud/snyk-cloud-issues/view-cloud-issues-in-the-snyk-web-ui), which provides a snapshot into the resources that are misconfigured. This allows security teams to easily identify and investigate issues. Users can view all of their infrastructure as code resources in a single, centralized interface.

## Powerful filters for triaging issues

You can filter issues based on parameters including rule severity and resource type, allowing you to target the most mission-critical resources. You can also group issues by rule or by resource, to more easily view relevant misconfigurations.

## Resource configuration snapshots

Integrated IaC records the configuration attributes of every resource it scans for a user's configuration file compared to the Current IaC Projects view. This lets users review a resource’s configuration at a given moment in time and see the attributes that have caused a specific issue.

## Connect Cloud back to IaC

Snyk Cloud and integrated IaC enable you to fix Cloud issues \[link to new feature page] directly in the IaC source code that was used to deploy the misconfigured Cloud resources, by linking a Cloud issue to the underlying IaC template via a SCM source code link.

