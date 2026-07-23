---
description: Snyk Code security rules for Visual Basic
---

# Visual Basic rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories                             | Autofixable |
| ------------------------------------------------------------ | ---------------- | ----------------------------------------------- | ----------- |
| Debug Features Enabled                                       | CWE-215          | OWASP:A10:2025                                  | Yes         |
| Usage of BinaryFormatter                                     | CWE-502          | CWE Top 25, OWASP:A08:2025                      | Yes         |
| Code Injection                                               | CWE-94           | CWE Top 25, OWASP:A05:2025                      | Yes         |
| Command Injection                                            | CWE-78           | CWE Top 25, OWASP:A05:2025                      | Yes         |
| Deserialization of Untrusted Data                            | CWE-502          | CWE Top 25, OWASP:A08:2025                      | Yes         |
| Hardcoded Secret                                             | CWE-547          | OWASP:A02:2025                                  | Yes         |
| Improper Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A05:2025                                  | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A04:2025                                  | Yes         |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A04:2025                                  | Yes         |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A04:2025                                  | Yes         |
| Use of Hardcoded Credentials                                 | CWE-798          | OWASP:A07:2025                                  | Yes         |
| Open Redirect                                                | CWE-601          | OWASP:A01:2025                                  | Yes         |
| Path Traversal                                               | CWE-23           | OWASP:A01:2025                                  | Yes         |
| Regular expression injection                                 | CWE-400, CWE-730 | OWASP-API:API4:2023                             | Yes         |
| Request Validation Disabled                                  | CWE-554          | None                                            | Yes         |
| SQL Injection                                                | CWE-89           | CWE Top 25, OWASP:A05:2025                      | Yes         |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023 | Yes         |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A04:2025                                  | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A02:2025                                  | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A02:2025                                  | Yes         |
| Cross-site Scripting (XSS)                                   | CWE-79           | CWE Top 25, OWASP:A05:2025                      | Yes         |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A02:2025                                  | Yes         |
| XML Injection                                                | CWE-91           | OWASP:A05:2025                                  | Yes         |
| XPath Injection                                              | CWE-643          | OWASP:A05:2025                                  | Yes         |
