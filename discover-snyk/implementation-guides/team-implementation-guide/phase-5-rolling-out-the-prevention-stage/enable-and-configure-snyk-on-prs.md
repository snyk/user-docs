---
description: How to enable and configure Snyk PR and MR Checks to block new vulnerabilities in code changes
nav_context: agnostic
---

# Enable and configure Snyk on PRs

[Snyk PR Checks](https://docs.snyk.io/scan-fix-and-prevent/prevent/pull-request-checks) prevent new security issues from entering your codebase by automatically scanning code changes when you submit a pull request (PR) or merge request (MR). PR Checks are available for open-source vulnerabilities, license compliance issues, and vulnerabilities in your own code.

If you import Projects through a source control integration, Snyk Open Source and Snyk Code PR Checks are a good place to start introducing gating.

{% hint style="info" %}
Snyk recommends announcing changes before you roll them out. For examples of how to message your developers, visit [Announce prevention measures](../../enterprise-implementation-guide/automate-prevention-measures.md#announce-prevention-measures).
{% endhint %}

## Implement Snyk Open Source PR Checks

Several features help you introduce PR Checks gradually to avoid friction with your development teams:

* Fail conditions: You can control whether the test fails if the PR itself is adding a dependency with issues (most common) or if the repository as a whole has any issues.

You can also customize what constitutes a "failed test." By default, the test does not filter based on severity or fixability, so PR tests can fail regularly. For Snyk Open Source, you can customize the criteria to fail the test:

* **Only fail for high or critical severity issues**
* **Only fail when the issues found have a fix available**

When you first enable this feature, Snyk recommends selecting both check boxes so that a test fails if Snyk finds an issue that is High or Critical severity and fixable. In this case, encourage the developer to fix the issue before proceeding.

PR tests are optional by default, so even if the test fails, the developer can continue and merge the PR. You configure whether a PR test is optional or blocking in your source control management platform, for example, with branch protection rules in GitHub.

## Implement Snyk Code PR Checks

When you enable static analysis of code changes, you can set the fail criteria to **High** (the highest severity).

When you first enable this feature, Snyk recommends selecting **High**. In this case, encourage developers to fix the issue before proceeding. As your codebase stabilizes and you work through the backlog, you can change the fail criteria to a lower severity (**Medium** or **Low**) to address broader issues or to match your internal policies.

## Use PR Checks for a phased rollout

It is common to have a phased rollout of these features. Using PR Checks as an example:

* Initially, run Snyk tests as optional checks, which you set in your source control settings. The results are displayed, but the developer is not blocked from merging the PR.
* Over time, as developers adapt to seeing these results and begin addressing the high-severity issues proactively, you can choose to start blocking PRs from being merged if there are any new highest severity issues or, in the case of Snyk Open Source, if a fix is available.

This phased rollout helps to decrease friction between your security and development teams.
