# snyk-delta

This tool provides the means to get the delta between two Snyk snapshots. This is especially useful when you are running CLI-based scans, such as in your local environment, git hooks, and so on.

`snyk-delta` compares snapshots to give details about:

* New vulnerabilities not found in the baseline snapshot
* New license issues not found in the baseline snapshot
* Dependency delta between the two snapshots:
  * Direct dependencies added and removed
  * Indirect dependencies added and removed
  * Flag path(s) carrying new vulnerabilities

## Prerequisites

* Snyk Enterprise plan (requires API)
* Your project to be monitored

## Installation

`npm i -g snyk-delta`

or

Grab a binary of your choice from [the release page](https://github.com/snyk-tech-services/snyk-delta/releases)

## Usage

You can use this tool inline, or as a standalone command.

### Inline operations

Use `snyk test --json --print-deps | snyk-delta`

* Possibly point to a specific snapshot by specifying org+project coordinates:\
  `snyk test --json --print-deps | snyk-delta --baselineOrg xxx --baselineProject xxx`
* Use `--setPassIfNoBaseline` if used with `snyk-prevent_commit_status` and the project is not monitored. This prevents`snyk-prevent_commit_status` from failing:\
  `setPassIfNoBaseline` default to false\
  `snyk test --json --print-deps | snyk-delta --baselineOrg xxx --baselineProject xxx --setPassIfNoBaseline true`

{% hint style="info" %}
The **BaselineProject** value is expected to be a UUID, not simply a name\
Check your Snyk Web UI or API to retrieve those UUIDs.
{% endhint %}

### Standalone operations

Use `snyk-delta --baselineOrg xxx --baselineProject xxx --currentOrg xxx --currentProject xxx --setPassIfNoBaseline false`

## Usage as module

```
import { getDelta } from 'snyk-delta'

const jsonResultsFromSnykTest = Read from file or pipe snyk test command

const result = await getDelta(jsonResultsFromSnykTest);
```

**Result** is a number:

* 0: no new issue
* 1: new issue(s) or when using strictMode and the unmonitored project has issues (see more details in [StrictMode](snyk-delta.md#strictmode))
* 2: for errors like invalid auth

Actual issue(s) details will be listed on stdout.

## Help

`-h` to list help

### StrictMode

When `snyk-delta` compares test results, it tries to find the same project, monitored on the Snyk platform. If no monitored project is found, `snyk-delta` returns all the issues found by the CLI scan, essentially acting as pass through.

The return code is 0 if no issue, 1 if issues.

### Caution

Usage as a module requires list of issues coming from Snyk CLI. Currently `snyk-delta` is not compatible with data coming straight from Snyk APIs.

### all-projects

`snyk-delta` does not currently support the `--all-projects` option, but you can try using `snyk_delta_all_projects.sh` as a workaround until it does.
