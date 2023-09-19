# Configure PR Checks

{% hint style="info" %}
PR Checks for Snyk Code is in Closed Beta and only available for Enterprise plans.
{% endhint %}

{% hint style="warning" %}
A PR Check is counted as a test within your Organization's test count, including automatic checks of new commits in an open pull request. See[ What counts as a test](https://support.snyk.io/hc/en-us/articles/360000925418-What-counts-as-a-test-). The number of tests allowed is determined by the [pricing plans](../../more-info/plans.md).
{% endhint %}

## Prerequisites for automated PR Checks

To check for open-source and licensing issues and code security, ensure that you have established the following:

* You have the Group Admin role so you have access to all integration settings. See [Member roles](../../snyk-admin/manage-permissions-and-roles/manage-member-roles.md).
* You have [set up a Git repository integration](../../integrations/git-repository-scm-integrations/). For help, see the Snyk Training course [Source code manager configuration](https://training.snyk.io/learn/course/source-code-manager-configurations).
* [Import a Project](../../getting-started/quickstart/import-a-project.md) to have a working Git repository.
* For code security (Snyk Code), meet all of the above conditions and then contact your Snyk representative to enable the feature for you.&#x20;

{% hint style="info" %}
PR Checks rely on webhooks from the SCM. Integration scope must include the ability to create webhooks.
{% endhint %}

## Types of Snyk scans supported

You can analyze the changed code with PR Checks as follows:

* (Beta) Snyk Code: Source code changes result in a vulnerability that exceeds a specified threshold. A full scan of the repository is done to determine if there are new vulnerabilities.&#x20;
* Snyk Open Source: Snyk analyzes dependency manifest or supported files for known security vulnerabilities that meet a threshold, such as exceeding severity, or checks to determine whether a fix is available.&#x20;
* Open Source license check: Snyk validates package licenses against the configured policy for license policy violations.

PR Checks also support all programming languages and frameworks supported by the Snyk Code and Open Source engines. For more information, see programming language support for [Snyk Code](../../scan-applications/snyk-code/snyk-code-language-and-framework-support.md) and [Open Source](../../scan-applications/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/).

## How configuration of PR Checks works

You can configure PR Checks either [at the Integration level](configure-pr-checks.md#configure-pr-checks-at-the-integration-level) for your Snyk Organization or [for specific Snyk Projects](configure-pr-checks.md#configure-pr-checks-at-the-project-level) in an Organization.

* In your Organization, you can have multiple repository integrations, but the feature works only for those integrations that have PR Checks configured.
* At the Project level, the settings are inherited from the integration by default, but you can configure custom settings.

## Configure PR Checks at the integration level

Configure PR Checks for a specific Git repository you have already integrated with Snyk, such as a GitHub repository.

The configuration settings apply to all Projects in that Organization. You can also extend the configuration to Projects with custom settings.

1. In the Snyk Web UI, navigate to **Settings >** **Integrations** and select your connected source code manager to open the settings configuration.
2. To check for code issues, configure and save the following changes:

* [ ] **Code Analysis**: Enable this option to fail the PR on new vulnerabilities detected in your Git repository. If the severity is higher than your threshold, the PR is not merged into the main branch.&#x20;
* [ ] **Fail conditions**: Select the severity threshold at which the PR fails. For example, if you select **Medium**, the PR fails on issues found at this level or higher, while it is merged for **Low** severity issues.

<figure><img src="../../.gitbook/assets/enable analyze code.png" alt="PR check settings to analyze code issues."><figcaption><p>PR check settings to analyze code issues</p></figcaption></figure>

{% hint style="info" %}
If you cannot see the **Code Analysis** section, ensure that your user has the Organization Admin role assigned and that the feature is enabled for Snyk Code. See the  [Prerequisites](configure-pr-checks.md#prerequisites).
{% endhint %}

3. To check for open-source and licensing issues, configure and save the following changes:

* [ ] **Open Source Security & Licenses**: Enable this option to fail the PR when open-source and licensing issues found in the proposed changes exceed your specified severity threshold. If the severity is higher than your threshold, the PR is not merged into the main branch.
* [ ] **Fail conditions**: Select one of the following PR failure conditions based on the security issues distribution.
  * [ ] **Only fail when the PR is adding a dependency with issues**: Set this condition to fail PR when there is at least one dependency with security issues.
  * [ ] **Fail if the repo has any issues**: Set this condition to fail a PR for any security issues found in the Git repository.
* [ ] **Only fail for high or critical severity issues**: Select additional failure conditions based on the severity threshold.
* [ ] **Only fail when the issues found have a fix available**: Set this condition on for the PR check to fail when the PR introduces new vulnerabilities that are fixable by Snyk. PR checks don't fail on newly introduced vulnerabilities if Snyk is unable to fix them.&#x20;

When switched on, this will cause the PR check to fail when the PR introduces new vulnerabilities that are fixable by Snyk. PR checks will not fail on newly introduced vulnerabilities if Snyk is unable to fix them.

<figure><img src="../../.gitbook/assets/Screenshot 2023-04-28 at 12.06.13 (1).png" alt="Pull request check settings to analyze open-source and licensing issues." width="563"><figcaption><p>PR check settings to analyze open-source and licensing issues</p></figcaption></figure>

4. Either click **Save** to save the changes, select the Save dropdown and click **Apply changes to all overridden Projects** to extend the current configuration to Projects with custom settings. For more information, see [Configure PR Checks at the Project level](configure-pr-checks.md#configure-pr-checks-at-the-project-level)).

## Configure PR Checks at the Project level

You can configure PR Checks to work only for specific Projects:

1. Navigate to **Projects** and expand the [target](../../snyk-admin/introduction-to-snyk-projects/#target) containing your Project.
2. Click a Project name to open it. Based on the Project type, you can choose the following:

* **package.json** to check for open-source and licensing issues.
* **Code analysis** to check for security issues in your code.

<figure><img src="../../.gitbook/assets/configure_pr_checks_project_level.png" alt="Project overview"><figcaption><p>Project overview</p></figcaption></figure>

6. Navigate to **Settings.**
7. On the left side, select your integration. For this example, GitHub has been integrated with Snyk.
8. Configure Project settings based on your Project type:

<details>

<summary>Configure for open source and licensing issues (click to expand)</summary>

1. In **Snyk test for pull request** select **Custom** to configure the settings.
2. Enable the option to fail the PR when open-source and licensing issues found in the proposed changes exceed your specified severity threshold.
3. Configure the following settings:

* [ ] **Fail conditions**: Select one of the following PR failure conditions based on the security issues distribution.
  * [ ] **Only fail when the PR is adding a dependency with issues**: Set this condition when there is at least one dependency with security issues.
  * [ ] **Fail if the repo has any issues**: Set this condition for any security issues found in the Git repository.
* [ ] **Only fail for high or critical severity issues**: Select additional failure conditions based on the severity threshold.
* [ ] **Only fail when the issues found have a fix available**: Set this condition on or more if the issues found have a dependency or package with a version in which the issue is fixed.

4. **Update Snyk pull request settings** to save changes.

</details>

<details>

<summary>Configure for code analysis (click to expand)</summary>

1. In **Snyk Code for pull request** select **Custom** to configure the settings.
2. Enable this option to fail the PR when the security issues found in the proposed changes exceed your specified severity threshold.
3. Configure the following settings:

* [ ] **Minimal severity to fail PR check**: Select the severity threshold at which the PR fails. For example, if you select **Medium**, the PR fails on issues found at this level or above, while it is merged for **Low** severity issues.

4. **Update Snyk pull request settings** to save changes.

</details>

###
