# Ignore vulnerabilities using Snyk CLI

{% hint style="info" %}
For [Snyk Open Source, ](https://docs.snyk.io/snyk-open-source)these options work by default.\
For [Snyk Container](https://docs.snyk.io/snyk-container), these options also work, but after registering an ignore, when you call `snyk test` or `snyk monitor`, you must use the `--policy-path=` option, for example: `snyk container test node --policy-path=.snyk.`\
For [Snyk Infrastructure as Code](https://support.snyk.io/hc/en-us/categories/360001342678-Infrastructure-as-code), see [IaC ignores using the snyk policy file](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/iac-ignores-using-the-.snyk-policy-file).\
For [Snyk Code](https://docs.snyk.io/snyk-code), these options are not yet implemented.
{% endhint %}

Sometimes, Snyk alerts you to a vulnerability that has no update or Snyk patch available, or that you do not believe to be currently exploitable in your application. When this happens you may want to tell Snyk to ignore the vulnerability for a certain period of time.

You can ignore a specific vulnerability in a project using the `snyk ignore` command.

`snyk ignore --id=<ISSUE_ID> [--expiry=<EXPIRY>] [--reason=<REASON>] [--policy-path=<PATH_TO_POLICY_FILE>] [<OPTIONS>]`

The `snyk ignore` command updates the `.snyk` file and supports the following options:

| **OPTION**                            | **DESCRIPTION**                                                                                                                                                                                                                                                                                                                                                                                       | **DEFAULT** | **REQUIRED** |
| ------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- | ------------ |
| `--id`                                | <p>The Snyk ID for the issue to ignore. Found by running <code>snyk test</code> and grabbing the last segment of the URL for a given vulnerability.</p><p><strong>Example:</strong> For the vulnerability found at <a href="https://snyk.io/vuln/npm:tough-cookie:20160722">https://snyk.io/vuln/npm:tough-cookie:20160722</a>, the Snyk ID is:</p><p><code>--id=npm:tough-cookie:20160722</code></p> | None        | Yes          |
| `--expiry`                            | <p>Expiry date in YYYY-MM-DD format (<a href="https://tools.ietf.org/html/rfc2822#page-14">RFC2822</a> and <a href="https://www.iso.org/iso-8601-date-and-time-format.html">ISO 8601</a> are supported).</p><p>Example: <code>--expiry=2017-04-30</code>.</p>                                                                                                                                         | 30 days     | No           |
| `--reason`                            | Human-readable \<REASON> to ignore this issue. Example: `reason='Not currently exploitable'`.                                                                                                                                                                                                                                                                                                         | None        | No           |
| `--policy-path=<PATH_TO_POLICY_FILE>` | Path to a .snyk policy file to pass manually.                                                                                                                                                                                                                                                                                                                                                         | None        | No           |
| `--path`                              | Path to resource for which to ignore the issue. Example: `path='tough-cookie@2.15.8'`                                                                                                                                                                                                                                                                                                                 | All         | No           |
