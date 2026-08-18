---
description: Snyk Code security rules for Rust
nav_context: agnostic
---

# Rust rules

{% hint style="info" %}
Code analysis support for Rust is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                                   | CWEs             | Security Categories                                 |
| ----------------------------------------------------------- | ---------------- | --------------------------------------------------- |
| Command Injection                                           | CWE-78           | CWE Top 25, OWASP:A05:2025                          |
| Cross-site Scripting (XSS)                                  | CWE-79           | CWE Top 25, OWASP:A05:2025                          |
| Hardcoded Secret                                            | CWE-547          | OWASP:A02:2025                                      |
| Inadequate Padding for Public Key Encryption                | CWE-326          | OWASP:A04:2025                                      |
| Insecure File Permissions                                   | CWE-732          | OWASP:A01:2025                                      |
| Observable Timing Discrepancy                               | CWE-208          | None                                                |
| Open Redirect                                               | CWE-601          | OWASP:A01:2025                                      |
| Origin Validation Error                                     | CWE-346, CWE-942 | OWASP:A02:2025, OWASP:A07:2025, OWASP-API:API8:2023 |
| Path Traversal                                              | CWE-23           | OWASP:A01:2025                                      |
| SQL Injection                                               | CWE-89           | CWE Top 25, OWASP:A05:2025                          |
| Server-Side Request Forgery (SSRF)                          | CWE-918          | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023     |
| Use of Hardcoded Passwords                                  | CWE-259, CWE-798 | OWASP:A07:2025                                      |
| Use of Insufficiently Random Values                         | CWE-330          | OWASP:A04:2025                                      |
| Use of Password Hash With Insufficient Computational Effort | CWE-916          | OWASP:A04:2025                                      |
| Use of a Broken or Risky Cryptographic Algorithm            | CWE-327          | OWASP:A04:2025                                      |
