---
description: Snyk Code security rules for Objective-C
nav_context: agnostic
---

# Objective-C rules

{% hint style="info" %}
Code analysis support for Objective-C is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s)**: The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Clear Text Logging                                           | CWE-200, CWE-312 | OWASP:A01, OWASP:A04   | Yes         |
| Client-Side Request Forgery (CSRF)                           | CWE-918          | Sans Top 25, OWASP:A10 | Yes         |
| Code Injection                                               | CWE-94           | OWASP:A03              | Yes         |
| Command Injection                                            | CWE-78           | Sans Top 25, OWASP:A03 | Yes         |
| Device Authentication Bypass                                 | CWE-287          | OWASP:A07              | Yes         |
| Improper Certificate Validation                              | CWE-295          | OWASP:A07              | Yes         |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A02              | Yes         |
| Information Exposure                                         | CWE-200          | OWASP:A01, OWASP:A04   | Yes         |
| Insecure Data Storage                                        | CWE-922          |                        | Yes         |
| Memory Corruption                                            | CWE-822          | OWASP:A03              | Yes         |
| Path Traversal                                               | CWE-23           | OWASP:A01              | Yes         |
| SQL Injection                                                | CWE-89           | Sans Top 25, OWASP:A03 | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | Yes         |
| Use of Hardcoded Cryptographic Key                           | CWE-321          | OWASP:A02              | Yes         |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A02              | Yes         |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A02              | Yes         |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A05              | Yes         |
