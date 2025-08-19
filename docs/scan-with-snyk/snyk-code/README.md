# Snyk Code

## Overview

Snyk Code is a developer-first static application security testing (SAST) solution. By scanning code in real-time and providing actionable insights directly in the developer workflow across IDEs, repositories, and CI/CD pipelines, you can identify and address vulnerabilities early on. The AI-based engine results in fewer false positives for your developers, improving code quality and security. You can scan your code using the following options:

* Snyk web UI (including [PR checks](../pull-requests/pull-request-checks/))
* [Snyk IDE](../../developer-tools/snyk-ide-plugins-and-extensions/)
* [Snyk CLI](../../developer-tools/snyk-cli/)
* [Snyk API](../../snyk-api/snyk-api.md)

The following table shows the Snyk Code features, including analysis, managing security issues in your code, and facilitating remediations within your development environment.

<table><thead><tr><th width="179">Feature</th><th>Description</th></tr></thead><tbody><tr><td>Issue filtering, sorting, and grouping</td><td><p>To identify the most common problems, you can filter issues based on their severity, programming language, priority score, and other criteria.</p><p>See <a href="manage-code-vulnerabilities/#filtering-existing-projects">Filter existing Projects</a>.</p></td></tr><tr><td>Priority Score</td><td><p>Sort by and prioritize the more important issues by incorporating factors such as issue prevalence, ease of fix, and risk factor into a single risk score.</p><p>See <a href="../../manage-risk/prioritize-issues-for-fixing/priority-score.md">Priority score</a>.</p></td></tr><tr><td>Data flow</td><td><p>Visualize the path of the issue from source to sink with a step-by-step flow.</p><p>See <a href="manage-code-vulnerabilities/breakdown-of-code-analysis.md">Data flow</a>.</p></td></tr><tr><td>Vulnerability</td><td><p>Learn more about the vulnerability through curated content that explains how the vulnerability was created, what the risk factors are, and popular mitigation strategies for it.</p><p>See <a href="manage-code-vulnerabilities/">Manage code vulnerabilities</a></p></td></tr><tr><td>Fix analysis</td><td><p>Gain insight and context by examining examples with links to actual code that fixes the same issues in similar data flows.</p><p>See <a href="manage-code-vulnerabilities/breakdown-of-code-analysis.md">Breakdown of Code analysis</a>.</p></td></tr><tr><td>Create Jira issue</td><td><p>Track and export Snyk issues to your Jira project.</p><p>See <a href="../../integrations/jira-and-slack-integrations/jira-integration.md#create-a-jira-issue">Create a Jira issue</a>.</p></td></tr><tr><td>Ignore issues</td><td><p>Configure Snyk to ignore suggested fixes for an issue to suppress specific warnings. For example, you may have deliberately used hard-coded passwords to test your routines in test code, or you are aware of an issue but have decided not to fix it.</p><p>See <a href="../../manage-risk/prioritize-issues-for-fixing/ignore-issues/">Ignore issues</a>.</p></td></tr><tr><td>Exclude files from the import process</td><td><p>Check for <code>DeepCode/Snyk</code> ignore files <code>.gitignore</code> <code>.dcignore</code> and read them if they exist. Using the information in these files, Snyk filters to identify only the files with <a href="../../supported-languages/supported-languages-package-managers-and-frameworks.md">the supported extensions</a> in the Project directory and not above the current Project directory. Snyk Code bundles these files that are smaller than 4 MB and sends them to Snyk. <code>,gitignore</code> exclusions are honored by the <code>snyk code test</code> CLI command.</p><p>See also <a href="../import-project-repository/exclude-directories-and-files-from-project-import.md">Exclude directories and files from the import process</a>.</p></td></tr><tr><td>Interfile analysis</td><td>This is available for <a href="../../supported-languages/supported-languages-package-managers-and-frameworks.md#code-analysis-snyk-code">all languages supported by Snyk Code</a> except Ruby.</td></tr></tbody></table>

## Deployment

