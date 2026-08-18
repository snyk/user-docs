---
description: Snyk Code security rules for JavaScript and TypeScript
nav_context: agnostic
---

# JavaScript and TypeScript rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                                                                                         | CWEs                    | Security Categories                                                    |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------- | ---------------------------------------------------------------------- |
| Disabling Strict Contextual escaping (SCE) could provide additional attack surface for Cross-site Scripting (XSS) | CWE-79                  | CWE Top 25, OWASP:A05:2025                                             |
| Use of a Broken or Risky Cryptographic Algorithm                                                                  | CWE-327                 | OWASP:A04:2025                                                         |
| Clear Text Sensitive Storage                                                                                      | CWE-200, CWE-312        | CWE Top 25, OWASP:A01:2025, OWASP:A06:2025, OWASP-API:API10:2023       |
| Code Injection                                                                                                    | CWE-94                  | CWE Top 25, OWASP:A05:2025                                             |
| Command Injection                                                                                                 | CWE-78                  | CWE Top 25, OWASP:A05:2025                                             |
| Cross-site Scripting (XSS)                                                                                        | CWE-79                  | CWE Top 25, OWASP:A05:2025                                             |
| Deserialization of Untrusted Data                                                                                 | CWE-502                 | CWE Top 25, OWASP:A08:2025                                             |
| Information Exposure                                                                                              | CWE-200                 | CWE Top 25, OWASP:A01:2025, OWASP-API:API10:2023                       |
| Electron Disable Security Warnings                                                                                | CWE-16                  | OWASP:A02:2025, OWASP-API:API8:2023                                    |
| Electron Insecure Web Preferences                                                                                 | CWE-16                  | OWASP:A02:2025, OWASP-API:API8:2023                                    |
| Electron Load Insecure Content                                                                                    | CWE-16                  | OWASP:A02:2025, OWASP-API:API8:2023                                    |
| Use of Externally-Controlled Format String                                                                        | CWE-134                 | None                                                                   |
| GraphQL Injection                                                                                                 | CWE-89                  | CWE Top 25, OWASP:A05:2025                                             |
| Improper Type Validation                                                                                          | CWE-1287                | None                                                                   |
| Hardcoded Secret                                                                                                  | CWE-547                 | OWASP:A02:2025                                                         |
| Cleartext Transmission of Sensitive Information                                                                   | CWE-319                 | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023              |
| Improper Code Sanitization                                                                                        | CWE-94, CWE-79, CWE-116 | CWE Top 25, OWASP:A05:2025                                             |
| Use of Password Hash With Insufficient Computational Effort                                                       | CWE-916                 | OWASP:A04:2025                                                         |
| Use of Insufficiently Random Values                                                                               | CWE-330                 | OWASP:A04:2025                                                         |
| Insecure TLS Configuration                                                                                        | CWE-327                 | OWASP:A04:2025                                                         |
| Insufficient postMessage Validation                                                                               | CWE-20                  | CWE Top 25, OWASP:A05:2025, OWASP-API:API10:2023, OWASP-Mobile:M4:2024 |
| Introspection Enabled                                                                                             | CWE-200                 | CWE Top 25, OWASP:A01:2025, OWASP-API:API10:2023                       |
| Insecure JWT Verification Method                                                                                  | CWE-347                 | OWASP:A04:2025                                                         |
| JWT Signature Verification Method Disabled                                                                        | CWE-347                 | OWASP:A04:2025                                                         |
| JWT 'none' Algorithm Supported                                                                                    | CWE-347                 | OWASP:A04:2025                                                         |
| Denial of Service (DoS) through Nested GraphQL Queries                                                            | CWE-400                 | OWASP-API:API4:2023                                                    |
| Unchecked Input for Loop Condition                                                                                | CWE-400, CWE-606        | OWASP-API:API4:2023                                                    |
| Observable Timing Discrepancy (Timing Attack)                                                                     | CWE-208                 | None                                                                   |
| Use of Hardcoded Credentials                                                                                      | CWE-798                 | OWASP:A07:2025                                                         |
| Use of Hardcoded Passwords                                                                                        | CWE-798, CWE-259        | OWASP:A07:2025                                                         |
| Allocation of Resources Without Limits or Throttling                                                              | CWE-770                 | CWE Top 25, OWASP-API:API4:2023                                        |
| NoSQL Injection                                                                                                   | CWE-943                 | None                                                                   |
| Buffer Over-read                                                                                                  | CWE-126                 | None                                                                   |
| Open Redirect                                                                                                     | CWE-601                 | OWASP:A01:2025                                                         |
| Path Traversal                                                                                                    | CWE-23                  | OWASP:A01:2025                                                         |
| Prototype Pollution                                                                                               | CWE-1321                | None                                                                   |
| Use dangerouslySetInnerHTML to Explicitly Handle XSS Risks                                                        | CWE-79                  | CWE Top 25, OWASP:A05:2025                                             |
| Weak Password Recovery Mechanism for Forgotten Password                                                           | CWE-640                 | OWASP:A07:2025                                                         |
| SQL Injection                                                                                                     | CWE-89                  | CWE Top 25, OWASP:A05:2025                                             |
| Server-Side Request Forgery (SSRF)                                                                                | CWE-918                 | CWE Top 25, OWASP:A01:2025, OWASP-API:API7:2023                        |
| Improper Neutralization of Directives in Statically Saved Code                                                    | CWE-96                  | OWASP:A05:2025                                                         |
| Origin Validation Error                                                                                           | CWE-942, CWE-346        | OWASP:A02:2025, OWASP:A07:2025, OWASP-API:API8:2023                    |
| Permissive Cross-domain Policy                                                                                    | CWE-942                 | OWASP:A02:2025, OWASP-API:API8:2023                                    |
| Improper Restriction of Rendered UI Layers or Frames                                                              | CWE-1021                | OWASP:A06:2025                                                         |
| Cryptographic Issues                                                                                              | CWE-310                 | None                                                                   |
| Unsafe JQuery Plugin                                                                                              | CWE-79, CWE-116         | CWE Top 25, OWASP:A05:2025                                             |
| Cross-Site Request Forgery (CSRF)                                                                                 | CWE-352                 | CWE Top 25, OWASP:A01:2025                                             |
| Sensitive Cookie Without 'HttpOnly' Flag                                                                          | CWE-1004                | OWASP:A02:2025                                                         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute                                                      | CWE-614                 | OWASP:A02:2025                                                         |
| XML External Entity (XXE) Injection                                                                               | CWE-611                 | OWASP:A02:2025                                                         |
| XPath Injection                                                                                                   | CWE-643                 | OWASP:A05:2025                                                         |
| Arbitrary File Write via Archive Extraction (Zip Slip)                                                            | CWE-22                  | CWE Top 25, OWASP:A01:2025                                             |
| Regular Expression Denial of Service (ReDoS)                                                                      | CWE-400                 | OWASP-API:API4:2023                                                    |
