---
description: Snyk Code security rules for Dart and Flutter
---

# Dart and Flutter rules

{% hint style="info" %}
Code analysis support for Dart is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s)**: The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category to which the rule belongs to, if any, and if it is included in [CWE Top 25](https://cwe.mitre.org/top25/), plus any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) and [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories                                              | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------------------------------------------------- | ----------- |
| Clear Text Logging                                           | CWE-200, CWE-312 | CWE Top 25, OWASP:A01:2025, OWASP:A06:2025, OWASP-API:API10:2023 | Yes         |
| Cleartext Transmission - HTTP Instead of HTTPS               | CWE-319          | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023        | Yes         |
| Code Injection                                               | CWE-94           | CWE Top 25, OWASP:A05:2025                                       | Yes         |
| File Access Enabled                                          | CWE-200          | CWE Top 25, OWASP:A01:2025, OWASP-API:API10:2023                 | Yes         |
| Improper Certificate Validation - SSL Verification Bypass    | CWE-295          | OWASP:A07:2025                                                   | Yes         |
| Insecure JWT Verification Method                             | CWE-347          | OWASP:A04:2025                                                   | Yes         |
| Insecure Storage Shared Keystore                             | CWE-922          | OWASP:A01:2025                                                   | Yes         |
| Insecure Token Storage                                       | CWE-798          | OWASP:A07:2025                                                   | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A02:2025                                                   | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A02:2025                                                   | Yes         |
| Use of Hardcoded Credentials                                 | CWE-798          | OWASP:A07:2025                                                   | Yes         |
| Use of Hardcoded Cryptographic Key                           | CWE-321          | OWASP:A04:2025                                                   | Yes         |
| Use of Insufficiently Random Values - Secrets                | CWE-330          | OWASP:A04:2025                                                   | Yes         |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A04:2025                                                   | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A04:2025                                                   | Yes         |
