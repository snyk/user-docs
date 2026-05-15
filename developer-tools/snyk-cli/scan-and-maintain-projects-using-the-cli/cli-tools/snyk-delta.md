# snyk-delta

This tool provides the means to get the delta between two Snyk Open Source snapshots. This is especially useful when you are running CLI-based scans, such as in your local environment, Git hooks, and so on.

`snyk-delta` compares snapshots to provide details about:

* New vulnerabilities not found in the baseline snapshot
* New license issues not found in the baseline snapshot
* Dependency delta between the two snapshots:
  * Direct dependencies added and removed
  * Indirect dependencies added and removed
  * Flag path(s) carrying new vulnerabilities

## Prerequisites

* Snyk Enterprise plan (requires the Snyk API)
* Your Project to be monitored
* A `SNYK_TOKEN`, either exported to the environment or provided inline to `snyk-delta`, or the API token set with the `snyk config` command. Run `snyk config get api` to see if the API token is set. OAuth authentication is not supported by `snyk-delta`.

## Installation

`npm i -g snyk-delta`

or

Download a binary of your choice from [the release page](https://github.com/snyk-tech-services/snyk-delta/releases)

## Usage

You can use this tool inline or as a standalone command.

### Inline operations

Use `snyk test --json --print-deps | snyk-delta`

* You may point to a specific snapshot by specifying org+project coordinates:\
  `snyk test --json --print-deps | snyk-delta --baselineOrg xxx --baselineProject xxx`
* Use `--setPassIfNoBaseline` if used with `snyk-prevent_commit_status` and the Project is not monitored. This prevents`snyk-prevent_commit_status` from failing:\
  `setPassIfNoBaseline` default to false\
  `snyk test --json --print-deps | snyk-delta --baselineOrg xxx --baselineProject xxx --setPassIfNoBaseline true`

{% hint style="info" %}
The **BaselineProject** value is expected to be a UUID, not a name.\
Check the Snyk Web UI or API to retrieve those UUIDs.
{% endhint %}

### Standalone operations

Use `snyk-delta --baselineOrg xxx --baselineProject xxx --currentOrg xxx --currentProject xxx --setPassIfNoBaseline false`

## Usage as module

```
import { getDelta } from 'snyk-delta'

const jsonResultsFromSnykTest = Read from file or pipe snyk test command

const result = await getDelta(jsonResultsFromSnykTest);
```

**The result** is a number:

* 0: no new issue
* 1: new issue(s) or when using StrictMode and the unmonitored Project has issues (See more details in [StrictMode](snyk-delta.md#strictmode).)
* 2: for errors like invalid auth

Details for issues will be listed on stdout.

## Help for snyk-delta

Use `-h` to display help.

### StrictMode

When `snyk-delta` compares test results, it tries to find the same Project monitored on the Snyk platform. If no monitored Project is found, `snyk-delta` returns all the issues found by the CLI scan, essentially acting as a pass-through.

The return code is 0 if no issue is found, and 1 if issues are found.

### Caution

Use as a module requires a list of issues coming from the Snyk CLI. `snyk-delta` is not compatible with data coming straight from Snyk APIs.

### all-projects

`snyk-delta` does not support the `--all-projects` option, but you can try using `snyk_delta_all_projects.sh` as a workaround until it does.

## Troubleshooting for snyk-delta

If you have trouble, you can try the following:

* Run the Snyk `test -d` step first and ensure it works.
* If you get a 401 error, check to see that you have a valid `SNYK_TOKEN` or run `snyk config get api` .\
  If you do not want to set the token in the environment, you can provide the `SNYK_TOKEN` as an inline command before the `snyk-delta` but after the | symbol: `snyk test --json | SNYK_TOKEN={token} snyk-delta ….{other arguments}`.
* If you are using the `delta allprojects` script, try removing that and testing the Projects individually.
* If no baseline is found, ensure there is an existing monitored Project first, and check the `.git` metadata if you are trying to match against an SCM-monitored Project.

If you need help, contact [Snyk Support](https://support.snyk.io).
