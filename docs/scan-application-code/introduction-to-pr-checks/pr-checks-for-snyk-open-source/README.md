# PR Checks for Snyk Open Source

{% hint style="info" %}
Snyk Code also provides PR Checks - see [PR Checks for Snyk Code](../pr-checks-for-snyk-code/).
{% endhint %}

### Introduction to PR Checks

Snyk completes a live test before and after on the branch with the pull request. This means the build only fails if a vulnerability is detected.

| Introduced Vulnerability                                                                                                                                                                           | Main Branch Vulnerability            |                          PR Check Result |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------ | ---------------------------------------: |
| <mark style="color:red;">Yes</mark>                                                                                                                                                                | <mark style="color:red;">Yes</mark>  |   <mark style="color:red;">Fail ❌</mark> |
| <mark style="color:red;">Yes</mark>                                                                                                                                                                | <mark style="color:green;">No</mark> |   <mark style="color:red;">Fail ❌</mark> |
| <mark style="color:green;">No</mark>                                                                                                                                                               | <mark style="color:red;">Yes</mark>  | <mark style="color:green;">Pass ✅</mark> |
| <mark style="color:green;">No</mark>                                                                                                                                                               | <mark style="color:green;">No</mark> | <mark style="color:green;">Pass ✅</mark> |
| <mark style="color:red;">Yes</mark> with[ _"Mark as successful in SCM"_](../pr-checks-for-snyk-code/viewing-and-working-with-the-pr-check-results-on-the-snyk-web-ui.md#\_ref105582006) _by admin_ | _Any (overridden)_                   | <mark style="color:green;">Pass ✅</mark> |

<figure><img src="../../../.gitbook/assets/scm-ci-cid.png" alt=""><figcaption><p>Testing PRs for vulnerability introduction falls in the CI/CD pipeline.</p></figcaption></figure>

There are two main troubleshooting situations to diagnose for Snyk's PR checks.

1. Passed when it should have failed: submit [Vulnerability Disclosure](https://snyk.io/vulnerability-disclosure/).
2.  Failed when it should have passed: Check security check output.

    <figure><img src="../../../.gitbook/assets/security-check.png" alt=""><figcaption></figcaption></figure>

> **Hot tip**: Mimic a PR check like this (changes only, not state of project)[ Snyk CLI](../../../snyk-cli/test-for-vulnerabilities/advanced-failing-of-builds-in-snyk-cli.md#fail-current-build-only-if-new-vulnerabilities-are-being-introduced).

When SAST issues are found in your PR, Snyk Code provides you with additional details on each detected issue and offers you fix examples to assist you in developing secure code. By clicking the discovered issues or the link next to them in your SCM, you can open the Snyk Web UI, and view the full details of each discovered vulnerability in your PR:

If you want to pass PRs that automatically failed due to vulnerabilities that were found in them, Snyk Code also enables you to [mark failed PRs as successful](../pr-checks-for-snyk-code/viewing-and-working-with-the-pr-check-results-on-the-snyk-web-ui.md#\_ref105582006) via the Snyk Web UI. Once you click the **Mark as successful in SCM** button on the Web UI, your failed PRs are considered as successful in the SCM, and can be merged into the target branch:

The Automatic PR Checks feature is applied only to repositories imported to Snyk from the integrated SCM. However, after the initial import, any new file or folder added to the imported repositories is included in the automatic PR Checks. The Automatic PR Checks feature can be enabled for your integrated SCM on the level of an entire Organization or on the level of a specific Project.

{% hint style="info" %}
Every PR check is considered as a “test” in the test count of the related Organization. New commits to on open PR branch are also checked automatically, and therefore these commit checks are also counted as “tests”.
{% endhint %}

#### Using PR Checks: workflow

The workflow of using the PR Checks feature is the same as it is for Snyk Code:

1. [Integrating your required SCM with Snyk.](../../snyk-code/getting-started-with-snyk-code/activating-snyk-code-using-the-web-ui/step-2-integrating-your-source-control-system-with-snyk-code.md)
2. [Importing from the integrated SCM to Snyk the repositories whose PRs you want Snyk Code to check.](../../snyk-code/getting-started-with-snyk-code/activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/)
3. [Enabling the Automatic PR Checks feature for Snyk Code in your Snyk-SCM integration on the level of an Organization or a specific Project.](../pr-checks-for-snyk-code/enabling-pr-checks-for-snyk-code.md)
4. [Viewing the results of the PR check in your SCM after the creation of a new PR.](../pr-checks-for-snyk-code/viewing-the-pr-checks-in-your-scm.md)
5. [Viewing additional details and fix examples for the issues that were discovered in the PR on the Snyk Web UI.](../pr-checks-for-snyk-code/viewing-and-working-with-the-pr-check-results-on-the-snyk-web-ui.md)
