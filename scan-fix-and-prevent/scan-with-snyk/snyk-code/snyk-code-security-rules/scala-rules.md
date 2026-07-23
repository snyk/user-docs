---
description: Snyk Code security rules for Scala
---

# Scala rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/2025/)(2025 edition) category to which the rule belongs to, if any, and if it is included in [CWE Top 25](https://cwe.mitre.org/top25/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories            | Autofixable |
| ------------------------------------------------------------ | ---------------- | ------------------------------ | ----------- |
| Use of Potentially Dangerous Function                        | CWE-676          | OWASP:A06:2025                 | Yes         |
| Cleartext Storage of Sensitive Information in a Cookie       | CWE-315          | OWASP:A02:2025                 | Yes         |
| Code Injection                                               | CWE-94           | CWE Top 25, OWASP:A05:2025     | Yes         |
| Command Injection                                            | CWE-78           | CWE Top 25, OWASP:A05:2025     | Yes         |
| Deserialization of Untrusted Data                            | CWE-502          | CWE Top 25, OWASP:A08:2025     | Yes         |
| Cross-Site Request Forgery (CSRF)                            | CWE-352          | CWE Top 25, OWASP:A01:2025     | Yes         |
| Information Exposure                                         | CWE-200          | CWE Top 25, OWASP:A01:2025     | Yes         |
| Cleartext Transmission of Sensitive Information              | CWE-319          | OWASP:A04:2025                 | Yes         |
| Indirect Command Injection via User Controlled Environment   | CWE-78           | CWE Top 25, OWASP:A05:2025     | Yes         |
| External Control of System or Configuration Setting          | CWE-15           | OWASP:A02:2025                 | Yes         |
| Process Control                                              | CWE-114          | OWASP:A05:2025                 | Yes         |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | OWASP:A07:2025                 | Yes         |
| Hardcoded Secret                                             | CWE-547          | OWASP:A02:2025                 | Yes         |
| Improper Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A05:2025                 | Yes         |
| Disabled Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A05:2025                 | Yes         |
| Inadequate Padding for AES encryption                        | CWE-326          | OWASP:A04:2025                 | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A04:2025                 | Yes         |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A04:2025                 | Yes         |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A04:2025                 | Yes         |
| Improper Validation of Certificate with Host Mismatch        | CWE-297          | OWASP:A07:2025                 | Yes         |
| Java Naming and Directory Interface (JNDI) Injection         | CWE-074          | OWASP:A05:2025                 | Yes         |
| Improper Authentication                                      | CWE-287          | OWASP:A07:2025                 | Yes         |
| LDAP Injection                                               | CWE-90           | OWASP:A05:2025                 | Yes         |
| Use of Hardcoded Credentials                                 | CWE-798          | OWASP:A07:2025                 | Yes         |
| The cipher text is equal to the provided input plain text    | CWE-311          | OWASP:A06:2025                 | Yes         |
| Use of Hardcoded, Security-relevant Constants                | CWE-547          | OWASP:A02:2025                 | Yes         |
| Open Redirect                                                | CWE-601          | OWASP:A01:2025                 | Yes         |
| Path Traversal                                               | CWE-23           | OWASP:A01:2025                 | Yes         |
| Regular expression injection                                 | CWE-400, CWE-730 | None                           | Yes         |
| Unprotected Storage of Credentials                           | CWE-256          | OWASP:A06:2025                 | Yes         |
| Server Information Exposure                                  | CWE-209          | OWASP:A10:2025                 | Yes         |
| Cross-site Scripting (XSS)                                   | CWE-79           | CWE Top 25, OWASP:A05:2025     | Yes         |
| Android World Writeable/Readable File Permission Found       | CWE-732          | OWASP:A01:2025                 | Yes         |
| SQL Injection                                                | CWE-89           | CWE Top 25, OWASP:A05:2025     | Yes         |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | CWE Top 25, OWASP:A01:2025     | Yes         |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A04:2025                 | Yes         |
| Observable Timing Discrepancy (Timing Attack)                | CWE-208          | None                           | Yes         |
| Origin Validation Error                                      | CWE-942, CWE-346 | OWASP:A02:2025, OWASP:A07:2025 | Yes         |
| Improper Certificate Validation                              | CWE-295          | OWASP:A07:2025                 | Yes         |
| Cryptographic Issues                                         | CWE-310          | None                           | Yes         |
| Trust Boundary Violation                                     | CWE-501          | OWASP:A06:2025                 | Yes         |
| Use of Externally-Controlled Format String                   | CWE-134          | None                           | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A02:2025                 | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A02:2025                 | Yes         |
| Insufficient Session Expiration                              | CWE-613          | OWASP:A07:2025                 | Yes         |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A02:2025                 | Yes         |
| XPath Injection                                              | CWE-643          | OWASP:A05:2025                 | Yes         |
