# Severity levels in findings

Snyk API & Web assigns a severity level to each finding to summarize its overall risk based on the following factors:

* The likelihood of the vulnerability being found and exploited.
* The skills required to exploit the vulnerability.
* The impact of exploiting the vulnerability.

For example, a vulnerability that is easy to find, easy to exploit, and has a high impact will likely be classified with a high severity.

Different findings for the same vulnerability can have different severity levels depending on the context in which Snyk API & Web finds the vulnerabilities. Multiple factors can influence this context, which Snyk API & Web considers to lower or raise the severity level. For example, the severity of a finding can be higher or lower depending on whether the scanned target has authentication configured.

## Severity levels

The following table describes the different severity levels.

| Severity | Description | Examples |
| --- | --- | --- |
| <figure><img src="../../../.gitbook/assets/severity-levels-critical.png" alt="Critical severity indicator"></figure> | These findings require immediate attention and remediation due to their potentially devastating impact. | SQL Injection, OS Command Injection |
| <figure><img src="../../../.gitbook/assets/severity-levels-high.png" alt="High severity indicator"></figure> | These findings may have a direct impact on the application security, either clients or service owners, for instance, by granting the attacker access to sensitive information. | Reflected Cross-Site Scripting, Path Traversal |
| <figure><img src="../../../.gitbook/assets/severity-levels-medium.png" alt="Medium severity indicator"></figure> | Medium findings do not usually have an immediate impact alone, but combined with other findings, may lead to a successful compromise of the application. | Cross-Site Request Forgery, Unencrypted Communications |
| <figure><img src="../../../.gitbook/assets/severity-levels-low.png" alt="Low severity indicator"></figure> | Findings where either the exploit is not trivial, or the finding cannot be exploited by itself. | Directory Listing, Clickjacking |

## Related information

Visit [Interpret target scan results](../start-scanning/interpret-target-scan-results.md) for more information about findings.
