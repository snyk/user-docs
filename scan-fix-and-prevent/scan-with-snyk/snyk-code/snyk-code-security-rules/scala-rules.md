---
description: Snyk Code security rules for Scala
nav_context: agnostic
---

# Scala rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                                    | CWEs             | Security Categories                                       |
| ------------------------------------------------------------ | ---------------- | --------------------------------------------------------- |
| Use of Potentially Dangerous Function                        | CWE-676          | OWASP:A06:2025                                            |
| Cleartext Storage of Sensitive Information in a Cookie       | CWE-315          | OWASP:A02:2025                                            |
| Code Injection                                               | CWE-94           | CWE Top 25, OWASP:A05:2025                                |
| Command Injection                                            | CWE-78           | CWE Top 25, OWASP:A05:2025                                |
| Deserialization of Untrusted Data                            | CWE-502          | CWE Top 25, OWASP:A08:2025                                |
| Cross-Site Request Forgery (CSRF)                            | CWE-352          | CWE Top 25, OWASP:A01:2025                                |
| Information Exposure                                         | CWE-200          | CWE Top 25, OWASP:A01:2025, OWASP-API:API10:2023          |
| Cleartext Transmission of Sensitive Information              | CWE-319          | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023 |
| Indirect Command Injection via User Controlled Environment   | CWE-78           | CWE Top 25, OWASP:A05:2025                                |
| External Control of System or Configuration Setting          | CWE-15           | OWASP:A02:2025                                            |
| Process Control                                              | CWE-114          | OWASP:A05:2025                                            |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | OWASP:A07:2025                                            |
| Hardcoded Secret                                             | CWE-547          | OWASP:A02:2025                                            |
| Improper Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A05:2025                                            |
| Disabled Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A05:2025                                            |
| Inadequate Padding for AES encryption                        | CWE-326          | OWASP:A04:2025                                            |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A04:2025                                            |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A04:2025                                            |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A04:2025                                            |
| Improper Validation of Certificate with Host Mismatch        | CWE-297          | OWASP:A07:2025                                            |
| Java Naming and Directory Interface (JNDI) Injection         | CWE-074          | OWASP:A05:2025                                            |
| Improper Authentication                                      | CWE-287          | OWASP:A07:2025                                            |
| LDAP Injection                                               | CWE-90           | OWASP:A05:2025                                            |
| Use of Hardcoded Credentials                                 | CWE-798          | OWASP:A07:2025                                            |
| The cipher text is equal to the provided input plain text    | CWE-311          | OWASP:A06:2025                                            |
| Use of Hardcoded, Security-relevant Constants                | CWE-547          | OWASP:A02:2025                                            |
| Open Redirect                                                | CWE-601          | OWASP:A01:2025                                            |
| Path Traversal                                               | CWE-23           | OWASP:A01:2025                                            |
| Regular expression injection                                 | CWE-400, CWE-730 | OWASP-API:API4:2023                                       |
| Unprotected Storage of Credentials                           | CWE-256          | OWASP:A06:2025                                            |
| Server Information Exposure                                  | CWE-209          | OWASP:A10:2025, OWASP-API:API8:2023                       |
| Cross-site Scripting (XSS)                                   | CWE-79           | CWE Top 25, OWASP:A05:2025                                |
| Android World Writeable/Readable File Permission Found       | CWE-732          | OWASP:A01:2025                                            |
| SQL Injection                                                | CWE-89           | CWE Top 25, OWASP:A05:2025                                |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023           |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A04:2025                                            |
| Observable Timing Discrepancy (Timing Attack)                | CWE-208          | None                                                      |
| Origin Validation Error                                      | CWE-942, CWE-346 | OWASP:A02:2025, OWASP:A07:2025, OWASP-API:API8:2023       |
| Improper Certificate Validation                              | CWE-295          | OWASP:A07:2025                                            |
| Cryptographic Issues                                         | CWE-310          | None                                                      |
| Trust Boundary Violation                                     | CWE-501          | OWASP:A06:2025                                            |
| Use of Externally-Controlled Format String                   | CWE-134          | None                                                      |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A02:2025                                            |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A02:2025                                            |
| Insufficient Session Expiration                              | CWE-613          | OWASP:A07:2025                                            |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A02:2025                                            |
| XPath Injection                                              | CWE-643          | OWASP:A05:2025                                            |
