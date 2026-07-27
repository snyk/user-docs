---
description: Snyk Code security rules for Groovy
nav_context: agnostic
---

# Groovy rules

{% hint style="info" %}
Code analysis support for Groovy is in Early Access and is available only with Enterprise plans. To enable the feature, see [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

Each rule includes the following information.

* **Rule name**: The Snyk name of the rule.
* **CWE(s)**: The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule name                                                   | CWE(s)           | Security categories    | Autofixable |
| ----------------------------------------------------------- | ---------------- | ---------------------- | ----------- |
| Code Injection                                              | CWE-94           | OWASP:A03              | Yes         |
| Command Injection                                           | CWE-78           | Sans Top 25, OWASP:A03 | Yes         |
| Deserialization of Untrusted Data                           | CWE-502          | OWASP:A08              | Yes         |
| Hardcoded Secret                                            | CWE-547          | OWASP:A05              | Yes         |
| Open Redirect                                               | CWE-601          | OWASP:A01              | Yes         |
| Path Traversal                                              | CWE-23           | OWASP:A01              | Yes         |
| SQL Injection                                               | CWE-89           | Sans Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                          | CWE-918          | Sans Top 25, OWASP:A10 | Yes         |
| Use of Hardcoded Passwords                                  | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | Yes         |
| Use of Hardcoded, Security-relevant Constants               | CWE-547          | OWASP:A05              | Yes         |
| XML External Entity (XXE) Injection                         | CWE-611          | OWASP:A05              | Yes         |
| Inadequate Padding for AES encryption                       | CWE-326          | OWASP:A02              | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm            | CWE-327          | OWASP:A02              | Yes         |
| Use of Password Hash With Insufficient Computational Effort | CWE-916          | OWASP:A02              | Yes         |
| Improper Certificate Validation                             | CWE-295          | OWASP:A07              | Yes         |
