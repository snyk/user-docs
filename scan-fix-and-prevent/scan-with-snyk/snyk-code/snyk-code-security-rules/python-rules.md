# Python rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                                  | CWE(s)           | Security Categories    | Autofixable |
| -------------------------------------------------------------------------- | ---------------- | ---------------------- | ----------- |
| Authentication over HTTP                                                   | CWE-319          | OWASP:A02              | Yes         |
| Binding to all network interfaces may open service to unintended traffic   | CWE-284          | OWASP:A01              | Yes         |
| Broken User Authentication                                                 | CWE-287          | Sans Top 25, OWASP:A07 | No          |
| Code Injection                                                             | CWE-94           | Sans Top 25, OWASP:A03 | Yes         |
| Command Injection                                                          | CWE-78           | Sans Top 25, OWASP:A03 | Yes         |
| Deserialization of Untrusted Data                                          | CWE-502          | Sans Top 25, OWASP:A08 | Yes         |
| Cross-Site Request Forgery (CSRF)                                          | CWE-352          | Sans Top 25, OWASP:A01 | Yes         |
| Password Requirements Not Enforced in Django Application                   | CWE-521          | OWASP:A07              | No          |
| Use of Hardcoded Cryptographic Initialization Value                        | CWE-329          | OWASP:A02              | No          |
| Use of Hardcoded Cryptographic Key                                         | CWE-321          | OWASP:A02              | No          |
| Hardcoded Secret                                                           | CWE-547          | OWASP:A05              | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm                           | CWE-327          | OWASP:A02              | No          |
| Insecure default value                                                     | CWE-453          | None                   | No          |
| Insecure File Permissions                                                  | CWE-732          | None                   | Yes         |
| Use of Password Hash With Insufficient Computational Effort                | CWE-916          | OWASP:A02              | Yes         |
| Insecure Temporary File                                                    | CWE-377          | OWASP:A01              | Yes         |
| Insecure Xml Parser                                                        | CWE-611          | OWASP:A05              | No          |
| Jinja auto-escape is set to false.                                         | CWE-79           | Sans Top 25, OWASP:A03 | Yes         |
| LDAP Injection                                                             | CWE-90           | OWASP:A03              | No          |
| Improper Handling of Insufficient Permissions or Privileges                | CWE-280          | OWASP:A04              | No          |
| Use of Hardcoded Credentials                                               | CWE-798          | Sans Top 25, OWASP:A07 | No          |
| Use of Hardcoded Passwords                                                 | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | No          |
| NoSQL Injection                                                            | CWE-943          | None                   | No          |
| Open Redirect                                                              | CWE-601          | OWASP:A01              | No          |
| Path Traversal                                                             | CWE-23           | OWASP:A01              | Yes         |
| Debug Mode Enabled                                                         | CWE-489          | None                   | Yes         |
| Improper Certificate Validation                                            | CWE-295          | OWASP:A07              | No          |
| Server Information Exposure                                                | CWE-209          | OWASP:A04              | No          |
| SQL Injection                                                              | CWE-89           | Sans Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                                         | CWE-918          | Sans Top 25, OWASP:A10 | No          |
| Improper Neutralization of Directives in Statically Saved Code             | CWE-96           | OWASP:A03              | No          |
| Inadequate Encryption Strength                                             | CWE-326          | OWASP:A02              | Yes         |
| Arbitrary File Write via Archive Extraction (Tar Slip)                     | CWE-22           | Sans Top 25, OWASP:A01 | No          |
| Origin Validation Error                                                    | CWE-942, CWE-346 | OWASP:A05, OWASP:A07   | No          |
| Cryptographic Issues                                                       | CWE-310          | OWASP:A02              | No          |
| Use of Insufficiently Random Values                                        | CWE-330          | OWASP:A02              | No          |
| Python 2 source code                                                       | CWE-1104         | OWASP:A06              | No          |
| Selection of Less-Secure Algorithm During Negotiation (SSL instead of TLS) | CWE-757          | OWASP:A02              | No          |
| Sensitive Cookie Without 'HttpOnly' Flag                                   | CWE-1004         | OWASP:A05              | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute               | CWE-614          | OWASP:A05              | Yes         |
| Cross-site Scripting (XSS)                                                 | CWE-79           | Sans Top 25, OWASP:A03 | Yes         |
| XPath Injection                                                            | CWE-643          | OWASP:A03              | No          |
| Regular Expression Denial of Service (ReDoS)                               | CWE-400          | None                   | No          |
