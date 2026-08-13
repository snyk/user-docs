---
description: Snyk Code security rules for C# and ASP.NET
nav_context: agnostic
---

# C# and ASP.NET rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2021 edition) category the rule belongs to, if any, and whether [SANS Top 25](https://www.sans.org/top25-software-errors/) includes it.
* **Autofixable**: Whether Snyk Agent Fix can fix the rule automatically. Snyk includes this column only for languages Snyk Agent Fix supports.

| Rule Name                                                         | CWEs             | Security Categories    | Autofixable |
| ----------------------------------------------------------------- | ---------------- | ---------------------- | ----------- |
| Anti-forgery token validation disabled                            | CWE-352          | SANS Top 25, OWASP:A01 | Yes         |
| Debug Features Enabled                                            | CWE-215          | None                   | Yes         |
| Usage of BinaryFormatter                                          | CWE-502          | SANS Top 25, OWASP:A08 | Yes         |
| Cleartext Storage of Sensitive Information in a Cookie            | CWE-315          | OWASP:A05              | Yes         |
| Code Injection                                                    | CWE-94           | SANS Top 25, OWASP:A03 | Yes         |
| Command Injection                                                 | CWE-78           | SANS Top 25, OWASP:A03 | Yes         |
| Deserialization of Untrusted Data                                 | CWE-502          | SANS Top 25, OWASP:A08 | Yes         |
| Hardcoded Secret                                                  | CWE-547          | OWASP:A05              | Yes         |
| Improper Neutralization of CRLF Sequences in HTTP Headers         | CWE-113          | OWASP:A03              | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm                  | CWE-327          | OWASP:A02              | Yes         |
| Use of Password Hash With Insufficient Computational Effort       | CWE-916          | OWASP:A02              | Yes         |
| Use of Insufficiently Random Values                               | CWE-330          | OWASP:A02              | Yes         |
| Insecure Data Transmission                                        | CWE-319          | OWASP:A02              | Yes         |
| LDAP Injection                                                    | CWE-90           | OWASP:A03              | Yes         |
| Log Forging                                                       | CWE-117          | OWASP:A09              | Yes         |
| Use of Hardcoded Credentials                                      | CWE-798          | SANS Top 25, OWASP:A07 | Yes         |
| Open Redirect                                                     | CWE-601          | OWASP:A01              | Yes         |
| Path Traversal                                                    | CWE-23           | OWASP:A01              | Yes         |
| Exposure of Private Personal Information to an Unauthorized Actor | CWE-359          | OWASP:A01              | Yes         |
| Regular expression injection                                      | CWE-400, CWE-730 | None                   | Yes         |
| Request Validation Disabled                                       | CWE-554          | None                   | Yes         |
| Information Exposure                                              | CWE-200          | OWASP:A01              | Yes         |
| SQL Injection                                                     | CWE-89           | SANS Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                                | CWE-918          | SANS Top 25, OWASP:A10 | Yes         |
| Inadequate Encryption Strength                                    | CWE-326          | OWASP:A02              | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                          | CWE-1004         | OWASP:A05              | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute      | CWE-614          | OWASP:A05              | Yes         |
| Cross-site Scripting (XSS)                                        | CWE-79           | SANS Top 25, OWASP:A03 | Yes         |
| XML External Entity (XXE) Injection                               | CWE-611          | OWASP:A05              | Yes         |
| XAML Injection                                                    | CWE-611          | OWASP:A05              | Yes         |
| XML Injection                                                     | CWE-91           | OWASP:A03              | Yes         |
| XPath Injection                                                   | CWE-643          | OWASP:A03              | Yes         |
| Arbitrary File Write via Archive Extraction (Zip Slip)            | CWE-22           | SANS Top 25, OWASP:A01 | Yes         |
