---
description: Snyk Code security rules for Apex
---

# Apex rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s)**: The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category to which the rule belongs to, if any, and if it is included in [CWE Top 25](https://cwe.mitre.org/top25/), plus any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) and [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories                                                  | Autofixable |
| ------------------------------------------------------------ | ---------------- | -------------------------------------------------------------------- | ----------- |
| Access Violation                                             | CWE-284, CWE-285 | CWE Top 25, OWASP:A01:2025, OWASP-API:API1:2023, OWASP-API:API5:2023 | Yes         |
| Clear Text Sensitive Storage                                 | CWE-200, CWE-312 | CWE Top 25, OWASP:A01:2025, OWASP:A06:2025, OWASP-API:API10:2023     | Yes         |
| Command Injection                                            | CWE-78           | CWE Top 25, OWASP:A05:2025                                           | Yes         |
| Improper Access Control: Email Content Injection             | CWE-284          | CWE Top 25, OWASP:A01:2025                                           | Yes         |
| Use of Hardcoded Credentials                                 | CWE-798, CWE-259 | OWASP:A07:2025                                                       | Yes         |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | OWASP:A07:2025                                                       | Yes         |
| Hardcoded Secret                                             | CWE-547          | OWASP:A02:2025                                                       | Yes         |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A04:2025                                                       | Yes         |
| Insecure Data Transmission                                   | CWE-319          | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023            | Yes         |
| Open Redirect                                                | CWE-601          | OWASP:A01:2025                                                       | Yes         |
| Cross-site Scripting (XSS)                                   | CWE-79           | CWE Top 25, OWASP:A05:2025                                           | Yes         |
| Regular expression injection                                 | CWE-400, CWE-730 | OWASP-API:API4:2023                                                  | Yes         |
| SOQL Injection                                               | CWE-89           | CWE Top 25, OWASP:A05:2025                                           | Yes         |
| SOSL Injection                                               | CWE-89           | CWE Top 25, OWASP:A05:2025                                           | Yes         |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023                      | Yes         |
| Unverified Password Change                                   | CWE-620          | OWASP:A07:2025                                                       | Yes         |
| Unsafe SOQL Concatenation                                    | CWE-89           | CWE Top 25, OWASP:A05:2025                                           | Yes         |
| Unsafe SOSL Concatenation                                    | CWE-89           | CWE Top 25, OWASP:A05:2025                                           | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A02:2025                                                       | Yes         |
| XML Injection                                                | CWE-91           | OWASP:A05:2025                                                       | Yes         |
