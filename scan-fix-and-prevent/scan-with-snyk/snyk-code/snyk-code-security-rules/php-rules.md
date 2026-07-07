# PHP rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Code Injection                                               | CWE-94           | Sans Top 25, OWASP:A03 | Yes         |
| Command Injection                                            | CWE-78           | Sans Top 25, OWASP:A03 | Yes         |
| Deserialization of Untrusted Data                            | CWE-502          | Sans Top 25, OWASP:A08 | Yes         |
| Improper Access Control: Email Content Injection             | CWE-284          | OWASP:A01              | Yes         |
| File Inclusion                                               | CWE-98           | OWASP:A03              | Yes         |
| Use of Hardcoded Credentials                                 | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | Yes         |
| Hardcoded Secret                                             | CWE-547          | OWASP:A05              | Yes         |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | Yes         |
| Inadequate Padding for Public Key Encryption                 | CWE-326          | OWASP:A02              | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A02              | Yes         |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | Yes         |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A02              | Yes         |
| Cross-site Scripting (XSS)                                   | CWE-79           | Sans Top 25, OWASP:A03 | Yes         |
| Allocation of Resources Without Limits or Throttling         | CWE-770          | None                   | Yes         |
| Open Redirect                                                | CWE-601          | OWASP:A01              | Yes         |
| Path Traversal                                               | CWE-23           | OWASP:A01              | Yes         |
| Information Exposure                                         | CWE-200          | OWASP:A01              | Yes         |
| SQL Injection                                                | CWE-89           | Sans Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | Sans Top 25, OWASP:A10 | Yes         |
| Origin Validation Error                                      | CWE-942, CWE-346 | OWASP:A05, OWASP:A07   | Yes         |
| Improper Restriction of Rendered UI Layers or Frames         | CWE-1021         | OWASP:A04              | Yes         |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A02              | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A05              | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | Yes         |
| XPath Injection                                              | CWE-643          | OWASP:A03              | Yes         |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A05              | Yes         |
| Arbitrary File Write via Archive Extraction (Zip Slip)       | CWE-22           | Sans Top 25, OWASP:A01 | Yes         |
| Regular Expression Denial of Service (ReDoS)                 | CWE-400          | None                   | Yes         |
