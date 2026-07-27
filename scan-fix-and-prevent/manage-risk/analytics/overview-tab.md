---
description: What the Overview tab shows in Snyk Analytics
nav_context: classic
---

# Overview tab

The **Overview** tab in **Snyk Analytics** provides a centralized visualization of your organization's security posture across Snyk Code, Snyk Container, Snyk Infrastructure as Code (IaC), and Snyk Open Source. Use this dashboard to identify coverage gaps, monitor historical risk trends, and evaluate remediation efficiency without drilling down into individual projects.

Dashboard data adjusts automatically based on your selected global filters and time range. The dashboard is divided into several data sections to help you focus on specific metrics.

## Dashboard widgets

### Key Performance Indicators

This section displays aggregate totals for the selected time range and calculates the percentage change compared to the previous equivalent period.

The following widgets are enabled by default:

* **Open issues**: The total count of unresolved vulnerabilities.
* **Ignored issues**: The total count of vulnerabilities suppressed by Snyk ignore policies.
* **New issues**: The volume of vulnerabilities introduced during the selected timeframe.
* **Resolved issues**: The total count of vulnerabilities fixed during the selected timeframe.
* **Unique vulnerabilities**: The distinct count of individual vulnerability types, regardless of how many times they appear across your environments.

Additionally, you can add the following widgets:

* **Projects Monitored**: Number of Snyk Projects continuously monitored for open-source vulnerabilities and license issues after using the `snyk monitor` CLI command.
* **Agentic Scans**: Number of Snyk Studio scans for the selected timeframe.
* **Developers running agentic scans: S**nyk Studio adoption rate.
* **Total PR checks**: Number of pull request (PR) checks for the selected timeframe.
* **PR check success rate**: Outcomes of PR checks. An increasing success rate over time demonstrates that developers produce more secure code earlier in the software development lifecycle.
* **Snyk Organizations**: Track Snyk rollout progress by viewing the total number of new Snyk Organizations created over time.
* **Zero-Day Open Issues**: Use this centralized dashboard to assess immediate risk and track your response during a major zero-day security incident.
* **Developers testing in the IDE/CLI**:  Number of developers who run a Snyk scan during the selected timeframe.
* **SCA Preventable issues**: It shows the number of new open-source (SCA) vulnerabilities introduced into your codebase that you can block earlier, for example, during pull request checks.
* **Tested repositories in CI/CD**: Number of repositories tested in the continuous integration and continuous delivery (CI/CD) pipeline.

### Coverage

The Coverage section helps you identify blind spots in your security rollout and ensure your assets are actively monitored.

The following widgets are enabled by default:

* **Repository coverage by Snyk product**: Identifies unmonitored assets. Use the **Not Scanned** segments to pinpoint exactly where you need to deploy Snyk to achieve full visibility.
* **Repository coverage percentage over time, by Snyk product**: Tracks your onboarding progress. Use this trend line to verify that you successfully integrate new repositories into Snyk over time.
* **Assets discovered over time, by type**: Monitors your environment's growth. Use this to ensure Snyk successfully detects new repositories as your organization scales.

Additionally, you can add the **Targets by integration** widget to track the number of repositories each integration imports into Snyk.

### Exposure

The Exposure section helps you triage risk by isolating where vulnerabilities concentrate and prioritizing them by severity.

The following widgets are enabled by default:

* **Organizations with the most open issues**: Rank business units or teams by risk volume. Use this chart to determine which teams require the most immediate remediation support or intervention.
* **Open issues over time, by severity**: Categorizes your risk backlog. Use this to verify that teams prioritize and resolve Critical and High severity issues faster than lower-severity issues.
* **Open issues by Snyk product**: Illustrates risk distribution across your stack. Use this to determine if your primary vulnerabilities lie in your proprietary code, open-source libraries, containers, or infrastructure configurations.

Additionally, you can add the following widgets:

* **Open issues by severity**: Categorizes open issues by severity level.
* **New issues by repository freshness**: Identifies risks based on the last commit (freshness) to highlight blind spots. You can configure the definitions for these statuses:

| Status       | Definition                                    |
| ------------ | --------------------------------------------- |
| **Active**   | Less than 3 months                            |
| **Inactive** | 3-6 months                                    |
| **Dormant**  | Over 6 months                                 |
| **N/A**      | The Group SCM integration detected no commits |

* **New issues (OWASP Top 10, 2025) by severity**: Group open issues by OWASP Top 10 (2025) category and severity.
* **New issues by asset class**: Enables prioritization of remediation efforts based on the business-criticality of different assets.
* **SCA CVEs with the most open issues**: Highlights high-impact open-source vulnerabilities (CVEs) with the most open issues identified through software composition analysis (SCA).

### Remediation

The Remediation section measures your organization's efficiency in fixing security flaws and helps track compliance with Service Level Agreements (SLAs).

The following widgets are enabled by default:

* **Issues identified and resolved over time**: Compares issue creation against issue resolution. Use this chart to determine whether your teams fix vulnerabilities faster than they introduce them.
* **Resolved issues by MTTR duration**: Groups your Mean Time to Resolution (MTTR) into time buckets. Use this to evaluate whether teams meet your organization's required security SLAs (for example, fixing critical issues in less than 15 days).
* **Organizations with the highest MTTR**: Rank teams by how long they take to fix issues. Use this to identify organizations that struggle with complex fixes or require workflow improvements.

Additionally, you can add the following widgets:

* **Top 5 vulnerabilities approaching eradication by CVE:** Highlights specific known vulnerabilities that are closest to being completely resolved across your environment.
* **Top 5 vulnerabilities approaching eradication by CWE:** Highlights broader categories of software weaknesses that your team is close to eliminating entirely.
* **Snyk-generated PRs merged vs open:** Compares the number of merged Snyk-generated pull requests to the number of open pull requests.

### Usage

The Usage section tracks developer engagement and measures the success of your shift-left prevention strategies.

The following widgets are enabled by default:

* **Snyk IDE and CLI usage over time (developer count)**: Measures proactive tool adoption. Use this to verify that developers actively test code locally before pushing it to your repositories.
* **Organizations introducing most new SCA preventable issues**: Pinpoints where known, avoidable open-source risks enter your codebase. Use this to target specific teams for training on secure coding practices or Snyk use.
* **New issues by introduction category**: Classifies the avoidability of new vulnerabilities. Use the **Preventable Issue** segment to measure how many issues Snyk scans could have blocked earlier, helping you justify stricter PR checks.

You can also add the **New SCA preventable issues introduced over time** widget. This widget classifies the avoidability of new vulnerabilities. Use the Preventable Issue segment to measure how many issues Snyk scans can block earlier to help you justify stricter pull request checks.

## Customize the dashboard

You can customize the layout to prioritize specific data:

* Show or hide analytics sections: Click **Edit Page Layout**.
* Reorder the pag&#x65;**:** Click and drag a section title.
* Select or remove widget&#x73;**:** Click **Edit section**.
* Move widget&#x73;**:** Use the hamburger menu.
* Save your changes in a saved vie&#x77;**:** Click **Save**.
