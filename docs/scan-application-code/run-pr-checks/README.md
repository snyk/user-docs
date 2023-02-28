# PR Checks

## Introduction

Use Snyk PR Checks to prevent new security issues from entering your codebase, by automatically scanning code changes in real-time, as soon as you submit a pull request (PR) in your source code manager (SCM).

## Why use PR Checks

| Benefits                                      | Details                                                                                                                                                                    |
| --------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Automatically scanned pull request            | Scan code changes in real-time as soon as a pull request is submitted, to catch potential issues before they go into production.                                           |
| Results displayed in your source code manager | Make use of security reviews and notes left by Snyk on your pull requests.                                                                                                 |
| Code change testing for security issues       | Test changes to your codebase for any security issues, ensuring that your code stays secure over time.                                                                     |
| Branch testing                                | Test branches before and after implementing changes to fail only if the new branch has introduced issues. Prevention of merging pull requests with failed security issues. |

## What to test for&#x20;

You can analyze PR Checks results in your SCM to test for dependency and licensing issues using [Snyk Open Source](../../products/snyk-open-source/), and code security using [Snyk Code](../../products/snyk-code/).

## How it works

The following diagram explains how Snyk Checks PRs in your development workflow.

<figure><img src="../../.gitbook/assets/PR checks diagram (development workflow).png" alt="Diagram showing where Snyk checks for pull requests in the development workflow."><figcaption><p>Pull request checks in the development workflow</p></figcaption></figure>

#### Step-by-step process

1. A developer creates a pull request (PR) in an SCM integrated with Snyk.
2. Snyk automatically scans the code changes in the PR for issues.
3. Snyk leaves security reviews and notes on the PR.
4. The developer can view the PR Checks results and fix identified issues before merging the code.
5. The PR Checks results appear as **Passed** or **Failed** directly in the SCM, preventing PRs from being merged with security issues.&#x20;

## What's next?

* [Configure PR Checks](configure-pr-checks.md)
* [Analyze PR Checks results](pr-checks-results.md)
* [Troubleshooting](troubleshooting.md)
