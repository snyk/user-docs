# Snyk AppRisk

Snyk AppRisk is a product that enables Application Security teams to implement, manage, and scale a modern, high-performing, developer security program. The product covers use cases under Application Security Posture Management (ASPM).

## Overview

Snyk AppRisk builds upon the [capabilities ](snyk-essentials.md#overview)of Snyk Essentials by offering:

* More sophisticated risk-based prioritization with runtime Insights - such as whether an app is deployed and public facing, and whether vulnerable packages found by Snyk Open Source are actually used in runtime.
* Manage security coverage for secrets detection tools: Assess coverage for secrets detection alongside Snyk's AST products.
* Scalability for diverse environments: Robust security management across various development environments.

The following videos show the capabilities of Snyk AppRisk.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737657002/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-2_-_v1_-_Snyk_AppRisk_Overview.mp4" %}
Snyk AppRisk additional capabilities overview
{% endembed %}

## Features

Snyk AppRisk provides additional features beyond [those available in Snyk Essentials](snyk-essentials.md#features):

* [Integrations](../integrations/connect-a-third-party-integration.md) with non-Snyk products to support security coverage and Insights use cases
* [Insights](../manage-risk/prioritize-issues-for-fixing/prioritization-for-snyk-essentials.md) for runtime context - providing risk factors to help with risk-based prioritization.
* [Analytics](../manage-risk/analytics/application-analytics.md) for reviewing and comparing assets and issues metrics at the level of asset classes, applications, or code owners.

## Prerequisites

* You are a Snyk Enterprise customer.
* Your account is entitled with access for Snyk AppRisk.
* You are a Group Administrator for the Group associated with Snyk AppRisk, or you are assigned a Group level role with permissions to View Group and Edit Snyk Essentials.
* The Group associated with Snyk AppRisk includes organizations that have onboarded Snyk application security products.
* You have the necessary permissions to onboard cloud-based SCM tools (Azure DevOps, GitHub, GitLab, and so on) to Snyk AppRisk for repository asset discovery.

{% hint style="info" %}
When you integrate a Git code repository with Snyk Essentials, you should use a secondary token with a broad, complete view of the code repository, not only of what you imported into Snyk. \
Use a secondary token to countercheck everything onboarded using Snyk. \
Using the secondary token reduces the likelihood of introducing a blind spot from a limited token at the Organization level configuration. \
The first import, synchronization, can take up to 24 hours to complete.
{% endhint %}

## Permissions

You can access Snyk AppRisk with one of the Group-level role permissions described here. To access the permissions, navigate to **View groups** > **Snyk Snyk Essentials permissions**.

* **View Snyk Essentials** - Grants you read-only access to Snyk Essentials.
* **Edit Snyk Essentials** - Grants you edit access to Snyk Essentials, for example, edit policies, edit asset classification, and add the integration.

A Group Administrator has the **Edit Snyk Essentials** permission assigned by default, and a Group Viewer has the **View Snyk Essentials** permission assigned by default.

{% hint style="info" %}
For more information on default user roles and permissions, see [Default user roles](../snyk-platform-administration/user-roles/pre-defined-roles.md).
{% endhint %}

## Login and Authentication

Log in and authenticate to Snyk using existing mechanisms (SSO, Google SAML, and so on).

## Accessing Snyk AppRisk

Ensure you are at the Group level to access the Snyk AppRisk options. From the Group level, you have centralized security management that enhances security and simplifies security procedures for Projects.

The Snyk AppRisk features are available at the [Group level from the Snyk Web UI](../discover-snyk/getting-started/snyk-web-ui.md#group-level).&#x20;

## Key concepts

Some of the key concepts for Snyk AppRisk are asset, class, coverage, and policy. For more details, see  [Snyk - Essentials Key concepts](snyk-essentials.md).

## Scanning methods

You can initiate a scan from the Web UI, the CLI, the API, or with PR Checks. For more details, see [Scanning methods](snyk-essentials.md#scanning-methods) and  [Start scanning](start-scanning.md).

## Capabilities and features

<table><thead><tr><th width="288"></th><th>Snyk Essentials</th><th>Snyk AppRisk</th></tr></thead><tbody><tr><td>Availability</td><td>Included with all Enterprise plans.</td><td>Available for sale; reach out to your account manager for more information.</td></tr><tr><td>Applicability</td><td>Helps Snyk customers better manage their developer-first application security program with Snyk.</td><td>Helps Snyk customers prioritize with Insights risk factors from runtime data sources, and manage and scale their developer-first application security program more holistically.</td></tr><tr><td>Integrations and Data Sources</td><td><ul><li>SCM</li><li>Application Context - Developer Portals, Service Catalogs, CMDBs</li><li>Jira integration</li></ul></td><td><p>In addition to Essentials:</p><ul><li>3rd party secrets</li><li>3rd party Runtime, observability, cloud, CNAPP</li><li>Snyk Runtime Sensor</li></ul></td></tr><tr><td>Discovery and visibility</td><td><ul><li>Discover and classify code-based assets and runtime-based assets (repositories, packages, images).</li><li>Issue counts identified by Snyk</li></ul></td><td><ul><li>Discover and classify code and runtime-based assets (repositories, packages, images).</li><li>Issue counts identified by Snyk and view issue counts identified by third-party tools.</li></ul></td></tr><tr><td>Security coverage management</td><td>Ensure these assets are covered by Snyk.</td><td>Ensure these assets are covered by Snyk or other third-party tools.</td></tr><tr><td>Issue prioritization</td><td>Prioritize with asset and application context, and automate risk management workflows for assets using policies.</td><td><p>Manage risk holistically by:</p><ul><li>Prioritizing issues posing greatest risk with runtime Insights.</li><li>Tracking and reporting upon program health with risk and coverage metrics via Application Analytics.</li></ul></td></tr></tbody></table>
