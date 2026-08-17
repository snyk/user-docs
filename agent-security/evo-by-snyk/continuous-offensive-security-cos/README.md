# Continuous Offensive Security (COS)

Snyk Continuous Offensive Security is an AI-driven pentesting engine that continuously tests your web applications, validates real vulnerabilities, and tracks findings over time.

Use the pages in this section to configure targets, run scans, and review findings.

### Assessment depth

Depth is determined by whether a source repository is linked. There is no separate setting to change.

| Depth        | Condition            | What you get                                                                                                                                                                           |
| ------------ | -------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Blackbox** | No repository linked | Agents test entirely from the outside. Remediation guidance is generic to the vulnerability class, with references                                                                     |
| **Greybox**  | Repository linked    | Agents use source code to understand structure and reach logic flaws that are invisible from the outside. Remediation guidance names the file, the line, and the suggested code change |

Link repositories to help Snyk find more issues and provide more actionable guidance on fixes.
