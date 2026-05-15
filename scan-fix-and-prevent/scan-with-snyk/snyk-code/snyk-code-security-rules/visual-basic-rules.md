# Visual Basic rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Debug Features Enabled                                       | CWE-215          | None                   | No          |
| Usage of BinaryFormatter                                     | CWE-502          | Sans Top 25, OWASP:A08 | No          |
| Code Injection                                               | CWE-94           | Sans Top 25, OWASP:A03 | No          |
| Command Injection                                            | CWE-78           | Sans Top 25, OWASP:A03 | No          |
| Deserialization of Untrusted Data                            | CWE-502          | Sans Top 25, OWASP:A08 | No          |
| Hardcoded Secret                                             | CWE-547          | OWASP:A05              | No          |
| Improper Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A03              | No          |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A02              | No          |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | No          |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A02              | No          |
| Use of Hardcoded Credentials                                 | CWE-798          | Sans Top 25, OWASP:A07 | No          |
| Open Redirect                                                | CWE-601          | OWASP:A01              | No          |
| Path Traversal                                               | CWE-23           | OWASP:A01              | No          |
| Regular expression injection                                 | CWE-400, CWE-730 | None                   | No          |
| Request Validation Disabled                                  | CWE-554          | None                   | No          |
| SQL Injection                                                | CWE-89           | Sans Top 25, OWASP:A03 | No          |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | Sans Top 25, OWASP:A10 | No          |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A02              | No          |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A05              | No          |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | No          |
| Cross-site Scripting (XSS)                                   | CWE-79           | Sans Top 25, OWASP:A03 | No          |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A05              | No          |
| XML Injection                                                | CWE-91           | OWASP:A03              | No          |
| XPath Injection                                              | CWE-643          | OWASP:A03              | No          |
