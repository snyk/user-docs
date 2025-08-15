# Pull Request checks

## Introduction to automated security scans with PR checks

The [Snyk PR Checks](configure-pull-request-checks.md) feature allows developers to auto-scan their PRs for issues before merging their code. As Snyk acts as an automated pseudo-team member (an “expert security reviewer”), it finds potential issues, leaving review notes on your PR before the code is committed.

While PRs are the points in the development process where code reviews happen, PR Checks allow security scanning to be integrated with developer workflows, empowering dev teams and helping to prevent security issues from occurring in deployed code.

As part of the [Snyk Pull Request Experience](pull-request-experience.md), the Snyk Issue Summary Comment feature allows you to receive a PR comment on each pull request, summarizing the most recent PR Check results according to the type of check and the count of findings by severity. This reduces context switching by displaying additional information about the PR Check scan results in the pull request.

## **Test the change**

The Snyk PR Checks feature allows you to test a change to the current codebase to see if that change introduces a problem. This change testing makes it easier to maintain the security of your codebase on an ongoing basis.

For developers, change-related flaws are relevant and easy to fix, and fixing change-related flaws rather than accumulated flaws makes rolling out secure code easier. You can detect security issues early in the development process, see the test results immediately after you write new code, and find and fix issues as they emerge, all in your native workflow.

## **Testing “before” and “after”**

The Snyk PR Checks feature runs live tests of the “before and after” branch with the PR and fails only if the new branch has more issues. This allows you to address problems that have been introduced since the last scan, for example, new vulnerabilities introduced externally. Snyk PR Checks are triggered by a change in your code, and find issues across the entire repository. Thus,  a PR check finds issues in your code as well as other issues introduced since the last Snyk scan.Use the Snyk PR Checks feature to prevent new security issues from entering your codebase by automatically scanning code changes in real time as soon as you submit a pull request (PR) in your source code manager (SCM).

## Why use PR Checks

| Benefits                                      | Details                                                                                                                                                                                                                                                     |
| --------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Automatically scanned pull request            | Scan code changes in real-time as soon as a pull request is submitted, to catch potential issues before they go into production.                                                                                                                            |
| Results displayed in your source code manager | Make use of security reviews and notes left by Snyk on your pull requests.                                                                                                                                                                                  |
| Code change testing for security issues       | Test changes to your codebase for any security issues, ensuring that your code stays secure over time.                                                                                                                                                      |
| Branch testing                                | <p>Test branches before and after implementing changes to fail only if the new branch has introduced issues. Prevent merging pull requests with failed security issues. <br>Note that Snyk monitors all pull requests made to the monitored repository.</p> |

## What to test for

You can analyze PR Checks results in your SCM to test for dependency and licensing issues using [Snyk Open Source](../../snyk-open-source/) and code security using [Snyk Code](../../snyk-code/).

## How PR checks work

The following diagram explains how Snyk Checks PRs in your development workflow.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-09-20 at 11.27.44.png" alt="Where Snyk checks for pull requests in the development workflow."><figcaption><p>Where Snyk checks for pull requests in the development workflow</p></figcaption></figure>

PR checks proceed as follows:

1. A developer creates a pull request (PR) in an SCM integrated with Snyk.
2. A webhook is triggered from the SCM to Snyk
3. Snyk automatically scans the code changes in the PR for issues.
4. Snyk leaves security reviews and notes on the PR.
5. The developer can view the PR Checks results and fix identified issues before merging the code.
6. The PR Checks results appear as **Passed** or **Failed** directly in the SCM, preventing PRs from being merged with security issues.

For more information on working with PR Checks, see the following pages:

* [Configure PR Checks](configure-pull-request-checks.md)
* [Pull Request Experience](pull-request-experience.md)
* [Analyze PR Checks results](analyze-pr-checks-results.md)
* [Troubleshoot PR Checks](troubleshoot-pr-checks.md)
