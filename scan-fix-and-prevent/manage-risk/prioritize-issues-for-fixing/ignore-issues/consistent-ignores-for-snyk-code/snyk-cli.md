---
description: How Consistent Ignores for Snyk Code work in the CLI
nav_context: classic
---

# Consistent Ignores for Snyk Code CLI

Ignores are taken into account in the Snyk CLI when `snyk code test` is run.

## Minimum version required

You must have at least Snyk CLI v1.1297.1 installed for Snyk Code Consistent Ignores. See [Install or update the Snyk CLI](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/install-the-snyk-cli).

## Setup

To take ignores into account, specify the Organization where the ignores reside.

[Group-level policies also cascade down to all Organizations](./#manage-ignores-at-the-group-level-through-snyk-code-security-policies). See [How to select the Organization to use in the CLI](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli).

Repository context is required for asset-scoped ignores to take effect. Policy-based ignores such as those based on CWE or Snyk Code Rule ID are still being applied regardless of repository context.

`snyk code test` automatically detects the repository context if a .git directory is present. If not, you can explicitly specify it using the `--remote-repo-url` option. To verify the Git URL, run `git remote -v`.

## Snyk CLI default ignore behavior

The CLI display output hides ignored results by default when you run `snyk code test`. It displays only unignored results and a summary table with the total number of issues (open and ignored).

<figure><img src="../../../../.gitbook/assets/snyk-cli-default-behaviour.png" alt=""><figcaption><p>Snyk CLI default ignore behavior</p></figcaption></figure>

## View ignores in Snyk CLI

After running `snyk code test`, the CLI displays a hint about using the `--include-ignores` parameter to show ignored results.

Running `snyk code test --include-ignores` shows ignored results with their metadata below the open results.

<figure><img src="../../../../.gitbook/assets/display-ignores-snyk-cli.png" alt=""><figcaption><p>Ignores in Snyk CLI</p></figcaption></figure>

## View JSON output

You can find the ignore metadata in the suppressions module of the SARIF output. Run `snyk code test --json` or `snyk code test --sarif` to view this output.

## Access the finding identifier in JSON and SARIF output

The finding identifier is included in the JSON and SARIF output of Snyk CLI. To view it, run `snyk code test --json` and navigate to `runs.results[n].fingerprints.snyk/assets/finding/v1` in the JSON output. See How Snyk Code identifies and tracks issues.

You can use this identifier to [create new ignores using API calls](api.md).

## Create ignores using the Snyk CLI

You can create an ignore for a Snyk Code finding from the command line using the `snyk ignore create` command. Snyk stores the ignore on the finding and applies it consistently across the CLI, IDE, and other integrations on the next test.

Creating ignores from the command line is an Early Access feature of the Ignore Approval Workflow. It applies to `snyk code test` runs from the CLI and IDE. It does not apply to SCM (stateful) tests run through the Import API, and it does not support CLI Upload projects.

Before you create an ignore, complete the [Setup](#setup) and confirm that Snyk Code Consistent Ignores is enabled for your Group or Organization. Commit and push your code to the remote repository so that reviewers can locate the finding. To identify the finding to ignore, obtain its finding identifier as described in [Access the finding identifier in JSON and SARIF output](#access-the-finding-identifier-in-json-and-sarif-output).

To create an ignore interactively and be prompted for each value, run `snyk ignore create` without options.

To create an ignore non-interactively, provide all required options. Use this form in scripts and CI/CD pipelines:

```
$ snyk ignore create \
  --finding-id=<FINDING_ID> \
  --ignore-type=<not-vulnerable|wont-fix|temporary-ignore> \
  --reason="<REASON>" \
  --expiration=<YYYY-MM-DD|never> \
  --org=<ORG_ID> \
  --remote-repo-url=<REPOSITORY_URL>
```

The following table describes the options for `snyk ignore create`.

| Option | Description |
| --- | --- |
| `--finding-id=<FINDING_ID>` | Finding to ignore. Required. |
| `--ignore-type=<TYPE>` | Reason category for the ignore: `not-vulnerable`, `wont-fix`, or `temporary-ignore`. Required. |
| `--reason=<REASON>` | Human-readable justification for the ignore. Required and must not be empty. |
| `--expiration=<VALUE>` | Expiration date in `YYYY-MM-DD` format, or `never`. Required in non-interactive mode. |
| `--org=<ORG_ID>` | Organization that holds the ignore. The value must be a valid Organization ID. |
| `--remote-repo-url=<REPOSITORY_URL>` | Repository URL for the finding. Snyk detects this automatically when a .git directory is present. Specify it explicitly when the repository has a different Git URL. To verify the URL, run `git remote -v`. |

## Ignores in CI/CD pipelines

As ignores are taken into account in Snyk CLI, the same applies when Snyk CLI is integrated into CI/CD pipelines. For example, if a pipeline uses the command `snyk code test –severity-threshold=high` and there are no unignored high-severity results, Snyk CLI will exit with a `0` (success) status code and the build will succeed.

The following example shows how Snyk Code detected high-severity hardcoded secrets, causing a GitHub Action workflow to fail with the exit code `1`.

<figure><img src="../../../../.gitbook/assets/snyk-code-github-actions-exit-code-1.png" alt=""><figcaption><p>High severity hardcoded secreted detected causing GitHub Action workflow to fail with exit code 1</p></figcaption></figure>

In a scenario with ignores applied through Group Policies, Snyk Code has successfully completed the scan, resulting in zero open issues, with the exit code `0`.

<figure><img src="../../../../.gitbook/assets/snyk-code-github-action-exit-code-0.png" alt=""><figcaption><p>High severity issues ignored causing GitHub Action workflow to succeed with exit code 0</p></figcaption></figure>
