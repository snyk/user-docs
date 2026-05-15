# COBOL rules

{% hint style="info" %}
Code analysis support for COBOL is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-platform-administration/snyk-preview).
{% endhint %}

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                           | CWE(s)           | Security Categories         | Autofixable |
| --------------------------------------------------- | ---------------- | --------------------------- | ----------- |
| Use of Hardcoded Cryptographic Key                  | CWE-321          | OWASP:A02:2021              | No          |
| Use of Hardcoded Cryptographic Initialization Value | CWE-321          | OWASP:A02:2021              | No          |
| No Dynamic SQL Clauses                              | CWE-89           | Sans Top 25, OWASP:A03:2021 | No          |
| Inadequate Encryption Strength - Small Key Size     | CWE-326          | OWASP:A02:2021              | No          |
| Weak Cryptographic Primitive                        | CWE-327          | OWASP:A02:2021              | No          |
| Clear Text Logging                                  | CWE-321          | OWASP:A02:2021              | No          |
| Hardcoded Secret                                    | CWE-547          | OWASP:A02:2021              | No          |
| Injection on Accept                                 | CWE-20           | SANS Top 25                 | No          |
| Insecure Debug Features Enabled                     | CWE-489, CWE-215 | OWASP:A05:2021              | No          |
| Insecure Data Transmission                          | CWE-319          | OWASP:A02:2021              | No          |
| SQL SELECT statement without WHERE clause           | CWE-668          | OWASP:A01:2021              | No          |
| Multiple CICS HANDLE ABEND Declarations             | CWE-755          | OWASP:A05:2021              | No          |
| Missing SQL Communication Area (SQLCA)              | CWE-391          | OWASP:A05:2021              | No          |
| Ignored Error Condition                             | CWE-391          | OWASP:A05:2021              | No          |
| Use of a Broken or Risky Cryptographic Algorithm    | CWE-327          | OWASP:A02:2021              | No          |
