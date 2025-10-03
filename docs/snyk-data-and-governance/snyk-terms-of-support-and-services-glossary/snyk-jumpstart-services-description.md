# Snyk Jumpstart Services Description

## Overview of Snyk Jumpstart

A Snyk Consultant will provide services to help the Customer accelerate its setup of Snyk products through assisted account configuration (the “Jumpstart Services”). The engagement will consist of knowledge transfer, paired with configuration guidance for your team. 

The objective is a working setup of Snyk and a Customer team that is well-prepared to continue its application security efforts.

**Jumpstart is recommended for**:

* Teams needing assistance with the setup of Snyk products
* Teams who are new to Snyk and have limited experience with security scanning
* Self-starters who are comfortable with extending and maintaining Snyk on their own, post-engagement

## Jumpstart services description

The Snyk Consultant will deliver the following services related to the setup of Snyk remotely as part of the Jumpstart Services to the Customer. Note that the Jumpstart Services will be delivered only for the Snyk Applications purchased on an Order Form at the same time as this service. References to Snyk Services not purchased by the Customer are hereby omitted unless otherwise noted. 

1. [Pre-engagement planning and preparation](snyk-jumpstart-services-description.md#pre-engagement-planning-and-preparation)
   1. Review deliverables per product module
   2. Review prerequisites per product module
   3. Confirm availability for Customer contacts based on [Customer prerequisites](snyk-jumpstart-customer-prerequisites.md)
2. [Snyk Platform configuration](snyk-jumpstart-services-description.md#snyk-platform-configuration)
3. [Snyk Open Source configuration](snyk-jumpstart-services-description.md#snyk-open-source-configuration)
4. [Snyk Code configuration](snyk-jumpstart-services-description.md#snyk-code-configuration)
5. [Snyk Container configuration](snyk-jumpstart-services-description.md#snyk-container-configuration)
6. [Snyk IaC configuration](snyk-jumpstart-services-description.md#snyk-iac-configuration)
7. [Snyk API & Web configuration](snyk-jumpstart-services-description.md#snyk-api-and-web-configuration-details)

## Pre-engagement planning and preparation

A pre-engagement call will be held before the start of the Jumpstart Services listed herein, to ensure the Customer understands the prerequisites required for engagement start, along with the deliverables to be completed during the engagement timeframe. Customer [prerequisites](snyk-jumpstart-customer-prerequisites.md), including resources, availability, and deliverables for each product module, will be reviewed and confirmed. The Customer acknowledges that complying with these prerequisites is its sole responsibility, and Snyk will not be responsible for any delays or failure to deliver the Jumpstart Services based on the Customer’s failure to meet these prerequisites.

## Snyk Platform configuration

### Delivery approach - Snyk Platform configuration

The Snyk delivery method is designed to ensure rapid value realization with Snyk. Throughout our collaboration, Snyk will guide the Customer in setting up a foundational configuration and equip them to expand this setup to other applications and integrations. Ensuring Snyk is correctly set up from the outset improves developer adoption and paves the way for long-term success.

#### Base SSO configuration and custom mapping

The Snyk Consultant will work with the Customer to configure SSO through SAML, Entra ID (formerly Azure AD), OIDC, or ADFS connection along with custom mapping of dynamically assigned users to Snyk Groups and Organizations based on data provided by the desired Identity Provider (IdP) to set up a scaled user provisioning and access model.

#### **Template Organization configuration**

The Snyk Consultant will work with the Customer to configure notification settings, Open Source settings, Code settings, and Snyk account structure. Time will be spent reviewing configuration details with the Customer to ensure the Customer understands and maintains them.

#### Coverage and visibility configuration

The Snyk Consultant will work with Customer to configure your Asset Inventory as follows:

* Configure one Group-level Source Code Management (SCM) integration, and if necessary, Snyk Broker for this integration
* Configure one application context provider integration
* Configure asset policies. The Snyk Consultant will guide the Customer through the setup and configuration of asset policies, such as:
  * Classify assets visible to Snyk
  * Identify coverage gaps based on purchased Snyk products
  * Notify by email when a criteria (e.g. coverage gap or new repo detected) is met

#### **Walkthrough of coverage and visibility use cases in Snyk Essentials**

The Snyk Consultant will educate the Customer on how to identify assets not currently being scanned by Snyk and how to group assets and issues based on asset classification.

#### **Walkthrough of prioritized issues in Snyk Essentials dashboard**

The Snyk Consultant will educate the Customer on how to filter and prioritize issues in the Asset Dashboard.

#### Developer training

The Snyk Consultant will run a 60-minute training session for Customer’s Developers to ensure all collaborator users know how to access Snyk results and initiate tests. This training will cover a range of tools, including the UI (app.snyk.io), IDE plugin, and CLI. Developer Training will be completed after the product-specific modules of Jumpstart have been completed.

### Target initiatives - Snyk Platform configuration

<table><thead><tr><th width="300">Platform - initial setup</th><th width="555">Outcome</th></tr></thead><tbody><tr><td>Base SSO connection configuration</td><td>Users can access Snyk with broadly defined role.</td></tr><tr><td>SSO Custom Mapping</td><td>Users can access Snyk based on role defined in IdP.</td></tr><tr><td>(Template Org) Custom Roles and service account configuration</td><td>Provide a templated organization to replicate and scale your setup of Snyk quickly.</td></tr><tr><td>(Template Org) Notification configuration</td><td>Provide a templated organization to replicate and scale your setup of Snyk quickly.</td></tr><tr><td>(Template Org) Products and features configuration</td><td></td></tr><tr><td>(Template Org) Jira or Slack App configuration</td><td></td></tr><tr><td>Coverage and visibility configuration</td><td>SCM integration is configured at Group level and starter policies are created to show coverage gaps, and asset classifications, respectively.</td></tr><tr><td>Walkthrough of coverage and visibility use cases in Snyk Essentials</td><td>Gain an understanding of how to identify assets that are not currently being scanned by one more Snyk controls, as well as how to group assets and issues based on asset classification.</td></tr><tr><td>Walkthrough of prioritized issues in Snyk Essentials</td><td>Gain an understanding of how to filter and prioritize issues in the Inventory Overview and Reports.</td></tr><tr><td>Developer training (1 session)</td><td>Delivered after product-specific modules are completed, ensuring a successful launch and initial adoption with Developers.</td></tr></tbody></table>

## Snyk Open Source configuration

### Delivery approach - Snyk Open Source configuration

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) through the UI import functionality, the [API Import tool](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/), or the Snyk CLI.

#### SCM integration settings

The Snyk Consultant will work with the Customer to configure SCM integration settings based on the Customer’s desired gating strategy.

#### SCM Broker installation

The Snyk Consultant will work with the Customer to install the Snyk Broker in a pre-determined environment that follows the [Snyk Broker system requirements](../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/).

#### Single pipeline configuration (direct Integration or CLI)

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk test` and `snyk monitor` commands to provide the Customer with an understanding of how to configure additional pipeline scans.

#### SBOM walkthrough (API and CLI)

The Snyk Consultant will educate the Customer on creating an SBOM through the Snyk API and the Snyk CLI.

#### Interpreting and actioning Open Source results

The Snyk Consultant will educate the Customer on understanding Snyk Open Source results through the CLI and Snyk UI and how to manage Snyk Open Source results using Snyk Reporting.

### Target initiatives - Snyk Open Source configuration

<table><thead><tr><th width="300">Snyk Open Source configuration</th><th width="306">Outcome</th></tr></thead><tbody><tr><td>Repository import (one SCM integration)</td><td>Import a maximum of 50 targets into Snyk using a <a href="../../developer-tools/scm-integrations/organization-level-integrations/">supported SCM integration</a> (GitHub, Azure Repos, Bitbucket, GitLab) or Snyk CLI</td></tr><tr><td>SCM integration settings</td><td>Configure SCM integration settings to the Customer’s desired gating settings.</td></tr><tr><td>SCM Broker installation</td><td>Install SCM Broker in a pre-determined customer environment based on Snyk system requirements.</td></tr><tr><td>Single pipeline configuration (direct integration OR CLI)</td><td>Configure a pipeline to run <code>snyk test</code> and <code>snyk monitor</code>.</td></tr><tr><td>SBOM Walkthrough (CLI and API)</td><td>Gain an understanding of generating an SBOM through Snyk using the CLI and API.</td></tr><tr><td>Interpreting and actioning Open Source results</td><td>Gain an understanding of how to view Open Source results in Snyk Reporting along with managing issues.</td></tr></tbody></table>



## Snyk Code configuration

### Delivery approach - Snyk Code configuration

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) either through the UI import functionality or through the [API Import tool](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/).&#x20;

#### SCM integration settings

The Snyk Consultant will work with the Customer to configure SCM integration settings based on the Customer’s desired gating strategy.

#### SCM Broker installation

The Snyk Consultant will work with the Customer to install the Snyk Broker in a pre-determined environment that follows the [Snyk Broker system requirements](../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/).

#### Interpreting and actioning Code results

The Snyk Consultant will educate the Customer on understanding Snyk Code results through the CLI and Snyk UI and how to manage Snyk Code results using Snyk Reporting.

### Target initiatives - Snyk Code configuration

<table><thead><tr><th width="250">Snyk Code Configuration</th><th>Outcome</th></tr></thead><tbody><tr><td>Repository import (one SCM integration)</td><td>Import a maximum of 50 targets into Snyk using a <a href="../../developer-tools/scm-integrations/organization-level-integrations/">supported SCM Integration</a> (GitHub, Azure Repos, Bitbucket, GitLab).</td></tr><tr><td>SCM integration settings</td><td>Configure SCM integration settings to the Customer’s desired gating settings.</td></tr><tr><td>SCM Broker installation</td><td>Install SCM Broker in a pre-determined customer environment based on Snyk system requirements.</td></tr><tr><td>Interpreting and actioning Code results</td><td>Gain an understanding of how to view Code results in Snyk Reporting along with managing issues.</td></tr></tbody></table>

## Snyk Container configuration

### Delivery approach - Snyk Container configuration

#### Container Registry import

The Snyk Consultant will work with the Customer to import their container images into Snyk (up to 50 targets) through the UI import functionality or Snyk CLI.

#### Single Broker Container Registry installation and configuration

The Snyk Consultant will work with the Customer to configure and install the Snyk Broker and CR Agent if needed for a single [supported Container Registry](https://snyk.io/integrations/?type=container-registries) integration.

#### Interpreting and actioning Snyk Container results

The Snyk Consultant will educate the Customer on understanding Snyk Container results through the CLI and Snyk UI and how to manage Snyk Container results using Snyk Reporting.

#### Single CI/CD CLI configuration

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk container test` and `snyk container monitor` commands to provide the Customer with an understanding of how to configure additional pipeline scans.

#### Custom Base Images walkthrough (UI and CLI)

The Snyk Consultant will educate the Customer on how to use the Snyk Custom Base Image Recommendation functionality.

### Target initiatives - Snyk Container configuration

<table><thead><tr><th width="269">Snyk Container Configuration</th><th>Outcome</th></tr></thead><tbody><tr><td>Single Broker Container Registry installation and configuration</td><td>Install Broker and CR Agent in a pre-determined customer environment based on [Snyk Broker system requirements](../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/).</td></tr><tr><td>Container Registry import (up to 50 targets)</td><td>Import a maximum of 50 targets into Snyk using a <a href="https://snyk.io/integrations/?type=container-registries">Supported Container Registry</a>.</td></tr><tr><td>Interpreting and actioning Container results</td><td>Gain an understanding of how to view Container results in Snyk Reporting along with managing issues.</td></tr><tr><td>Single CI/CD CLI integration</td><td>Configure a single pipeline to <code>test</code> and <code>monitor</code> for Snyk Container.</td></tr><tr><td>Custom Base Images walkthrough</td><td>Gain an understanding of how to use the Custom Base Image Recommendations functionality.</td></tr></tbody></table>

## Snyk IaC Configuration

### Delivery approach - Snyk IaC configuration

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) through the UI import functionality, the [API Import tool](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/), or the Snyk CLI.

#### IaC Settings

The Snyk Consultant will work with the Customer to configure IaC settings based on the Customer’s preferences. 

#### SCM Broker Installation

The Snyk Consultant will work with the Customer to install the Snyk Broker in a pre-determined environment that follows the [Snyk Broker system requirements](../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/).

#### Single Pipeline CI/CD CLI configuration 

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk iac test` and `snyk iac test --report` commands to provide the Customer with an understanding of how to configure additional pipeline scans and fix misconfigurations.

#### Interpreting and actioning IaC results

The Snyk Consultant will educate the Customer on understanding Snyk IaC results through the CLI and Snyk UI and how to manage Snyk IaC results using Snyk Reporting.


### Target initiatives - Snyk IaC configuration

| Snyk IaC configuration                  | Outcome                                                                                                             |
| --------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| Repository import (one SCM integration) | Import a maximum of 50 targets into Snyk using a supported SCM integration (GitHub, Azure Repos, Bitbucket, GitLab) or Snyk CLI. |
| IaC Settings                | Configure IaC settings to Customer’s desired preferences.                                       |
| SCM Broker installation                 | Install SCM Broker in a pre-determined customer environment based on Snyk System Requirements.                      |
| Single Pipeline CI/CD CLI configuration | Configure a pipeline to run snyk iac test and snyk iac test --report.                                              |
| Interpreting and actioning IaC Results  | Gain an understanding of how to view IaC results in Snyk Reporting along with managing misconfigurations.           |


## Snyk API & Web configuration details

### Delivery approach - Snyk API & Web configuration

#### **Web Target configuration, including authenticated scans**

The Snyk Consultant will work with the Customer to configure [Web Targets](https://help.probely.com/en/articles/3292779-how-to-set-up-target-authentication-with-a-login-form?q=Web+Target+configuration) (up to three web applications) to be scanned by Snyk API & Web. This includes configuring the authentication for each Target where necessary, such as using a login form or recorded login sequence. Snyk will also help to set up a single navigation sequence for each web application.

#### **API Target configuration**

The Snyk Consultant will work with the Customer to [configure API collections](https://help.probely.com/en/articles/8178059-how-to-configure-an-api-target-postman-collection) (maximum of one collection) to be scanned by Snyk API & Web, using a Postman Collection or OpenAPI definition.

#### **Domain Ownership Verification**

The Snyk Consultant will work with the Customer to complete [Domain Ownership Verification](https://help.probely.com/en/articles/3289281-how-to-verify-the-ownership-of-a-domain-using-a-txt-file) for one domain. This can be achieved by a .txt file, a TXT record, a CNAME record, or a meta tag.

#### **Scanning Agent configuration**

The Snyk Consultant will work with the Customer to configure the [Scanning Agent](https://help.probely.com/en/articles/4615595-how-to-scan-internal-applications-with-a-scanning-agent) using Docker, Docker-Compose, or Kubernetes. This is required only if there are Targets to be scanned that are not internet-accessible, and there is a maximum of one agent to be configured.

#### **Target scanning in CI/CD configuration**

The Snyk Consultant will work with the Customer to configure an example workflow using the API & Web CLI that allows you to trigger a Target scan when a repo is updated. This can include the opportunity of failing the pipeline if the scan finds vulnerabilities of High severity.

#### **Issue ticketing integration configuration**

The Snyk Consultant will work with the Customer to configure a single issue ticketing integration (e.g. Jira or Azure Boards).

#### **Review of Target Scan Results**

The Snyk Consultant will educate the Customer on understanding [DAST scan results](https://help.probely.com/en/articles/6843262-how-to-interpret-target-scan-results) for Web and API Targets in the Snyk API & Web UI, including the different reporting functionality that is available in the tool.


### Target initiatives - Snyk API & Web configuration

| Snyk API & Web configuration                                                           | Outcome                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Web Target configuration, including Authenticated Scans (up to three web applications) | A maximum of three web applications are configured in Snyk API & Web for scanning, including Authentication. For each application there is a maximum of one login sequence and one navigation sequence). |
| API Target configuration (up to one API collection)                                    | One API collection is configured in Snyk API & Web for scanning.                                                                                                                                         |
| Domain Ownership Verification                                                          | Domain Ownership Verification is completed for one domain.                                                                                                                                               |
| Scanning Agent Configuration (one agent)                                               | A maximum of one Scanning Agent is configured to enable Snyk API & Web to scan internal applications without internet access.                                                                            |
| Target scanning in CI/CD configuration                                               | A single pipeline will be configured that allows you to trigger a test on an existing Web or API Target.                                                                            |
| Issue ticketing integration configuration                                               | A single integration will be configured that allows tickets to be created when findings of a certain severity are discovered.                                                                            |
| Review of Target Scan Results                                                          | Target Scan Results are reviewed to ensure the Customer understands how to process the DAST findings and prioritize fixing the issues.                                                                   |


## Snyk Learning Management Add-On configuration details

### Delivery approach - Snyk Learning Management Add-On configuration

#### **Best Practices and Content Overview**

The Snyk Consultant will share best practices for administering a security awareness program within the Snyk Learn platform, including an overview of the types of content available in Snyk Learn that can be leveraged. Provide insights and recommendations to help the Customer shape their program strategy and select appropriate learning materials.

#### **Custom Roles and User Management**

The Snyk Consultant will guide the Customer on creating custom roles, such as a "Learning Manager" role, with tailored permissions related to the Learning Management add-on. The Consultant will explain the process of assigning these roles to users within the Customer's Organizations, ensuring appropriate access and responsibilities for managing and overseeing the Snyk Learn program.

#### **Configuring Assignments using a Learning Organization**

The Snyk Consultant will educate the Customer on creating Snyk Learn Assignments, setting a due date, and using a ‘Learning Organization’ to specify who needs to complete the Assignment. This includes optional settings, such as resetting user progress to force the lesson to be retaken for compliance reasons.

#### **Tracking Assignments and controlling access**

The Snyk Consultant will walk through how to track the progress for active Assignments to review which employees have completed the course. This also includes reviewing which users can access the main Assignments page where they are tracked, created, and removed. An overview of the different reports available in the tool will help to track progress on usage.

### Target initiatives - Snyk Learning Management Add-On configuration

| Snyk Learning Management Add-On configuration                  | Outcome                                                                                                             |
| --------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| Best Practices and Content Overview | Learning Management administrators understand the content available in Snyk Learn and have an initial plan for how to use the platform with their users. |
| Custom Roles and User Management                | Custom Roles are created and assigned to relevant users in order to provide access to the Learning Management add-on to the relevant users.                                       |
| Configuring Assignments using a Learning Organization                 | Gain an understanding of how to create and send Assignments to members of an Organization.                      |
| Tracking Assignments and controlling access | Gain an understanding of how to review in-progress Assignments and control who has access to review/create/delete Assignments based on Custom Roles.                                              |





## Timeline for Snyk Jumpstart delivery

Snyk Jumpstart delivery is a 60-day engagement that begins with the Jumpstart Kickoff call. 

For customers starting their journey with Snyk, Jumpstart includes initial Platform Configuration and each product module that has been purchased.

If the Jumpstart project is for an existing Snyk customer, the Platform Configuration content is adjusted to reflect this, with the Snyk consultant providing a ‘healthcheck’ review of the existing configuration, before working through the onboarding of new features or products. 

## Additional terms

The fees for this project will be a fixed price. Services will be invoiced in full at the time of purchase and are non-refundable.&#x20;

Upon purchase, the Snyk implementation team will coordinate a pre-engagement call at a time that works for both parties. Note that Jumpstart services expire 120 days after our consultant's first outreach, regardless of customer engagement. All Jumpstart Services automatically conclude either upon completion or eight weeks following the kick-off call, whichever occurs first.

Unless otherwise agreed to by the parties in writing, a) services must be scheduled as agreed during the pre-engagement call; b) product modules will be delivered consecutively; and c) all services will occur during normal business hours.&#x20;

All services will be performed remotely. Any onsite time requires Snyk’s prior consent and will be subject to additional fees and expenses to be agreed in advance.

## Key assumptions

The following assumptions are reflected in the services outlined in this Jumpstart Services description:

1. All services will be performed remotely using video conferencing software such as Zoom.&#x20;
2. The Customer must provide prompt feedback on all deliverables.
3. The Customer will appoint one subject matter expert who will be the point of contact for the Jumpstart Services. This subject matter expert must be available to work remotely with the Snyk Consultant for the entirety of the engagement.
4. The Customer will provide Snyk with documentation and access to subject matter experts for non-Snyk systems and software if required within the scope of the engagement.
5. The Customer will have identified key personnel prior to the beginning of the engagement.
6. Services will be scheduled and delivered during Snyk’s normal business hours: 8 am to 5 pm local time zone Monday through Friday (Sunday through Thursday where applicable based on the region of the assigned Snyk Consultant).
7. The Customer will provide prompt access to all systems and resources that Snyk will need in order to complete the work.&#x20;
8. Snyk does not provide support for third-party software that is used as part of the Snyk solution, such as version control systems, repository management, trouble ticketing systems, packaging, and other software that is not part of the Snyk stack.&#x20;
9. If a Broker is required, the Customer will have all system requirements before services start.
10. All services and communications will be conducted in English.
