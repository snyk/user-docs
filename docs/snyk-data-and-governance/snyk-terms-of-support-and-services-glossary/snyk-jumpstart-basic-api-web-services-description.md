# Snyk Jumpstart Basic for API & Web Services Description

## Overview of Snyk Jumpstart Basic for API & Web

A Snyk Consultant will provide services to help the Customer accelerate its setup of Snyk API & Web through assisted account configuration (the “Jumpstart Basic for API & Web Services”). The engagement will consist of knowledge transfer, paired with configuration guidance for your team.

The objective is a working setup of Snyk API & Web, ready for the Customer team to start scanning and monitoring their targets with DAST scanning.

**Jumpstart Basic for API & Web is recommended for**:

* Teams needing assistance with the setup of Snyk API & Web

## Jumpstart Basic for API & Web services description

The Snyk Consultant will deliver the following services related to the setup of Snyk remotely as part of the Jumpstart Basic for API & Web Services to the Customer. Note that the Jumpstart Basic for API & Web Services will be delivered only for this single Snyk product.

1. [Pre-engagement planning and preparation](#pre-engagement-planning-and-preparation)
2. [Snyk API & Web configuration](#snyk-api-and-web-configuration-details)

## Pre-engagement planning and preparation

The customer should review the [prerequisites](https://docs.snyk.io/snyk-data-and-governance/snyk-terms-of-support-and-services-glossary/snyk-jumpstart-customer-prerequisites#snyk-api-and-web-prerequisites), including resources, availability, and deliverables for the Snyk API & Web module, before starting their services. The Customer acknowledges that complying with these prerequisites is its sole responsibility, and Snyk will not be responsible for any delays or failure to deliver the Jumpstart Basic for API & Web Services based on the Customer’s failure to meet these prerequisites.

## Snyk API & Web configuration details

### Delivery approach - Snyk API & Web configuration

#### Web Target configuration, including authenticated scans

The Snyk Consultant will work with the Customer to configure [Web Targets](https://help.probely.com/en/articles/3292779-how-to-set-up-target-authentication-with-a-login-form?q=Web+Target+configuration) (up to three web applications) to be scanned by Snyk API & Web. This includes configuring the authentication for each Target where necessary, such as using a login form or recorded login sequence. Snyk will also help to set up a single navigation sequence for each web application.

#### API Target configuration

The Snyk Consultant will work with the Customer to [configure API collections](https://help.probely.com/en/articles/8178059-how-to-configure-an-api-target-postman-collection) (maximum of one collection) to be scanned by Snyk API & Web, using a Postman Collection or OpenAPI definition.

#### Domain Ownership Verification

The Snyk Consultant will work with the Customer to complete [Domain Ownership Verification](https://help.probely.com/en/articles/3289281-how-to-verify-the-ownership-of-a-domain-using-a-txt-file) for one domain. This can be achieved by a .txt file, a TXT record, a CNAME record, or a meta tag.

#### Scanning Agent configuration

The Snyk Consultant will work with the Customer to configure the [Scanning Agent](https://help.probely.com/en/articles/4615595-how-to-scan-internal-applications-with-a-scanning-agent) using Docker, Docker-Compose, or Kubernetes. This is required only if there are Targets to be scanned that are not internet-accessible, and there is a maximum of one agent to be configured.

#### Target scanning in CI/CD configuration

The Snyk Consultant will work with the Customer to configure an example workflow using the API & Web CLI that allows you to trigger a Target scan when a repo is updated. This can include the opportunity of failing the pipeline if the scan finds vulnerabilities of High severity.

#### Issue ticketing integration configuration

The Snyk Consultant will work with the Customer to configure a single issue ticketing integration (e.g. Jira or Azure Boards).

#### Review of Target Scan Results

The Snyk Consultant will educate the Customer on understanding [DAST scan results](https://help.probely.com/en/articles/6843262-how-to-interpret-target-scan-results) for Web and API Targets in the Snyk API & Web UI, including the different reporting functionality that is available in the tool.

### Target initiatives - Snyk API & Web configuration

| Snyk API & Web configuration                                                           | Outcome                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Web Target configuration, including Authenticated Scans (up to three web applications) | A maximum of three web applications are configured in Snyk API & Web for scanning, including Authentication. For each application there is a maximum of one login sequence and one navigation sequence). |
| API Target configuration (up to one API collection)                                    | One API collection is configured in Snyk API & Web for scanning.                                                                                                                                         |
| Domain Ownership Verification                                                          | Domain Ownership Verification is completed for one domain.                                                                                                                                               |
| Scanning Agent Configuration (one agent)                                               | A maximum of one Scanning Agent is configured to enable Snyk API & Web to scan internal applications without internet access.                                                                            |
| Target scanning in CI/CD configuration                                                 | A single pipeline will be configured that allows you to trigger a test on an existing Web or API Target.                                                                                                 |
| Issue ticketing integration configuration                                              | A single integration will be configured that allows tickets to be created when findings of a certain severity are discovered.                                                                            |
| Review of Target Scan Results                                                          | Target Scan Results are reviewed to ensure the Customer understands how to process the DAST findings and prioritize fixing the issues.                                                                   |

## Timeline for Snyk Jumpstart Basic for API & Web delivery

Snyk Jumpstart Basic for API & Web delivery is a 30-day engagement that begins with the Jumpstart Basic for API & Web Kickoff call.

## Additional terms

The fees for this project will be a fixed price. Services will be invoiced in full at the time of purchase and are non-refundable.

Upon purchase, the Snyk implementation team will coordinate a Kickoff call at a time that works for both parties. Note that Jumpstart Basic for API & Web services expire 120 days after our consultant's first outreach, regardless of customer engagement. All Jumpstart Basic for API & Web Services automatically conclude either upon completion or eight weeks following the kick-off call, whichever occurs first.

Unless otherwise agreed to by the parties in writing, a) services must be scheduled as agreed during the pre-engagement call; b) product modules will be delivered consecutively; and c) all services will occur during normal business hours.

All services will be performed remotely. Any onsite time requires Snyk’s prior consent and will be subject to additional fees and expenses to be agreed in advance.

## Key assumptions

The following assumptions are reflected in the services outlined in this Jumpstart Basic for API & Web Services description:

1. All services will be performed remotely using video conferencing software such as Google Meet.
2. The Customer must provide prompt feedback on all deliverables.
3. The Customer will appoint one subject matter expert who will be the point of contact for the Jumpstart Basic for API & Web Services. This subject matter expert must be available to work remotely with the Snyk Consultant for the entirety of the engagement.
4. The Customer will provide Snyk with documentation and access to subject matter experts for non-Snyk systems and software if required within the scope of the engagement.
5. The Customer will have identified key personnel prior to the beginning of the engagement.
6. Services will be scheduled and delivered during Snyk’s normal business hours: 8 am to 5 pm local time zone Monday through Friday (Sunday through Thursday where applicable based on the region of the assigned Snyk Consultant).
7. The Customer will provide prompt access to all systems and resources that Snyk will need in order to complete the work.
8. Snyk does not provide support for third-party software that is used as part of the Snyk solution, such as version control systems, repository management, trouble ticketing systems, packaging, and other software that is not part of the Snyk stack.
9. If a Broker is required, the Customer will have all system requirements before services start.
10. All services and communications will be conducted in English.