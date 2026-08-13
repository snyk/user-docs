---
description: Snyk Code security rules for XML
nav_context: agnostic
---

# XML rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/) (2025 edition) category the rule belongs to, if any, and whether [SANS Top 25](https://www.sans.org/top25-software-errors/) includes it.
* **Autofixable**: Whether Snyk Agent Fix can fix the rule automatically. Snyk includes this column only for languages Snyk Agent Fix supports.

| Rule Name                                                    | CWEs             | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Android Debug Mode Enabled                                   | CWE-489          | None                   | Yes         |
| Debug Features Enabled                                       | CWE-215          | None                   | Yes         |
| Generation of Error Message Containing Sensitive Information | CWE-209          | OWASP:A04              | Yes         |
| Improper Restriction of Rendered UI Layers or Frames         | CWE-1021         | OWASP:A04              | Yes         |
| ASP SSL Disabled                                             | CWE-319          | OWASP:A02              | Yes         |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | SANS Top 25, OWASP:A07 | Yes         |
| Request Validation Disabled                                  | CWE-554          | None                   | Yes         |
| Struts Development Mode Enabled                              | CWE-489          | None                   | Yes         |
