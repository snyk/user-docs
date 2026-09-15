---
description: Snyk Code security rules for Python
nav_context: agnostic
---

# Python rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                                                  | CWEs             | Security Categories                                       |
| -------------------------------------------------------------------------- | ---------------- | --------------------------------------------------------- |
| Authentication over HTTP                                                   | CWE-319          | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023 |
| Binding to all network interfaces may open service to unintended traffic   | CWE-284          | CWE Top 25, OWASP:A01:2025                                |
| Broken User Authentication                                                 | CWE-287          | OWASP:A07:2025                                            |
| Code Injection                                                             | CWE-94           | CWE Top 25, OWASP:A05:2025                                |
| Command Injection                                                          | CWE-78           | CWE Top 25, OWASP:A05:2025                                |
| Deserialization of Untrusted Data                                          | CWE-502          | CWE Top 25, OWASP:A08:2025                                |
| Cross-Site Request Forgery (CSRF)                                          | CWE-352          | CWE Top 25, OWASP:A01:2025                                |
| Password Requirements Not Enforced in Django Application                   | CWE-521          | OWASP:A07:2025                                            |
| Use of Hardcoded Cryptographic Initialization Value                        | CWE-329          | OWASP:A04:2025                                            |
| Use of Hardcoded Cryptographic Key                                         | CWE-321          | OWASP:A04:2025                                            |
| Hardcoded Secret                                                           | CWE-547          | OWASP:A02:2025                                            |
| Use of a Broken or Risky Cryptographic Algorithm                           | CWE-327          | OWASP:A04:2025                                            |
| Insecure default value                                                     | CWE-453          | None                                                      |
| Insecure File Permissions                                                  | CWE-732          | OWASP:A01:2025                                            |
| Use of Password Hash With Insufficient Computational Effort                | CWE-916          | OWASP:A04:2025                                            |
| Insecure Temporary File                                                    | CWE-377          | OWASP:A01:2025                                            |
| Insecure Xml Parser                                                        | CWE-611          | OWASP:A02:2025                                            |
| Jinja auto-escape is set to false.                                         | CWE-79           | CWE Top 25, OWASP:A05:2025                                |
| LDAP Injection                                                             | CWE-90           | OWASP:A05:2025                                            |
| Improper Handling of Insufficient Permissions or Privileges                | CWE-280          | OWASP:A10:2025                                            |
| Use of Hardcoded Credentials                                               | CWE-798          | OWASP:A07:2025                                            |
| Use of Hardcoded Passwords                                                 | CWE-798, CWE-259 | OWASP:A07:2025                                            |
| NoSQL Injection                                                            | CWE-943          | None                                                      |
| Open Redirect                                                              | CWE-601          | OWASP:A01:2025                                            |
| Path Traversal                                                             | CWE-23           | OWASP:A01:2025                                            |
| Debug Mode Enabled                                                         | CWE-489          | OWASP:A02:2025                                            |
| Improper Certificate Validation                                            | CWE-295          | OWASP:A07:2025                                            |
| Server Information Exposure                                                | CWE-209          | OWASP:A10:2025, OWASP-API:API8:2023                       |
| SQL Injection                                                              | CWE-89           | CWE Top 25, OWASP:A05:2025                                |
| Server-Side Request Forgery (SSRF)                                         | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023           |
| Improper Neutralization of Directives in Statically Saved Code             | CWE-96           | OWASP:A05:2025                                            |
| Inadequate Encryption Strength                                             | CWE-326          | OWASP:A04:2025                                            |
| Arbitrary File Write via Archive Extraction (Tar Slip)                     | CWE-22           | CWE Top 25, OWASP:A01:2025                                |
| Origin Validation Error                                                    | CWE-942, CWE-346 | OWASP:A02:2025, OWASP:A07:2025, OWASP-API:API8:2023       |
| Cryptographic Issues                                                       | CWE-310          | None                                                      |
| Use of Insufficiently Random Values                                        | CWE-330          | OWASP:A04:2025                                            |
| Python 2 source code                                                       | CWE-1104         | OWASP:A03:2025                                            |
| Selection of Less-Secure Algorithm During Negotiation (SSL instead of TLS) | CWE-757          | OWASP:A04:2025                                            |
| Sensitive Cookie Without 'HttpOnly' Flag                                   | CWE-1004         | OWASP:A02:2025                                            |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute               | CWE-614          | OWASP:A02:2025                                            |
| Cross-site Scripting (XSS)                                                 | CWE-79           | CWE Top 25, OWASP:A05:2025                                |
| XPath Injection                                                            | CWE-643          | OWASP:A05:2025                                            |
| Regular Expression Denial of Service (ReDoS)                               | CWE-400          | OWASP-API:API4:2023                                       |
