# Snyk Code

## Overview

Snyk Code is a security tool that is fast and accurate and produces fewer false positives, making it easier for developers to remediate issues and build secure software.&#x20;

You can scan your code using the following options:

* [Snyk Web UI](../../getting-started/snyk-web-ui.md) (including [PR checks](../run-pr-checks/))
* [Snyk IDE](../../integrate-with-snyk/use-snyk-in-your-ide/)
* [Snyk CLI](../../snyk-cli/)
* [Snyk API](../../snyk-api/)

The following table shows the Snyk Code features, including analysis, managing security issues in your code, and facilitating remediations within your development environment.

<table><thead><tr><th width="179">Feature</th><th>Description</th></tr></thead><tbody><tr><td><strong>Issue filtering, sorting, and grouping</strong></td><td><p>To identify the most common problems, you can filter issues based on their severity, programming language, priority score, and other criteria. </p><p></p><p>See <a href="manage-code-vulnerabilities/#filtering-existing-projects">Filter existing Projects</a>.</p></td></tr><tr><td><strong>Priority Score</strong></td><td><p>Sort by and prioritize the more important issues by incorporating factors such as issue prevalence, ease of fix, and risk factor into a single risk score.</p><p></p><p>See <a href="../find-and-manage-priority-issues/priority-score.md">Priority score</a>.</p></td></tr><tr><td><strong>Data flow</strong></td><td><p>Visualize the path of the issue from source to sink with a step-by-step flow.</p><p></p><p>See <a href="manage-code-vulnerabilities/breakdown-of-code-analysis.md">Data flow</a>.</p></td></tr><tr><td><strong>Vulnerability</strong></td><td><p>Learn more about the vulnerability through curated content that explains how the vulnerability was created, what the risk factors are, and popular mitigation strategies for it.</p><p></p><p>See <a href="manage-code-vulnerabilities/">VManage code vulnerabilities</a></p></td></tr><tr><td><strong>Fix analysis</strong></td><td><p>Gain insight and context by examining examples with links to actual code that fixes the same issues in similar data flows.</p><p></p><p>See <a href="manage-code-vulnerabilities/breakdown-of-code-analysis.md">Breakdown of Code analysis</a>.</p></td></tr><tr><td><strong>Create Jira issue</strong></td><td><p>Track and export Snyk issues to your Jira project.</p><p></p><p>See <a href="../../integrate-with-snyk/notification-and-ticketing-systems-integrations/jira-integration.md#create-a-jira-issue">Create a Jira issue</a>.</p></td></tr><tr><td><strong>Ignore issues</strong></td><td><p>Configure Snyk to ignore suggested fixes for an issue to suppress specific warnings. For example, you may have deliberately used hard-coded passwords to test your routines in test code, or you are aware of an issue but have decided not to fix it.</p><p></p><p>See <a href="../../manage-risk/prioritize-your-issues/ignore-issues/">Ignore issues</a>.</p></td></tr><tr><td><strong>Exclude files from the import process</strong></td><td><p>Check for <code>DeepCode/Snyk</code> ignore files <code>.gitignore</code> <code>.dcignore</code> and read them if they exist. Using the information in these files, Snyk filters to identify only the files with <a href="../../getting-started/supported-languages-and-frameworks/">the supported extensions</a> in the Project directory and not above the current Project directory. Snyk Code bundles these files that are smaller than 4 MB and sends them to Snyk. <code>,gitignore</code> exclusions are honored by the <code>snyk code test</code> CLI command.</p><p></p><p>See also <a href="import-repository-to-snyk/excluding-directories-and-files-from-the-import-process.md">Exclude directories and files from the import process</a>.</p></td></tr></tbody></table>

## Deployment

