---
description: Snyk Code security rules for Scala
nav_context: agnostic
---

# Scala rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2021 edition) category the rule belongs to, if any, and whether [SANS Top 25](https://www.sans.org/top25-software-errors/) includes it.
* **Autofixable**: Whether Snyk Agent Fix can fix the rule automatically. Snyk includes this column only for languages Snyk Agent Fix supports.

| Rule Name                                                    | CWEs             | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Use of Potentially Dangerous Function                        | CWE-676          | None                   | Yes         |
| Cleartext Storage of Sensitive Information in a Cookie       | CWE-315          | OWASP:A05              | Yes         |
| Code Injection                                               | CWE-94           | SANS Top 25, OWASP:A03 | Yes         |
| Command Injection                                            | CWE-78           | SANS Top 25, OWASP:A03 | Yes         |
| Deserialization of Untrusted Data                            | CWE-502          | SANS Top 25, OWASP:A08 | Yes         |
| Cross-Site Request Forgery (CSRF)                            | CWE-352          | SANS Top 25, OWASP:A01 | Yes         |
| Information Exposure                                         | CWE-200          | OWASP:A01              | Yes         |
| Cleartext Transmission of Sensitive Information              | CWE-319          | OWASP:A02              | Yes         |
| Indirect Command Injection via User Controlled Environment   | CWE-78           | SANS Top 25, OWASP:A03 | Yes         |
| External Control of System or Configuration Setting          | CWE-15           | OWASP:A05              | Yes         |
| Process Control                                              | CWE-114          | None                   | Yes         |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | SANS Top 25, OWASP:A07 | Yes         |
| Hardcoded Secret                                             | CWE-547          | OWASP:A05              | Yes         |
| Improper Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A03              | Yes         |
| Disabled Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A03              | Yes         |
| Inadequate Padding for AES encryption                        | CWE-326          | OWASP:A02              | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A02              | Yes         |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | Yes         |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A02              | Yes         |
| Improper Validation of Certificate with Host Mismatch        | CWE-297          | OWASP:A07              | Yes         |
| Java Naming and Directory Interface (JNDI) Injection         | CWE-074          | None                   | Yes         |
| Improper Authentication                                      | CWE-287          | SANS Top 25, OWASP:A07 | Yes         |
| LDAP Injection                                               | CWE-90           | OWASP:A03              | Yes         |
| Use of Hardcoded Credentials                                 | CWE-798          | SANS Top 25, OWASP:A07 | Yes         |
| The cipher text is equal to the provided input plain text    | CWE-311          | OWASP:A04              | Yes         |
| Use of Hardcoded, Security-relevant Constants                | CWE-547          | OWASP:A05              | Yes         |
| Open Redirect                                                | CWE-601          | OWASP:A01              | Yes         |
| Path Traversal                                               | CWE-23           | OWASP:A01              | Yes         |
| Regular expression injection                                 | CWE-400, CWE-730 | None                   | Yes         |
| Unprotected Storage of Credentials                           | CWE-256          | OWASP:A04              | Yes         |
| Server Information Exposure                                  | CWE-209          | OWASP:A04              | Yes         |
| Cross-site Scripting (XSS)                                   | CWE-79           | SANS Top 25, OWASP:A03 | Yes         |
| Android World Writeable/Readable File Permission Found       | CWE-732          | None                   | Yes         |
| SQL Injection                                                | CWE-89           | SANS Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | SANS Top 25, OWASP:A10 | Yes         |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A02              | Yes         |
| Observable Timing Discrepancy (Timing Attack)                | CWE-208          | None                   | Yes         |
| Origin Validation Error                                      | CWE-942, CWE-346 | OWASP:A05, OWASP:A07   | Yes         |
| Improper Certificate Validation                              | CWE-295          | OWASP:A07              | Yes         |
| Cryptographic Issues                                         | CWE-310          | OWASP:A02              | Yes         |
| Trust Boundary Violation                                     | CWE-501          | OWASP:A04              | Yes         |
| Use of Externally-Controlled Format String                   | CWE-134          | None                   | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A05              | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | Yes         |
| Insufficient Session Expiration                              | CWE-613          | OWASP:A07              | Yes         |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A05              | Yes         |
| XPath Injection                                              | CWE-643          | OWASP:A03              | Yes         |
