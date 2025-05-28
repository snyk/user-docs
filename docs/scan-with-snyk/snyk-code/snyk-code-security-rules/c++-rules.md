# C++ rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                                | CWE(s)           | Security Categories    | Autofixable |
| ------------------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Memory Allocation Of String Length                                       | CWE-170          | None                   | Yes         |
| Insecure Anonymous LDAP Binding                                          | CWE-287          | Sans Top 25, OWASP:A07 | No          |
| Buffer Overflow                                                          | CWE-122          | None                   | Yes         |
| Division By Zero                                                         | CWE-369          | None                   | No          |
| Missing Release of File Descriptor or Handle after Effective Lifetime    | CWE-775          | None                   | Yes         |
| Command Injection                                                        | CWE-78           | Sans Top 25, OWASP:A03 | No          |
| Dereference of a NULL Pointer                                            | CWE-476          | Sans Top 25            | No          |
| Double Free                                                              | CWE-415          | None                   | Yes         |
| Use of Externally-Controlled Format String                               | CWE-134          | None                   | Yes         |
| Use of Hardcoded Cryptographic Key                                       | CWE-321          | OWASP:A02              | No          |
| Improper Null Termination                                                | CWE-170          | None                   | No          |
| Use of Password Hash With Insufficient Computational Effort              | CWE-916          | OWASP:A02              | Yes         |
| Integer Overflow                                                         | CWE-190          | Sans Top 25            | No          |
| LDAP Injection                                                           | CWE-90           | OWASP:A03              | No          |
| Missing Release of Memory after Effective Lifetime                       | CWE-401          | None                   | Yes         |
| An optimizing compiler may remove memset non-zero leaving data in memory | CWE-1330         | None                   | No          |
| Potential Negative Number Used as Index                                  | CWE-125, CWE-787 | Sans Top 25            | No          |
| Path Traversal                                                           | CWE-23           | OWASP:A01              | No          |
| Exposure of Private Personal Information to an Unauthorized Actor        | CWE-359          | OWASP:A01              | No          |
| Size Used as Index                                                       | CWE-125, CWE-787 | Sans Top 25            | Yes         |
| SQL Injection                                                            | CWE-89           | Sans Top 25, OWASP:A03 | No          |
| Server-Side Request Forgery (SSRF)                                       | CWE-918          | Sans Top 25, OWASP:A10 | No          |
| Inadequate Encryption Strength                                           | CWE-326          | OWASP:A02              | Yes         |
| Potential buffer overflow from usage of unsafe function                  | CWE-122          | None                   | Yes         |
| Use of Expired File Descriptor                                           | CWE-910          | None                   | No          |
| Use After Free                                                           | CWE-416          | Sans Top 25            | No          |
| User Controlled Pointer                                                  | CWE-1285         | None                   | No          |
| Authentication Bypass by Spoofing                                        | CWE-290          | OWASP:A07              | No          |
| Cross-site Scripting (XSS)                                               | CWE-79           | Sans Top 25, OWASP:A03 | No          |
| XML External Entity (XXE) Injection                                      | CWE-611          | OWASP:A05              | No          |
| XPath Injection                                                          | CWE-643          | OWASP:A03              | No          |
