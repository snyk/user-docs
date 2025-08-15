# Snyk Jumpstart Services Description

## Overview of Snyk Jumpstart

A Snyk Consultant will provide services to help the  Customer accelerate its setup of Snyk products through assisted account configuration (the “Jumpstart Services”). The engagement will consist of knowledge transfer, paired with configuration guidance for your team.&#x20;

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
7. [Snyk Essentials configuration](snyk-jumpstart-services-description.md#snyk-essentials-configuration)
8. [Snyk AppRisk configuration](snyk-jumpstart-services-description.md#snyk-apprisk-configuration)
9. [Snyk API & Web configuration](snyk-jumpstart-services-description.md#snyk-api-and-web-configuration-details)

## Pre-engagement planning and preparation

A pre-engagement cal**l** will be held before the start of the Jumpstart Services listed herein, to ensure the Customer understands the prerequisites required for engagement start along with the deliverables completed during the engagement timeframe. Customer [prerequisites](snyk-jumpstart-customer-prerequisites.md), including resources, availability, and deliverables for each product module, will be reviewed and confirmed. The Customer acknowledges that complying with these prerequisites is its sole responsibility, and Snyk will not be responsible for any delays or failure to deliver the Jumpstart Services based on the Customer’s failure to meet these prerequisites.

## Snyk Platform configuration

### Delivery approach - Snyk Platform configuration

The Snyk delivery method is designed to ensure rapid value realization with Snyk. Throughout our collaboration, Snyk will guide the Customer in setting up a foundational configuration and equip them to expand this setup to other applications and integrations. Ensuring Snyk is correctly set up from the outset improves developer adoption and paves the way for long-term success.

#### Base SSO configuration and custom mapping

The Snyk Consultant will work with the Customer to configure SSO through SAML, Entra ID (formerly Azure AD), OIDC, or ADFS connection along with custom mapping of dynamically assigned users to Snyk Groups and Organizations based on data provided by the desired Identity Provider (IdP) to set up a scaled user provisioning and access model.

#### **Template Organization configuration**

The Snyk Consultant will work with the Customer to configure notification settings, language settings, and Snyk account structure. Time will be spent reviewing configuration details with the Customer to ensure the Customer understands and maintains them.

#### Admin training

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

#### Developer training

The Snyk Consultant will run a 60-minute training session for Customer’s Developers to ensure all collaborator users know how to access Snyk results and initiate tests. This training will cover a range of tools, including the UI (app.snyk.io), IDE plugin, and CLI. Developer Training will be completed after the product-specific modules of Jumpstart have been completed.

### Target initiatives - Snyk Platform configuration

<table><thead><tr><th width="176">Platform - initial setup</th><th width="555">Outcome</th></tr></thead><tbody><tr><td>Base SSO connection configuration</td><td>Users can access Snyk based on role.</td></tr><tr><td>SSO Custom Mapping</td><td>Users can access Snyk based on role.</td></tr><tr><td>(Template Org) Custom Roles and service account configuration</td><td>Provide a templated organization to replicate and scale your setup of Snyk quickly.</td></tr><tr><td>(Template Org) Notification configuration</td><td>Provide a templated organization to replicate and scale your setup of Snyk quickly.</td></tr><tr><td>(Template Org) Language settings configuration</td><td></td></tr><tr><td>(Template Org) Jira or Slack App configuration</td><td></td></tr><tr><td>(Template Org) Account Organization and Group configuration</td><td></td></tr><tr><td>Admin training (1 session)</td><td>Group and Organisation administrators are comfortable using Snyk to import their targets and complete the initial configuration for each team.</td></tr><tr><td>Developer training (1 session)</td><td>Delivered after product-specific modules are completed, ensuring a successful launch and initial adoption with Developers.</td></tr></tbody></table>

## Snyk Open Source configuration

### Delivery approach - Snyk Open Source configuration

The Snyk delivery method is designed to ensure rapid value realization with Snyk. Throughout our collaboration, Snyk will guide the Customer in setting up a foundational configuration and equip the Customer to expand this setup to other applications and integrations. Ensuring Snyk is correctly set up from the outset improves developer adoption and paves the way for long-term success.

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) either through the UI import functionality or through the [API Import tool](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/).

#### SCM integration settings

The Snyk Consultant will work with the Customer to configure SCM integration settings based on the Customer’s desired gating strategy.

#### SCM Broker installation

The Snyk Consultant will work with the Customer to install the Snyk Broker in a pre-determined environment that follows the [Snyk Broker system requirements](../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/).

#### Snyk API Import and SCM sync

The Consultant will review the Snyk API Import script to ensure the Customer understands how to import additional Projects into Snyk and keep their SCM integration in sync with incoming changes to manifests.

#### Single pipeline configuration (direct Integration or CLI)

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk test` and `snyk monitor` commands to provide the Customer with an understanding of how to configure additional pipeline scans.

#### SBOM walkthrough (API and CLI)

The Snyk Consultant will educate the Customer on creating an SBOM through the Snyk API and the Snyk CLI.

#### Interpreting and actioning Open Source results

The Snyk Consultant will educate the Customer on understanding Snyk Open Source results through the CLI and Snyk UI and how to manage Snyk Open Source results using Snyk Reporting.

#### Documentation close-out

The Customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, the document offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document puts forth practical recommendations and actionable next steps that will aid the Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document provides an essential guide for Customers to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk Open Source configuration

<table><thead><tr><th width="300">Snyk Open Source configuration</th><th width="306">Outcome</th></tr></thead><tbody><tr><td>Repository import (one SCM integration)</td><td>Import a maximum of 50 targets into Snyk using a <a href="../../developer-tools/scm-integrations/organization-level-integrations/">supported SCM integration</a> (GitHub, Azure Repos, Bitbucket, GitLab).</td></tr><tr><td>SCM integration settings</td><td>Configure SCM integration settings to the Customer’s desired gating settings.</td></tr><tr><td>SCM Broker installation</td><td>Install SCM Broker in a pre-determined customer environment based on Snyk system requirements.</td></tr><tr><td>Snyk Tools - API Import and SCM Sync</td><td>Gain an understanding of how to use the Snyk API Import script to import additional targets and keep repos in Sync (GHE only).</td></tr><tr><td>Single pipeline configuration (direct integration OR CLI)</td><td>Configure a pipeline to run <code>snyk test</code> and <code>snyk monitor</code>.</td></tr><tr><td>SBOM Walkthrough (CLI and API)</td><td>Gain an understanding of generating an SBOM through Snyk using the CLI and API.</td></tr><tr><td>Interpreting and actioning Open Source results</td><td>Gain an understanding of how to view Open Source results in Snyk Reporting along with managing issues.</td></tr><tr><td>Documentation close-out</td><td>Gain an understanding of work completed along with a runbook for onboarding additional projects.</td></tr></tbody></table>

## Snyk Code configuration

### Delivery approach - Snyk Code configuration

The Snyk delivery method is tailored to ensure rapid value realization with Snyk. Throughout our collaboration, Snyk will guide the Customer in setting up a foundational configuration and equip the Customer to expand this setup to other applications and integrations. Ensuring Snyk is correctly set up from the outset improves developer adoption and paves the way for long-term success.

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) either through the UI import functionality or through the [API Import tool](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/).&#x20;

#### SCM integration settings

The Snyk Consultant will work with the Customer to configure SCM integration settings based on the Customer’s desired gating strategy.

#### SCM Broker installation

The Snyk Consultant will work with the Customer to install the Snyk Broker in a pre-determined environment that follows the [Snyk Broker system requirements](../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/).

#### Snyk API Import and SCM Sync

The Consultant will review the Snyk API Import script to ensure the Customer understands how to import additional Projects into Snyk and keep their SCM integration in sync with incoming changes to manifests.

#### Interpreting and actioning Code results

The Snyk Consultant will educate the Customer on understanding Snyk Code results through the CLI and Snyk UI and how to manage Snyk Code results using Snyk Reporting.

#### Documentation close-out

The Customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, the document offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document puts forth practical recommendations and actionable next steps that will aid the Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document provides an essential guide for customers to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk Code configuration

<table><thead><tr><th width="250">Snyk Code Configuration</th><th>Outcome</th></tr></thead><tbody><tr><td>Repository import (one SCM integration)</td><td>Import a maximum of 50 targets into Snyk using a <a href="../../developer-tools/scm-integrations/organization-level-integrations/">supported SCM Integration</a>.</td></tr><tr><td>SCM integration settings</td><td>Configure SCM integration settings to the Customer’s desired gating settings.</td></tr><tr><td>SCM Broker installation</td><td>Install SCM Broker in a pre-determined customer environment based on Snyk system requirements.</td></tr><tr><td>Snyk Tools - API Import and SCM Sync</td><td>Gain an understanding of how to use the Snyk API Import script to import additional targets and keep their repos in sync (GHE only).</td></tr><tr><td>Interpreting and actioning Code results</td><td>Gain an understanding of how to view Code results in Snyk Reporting along with managing issues.</td></tr><tr><td>Documentation close-out</td><td>Gain an understanding of work completed along with a runbook for onboarding additional projects.</td></tr></tbody></table>

## Snyk Container configuration

### Delivery approach - Snyk Container configuration

The Snyk delivery method is tailored to ensure rapid value realization with Snyk. Throughout our collaboration, Snyk will guide the Customer in setting up a foundational configuration and equip them to expand this setup to other applications and integrations. Ensuring Snyk is correctly set up from the outset improves developer adoption and paves the way for long-term success.

#### Single Broker Container registry installation and configuration

The Snyk Consultant will work with the Customer to configure and install Snyk Broker if needed for a single [supported Container Registry](https://snyk.io/integrations/?type=container-registries) integration.

#### Container registry import

The Snyk Consultant will work with the Customer to import their container images into Snyk (up to 50 targets) through the UI import functionality.

#### Interpreting and actioning Snyk Container results

The Snyk Consultant will educate the Customer on understanding Snyk Container results through the CLI and Snyk UI and how to manage Snyk Container results using Snyk Reporting.

#### Single CI/CD CLI configuration

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk container test` and `snyk container monitor` commands to provide the Customer with an understanding of how to configure additional pipeline scans.

#### Custom Base Images walkthrough (UI and CLI)

The Snyk Consultant will educate the Customer on how to use the Snyk Custom Base Image Recommendation functionality both in the Snyk UI and CLI.

#### Documentation close-out

The Customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, the document offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document puts forth practical recommendations and actionable next steps that will aid Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document provides an essential guide for customers to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk Container configuration

<table><thead><tr><th width="269">Snyk Container Configuration</th><th>Outcome</th></tr></thead><tbody><tr><td>Single Broker Container Registry installation and configuration</td><td>Single Broker installed and configured for a <a href="https://snyk.io/integrations/?type=container-registries">Supported Container Registry</a>.</td></tr><tr><td>Container Registry import (up to 50 targets)</td><td>Import a maximum of 50 targets into Snyk using a <a href="https://snyk.io/integrations/?type=container-registries">Supported Container Registry</a>.</td></tr><tr><td>Interpreting and actioning Container results</td><td>Gain an understanding of how to view Container results in Snyk Reporting along with managing issues.</td></tr><tr><td>Single CI/CD CLI rntegration</td><td>Configure a single pipeline to <code>test</code> and <code>monitor</code> for Snyk Container.</td></tr><tr><td>Custom Base Images walkthrough (UI and CLI)</td><td>Gain an understanding of how to use the Custom Base Image Recommendations functionality through the UI and CLI.</td></tr><tr><td>Documentation close-out</td><td>Gain an understanding of work completed along with a runbook for onboarding additional projects.</td></tr></tbody></table>

## Snyk IaC Configuration

### Delivery approach - Snyk IaC configuration

#### Repository import

The Snyk Consultant will work with the Customer to import their repositories into Snyk (up to 50 targets) using the UI import functionality to import into the Customer’s SCM integration.

#### Interpreting and actioning IaC results

The Snyk Consultant will educate the Customer on understanding Snyk IaC results through the CLI and Snyk UI and how to manage Snyk IaC results using Snyk Reporting.

#### SCM integration settings

The Snyk Consultant will work with the Customer to configure SCM Integration settings based on the Customer’s desired gating strategy.

#### SCM Broker installation

The Snyk Consultant will work with the Customer to install the Snyk Broker in a predetermined environment that follows the [Snyk Broker system requirements](../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/).

#### Single pipeline CI/CD CLI configuration

The Snyk Consultant will work with the Customer to configure a single pipeline to run the `snyk iac test` and `snyk iac test --report` commands to provide the Customer with an understanding of how to configure additional pipeline scans and fix misconfigurations.

#### Documentation close-out

The customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, the document offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document puts forth practical recommendations and actionable next steps that will aid the Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document provides an essential guide for customers to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk IaC configuration

| Snyk IaC configuration                  | Outcome                                                                                                             |
| --------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| Repository import (one SCM integration) | Import a maximum of 50 targets into Snyk via a [supported SCM Integration](https://snyk.io/integrations/?type=scm). |
| Interpreting and actioning IaC Results  | Gain an understanding of how to view IaC results in Snyk Reporting along with managing misconfigurations.           |
| SCM integration settings                | Configure SCM integration settings to the Customer’s desired gating settings.                                       |
| SCM Broker installation                 | Install SCM Broker in a pre-determined customer environment based on Snyk System Requirements.                      |
| Single Pipeline CI/CD CLI configuration | Configure a single pipeline to test and provide a report for Snyk IaC.                                              |
| Documentation close-out                 | Gain an understanding of work completed along with a runbook for onboarding additional projects.                    |

## Snyk Essentials configuration

This portion of the Jumpstart service is part of the Platform Configuration, and prepares Customers to better operationalize and scale usage of Snyk with broad application visibility and security coverage management.

### Delivery approach - Snyk Essentials configuration

#### Coverage and visibility configuration

The Snyk Consultant will work with Customer to configure Snyk Essentials as follows:

* Configure one Source Code Management (SCM) integration, and if necessary, Snyk Broker for this integration
* Configure one application context provider integration
* Configure asset policies. The Snyk Consultant will guide the Customer through the setup and configuration of the following asset policies:
  * Classify assets visible to Snyk
  * Identify coverage gaps based on purchased Snyk products
  * Notify by email (or similar) when a criteria (for example, coverage gap) is met

#### **Walkthrough of coverage and visibility use cases in Snyk Essentials**

The Snyk Consultant will educate the Customer on how to identify assets not being scanned by Snyk and how to group assets and issues based on asset classification.

#### **Walkthrough of prioritized issues in Snyk Essentials dashboard**

The Snyk Consultant will educate the Customer on how to filter and prioritize issues in the Asset Dashboard.

#### Documentation close-out

The Customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, it offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document puts forth practical recommendations and actionable next steps that will aid Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document is an essential guide for customers to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk Essentials configuration

<table><thead><tr><th width="382">Snyk Essentials Prioritization Configuration</th><th width="299">Outcome</th></tr></thead><tbody><tr><td>Coverage and visibility configuration</td><td>SCM integration is configured in Snyk Essentials and two starter policies are created to show coverage gaps and asset classifications respectively.</td></tr><tr><td>Walkthrough of coverage and visibility use cases in Snyk Essentials</td><td>Gain an understanding of how to identify assets that are not currently being scanned by one more Snyk controls, as well as how to group assets and issues based on asset classification.</td></tr><tr><td>Walkthrough of prioritized issues in Snyk Essentials</td><td>Gain an understanding of how to filter and prioritize issues in the Asset Dashboard.</td></tr><tr><td>Documentation close-out</td><td>Gain an understanding of work completed.</td></tr></tbody></table>

## Snyk AppRisk configuration

### Delivery approach - Snyk AppRisk configuration

**Third-party coverage and visibility configuration**

The Snyk Consultant will work with the Customer to configure Snyk AppRisk as follows:

* Configure one third-party integration (for example, Secrets)
* Configure an asset policy leveraging newly configured integration

**Snyk Runtime Sensor installation**&#x20;

The Snyk Consultant will work with the Customer to install the Snyk Runtime Sensor in a predetermined environment that follows the [Snyk Runtime Sensor system requirements](../../integrations/snyk-runtime-sensor.md). The Snyk Runtime Sensor allows the Customer to prioritize issues based on risk factors of **Deployed** and **Loaded** **Package**.

#### Tagging of Projects for the issues prioritization feature (up to 50 Targets)

The Snyk consultant will provide guidance on how the Customer can add Project tags to Targets so Open Source, Code, and Container Projects are linked for the Snyk AppRisk Issues prioritization feature. This will include using a script for tagging SCM-monitored Projects (Open Source and Code) and the CLI for tagging Container Projects.

**Walkthrough of prioritized issues in AppRisk**

The Snyk Consultant will educate the Customer on how to filter and prioritize issues in the Asset Dashboard, using the new risk factors detected by the Snyk Runtime Sensor.

**Documentation close-out**&#x20;

The Customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, the document offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document sets out practical recommendations and actionable next steps that will aid the Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document is an essential guide for the Customer to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk AppRisk configuration

| Snyk AppRisk Configuration                                                   | Outcome                                                                                                                            |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| Third-party coverage and visibility configuration                            | Third-Party integration is configured in AppRisk and a policy is created to show how the new data can be processed by AppRisk.     |
| <p>Snyk Runtime Sensor installation</p><p><br></p>                           | Install Snyk Runtime Sensor on a predetermined customer environment based on Snyk System Requirements.                             |
| Tagging of Projects for the issues prioritization feature (up to 50 targets) | Targets imported through an SCM integration must be tagged with tags that match the CLI-imported Container Projects.               |
| Walkthrough of prioritized issues in AppRisk                                 | Gain an understanding of how to filter and prioritize issues in the Asset Dashboard using deployed and loaded package risk factors |
| Documentation close-out                                                      | Gain an understanding of work completed                                                                                            |

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

#### **Review of Target Scan Results**

The Snyk Consultant will educate the Customer on understanding a [DAST scan](https://help.probely.com/en/articles/6843262-how-to-interpret-target-scan-results) for Web and API Targets in the Snyk API & Web UI, including the different reporting functionality that is available in the tool.

#### Documentation close-out

The Customer will be provided with a document that provides a comprehensive overview of the professional services rendered by Snyk during the engagement. Spanning the period from the engagement's start to its conclusion, the document offers insights into account configuration, repository onboarding, and integrations. More than just a retrospective, the document puts forth practical recommendations and actionable next steps that will aid the Customer in optimizing their use of Snyk for improved application security. By detailing both the accomplishments and the roadmap ahead, this document provides an essential guide for customers to realize the full potential of their investment in Snyk.

### Target initiatives - Snyk API & Web configuration

| Snyk API & Web configuration                                                           | Outcome                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Web Target configuration, including Authenticated Scans (up to three web applications) | A maximum of three web applications are configured in Snyk API & Web for scanning, including Authentication. For each application there is a maximum of one login sequence and one navigation sequence). |
| API Target configuration (up to one API collection)                                    | One API collection is configured in Snyk API & Web for scanning.                                                                                                                                         |
| Domain Ownership Verification                                                          | Domain Ownership Verification is completed for one domain.                                                                                                                                               |
| Scanning Agent Configuration (one agent)                                               | A maximum of one Scanning Agent is configured to enable Snyk API & Web to scan internal applications without internet access.                                                                            |
| Review of Target Scan Results                                                          | Target Scan Results are reviewed to ensure the Customer understands how to process the DAST findings and prioritize fixing the issues.                                                                   |
| Documentation close-out                                                                | Gain an understanding of work completed.                                                                                                                                                                 |

## Timeline for Snyk Jumpstart delivery

Snyk Jumpstart delivery is an eight-week engagement that begins with the Pre-engagement call.&#x20;

This engagement will include initial Platform Configuration and each product module that has been purchased. Modules have been designed to be delivered consecutively during the Product Deployment stage, as shown below:

<figure><img src="../../.gitbook/assets/Screenshot 2025-02-11 at 3.52.59 PM.png" alt=""><figcaption><p>Snyk Jumpstart timeline</p></figcaption></figure>

## Additional terms

The fees for this project will be a fixed price. Services will be invoiced in full at the time of purchase and are non-refundable.&#x20;

Upon purchase, the Snyk implementation team will coordinate a pre-engagement call at a time that works for both parties. Note that Jumpstart services expire 120 days after our consultant's first outreach, regardless of customer engagement. All Jumpstart Services automatically conclude either upon completion or eight weeks following the kick-off call, whichever occurs first.

Unless otherwise agreed to by the parties in writing, a) services must be scheduled as agreed during the pre-engagement call; b) product modules will be delivered consecutively; and c) all services will occur during normal business hours.&#x20;

