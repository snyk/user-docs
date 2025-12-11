# Snyk Jumpstart Basic Services Description

## Overview of Snyk Jumpstart Basic

A Snyk Consultant will provide services to help the Customer accelerate its setup of Snyk products through assisted account configuration (the “Jumpstart Basic Services”). The engagement will consist of knowledge transfer, paired with configuration guidance for your team.

The objective is a working setup of Snyk, ready for the Customer team to start engaging with their developers to build their application security efforts.

**Jumpstart Basic is recommended for**:

* Teams needing assistance with the setup of Snyk products

## Jumpstart Basic services description

The Snyk Consultant will deliver the following services related to the setup of Snyk remotely as part of the Jumpstart Basic Services to the Customer. Note that the Jumpstart Basic Services will be delivered only for the Snyk Applications purchased on an Order Form at the same time as this service. References to Snyk Services not purchased by the Customer are hereby omitted unless otherwise noted.

1. [Pre-engagement planning and preparation](#pre-engagement-planning-and-preparation)
   1. Review deliverables per product module
   2. Review prerequisites per product module
   3. Confirm availability for Customer contacts based on [Customer prerequisites](https://docs.snyk.io/snyk-data-and-governance/snyk-terms-of-support-and-services-glossary/snyk-jumpstart-customer-prerequisites)
2. [Snyk Platform configuration](#snyk-platform-configuration)
3. [Snyk Open Source configuration](#snyk-open-source-configuration)
4. [Snyk Code configuration](#snyk-code-configuration)
5. [Snyk Container configuration](#snyk-container-configuration)
6. [Snyk IaC configuration](#snyk-iac-configuration)

## Pre-engagement planning and preparation

The customer should review the [prerequisites](https://docs.snyk.io/snyk-data-and-governance/snyk-terms-of-support-and-services-glossary/snyk-jumpstart-customer-prerequisites), including resources, availability, and deliverables for each product module, before starting their services. The Customer acknowledges that complying with these prerequisites is its sole responsibility, and Snyk will not be responsible for any delays or failure to deliver the Jumpstart Basic Services based on the Customer’s failure to meet these prerequisites.

## Snyk Platform configuration

### Delivery approach - Snyk Platform configuration

The Snyk delivery method is designed to ensure rapid value realization with Snyk. Throughout our collaboration, Snyk will guide the Customer in setting up a foundational configuration and equip them to expand this setup to other applications and integrations. Ensuring Snyk is correctly set up from the outset improves developer adoption and paves the way for long-term success.

#### Base SSO configuration

The Snyk Consultant will work with the Customer to configure SSO through a SAML connection, allowing users to gain access to the Snyk Group with a simple 'org collaborator' or 'group member' permissioning structure.

#### Template Organization configuration

The Snyk Consultant will work with the Customer to configure notification settings, Open Source settings, Code settings, and Snyk account structure. Time will be spent reviewing configuration details with the Customer to ensure the Customer understands and maintains them.

#### Coverage and visibility configuration

The Snyk Consultant will work with Customer to configure your Asset Inventory as follows:

* Configure one Group-level Source Code Management (SCM) integration
* Configure asset policies. The Snyk Consultant will guide the Customer through the setup and configuration of asset policies, such as:
  * Classify assets visible to Snyk
  * Identify coverage gaps based on purchased Snyk products
  * Notify by email when a criteria (e.g. coverage gap or new repo detected) is met

#### Walkthrough of coverage and visibility use cases in Snyk Essentials

The Snyk Consultant will educate the Customer on how to identify assets not currently being scanned by Snyk and how to group assets and issues based on asset classification.

#### Walkthrough of prioritized issues in Snyk Essentials dashboard

The Snyk Consultant will educate the Customer on how to filter and prioritize issues in the Asset Dashboard.

### Target initiatives - Snyk Platform configuration

<table><thead><tr><th width="300">Platform - initial setup</th><th width="555">Outcome</th></tr></thead><tbody><tr><td>Base SSO connection configuration</td><td>Users can access Snyk with broadly defined role.</td></tr><tr><td>(Template Org) Custom Roles and service account configuration</td><td>Provide a templated organization to replicate and scale your setup of Snyk quickly.</td></tr><tr><td>(Template Org) Notification configuration</td><td>Provide a templated organization to replicate and scale your setup of Snyk quickly.</td></tr><tr><td>(Template Org) Products and features configuration</td><td></td></tr><tr><td>(Template Org) Jira or Slack App configuration</td><td></td></tr><tr><td>Coverage and visibility configuration</td><td>SCM integration is configured at Group level and starter policies are created to show coverage gaps, and asset classifications, respectively.</td></tr><tr><td>Walkthrough of coverage and visibility use cases in Snyk Essentials</td><td>Gain an understanding of how to identify assets that are not currently being scanned by one more Snyk controls, as well as how to group assets and issues based on asset classification.</td></tr><tr><td>Walkthrough of prioritized issues in Snyk Essentials</td><td>Gain an understanding of how to filter and prioritize issues in the Inventory Overview and Reports.</td></tr></tbody></table>

## Snyk Open Source configuration

### Delivery approach - Snyk Open Source configuration

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) through the UI import functionality or the Snyk CLI.

#### SCM integration settings

The Snyk Consultant will work with the Customer to configure SCM integration settings based on the Customer’s desired gating strategy.

#### Single pipeline configuration (direct Integration or CLI)

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk test` and `snyk monitor` commands to provide the Customer with an understanding of how to configure additional pipeline scans.

#### Interpreting and actioning Open Source results

The Snyk Consultant will educate the Customer on understanding Snyk Open Source results through the CLI and Snyk UI and how to manage Snyk Open Source results using Snyk Reporting.

### Target initiatives - Snyk Open Source configuration

<table><thead><tr><th width="300">Snyk Open Source configuration</th><th width="306">Outcome</th></tr></thead><tbody><tr><td>Repository import (one SCM integration)</td><td>Import a maximum of 50 targets into Snyk using a <a href="../../developer-tools/scm-integrations/organization-level-integrations">supported SCM integration</a> (GitHub, Azure Repos, Bitbucket, GitLab) or Snyk CLI</td></tr><tr><td>SCM integration settings</td><td>Configure SCM integration settings to the Customer’s desired gating settings.</td></tr><tr><td>Single pipeline configuration (direct integration OR CLI)</td><td>Configure a pipeline to run <code>snyk test</code> and <code>snyk monitor</code>.</td></tr><tr><td>Interpreting and actioning Open Source results</td><td>Gain an understanding of how to view Open Source results in Snyk Reporting along with managing issues.</td></tr></tbody></table>

## Snyk Code configuration

### Delivery approach - Snyk Code configuration

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) either through the UI import functionality or the Snyk CLI.

#### SCM integration settings

The Snyk Consultant will work with the Customer to configure SCM integration settings based on the Customer’s desired gating strategy.

#### Interpreting and actioning Code results

The Snyk Consultant will educate the Customer on understanding Snyk Code results through the CLI and Snyk UI and how to manage Snyk Code results using Snyk Reporting.

### Target initiatives - Snyk Code configuration

<table><thead><tr><th width="300">Snyk Code Configuration</th><th>Outcome</th></tr></thead><tbody><tr><td>Repository import (one SCM integration)</td><td>Import a maximum of 50 targets into Snyk using a <a href="../../developer-tools/scm-integrations/organization-level-integrations">supported SCM Integration</a> (GitHub, Azure Repos, Bitbucket, GitLab).</td></tr><tr><td>SCM integration settings</td><td>Configure SCM integration settings to the Customer’s desired gating settings.</td></tr><tr><td>Interpreting and actioning Code results</td><td>Gain an understanding of how to view Code results in Snyk Reporting along with managing issues.</td></tr></tbody></table>

## Snyk Container configuration

### Delivery approach - Snyk Container configuration

#### Container Registry import

The Snyk Consultant will work with the Customer to import their container images into Snyk (up to 50 targets) through the UI import functionality or Snyk CLI.

#### Single CI/CD CLI configuration

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk container test` and `snyk container monitor` commands to provide the Customer with an understanding of how to configure additional pipeline scans.

#### Interpreting and actioning Snyk Container results

The Snyk Consultant will educate the Customer on understanding Snyk Container results through the CLI and Snyk UI and how to manage Snyk Container results using Snyk Reporting.

### Target initiatives - Snyk Container configuration

<table><thead><tr><th width="300">Snyk Container Configuration</th><th>Outcome</th></tr></thead><tbody><tr><td>Container Registry import (up to 50 targets)</td><td>Import a maximum of 50 targets into Snyk using a <a href="https://snyk.io/integrations/?type=container-registries">Supported Container Registry</a>.</td></tr><tr><td>Single CI/CD CLI integration</td><td>Configure a single pipeline to <code>test</code> and <code>monitor</code> for Snyk Container.</td></tr><tr><td>Interpreting and actioning Container results</td><td>Gain an understanding of how to view Container results in Snyk Reporting along with managing issues.</td></tr></tbody></table>

## Snyk IaC Configuration

### Delivery approach - Snyk IaC configuration

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) through the UI import functionality or the Snyk CLI.

#### IaC Settings

The Snyk Consultant will work with the Customer to configure IaC settings based on the Customer’s preferences.

#### Single Pipeline CI/CD CLI configuration

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk iac test` and `snyk iac test --report` commands to provide the Customer with an understanding of how to configure additional pipeline scans and fix misconfigurations.

#### Interpreting and actioning IaC results

The Snyk Consultant will educate the Customer on understanding Snyk IaC results through the CLI and Snyk UI and how to manage Snyk IaC results using Snyk Reporting.

### Target initiatives - Snyk IaC configuration

| Snyk IaC configuration                  | Outcome                                                                                                                          |
| --------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| Repository import (one SCM integration) | Import a maximum of 50 targets into Snyk using a supported SCM integration (GitHub, Azure Repos, Bitbucket, GitLab) or Snyk CLI. |
| IaC Settings                            | Configure IaC settings to Customer’s desired preferences.                                                                        |
| Single Pipeline CI/CD CLI configuration | Configure a pipeline to run snyk iac test and snyk iac test --report.                                                            |
| Interpreting and actioning IaC Results  | Gain an understanding of how to view IaC results in Snyk Reporting along with managing misconfigurations.                        |


## Timeline for Snyk Jumpstart Basic delivery

Snyk Jumpstart Basic delivery is a 30-day engagement that begins with the Jumpstart Basic Kickoff call.

For customers starting their journey with Snyk, Jumpstart Basic includes initial Platform Configuration and each product module that has been purchased.

If the Jumpstart Basic project is for an existing Snyk customer, the Platform Configuration content is adjusted to reflect this, with the Snyk consultant providing a ‘healthcheck’ review of the existing configuration, before working through the onboarding of new features or products.


## Additional terms

The fees for this project will be a fixed price. Services will be invoiced in full at the time of purchase and are non-refundable.

Upon purchase, the Snyk implementation team will coordinate a pre-engagement call at a time that works for both parties. Note that Jumpstart Basic services expire 60 days after our consultant's first outreach, regardless of customer engagement. All Jumpstart Basic Services automatically conclude either upon project completion or 30 days following the kick-off call, whichever occurs first.

Unless otherwise agreed to by the parties in writing, a) services must be scheduled as agreed prior to the Kickoff call; b) product modules will be delivered consecutively; and c) all services will occur during normal business hours.

All services will be performed remotely.

## Key assumptions

The following assumptions are reflected in the services outlined in this Jumpstart Basic Services description:

1. All services will be performed remotely using video conferencing software such as Google Meet.
2. The Customer must provide prompt feedback on all deliverables.
3. The Customer will appoint one subject matter expert who will be the point of contact for the Jumpstart Basic Services. This subject matter expert must be available to work remotely with the Snyk Consultant for the entirety of the engagement.
4. The Customer will provide Snyk with documentation and access to subject matter experts for non-Snyk systems and software if required within the scope of the engagement.
5. The Customer will have identified key personnel prior to the beginning of the engagement.
6. Services will be scheduled and delivered during Snyk’s normal business hours: 8 am to 5 pm local time zone Monday through Friday (Sunday through Thursday where applicable based on the region of the assigned Snyk Consultant).
7. The Customer will provide prompt access to all systems and resources that Snyk will need in order to complete the work.
8. Snyk does not provide support for third-party software that is used as part of the Snyk solution, such as version control systems, repository management, trouble ticketing systems, packaging, and other software that is not part of the Snyk stack.
9. All services and communications will be conducted in English.