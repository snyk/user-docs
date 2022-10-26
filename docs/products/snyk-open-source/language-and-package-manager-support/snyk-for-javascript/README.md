# Snyk for JavaScript

You can use Snyk to scan your JavaScript projects.

Snyk supports both npm and Yarn. See these docs pages:

* [Snyk for npm](snyk-for-npm.md)
* [Snyk for Yarn](snyk-for-yarn.md)

The following are the **limitations and workarounds** for Snyk support for JavaScript.

{% hint style="info" %}
Snyk currently does not support **Lerna**. However, if your project is set up using Yarn Workspaces, you can scan the project in the same way you scan any Yarn Workspaces project. If your Lerna project is set up using Yarn Workspaces, you can run `snyk test` and `snyk monitor` using the commands that follow.
{% endhint %}

For each example-package you can use the following:

`snyk monitor --file=packages/example-package/package.json`

Alternatively, you can specify a script to automate scanning nested package.json files:

`ls packages | xargs -I PKG_NAME snyk monitor --file=packages/PKG_NAME/package.json`
