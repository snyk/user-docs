---
description: Snyk Code security rules for Ruby
nav_context: agnostic
---

# Ruby rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2021 edition) category the rule belongs to, if any, and whether [SANS Top 25](https://www.sans.org/top25-software-errors/) includes it.
* **Autofixable**: Whether Snyk Agent Fix can fix the rule automatically. Snyk includes this column only for languages Snyk Agent Fix supports.

| Rule Name                                                                      | CWEs                                                        | Security Categories                                     | Autofixable |
| ------------------------------------------------------------------------------ | ----------------------------------------------------------- | ------------------------------------------------------- | ----------- |
| Code Injection                                                                 | CWE-94                                                      | SANS Top 25, OWASP:A03                                  | Yes         |
| Command Injection                                                              | CWE-78                                                      | SANS Top 25, OWASP:A03                                  | Yes         |
| Remote Code Execution via Endpoint                                             | CWE-94                                                      | SANS Top 25, OWASP:A03                                  | Yes         |
| Deserialization of Untrusted Data                                              | CWE-502                                                     | SANS Top 25, OWASP:A08                                  | Yes         |
| Use of Hardcoded Credentials                                                   | CWE-798, CWE-259                                            | SANS Top 25, OWASP:A07                                  | Yes         |
| Use of Hardcoded Cryptographic Key                                             | CWE-321                                                     | OWASP:A02                                               | Yes         |
| Hardcoded Secret                                                               | CWE-547                                                     | OWASP:A05                                               | Yes         |
| Use of Hardcoded Passwords                                                     | CWE-798, CWE-259                                            | SANS Top 25, OWASP:A07                                  | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm                               | CWE-327                                                     | OWASP:A02                                               | Yes         |
| Use of Password Hash With Insufficient Computational Effort                    | CWE-916                                                     | OWASP:A02                                               | Yes         |
| Use of Insufficiently Random Values                                            | CWE-330                                                     | OWASP:A02                                               | Yes         |
| Sinatra Protection Layers Disabled                                             | CWE-16, CWE-352, CWE-79, CWE-693, CWE-1021, CWE-35, CWE-348 | SANS Top 25, OWASP:A05, OWASP:A03, OWASP:A01, OWASP:A04 | Yes         |
| Insecure Data Transmission                                                     | CWE-319                                                     | OWASP:A02                                               | Yes         |
| Improper Input Validation                                                      | CWE-20                                                      | SANS Top 25, OWASP:A03                                  | Yes         |
| Improperly Controlled Modification of Dynamically-Determined Object Attributes | CWE-915                                                     | OWASP:A08                                               | Yes         |
| Selection of Less-Secure Algorithm During Negotiation (Force SSL)              | CWE-311, CWE-757                                            | OWASP:A02, OWASP:A04                                    | Yes         |
| Open Redirect                                                                  | CWE-601                                                     | OWASP:A01                                               | Yes         |
| Path Traversal                                                                 | CWE-23                                                      | OWASP:A01                                               | Yes         |
| Unsafe Reflection                                                              | CWE-470                                                     | OWASP:A03                                               | Yes         |
| Improper Certificate Validation                                                | CWE-295                                                     | OWASP:A07                                               | Yes         |
| Information Exposure                                                           | CWE-200                                                     | OWASP:A01                                               | Yes         |
| Session Manipulation                                                           | CWE-285                                                     | OWASP:A01                                               | Yes         |
| SQL Injection                                                                  | CWE-89                                                      | SANS Top 25, OWASP:A03                                  | Yes         |
| Improper Neutralization of Directives in Statically Saved Code                 | CWE-96                                                      | OWASP:A03                                               | Yes         |
| No Weak Password Requirements                                                  | CWE-521                                                     | OWASP:A07                                               | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                                       | CWE-1004                                                    | OWASP:A05                                               | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute                   | CWE-614                                                     | OWASP:A05                                               | Yes         |
| Incorrect regular expression for validating values                             | CWE-1286                                                    | None                                                    | Yes         |
| Cross-site Scripting (XSS)                                                     | CWE-79                                                      | SANS Top 25, OWASP:A03                                  | Yes         |
| XML External Entity (XXE) Injection                                            | CWE-611                                                     | OWASP:A05                                               | Yes         |
| XPath Injection                                                                | CWE-643                                                     | OWASP:A03                                               | Yes         |
| Regular Expression Denial of Service (ReDoS)                                   | CWE-400                                                     | None                                                    | Yes         |
