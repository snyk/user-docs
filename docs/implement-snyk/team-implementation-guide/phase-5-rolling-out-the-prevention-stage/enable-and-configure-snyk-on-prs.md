# Enable and configure Snyk on PRs

## Use PR Checks to introduce gating

[Snyk Pull Request(PR)/Merge Request (MR) Checks](../../../scan-using-snyk/pull-requests/pull-request-checks/) allow you to prevent new security issues from entering your codebase, by automatically scanning code changes when you submit a pull request (PR). PR Checks are available for open-source vulnerabilities, license compliance issues, and your own code issues.

If you import Projects through a source control integration, then Snyk Open Source PR Checks is a good place to start introducing gating.&#x20;

{% hint style="info" %}
Snyk recommendes announcing the changes prior to rolling out the changes. See [Announcement Templates](../../enterprise-implementation-guide/phase-6-rolling-out-the-prevention-stage/announcement-templates-for-prevention.md) for examples of how to message your developers.
{% endhint %}

### Implementing PR Checks

There are a number of different features available that can be used to help you gradually introduce the feature to avoid friction with your development teams:

* Fail conditions: You can control whether the test will "fail" if the PR itself is adding a dependency with issues (most common) or if the repository as a whole has any issues.

The criteria of what constitutes a "failed test" can also be customized. By default, the test does not filter based on severity or fixability, which can mean that PR tests will regularly fail. You can customize what the criteria are to fail the test:

* Only fail for High or Critical severity issues (Snyk Open Source and Snyk Code)
* Only fail when the issues found have a fix available (Snyk Open Source).

When you first enable this feature, Snyk suggests ticking both of these boxes so that a test would fail if a ‘High or Critical’ and ‘fixable’ issue is found. In this case, you would want to encourage to developer to fix the issue before proceeding.

These PR tests are optional by default, meaning that even if the test fails, the developer may be able to continue and merge the PR. Controlling whether a PR test is optional or blocking is configured within your source control management platform, such as GitHub’s branch protection rules.

### Using PR Checks for a phased rollout

It is common to have a phased rollout of these features. Using PR checks as an example:

* You may initially run Snyk tests and set, via your source control settings, as optional checks. The results are displayed, but the developer is not blocked from merging the PR.&#x20;
* Over time, as developers adapt to seeing these results and begin addressing the critical issues proactively, you can choose to start blocking PRs from being merged if there are any new High or Critical severity issues, or in the case of Snyk Open Source, if there is a fix available.&#x20;

This phased rollout helps to decrease friction between your security and development teams.

## Also see

Here is a link to a webinar that covers PR checks in more detail and includes an example of how you can gradually introduce this feature: [Dev-First Prevention Strategies Using PR Checks](https://www.youtube.com/watch?v=6x33EJW\_d\_E).
