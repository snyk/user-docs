---
description: Snyk Code security rules for COBOL
---

# COBOL rules

{% hint style="info" %}
Code analysis support for COBOL is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/2025/)(2025 edition) category to which the rule belongs to, if any, and if it is included in [CWE Top 25](https://cwe.mitre.org/top25/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                           | CWE(s)           | Security Categories            | Autofixable |
| --------------------------------------------------- | ---------------- | ------------------------------ | ----------- |
| Use of Hardcoded Cryptographic Key                  | CWE-321          | OWASP:A04:2025                 | Yes         |
| Use of Hardcoded Cryptographic Initialization Value | CWE-321          | OWASP:A04:2025                 | Yes         |
| No Dynamic SQL Clauses                              | CWE-89           | CWE Top 25, OWASP:A05:2025     | Yes         |
| Inadequate Encryption Strength - Small Key Size     | CWE-326          | OWASP:A04:2025                 | Yes         |
| Weak Cryptographic Primitive                        | CWE-327          | OWASP:A04:2025                 | Yes         |
| Clear Text Logging                                  | CWE-321          | OWASP:A04:2025                 | Yes         |
| Hardcoded Secret                                    | CWE-547          | OWASP:A02:2025                 | Yes         |
| Injection on Accept                                 | CWE-20           | CWE Top 25, OWASP:A05:2025     | Yes         |
| Insecure Debug Features Enabled                     | CWE-489, CWE-215 | OWASP:A02:2025, OWASP:A10:2025 | Yes         |
| Insecure Data Transmission                          | CWE-319          | OWASP:A04:2025                 | Yes         |
| SQL SELECT statement without WHERE clause           | CWE-668          | OWASP:A01:2025                 | Yes         |
| Multiple CICS HANDLE ABEND Declarations             | CWE-755          | OWASP:A10:2025                 | Yes         |
| Missing SQL Communication Area (SQLCA)              | CWE-391          | OWASP:A10:2025                 | Yes         |
| Ignored Error Condition                             | CWE-391          | OWASP:A10:2025                 | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm    | CWE-327          | OWASP:A04:2025                 | Yes         |
