---
description: Snyk Code security rules for Rust
---

# Rust rules

{% hint style="info" %}
Code analysis support for Rust is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s)**: The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category to which the rule belongs to, if any, and if it is included in [CWE Top 25](https://cwe.mitre.org/top25/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                   | CWE(s)           | Security Categories            | Autofixable |
| ----------------------------------------------------------- | ---------------- | ------------------------------ | ----------- |
| Command Injection                                           | CWE-78           | CWE Top 25, OWASP:A05:2025     | Yes         |
| Cross-site Scripting (XSS)                                  | CWE-79           | CWE Top 25, OWASP:A05:2025     | Yes         |
| Hardcoded Secret                                            | CWE-547          | OWASP:A02:2025                 | Yes         |
| Inadequate Padding for Public Key Encryption                | CWE-326          | OWASP:A04:2025                 | Yes         |
| Insecure File Permissions                                   | CWE-732          | OWASP:A01:2025                 | Yes         |
| Observable Timing Discrepancy                               | CWE-208          | None                           | Yes         |
| Open Redirect                                               | CWE-601          | OWASP:A01:2025                 | Yes         |
| Origin Validation Error                                     | CWE-346, CWE-942 | OWASP:A02:2025, OWASP:A07:2025 | Yes         |
| Path Traversal                                              | CWE-23           | OWASP:A01:2025                 | Yes         |
| SQL Injection                                               | CWE-89           | CWE Top 25, OWASP:A05:2025     | Yes         |
| Server-Side Request Forgery (SSRF)                          | CWE-918          | CWE Top 25, OWASP:A01:2025     | Yes         |
| Use of Hardcoded Passwords                                  | CWE-259, CWE-798 | OWASP:A07:2025                 | Yes         |
| Use of Insufficiently Random Values                         | CWE-330          | OWASP:A04:2025                 | Yes         |
| Use of Password Hash With Insufficient Computational Effort | CWE-916          | OWASP:A04:2025                 | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm            | CWE-327          | OWASP:A04:2025                 | Yes         |
