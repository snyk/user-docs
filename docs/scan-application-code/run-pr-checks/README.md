# Run PR Checks

Use the **Snyk PR Checks** feature to prevent new security issues from entering your codebase, by automatically scanning code changes in real-time, as soon as a developer submits a PR (Pull Request).

### Introduction

**Snyk PR Checks** allows developers to auto-scan their PRs for issues, before merging their code. Snyk acts as an automated pseudo-team member (an “expert security reviewer”), to find any potential issues, leaving review notes on your PR, before the code is committed.

PRs are the points in the development process where code reviews happen. So PR Checks allows security scanning to meet the developers in their native environment, seamlessly integrating with developer workflows, empowering your dev teams, and helping to prevent security issues occurring in deployed code.

<figure><img src="../../.gitbook/assets/Screenshot 2022-09-20 at 11.27.44.png" alt=""><figcaption></figcaption></figure>

The PR Checks feature is available for [Snyk Open Source](pr-checks-for-snyk-open-source/) and [Snyk Code](pr-checks-for-snyk-code/).

#### Test the change

Snyk PR Checks allows you to test a **change** to the current codebase - to see if that change introduces a problem. This change testing makes it easier to maintain the security of your codebase, on an ongoing basis.

For developers, change-related flaws are naturally far more relevant and easy to fix, making secure code much easier to roll out. You can detect security issues at an early stage in the development process, see the test results immediately after you write new code, and find and fix issues as they emerge, all in your native workflow.

{% hint style="info" %}
You can also use Snyk products such as Snyk Open Source to scan the overall state of your current codebase - finding problems in your current repos, regardless of where the problem came from originally. This state testing helps address your backlog of vulnerabilities in your existing code.
{% endhint %}

#### Testing “before” and “after”

Snyk PR checks runs live tests of the “before and after” branch with the PR, and fails only if the new branch has more issues. This allows you to cater for problems that occurred externally since the last scan (for example, new vulnerabilities introduced externally); Snyk PR Checks only identifies problems with your specific code change.

### How it works

Snyk PR Checks automatically scans your source code PRs after you create them, including each additional commit you make to that PR. PR Checks display scan results in the SCM (either passed or failed).

If security vulnerabilities are found, Snyk PR Checks automatically fails the scanned PRs, preventing new security issues from entering into your codebase.

#### Example: Snyk Code

<figure><img src="../../.gitbook/assets/image2.png" alt=""><figcaption></figcaption></figure>

#### Example: Snyk Open Source

<figure><img src="../../.gitbook/assets/image3.png" alt=""><figcaption></figcaption></figure>

### More details

* Scan PRs for security and license issue: [PR Checks for Snyk Open Source](pr-checks-for-snyk-open-source/).
* Scan your application code PRs: [PR Checks for Snyk Code](pr-checks-for-snyk-code/).
