# Using Snyk Code via CLI

{% hint style="info" %}
Make sure Snyk Code is enabled for your organization, see [Getting started with Snyk Code](../../../getting-started/getting-started-snyk-products/getting-started-with-snyk-code.md#stage-1-enable-snyk-code) for details.
{% endhint %}

The Snyk Command Line Interface ([CLI](../../../features/snyk-cli/)) for Snyk Code helps you find and fix security flaws in your code on your local machine.

To use the CLI you must first [install](../../../features/snyk-cli/install-the-snyk-cli/) it and then [authenticate](../../../features/snyk-cli/commands/auth.md).

## **Test a project or folder**

To test the current folder, run `snyk code test` with no parameters.

To test another context, run `snyk code test <my-folder-path>` with a path to a file or folder as the parameter.

All sub-folders inside the specified folder are also scanned.

## "Not supported" message

{% hint style="danger" %}
**Snyk Code is not supported for org** _**your.org**_

A minimum version of `1.716.0` is required to use CLI for Snyk Code.
{% endhint %}
