# Groovy rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                     | CWE(s)           | Security Categories    | Autofixable |
| --------------------------------------------- | ---------------- | ---------------------- | ----------- |
| Code Injection                                | CWE-94           | OWASP:A03              | No          |
| Command Injection                             | CWE-78           | Sans Top 25, OWASP:A03 | No          |
| Deserialization of Untrusted Data             | CWE-502          | OWASP:A08              | No          |
| Hardcoded Secret                              | CWE-547          | OWASP:A05              | No          |
| Open Redirect                                 | CWE-601          | OWASP:A01              | No          |
| Path Traversal                                | CWE-23           | OWASP:A01              | No          |
| SQL Injection                                 | CWE-89           | Sans Top 25, OWASP:A03 | No          |
| Server-Side Request Forgery (SSRF)            | CWE-918          | Sans Top 25, OWASP:A10 | No          |
| Use of Hardcoded Passwords                    | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | No          |
| Use of Hardcoded, Security-relevant Constants | CWE-547          | OWASP:A05              | No          |
| XML External Entity (XXE) Injection           | CWE-611          | OWASP:A05              | No          |