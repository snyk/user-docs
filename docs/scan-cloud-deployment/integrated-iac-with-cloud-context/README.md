# Integrated IaC with cloud context

{% hint style="info" %}
**Feature availability**\
Integrated Infrastructure as Code (IaC) is a new version of Snyk IaC, with cloud context that secures cloud configurations across the entire SDLC, from code to deployed cloud environments. Integrated IaC is currently in **closed beta**. Reach out to your account team if you would like access.
{% endhint %}

Snyk cloud context helps users find, prioritize, and fix cloud misconfigurations by providing for scanning and testing deployed cloud resources with a library of predefined security rules. Scanning generates a record of issues that can be filtered and triaged.

Integrated Infrastructure as Code (IaC) is a new version of Snyk IaC with Snyk cloud context. Integrated IaC secures cloud configurations across the entire SDLC, from code to deployed AWS, Azure, and Google Cloud environments. Integrated IaC has the cloud context interfaces, workflows, policy engine, and data model.

## **What is Integrated IaC?**

Integrated Infrastructure as Code (IaC) is a new version of Snyk IaC that secures cloud configurations across the entire SDLC, from code to deployed cloud environments. &#x20;

Integrated Infrastructure as Code (IaC) helps developers to write secure infrastructure configurations, with improvements including:

* Improved accuracy of issue detection with line numbers in the CLI.
* Support for HashiCorp Terraform modules and variables for SCM repositories and the CLI.
* Powerful code-to-cloud capabilities, including the ability to [fix cloud issues in IaC](fix-cloud-issues-in-integrated-iac.md) and run IaC tests in the CLI with cloud context to reduce noise.

## Features of Integrated IaC

### Developer-first security for cloud configurations

With Integrated IaC and cloud context, Snyk extends its focus on developer security to enable developers to build and deploy secure cloud configurations from IaC in the IDE through to production. Snyk has the same interfaces across the UI and API for cloud configurations in code and in deployed environments.&#x20;

### Code to cloud with the Snyk unified policy engine

The Snyk unified policy engine enables consistent security testing across IaC files and deployed cloud resources, giving your security teams assurance that misconfigurations are correctly identified. In combination with the other Snyk products, this enables you to find and fix issues throughout the entire software development lifecycle (SDLC).

### Connect Cloud back to IaC for faster fixes

Fix cloud issues directly in the IaC source code that was used to deploy the misconfigured cloud resources by viewing the underlying SCM source code link for a given Cloud issue. With Snyk, you can view resource configurations alongside IaC-specific remediation guidance.

### Security and compliance posture for security teams

Snyk cloud context and Integrated IaC provide visibility into an Organization’s entire cloud footprint, allowing security teams to identify and investigate issues easily. Users view reports and focus on issues with impact on compliance, with support for more than 10+ compliance standards, including CIS Benchmarks, PCI-DSS, SOC 2, and more.

## Unified Policy Engine and Ruleset

Integrated IaC uses the same unified policy engine as cloud context, rather than the current IaC policy engine. This means you can apply the same security rules to both IaC files and runtime cloud resources.

## New Issues UI

Integrated IaC has focused on providing visibility into an Organization's issues across all its configuration templates. As a result, traditional Projects are currently not supported in Integrated IaC. Instead, [issues](snyk-cloud-issues/) can be found using the new [Cloud tab](snyk-cloud-issues/view-cloud-issues-in-the-snyk-web-ui.md), which provides a snapshot of the resources that are misconfigured. This allows security teams to identify and investigate issues easily. Users can view all of their infrastructure as code resources in a single, centralized interface.

## Powerful filters for triaging issues

You can filter issues based on parameters, including rule severity and resource type, allowing you to target the most mission-critical resources. You can also group issues by rule or by resource to more easily view relevant misconfigurations.

## Resource configuration snapshots

Integrated IaC records the configuration attributes of every resource it scans for a user's configuration file compared to the Current IaC Projects view. This lets users review a resource’s configuration at a given moment in time and see the attributes that have caused a specific issue.

## Connect Cloud back to IaC

Snyk cloud context and integrated IaC enable you to [fix Cloud issues directly in the IaC source code](fix-cloud-issues-in-integrated-iac.md) that was used to deploy the misconfigured Cloud resources, by linking a Cloud issue to the underlying IaC template via an SCM source code link.
