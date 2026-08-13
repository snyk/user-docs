---
description: Snyk Code security rules for Apex
nav_context: agnostic
---

# Apex rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2025 edition) category the rule belongs to, if any, and whether [SANS Top 25](https://www.sans.org/top25-software-errors/) includes it.
* **Autofixable**: Whether Snyk Agent Fix can fix the rule automatically. Snyk includes this column only for languages Snyk Agent Fix supports.

| Rule Name                                                    | CWEs             | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Access Violation                                             | CWE-284, CWE-285 | OWASP:A01              | Yes         |
| Clear Text Sensitive Storage                                 | CWE-200, CWE-312 | OWASP:A01, OWASP:A04   | Yes         |
| Command Injection                                            | CWE-78           | SANS Top 25, OWASP:A03 | Yes         |
| Improper Access Control: Email Content Injection             | CWE-284          | OWASP:A01              | Yes         |
| Use of Hardcoded Credentials                                 | CWE-798, CWE-259 | SANS Top 25, OWASP:A07 | Yes         |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | SANS Top 25, OWASP:A07 | Yes         |
| Hardcoded Secret                                             | CWE-547          | OWASP:A05              | Yes         |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | Yes         |
| Insecure Data Transmission                                   | CWE-319          | OWASP:A02              | Yes         |
| Open Redirect                                                | CWE-601          | OWASP:A01              | Yes         |
| Cross-site Scripting (XSS)                                   | CWE-79           | SANS Top 25, OWASP:A03 | Yes         |
| Regular expression injection                                 | CWE-400, CWE-730 | None                   | Yes         |
| SOQL Injection                                               | CWE-89           | SANS Top 25, OWASP:A03 | Yes         |
| SOSL Injection                                               | CWE-89           | SANS Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | SANS Top 25, OWASP:A10 | Yes         |
| Unverified Password Change                                   | CWE-620          | OWASP:A07              | Yes         |
| Unsafe SOQL Concatenation                                    | CWE-89           | SANS Top 25, OWASP:A03 | Yes         |
| Unsafe SOSL Concatenation                                    | CWE-89           | SANS Top 25, OWASP:A03 | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | Yes         |
| XML Injection                                                | CWE-91           | OWASP:A03              | Yes         |
