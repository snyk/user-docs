# IDE plugin scan fails on Windows systems with .exe download blocking

## **Problem**

When you are using Snyk plugins in an IDE on Windows systems, the Snyk scan may fail.&#x20;

You may see an error message in the logs that indicates `Error running command snyk.start: command 'snyk.workspace.scan' not found. This is likely caused by the extension that contributes snyk.start.`

## **Cause**

This issue is often seen on Windows systems that have a company policy in a firewall or other tool to prevent automatic `.exe` download by showing an interactive popup requesting approval.&#x20;

This popup is not shown in the IDE, and the `.exe` file cannot be downloaded. This may result in a 0KB file being created in the download location.&#x20;

To determine if this is the case, navigate to the location for the CLI or Language Server shown in your Snyk plugin settings in the IDE. If either the CLI or Language Server files are 0KB, this is likely the issue.

## **Solution**

The preferred solution for this is to allowlist `downloads.snyk.io` and `static.snyk.io` in the baseline policy for `.exe` files, so that downloads from this location are exempt from `.exe` popup approvals for all users.&#x20;

You can use the following temporary solutions in the interim if the preferred solution will take time to approve.&#x20;

* Manually download the relevant `.exe` file directly from either [Snyk CLI Releases](https://github.com/snyk/cli/releases) or [Snyk Language Server Releases](https://github.com/snyk/snyk-ls/releases), and place it in the location listed in your plugin settings. Ensure that you turn off the option to **Automatically Manage Dependencies**.

or

* If  `CLI exe` is downloaded, you can use it to provide the language server function. Redirect the language server location to the CLI location in the plugin settings, and ensure that you turn off the option to **Automatically Manage Dependencies**.&#x20;
