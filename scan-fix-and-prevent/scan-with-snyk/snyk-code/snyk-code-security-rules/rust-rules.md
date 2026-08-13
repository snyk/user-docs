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
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2021 edition) category the rule belongs to, if any, and whether [SANS Top 25](https://www.sans.org/top25-software-errors/) includes it.
* **Autofixable**: Whether Snyk Agent Fix can fix the rule automatically. Snyk includes this column only for languages Snyk Agent Fix supports.

| Rule Name                                                   | CWEs             | Security Categories    | Autofixable |
| ----------------------------------------------------------- | ---------------- | ---------------------- | ----------- |
| Command Injection                                           | CWE-78           | SANS Top 25, OWASP:A03 | Yes         |
| Cross-site Scripting (XSS)                                  | CWE-79           | SANS Top 25, OWASP:A03 | Yes         |
| Hardcoded Secret                                            | CWE-547          | OWASP:A05              | Yes         |
| Inadequate Padding for Public Key Encryption                | CWE-326          | OWASP:A02              | Yes         |
| Insecure File Permissions                                   | CWE-732          | OWASP:A05              | Yes         |
| Observable Timing Discrepancy                               | CWE-208          | OWASP:A02              | Yes         |
| Open Redirect                                               | CWE-601          | OWASP:A01              | Yes         |
| Origin Validation Error                                     | CWE-346, CWE-942 | OWASP:A05              | Yes         |
| Path Traversal                                              | CWE-23           | OWASP:A01              | Yes         |
| SQL Injection                                               | CWE-89           | SANS Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                          | CWE-918          | SANS Top 25, OWASP:A10 | Yes         |
| Use of Hardcoded Passwords                                  | CWE-259, CWE-798 | SANS Top 25, OWASP:A07 | Yes         |
| Use of Insufficiently Random Values                         | CWE-330          | OWASP:A02              | Yes         |
| Use of Password Hash With Insufficient Computational Effort | CWE-916          | OWASP:A02              | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm            | CWE-327          | OWASP:A02              | Yes         |
