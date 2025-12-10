# Snyk Essentials

Snyk Essentials helps AppSec teams better operationalize and scale use of Snyk with broad application visibility and security coverage management.

## Overview

Snyk Essentials enables:

* Automated app asset discovery: Continually discover application assets and classify them by business context, ensuring security is in sync with development.
* Tailored security controls: Define and manage appropriate security and compliance requirements, and verify the correct controls are in place.
* Risk-based prioritization: Assess risk for each app based on application context and best-in-class security analysis and fix guidance to focus developer remediation efforts on issues that matter most to the business.

The following videos show the capabilities of Snyk Essentials.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737657001/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-1_-_v1_-_Snyk_Essentials_Overview.mp4" %}
Snyk Essentials core capabilities overview
{% endembed %}

## Features

Snyk Essentials includes the following features:

* [Integrations](../developer-tools/scm-integrations/group-level-integrations/) to support ingesting data from SCM tools for asset discovery, Snyk Application Security Testing products for security controls coverage, and ticketing or notification tools for policy actions. Use the the Integration page to add and set up new integrations.
* [Policies](../manage-risk/policies/assets-policies/) to classify and tag assets with business context and configure actions using a Policy Builder UI.
* [Inventory](../manage-assets/manage-assets.md) layouts for managing assets and viewing Snyk coverage.
* A dashboard to view, add, and customize widgets.

## Prerequisites

* You are a Snyk Enterprise customer.
* You have the necessary permissions to onboard cloud-based SCM tools (Azure DevOps, GitHub, GitLab, and so on) to Snyk Essentials for repository asset discovery.

{% hint style="info" %}
When you integrate a Git code repository with Snyk Essentials, you should use a secondary token with a broad, complete view of the code repository, not only of what you imported into Snyk. \\

Use a secondary token to countercheck everything onboarded using Snyk. \\

Using the secondary token reduces the likelihood of introducing a blind spot from a limited token at the Organization level configuration. \\

The first import, synchronization, can take up to 24 hours to complete.
{% endhint %}

## Permissions

Snyk Essentials is included in the Snyk Enterprise plan. You do not need additional permissions to access it. For more information on default user roles and permissions, see [Default user roles](../snyk-platform-administration/user-roles/pre-defined-roles.md).

## Login and authentication

Log in and authenticate to Snyk using existing mechanisms (SSO, Google SAML, and so on).

## Key concepts

### Asset

meaningful, real-world components in an application’s SDLC, where meaningful means either carries a risk or aggregates risk of other components (for example, repositories that contain packages) and real-world means that the concept exists outside of Snyk, for example, repository (which is a generally applicable term).

### **Controls**

The security controls associated with the asset. Navigate to the [Coverage controls](../manage-risk/policies/assets-policies/use-cases-for-policies/coverage-control-policy.md) section to see all available statuses for security controls.

### **Coverage**

An assessment of whether applicable assets are scanned and tested by security tools (like Snyk Open Source, for instance), as it relates to an application security program. A type of policy that allows you to specify what controls should be applied and, optionally, how often it needs to be run.

### **Tags**

A way to categorize assets. Helps you recognize or handle assets differently according to mutual properties. Assets can be filtered by their tags in the inventory or when creating policy rules. A tag can be automatically assigned to an asset, or the asset can be tagged by a policy you created. GitHub and GitLab topics are treated as asset tags, and you can use them to create policies.

### **Class**

A way to assign business context to assets and categorize an asset based on the business criticality. Assets can be assigned Classes A, B, C, or D, where Class A (assets that are business critical, deal with sensitive data, subject to compliance, and so on) is the most important and Class D (test apps, sandbox environments, and so on) the least important. Assets are assigned Class C by default. A class can be used in policies as well as defined in a policy.

### **Policy**

A way to automate actions in certain conditions, like classifying and tagging assets with business context. You can also use a policy to configure actions like sending a message or setting the coverage gap control using a Policy builder UI.

## Scanning methods

You can initiate a scan from the Web UI, the CLI, the API, or with PR Checks. See [Start scanning](start-scanning.md) for more details.

If you initiate your scans using the CLI, you might encounter one of the following situations:

1. If you have a `.git` folder available in the directory that the CLI is scanning, then the `git remoteurl` is picked up automatically for Snyk Open Source, Snyk Container, and Snyk IaC.

{% hint style="info" %}
Snyk Code does not automatically pick up the `git remoteurl`, even if the `.git` folder is available in the directory scanned by the CLI.
{% endhint %}

2. If you do not have a `.git` folder available in the directory that the CLI is scanning, you can use different test or monitor commands to achieve the same result:
   * [`snyk monitor`](../developer-tools/snyk-cli/commands/monitor.md#remote-repo-url-less-than-url-greater-than), for Snyk Open Source
   * [`snyk iac test`](../developer-tools/snyk-cli/commands/iac-test.md#remote-repo-url-less-than-url-greater-than) - also requires the `--report` command
   * `snyk container monitor` - no options available.
   * `snyk code test` - no options available.
