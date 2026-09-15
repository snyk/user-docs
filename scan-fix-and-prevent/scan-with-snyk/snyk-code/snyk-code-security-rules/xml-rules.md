---
description: Snyk Code security rules for XML
nav_context: agnostic
---

# XML rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

| Rule Name                                                    | CWEs             | Security Categories                                       |
| ------------------------------------------------------------ | ---------------- | --------------------------------------------------------- |
| Android Debug Mode Enabled                                   | CWE-489          | OWASP:A02:2025                                            |
| Debug Features Enabled                                       | CWE-215          | OWASP:A10:2025                                            |
| Generation of Error Message Containing Sensitive Information | CWE-209          | OWASP:A10:2025, OWASP-API:API8:2023                       |
| Improper Restriction of Rendered UI Layers or Frames         | CWE-1021         | OWASP:A06:2025                                            |
| ASP SSL Disabled                                             | CWE-319          | OWASP:A04:2025, OWASP-API:API8:2023, OWASP-API:API10:2023 |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | OWASP:A07:2025                                            |
| Request Validation Disabled                                  | CWE-554          | None                                                      |
| Struts Development Mode Enabled                              | CWE-489          | OWASP:A02:2025                                            |
