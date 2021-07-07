# Ignoring issues not prioritized for your project

* [ Evaluating and prioritizing vulnerabilities](/hc/en-us/articles/360006113978-Evaluating-and-prioritizing-vulnerabilities)
* [ Remediate your vulnerabilities](/hc/en-us/articles/360006113798-Remediate-your-vulnerabilities)
* [ Upgrading package versions to fix](/hc/en-us/articles/360005993658-Upgrading-package-versions-to-fix)
* [ Snyk patches to fix](/hc/en-us/articles/360004032437-Snyk-patches-to-fix)
* [ Ignoring issues not prioritized for your project](/hc/en-us/articles/360004002718-Ignoring-issues-not-prioritized-for-your-project)

##  Ignoring issues not prioritized for your project

### Ignoring issues in the UI

Each issue card has an **Ignore** button that opens up a dialog where you can select why you want to ignore the issue, and how long to ignore it.

If you select **Ignore temporarily,** then you can check the **Until fix is available** checkbox:

This will resurface the vulnerability as soon as we have a fix for it, and you can optionally give additional details on why you’re ignoring the issue. This is checked by default if there is currently no remediation available for this issue.

An issue is ignored until ANY of the conditions happen - either the ignore period expires, OR the vuln becomes fixable.

When you ignore an issue in our UI, it will show who ignored it and allow you to edit or unignore it.

### Ignoring issues in the CLI

Suppressing issues is possible via the CLI. For node.js projects, you can use Snyk wizard, which will give you the option of ignoring the vulnerability for a period of 30 days. For other supported languages,  or to specify a different duration, use **snyk ignore**.

`snyk ignore --id='npm:braces:20180219' --expiry='2018-04-01' --reason='testing'`

When using Snyk wizard or Snyk ignore the .snyk policy file is updated with the path and given reason \(if one was provided\). Here’s an example:

```text
'npm:moment:20170905':
- moment:
reason: The reason given
expires: '2017-12-29T16:10:16.946Z'
```

See [Ignore vulnerabilities](https://support.snyk.io/hc/en-us/articles/360003851317-Ignore-vulnerabilities) for more details of ignoring issues in the CLI.

### Set who can configure ignore settings

Since suppressing vulnerabilities carries a level of risk, there is a setting that allows you to decide whether this feature is available to admins only. To set this, go to your organization settings, and select Admin only in the “Ignores” section. When you enable Admin only for ignores, this will also disable ignores from being added via the CLI because we are unable to prevent non-admins from ignoring issues in this environment.

You can also choose to set the more details field to be a compulsory field when an issue is being ignored, requiring the user to enter a reason for each ignore.

### Ignoring issues when necessary

Ignoring security issues shouldn’t be the default action, but it is sometimes necessary. The best practice is to fix or patch vulnerabilities, or to remove the vulnerable dependency, but there may still be reasons why you would want to suppress an issue – for example, if an issue doesn’t currently have a fix, you might want to snooze it until it does.

Some issues are irrelevant for certain projects \(e.g. a DDOS attack for an internal service\). Other times, an issue has a path that makes it non-exploitable. In these scenarios, Snyk would still recommend fixing the issue where possible, as just because the vulnerability is not exploitable today via the current code paths, it does not mean it will not be in the future if the code changes.

Issues can be ignored and viewed via the snyk.io UI using the issue filter on your project, via the Snyk APIs and via the CLI.

If you have access to our Reports feature, you will also be able to see an overview of how many issues in your organization’s projects are ignored, along with an option to filter these so you can drill down into each one. If the issue was ignored in our UI, we include a credit for additional accountability, so you can see who initiated it.

See [Reports](https://support.snyk.io/hc/en-us/categories/360000598418-Reports) for more details.

### Ignoring issues in Snyk Code

For [Snyk Code](https://support.snyk.io/hc/en-us/categories/360003257537-Snyk-Code), ignore functionality may capture a wider range of issues than other products.

Snyk Code's static code analysis transforms the input code into an **intermediate representation**, which captures the flow of code, but abstracts away some details. Snyk Code uses this representation to recognize the same issue even when you refactor your code or rename a variable.

So when you ignore an issue, Snyk Code can also ignore that issue if it occurs in multiple places in your code, even with minor code changes. This avoids seeing multiple duplicate reports for pieces of code with the same ignored issue.

As an example, the following two code snippets \(despite textual differences\) denounce the same issue, as we only renamed the variables:

```text
var fs = require('fs');
var logFileName = req.query.file || 'standard_log.log';
var logfile = fs.readFile(logFileName, "utf8", function(err, data) {...
```

```text
var filesystem = require('fs');
var generalLogFileName = req.query.file || 'standard_log.log'; 
var handleLogFile = filesystem.readFile(generalLogFileName, "utf8", function(err, data) {...
```

See [Secure coding with Snyk Code: Ignore functionality with a twist](https://snyk.io/blog/secure-coding-snyk-code-ignore-feature-twist/) for more details.

