---
description: Understanding the integration strategy
---

# Overview of IDE plugins

This document shows how to extract Snyk scan results and make them available in your IDE plugin by using the Snyk CLI. The IDE plugin downloads the CLI in the background and periodically checks for newer versions and downloads those when available.

Snyk SDK makes it easier to implement and download periodic updates. While this integration is under development, a [Java implementation](https://github.com/jenkinsci/snyk-security-scanner-plugin/blob/master/src/main/java/io/snyk/jenkins/tools/internal/DownloadService.java) is available.

{% hint style="info" %}
Download the Snyk CLI for IDE usage whether or not the developer's machine has the Snyk CLI installed.
{% endhint %}

To use the Snyk CLI you must authenticate. To check whether the user is previously authenticated on this machine, run `snyk config get api` and check for a return of a UUID. If needed run `snyk auth` and follow the prompts to sign up and log in. The process stores the Snyk token locally on the user's machine. For more information see [Authenticating the user to Snyk](overview.md#6689c939-0bff-4d30-9480-b62179889e37).

Use the Snyk CLI `test` command with its `--json` option to convert the output into machine-readable format. For every supported dependency manifest file in a project you need to run `snyk test --file=<manifest file name> --json` , parse the results, and display them in the intended places inside the IDE user interface.

Rerun `snyk test` periodically or when a triggering event happens: a manifest file has been edited, a day has passed since the last scan, and any time you explicitly reinstall the dependencies (`mvn install`).

In summary the steps are as follows:

1. Check for the presence of the Snyk CLI or prompt the user to install it (one time).
2. Check if the user is authenticated to the CLI and run `snyk auth` if not (one time).
3. Scan every manifest file using `snyk test --file=<manifest file name> --json` parsing the output and incorporating it into the IDE user interface.
4. Rerun a scan periodically or when a triggering event happens; see [Scanning with IDE plugins](scanning.md#607b2cd8-2fb5-49ee-8473-319a42b8c421).
