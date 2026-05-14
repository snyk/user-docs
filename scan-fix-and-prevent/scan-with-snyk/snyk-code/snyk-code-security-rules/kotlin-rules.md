# Kotlin rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Android World Writeable/Readable File Permission Found       | CWE-732          | None                   | No          |
| Use of Potentially Dangerous Function                        | CWE-676          | None                   | No          |
| Cleartext Storage of Sensitive Information in a Cookie       | CWE-315          | OWASP:A05              | No          |
| Code Injection                                               | CWE-94           | Sans Top 25, OWASP:A03 | No          |
| Command Injection                                            | CWE-78           | Sans Top 25, OWASP:A03 | No          |
| Deserialization of Untrusted Data                            | CWE-502          | Sans Top 25, OWASP:A08 | No          |
| Cross-Site Request Forgery (CSRF)                            | CWE-352          | Sans Top 25, OWASP:A01 | No          |
| Information Exposure                                         | CWE-200          | OWASP:A01              | No          |
| Cleartext Transmission of Sensitive Information              | CWE-319          | OWASP:A02              | No          |
| Indirect Command Injection via User Controlled Environment   | CWE-78           | Sans Top 25, OWASP:A03 | No          |
| External Control of System or Configuration Setting          | CWE-15           | OWASP:A05              | No          |
| Process Control                                              | CWE-114          | None                   | No          |
| File Access Enabled                                          | CWE-200          | OWASP:A01              | No          |
| Android Fragment Injection                                   | CWE-470          | OWASP:A03              | No          |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | No          |
| Hardcoded Secret                                             | CWE-547          | OWASP:A05              | No          |
| Improper Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A03              | No          |
| Disabled Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A03              | No          |
| Inadequate Padding for AES encryption                        | CWE-326          | OWASP:A02              | No          |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A02              | No          |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | No          |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A02              | No          |
| Android Intent Forwarding                                    | CWE-940          | OWASP:A07              | No          |
| Improper Validation of Certificate with Host Mismatch        | CWE-297          | OWASP:A07              | No          |
| JavaScript Enabled                                           | CWE-79           | Sans Top 25, OWASP:A03 | No          |
| Java Naming and Directory Interface (JNDI) Injection         | CWE-074          | None                   | No          |
| Improper Authentication                                      | CWE-287          | Sans Top 25, OWASP:A07 | No          |
| LDAP Injection                                               | CWE-90           | OWASP:A03              | No          |
| Use of Hardcoded Credentials                                 | CWE-798          | Sans Top 25, OWASP:A07 | No          |
| The cipher text is equal to the provided input plain text    | CWE-311          | OWASP:A04              | No          |
| Use of Sticky broadcasts                                     | CWE-265          | None                   | No          |
| Use of Hardcoded, Security-relevant Constants                | CWE-547          | OWASP:A05              | No          |
| Open Redirect                                                | CWE-601          | OWASP:A01              | No          |
| Path Traversal                                               | CWE-23           | OWASP:A01              | No          |
| Regular expression injection                                 | CWE-400, CWE-730 | None                   | No          |
| Unprotected Storage of Credentials                           | CWE-256          | OWASP:A04              | No          |
| Incorrect Permission Assignment                              | CWE-732          | None                   | No          |
| Improper Certificate Validation                              | CWE-295          | OWASP:A07              | No          |
| Server Information Exposure                                  | CWE-209          | OWASP:A04              | No          |
| Improper Handling of Insufficient Permissions or Privileges  | CWE-280          | OWASP:A04              | No          |
| Cross-site Scripting (XSS)                                   | CWE-79           | Sans Top 25, OWASP:A03 | No          |
| Unrestricted Android Broadcast                               | CWE-862          | Sans Top 25, OWASP:A01 | No          |
| SQL Injection                                                | CWE-89           | Sans Top 25, OWASP:A03 | No          |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | Sans Top 25, OWASP:A10 | No          |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A02              | No          |
| Code Execution via Third Party Package Context               | CWE-94           | Sans Top 25, OWASP:A03 | No          |
| Code Execution via Third Party Package Installation          | CWE-940          | OWASP:A07              | No          |
| Observable Timing Discrepancy (Timing Attack)                | CWE-208          | None                   | No          |
| Origin Validation Error                                      | CWE-942, CWE-346 | OWASP:A05, OWASP:A07   | No          |
| Cryptographic Issues                                         | CWE-310          | OWASP:A02              | No          |
| Trust Boundary Violation                                     | CWE-501          | OWASP:A04              | No          |
| Unauthorized File Access                                     | CWE-79           | Sans Top 25, OWASP:A03 | No          |
| Android Uri Permission Manipulation                          | CWE-266          | OWASP:A04              | No          |
| Use of Externally-Controlled Format String                   | CWE-134          | None                   | No          |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A05              | No          |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | No          |
| Insufficient Session Expiration                              | CWE-613          | OWASP:A07              | No          |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A05              | No          |
| XPath Injection                                              | CWE-643          | OWASP:A03              | No          |
