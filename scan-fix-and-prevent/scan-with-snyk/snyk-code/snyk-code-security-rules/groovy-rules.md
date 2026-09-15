---
description: Snyk Code security rules for Groovy
nav_context: agnostic
---

# Groovy rules

{% hint style="info" %}
Code analysis support for Groovy is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                                   | CWEs             | Security Categories                             |
| ----------------------------------------------------------- | ---------------- | ----------------------------------------------- |
| Code Injection                                              | CWE-94           | CWE Top 25, OWASP:A05:2025                      |
| Command Injection                                           | CWE-78           | CWE Top 25, OWASP:A05:2025                      |
| Deserialization of Untrusted Data                           | CWE-502          | CWE Top 25, OWASP:A08:2025                      |
| Hardcoded Secret                                            | CWE-547          | OWASP:A02:2025                                  |
| Open Redirect                                               | CWE-601          | OWASP:A01:2025                                  |
| Path Traversal                                              | CWE-23           | OWASP:A01:2025                                  |
| SQL Injection                                               | CWE-89           | CWE Top 25, OWASP:A05:2025                      |
| Server-Side Request Forgery (SSRF)                          | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023 |
| Use of Hardcoded Passwords                                  | CWE-798, CWE-259 | OWASP:A07:2025                                  |
| Use of Hardcoded, Security-relevant Constants               | CWE-547          | OWASP:A02:2025                                  |
| XML External Entity (XXE) Injection                         | CWE-611          | OWASP:A02:2025                                  |
| Inadequate Padding for AES encryption                       | CWE-326          | OWASP:A04:2025                                  |
| Use of a Broken or Risky Cryptographic Algorithm            | CWE-327          | OWASP:A04:2025                                  |
| Use of Password Hash With Insufficient Computational Effort | CWE-916          | OWASP:A04:2025                                  |
| Improper Certificate Validation                             | CWE-295          | OWASP:A07:2025                                  |
