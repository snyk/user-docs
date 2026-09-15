---
description: Snyk Code security rules for C# and ASP.NET
nav_context: agnostic
---

# C# and ASP.NET rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                                         | CWEs             | Security Categories                                       |
| ----------------------------------------------------------------- | ---------------- | --------------------------------------------------------- |
| Anti-forgery token validation disabled                            | CWE-352          | CWE Top 25, OWASP:A01:2025                                |
| Debug Features Enabled                                            | CWE-215          | OWASP:A10:2025                                            |
| Usage of BinaryFormatter                                          | CWE-502          | CWE Top 25, OWASP:A08:2025                                |
| Cleartext Storage of Sensitive Information in a Cookie            | CWE-315          | OWASP:A02:2025                                            |
| Code Injection                                                    | CWE-94           | CWE Top 25, OWASP:A05:2025                                |
| Command Injection                                                 | CWE-78           | CWE Top 25, OWASP:A05:2025                                |
| Deserialization of Untrusted Data                                 | CWE-502          | CWE Top 25, OWASP:A08:2025                                |
| Hardcoded Secret                                                  | CWE-547          | OWASP:A02:2025                                            |
| Improper Neutralization of CRLF Sequences in HTTP Headers         | CWE-113          | OWASP:A05:2025                                            |
| Use of a Broken or Risky Cryptographic Algorithm                  | CWE-327          | OWASP:A04:2025                                            |
| Use of Password Hash With Insufficient Computational Effort       | CWE-916          | OWASP:A04:2025                                            |
| Use of Insufficiently Random Values                               | CWE-330          | OWASP:A04:2025                                            |
| Insecure Data Transmission                                        | CWE-319          | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023 |
| LDAP Injection                                                    | CWE-90           | OWASP:A05:2025                                            |
| Log Forging                                                       | CWE-117          | OWASP:A09:2025                                            |
| Use of Hardcoded Credentials                                      | CWE-798          | OWASP:A07:2025                                            |
| Open Redirect                                                     | CWE-601          | OWASP:A01:2025                                            |
| Path Traversal                                                    | CWE-23           | OWASP:A01:2025                                            |
| Exposure of Private Personal Information to an Unauthorized Actor | CWE-359          | OWASP:A01:2025                                            |
| Regular expression injection                                      | CWE-400, CWE-730 | OWASP-API:API4:2023                                       |
| Request Validation Disabled                                       | CWE-554          | None                                                      |
| Information Exposure                                              | CWE-200          | CWE Top 25, OWASP:A01:2025, OWASP-API:API10:2023          |
| SQL Injection                                                     | CWE-89           | CWE Top 25, OWASP:A05:2025                                |
| Server-Side Request Forgery (SSRF)                                | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023           |
| Inadequate Encryption Strength                                    | CWE-326          | OWASP:A04:2025                                            |
| Sensitive Cookie Without 'HttpOnly' Flag                          | CWE-1004         | OWASP:A02:2025                                            |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute      | CWE-614          | OWASP:A02:2025                                            |
| Cross-site Scripting (XSS)                                        | CWE-79           | CWE Top 25, OWASP:A05:2025                                |
| XML External Entity (XXE) Injection                               | CWE-611          | OWASP:A02:2025                                            |
| XAML Injection                                                    | CWE-611          | OWASP:A02:2025                                            |
| XML Injection                                                     | CWE-91           | OWASP:A05:2025                                            |
| XPath Injection                                                   | CWE-643          | OWASP:A05:2025                                            |
| Arbitrary File Write via Archive Extraction (Zip Slip)            | CWE-22           | CWE Top 25, OWASP:A01:2025                                |
