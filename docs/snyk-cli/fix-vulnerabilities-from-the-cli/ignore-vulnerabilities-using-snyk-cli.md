# Ignore vulnerabilities using Snyk CLI

{% hint style="info" %}
For [Snyk Open Source, ](https://support.snyk.io/hc/en-us/categories/360003049458-Snyk-Open-Source)these commands work by default.  
For [Snyk Container](https://support.snyk.io/hc/en-us/categories/360000583498-Snyk-Container), these commands also work, but after registering an ignore, when you call **snyk test** or **snyk monitor**, you must use the **--policy-file=** argument: for example: **snyk container test node --policy-file=.snyk**  
For [Snyk Infrastructure as Code](https://support.snyk.io/hc/en-us/categories/360001342678-Infrastructure-as-code), see [IaC ignores using the snyk policy file](https://support.snyk.io/hc/en-us/articles/4405381463313-IaC-ignores-using-the-snyk-policy-file).  
For  [Snyk Code](https://support.snyk.io/hc/en-us/categories/360003257537-Snyk-Code), these commands are not yet implemented.
{% endhint %}

Sometimes, Snyk may alert you to a vulnerability that has no update or Snyk patch available, or that you do not believe to be currently exploitable in your application. In this case, you may want to tell Snyk to ignore the vulnerability for a certain period of time.

You can ignore a specific vulnerability in a project, using `snyk ignore`:

```text
snyk ignore --id=IssueID [--expiry=expiry] [--reason='reason for ignoring']
```

**Options**

`snyk ignore` accepts three options:

<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>OPTION</b>
      </th>
      <th style="text-align:left"><b>DESCRIPTION</b>
      </th>
      <th style="text-align:left"><b>DEFAULT</b>
      </th>
      <th style="text-align:left"><b>REQUIRED</b>
      </th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">--id</td>
      <td style="text-align:left">
        <p>The Snyk ID for the issue to ignore. Found by running <code>snyk test</code> and
          grabbing the last segment of the URL for a given vulnerability.</p>
        <p><b>Example:</b> For the vulnerability found at <a href="https://snyk.io/vuln/npm:tough-cookie:20160722">https://snyk.io/vuln/npm:tough-cookie:20160722</a>,
          you</p>
        <p>would use:</p>
        <p><b>--id=npm:tough-cookie:20160722</b>
        </p>
      </td>
      <td style="text-align:left">None</td>
      <td style="text-align:left"><b>Yes</b>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">--expiry</td>
      <td style="text-align:left">
        <p>The expiry date string, according to <a href="https://tools.ietf.org/html/rfc2822#page-14">RFC2822</a>.</p>
        <p><b>Example: --expiry=2017-04-30</b>
        </p>
      </td>
      <td style="text-align:left">30 days</td>
      <td style="text-align:left">No</td>
    </tr>
    <tr>
      <td style="text-align:left">--reason</td>
      <td style="text-align:left">
        <p>The reason for ignoring the issue.</p>
        <p><b>Example: --reason=&apos;Not currently exploitable.&apos;</b>
        </p>
      </td>
      <td style="text-align:left">None given</td>
      <td style="text-align:left">No</td>
    </tr>
  </tbody>
</table>

