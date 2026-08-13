---
description: Snyk Code security rules for Dart and Flutter
nav_context: agnostic
---

# Dart and Flutter rules

{% hint style="info" %}
Code analysis support for Dart is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2021 edition) category the rule belongs to, if any, and whether [SANS Top 25](https://www.sans.org/top25-software-errors/) includes it.
* **Autofixable**: Whether Snyk Agent Fix can fix the rule automatically. Snyk includes this column only for languages Snyk Agent Fix supports.

| Rule Name                                                    | CWEs             | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Clear Text Logging                                           | CWE-200, CWE-312 | OWASP:A01, OWASP:A04   | Yes         |
| Cleartext Transmission - HTTP Instead of HTTPS               | CWE-319          | OWASP:A02              | Yes         |
| Code Injection                                               | CWE-94           | OWASP:A03              | Yes         |
| File Access Enabled                                          | CWE-200          | OWASP:A01, OWASP:A04   | Yes         |
| Improper Certificate Validation - SSL Verification Bypass    | CWE-295          | OWASP:A07              | Yes         |
| Insecure JWT Verification Method                             | CWE-347          | OWASP:A02              | Yes         |
| Insecure Storage Shared Keystore                             | CWE-922          |                        | Yes         |
| Insecure Token Storage                                       | CWE-798          | SANS Top 25, OWASP:A07 | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A05              | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | Yes         |
| Use of Hardcoded Credentials                                 | CWE-798          | SANS Top 25, OWASP:A07 | Yes         |
| Use of Hardcoded Cryptographic Key                           | CWE-321          | OWASP:A02              | Yes         |
| Use of Insufficiently Random Values - Secrets                | CWE-330          | OWASP:A02              | Yes         |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A02              | Yes         |