All services will be performed remotely. Any onsite time requires Snyk’s prior consent and will be subject to additional fees and expenses to be agreed in advance.&#x20;

## Key assumptions

The following assumptions are reflected in the services outlined in this Jumpstart Services description:

1. All services will be performed remotely using video conferencing software such as Zoom.&#x20;
2. The Customer must provide prompt feedback on all deliverables.
3. &#x20;The Customer will appoint one subject matter expert who will be the point of contact for the Jumpstart Services. This subject matter expert must be available to work remotely with the Snyk Consultant for the entirety of the engagement.
4. The Customer will provide Snyk with documentation and access to subject matter experts for non-Snyk systems and software if required within the scope of the engagement.
5. The Customer will have identified key personnel prior to the beginning of the engagement.
6. Services will be scheduled and delivered during Snyk’s normal business hours: 8 am to 5 pm local time zone Monday through Friday (Sunday through Thursday where applicable based on region of the assigned Snyk Consultant).
7. The Customer will provide prompt access to all systems and resources that Snyk will need in order to complete the work.&#x20;
8. Snyk does not provide support for third-party software that is used as part of the Snyk solution, such as version control systems, repository management, trouble ticketing systems, packaging, and other software that is not part of the Snyk stack.&#x20;
9. If a Broker is required, the Customer will have all system requirements before services start.
10. All services and communications will be conducted in English.
