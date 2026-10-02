---
description: >-
  The snyk fix --agentic command, which applies and verifies vulnerability
  fixes
---

# Fix

`snyk fix --agentic` is a new LLM-driven command. It is distinct from the legacy `snyk fix` command, which used a deterministic fix engine and is deprecated.

## Prerequisites

To use the `snyk fix --agentic` command:

* Install the latest version of the [Snyk CLI](../install-the-snyk-cli/README.md)
* [Authenticate](auth.md) your machine with the Snyk CLI using `snyk auth`
* Configure access to an LLM provider. Not every provider uses an API key. Visit [Remediation Agent](../../../scan-fix-and-prevent/fix/remediation-agent.md) for the supported providers and their setup instructions.

## Usage

`snyk fix --agentic --experimental --sca|--sast|--container [<OPTIONS>] [<PATH>]`

## Description

The `snyk fix --agentic` command is an LLM-driven vulnerability remediation command. It scans your Project for vulnerabilities, generates a fix plan enriched with Snyk security intelligence, and walks you through applying and verifying each fix interactively. Use `--auto-approve` to apply fixes without prompting.

Choose what to remediate with a product flag:

* `--sca`: Snyk Open Source. Fixes vulnerable dependencies by bumping versions and adding overrides.
* `--sast`: Snyk Code. Applies Snyk Agent Fix suggestions for source code vulnerabilities and rescans to confirm the vulnerability is resolved.
* `--container`: Snyk Container. Bumps Dockerfile base images and patches OS packages, then builds and rescans the image to verify the result.

Pass exactly one of `--sca`, `--sast`, or `--container`. The `--experimental` flag is required alongside `--agentic`.

Visit [Remediation Agent](../../../scan-fix-and-prevent/fix/remediation-agent.md) for conceptual documentation, including setup instructions and supported IDEs.

The Remediation Agent is under rapid development. Some options on this page are available only in the latest preview release of the Snyk CLI. Run `snyk version` to check which version you have.

## Exit codes

Possible exit codes and their meaning:

**0**: Success. Vulnerabilities were remediated.\
**1**: Action needed. Some vulnerabilities could not be fixed.\
**2**: Failure. An error occurred. Use `-d` to output the debug logs.

## Configure the Snyk CLI

You can use environment variables to configure the Snyk CLI and set variables for connecting with the Snyk API. For more information see [Configure the Snyk CLI](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)

## Code execution warning

Before scanning your code, review the [Code execution warning for Snyk CLI](https://docs.snyk.io/snyk-cli/code-execution-warning-for-snyk-cli)

## Debug

Use the `-d` option to output the debug logs.

## Options

### `--agentic`

Enable the LLM-driven fix flow. Required to use this command in agentic mode.

### `--experimental`

Required alongside `--agentic`. Acknowledges the release status of the command.

### `--sca`

Remediate Snyk Open Source (dependency) vulnerabilities. Mutually exclusive with `--sast` and `--container`.

### `--sast`

Remediate Snyk Code (source code) vulnerabilities. Mutually exclusive with `--sca` and `--container`.

### `--container`

Remediate Snyk Container vulnerabilities by editing a single Dockerfile: bump the base image and patch OS packages. Mutually exclusive with `--sca` and `--sast`, and requires `--agentic`.

**Note:** Because the container flow selects fixes at the Dockerfile and package level rather than per vulnerability, it does not accept `--issue-ids`, `--exclude-ids`, `--severity-threshold`, or `--severity-filter`. Passing any of them ends the run before the scan starts.

### `--provider=<PROVIDER>`

LLM provider to use.\
Allowed values: `anthropic`, `openai`, `bedrock`, `vertex`, `litellm`, `ollama`\
Default: `anthropic`

Example: `snyk fix --agentic --experimental --sca --provider=openai`

### `--model=<MODEL>`

Model ID, which overrides the provider's default model. Required for Ollama, for example `llama3.1`. For Vertex AI, use a Gemini or Claude model name, for example `gemini-2.5-flash`. Visit [Configure a model provider](../../../scan-fix-and-prevent/fix/remediation-agent.md#configure-a-model-provider) for Amazon Bedrock model ID guidance.

Example: `snyk fix --agentic --experimental --sca --provider=ollama --model=llama3.1`

### `--dry-run`

Show the fix plan without applying changes.

Example: `snyk fix --agentic --experimental --sca --dry-run`

### `--auto-approve`

Run without prompting.\
**Important:** In addition to approving fixes, this option trusts the folder, continues after test failures, skips advisories, and approves partial plans when the agent reaches a budget warning.

### `--issue-ids=<ID>[,<ID>]...`

Comma-separated list of IDs to fix, accepting both SCA vulnerability IDs and SAST finding IDs.\
**Note:** Requires `--auto-approve`, and does not combine with `--exclude-ids`. The command exits with a nonzero code if any requested ID is not fixed.

Example: `snyk fix --agentic --experimental --sca --issue-ids=SNYK-JS-FOO-123,SNYK-JS-BAR-456`

### `--exclude-ids=<ID>[,<ID>]...`

Comma-separated list of IDs to prune from the run, so the agent fixes everything except these. Accepts both SCA vulnerability IDs and SAST finding IDs.\
**Note:** Requires `--auto-approve`, and does not combine with `--issue-ids`. An SCA fix group is skipped only when all of its IDs are excluded.

### `--severity-threshold=<SEVERITY>`

Fix only vulnerabilities at or above this severity.\
Allowed values: `low`, `medium`, `high`, `critical`\
**Note:** Requires `--auto-approve`.

Example: `snyk fix --agentic --experimental --sca --severity-threshold=high`

### `--severity-filter=<SEVERITY>[,<SEVERITY>]...`

Fix only vulnerabilities whose severity is exactly one of these.\
Allowed values: `low`, `medium`, `high`, `critical`\
**Note:** Unlike `--severity-threshold`, this is an exact match, and it applies to interactive runs as well as auto-approved ones.

Example: `snyk fix --agentic --experimental --sca --severity-filter=high,critical`

### `--breakability-filter=<RATING>[,<RATING>]...`

Fix only SCA fixes whose assessed breakability is exactly one of these.\
Allowed values: `low`, `medium`, `high`\
**Note:** Applies to Snyk Open Source only, so it skips the SAST leg.

Example: `snyk fix --agentic --experimental --sca --breakability-filter=low`

### `--no-breakability`

Skip the Snyk Breakability API and use the local heuristic only.

### `--enable-revert`

Roll a fix's edits back when it fails, instead of keeping them.\
Default: off, so failed fixes are kept and reported\
**Note:** A test failure always keeps the fix.

### `--fix-report=<PATH>`

Write a JSON report of fixed and not-fixed vulnerabilities, including ID, title, and severity, to this file path.

### `--ci`

Produce plain output suited to a CI log, with no colors or spinner.\
**Note:** This option also disables parent-package fix discovery.

### `--agent-max-iterations=<N>`

Cap how many iterations the agent runs per fix.\
Default: `0`, which uses the built-in limit of 50

### `--max-validation-attempts=<N>`

Cap the number of validation retry attempts per fix. After this many failed test runs, the agent skips the fix and moves on.\
Default: `0`, which uses the built-in default

### `--additional-params="<PARAMS>"`

Extra arguments forwarded to the Snyk Open Source scan.

Example: `snyk fix --agentic --experimental --sca --additional-params="--exclude=vendor --detection-depth=3"`

### `-d`, `--debug`

Print debug information to stderr.
