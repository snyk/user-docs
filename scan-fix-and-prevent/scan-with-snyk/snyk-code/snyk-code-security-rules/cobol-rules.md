---
description: Snyk Code security rules for COBOL
nav_context: agnostic
---

# COBOL rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                           | CWEs             | Security Categories                                                    |
| --------------------------------------------------- | ---------------- | ---------------------------------------------------------------------- |
| Use of Hardcoded Cryptographic Key                  | CWE-321          | OWASP:A04:2025                                                         |
| Use of Hardcoded Cryptographic Initialization Value | CWE-321          | OWASP:A04:2025                                                         |
| No Dynamic SQL Clauses                              | CWE-89           | CWE Top 25, OWASP:A05:2025                                             |
| Inadequate Encryption Strength - Small Key Size     | CWE-326          | OWASP:A04:2025                                                         |
| Weak Cryptographic Primitive                        | CWE-327          | OWASP:A04:2025                                                         |
| Clear Text Logging                                  | CWE-321          | OWASP:A04:2025                                                         |
| Hardcoded Secret                                    | CWE-547          | OWASP:A02:2025                                                         |
| Injection on Accept                                 | CWE-20           | CWE Top 25, OWASP:A05:2025, OWASP-API:API10:2023, OWASP-Mobile:M4:2024 |
| Insecure Debug Features Enabled                     | CWE-489, CWE-215 | OWASP:A02:2025, OWASP:A10:2025                                         |
| Insecure Data Transmission                          | CWE-319          | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023              |
| SQL SELECT statement without WHERE clause           | CWE-668          | OWASP:A01:2025                                                         |
| Multiple CICS HANDLE ABEND Declarations             | CWE-755          | OWASP:A10:2025                                                         |
| Missing SQL Communication Area (SQLCA)              | CWE-391          | OWASP:A10:2025                                                         |
| Ignored Error Condition                             | CWE-391          | OWASP:A10:2025                                                         |
| Use of a Broken or Risky Cryptographic Algorithm    | CWE-327          | OWASP:A04:2025                                                         |
