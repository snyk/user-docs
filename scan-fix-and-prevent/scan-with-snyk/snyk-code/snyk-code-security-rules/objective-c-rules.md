---
description: Snyk Code security rules for Objective-C
---

# Objective-C rules

{% hint style="info" %}
Code analysis support for Objective-C is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s)**: The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories                                              | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------------------------------------------------- | ----------- |
| Clear Text Logging                                           | CWE-200, CWE-312 | CWE Top 25, OWASP:A01:2025, OWASP:A06:2025, OWASP-API:API10:2023 | Yes         |
| Client-Side Request Forgery (CSRF)                           | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023                  | Yes         |
| Code Injection                                               | CWE-94           | CWE Top 25, OWASP:A05:2025                                       | Yes         |
| Command Injection                                            | CWE-78           | CWE Top 25, OWASP:A05:2025                                       | Yes         |
| Device Authentication Bypass                                 | CWE-287          | OWASP:A07:2025                                                   | Yes         |
| Improper Certificate Validation                              | CWE-295          | OWASP:A07:2025                                                   | Yes         |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A04:2025                                                   | Yes         |
| Information Exposure                                         | CWE-200          | CWE Top 25, OWASP:A01:2025, OWASP-API:API10:2023                 | Yes         |
| Insecure Data Storage                                        | CWE-922          | OWASP:A01:2025                                                   | Yes         |
| Memory Corruption                                            | CWE-822          | None                                                             | Yes         |
| Path Traversal                                               | CWE-23           | OWASP:A01:2025                                                   | Yes         |
| SQL Injection                                                | CWE-89           | CWE Top 25, OWASP:A05:2025                                       | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A02:2025                                                   | Yes         |
| Use of Hardcoded Cryptographic Key                           | CWE-321          | OWASP:A04:2025                                                   | Yes         |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A04:2025                                                   | Yes         |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A04:2025                                                   | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A04:2025                                                   | Yes         |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A02:2025                                                   | Yes         |
