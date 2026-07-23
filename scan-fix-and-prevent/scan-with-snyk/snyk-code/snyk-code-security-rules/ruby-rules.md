---
description: Snyk Code security rules for Ruby
---

# Ruby rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/2025/)(2025 edition) category to which the rule belongs to, if any, and if it is included in [CWE Top 25](https://cwe.mitre.org/top25/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                                      | CWE(s)                                                      | Security Categories                                                        | Autofixable |
| ------------------------------------------------------------------------------ | ----------------------------------------------------------- | -------------------------------------------------------------------------- | ----------- |
| Code Injection                                                                 | CWE-94                                                      | CWE Top 25, OWASP:A05:2025                                                 | Yes         |
| Command Injection                                                              | CWE-78                                                      | CWE Top 25, OWASP:A05:2025                                                 | Yes         |
| Remote Code Execution via Endpoint                                             | CWE-94                                                      | CWE Top 25, OWASP:A05:2025                                                 | Yes         |
| Deserialization of Untrusted Data                                              | CWE-502                                                     | CWE Top 25, OWASP:A08:2025                                                 | Yes         |
| Use of Hardcoded Credentials                                                   | CWE-798, CWE-259                                            | OWASP:A07:2025                                                             | Yes         |
| Use of Hardcoded Cryptographic Key                                             | CWE-321                                                     | OWASP:A04:2025                                                             | Yes         |
| Hardcoded Secret                                                               | CWE-547                                                     | OWASP:A02:2025                                                             | Yes         |
| Use of Hardcoded Passwords                                                     | CWE-798, CWE-259                                            | OWASP:A07:2025                                                             | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm                               | CWE-327                                                     | OWASP:A04:2025                                                             | Yes         |
| Use of Password Hash With Insufficient Computational Effort                    | CWE-916                                                     | OWASP:A04:2025                                                             | Yes         |
| Use of Insufficiently Random Values                                            | CWE-330                                                     | OWASP:A04:2025                                                             | Yes         |
| Sinatra Protection Layers Disabled                                             | CWE-16, CWE-352, CWE-79, CWE-693, CWE-1021, CWE-35, CWE-348 | CWE Top 25, OWASP:A01:2025, OWASP:A02:2025, OWASP:A05:2025, OWASP:A06:2025 | Yes         |
| Insecure Data Transmission                                                     | CWE-319                                                     | OWASP:A04:2025                                                             | Yes         |
| Improper Input Validation                                                      | CWE-20                                                      | CWE Top 25, OWASP:A05:2025                                                 | Yes         |
| Improperly Controlled Modification of Dynamically-Determined Object Attributes | CWE-915                                                     | OWASP:A08:2025                                                             | Yes         |
| Selection of Less-Secure Algorithm During Negotiation (Force SSL)              | CWE-311, CWE-757                                            | OWASP:A04:2025, OWASP:A06:2025                                             | Yes         |
| Open Redirect                                                                  | CWE-601                                                     | OWASP:A01:2025                                                             | Yes         |
| Path Traversal                                                                 | CWE-23                                                      | OWASP:A01:2025                                                             | Yes         |
| Unsafe Reflection                                                              | CWE-470                                                     | OWASP:A05:2025                                                             | Yes         |
| Improper Certificate Validation                                                | CWE-295                                                     | OWASP:A07:2025                                                             | Yes         |
| Information Exposure                                                           | CWE-200                                                     | CWE Top 25, OWASP:A01:2025                                                 | Yes         |
| Session Manipulation                                                           | CWE-285                                                     | OWASP:A01:2025                                                             | Yes         |
| SQL Injection                                                                  | CWE-89                                                      | CWE Top 25, OWASP:A05:2025                                                 | Yes         |
| Improper Neutralization of Directives in Statically Saved Code                 | CWE-96                                                      | OWASP:A05:2025                                                             | Yes         |
| No Weak Password Requirements                                                  | CWE-521                                                     | OWASP:A07:2025                                                             | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                                       | CWE-1004                                                    | OWASP:A02:2025                                                             | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute                   | CWE-614                                                     | OWASP:A02:2025                                                             | Yes         |
| Incorrect regular expression for validating values                             | CWE-1286                                                    | None                                                                       | Yes         |
| Cross-site Scripting (XSS)                                                     | CWE-79                                                      | CWE Top 25, OWASP:A05:2025                                                 | Yes         |
| XML External Entity (XXE) Injection                                            | CWE-611                                                     | OWASP:A02:2025                                                             | Yes         |
| XPath Injection                                                                | CWE-643                                                     | OWASP:A05:2025                                                             | Yes         |
| Regular Expression Denial of Service (ReDoS)                                   | CWE-400                                                     | None                                                                       | Yes         |
