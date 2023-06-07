# Integrated Infrastructure as Code

{% hint style="info" %}
**Feature availability**\
Integrated Infrastructure as Code (IaC) is a new version of Snyk IaC, with cloud context that secures cloud configurations across the entire SDLC, from code to deployed cloud environments. Integrated IaC is currently in **closed beta**. Reach out to your account team if you would like access.
{% endhint %}

## **What is Integrated IaC?**

Integrated Infrastructure as Code (IaC) is a new version of Snyk IaC that secures cloud configurations across the entire SDLC, from code to deployed cloud environments. &#x20;

Integrated Infrastructure as Code (IaC) helps developers to write secure infrastructure configurations, with improvements including:

* Improved accuracy of issue detection with line numbers in the CLI.
* Support for HashiCorp Terraform modules and variables for SCM repositories and the CLI.
* Powerful code-to-cloud capabilities, including the ability to [fix cloud issues in IaC](fix-cloud-issues-in-integrated-iac.md) and run IaC tests in the CLI with cloud context to reduce noise.

## Unified Policy Engine and Ruleset

Integrated IaC uses the same unified policy engine as cloud context, rather than the current IaC policy engine. This means you can apply the same security rules to both IaC files and runtime cloud resources.

## New Issues UI

Integrated IaC has focused on providing visibility into an Organization's issues across all its configuration templates. As a result, traditional Projects are currently not supported in Integrated IaC. Instead, [issues](../snyk-cloud-issues/) can be found using the new [Cloud tab](../snyk-cloud-issues/view-cloud-issues-in-the-snyk-web-ui.md), which provides a snapshot of the resources that are misconfigured. This allows security teams to identify and investigate issues easily. Users can view all of their infrastructure as code resources in a single, centralized interface.

## Powerful filters for triaging issues

You can filter issues based on parameters, including rule severity and resource type, allowing you to target the most mission-critical resources. You can also group issues by rule or by resource to more easily view relevant misconfigurations.

## Resource configuration snapshots

Integrated IaC records the configuration attributes of every resource it scans for a user's configuration file compared to the Current IaC Projects view. This lets users review a resource’s configuration at a given moment in time and see the attributes that have caused a specific issue.

## Connect Cloud back to IaC

Snyk cloud context and integrated IaC enable you to [fix Cloud issues directly in the IaC source code](fix-cloud-issues-in-integrated-iac.md) that was used to deploy the misconfigured Cloud resources, by linking a Cloud issue to the underlying IaC template via an SCM source code link.

