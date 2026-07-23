---
description: Snyk Code security rules for Go
---

# Go rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/2025/)(2025 edition) category to which the rule belongs to, if any, and if it is included in [CWE Top 25](https://cwe.mitre.org/top25/), plus any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) and [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                      | CWE(s)           | Security Categories                                              | Autofixable |
| -------------------------------------------------------------- | ---------------- | ---------------------------------------------------------------- | ----------- |
| Clear Text Logging                                             | CWE-200, CWE-312 | CWE Top 25, OWASP:A01:2025, OWASP:A06:2025, OWASP-API:API10:2023 | Yes         |
| Command Injection                                              | CWE-78           | CWE Top 25, OWASP:A05:2025                                       | Yes         |
| Improper Access Control: Email Content Injection               | CWE-284          | CWE Top 25, OWASP:A01:2025                                       | Yes         |
| Generation of Error Message Containing Sensitive Information   | CWE-209          | OWASP:A10:2025, OWASP-API:API8:2023                              | Yes         |
| Hardcoded Secret                                               | CWE-547          | OWASP:A02:2025                                                   | Yes         |
| Use of Hardcoded Passwords                                     | CWE-798, CWE-259 | OWASP:A07:2025                                                   | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm               | CWE-327          | OWASP:A04:2025                                                   | Yes         |
| Use of Password Hash With Insufficient Computational Effort    | CWE-916          | OWASP:A04:2025                                                   | Yes         |
| Insecure TLS Configuration                                     | CWE-327          | OWASP:A04:2025                                                   | Yes         |
| Use of Insufficiently Random Values                            | CWE-330          | OWASP:A04:2025                                                   | Yes         |
| Use of Hardcoded Credentials                                   | CWE-798          | OWASP:A07:2025                                                   | Yes         |
| Open Redirect                                                  | CWE-601          | OWASP:A01:2025                                                   | Yes         |
| Path Traversal                                                 | CWE-23           | OWASP:A01:2025                                                   | Yes         |
| SQL Injection                                                  | CWE-89           | CWE Top 25, OWASP:A05:2025                                       | Yes         |
| Server-Side Request Forgery (SSRF)                             | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023                  | Yes         |
| Improper Neutralization of Directives in Statically Saved Code | CWE-96           | OWASP:A05:2025                                                   | Yes         |
| Improper Certificate Validation                                | CWE-295          | OWASP:A07:2025                                                   | Yes         |
| Inadequate Encryption Strength                                 | CWE-326          | OWASP:A04:2025                                                   | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                       | CWE-1004         | OWASP:A02:2025                                                   | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute   | CWE-614          | OWASP:A02:2025                                                   | Yes         |
| Cross-site Scripting (XSS)                                     | CWE-79           | CWE Top 25, OWASP:A05:2025                                       | Yes         |
| XPath Injection                                                | CWE-643          | OWASP:A05:2025                                                   | Yes         |
