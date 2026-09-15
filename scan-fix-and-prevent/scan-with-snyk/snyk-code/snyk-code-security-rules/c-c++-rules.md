---
description: Snyk Code security rules for C and C++
nav_context: agnostic
---

# C++ rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                                                | CWEs             | Security Categories                             |
| ------------------------------------------------------------------------ | ---------------- | ----------------------------------------------- |
| Memory Allocation Of String Length                                       | CWE-170          | None                                            |
| Insecure Anonymous LDAP Binding                                          | CWE-287          | OWASP:A07:2025                                  |
| Buffer Overflow                                                          | CWE-122          | CWE Top 25                                      |
| Division By Zero                                                         | CWE-369          | OWASP:A10:2025                                  |
| Missing Release of File Descriptor or Handle after Effective Lifetime    | CWE-775          | None                                            |
| Command Injection                                                        | CWE-78           | CWE Top 25, OWASP:A05:2025                      |
| Dereference of a NULL Pointer                                            | CWE-476          | CWE Top 25, OWASP:A10:2025                      |
| Double Free                                                              | CWE-415          | None                                            |
| Use of Externally-Controlled Format String                               | CWE-134          | None                                            |
| Use of Hardcoded Cryptographic Key                                       | CWE-321          | OWASP:A04:2025                                  |
| Improper Null Termination                                                | CWE-170          | None                                            |
| Use of Password Hash With Insufficient Computational Effort              | CWE-916          | OWASP:A04:2025                                  |
| Integer Overflow                                                         | CWE-190          | None                                            |
| LDAP Injection                                                           | CWE-90           | OWASP:A05:2025                                  |
| Missing Release of Memory after Effective Lifetime                       | CWE-401          | None                                            |
| An optimizing compiler may remove memset non-zero leaving data in memory | CWE-1330         | None                                            |
| Potential Negative Number Used as Index                                  | CWE-125, CWE-787 | CWE Top 25                                      |
| Path Traversal                                                           | CWE-23           | OWASP:A01:2025                                  |
| Exposure of Private Personal Information to an Unauthorized Actor        | CWE-359          | OWASP:A01:2025                                  |
| Size Used as Index                                                       | CWE-125, CWE-787 | CWE Top 25                                      |
| SQL Injection                                                            | CWE-89           | CWE Top 25, OWASP:A05:2025                      |
| Server-Side Request Forgery (SSRF)                                       | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023 |
| Inadequate Encryption Strength                                           | CWE-326          | OWASP:A04:2025                                  |
| Potential buffer overflow from usage of unsafe function                  | CWE-122          | CWE Top 25                                      |
| Use of Expired File Descriptor                                           | CWE-910          | None                                            |
| Use After Free                                                           | CWE-416          | CWE Top 25                                      |
| User Controlled Pointer                                                  | CWE-1285         | None                                            |
| Authentication Bypass by Spoofing                                        | CWE-290          | OWASP:A07:2025                                  |
| Cross-site Scripting (XSS)                                               | CWE-79           | CWE Top 25, OWASP:A05:2025                      |
| XML External Entity (XXE) Injection                                      | CWE-611          | OWASP:A02:2025                                  |
| XPath Injection                                                          | CWE-643          | OWASP:A05:2025                                  |