| Deployment                                                       | Description                                                                                                                                                                                                                                                                           |
| ---------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Full SaaS solution                                               | <p>Get the most out of Snyk Code with a native Git repository integration, easy onboarding, and continuous updates.<br><br>See <a href="configure-snyk-code.md">Configure Snyk Code</a>.</p>                                                                                          |
| SaaS with a self-hosted Git server (requires Snyk Broker)        | For customers with SCMs that are not publicly accessible from the internet and want to connect Snyk Code with their local self-hosted SCM.                                                                                                                                            |
| Local no-upload implementation (requires Snyk Code Local Engine) | <p>For customers with a stricter upload policy.</p><p><br>This deployment method requires more maintenance and receives slower updates than the SaaS options, but it does not require any code upload.<br><br>See <a href="snyk-code-local-engine.md">Snyk Code Local Engine</a>.</p> |

## AI Engine

Snyk Code is powered by a semantic, AI-based analysis engine and can analyze the following in your code:

* API usage: Identifies multiple potential issues, including API misuses, null dereferences, and type mismatches, by modeling the use of memory in variables and references. This mechanism can also identify the use of insecure functions.
* Coding issues: Finds problems such as dead code, branches that are predefined, and branches having the same code on each side.
* Control flow: Identifies null dereference or race conditions by modeling each possible control flow in the application.
* Data flow: Follows the flow of data within the application from the source to the sink. Combined with AI-based learning of external insecure data sources, data sinks, and sanitation functions, this enables a strong taint analysis.
* Hardcoded secrets: Hardcoded secrets detection rules are invoked during SAST scans but do not act as a standalone secrets scanning tool, as this is done through our partnership with third-party tools. See our Snyk Learn lessons on [GitGuardian](https://learn.snyk.io/lesson/snyk-apprisk-gitguardian/) and [Nightfall AI](https://learn.snyk.io/lesson/snyk-apprisk-nightfall-ai/).

<figure><img src="../../.gitbook/assets/Introduction - AI Engine - Hardcoded secrets.png" alt="Hardcoded secret found"><figcaption><p>Hardcoded secret found</p></figcaption></figure>

* Point-to analysis: Identifies multiple potential issues, including buffer overruns, null dereferences, and type mismatches, by modeling memory use in variables and references.
* Type inference: Determines the initial type and its changes. This is of special interest for dynamically typed languages.
* Value ranges: Infers possible values for variables used to call functions to track off-by-one errors in arrays, division-by-zero errors, and null dereferences.

## Supported integrations

* IDE
  * [JetBrains plugin](../../developer-tools/snyk-ide-plugins-and-extensions/jetbrains-plugin/)
  * [Visual Studio Code extension](../../developer-tools/snyk-ide-plugins-and-extensions/visual-studio-code-extension/)
  * [Visual Studio extension](../../developer-tools/snyk-ide-plugins-and-extensions/visual-studio-extension/)
* Git repository: With repository monitoring integration, you can actively manage your Code Projects using the existing native import flow and tools. You can view and prioritize security issues discovered in your source code. Additionally, you can initiate a retest of any Project and examine the historical snapshots to track changes over time.&#x20;

{% hint style="info" %}
Snyk Code analysis can be applied to every pull request you create in your Git repository before you merge it into the target branch. See [PR Checks](../pull-requests/pull-request-checks/).
{% endhint %}

* CLI and CI/CD: [Using the CLI](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/) helps you find and fix security flaws in your code on your local machine or in your CI/CD.
* APIs and extensibility: Query Code Projects and issues using the Snyk [REST API](../../snyk-api/reference/).
* Notifications: [Integrate with Jira](../../integrations/jira-and-slack-integrations/jira-integration.md) to export data to Jira issues.

## Supported languages

Snyk Code supports many [languages and frameworks](../../supported-languages/supported-languages-package-managers-and-frameworks.md).

## What's next?

* [Create or log in to a Snyk account](../../discover-snyk/getting-started/#create-or-log-in-to-a-snyk-account)
* [Set up an integration](../../discover-snyk/getting-started/#set-up-a-snyk-integration)
* [Import a Project](../../discover-snyk/getting-started/#import-a-project-to-scan-and-identify-issues)
* [Configure Snyk Code in Snyk Web UI](configure-snyk-code.md)
