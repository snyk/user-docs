---
description: Snyk Code security rules for XML
---

# XML rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s)**: The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category to which the rule belongs to, if any, and if it is included in [CWE Top 25](https://cwe.mitre.org/top25/), plus any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) and [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories                                       | Autofixable |
| ------------------------------------------------------------ | ---------------- | --------------------------------------------------------- | ----------- |
| Android Debug Mode Enabled                                   | CWE-489          | OWASP:A02:2025                                            | Yes         |
| Debug Features Enabled                                       | CWE-215          | OWASP:A10:2025                                            | Yes         |
| Generation of Error Message Containing Sensitive Information | CWE-209          | OWASP:A10:2025, OWASP-API:API8:2023                       | Yes         |
| Improper Restriction of Rendered UI Layers or Frames         | CWE-1021         | OWASP:A06:2025                                            | Yes         |
| ASP SSL Disabled                                             | CWE-319          | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023 | Yes         |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | OWASP:A07:2025                                            | Yes         |
| Request Validation Disabled                                  | CWE-554          | None                                                      | Yes         |
| Struts Development Mode Enabled                              | CWE-489          | OWASP:A02:2025                                            | Yes         |
