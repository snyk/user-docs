# Using Automatic PR Checks for securing your source code

{% hint style="info" %}
* The **Automatic PR Checks** feature in Snyk Code is currently in Closed Beta. If you would like to activate it in your organization, contact your CSM.
* The **Automatic PR Checks** feature in Snyk Code is currently NOT supported when using the Snyk Broker – Code Agent deployment method.
{% endhint %}

### Understanding the Automatic PR Checks feature

**Note**: The Automatic PR Checks feature is applicable to all supported SCMs.

The **Automatic PR Checks** feature enables you to apply Snyk Code test to every Pull Request you are creating in your integrated SCM, before merging it into the target branch. By using the Automatic PR Checks feature, you can detect security issues at an early stage in your development process, seeing the test results shortly after you write new code, and identifying and fixing issues as they emerge in your native workflow.

When the Automatic PR Checks feature is enabled, Snyk Code automatically scans your source code PRs once they are created, in search of security vulnerabilities. Every additional commit that will be made to the open PRs, will be scanned automatically as well. When the scan is completed, Snyk Code displays its findings in the SCM – either passed or failed. If security vulnerabilities are found, Snyk Code automatically fails the scanned PRs, thus preventing new security issues from entering into your code:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - GitHub - Some Checks Failed - Intro.png>)

When SAST issues are found in your PR, Snyk Code provides you with additional details on each detected issue, and offers you fix examples to assist you in developing secure code. By clicking the discovered issues or the link next to them in your SCM, you can open the Snyk Web UI, and view the full details of each discovered vulnerability in your PR:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Web UI - Intro (1).png>)

If you want to pass PRs that were automatically failed due to vulnerabilities that were found in them, Snyk Code also enables you to [mark failed PRs as successful](viewing-and-working-with-the-pr-check-results-on-the-snyk-web-ui.md#\_ref105582006) via the Snyk Web UI. Once you click the **Mark as successful in SCM** button on the Web UI, your failed PRs are considered as successful in the SCM, and can be merged into the target branch:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Mark as successful - On GitHub (1).png>)

The Automatic PR Checks feature is applied only to repositories that were imported to Snyk from the integrated SCM. However, after the initial import, any new file or folder that will be added in the SCM to the imported repositories, will be included in the automatic PR Checks. The Automatic PR Checks feature can be enabled for your integrated SCM on the level of an entire organization or on the level of a specific Project.

**Note**: Every PR check is considered as a “test” in the test count of the related organization. New commits to on open PR branch are also checked automatically, and therefore these commit checks are also counted as “tests”.

The workflow of using the Automatic PR Checks feature is as follows:

1. [Integrating your required SCM with Snyk.](../getting-started-with-snyk-code/activating-snyk-code-using-the-web-ui/step-2-integrating-your-source-control-system-with-snyk-code.md)
2. [Importing from the integrated SCM to Snyk the repositories whose PRs you want Snyk Code to check.](../getting-started-with-snyk-code/activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/)
3. [Enabling the Automatic PR Checks feature for Snyk Code in your Snyk-SCM integration on the level of an organization or a specific Project.](enabling-the-automatic-pr-checks.md)
4. [Viewing the results of the PR check in your SCM after the creation of a new PR.](viewing-the-pr-checks-in-your-scm.md)
5. [Viewing additional details and fix examples for the issues that were discovered in the PR on the Snyk Web UI.](viewing-and-working-with-the-pr-check-results-on-the-snyk-web-ui.md)
