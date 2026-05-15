# Snyk Penetration Testing service description

## Overview

The Snyk Penetration Testing service delivers an expert-driven, end-to-end security evaluation of your web applications. We go beyond automated scans to identify real-world vulnerabilities and potential business impact. Our methodology encompasses scoping, threat modeling, design review, custom test plan development, rigorous hands-on testing, and continuous reporting, culminating in clear remediation advice and a validation retest. Our testing is primarily focused on identifying the OWASP Top 10 list of most common web application vulnerabilities.

For customers with AI and Large Language Model-powered applications, **AI/LLM Penetration Testing is available as an optional, separately scoped subset of this service**. This specialized testing assesses model behavior, LLM APIs and orchestration, integration layers (like function calling and RAG), data privacy controls, and the surrounding SDLC/infrastructure. AI/LLM Penetration Testing is focused on the OWASP LLM Top 10.

### Key benefits

* **Holistic Web Application Security Assessment**: Evaluates your web applications to uncover risks beyond traditional automated testing.
* **Threat-Led Approach**: Focuses on real-world attack scenarios and potential business impact.
* **Actionable Remediation Guidance**: Provides clear, prioritized fixes mapped to relevant security standards.
* **Reproducible Evidence**: Backs findings with concrete proof-of-concept exploits and supporting data.
* **Comprehensive Coverage**: Assesses against the OWASP Top 10 vulnerabilities and other common web application risks.
* **Validation of Remediation**: Includes a no-cost retest to confirm the effectiveness of implemented fixes.
* **Clear Reporting for All Stakeholders**: Delivers executive summaries and detailed technical reports.

### Web application penetration testing activities

Our comprehensive methodology ensures a thorough evaluation of your web application:

* **Scoping & Threat Modeling**: Collaborative meeting to define the scope of the engagement and develop a threat model specific to your web application implementation.
* **Design Review & Test Planning**: Analysis of your web application architecture and the creation of a custom test plan tailored to your specific systems and identified threats, with a focus on the OWASP Top 10.
* **Fieldwork**: Rigorous hands-on testing across key areas, including but not limited to:
  * **Input Validation & Sanitization**: Identifying vulnerabilities like Injection (SQL, Command, LDAP), Cross-Site Scripting (XSS).
  * **Broken Access Control**: Assessing for vertical and horizontal privilege escalation, insecure direct object references.
  * **Security Misconfiguration**: Discovering insecure defaults, unpatched flaws, open cloud storage, misconfigured HTTP headers.
  * **Insecure Design**: Evaluating for logical flaws, lack of business logic defenses.
  * **Vulnerable and Outdated Components**: Identifying known vulnerabilities in libraries, frameworks, and other components.
  * **Identification and Authentication Failures**: Assessing for weak authentication mechanisms, session management flaws.
  * **Software and Data Integrity Failure**s: Examining for insecure deserialization, insecure updates.
  * **Server-Side Request Forgery (SSRF)**: Detecting vulnerabilities allowing server-side requests to unauthorized locations.
* **Fieldwork for Penetration Testing focused on AI/LLM implementations** includes testing of five key layers:
  * **Model Behavior**: Assessing for jailbreaks, data leakage, guardrail bypass, extraction, and model-level vulnerabilities.
  * **LLM APIs/Orchestration**: Identifying traditional AppSec issues as well as vulnerabilities specific to agentic and LLM API interactions.
  * **Integration Layers**: Evaluating the security of function calling and Retrieval-Augmented Generation (RAG) pipelines.
  * **Data Privacy Controls**: Examining measures protecting user and training data.
  * **SDLC/Infrastructure**: Assessing the security of the surrounding development lifecycle and infrastructure (CI/CD, servers, cloud).
* **Continuous Reporting**: Regular updates on findings throughout the testing process.
* **Executive & Technical Readouts**: Presentation of findings with a focus on business impact for leadership and detailed technical information for security teams. Findings are mapped to standards like OWASP Top 10, OWASP API Top 10, OWASP LLM Top 10 and others as relevant.
* **Remediation Support & Retest**: Guidance on addressing identified vulnerabilities, followed by a no-cost retest to validate the effectiveness of implemented fixes and a formal attestation of the testing.

### Key deliverables

Upon completion of the engagement, you will receive:

* **Executive Summary**: A high-level overview of the security posture and key findings for leadership.
* **Detailed Technical Report**: In-depth documentation of all identified vulnerabilities, their impact, and remediation recommendations.
* **Evidence Bundle**: Supporting materials, including logs and proof-of-concept exploits.
* **Attestation Letter**: Formal confirmation of the penetration testing engagement and the validation of remediated findings (post-retest).

### Engagement timeline

Typical engagement timelines for single web applications are as follows (timelines scale with system complexity):

* ½ day kickoff
* ½ day of design review and planning
* 4 or 9 days of testing with rolling reporting
* ½ day final readout
* 1 day for retest and attestation

### Key assumptions

To ensure a successful and effective penetration testing engagement, we require the following from the customer:

* **Access to Test Environment**: Provision of a stable and representative test environment for the applications and related infrastructure.
* **Test Credentials and Access**: Provision of necessary credentials and access permissions to interact with all in-scope application components, APIs, user interfaces, and relevant data stores.
* **Detailed System Information**: Provision of relevant documentation and information about the application or LLM architecture, data flows, APIs, training data sources (if applicable and in scope), and any existing security controls.
* **Clear Communication Channels**: Establishment of clear and responsive communication channels with designated technical contacts.
* **Availability of Technical Personnel**: Availability of knowledgeable technical personnel to answer questions and assist with troubleshooting.
* **Defined Scope Confirmation**: Formal confirmation and agreement on the testing scope.
* **Awareness of Testing Activities**: Ensuring relevant internal teams are aware of the scheduled testing.
* **Prompt Response to Findings**: Commitment to reviewing findings and engaging in remediation discussions.

### Additional terms

The customer acknowledges that the services referred to herein may be performed either by Snyk personnel or by a Snyk certified partner under the direction and supervision of Snyk. In either event, Snyk remains fully responsible and liable for the performance of these services and its partner's compliance with the terms of the underlying Agreement between Snyk and the customer.
