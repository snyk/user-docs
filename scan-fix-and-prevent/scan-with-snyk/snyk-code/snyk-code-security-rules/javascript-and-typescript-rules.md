# JavaScript and TypeScript rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                                                                         | CWE(s)                  | Security Categories    | Autofixable |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------- | ---------------------- | ----------- |
| Disabling Strict Contextual escaping (SCE) could provide additional attack surface for Cross-site Scripting (XSS) | CWE-79                  | Sans Top 25, OWASP:A03 | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm                                                                  | CWE-327                 | OWASP:A02              | No          |
| Clear Text Sensitive Storage                                                                                      | CWE-200, CWE-312        | OWASP:A01, OWASP:A04   | No          |
| Code Injection                                                                                                    | CWE-94                  | Sans Top 25, OWASP:A03 | Yes         |
| Command Injection                                                                                                 | CWE-78                  | Sans Top 25, OWASP:A03 | Yes         |
| Cross-site Scripting (XSS)                                                                                        | CWE-79                  | Sans Top 25, OWASP:A03 | Yes         |
| Deserialization of Untrusted Data                                                                                 | CWE-502                 | Sans Top 25, OWASP:A08 | No          |
| Information Exposure                                                                                              | CWE-200                 | OWASP:A01              | Yes         |
| Electron Disable Security Warnings                                                                                | CWE-16                  | OWASP:A05              | No          |
| Electron Insecure Web Preferences                                                                                 | CWE-16                  | OWASP:A05              | Yes         |
| Electron Load Insecure Content                                                                                    | CWE-16                  | OWASP:A05              | Yes         |
| Use of Externally-Controlled Format String                                                                        | CWE-134                 | None                   | Yes         |
| GraphQL Injection                                                                                                 | CWE-89                  | Sans Top 25, OWASP:A03 | No          |
| Improper Type Validation                                                                                          | CWE-1287                | None                   | Yes         |
| Hardcoded Secret                                                                                                  | CWE-547                 | OWASP:A05              | Yes         |
| Cleartext Transmission of Sensitive Information                                                                   | CWE-319                 | OWASP:A02              | Yes         |
| Improper Code Sanitization                                                                                        | CWE-94, CWE-79, CWE-116 | Sans Top 25, OWASP:A03 | No          |
| Use of Password Hash With Insufficient Computational Effort                                                       | CWE-916                 | OWASP:A02              | Yes         |
| Use of Insufficiently Random Values                                                                               | CWE-330                 | OWASP:A02              | No          |
| Insecure TLS Configuration                                                                                        | CWE-327                 | OWASP:A02              | Yes         |
| Insufficient postMessage Validation                                                                               | CWE-20                  | Sans Top 25, OWASP:A03 | Yes         |
| Introspection Enabled                                                                                             | CWE-200                 | OWASP:A01              | No          |
| Insecure JWT Verification Method                                                                                  | CWE-347                 | OWASP:A02              | No          |
| JWT Signature Verification Method Disabled                                                                        | CWE-347                 | OWASP:A02              | No          |
| JWT 'none' Algorithm Supported                                                                                    | CWE-347                 | OWASP:A02              | No          |
| Denial of Service (DoS) through Nested GraphQL Queries                                                            | CWE-400                 | None                   | Yes         |
| Unchecked Input for Loop Condition                                                                                | CWE-400, CWE-606        | None                   | No          |
| Observable Timing Discrepancy (Timing Attack)                                                                     | CWE-208                 | None                   | No          |
| Use of Hardcoded Credentials                                                                                      | CWE-798                 | Sans Top 25, OWASP:A07 | Yes         |
| Use of Hardcoded Passwords                                                                                        | CWE-798, CWE-259        | Sans Top 25, OWASP:A07 | Yes         |
| Allocation of Resources Without Limits or Throttling                                                              | CWE-770                 | None                   | Yes         |
| NoSQL Injection                                                                                                   | CWE-943                 | None                   | No          |
| Buffer Over-read                                                                                                  | CWE-126                 | None                   | No          |
| Open Redirect                                                                                                     | CWE-601                 | OWASP:A01              | Yes         |
| Path Traversal                                                                                                    | CWE-23                  | OWASP:A01              | Yes         |
| Prototype Pollution                                                                                               | CWE-1321                | None                   | No          |
| Use dangerouslySetInnerHTML to Explicitly Handle XSS Risks                                                        | CWE-79                  | Sans Top 25, OWASP:A03 | Yes         |
| Weak Password Recovery Mechanism for Forgotten Password                                                           | CWE-640                 | OWASP:A07              | No          |
| SQL Injection                                                                                                     | CWE-89                  | Sans Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                                                                                | CWE-918                 | Sans Top 25, OWASP:A10 | No          |
| Improper Neutralization of Directives in Statically Saved Code                                                    | CWE-96                  | OWASP:A03              | No          |
| Origin Validation Error                                                                                           | CWE-942, CWE-346        | OWASP:A05, OWASP:A07   | Yes         |
| Permissive Cross-domain Policy                                                                                    | CWE-942                 | OWASP:A05              | Yes         |
| Improper Restriction of Rendered UI Layers or Frames                                                              | CWE-1021                | OWASP:A04              | No          |
| Cryptographic Issues                                                                                              | CWE-310                 | OWASP:A02              | Yes         |
| Unsafe JQuery Plugin                                                                                              | CWE-79, CWE-116         | Sans Top 25, OWASP:A03 | No          |
| Cross-Site Request Forgery (CSRF)                                                                                 | CWE-352                 | Sans Top 25, OWASP:A01 | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                                                                          | CWE-1004                | OWASP:A05              | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute                                                      | CWE-614                 | OWASP:A05              | Yes         |
| XML External Entity (XXE) Injection                                                                               | CWE-611                 | OWASP:A05              | No          |
| XPath Injection                                                                                                   | CWE-643                 | OWASP:A03              | No          |
| Arbitrary File Write via Archive Extraction (Zip Slip)                                                            | CWE-22                  | Sans Top 25, OWASP:A01 | No          |
| Regular Expression Denial of Service (ReDoS)                                                                      | CWE-400                 | None                   | Yes         |
