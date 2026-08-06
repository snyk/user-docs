---
description: How to use the Snyk CLI to scan your source code for hardcoded secrets
nav_context: agnostic
---

# Snyk CLI for Snyk Secrets

The [Snyk Command Line Interface](../../) (CLI) enables you to bring the functionality of [Snyk Secrets](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-secrets) into your development workflow. Using the Snyk CLI, you can scan your source code locally, in a pre-commit hook, or in your CI/CD pipeline to find hardcoded secrets such as API keys, passwords, and tokens.

## Prerequisites for using the Snyk CLI with Snyk Secrets

Before using the Snyk CLI to scan for secrets, verify you have the following prerequisites:

* A Snyk account.
* The Secrets feature enabled for your Snyk Organization. If you receive a `SNYK-CLI-0016` error, contact your Snyk account manager.
* The Snyk CLI installed and authenticated.
  * For instructions, see [Install or update the Snyk CLI](../../install-the-snyk-cli/) and [Authenticate the Snyk CLI](../../authenticate-to-use-the-cli.md).
  * Snyk recommends using the latest version of the CLI.

## Using the Snyk CLI for secrets scans

To scan your code for hardcoded secrets using the Snyk CLI, use the [`snyk secrets test`](../../commands/secrets-test.md) command.\
For more information, see [Secrets scanning in the Snyk CLI](secrets-scanning-in-the-snyk-cli.md).

That page also explains how to do the following:

* [Ignore findings](secrets-scanning-in-the-snyk-cli.md#ignore-findings) that are placeholders, revoked keys, or won't fix scenarios.
* [Review ignored secrets](secrets-scanning-in-the-snyk-cli.md#review-ignored-secrets) to audit your codebase.
* [Scan with a pre-commit hook](secrets-scanning-in-the-snyk-cli.md#scan-with-a-pre-commit-hook) to catch secrets before they reach a commit.
