---
description: Snyk Code security rules for Go
nav_context: agnostic
---

# Go rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2021 edition) category the rule belongs to, if any, and whether [SANS Top 25](https://www.sans.org/top25-software-errors/) includes it.
* **Autofixable**: Whether Snyk Agent Fix can fix the rule automatically. Snyk includes this column only for languages Snyk Agent Fix supports.

| Rule Name                                                      | CWEs             | Security Categories    | Autofixable |
| -------------------------------------------------------------- | ---------------- | ---------------------- | ----------- |
| Clear Text Logging                                             | CWE-200, CWE-312 | OWASP:A01, OWASP:A04   | Yes         |
| Command Injection                                              | CWE-78           | SANS Top 25, OWASP:A03 | Yes         |
| Improper Access Control: Email Content Injection               | CWE-284          | OWASP:A01              | Yes         |
| Generation of Error Message Containing Sensitive Information   | CWE-209          | OWASP:A04              | Yes         |
| Hardcoded Secret                                               | CWE-547          | OWASP:A05              | Yes         |
| Use of Hardcoded Passwords                                     | CWE-798, CWE-259 | SANS Top 25, OWASP:A07 | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm               | CWE-327          | OWASP:A02              | Yes         |
| Use of Password Hash With Insufficient Computational Effort    | CWE-916          | OWASP:A02              | Yes         |
| Insecure TLS Configuration                                     | CWE-327          | OWASP:A02              | Yes         |
| Use of Insufficiently Random Values                            | CWE-330          | OWASP:A02              | Yes         |
| Use of Hardcoded Credentials                                   | CWE-798          | SANS Top 25, OWASP:A07 | Yes         |
| Open Redirect                                                  | CWE-601          | OWASP:A01              | Yes         |
| Path Traversal                                                 | CWE-23           | OWASP:A01              | Yes         |
| SQL Injection                                                  | CWE-89           | SANS Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                             | CWE-918          | SANS Top 25, OWASP:A10 | Yes         |
| Improper Neutralization of Directives in Statically Saved Code | CWE-96           | OWASP:A03              | Yes         |
| Improper Certificate Validation                                | CWE-295          | OWASP:A07              | Yes         |
| Inadequate Encryption Strength                                 | CWE-326          | OWASP:A02              | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                       | CWE-1004         | OWASP:A05              | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute   | CWE-614          | OWASP:A05              | Yes         |
| Cross-site Scripting (XSS)                                     | CWE-79           | SANS Top 25, OWASP:A03 | Yes         |
| XPath Injection                                                | CWE-643          | OWASP:A03              | Yes         |
