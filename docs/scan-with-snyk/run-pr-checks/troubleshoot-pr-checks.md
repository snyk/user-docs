# Troubleshoot PR Checks

{% hint style="info" %}
If you use `###` in the description of the PR , it will be blocked and the PR check will not take place
{% endhint %}

## General troubleshooting for PR Checks

The following table lists general issues with PR Checks and how to address them.

| Scenario                                                              | Description                                                                                                       | Action                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| --------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| PR Check not triggered.                                               | The repository is imported to Snyk, but when a PR is raised it does not trigger a PR Check.                       | <ol><li>Check the <a href="configure-pr-checks.md">Project and Integration settings</a> to make sure PR checks are configured. </li><li>Check the Project Settings integration page for a banner advising you that a webhook is <a data-footnote-ref href="#user-content-fn-1">absent</a>. <br><br>Double-check integration permission scopes (see <a href="../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/">Git repositories</a>). <br><br>If you still cannot create a webhook, <a href="https://support.snyk.io/hc/en-us/requests/new">contact support</a>.</li></ol> |
| PR Check is expected but does not run.                                | The PR check is listed in the Git repository (SCM) as expected but never completes.                               | <p>This issue is generally caused by a Branch Protection rule requiring the PR check. If the Project has been disabled or removed from Snyk, the PR check will not run, but the branch protection rule is still in force until removed or edited. <br><br>Check for Branch Protection rules and confirm that the Project is imported and active. </p>                                                                                                                                                                                                                                            |
| Multiple Security and Licence PR Checks run on a single Pull Request. | When a PR is submitted, multiple Snyk PR Checks of the same type run against it, possibly with different results. | <p>If a repository is imported into multiple Snyk Organizations, PR checks will run on the repository for any configured Organization.<br><br>Check the name of the PR check as it includes the Organization name against which the check is run. Alternatively, selecting the PR Check details will take you to the results for the relevant Organization.</p>                                                                                                                                                                                                                                  |

## Open-source and licensing checks

If you come across false positive or false negative results, you can take action to diagnose and report the issue.&#x20;

| Scenario       | Description                                                                                             | Action                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| -------------- | ------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| False positive | The result is marked as Failed by Snyk because it has identified an issue that does not actually exist. | <p><a href="https://support.snyk.io/hc/en-us/requests/new">Contact support</a> to update the dependency version if Snyk has misidentified an issue for a package version.</p><p>If you want to push the changes forward, you can mark the result as successful. For more einformaiton, see <a href="analyze-pr-checks-results.md#example-fix-dependency-issues-with-pr-checks">Example: fix dependency issues with PR Checks</a>).</p> |
| False negative | The result is marked as Passed by Snyk because it failed to detect an issue that actually exists.       | <p>To address the issue, you can submit a <a href="https://snyk.io/vulnerability-disclosure/">vulnerability disclosure</a>.<br><br>If Snyk did not detect the vulnerability due to a misidentified package or the absence of code trace, <a href="https://support.snyk.io/hc/en-us/requests/new">contact support</a>.</p>                                                                                                              |

## Code analysis checks

The following table lists code analysis errors and how to address them.

| Error                                               | Description                                                                                                                      | Action                                                                                                                                                          |
| --------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Failed to start code analysis.                      | <p>Error causes:</p><ul><li>The PR Checks cannot be created in the database.</li><li>The commit status cannot be sent.</li></ul> | Wait a few minutes, then [try again](troubleshoot-pr-checks.md#re-run-pr-checks-results).                                                                       |
| Could not complete the PR analysis.                 | The PR Checks result has an unexpected status.                                                                                   | Wait a few minutes, then [try again](troubleshoot-pr-checks.md#re-run-pr-checks-results) or [mark as successful](troubleshoot-pr-checks.md#mark-as-successful). |
| Failed to analyze code.                             | <p>Error causes:</p><ul><li>The analysis cannot be completed.</li><li>The commit status cannot be sent.</li></ul>                | Wait a few minutes, then [try again](troubleshoot-pr-checks.md#re-run-pr-checks-results) or [mark as successful](troubleshoot-pr-checks.md#mark-as-successful). |
| Upstream rate limit triggered while analyzing code. | The Git server rate limit has been reached and the repository cannot be read.                                                    | Wait a few minutes, then [try again](troubleshoot-pr-checks.md#re-run-pr-checks-results) or [mark as successful](troubleshoot-pr-checks.md#mark-as-successful). |
| No valid credentials to perform code analysis.      | The personal access token or OAuth is not recognized or the user access is not provisioned.                                      | Revise your configuration on the Git repository side for any credential issues.                                                                                 |

## What to do if there are errors

### Re-run PR Checks results

To re-run PR Checks results:

* Create an empty commit for example with `git commit –allow-empty`&#x20;
* Create a new commit with a fix or additional capability
* Close and re-open the pull request in your connected Git repository (for example, GitHub).

### Mark as successful

Provide specific users or roles with the capability to pass the PR Check when errors happen. This can be done through the Snyk link in the PR Check and Marking as successful.&#x20;

[^1]: [andrei.onciu](https://app.gitbook.com/u/Ge4Ptulj2Hb65sgVH1Uanpl5nZv2 "mention") the link here is not&#x20;
