# Overview

Snyk helps you prioritize issues for fixing by distinguishing between insignificant alerts and actionable threats. Not all vulnerabilities pose the same level of danger.&#x20;

Snyk uses contextual signals, such as code reachability and deployment environments, and prioritization methods, including exploit maturity, severity levels, and holistic risk scoring.

## Prioritize issues for fixing

* [Risk Score](prioritize-issues-for-fixing/risk-score.md) and [Priority Score](prioritize-issues-for-fixing/priority-score.md): Use Risk Score and Priority Score to rank issues for remediation. Learn how they differ in [Priority Score vs Risk Score](prioritize-issues-for-fixing/priority-score-vs-risk-score.md).
* [Reachability Analysis](prioritize-issues-for-fixing/reachability-analysis.md): Snyk determines whether your application's code calls specific functions or modules that contain a vulnerability. This helps you prioritize fixes based on whether the threat is executable in your environment.
* [Exploit maturity](prioritize-issues-for-fixing/view-exploits.md): Track and display real-world exploits of a vulnerability, including active attacks or proof-of-concept, and prioritize fixes based on proven, active threats.
* [Malicious packages](prioritize-issues-for-fixing/malicious-packages.md): Malicious dependencies are an increasingly common method for executing software supply chain attacks.&#x20;
* [Severity levels](prioritize-issues-for-fixing/severity-levels.md): Snyk categorizes vulnerabilities into Critical, High, Medium, or Low. These levels are based on industry-standard scoring frameworks (like CVSS and CCSS) to quickly communicate the potential impact and risk of an exploit.
* [Application Context](../developer-tools/scm-integrations/application-context-for-scm-integrations/): Bring the application context with [Assets and risk factors](prioritize-issues-for-fixing/assets-and-risk-factors/), and [Set up Insights](prioritize-issues-for-fixing/set-up-insights/).

## Snyk Policies

Snyk [Policies](policies/) automate security and license governance. Define rules that automatically adjust issue severities or set compliance requirements across your Projects.

## Snyk Analytics

[Analytics](analytics/) provides a centralized reporting hub for Enterprise customers. Build custom dashboards and track security metrics across teams for high-level visibility into risk and remediation progress.

The [Reports tab](analytics/reports-tab/) organizes data into specific categories: **Exposure & coverage**, **Remediation**, **Prevention**, **Compliance**, and **Education**. Filter, save custom views, and export data (PDF/CSV) to facilitate conversations between security and development teams.

## Dependencies and licenses

View [dependencies](dependencies-and-licenses/view-dependencies.md) and [license](dependencies-and-licenses/view-licenses.md) details for all Projects in your Group or Organization by selecting the **Dependencies** option in your Group or Organization menu.
