# Enable and configure Snyk on PRs

## Use PR Checks to introduce gating

[Snyk Pull Request(PR)/Merge Request (MR) Checks](../../../scan-with-snyk/pull-requests/pull-request-checks/) allow you to prevent new security issues from entering your codebase, by automatically scanning code changes when you submit a pull request (PR). PR Checks are available for open-source vulnerabilities, license compliance issues, and your own code issues.

If you import Projects through a source control integration, then Snyk Open Source and Snyk Code PR Checks is a good place to start introducing gating.&#x20;

{% hint style="info" %}
Snyk recommendes announcing the changes prior to rolling out the changes. See [Announcement Templates](../../enterprise-implementation-guide/phase-6-rolling-out-the-prevention-stage/announcement-templates-for-prevention.md) for examples of how to message your developers.
{% endhint %}

### Implementing Open Source PR Checks

There are a number of different features available that can be used to help you gradually introduce the feature to avoid friction with your development teams:

* Fail conditions: You can control whether the test will "fail" if the PR itself is adding a dependency with issues (most common) or if the repository as a whole has any issues.

The criteria of what constitutes a "failed test" can also be customized. By default, the test does not filter based on severity or fixability, which can mean that PR tests will regularly fail. For Snyk Open Source you can customize what the criteria are to fail the test:

* **Only fail for high or critical severity issues**
* **Only fail when the issues found have a fix available**

When you first enable this feature, Snyk suggests ticking both of these boxes so that a test would fail if a **High or Critical** and **fixable** issue is found. In this case, you would want to encourage to developer to fix the issue before proceeding.

These PR tests are optional by default, meaning that even if the test fails, the developer may be able to continue and merge the PR. Controlling whether a PR test is optional or blocking is configured within your source control management platform, such as GitHub’s branch protection rules.

### Implementing Code PR Checks

When enabling static analysis of code changes and therefore new vulnerabilities, you can select the fail criteria to be **High** (the highest severity).

When you first enable this feature, Snyk suggests selecting **High** if an issue is found. In this case, you would want to encourage developers to fix the issue before proceeding. As your code base stabilizes and you have worked through the backlog, you can change the fail criteria to lower the severity (to either **Medium** or **Low)** in order to address broader issues or to match your internal policies.

### Using PR Checks for a phased rollout

It is common to have a phased rollout of these features. Using PR checks as an example:

* You may initially run Snyk tests and set, via your source control settings, as optional checks. The results are displayed, but the developer is not blocked from merging the PR.&#x20;
* Over time, as developers adapt to seeing these results and begin addressing the high issues proactively, you can choose to start blocking PRs from being merged if there are any new highest severity issues or, in the case of Snyk Open Source, if a fix is available.&#x20;

This phased rollout helps to decrease friction between your security and development teams.
