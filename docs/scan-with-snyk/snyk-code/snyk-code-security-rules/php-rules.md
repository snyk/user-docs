# PHP rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Code Injection                                               | CWE-94           | Sans Top 25, OWASP:A03 | No          |
| Command Injection                                            | CWE-78           | Sans Top 25, OWASP:A03 | No          |
| Deserialization of Untrusted Data                            | CWE-502          | Sans Top 25, OWASP:A08 | No          |
| Improper Access Control: Email Content Injection             | CWE-284          | OWASP:A01              | No          |
| File Inclusion                                               | CWE-98           | OWASP:A03              | No          |
| Use of Hardcoded Credentials                                 | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | No          |
| Hardcoded Secret                                             | CWE-547          | OWASP:A05              | No          |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | No          |
| Inadequate Padding for Public Key Encryption                 | CWE-326          | OWASP:A02              | No          |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A02              | No          |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | No          |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A02              | No          |
| Cross-site Scripting (XSS)                                   | CWE-79           | Sans Top 25, OWASP:A03 | No          |
| Allocation of Resources Without Limits or Throttling         | CWE-770          | None                   | No          |
| Open Redirect                                                | CWE-601          | OWASP:A01              | No          |
| Path Traversal                                               | CWE-23           | OWASP:A01              | No          |
| Information Exposure                                         | CWE-200          | OWASP:A01              | No          |
| SQL Injection                                                | CWE-89           | Sans Top 25, OWASP:A03 | No          |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | Sans Top 25, OWASP:A10 | No          |
| Origin Validation Error                                      | CWE-942, CWE-346 | OWASP:A05, OWASP:A07   | No          |
| Improper Restriction of Rendered UI Layers or Frames         | CWE-1021         | OWASP:A04              | No          |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A02              | No          |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A05              | No          |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | No          |
| XPath Injection                                              | CWE-643          | OWASP:A03              | No          |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A05              | No          |
| Arbitrary File Write via Archive Extraction (Zip Slip)       | CWE-22           | Sans Top 25, OWASP:A01 | No          |
| Regular Expression Denial of Service (ReDoS)                 | CWE-400          | None                   | No          |
