# Snyk for JavaScript

You can use Snyk to scan your JavaScript Projects.

Snyk supports both npm and Yarn. See [Snyk for npm](snyk-for-npm.md) and [Snyk for Yarn](snyk-for-yarn.md).

The following are the **limitations and workarounds** for Snyk support for JavaScript.

{% hint style="info" %}
Snyk currently does not support **Lerna**. However, if your Project is set up using Yarn Workspaces, you can scan the Project in the same way you scan any Yarn Workspaces Project.
{% endhint %}

If your Lerna Project is set up using Yarn Workspaces, you can run `snyk test` and `snyk monitor` as follows:

For each example-package you can use\
`snyk monitor --file=packages/example-package/package.json`

Alternatively, you can specify a script to automate scanning nested `package.json` files:\
`ls packages | xargs -I PKG_NAME snyk monitor --file=packages/PKG_NAME/package.json`
