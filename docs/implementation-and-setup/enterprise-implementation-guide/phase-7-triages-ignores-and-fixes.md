# Manage and remediate issues

After establishing visibility and prevention measures, focus on managing your existing vulnerability backlog. This stage involves defining a fix strategy, prioritizing issues based on risk, and operationalizing the remediation process in your development teams. Effective management ensures that your security posture improves over time rather than just maintaining the status quo.

To manage and remediate issues, follow these stages:

1. [Define a fix strategy:](phase-7-triages-ignores-and-fixes.md#define-a-fix-strategy) Identify focus areas and group work by team responsibilities.
2. [Apply prioritization filters:](phase-7-triages-ignores-and-fixes.md#prioritize-issues) Use Snyk metrics to identify the most critical issues.
3. [Establish remediation workflows:](phase-7-triages-ignores-and-fixes.md#establish-remediation-workflows) Integrate with ticket systems like Jira and manage exceptions using the ignore feature.
4. [Monitor adoption and progress:](phase-7-triages-ignores-and-fixes.md#monitor-adoption-and-progress) Use Snyk Reports to track resolution trends and team engagement.

## Define a fix strategy

{% hint style="success" %}
Key decision: Select a focus area based on business impact. Start with one high-traffic Organization or public-facing application to demonstrate success before scaling.
{% endhint %}

Analyze your imported Projects to determine where to start:

* **Identify critical assets**: Use discovery data obtained from configuring global settings to prioritize business-critical applications.
* **Group by team**: Assign Open Source and Snyk Code issues to development teams, and Container or IaC issues to DevOps teams.
* **Use metadata**: Filter Projects using attributes or tags to focus on specific business units or application types.

## Prioritize issues

{% hint style="success" %}
Key decision: Choose a primary metric for triage. Use **Priority Score** (900+) for a risk-based approach, or **Severity** (Critical/High) for a policy-based approach.
{% endhint %}

Use Snyk filters iteratively to build your plan:

* **Risk and priority scores**: Start with scores of 900–1000 and work downward.
* **Severity**: filter for **Critical** and **High** issues.

{% hint style="info" %}
For Snyk Open Source, prioritize critical issues with a **Fixable** filter to identify quick wins.
{% endhint %}

* **Exploit maturity**: Focus on issues with **Mature** or **Proof of Concept** exploit code to address the most reachable threats.
* **Issue type**: Run targeted vulnerability campaigns to eliminate specific types of flaws across all Projects, such as SQL injection (using CWE filters).

## Establish remediation workflows

{% hint style="success" %}
**Key decision**: Automate ticket creation. Use the `jira-tickets-for-new-vulns` tool or the Snyk Security in Jira Cloud plugin to ensure visibility in developer backlogs.
{% endhint %}

Operationalize the fix process using:

* **Jira integration**: Automatically create tickets for vulnerabilities that meet your priority criteria.
* **Ignore policy**: Use the **Ignore** feature for issues that cannot be fixed immediately due to environmental context or breaking changes.

{% hint style="info" %}
Ensure to always include a detailed reason and always set an expiration date (monthly or quarterly) for review.
{% endhint %}

* **Permissions**: Restrict ignore permissions to Organization admins in the general settings to maintain oversight.

## Monitor adoption and progress

{% hint style="success" %}
**Key decision**: Track **Resolved** versus **New** issues. A healthy program shows a downward trend in the total backlog over time.
{% endhint %}

Use the **Reports** tab to audit your progress using:

* **Issues summary**: View the **Risk Breakdown** to see open, new, and resolved issues.
* **Adoption tracking**: Identify which Organizations are most active in resolving issues to recognize successful teams or provide extra support where needed.
* **Organization reports**: Allow local admins to identify recurring vulnerabilities common across their specific repositories.
