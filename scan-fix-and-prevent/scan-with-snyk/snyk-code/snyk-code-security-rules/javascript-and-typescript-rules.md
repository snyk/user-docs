---
description: Snyk Code security rules for JavaScript and TypeScript
nav_context: agnostic
---

# JavaScript and TypeScript rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2021 edition) category the rule belongs to, if any, and whether [SANS Top 25](https://www.sans.org/top25-software-errors/) includes it.
* **Autofixable**: Whether Snyk Agent Fix can fix the rule automatically. Snyk includes this column only for languages Snyk Agent Fix supports.

| Rule Name                                                                                                         | CWEs                    | Security Categories    | Autofixable |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------- | ---------------------- | ----------- |
| Disabling Strict Contextual escaping (SCE) could provide additional attack surface for Cross-site Scripting (XSS) | CWE-79                  | SANS Top 25, OWASP:A03 | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm                                                                  | CWE-327                 | OWASP:A02              | Yes         |
| Clear Text Sensitive Storage                                                                                      | CWE-200, CWE-312        | OWASP:A01, OWASP:A04   | Yes         |
| Code Injection                                                                                                    | CWE-94                  | SANS Top 25, OWASP:A03 | Yes         |
| Command Injection                                                                                                 | CWE-78                  | SANS Top 25, OWASP:A03 | Yes         |
| Cross-site Scripting (XSS)                                                                                        | CWE-79                  | SANS Top 25, OWASP:A03 | Yes         |
| Deserialization of Untrusted Data                                                                                 | CWE-502                 | SANS Top 25, OWASP:A08 | Yes         |
| Information Exposure                                                                                              | CWE-200                 | OWASP:A01              | Yes         |
| Electron Disable Security Warnings                                                                                | CWE-16                  | OWASP:A05              | Yes         |
| Electron Insecure Web Preferences                                                                                 | CWE-16                  | OWASP:A05              | Yes         |
| Electron Load Insecure Content                                                                                    | CWE-16                  | OWASP:A05              | Yes         |
| Use of Externally-Controlled Format String                                                                        | CWE-134                 | None                   | Yes         |
| GraphQL Injection                                                                                                 | CWE-89                  | SANS Top 25, OWASP:A03 | Yes         |
| Improper Type Validation                                                                                          | CWE-1287                | None                   | Yes         |
| Hardcoded Secret                                                                                                  | CWE-547                 | OWASP:A05              | Yes         |
| Cleartext Transmission of Sensitive Information                                                                   | CWE-319                 | OWASP:A02              | Yes         |
| Improper Code Sanitization                                                                                        | CWE-94, CWE-79, CWE-116 | SANS Top 25, OWASP:A03 | Yes         |
| Use of Password Hash With Insufficient Computational Effort                                                       | CWE-916                 | OWASP:A02              | Yes         |
| Use of Insufficiently Random Values                                                                               | CWE-330                 | OWASP:A02              | Yes         |
| Insecure TLS Configuration                                                                                        | CWE-327                 | OWASP:A02              | Yes         |
| Insufficient postMessage Validation                                                                               | CWE-20                  | SANS Top 25, OWASP:A03 | Yes         |
| Introspection Enabled                                                                                             | CWE-200                 | OWASP:A01              | Yes         |
| Insecure JWT Verification Method                                                                                  | CWE-347                 | OWASP:A02              | Yes         |
| JWT Signature Verification Method Disabled                                                                        | CWE-347                 | OWASP:A02              | Yes         |
| JWT 'none' Algorithm Supported                                                                                    | CWE-347                 | OWASP:A02              | Yes         |
| Denial of Service (DoS) through Nested GraphQL Queries                                                            | CWE-400                 | None                   | Yes         |
| Unchecked Input for Loop Condition                                                                                | CWE-400, CWE-606        | None                   | Yes         |
| Observable Timing Discrepancy (Timing Attack)                                                                     | CWE-208                 | None                   | Yes         |
| Use of Hardcoded Credentials                                                                                      | CWE-798                 | SANS Top 25, OWASP:A07 | Yes         |
| Use of Hardcoded Passwords                                                                                        | CWE-798, CWE-259        | SANS Top 25, OWASP:A07 | Yes         |
| Allocation of Resources Without Limits or Throttling                                                              | CWE-770                 | None                   | Yes         |
| NoSQL Injection                                                                                                   | CWE-943                 | None                   | Yes         |
| Buffer Over-read                                                                                                  | CWE-126                 | None                   | Yes         |
| Open Redirect                                                                                                     | CWE-601                 | OWASP:A01              | Yes         |
| Path Traversal                                                                                                    | CWE-23                  | OWASP:A01              | Yes         |
| Prototype Pollution                                                                                               | CWE-1321                | None                   | Yes         |
| Use dangerouslySetInnerHTML to Explicitly Handle XSS Risks                                                        | CWE-79                  | SANS Top 25, OWASP:A03 | Yes         |
| Weak Password Recovery Mechanism for Forgotten Password                                                           | CWE-640                 | OWASP:A07              | Yes         |
| SQL Injection                                                                                                     | CWE-89                  | SANS Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                                                                                | CWE-918                 | SANS Top 25, OWASP:A10 | Yes         |
| Improper Neutralization of Directives in Statically Saved Code                                                    | CWE-96                  | OWASP:A03              | Yes         |
| Origin Validation Error                                                                                           | CWE-942, CWE-346        | OWASP:A05, OWASP:A07   | Yes         |
| Permissive Cross-domain Policy                                                                                    | CWE-942                 | OWASP:A05              | Yes         |
| Improper Restriction of Rendered UI Layers or Frames                                                              | CWE-1021                | OWASP:A04              | Yes         |
| Cryptographic Issues                                                                                              | CWE-310                 | OWASP:A02              | Yes         |
| Unsafe JQuery Plugin                                                                                              | CWE-79, CWE-116         | SANS Top 25, OWASP:A03 | Yes         |
| Cross-Site Request Forgery (CSRF)                                                                                 | CWE-352                 | SANS Top 25, OWASP:A01 | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                                                                          | CWE-1004                | OWASP:A05              | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute                                                      | CWE-614                 | OWASP:A05              | Yes         |
| XML External Entity (XXE) Injection                                                                               | CWE-611                 | OWASP:A05              | Yes         |
| XPath Injection                                                                                                   | CWE-643                 | OWASP:A03              | Yes         |
| Arbitrary File Write via Archive Extraction (Zip Slip)                                                            | CWE-22                  | SANS Top 25, OWASP:A01 | Yes         |
| Regular Expression Denial of Service (ReDoS)                                                                      | CWE-400                 | None                   | Yes         |
