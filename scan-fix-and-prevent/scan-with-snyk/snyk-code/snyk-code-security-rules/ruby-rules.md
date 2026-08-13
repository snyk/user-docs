---
description: Snyk Code security rules for Ruby
nav_context: agnostic
---

# Ruby rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                                                      | CWEs                                                        | Security Categories                                                                             |
| ------------------------------------------------------------------------------ | ----------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| Code Injection                                                                 | CWE-94                                                      | CWE Top 25, OWASP:A05:2025                                                                      |
| Command Injection                                                              | CWE-78                                                      | CWE Top 25, OWASP:A05:2025                                                                      |
| Remote Code Execution via Endpoint                                             | CWE-94                                                      | CWE Top 25, OWASP:A05:2025                                                                      |
| Deserialization of Untrusted Data                                              | CWE-502                                                     | CWE Top 25, OWASP:A08:2025                                                                      |
| Use of Hardcoded Credentials                                                   | CWE-798, CWE-259                                            | OWASP:A07:2025                                                                                  |
| Use of Hardcoded Cryptographic Key                                             | CWE-321                                                     | OWASP:A04:2025                                                                                  |
| Hardcoded Secret                                                               | CWE-547                                                     | OWASP:A02:2025                                                                                  |
| Use of Hardcoded Passwords                                                     | CWE-798, CWE-259                                            | OWASP:A07:2025                                                                                  |
| Use of a Broken or Risky Cryptographic Algorithm                               | CWE-327                                                     | OWASP:A04:2025                                                                                  |
| Use of Password Hash With Insufficient Computational Effort                    | CWE-916                                                     | OWASP:A04:2025                                                                                  |
| Use of Insufficiently Random Values                                            | CWE-330                                                     | OWASP:A04:2025                                                                                  |
| Sinatra Protection Layers Disabled                                             | CWE-16, CWE-352, CWE-79, CWE-693, CWE-1021, CWE-35, CWE-348 | CWE Top 25, OWASP:A01:2025, OWASP:A02:2025, OWASP:A05:2025, OWASP:A06:2025, OWASP-API:API8:2023 |
| Insecure Data Transmission                                                     | CWE-319                                                     | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023                                       |
| Improper Input Validation                                                      | CWE-20                                                      | CWE Top 25, OWASP:A05:2025, OWASP-API:API10:2023, OWASP-Mobile:M4:2024                          |
| Improperly Controlled Modification of Dynamically-Determined Object Attributes | CWE-915                                                     | OWASP:A08:2025, OWASP-API:API3:2023                                                             |
| Selection of Less-Secure Algorithm During Negotiation (Force SSL)              | CWE-311, CWE-757                                            | OWASP:A04:2025, OWASP:A06:2025                                                                  |
| Open Redirect                                                                  | CWE-601                                                     | OWASP:A01:2025                                                                                  |
| Path Traversal                                                                 | CWE-23                                                      | OWASP:A01:2025                                                                                  |
| Unsafe Reflection                                                              | CWE-470                                                     | OWASP:A05:2025                                                                                  |
| Improper Certificate Validation                                                | CWE-295                                                     | OWASP:A07:2025                                                                                  |
| Information Exposure                                                           | CWE-200                                                     | CWE Top 25, OWASP:A01:2025, OWASP-API:API10:2023                                                |
| Session Manipulation                                                           | CWE-285                                                     | OWASP:A01:2025, OWASP-API:API1:2023, OWASP-API:API5:2023                                        |
| SQL Injection                                                                  | CWE-89                                                      | CWE Top 25, OWASP:A05:2025                                                                      |
| Improper Neutralization of Directives in Statically Saved Code                 | CWE-96                                                      | OWASP:A05:2025                                                                                  |
| No Weak Password Requirements                                                  | CWE-521                                                     | OWASP:A07:2025                                                                                  |
| Sensitive Cookie Without 'HttpOnly' Flag                                       | CWE-1004                                                    | OWASP:A02:2025                                                                                  |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute                   | CWE-614                                                     | OWASP:A02:2025                                                                                  |
| Incorrect regular expression for validating values                             | CWE-1286                                                    | None                                                                                            |
| Cross-site Scripting (XSS)                                                     | CWE-79                                                      | CWE Top 25, OWASP:A05:2025                                                                      |
| XML External Entity (XXE) Injection                                            | CWE-611                                                     | OWASP:A02:2025                                                                                  |
| XPath Injection                                                                | CWE-643                                                     | OWASP:A05:2025                                                                                  |
| Regular Expression Denial of Service (ReDoS)                                   | CWE-400                                                     | OWASP-API:API4:2023                                                                             |
