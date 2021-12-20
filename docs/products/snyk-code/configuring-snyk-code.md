# Configuring Snyk Code

To use Snyk Code, you have to make sure it's enabled for your organization and then import/re-import your repositories via the UI or API.

{% hint style="info" %}
See [Getting started with Snyk Code](https://docs.snyk.io/getting-started/getting-started-snyk-products/getting-started-with-snyk-code) to enable Code and start scanning your projects.
{% endhint %}

Once you have Code enabled, the following will happen:

1. Every time a new repository is imported (or re-imported), if it contains any source code files, a "Code Analysis" project is created for it
2. You may now run Snyk Code tests against this organization via the CLI
3. You may now run Snyk Code tests against this organization via one of the IDE plugins

![](../../.gitbook/assets/screenshot\_2021-06-17\_at\_13.23.19.png)
