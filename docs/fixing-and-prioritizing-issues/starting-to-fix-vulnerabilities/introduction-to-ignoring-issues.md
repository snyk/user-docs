# Introduction to ignoring issues

{% hint style="info" %}
See [Ignoring issues not prioritized for your project](fixing-and-prioritizing-issues/issue-management/ignore-issues) for more details of ignoring issues.
{% endhint %}

## Snyk UI

Each issue card has an **Ignore** button allowing you to ignore that issue:

![](../../.gitbook/assets/new-ignore-2.png)

## Snyk CLI

In the Snyk CLI, you can ignore issues using **snyk ignore**. For example:

`snyk ignore --id='npm:braces:20180219' --expiry='2018-04-01' --reason='testing'`

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

