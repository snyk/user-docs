# Snyk Jumpstart Basic Services Description

A Snyk Consultant will provide services to help the Customer accelerate its setup of Snyk products through assisted account configuration (the “Jumpstart Basic Services”). The engagement will consist of knowledge transfer, paired with configuration guidance for your team.

The objective is to establish a working setup of Snyk, ready for the Customer team to collaborate with their developers to enhance their application security efforts.

{% hint style="info" %}
Jumpstart Basic is recommended for teams that need assistance with setting up Snyk products.
{% endhint %}

## Jumpstart Basic Services Description

The Snyk Consultant delivers the following Jumpstart Basic Services for remote Snyk setup. These services apply only to Snyk applications purchased concurrently on an Order Form. References to unpurchased Snyk Services are omitted unless otherwise noted.

1. [Pre-engagement planning and preparation](snyk-jumpstart-basic-services-description.md#pre-engagement-planning-and-preparation)
   1. Review deliverables per product module
   2. Review prerequisites per product module
   3. Confirm availability for Customer contacts based on [Customer prerequisites](https://docs.snyk.io/snyk-data-and-governance/snyk-terms-of-support-and-services-glossary/snyk-jumpstart-customer-prerequisites)
2. [Snyk Platform configuration](snyk-jumpstart-basic-services-description.md#snyk-platform-configuration)
3. [Snyk Open Source configuration](snyk-jumpstart-basic-services-description.md#snyk-open-source-configuration)
4. [Snyk Code configuration](snyk-jumpstart-basic-services-description.md#snyk-code-configuration)
5. [Snyk Container configuration](snyk-jumpstart-basic-services-description.md#snyk-container-configuration)
6. [Snyk IaC configuration](https://gist.github.com/cgibbs-snyk/2b73b280c946c3323cde65b4e2cfd88f#snyk-iac-configuration)

## Pre-engagement planning and preparation

Review the [prerequisites](https://docs.snyk.io/snyk-data-and-governance/snyk-terms-of-support-and-services-glossary/snyk-jumpstart-customer-prerequisites), including resources, availability, and deliverables for each product module, before starting your services. You are responsible for meeting these prerequisites. Snyk is not responsible for delays or undelivered Jumpstart Basic Services if you do not meet these prerequisites.

## Snyk Platform configuration

#### **Base SSO configuration**

A Snyk Consultant configures SSO using SAML to give users access to the Snyk Group as Org collaborators or Group members.

#### **Template Organization configuration**

A Snyk Consultant helps you configure notification settings, Snyk Open Source settings, Snyk Code settings, and your Snyk account structure. The Consultant reviews these details with you, ensuring you understand and maintain them.

#### **Coverage and visibility configuration**

A Snyk Consultant configures your Asset Inventory. This includes:

* Configuring one Group-level Source Code Management (SCM) integration.
* Guiding you through configuring asset policies. These policies include:
  * Classifying assets visible to Snyk.
  * Identifying coverage gaps based on purchased Snyk products.
  * Sending email notifications when criteria are met (for example, coverage gaps or new repositories).

#### **Walkthrough of coverage and visibility use cases**&#x20;

A Snyk Consultant shows you how to identify assets Snyk does not scan and how to group assets and issues using asset classification.

#### **Walkthrough of prioritized issues**

A Snyk Consultant educates customers on filtering and prioritizing issues in the Asset Dashboard.

#### Delivery approach

The Snyk delivery method helps you achieve value quickly. Snyk guides you to set up a foundational configuration and equips you to expand it to other applications and integrations. An initial, correct Snyk setup improves developer adoption and supports long-term success.

### Target initiatives

<table><thead><tr><th width="300">Platform - initial setup</th><th width="555">Outcome</th></tr></thead><tbody><tr><td>Base SSO connection configuration</td><td>Users can access Snyk with broadly defined role.</td></tr><tr><td>(Template Org) Custom Roles and service account configuration</td><td>Templated organizations help you quickly replicate and scale your Snyk setup.</td></tr><tr><td>(Template Org) Notification configuration</td><td>Provide a templated organization to quickly replicate and scale your notification setup.</td></tr><tr><td>(Template Org) Products and features configuration</td><td>Provide a templated organization to quickly replicate and scale your product and features setup.</td></tr><tr><td>(Template Org) Jira or Slack App configuration</td><td>Provide a templated organization to quickly replicate and scale your Jira or Slack setup.</td></tr><tr><td>Coverage and visibility configuration</td><td>Snyk configures SCM integration for the Group and creates starter policies to show coverage gaps, and asset classifications.</td></tr><tr><td>Walkthrough of coverage and visibility use cases</td><td>Identify assets not scanned by Snyk controls, and group assets and issues by asset classification.</td></tr><tr><td>Walkthrough of prioritized issues</td><td>Filter and prioritize issues using the Inventory Overview and Reports options.</td></tr></tbody></table>

## Snyk Open Source configuration

#### **Repository import**

Snyk consultants work with you to import your repositories (up to 50 targets) into Snyk using the Snyk UI or Snyk CLI.

#### **SCM integration settings**

A Snyk consultant helps you configure SCM integration settings using your chosen gating strategy.

#### **Single pipeline configuration (direct Integration or CLI)**

A Snyk consultant works with you to configure a single pipeline to run the `snyk test` and `snyk monitor` commands, providing you with an understanding of how to configure additional pipeline scans.

#### **Interpreting and actioning Open Source results**

A Snyk consultant educates you on understanding Snyk Open Source results through the CLI and Snyk Web UI, and how to manage Snyk Open Source results using the Reports option.

### Target initiatives

<table><thead><tr><th width="300">Snyk Open Source configuration</th><th width="306">Outcome</th></tr></thead><tbody><tr><td>Repository import (one SCM integration)</td><td>Import a maximum of 50 targets into Snyk using a <a href="https://gist.github.com/developer-tools/scm-integrations/organization-level-integrations">supported SCM integration</a> (GitHub, Azure Repos, Bitbucket, GitLab) or Snyk CLI</td></tr><tr><td>SCM integration settings</td><td>Configure SCM integration settings to your gating settings.</td></tr><tr><td>Single pipeline configuration (direct integration OR CLI)</td><td>Configure a pipeline to run <code>snyk test</code> and <code>snyk monitor</code>.</td></tr><tr><td>Interpreting and actioning Open Source results</td><td>View open-source results and manage issues in the Reports option.</td></tr></tbody></table>

## Snyk Code configuration

#### **Repository import**

A Snyk consultant helps you to import your repositories into Snyk (up to 50 targets) either through the Snyk Web UI import functionality or the Snyk CLI.

#### **SCM integration settings**

A Snyk consultant works with you to configure SCM integration settings based on your desired gating strategy.

#### **Interpreting and actioning Code results**

A Snyk consultant helps you understand the Snyk Code results through the CLI and Snyk Web UI and how to manage Snyk Code results using the Reports option.

### Target initiatives

<table><thead><tr><th width="300">Snyk Code Configuration</th><th>Outcome</th></tr></thead><tbody><tr><td>Repository import (one SCM integration)</td><td>Import a maximum of 50 targets into Snyk using a <a href="https://gist.github.com/developer-tools/scm-integrations/organization-level-integrations">supported SCM Integration</a> (GitHub, Azure Repos, Bitbucket, GitLab).</td></tr><tr><td>SCM integration settings</td><td>Configure SCM integration settings to match your gating settings.</td></tr><tr><td>Interpreting and actioning Code results</td><td>View Snyk Code results and manage issues in the Reports option.</td></tr></tbody></table>

## Snyk Container configuration

#### **Container Registry import**

A Snyk Consultant helps you to import your container images into Snyk (up to 50 targets) through the Snyk Web UI import functionality or Snyk CLI.

#### **Single CI/CD CLI configuration**

A Snyk Consultant works with you to configure a single pipeline to run the `snyk container test` and `snyk container monitor` commands to provide you with an understanding of how to configure additional pipeline scans.

#### **Interpreting and actioning Snyk Container results**

A Snyk Consultant helps you understand the Snyk Container results through the CLI and Snyk Web UI, and how to manage Snyk Container results using the Reports option.

### Target initiatives

<table><thead><tr><th width="300">Snyk Container Configuration</th><th>Outcome</th></tr></thead><tbody><tr><td>Container Registry import (up to 50 targets)</td><td>Import a maximum of 50 targets into Snyk using a <a href="https://snyk.io/integrations/?type=container-registries">Supported Container Registry</a>.</td></tr><tr><td>Single CI/CD CLI integration</td><td>Configure a single pipeline to <code>test</code> and <code>monitor</code> for Snyk Container.</td></tr><tr><td>Interpreting and actioning Container results</td><td>Gain an understanding of how to view Container results in Snyk Reporting along with managing issues.</td></tr></tbody></table>

## Snyk IaC Configuration

#### **Repository import**

A Snyk Consultant helps you to import your repositories into Snyk (up to 50 targets) through the Snyk Web UI import functionality or the Snyk CLI.

#### **IaC Settings**

A Snyk Consultant helps you to configure IaC settings.

#### **Single Pipeline CI/CD CLI configuration**

A Snyk Consultant helps you to configure a single pipeline to run the `snyk iac test` and `snyk iac test --report` commands to provide the Customer with an understanding of how to configure additional pipeline scans and fix misconfigurations.

#### **Interpreting and actioning IaC results**

The Snyk Consultant will educate the Customer on understanding Snyk IaC results through the CLI and Snyk UI and how to manage Snyk IaC results using Snyk Reporting.

### Target initiatives

| Snyk IaC configuration                  | Outcome                                                                                                                          |
| --------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| Repository import (one SCM integration) | Import a maximum of 50 targets into Snyk using a supported SCM integration (GitHub, Azure Repos, Bitbucket, GitLab) or Snyk CLI. |
| IaC Settings                            | Configure IaC settings based on your custom preferences.                                                                         |
| Single Pipeline CI/CD CLI configuration | Configure a pipeline to run `snyk iac test` and `snyk iac test --report`.                                                        |
| Interpreting and actioning IaC Results  | View IaC results and manage misconfigurations using the Reports option.                                                          |

## Timeline for Snyk Jumpstart Basic delivery

Snyk Jumpstart Basic is a 30-day engagement, starting with a Jumpstart Basic Kickoff call.

For new Snyk customers, Jumpstart Basic includes initial Platform Configuration and all purchased product modules.

For existing Snyk customers, a Snyk Consultant adjusts the Platform Configuration content. This includes a health check review of the existing configuration before onboarding new features or products.

## Additional terms

Project fees are fixed. Services are invoiced in full at purchase and are non-refundable.

After purchase, the Snyk implementation team coordinates a pre-engagement call at a mutually agreed time. Jumpstart Basic services expire 60 days after the first Snyk Consultant outreach, regardless of customer engagement. Jumpstart Basic services conclude automatically after project completion or 30 days after the Kickoff call, whichever occurs first.

Unless otherwise agreed to by the parties in writing, a) services must be scheduled as agreed during the  call; b) product modules will be delivered consecutively; and c) all services will occur during normal business hours.

All services are remote.

## Key assumptions

The following assumptions are reflected in the services outlined in this Jumpstart Basic Services description:

* All services will be performed remotely using video conferencing.
* The customer provides prompt feedback on deliverables.
* The customer appoints one subject matter expert (SME) as the contact for Jumpstart Basic Services. The SME is available to work remotely with the Snyk Consultant throughout the engagement.
* The customer provides Snyk with documentation and access to SMEs for non-Snyk systems and software when required.
* The customer identifies key personnel before the engagement starts.
* Snyk delivers services during normal business hours: Monday through Friday, 8 am to 5 pm local time. Regional variations apply (for example, Sunday through Thursday). Regional variations are based on the region of the assigned Snyk consultant.
* The customer provides prompt access to all systems and resources Snyk needs to complete the work.
* Snyk does not support third-party software used with the Snyk solution, for example, version control systems, repository management, or trouble-ticketing systems.
* Snyk conducts all services and communications in English.
