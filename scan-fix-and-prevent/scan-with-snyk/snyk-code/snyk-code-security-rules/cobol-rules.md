---
description: Snyk Code security rules for COBOL
nav_context: agnostic
---

# COBOL rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2021 edition) category the rule belongs to, if any, and whether [SANS Top 25](https://www.sans.org/top25-software-errors/) includes it.
* **Autofixable**: Whether Snyk Agent Fix can fix the rule automatically. Snyk includes this column only for languages Snyk Agent Fix supports.

| Rule Name                                           | CWEs             | Security Categories         | Autofixable |
| --------------------------------------------------- | ---------------- | --------------------------- | ----------- |
| Use of Hardcoded Cryptographic Key                  | CWE-321          | OWASP:A02:2021              | Yes         |
| Use of Hardcoded Cryptographic Initialization Value | CWE-321          | OWASP:A02:2021              | Yes         |
| No Dynamic SQL Clauses                              | CWE-89           | SANS Top 25, OWASP:A03:2021 | Yes         |
| Inadequate Encryption Strength - Small Key Size     | CWE-326          | OWASP:A02:2021              | Yes         |
| Weak Cryptographic Primitive                        | CWE-327          | OWASP:A02:2021              | Yes         |
| Clear Text Logging                                  | CWE-321          | OWASP:A02:2021              | Yes         |
| Hardcoded Secret                                    | CWE-547          | OWASP:A02:2021              | Yes         |
| Injection on Accept                                 | CWE-20           | SANS Top 25                 | Yes         |
| Insecure Debug Features Enabled                     | CWE-489, CWE-215 | OWASP:A05:2021              | Yes         |
| Insecure Data Transmission                          | CWE-319          | OWASP:A02:2021              | Yes         |
| SQL SELECT statement without WHERE clause           | CWE-668          | OWASP:A01:2021              | Yes         |
| Multiple CICS HANDLE ABEND Declarations             | CWE-755          | OWASP:A05:2021              | Yes         |
| Missing SQL Communication Area (SQLCA)              | CWE-391          | OWASP:A05:2021              | Yes         |
| Ignored Error Condition                             | CWE-391          | OWASP:A05:2021              | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm    | CWE-327          | OWASP:A02:2021              | Yes         |
