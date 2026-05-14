# Apex rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Access Violation                                             | CWE-284, CWE-285 | OWASP:A01              | Yes         |
| Clear Text Sensitive Storage                                 | CWE-200, CWE-312 | OWASP:A01, OWASP:A04   | No          |
| Command Injection                                            | CWE-78           | Sans Top 25, OWASP:A03 | No          |
| Improper Access Control: Email Content Injection             | CWE-284          | OWASP:A01              | No          |
| Use of Hardcoded Credentials                                 | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | No          |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | No          |
| Hardcoded Secret                                             | CWE-547          | OWASP:A05              | No          |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | Yes         |
| Insecure Data Transmission                                   | CWE-319          | OWASP:A02              | No          |
| Open Redirect                                                | CWE-601          | OWASP:A01              | No          |
| Cross-site Scripting (XSS)                                   | CWE-79           | Sans Top 25, OWASP:A03 | No          |
| Regular expression injection                                 | CWE-400, CWE-730 | None                   | No          |
| SOQL Injection                                               | CWE-89           | Sans Top 25, OWASP:A03 | No          |
| SOSL Injection                                               | CWE-89           | Sans Top 25, OWASP:A03 | No          |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | Sans Top 25, OWASP:A10 | No          |
| Unverified Password Change                                   | CWE-620          | OWASP:A07              | No          |
| Unsafe SOQL Concatenation                                    | CWE-89           | Sans Top 25, OWASP:A03 | No          |
| Unsafe SOSL Concatenation                                    | CWE-89           | Sans Top 25, OWASP:A03 | No          |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | Yes         |
| XML Injection                                                | CWE-91           | OWASP:A03              | No          |
