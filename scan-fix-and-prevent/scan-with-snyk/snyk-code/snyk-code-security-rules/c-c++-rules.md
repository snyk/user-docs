---
description: Snyk Code security rules for C and C++
nav_context: agnostic
---

# C++ rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2021 edition) category the rule belongs to, if any, and whether [SANS Top 25](https://www.sans.org/top25-software-errors/) includes it.
* **Autofixable**: Whether Snyk Agent Fix can fix the rule automatically. Snyk includes this column only for languages Snyk Agent Fix supports.

| Rule Name                                                                | CWEs             | Security Categories    | Autofixable |
| ------------------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Memory Allocation Of String Length                                       | CWE-170          | None                   | Yes         |
| Insecure Anonymous LDAP Binding                                          | CWE-287          | SANS Top 25, OWASP:A07 | Yes         |
| Buffer Overflow                                                          | CWE-122          | None                   | Yes         |
| Division By Zero                                                         | CWE-369          | None                   | Yes         |
| Missing Release of File Descriptor or Handle after Effective Lifetime    | CWE-775          | None                   | Yes         |
| Command Injection                                                        | CWE-78           | SANS Top 25, OWASP:A03 | Yes         |
| Dereference of a NULL Pointer                                            | CWE-476          | SANS Top 25            | Yes         |
| Double Free                                                              | CWE-415          | None                   | Yes         |
| Use of Externally-Controlled Format String                               | CWE-134          | None                   | Yes         |
| Use of Hardcoded Cryptographic Key                                       | CWE-321          | OWASP:A02              | Yes         |
| Improper Null Termination                                                | CWE-170          | None                   | Yes         |
| Use of Password Hash With Insufficient Computational Effort              | CWE-916          | OWASP:A02              | Yes         |
| Integer Overflow                                                         | CWE-190          | SANS Top 25            | Yes         |
| LDAP Injection                                                           | CWE-90           | OWASP:A03              | Yes         |
| Missing Release of Memory after Effective Lifetime                       | CWE-401          | None                   | Yes         |
| An optimizing compiler may remove memset non-zero leaving data in memory | CWE-1330         | None                   | Yes         |
| Potential Negative Number Used as Index                                  | CWE-125, CWE-787 | SANS Top 25            | Yes         |
| Path Traversal                                                           | CWE-23           | OWASP:A01              | Yes         |
| Exposure of Private Personal Information to an Unauthorized Actor        | CWE-359          | OWASP:A01              | Yes         |
| Size Used as Index                                                       | CWE-125, CWE-787 | SANS Top 25            | Yes         |
| SQL Injection                                                            | CWE-89           | SANS Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                                       | CWE-918          | SANS Top 25, OWASP:A10 | Yes         |
| Inadequate Encryption Strength                                           | CWE-326          | OWASP:A02              | Yes         |
| Potential buffer overflow from usage of unsafe function                  | CWE-122          | None                   | Yes         |
| Use of Expired File Descriptor                                           | CWE-910          | None                   | Yes         |
| Use After Free                                                           | CWE-416          | SANS Top 25            | Yes         |
| User Controlled Pointer                                                  | CWE-1285         | None                   | Yes         |
| Authentication Bypass by Spoofing                                        | CWE-290          | OWASP:A07              | Yes         |
| Cross-site Scripting (XSS)                                               | CWE-79           | SANS Top 25, OWASP:A03 | Yes         |
| XML External Entity (XXE) Injection                                      | CWE-611          | OWASP:A05              | Yes         |
| XPath Injection                                                          | CWE-643          | OWASP:A03              | Yes         |
