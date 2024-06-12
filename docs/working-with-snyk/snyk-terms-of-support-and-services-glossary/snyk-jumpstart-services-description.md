# Snyk Jumpstart Services Description

## Overview of Snyk Jumpstart

A Snyk Consultant will provide services to help the Customer accelerate the setup of Snyk products through assisted account configuration (the “Jumpstart Services”). The engagement will consist of knowledge transfer, paired with configuration guidance for your team.

The objective is a working setup of Snyk and a Customer team that is well-prepared to continue its application security efforts.

## Recommended for

* Teams needing assistance with the setup of Snyk products
* Teams who are new to Snyk and have limited experience with security scanning
* Self-starters who are comfortable with extending and maintaining Snyk on their own, post-engagement

## Jumpstart Services description

The Snyk Consultant will deliver the following services related to the setup of Snyk remotely as part of the Jumpstart Services to the Customer. Note that the Jumpstart Services will be delivered only for the Snyk Applications purchased on an Order Form at the same time as this service. References to Snyk Services not purchased by the Customer are hereby omitted unless otherwise noted.

1. [Pre-engagement planning and preparation](snyk-jumpstart-services-description.md#pre-engagement-planning-and-preparation)
   1. Review deliverables per product module
   2. Review prerequisites per product module
   3. Confirm full-time availability for Customer contacts based on Customer prerequisites
2. [Snyk Platform configuration](snyk-jumpstart-services-description.md#snyk-platform-configuration)
3. [Snyk Open Source configuration](snyk-jumpstart-services-description.md#snyk-open-source-configuration)
4. [Snyk Code configuration](snyk-jumpstart-services-description.md#snyk-code-configuration)
5. [Snyk Container configuration](snyk-jumpstart-services-description.md#snyk-container-configuration)
6. [Snyk IaC configuration](snyk-jumpstart-services-description.md#snyk-iac-configuration)
7. [Snyk AppRisk prioritization configuration](snyk-jumpstart-services-description.md#snyk-apprisk-essentials-prioritization-configuration)

## Pre-engagement planning and preparation

### Pre-engagement call

A Pre-engagement call will be held before the start of the Jumpstart Services listed herein, to ensure the Customer understands the prerequisites required for engagement start along with the deliverables completed during the engagement timeframe. Customer prerequisites, including resources, availability, and deliverables for each product module, will be reviewed and confirmed. The Customer acknowledges that complying with these prerequisites is its sole responsibility, and Snyk will not be responsible for any delays or failure to deliver the Jumpstart Services based on the Customer’s failure to meet these prerequisites.

## Snyk Platform configuration

### Delivery approach - Snyk Platform configuration

The Snyk delivery method is designed to ensure rapid value realization with Snyk. Throughout our collaboration, Snyk will guide the Customer in setting up a foundational configuration and equip them to expand this setup to other applications and integrations. Ensuring Snyk is correctly set up from the outset improves developer adoption and paves the way for long-term success.

#### Base SSO Configuration and Custom Mapping

The Snyk Consultant will work with the Customer to configure SSO through SAML, Entra ID (formerly Azure AD), OIDC, or ADFS connection along with custom mapping of dynamically assigned users to Snyk Groups and Organizations based on data provided by the desired Identity Provider (IdP) to set up a scaled user provisioning and access model.

#### **Template Organization configuration**

The Snyk Consultant will work with the Customer to configure notification settings, language settings, and Snyk account structure. Time will be spent reviewing configuration details with the Customer to ensure the Customer understands and maintains them.

#### Admin Training

The Snyk Consultant will run a 60-minute training session to ensure Customer administrator users know how to configure essential settings in Snyk based on the maturity of their security program.

Training topics include:

* Navigating the Snyk UI
* Integration Settings
* Notification Settings
* Enabling new features
* User roles and permissions
* Filtering and prioritizing issues
* Viewing results in Snyk reporting
* Remediation workflow
* Ignoring issues in the Snyk UI
* PR Check walkthrough
* Security and license policies
* Accessing additional training materials (learn.snyk.io)
* Project Collections and views

### Target Initiatives - Snyk Platform configuration

<table><thead><tr><th width="248">Platform - initial setup</th><th width="410">Outcome</th><th>Hours</th></tr></thead><tbody><tr><td>Base SSO connection configuration</td><td>Users can access Snyk based on role.</td><td>1</td></tr><tr><td>SSO Custom Mapping</td><td>Users can access Snyk based on role.</td><td>4</td></tr><tr><td>(Template Org) Custom Roles and service account configuration</td><td>Provide a templated organization to replicate and scale your setup of Snyk quickly.</td><td>.5</td></tr><tr><td>(Template Org) Notification configuration</td><td>Provide a templated organization to replicate and scale your setup of Snyk quickly.</td><td>.5</td></tr><tr><td>(Template Org) Language settings configuration</td><td></td><td>.5</td></tr><tr><td>(Template Org) Jira or Slack App configuration</td><td></td><td>1</td></tr><tr><td>(Template Org) Account Organization and Group configuration</td><td></td><td>.5</td></tr><tr><td>Admin training</td><td></td><td>1</td></tr><tr><td><strong>TOTAL HOURS</strong></td><td></td><td><strong>9</strong></td></tr></tbody></table>

## Snyk Open Source configuration

### Delivery approach - Snyk Open Source configuration

The Snyk delivery method is designed to ensure rapid value realization with Snyk. Throughout our collaboration, Snyk will guide the Customer in setting up a foundational configuration and equip the Customer to expand this setup to other applications and integrations. Ensuring Snyk is correctly set up from the outset improves developer adoption and paves the way for long-term success.

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) either through the UI import functionality or through the [API Import tool](../../snyk-api/snyk-tools/tool-snyk-api-import/).

#### SCM integration settings

The Snyk Consultant will work with the Customer to configure SCM integration settings based on the Customer’s desired gating strategy.

#### SCM Broker installation

The Snyk Consultant will work with the Customer to install the Snyk Broker in a pre-determined environment that follows the [Snyk Broker system requirements](../../enterprise-configuration/snyk-broker/prepare-snyk-broker-for-deployment.md).

#### Snyk API Import and SCM snyc

The Consultant will review the Snyk API Import script to ensure the Customer understands how to import additional Projects into Snyk and keep their SCM integration in sync with incoming changes to manifests.

#### Single pipeline configuration (dIrect Integration or CLI)

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk test` and `snyk monitor` commands to provide the Customer with an understanding of how to configure additional pipeline scans.

#### SBOM walkthrough (API and CLI)

The Snyk Consultant will educate the Customer on creating an SBOM through the Snyk API and the Snyk CLI.

#### Interpreting and actioning Open Source results

The Snyk Consultant will educate the Customer on understanding Snyk Open Source results through the CLI and Snyk UI and how to manage Snyk Open Source results using Snyk Reporting.

#### Documentation close-out

The Customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, the document offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document puts forth practical recommendations and actionable next steps that will aid the Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document provides an essential guide for Customers to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk Open Source

<table><thead><tr><th width="300">Snyk Open Source configuration</th><th width="306">Outcome</th><th>Hours</th></tr></thead><tbody><tr><td>Repository import (SCM-only up to 50 targets)</td><td>Import a maximum of 50 targets into Snyk using a <a href="../../scm-ide-and-ci-cd-workflow-and-integrations/git-repositories-scms-integrations-with-snyk/">supported SCM integration</a> (GitHub, Azure Repos, Bitbucket, GitLab).</td><td>2</td></tr><tr><td>SCM integration settings</td><td>Configure SCM integration settings to the Customer’s desired gating settings.</td><td>.5</td></tr><tr><td>SCM Broker installation</td><td>Install SCM Broker in a pre-determined customer environment based on Snyk system requirements.</td><td>2</td></tr><tr><td>Snyk Tools - API Import and SCM Sync</td><td>Gain an understanding of how to use the Snyk API Import script to import additional targets and keep repos in Sync (GHE only).</td><td>2</td></tr><tr><td>Single pipeline configuration (direct integration OR CLI)</td><td>Configure a pipeline to run <code>snyk test</code> and <code>snyk monitor</code>.</td><td>3</td></tr><tr><td>SBOM Walkthrough (CLI and API)</td><td>Gain an understanding of generating an SBOM through Snyk using the CLI and API.</td><td>1</td></tr><tr><td>Interpreting and actioning Open Source results</td><td>Gain an understanding of how to view Open Source results in Snyk Reporting along with managing issues.</td><td>1.5</td></tr><tr><td>Documentation close-out</td><td>Gain an understanding of work completed along with a runbook for onboarding additional projects.</td><td>2</td></tr><tr><td><strong>TOTAL HOURS</strong></td><td></td><td><strong>14</strong></td></tr></tbody></table>

## Snyk Code configuration

### Delivery approach - Snyk Code configuration

The Snyk delivery method is tailored to ensure rapid value realization with Snyk. Throughout our collaboration, Snyk will guide the Customer in setting up a foundational configuration and equip the Customer to expand this setup to other applications and integrations. Ensuring Snyk is correctly set up from the outset improves developer adoption and paves the way for long-term success.

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) either through the UI import functionality or through the [API Import tool](../../snyk-api/snyk-tools/tool-snyk-api-import/). Repos will be imported into Snyk Organizations that either mirror the Customer’s SCM organization structure or through custom Organizations configured in the API-import script.

#### SCM integration settings

The Snyk Consultant will work with the Customer to configure SCM integration settings based on the Customer’s desired gating strategy.

#### SCM Broker installation

The Snyk Consultant will work with the Customer to install the Snyk Broker in a pre-determined environment that follows the [Snyk Broker system requirements](../../enterprise-configuration/snyk-broker/prepare-snyk-broker-for-deployment.md).

#### Snyk API Import and SCM Sync

The Consultant will review the Snyk API Import script to ensure the Customer understands how to import additional Projects into Snyk and keep their SCM integration in sync with incoming changes to manifests.

#### Interpreting and actioning Code results

The Snyk Consultant will educate the Customer on understanding Snyk Code results through the CLI and Snyk UI and how to manage Snyk Code results using Snyk Reporting.

#### Documentation close-out

The Customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, the document offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document puts forth practical recommendations and actionable next steps that will aid the Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document provides an essential guide for customers to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk Code

<table><thead><tr><th width="250">Snyk Code Configuration</th><th>Outcome</th><th>Hours</th></tr></thead><tbody><tr><td>Repository import (SCM-only up to 50 targets)</td><td>Import a maximum of 50 targets into Snyk using a <a href="../../scm-ide-and-ci-cd-workflow-and-integrations/git-repositories-scms-integrations-with-snyk/">supported SCM Integration</a>.</td><td>2</td></tr><tr><td>SCM integration settings</td><td>Configure SCM integration settings to the Customer’s desired gating settings.</td><td>.5</td></tr><tr><td>SCM Broker installation</td><td>Install SCM Broker in a pre-determined customer environment based on Snyk system requirements.</td><td>2</td></tr><tr><td>Snyk Tools - API Import and SCM Sync</td><td>Gain an understanding of how to use the Snyk API Import script to import additional targets and keep their repos in sync (GHE only).</td><td>2</td></tr><tr><td>Interpreting and actioning Code results</td><td>Gain an understanding of how to view Code results in Snyk Reporting along with managing issues.</td><td>1.5</td></tr><tr><td>Documentation close-out</td><td>Gain an understanding of work completed along with a runbook for onboarding additional projects.</td><td>2</td></tr><tr><td><strong>TOTAL HOURS</strong></td><td></td><td><strong>10</strong></td></tr></tbody></table>

## Snyk Container configuration

### Delivery approach - Snyk Container configuration

The Snyk delivery method is tailored to ensure rapid value realization with Snyk. Throughout our collaboration, Snyk will guide the Customer in setting up a foundational configuration and equip them to expand this setup to other applications and integrations. Ensuring Snyk is correctly set up from the outset improves developer adoption and paves the way for long-term success.

#### Single Broker Container Registry installation and configuration

The Snyk Consultant will work with the Customer to configure and install Snyk Broker if needed for a single [supported Container Registry](https://snyk.io/integrations/?type=container-registries) integration.

#### Container Registry import

The Snyk Consultant will work with the Customer to import their container images into Snyk (up to 50 targets) through the UI import functionality.

#### Interpreting and actioning Snyk Container results

The Snyk Consultant will educate the Customer on understanding Snyk Container results through the CLI and Snyk UI and how to manage Snyk Container results using Snyk Reporting.

#### Single CI/CD CLI configuration

The Snyk Consultant will work with the Customer to configure a single pipeline to run the snyk test and `snyk container monitor` commands to provide the Customer with an understanding of how to configure additional pipeline scans.

#### Custom Base Images walkthrough (UI and CLI)

The Snyk Consultant will educate the Customer on how to use the Snyk Custom Base Image Recommendation functionality both in the Snyk UI and CLI.

#### Documentation close-out

The Customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, the document offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document puts forth practical recommendations and actionable next steps that will aid Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document provides an essential guide for customers to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk Container

<table><thead><tr><th width="269">Snyk Container Configuration</th><th>Outcome</th><th>Hours</th></tr></thead><tbody><tr><td>Single Broker Container Registry installation and configuration</td><td>Single Broker installed and configured for a <a href="https://snyk.io/integrations/?type=container-registries">Supported Container Registry</a>.</td><td>3</td></tr><tr><td>Container Registry import (up to 50 targets)</td><td>Import a maximum of 50 targets into Snyk using a <a href="https://snyk.io/integrations/?type=container-registries">Supported Container Registry</a>.</td><td>2</td></tr><tr><td>Interpreting and actioning Container results</td><td>Gain an understanding of how to view Container results in Snyk Reporting along with managing issues.</td><td>2</td></tr><tr><td>Single CI/CD CLI rntegration</td><td>Configure a single pipeline to <code>test</code> and <code>monitor</code> for Snyk Container.</td><td>3</td></tr><tr><td>Custom Base Images walkthrough (UI and CLI)</td><td>Gain an understanding of how to use the Custom Base Image Recommendations functionality through the UI and CLI.</td><td>2</td></tr><tr><td>Documentation close-out</td><td>Gain an understanding of work completed along with a runbook for onboarding additional projects.</td><td>2</td></tr><tr><td><strong>TOTAL HOURS</strong></td><td></td><td><strong>14</strong></td></tr></tbody></table>

## Snyk IAC+ Configuration

### Delivery approach - Snyk IAC+ configuration

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) using the UI import functionality to import into the Customer’s SCM integration.

#### Interpreting and actioning IAC+ results

The Snyk Consultant will educate the Customer on understanding Snyk IaC+ results through the CLI and Snyk UI and how to manage Snyk IAC+ results using Snyk Reporting.

#### SCM integration settings

The Snyk Consultant will work with the Customer to configure SCM Integration settings based on the Customer’s desired gating strategy.

#### SCM Broker installation

The Snyk Consultant will work with the Customer to install the Snyk Broker in a predetermined environment that follows the [Snyk Broker system requirements](../../enterprise-configuration/snyk-broker/prepare-snyk-broker-for-deployment.md).

#### Single pipeline CI/CD CLI configuration

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk iac test` command to provide the Customer with an understanding of how to configure additional pipeline scans and fix cloud issues.

#### Configure cloud environments (up to three environments)

The Snyk Consultant will work with the Customer to configure up to three (3) Cloud environments (Azure, GCP, AWS) in Snyk using the UI.

#### Documentation close-out

The customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, the document offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document puts forth practical recommendations and actionable next steps that will aid the Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document provides an essential guide for customers to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk IAC+ configuration

| Snyk IAC+ configuration                                 | Outcome                                                                                                             | Hours  |
| ------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------ |
| Repository import (SCM only up to 50 targets)           | Import a maximum of 50 targets into Snyk via a [supported SCM Integration](https://snyk.io/integrations/?type=scm). | 2      |
| Interpreting and actioning IAC+ Results                 | Gain an understanding of how to view IaC+ results in Snyk Reporting along with managing misconfigurations.          | 1.5    |
| SCM integration settings                                | Configure SCM integration settings to the Customer’s desired gating settings.                                       | .5     |
| SCM Broker installation                                 | Install SCM Broker in a pre-determined customer environment based on Snyk System Requirements.                      | 2      |
| Single Pipeline CI/CD CLI configuration                 | Configure a single pipeline to test and provide a report for Snyk IAC+.                                             | 4      |
| Configure Cloud environments (up to three environments) | Customer will configure up to three Cloud Eenvironments (Azure, GCP, AWS) with Snyk.                                | 3      |
| Documentation close-out                                 | Gain an understanding of work completed along with a runbook for onboarding additional projects.                    | 2      |
| **TOTAL HOURS**                                         |                                                                                                                     | **15** |

## Snyk AppRisk prioritization configuration

This portion of the Jumpstart service is applicable to Customers looking to achieve [risk-based prioritization](../../manage-risk/prioritize-issues-for-fixing/prioritization-for-snyk-apprisk.md).

### Delivery approach - Snyk AppRisk configuration

#### Coverage and visibility configuration

The Snyk Consultant will work with the Customer to configure their Source Code Management system integration(s) and create two starter policies to show coverage gaps and asset classification, respectively.

#### **Walkthrough of coverage and visibility use cases in AppRisk**

The Snyk Consultant will educate the Customer on how to identify assets that are not currently being scanned by Snyk, as well as how to group assets and issues based on asset classification.

#### Snyk Container - single CICD CLI configuration

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk container test` and `snyk container monitor` commands and provide the Customer with an understanding of how to configure additional pipeline scans.

#### **Snyk Container - single Broker Container Registry installation and configuration**

The Snyk Consultant will work with the Customer to configure and install Broker, if needed, for a single [supported Container Registry](https://snyk.io/integrations/?type=container-registries) integration.

#### **Snyk Container - Container Registry Import**

The Snyk Consultant will work with the Customer to import the set of Container Rargets for container images running in a single Kubernetes cluster either through UI Import functionality or using the Snyk API Import tool.

#### **Single Kubernetes Connector for AppRisk Installation**

The Snyk Consultant will work with the Customer to install the [Kubernetes Connector for AppRisk](../../manage-risk/prioritize-issues-for-fixing/set-up-insights-for-snyk-apprisk/set-up-insights-kubernetes-connector.md) on a predetermined Kubernetes Cluster that meets the Snyk system requirements.

#### **Walkthrough of Prioritized Issues in AppRisk Dashboard**

The Snyk Consultant will educate the Customer on how to filter and prioritize issues in the AppRisk Dashboard using deployed and public-facing risk factors.

#### Documentation close-out

The Customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, it offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document puts forth practical recommendations and actionable next steps that will aid Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document is an essential guide for customers to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk AppRisk prioritization

<table><thead><tr><th width="322">Snyk AppRisk Prioritization Configuration</th><th width="299">Outcome</th><th>Hours</th></tr></thead><tbody><tr><td>Coverage and visibility configuration</td><td>SCM integration is configured in AppRisk and two starter policies are created to show coverage gaps and asset classifications respectively.</td><td>2</td></tr><tr><td>Walk-through of coverage and visibility use cases in AppRisk</td><td>Gain an understanding of how to identify assets that are not currently being scanned by one more Snyk controls, as well as how to group assets and issues based on asset classification</td><td>.5</td></tr><tr><td><p>Snyk Container for AppRisk *</p><p>Setup of one of the following integration methods</p><ul><li>Single Broker Container Registry installation and donfiguration</li><li>Single CI/CD CLI Integration</li></ul></td><td><p>One of:</p><p>Single Broker installed and configured for a <a href="https://snyk.io/integrations/?type=container-registries">Supported Container Registry</a> or</p><p>configure a single pipeline to <code>test</code> and <code>monitor</code> for Snyk Container, including application of component tags according to the <a href="../../manage-risk/prioritize-issues-for-fixing/set-up-insights-for-snyk-apprisk/set-up-insights-associating-snyk-open-source-code-and-container-projects.md">product requirements</a>.</p></td><td>2</td></tr><tr><td><p>Snyk Container for AppRisk *</p><p>Container Registry import</p></td><td><p>Import the set of Container targets</p><p>into Snyk via a <a href="https://snyk.io/integrations/?type=container-registries">Supported Container Registry</a>.</p></td><td>2</td></tr><tr><td>Single Kubernetes Connector for AppRisk installation</td><td><a href="../../manage-risk/prioritize-issues-for-fixing/set-up-insights-for-snyk-apprisk/set-up-insights-kubernetes-connector.md">Kubernetes connector for AppRisk</a> is installed in a single Kubernetes cluster</td><td>3</td></tr><tr><td>Component tagging automation</td><td>Tagging rules are configured for Snyk Open Source, Code, and Container projects (if using Container Registry integration) according to the <a href="../../manage-risk/prioritize-issues-for-fixing/set-up-insights-for-snyk-apprisk/set-up-insights-associating-snyk-open-source-code-and-container-projects.md">product requirements</a>.<br>Application of tags is automed through the<a href="https://github.com/snyk-labs/snyk-tags-tool/blob/main/docs/components.md"> snyk-tags</a> tool in the Customer environment to tag new projects periodically.</td><td>8</td></tr><tr><td>Walk-through of prioritized issues in AppRisk</td><td>Gain an understanding of how to filter and prioritize issues in the AppRisk Dashboard using deployed and public-facing risk factors.</td><td>1.5</td></tr><tr><td>Documentation close-out</td><td>Gain an understanding of work completed along with a runbook for onboarding additional projects.</td><td>2</td></tr><tr><td><strong>TOTAL HOURS</strong></td><td></td><td><strong>21</strong></td></tr></tbody></table>

{% hint style="info" %}
\* _Snyk Container for AppRisk_ steps are required when the Customer does not already have Snyk Container on contract.
{% endhint %}

## Timeline for Snyk Jumpstart delivery

Snyk Jumpstart delivery will include initial platform configuration and each product module that has been purchased. Modules have been designed to be delivered consecutively and over consecutive hours.

| Product                     | Total time   |
| --------------------------- | ------------ |
| Snyk Platform configuration | 9 hours      |
| Snyk Open Source            | 14 hours     |
| Snyk Code                   | 10 hours     |
| Snyk Container              | 14 hours     |
| Snyk IAC+                   | 15 hours     |
| Snyk AppRisk Prioritization | 21 hours     |
| **TOTAL HOURS**             | **83 hours** |

## Additional terms

The fees for this project will be a fixed price. Services will be invoiced in full at the time of purchase and are non-refundable.

The Customer will engage Snyk for a kickoff call within 30 days of the contract start date at a time that is mutually agreed upon by the parties. Snyk Jumpstart must be delivered within 120 days of execution of the applicable Order Form, regardless of when or if the Customer engages Snyk for the kickoff call.

Unless otherwise agreed to by the parties in writing, a) services must be scheduled over consecutive hours; b) product modules will be delivered consecutively; and c) scheduled during normal business hours.

All services will be performed remotely. Any onsite time requires Snyk’s prior consent and will be subject to additional fees and expenses to be paid in accordance with the Snyk Travel and Expense Policy.

## Key assumptions

The following assumptions are reflected in the services outlined in this Jumpstart Services description:

1. All services will be performed remotely using video conferencing software such as Zoom.
2. The Customer must provide prompt feedback on all deliverables.
3. The Customer’s Snyk subject matter expert must be available to work remotely with the Snyk consultant for the entirety of the engagement.
4. The Customer will provide Snyk with documentation and access to subject matter experts for non-Snyk systems and software if required within the scope of the engagement.
5. The Customer will have identified key personnel prior to the beginning of the engagement.
6. Services will be scheduled and delivered during Snyk’s normal business hours, 8 am to 5 pm local time.
7. The Customer will provide prompt access to all systems and resources that Snyk will need in order to complete the work.
8. Snyk does not provide support for third-party software that is used as part of the Snyk solution, such as version control systems, repository management, trouble ticketing systems, packaging, and other software that is not part of the Snyk stack.
9. If a Broker is required, the Customer will have all system requirements before services start.
10. All services and communications will be conducted in English language.
