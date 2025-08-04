# Enable and configure Snyk on PRs

## Use PR Checks to introduce gating

[Snyk Pull Request(PR) or Merge Request (MR) Checks](../../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/) allow you to prevent new security issues from entering your codebase by automatically scanning code changes when you submit a pull request (PR). PR Checks are available for open-source vulnerabilities, license compliance issues, and your own code issues.

If you import Projects through a source control integration, then Snyk Open Source PR Checks is a good place to start introducing gating.&#x20;

{% hint style="info" %}
It is recommended that you announce the changes prior to rolling them out. See [Announcement Templates](announcement-templates-for-prevention.md) for examples of how to message your developers.
{% endhint %}

## Implement PR Checks

You can use fail conditions to help you gradually introduce PR Checks to avoid friction with your development teams.

Fail conditions allow you to control whether the test will fail if the PR itself adds a dependency with issues, the most common circumstance, or if the repository as a whole has issues.

The criteria for what constitutes a failed test can also be customized. By default, the test does not filter based on severity or fixability, which can mean that PR tests will regularly fail. You can customize  the criteria to fail the test:

* Fail only for High or Critical severity issues; function available for Snyk Open Source and Snyk Code
* Fail only when the issues found have a fix available; function available for Snyk Open Source

When you first enable PR Checks, Snyk suggests ticking both of these boxes so that a test would fail if a High or Critical and fixable issue is found. In this case, you would want to encourage to developer to fix the issue before proceeding.

These PR tests are optional by default, meaning that even if the test fails, the developer may be able to continue and merge the PR. Controlling whether a PR test is optional or blocking is configured within your source control management platform, such as GitHub’s branch protection rules.

## Use PR Checks for a phased rollout

It is common to have a phased rollout of Snyk features. Using PR checks as an example:

* You may initially run Snyk tests and set them through your source control settings as optional checks. The results are displayed, but the developer is not blocked from merging the PR.&#x20;
* Over time, as developers adapt to seeing these results and begin addressing the critical issues proactively, you can choose to start blocking PRs from being merged if there are any new High or Critical severity issues, or in the case of Snyk Open Source, if a fix is available.&#x20;

This phased rollout helps to decrease friction between your security and development teams.

## Additional information

The webinar [Dev-First Prevention Strategies Using PR Checks](https://www.youtube.com/watch?v=6x33EJW_d_E) covers PR checks in more detail and includes an example of how you can gradually introduce this feature.
