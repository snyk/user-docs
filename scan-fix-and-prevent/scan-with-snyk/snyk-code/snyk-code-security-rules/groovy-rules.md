---
description: Snyk Code security rules for Groovy
---

# Groovy rules

{% hint style="info" %}
Code analysis support for Groovy is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

Each rule includes the following information.

* **Rule name**: The Snyk name of the rule.
* **CWE(s)**: The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category to which the rule belongs to, if any, and if it is included in [CWE Top 25](https://cwe.mitre.org/top25/), plus any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) and [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule name                                                   | CWE(s)           | Security categories                             | Autofixable |
| ----------------------------------------------------------- | ---------------- | ----------------------------------------------- | ----------- |
| Code Injection                                              | CWE-94           | CWE Top 25, OWASP:A05:2025                      | Yes         |
| Command Injection                                           | CWE-78           | CWE Top 25, OWASP:A05:2025                      | Yes         |
| Deserialization of Untrusted Data                           | CWE-502          | CWE Top 25, OWASP:A08:2025                      | Yes         |
| Hardcoded Secret                                            | CWE-547          | OWASP:A02:2025                                  | Yes         |
| Open Redirect                                               | CWE-601          | OWASP:A01:2025                                  | Yes         |
| Path Traversal                                              | CWE-23           | OWASP:A01:2025                                  | Yes         |
| SQL Injection                                               | CWE-89           | CWE Top 25, OWASP:A05:2025                      | Yes         |
| Server-Side Request Forgery (SSRF)                          | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023 | Yes         |
| Use of Hardcoded Passwords                                  | CWE-798, CWE-259 | OWASP:A07:2025                                  | Yes         |
| Use of Hardcoded, Security-relevant Constants               | CWE-547          | OWASP:A02:2025                                  | Yes         |
| XML External Entity (XXE) Injection                         | CWE-611          | OWASP:A02:2025                                  | Yes         |
| Inadequate Padding for AES encryption                       | CWE-326          | OWASP:A04:2025                                  | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm            | CWE-327          | OWASP:A04:2025                                  | Yes         |
| Use of Password Hash With Insufficient Computational Effort | CWE-916          | OWASP:A04:2025                                  | Yes         |
| Improper Certificate Validation                             | CWE-295          | OWASP:A07:2025                                  | Yes         |
