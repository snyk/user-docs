# Run PR Checks

Use Snyk PR Checks to prevent new security issues from entering your codebase by automatically scanning code changes in real time as soon as you submit a pull request (PR) in your source code manager (SCM).

## Why use PR Checks

| Benefits                                      | Details                                                                                                                                                                                                                                                 |
| --------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Automatically scanned pull request            | Scan code changes in real-time as soon as a pull request is submitted, to catch potential issues before they go into production.                                                                                                                        |
| Results displayed in your source code manager | Make use of security reviews and notes left by Snyk on your pull requests.                                                                                                                                                                              |
| Code change testing for security issues       | Test changes to your codebase for any security issues, ensuring that your code stays secure over time.                                                                                                                                                  |
| Branch testing                                | <p>Test branches before and after implementing changes to fail only if the new branch has introduced issues. Prevent merging pull requests with failed security issues. <br>Note that Snyk monitors all pull requests made to the monitored branch.</p> |

## What to test for

You can analyze PR Checks results in your SCM to test for dependency and licensing issues using [Snyk Open Source](../snyk-open-source/) and code security using [Snyk Code](../snyk-code/).

## How PR checks work

The following diagram explains how Snyk Checks PRs in your development workflow.

<figure><img src="../../.gitbook/assets/Screenshot 2022-09-20 at 11.27.44 (1) (1).png" alt="Where Snyk checks for pull requests in the development workflow."><figcaption><p>Where Snyk checks for pull requests in the development workflow</p></figcaption></figure>

PR checks proceed as follows:

1. A developer creates a pull request (PR) in an SCM integrated with Snyk.
2. A webhook is triggered from the SCM to Snyk
3. Snyk automatically scans the code changes in the PR for issues.
4. Snyk leaves security reviews and notes on the PR.
5. The developer can view the PR Checks results and fix identified issues before merging the code.
6. The PR Checks results appear as **Passed** or **Failed** directly in the SCM, preventing PRs from being merged with security issues.

For more information on how PR Checks work, see [Introduction to automated security scans with PR Checks](../../scan-applications/run-pr-checks/introduction-to-automated-security-scans-with-pr-checks.md).

For more information on working with PR Checks, see the following pages:

* [Configure PR Checks](configure-pr-checks.md)
* [Analyze PR Checks results](../../scan-applications/run-pr-checks/analyze-pr-checks-results.md)
* [Troubleshooting](../../scan-applications/run-pr-checks/troubleshooting.md)
