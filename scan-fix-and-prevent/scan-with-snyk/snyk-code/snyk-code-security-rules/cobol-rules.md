---
description: Snyk Code security rules for COBOL
nav_context: agnostic
---

# COBOL rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).

| Rule Name                                           | CWE(s)           | Security Categories         |
| --------------------------------------------------- | ---------------- | --------------------------- |
| Use of Hardcoded Cryptographic Key                  | CWE-321          | OWASP:A02:2021              |
| Use of Hardcoded Cryptographic Initialization Value | CWE-321          | OWASP:A02:2021              |
| No Dynamic SQL Clauses                              | CWE-89           | Sans Top 25, OWASP:A03:2021 |
| Inadequate Encryption Strength - Small Key Size     | CWE-326          | OWASP:A02:2021              |
| Weak Cryptographic Primitive                        | CWE-327          | OWASP:A02:2021              |
| Clear Text Logging                                  | CWE-321          | OWASP:A02:2021              |
| Hardcoded Secret                                    | CWE-547          | OWASP:A02:2021              |
| Injection on Accept                                 | CWE-20           | SANS Top 25                 |
| Insecure Debug Features Enabled                     | CWE-489, CWE-215 | OWASP:A05:2021              |
| Insecure Data Transmission                          | CWE-319          | OWASP:A02:2021              |
| SQL SELECT statement without WHERE clause           | CWE-668          | OWASP:A01:2021              |
| Multiple CICS HANDLE ABEND Declarations             | CWE-755          | OWASP:A05:2021              |
| Missing SQL Communication Area (SQLCA)              | CWE-391          | OWASP:A05:2021              |
| Ignored Error Condition                             | CWE-391          | OWASP:A05:2021              |
| Use of a Broken or Risky Cryptographic Algorithm    | CWE-327          | OWASP:A02:2021              |
