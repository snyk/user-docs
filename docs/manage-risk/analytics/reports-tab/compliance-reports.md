# Compliance reports

The Compliance reports section includes the following reports:

* [CWE Top 10 KEV](compliance-reports.md#cwe-top-10-kev-report) report
* [CWE Top 25](compliance-reports.md#cwe-top-25-report) report
* [OWASP Top 10](compliance-reports.md#owasp-top-10-report) report
* [PCI-DSS v4.0.1 ](compliance-reports.md#pci-dss-v4.0.1-report)report

## CWE Top 10 KEV report

The [CWE Top 10 KEV Weaknesses](https://cwe.mitre.org/top25/archive/2023/2023_kev_list.html) list identifies the top ten CWEs in the Cybersecurity and Infrastructure Security Agency’s (CISA) [Known Exploited Vulnerabilities](https://www.cisa.gov/known-exploited-vulnerabilities-catalog) (KEV) Catalog, a database of security flaws in software applications and weaknesses that have been exposed and leveraged by attackers.

The report is based on the version released in 2023 by Mitre. The supported products are Snyk Open Source, Snyk Container, and Snyk Code.

## CWE Top 25 report

The [CWE Top 25](https://cwe.mitre.org/top25/) Most Dangerous Software Weaknesses is a list that demonstrates the current most common and impactful software weaknesses based on Common Vulnerabilities and Exposures (CVEs) severity and their exploitation potential.

The report is based on the latest version released in 2023 by Mitre. The supported products are Snyk Open Source, Snyk Container, and Snyk Code.

## OWASP Top 10 report

The [OWASP Top 10](https://owasp.org/www-project-top-ten/) is a standard awareness document for developers and web application security. It represents a broad consensus about the most critical security risks for web applications and is globally recognized by developers as the first step towards more secure coding.

Each control in the list (A1, A2, and so on) is based on a list of Common Weakness Enumerations (CWEs). For example, [A01:2021 – Broken Access Control](https://owasp.org/Top10/A01_2021-Broken_Access_Control/) is based on a list of 34 CWEs.

The CWEs are mapped to Snyk-IDs (), which are mapped to issues.

For example, the critical vulnerability [SNYK-JAVA-ORGAPACHELOGGINGLOG4J-2314720](https://security.snyk.io/vuln/SNYK-JAVA-ORGAPACHELOGGINGLOG4J-2314720) is classified as [CWE-94](https://cwe.mitre.org/data/definitions/94.html), which is part of the OWASP TOP 10 [A03:2021 - Injection](https://owasp.org/Top10/A03_2021-Injection/). All the issues related to this vulnerability will be under the A03 category.

Learn more by using the [OWASP TOP 10 Learning path](https://learn.snyk.io/learning-paths/owasp-top-10/) on Snyk Learn.

The report is based on the latest mapping released in 2021. The supported products are Snyk Open Source, Snyk Container, and Snyk Code.

## PCI-DSS v4.0.1 report

{% hint style="info" %}
**Release status**

The PCI-DSS v4.0.1 report is in Early Access and available only with Enterprise plans.
{% endhint %}

PCI Security Standards are technical and operational requirements created by the PCI Security Standards Council (PCI SSC) to safeguard cardholder data. These standards apply to all entities that store, process, or transmit this information and include requirements for software developers and manufacturers.\
\
The Council manages these standards, while compliance is enforced by founding members: American Express, Discover Financial Services, JCB, MasterCard, and Visa Inc.

Snyk PCI-DSS v4.0.1 Report is designed to help you:

* Estimate readiness for meeting the PCI-DSS AppSec requirements for SCA and SAST based on the Snyk scan results.
* Provide evidence that the Organization is meeting the PCI-DSS AppSec requirements for SCA and SAST vulnerabilities.
* Prioritize issues to improve PCI-DSS compliance readiness.

The report identifies PCI-DSS risks and violations based on the following PCI-DSS v4.0.1 requirements:

1. **Requirement 6.2.4:** Engineers use various techniques to prevent or mitigate common software attacks and related vulnerabilities in bespoke and custom software. This includes but is not limited to the following methods:
   * Injection attacks, including SQL, LDAP, XPath, or other command, parameter, object, fault, or injection-type flaws.
   * Attacks on data and data structures, including attempts to manipulate buffers, pointers, input data, or shared data.
   * Attacks on cryptography usage, including attempts to exploit weak, insecure, or inappropriate cryptographic implementations, algorithms, cipher suites, or modes of operation.
   * Attacks on business logic, including attempts to abuse or bypass application features and functionalities through the manipulation of APIs, communication protocols and channels, client-side functionality, or other system or application functions and resources. This includes cross-site scripting (XSS) and cross-site request forgery (CSRF).
   * Attacks on access control mechanisms, including attempts to bypass or abuse identification, authentication, or authorization mechanisms or attempts to exploit weaknesses in the implementation of such mechanisms.
   * Attacks using any “high-risk” vulnerabilities identified in the vulnerability identification process, as defined in Requirement 6.3.1.
2. **Requirement 6.3.3:** All system components are protected from known vulnerabilities by installing applicable security patches and updates as follows:
   * Patches and updates for critical vulnerabilities, identified according to the risk ranking process at Requirement 6.3.1 are installed within one month of release.

### Snyk Violation Analysis based on PCI-DSS attack categories

As the standard does not explicitly define specific CWEs or CVEs, Snyk provides an analysis based on leading CWEs associated with the named attack categories. Below are the CWEs categorized by attack type:

#### Injection Attack Violations Summary

The following list provides an association between the identified attack categories and the CWEs associated with each category:

* SQL Injection: CWE-89
* LDAP Injection: CWE-90
* XML Injection (XPath Injection): CWE-91
* Command Injection: CWE-77
* Use of Unsafe Reflection: CWE-470

#### Attacks on Data and Data Structures Violations Summary

The following list provides an association between the identified attack categories and the CWEs associated with each category:

* Buffer Overflow: CWE-120
* NULL Pointer Dereference: CWE-476
* Double Free: CWE-415
* Concurrent Execution using Shared Resource with Improper Synchronization (‘Race Condition’): CWE-362

#### Attacks on Cryptography Usage Violations Summary

The following list provides an association between the identified attack categories and the CWEs associated with each category:

* Use of a Broken or Risky Cryptographic Algorithm: CWE-327
* Use of Insufficiently Random Values: CWE-330
* Improper Verification of Cryptographic Signature: CWE-347
* Cleartext Transmission of Sensitive Information: CWE-319
* Use of Hard-coded Cryptographic Key: CWE-321

#### Attacks on Business Logic Violations Summary

The following list provides an association between the identified attack categories and the CWEs associated with each category:

* Server-Side Request Forgery (SSRF): CWE-918
* Cross-Site Request Forgery (CSRF): CWE-352
* Cross-Site Scripting (XSS): CWE-79
* Origin Validation Error: CWE-346
* Improper Authorization: CWE-285
* Exposure of Sensitive Information to an Unauthorized Actor: CWE-200

#### Attacks on Access Control Mechanisms Violations Summary

The following list provides an association between the identified attack categories and the CWEs associated with each category:

* Improper Authentication: CWE-287
* Improper Access Control: CWE-284
* Incorrect Authorization: CWE-863
* Authorization Bypass Through User-Controlled Key: CWE-639
* Missing Authentication for Critical Function: CWE-306
* Incorrect Implementation of Authentication Algorithm: CWE-303

#### Attacks on Access Control Mechanisms Violations Summary

The Missing Authorization attack category is associated with CWE-862.

### PCI-DSS v4.0.1 Guidance

The report is filtered by default on open issues of critical severity. Those filters are also applicable when exporting the report to PDF.

#### PCI-DSS Readiness Trend

The PCI-DSS Readiness Trend is designed to help you track your progress toward eliminating PCI-DSS violations. A violation is defined as a critical vulnerability elected by the PCI-DSS attack categories (as explained in Requirement 6.2.4) that is more than 30 days old, as stated in Requirement 6.3.3.

The number on the left indicates the violation status and the progress made in the last seven days.

The trend shows all vulnerabilities per Requirement 6.2.4, categorized by age bucket. This allows for quick identification of potential violations and vulnerabilities that may soon become violations.

#### Attack category breakdown

The breakdown table helps identify the number of vulnerabilities by attack category (as per requirement 6.2.4) or by Snyk Organization based on the relevant age bucket.

Use the table to pinpoint major attack categories or Snyk Organizations that lead to PCI-DSS violations. You can click on the figures to explore the specific issues in more detail.

{% hint style="info" %}
After you investigate and see the actual issues behind the figures, you may proceed by:

* Vulnerability triage and prioritization.
* Conclude the prevalent CWEs and CVEs by sorting on the CWE/CVE column and filtering those CWEs/CVEs in the [Vulnerabilities Detail Report](compliance-reports.md#vulnerabilities-detail-report) to surface all the vulnerability occurrences across targets and Projects.
* Run a vulnerability eradication campaign or assign Snyk Learn training to relevant engineering teams.
{% endhint %}
