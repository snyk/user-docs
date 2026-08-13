---
description: Snyk Code security rules for Go
nav_context: agnostic
---

# Go rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                                      | CWEs             | Security Categories                                              |
| -------------------------------------------------------------- | ---------------- | ---------------------------------------------------------------- |
| Clear Text Logging                                             | CWE-200, CWE-312 | CWE Top 25, OWASP:A01:2025, OWASP:A06:2025, OWASP-API:API10:2023 |
| Command Injection                                              | CWE-78           | CWE Top 25, OWASP:A05:2025                                       |
| Improper Access Control: Email Content Injection               | CWE-284          | CWE Top 25, OWASP:A01:2025                                       |
| Generation of Error Message Containing Sensitive Information   | CWE-209          | OWASP:A10:2025, OWASP-API:API8:2023                              |
| Hardcoded Secret                                               | CWE-547          | OWASP:A02:2025                                                   |
| Use of Hardcoded Passwords                                     | CWE-798, CWE-259 | OWASP:A07:2025                                                   |
| Use of a Broken or Risky Cryptographic Algorithm               | CWE-327          | OWASP:A04:2025                                                   |
| Use of Password Hash With Insufficient Computational Effort    | CWE-916          | OWASP:A04:2025                                                   |
| Insecure TLS Configuration                                     | CWE-327          | OWASP:A04:2025                                                   |
| Use of Insufficiently Random Values                            | CWE-330          | OWASP:A04:2025                                                   |
| Use of Hardcoded Credentials                                   | CWE-798          | OWASP:A07:2025                                                   |
| Open Redirect                                                  | CWE-601          | OWASP:A01:2025                                                   |
| Path Traversal                                                 | CWE-23           | OWASP:A01:2025                                                   |
| SQL Injection                                                  | CWE-89           | CWE Top 25, OWASP:A05:2025                                       |
| Server-Side Request Forgery (SSRF)                             | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023                  |
| Improper Neutralization of Directives in Statically Saved Code | CWE-96           | OWASP:A05:2025                                                   |
| Improper Certificate Validation                                | CWE-295          | OWASP:A07:2025                                                   |
| Inadequate Encryption Strength                                 | CWE-326          | OWASP:A04:2025                                                   |
| Sensitive Cookie Without 'HttpOnly' Flag                       | CWE-1004         | OWASP:A02:2025                                                   |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute   | CWE-614          | OWASP:A02:2025                                                   |
| Cross-site Scripting (XSS)                                     | CWE-79           | CWE Top 25, OWASP:A05:2025                                       |
| XPath Injection                                                | CWE-643          | OWASP:A05:2025                                                   |
