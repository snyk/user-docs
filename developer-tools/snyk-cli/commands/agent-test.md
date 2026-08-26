---
description: The snyk agent test command that runs Open Source, Code, and Secrets scans together
---

# Agent test

{% hint style="warning" %}
`snyk agent test` is experimental. Commands, flags, and output format may change in future releases.
{% endhint %}

## Prerequisites

**Note:** Requires Snyk CLI v1.1307.0 or later.

## Usage

`snyk agent test [<OPTIONS>]`

## Description

The `snyk agent test` command runs Snyk Open Source, Snyk Code, and Snyk Secrets scans together in a single pass, returning token-optimized output intended for AI coding agents rather than direct human review

## Exit codes

Possible exit codes and their meaning:

**0**: success (scan completed), no issues found\
**1**: action_needed (scan completed), issues found\
**2**: failure, try to re-run the command. Use `-d` to output the debug logs.

## Configure the Snyk CLI

You can use environment variables to configure the Snyk CLI and set variables for connecting with the Snyk API. For more information see [Configure the Snyk CLI](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)

## Code execution warning

Before scanning your code, review the [Code execution warning for Snyk CLI](https://docs.snyk.io/snyk-cli/code-execution-warning-for-snyk-cli)

## Debug

Use the `-d` option to output the debug logs.

## Options

{% hint style="warning" %}
Pending confirmation from engineering. Add each confirmed option below as its own H3 heading, following the pattern used in other command pages (for example `--org=<ORG_ID>` in [code test](code-test.md#org-org_id)).
{% endhint %}

## Examples for the snyk agent test command

### Run a combined scan

```bash
$ snyk agent test
```
