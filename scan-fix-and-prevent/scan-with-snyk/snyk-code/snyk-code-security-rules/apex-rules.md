---
description: Snyk Code security rules for Apex
nav_context: agnostic
---

# Apex rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                                    | CWEs             | Security Categories                                                  |
| ------------------------------------------------------------ | ---------------- | -------------------------------------------------------------------- |
| Access Violation                                             | CWE-284, CWE-285 | CWE Top 25, OWASP:A01:2025, OWASP-API:API1:2023, OWASP-API:API5:2023 |
| Clear Text Sensitive Storage                                 | CWE-200, CWE-312 | CWE Top 25, OWASP:A01:2025, OWASP:A06:2025, OWASP-API:API10:2023     |
| Command Injection                                            | CWE-78           | CWE Top 25, OWASP:A05:2025                                           |
| Improper Access Control: Email Content Injection             | CWE-284          | CWE Top 25, OWASP:A01:2025                                           |
| Use of Hardcoded Credentials                                 | CWE-798, CWE-259 | OWASP:A07:2025                                                       |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | OWASP:A07:2025                                                       |
| Hardcoded Secret                                             | CWE-547          | OWASP:A02:2025                                                       |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A04:2025                                                       |
| Insecure Data Transmission                                   | CWE-319          | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023            |
| Open Redirect                                                | CWE-601          | OWASP:A01:2025                                                       |
| Cross-site Scripting (XSS)                                   | CWE-79           | CWE Top 25, OWASP:A05:2025                                           |
| Regular expression injection                                 | CWE-400, CWE-730 | OWASP-API:API4:2023                                                  |
| SOQL Injection                                               | CWE-89           | CWE Top 25, OWASP:A05:2025                                           |
| SOSL Injection                                               | CWE-89           | CWE Top 25, OWASP:A05:2025                                           |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023                      |
| Unverified Password Change                                   | CWE-620          | OWASP:A07:2025                                                       |
| Unsafe SOQL Concatenation                                    | CWE-89           | CWE Top 25, OWASP:A05:2025                                           |
| Unsafe SOSL Concatenation                                    | CWE-89           | CWE Top 25, OWASP:A05:2025                                           |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A02:2025                                                       |
| XML Injection                                                | CWE-91           | OWASP:A05:2025                                                       |
