---
description: Understanding the integration strategy
---

# Overview

For the purpose of extracting Snyk scan results and making them available in your IDE plugin, we will make use of the Snyk CLI which the IDE plugin should download in the background and periodically check for newer versions and download those when available.

Snyk is working on a SDK to make it easier to implement the download and period updates. In the meantime you can get a taste of a Java [implementation here](https://github.com/jenkinsci/snyk-security-scanner-plugin/blob/master/src/main/java/io/snyk/jenkins/tools/internal/DownloadService.java):

{% hint style="info" %}
You should download the Snyk CLI for the IDE usage irrespective of whether the developer has snyk CLI installed on their machine as well.
{% endhint %}

Next, we need to make sure the user is authenticated into their snyk account, and if not run `snyk auth` which will take the user through a signup/login experience and will store their Snyk token locally on the machine.

For the purpose of scanning we leverage the Snyk CLI's `test` command with its `--json` flag to convert the output into machine readable format. In short, for every supported dependency manifest file in a project you need to run `snyk test --file=<manifest file name> --json` , parse the results and display them in the intended places inside the IDE user interface.

Rerun the snyk scan periodically or when a triggering event happens \(see [When to rerun](https://www.notion.so/snyk/How-to-Build-an-IDE-plugin-that-incorporates-Snyk-scanning-6aa2c0a9291e405bb8b26431039fc21c#607b2cd82fb549ee8473319a42b8c421) a snyk scan below\).

We can summarize the steps are as follows:

1. Check for the existence of the Snyk CLI, or prompt user to install it \(one time\).
2. Check if the user is authenticated to the CLI and run `snyk auth` if not \(one time\).
3. Scan every manifest file using `snyk test --file=<manifest file name> --json` parsing the output and incorporated it into the IDE user interface.
4. Rerun a scan periodically or when a triggering event happens.

## Authenticating the user to Snyk <a id="6689c939-0bff-4d30-9480-b62179889e37"></a>

In order to run the security scans using the Snyk CLI the user will need to authenticate to Snyk first. To check whether the user is previously authenticated on this machine you can run `snyk config get api` and check for a return of a UUID.

If the user is not authenticated run the `snyk auth` command which will take the user through a signup/login experience and will store their Snyk token locally on the machine.

##  <a id="76377ada-806b-4416-b68e-429e00cf0db8"></a>

