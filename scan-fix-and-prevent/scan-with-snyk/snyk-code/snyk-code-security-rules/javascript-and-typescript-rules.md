---
description: Snyk Code security rules for JavaScript and TypeScript
---

# JavaScript and TypeScript rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/2025/)(2025 edition) category to which the rule belongs to, if any, and if it is included in [CWE Top 25](https://cwe.mitre.org/top25/), plus any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) and [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                                                                         | CWE(s)                  | Security Categories                                                    | Autofixable |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------- | ---------------------------------------------------------------------- | ----------- |
| Disabling Strict Contextual escaping (SCE) could provide additional attack surface for Cross-site Scripting (XSS) | CWE-79                  | CWE Top 25, OWASP:A05:2025                                             | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm                                                                  | CWE-327                 | OWASP:A04:2025                                                         | Yes         |
| Clear Text Sensitive Storage                                                                                      | CWE-200, CWE-312        | CWE Top 25, OWASP:A01:2025, OWASP:A06:2025, OWASP-API:API10:2023       | Yes         |
| Code Injection                                                                                                    | CWE-94                  | CWE Top 25, OWASP:A05:2025                                             | Yes         |
| Command Injection                                                                                                 | CWE-78                  | CWE Top 25, OWASP:A05:2025                                             | Yes         |
| Cross-site Scripting (XSS)                                                                                        | CWE-79                  | CWE Top 25, OWASP:A05:2025                                             | Yes         |
| Deserialization of Untrusted Data                                                                                 | CWE-502                 | CWE Top 25, OWASP:A08:2025                                             | Yes         |
| Information Exposure                                                                                              | CWE-200                 | CWE Top 25, OWASP:A01:2025, OWASP-API:API10:2023                       | Yes         |
| Electron Disable Security Warnings                                                                                | CWE-16                  | OWASP:A02:2025, OWASP-API:API8:2023                                    | Yes         |
| Electron Insecure Web Preferences                                                                                 | CWE-16                  | OWASP:A02:2025, OWASP-API:API8:2023                                    | Yes         |
| Electron Load Insecure Content                                                                                    | CWE-16                  | OWASP:A02:2025, OWASP-API:API8:2023                                    | Yes         |
| Use of Externally-Controlled Format String                                                                        | CWE-134                 | None                                                                   | Yes         |
| GraphQL Injection                                                                                                 | CWE-89                  | CWE Top 25, OWASP:A05:2025                                             | Yes         |
| Improper Type Validation                                                                                          | CWE-1287                | None                                                                   | Yes         |
| Hardcoded Secret                                                                                                  | CWE-547                 | OWASP:A02:2025                                                         | Yes         |
| Cleartext Transmission of Sensitive Information                                                                   | CWE-319                 | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023              | Yes         |
| Improper Code Sanitization                                                                                        | CWE-94, CWE-79, CWE-116 | CWE Top 25, OWASP:A05:2025                                             | Yes         |
| Use of Password Hash With Insufficient Computational Effort                                                       | CWE-916                 | OWASP:A04:2025                                                         | Yes         |
| Use of Insufficiently Random Values                                                                               | CWE-330                 | OWASP:A04:2025                                                         | Yes         |
| Insecure TLS Configuration                                                                                        | CWE-327                 | OWASP:A04:2025                                                         | Yes         |
| Insufficient postMessage Validation                                                                               | CWE-20                  | CWE Top 25, OWASP:A05:2025, OWASP-API:API10:2023, OWASP-Mobile:M4:2024 | Yes         |
| Introspection Enabled                                                                                             | CWE-200                 | CWE Top 25, OWASP:A01:2025, OWASP-API:API10:2023                       | Yes         |
| Insecure JWT Verification Method                                                                                  | CWE-347                 | OWASP:A04:2025                                                         | Yes         |
| JWT Signature Verification Method Disabled                                                                        | CWE-347                 | OWASP:A04:2025                                                         | Yes         |
| JWT 'none' Algorithm Supported                                                                                    | CWE-347                 | OWASP:A04:2025                                                         | Yes         |
| Denial of Service (DoS) through Nested GraphQL Queries                                                            | CWE-400                 | OWASP-API:API4:2023                                                    | Yes         |
| Unchecked Input for Loop Condition                                                                                | CWE-400, CWE-606        | OWASP-API:API4:2023                                                    | Yes         |
| Observable Timing Discrepancy (Timing Attack)                                                                     | CWE-208                 | None                                                                   | Yes         |
| Use of Hardcoded Credentials                                                                                      | CWE-798                 | OWASP:A07:2025                                                         | Yes         |
| Use of Hardcoded Passwords                                                                                        | CWE-798, CWE-259        | OWASP:A07:2025                                                         | Yes         |
| Allocation of Resources Without Limits or Throttling                                                              | CWE-770                 | CWE Top 25, OWASP-API:API4:2023                                        | Yes         |
| NoSQL Injection                                                                                                   | CWE-943                 | None                                                                   | Yes         |
| Buffer Over-read                                                                                                  | CWE-126                 | None                                                                   | Yes         |
| Open Redirect                                                                                                     | CWE-601                 | OWASP:A01:2025                                                         | Yes         |
| Path Traversal                                                                                                    | CWE-23                  | OWASP:A01:2025                                                         | Yes         |
| Prototype Pollution                                                                                               | CWE-1321                | None                                                                   | Yes         |
| Use dangerouslySetInnerHTML to Explicitly Handle XSS Risks                                                        | CWE-79                  | CWE Top 25, OWASP:A05:2025                                             | Yes         |
| Weak Password Recovery Mechanism for Forgotten Password                                                           | CWE-640                 | OWASP:A07:2025                                                         | Yes         |
| SQL Injection                                                                                                     | CWE-89                  | CWE Top 25, OWASP:A05:2025                                             | Yes         |
| Server-Side Request Forgery (SSRF)                                                                                | CWE-918                 | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023                        | Yes         |
| Improper Neutralization of Directives in Statically Saved Code                                                    | CWE-96                  | OWASP:A05:2025                                                         | Yes         |
| Origin Validation Error                                                                                           | CWE-942, CWE-346        | OWASP:A02:2025, OWASP:A07:2025, OWASP-API:API8:2023                    | Yes         |
| Permissive Cross-domain Policy                                                                                    | CWE-942                 | OWASP:A02:2025, OWASP-API:API8:2023                                    | Yes         |
| Improper Restriction of Rendered UI Layers or Frames                                                              | CWE-1021                | OWASP:A06:2025                                                         | Yes         |
| Cryptographic Issues                                                                                              | CWE-310                 | None                                                                   | Yes         |
| Unsafe JQuery Plugin                                                                                              | CWE-79, CWE-116         | CWE Top 25, OWASP:A05:2025                                             | Yes         |
| Cross-Site Request Forgery (CSRF)                                                                                 | CWE-352                 | CWE Top 25, OWASP:A01:2025                                             | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                                                                          | CWE-1004                | OWASP:A02:2025                                                         | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute                                                      | CWE-614                 | OWASP:A02:2025                                                         | Yes         |
| XML External Entity (XXE) Injection                                                                               | CWE-611                 | OWASP:A02:2025                                                         | Yes         |
| XPath Injection                                                                                                   | CWE-643                 | OWASP:A05:2025                                                         | Yes         |
| Arbitrary File Write via Archive Extraction (Zip Slip)                                                            | CWE-22                  | CWE Top 25, OWASP:A01:2025                                             | Yes         |
| Regular Expression Denial of Service (ReDoS)                                                                      | CWE-400                 | OWASP-API:API4:2023                                                    | Yes         |