| Deployment                                                           | Description                                                                                                                                                                                                                                                                             |
| -------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Full SaaS solution**                                               | <p>Get the most out of Snyk Code with a native Git repository integration, easy onboarding, and continuous updates. <br><br>See <a href="configure-snyk-code.md">Configure Snyk Code</a>.</p>                                                                                           |
| **SaaS with a self-hosted Git server (requires Snyk Broker)**        | <p>For customers with SCMs that are not publicly accessible from the internet and want to connect Snyk Code with their local self-hosted SCM. <br><br>See <a href="../../enterprise-configuration/snyk-broker/snyk-broker-code-agent/">Snyk Broker - Code Agent</a>.</p>                |
| **Local no-upload implementation (requires Snyk Code Local Engine)** | <p>For customers with a stricter upload policy. </p><p><br>This deployment method requires more maintenance and receives slower updates than the SaaS options, but it does not require any code upload. <br><br>See <a href="snyk-code-local-engine.md">Snyk Code Local Engine</a>.</p> |

## AI Engine

Snyk Code is powered by a semantic, AI-based analysis engine and can analyze the following in your code:

* **API usage:** Identifies multiple potential issues, including API misuses, null dereferences, and type mismatches, by modeling the use of memory in variables and references. This mechanism can also identify the use of insecure functions.
* **Coding issues:** Finds problems such as dead code, branches that are predefined, and branches having the same code on each side.
* **Control flow:** Identifies null dereference or race conditions by modeling each possible control flow in the application.
* **Data flow:** Follows the flow of data within the application from the source to the sink. Combined with AI-based learning of external insecure data sources, data sinks, and sanitation functions, this enables a strong taint analysis.
* **Hardcoded secrets:** Hardcoded secrets detection rules are invoked during SAST scans but do not act as a standalone secrets scanning tool. For an enhanced secrets solution, see our partnership with [GitGuardian](https://snyk.io/blog/supercharge-app-security-code-to-cloud/).

<figure><img src="../../.gitbook/assets/Introduction - AI Engine - Hardcoded secrets.png" alt="Hardcoded secret found"><figcaption><p>Hardcoded secret found</p></figcaption></figure>

* **Point-to analysis:** Identifies multiple potential issues, including buffer overruns, null dereferences, and type mismatches, by modeling memory use in variables and references.
* **Type inference:** Determines the initial type and its changes. This is of special interest for dynamically typed languages.
* **Value ranges:** Infers possible values for variables used to call functions to track off-by-one errors in arrays, division-by-zero errors, and null dereferences.

## Supported integrations

* **IDE**&#x20;
  * [JetBrains plugins](../../integrate-with-snyk/use-snyk-in-your-ide/jetbrains-plugins/)
  * [Visual Studio Code extension](../../integrate-with-snyk/use-snyk-in-your-ide/visual-studio-code-extension/)
  * [Visual Studio extension](../../integrate-with-snyk/use-snyk-in-your-ide/visual-studio-extension/)
* **Git repository**: With repository monitoring integration, you can actively manage your Code Projects using the existing native import flow and tools at your disposal. You can view and prioritize security issues discovered in your source code. Additionally, you can initiate a retest of any Project and examine the historical snapshots to track changes over time. See [Supported Git repositories](../../integrate-with-snyk/git-repository-and-ci-cd-integrations-comparisons.md).

{% hint style="info" %}
Snyk Code analysis can be applied to every pull request you create in your Git repository before you merge it into the target branch. See [PR Checks](../run-pr-checks/).
{% endhint %}

* **CLI and CI/CD**: [Using the CLI ](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/)helps you find and fix security flaws in your code on your local machine or in your CI/CD.
* **APIs and extensibility**: Query Code Projects and issues using the [Snyk REST API](https://apidocs.snyk.io/#overview).
* **Notifications**: [Integrate with Jira](../../integrate-with-snyk/notification-and-ticketing-systems-integrations/jira-integration.md) to export data to Jira issues.

## Supported languages

Snyk Code supports many [languages and frameworks](../../getting-started/supported-languages-and-frameworks/).

## What's next?

* [Create or log in to a Snyk account](../../getting-started/quickstart/create-or-log-in-to-a-snyk-account.md)
* [Set up an integration](../../getting-started/quickstart/set-up-an-integration.md)
* [Import a Project](../../getting-started/quickstart/import-a-project.md)
* [Configure Snyk Code in Snyk Web UI](configure-snyk-code.md)

