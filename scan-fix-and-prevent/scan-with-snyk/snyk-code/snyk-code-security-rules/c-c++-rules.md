---
description: Snyk Code security rules for C and C++
---

# C++ rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                                | CWE(s)           | Security Categories                             | Autofixable |
| ------------------------------------------------------------------------ | ---------------- | ----------------------------------------------- | ----------- |
| Memory Allocation Of String Length                                       | CWE-170          | None                                            | Yes         |
| Insecure Anonymous LDAP Binding                                          | CWE-287          | OWASP:A07:2025                                  | Yes         |
| Buffer Overflow                                                          | CWE-122          | CWE Top 25                                      | Yes         |
| Division By Zero                                                         | CWE-369          | OWASP:A10:2025                                  | Yes         |
| Missing Release of File Descriptor or Handle after Effective Lifetime    | CWE-775          | None                                            | Yes         |
| Command Injection                                                        | CWE-78           | CWE Top 25, OWASP:A05:2025                      | Yes         |
| Dereference of a NULL Pointer                                            | CWE-476          | CWE Top 25, OWASP:A10:2025                      | Yes         |
| Double Free                                                              | CWE-415          | None                                            | Yes         |
| Use of Externally-Controlled Format String                               | CWE-134          | None                                            | Yes         |
| Use of Hardcoded Cryptographic Key                                       | CWE-321          | OWASP:A04:2025                                  | Yes         |
| Improper Null Termination                                                | CWE-170          | None                                            | Yes         |
| Use of Password Hash With Insufficient Computational Effort              | CWE-916          | OWASP:A04:2025                                  | Yes         |
| Integer Overflow                                                         | CWE-190          | None                                            | Yes         |
| LDAP Injection                                                           | CWE-90           | OWASP:A05:2025                                  | Yes         |
| Missing Release of Memory after Effective Lifetime                       | CWE-401          | None                                            | Yes         |
| An optimizing compiler may remove memset non-zero leaving data in memory | CWE-1330         | None                                            | Yes         |
| Potential Negative Number Used as Index                                  | CWE-125, CWE-787 | CWE Top 25                                      | Yes         |
| Path Traversal                                                           | CWE-23           | OWASP:A01:2025                                  | Yes         |
| Exposure of Private Personal Information to an Unauthorized Actor        | CWE-359          | OWASP:A01:2025                                  | Yes         |
| Size Used as Index                                                       | CWE-125, CWE-787 | CWE Top 25                                      | Yes         |
| SQL Injection                                                            | CWE-89           | CWE Top 25, OWASP:A05:2025                      | Yes         |
| Server-Side Request Forgery (SSRF)                                       | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023 | Yes         |
| Inadequate Encryption Strength                                           | CWE-326          | OWASP:A04:2025                                  | Yes         |
| Potential buffer overflow from usage of unsafe function                  | CWE-122          | CWE Top 25                                      | Yes         |
| Use of Expired File Descriptor                                           | CWE-910          | None                                            | Yes         |
| Use After Free                                                           | CWE-416          | CWE Top 25                                      | Yes         |
| User Controlled Pointer                                                  | CWE-1285         | None                                            | Yes         |
| Authentication Bypass by Spoofing                                        | CWE-290          | OWASP:A07:2025                                  | Yes         |
| Cross-site Scripting (XSS)                                               | CWE-79           | CWE Top 25, OWASP:A05:2025                      | Yes         |
| XML External Entity (XXE) Injection                                      | CWE-611          | OWASP:A02:2025                                  | Yes         |
| XPath Injection                                                          | CWE-643          | OWASP:A05:2025                                  | Yes         |
