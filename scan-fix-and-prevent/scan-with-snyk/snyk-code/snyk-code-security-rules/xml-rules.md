# XML rules

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **Security Categories**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Autofixable**: Security rules that are autofixable by Snyk Agent Fix. This information is included only for the supported programming languages.

| Rule Name                                                    | CWE(s)           | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Android Debug Mode Enabled                                   | CWE-489          | None                   | No          |
| Debug Features Enabled                                       | CWE-215          | None                   | No          |
| Generation of Error Message Containing Sensitive Information | CWE-209          | OWASP:A04              | No          |
| Improper Restriction of Rendered UI Layers or Frames         | CWE-1021         | OWASP:A04              | No          |
| ASP SSL Disabled                                             | CWE-319          | OWASP:A02              | No          |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | No          |
| Request Validation Disabled                                  | CWE-554          | None                   | No          |
| Struts Development Mode Enabled                              | CWE-489          | None                   | No          |
