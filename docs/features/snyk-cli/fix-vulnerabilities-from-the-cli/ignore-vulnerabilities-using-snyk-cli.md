# Ignore vulnerabilities using Snyk CLI

{% hint style="info" %}
For [Snyk Open Source, ](https://docs.snyk.io/snyk-open-source)these commands work by default.\
For [Snyk Container](https://docs.snyk.io/snyk-container), these commands also work, but after registering an ignore, when you call **snyk test** or **snyk monitor**, you must use the **--policy-path=** argument: for example: **snyk container test node --policy-path=.snyk**\
For [Snyk Infrastructure as Code](https://support.snyk.io/hc/en-us/categories/360001342678-Infrastructure-as-code), see [IaC ignores using the snyk policy file](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/iac-ignores-using-the-.snyk-policy-file).\
For [Snyk Code](https://docs.snyk.io/snyk-code), these commands are not yet implemented.
{% endhint %}

Sometimes, Snyk may alert you to a vulnerability that has no update or Snyk patch available, or that you do not believe to be currently exploitable in your application. In this case, you may want to tell Snyk to ignore the vulnerability for a certain period of time.

You can ignore a specific vulnerability in a project, using `snyk ignore`:

```
snyk ignore --id=IssueID [--expiry=expiry] [--path='path > to > resource'] [--reason='reason for ignoring']
```

**Options**

`snyk ignore` accepts three options:

| **OPTION** | **DESCRIPTION**                                                                                                                                                                                                                                                                                                                                                                                                | **DEFAULT** | **REQUIRED** |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- | ------------ |
| --id       | <p>The Snyk ID for the issue to ignore. Found by running <code>snyk test</code> and grabbing the last segment of the URL for a given vulnerability.</p><p><strong>Example:</strong> For the vulnerability found at <a href="https://snyk.io/vuln/npm:tough-cookie:20160722">https://snyk.io/vuln/npm:tough-cookie:20160722</a>, you</p><p>would use:</p><p><strong>--id=npm:tough-cookie:20160722</strong></p> | None        | **Yes**      |
| --expiry   | <p>The expiry date string, according to <a href="https://tools.ietf.org/html/rfc2822#page-14">RFC2822</a>.</p><p><strong>Example: --expiry=2017-04-30</strong></p>                                                                                                                                                                                                                                             | 30 days     | No           |
| --reason   | <p>The reason for ignoring the issue.</p><p><strong>Example: --reason='Not currently exploitable.'</strong></p>                                                                                                                                                                                                                                                                                                | None given  | No           |
| --path     | <p>Resource path to ignore the issue for.<br><strong>Example: --path='tough-cookie@2.15.8'</strong></p>                                                                                                                                                                                                                                                                                                        | \*          | No           |

