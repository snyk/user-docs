# Snyk CLI for IaC

{% hint style="info" %}
To use the [IaC+](../../../scan-using-snyk/scan-infrastructure/introduction-to-iac+/) version of the Snyk CLI, install Snyk CLI v1.1022.0 or later.
{% endhint %}

## Overview

To use the CLI, you must first [install](../../install-or-update-the-snyk-cli/) it and then [authenticate](../../commands/auth.md).

With Snyk Infrastructure as Code, you can test your configuration files directly from the CLI. See the following pages for details:

* [Test your IaC files](test-your-iac-files/)
* [Share CLI results with the Snyk Web UI](share-cli-results-with-the-snyk-web-ui.md)
* [IaC ignores using the `.snyk` policy file](iac-ignores-using-the-.snyk-policy-file.md)
* [IaC exclusions using the command line](iac-exclusions-using-the-command-line.md)
* [Understanding the IaC CLI test results](understand-the-iac-cli-test-results/) (has information about reports)

## Regularly testing IaC files

Snyk Infrastructure as Code has no equivalent command to `snyk monitor` because the CLI does not send IaC source files back to the platform for periodic testing.

For IaC CLI results to appear in the Snyk Web UI, use [`snyk iac test --report`](https://docs.snyk.io/products/snyk-infrastructure-as-code/share-cli-results-with-the-snyk-web-ui) to capture a one-time snapshot. Optionally, run the command on a recurring schedule to regularly test your IaC files.

Alternatively, you can add an [SCM integration](../../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/), and Snyk will monitor and test a given Git repository on a recurring basis.

## Using Snyk behind a proxy

If you are using a proxy, see [Proxy configuration for Snyk CLI](../../configure-the-snyk-cli/proxy-configuration-for-snyk-cli.md).

For IaC scans specifically, you must also whitelist the \*.snyk.io address, as explained[ ](https://support.snyk.io/hc/en-us/articles/360002153077-How-can-we-whitelist-Snyk-IP-addresses-)on the page [How can we whitelist Snyk IP addresses?](https://support.snyk.io/hc/en-us/articles/360002153077-How-can-we-whitelist-Snyk-IP-addresses-)
