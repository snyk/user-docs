# Ignore issues

You can set ignores individually or as actions - see [Setting ignores](ignore-issues.md#setting-ignores) for details.

If you do not want to fix a vulnerability or license issue, and don't want to see that issue in scan results, Snyk allows you to ignore it, either temporarily or permanently.

## Deciding to ignore issues

**Note:** Ignoring security issues should not be the default action, but it is sometimes necessary. The best practice is to fix or patch vulnerabilities, or to remove the vulnerable dependency.

However, there may be reasons why you would want to suppress an issue. For example, if an issue does not currently have a fix, you might want to ignore it until there is a fix.

Some issues are irrelevant for certain projects (for example, a distributed denial-of-service (DDoS) attack for an internal service). Other times, an issue has a path that makes it non-exploitable. In these scenarios, you should still fix the issue. A vulnerability that is not exploitable today may be exploitable in future.

## Setting ignores

Issues can be ignored and viewed through the following means:

* [Snyk Web UI](ignore-issues.md#ignoring-issues-in-the-web-ui)
* [Snyk CLI](ignore-issues.md#ignoring-issues-in-the-cli)
* [Scanning from the CLI or CI/CD, ignoring in the Web UI](ignore-issues.md#scanning-from-the-cli-or-ci-cd-ignoring-in-the-ui)
* [The .snyk file](ignore-issues.md#ignoring-issues-with-the-.snyk-file)
* [Policy actions](ignore-issues.md#ignoring-issues-with-policy-actions)

### Ignoring issues in the Web UI

Each issue card has an **Ignore** button that opens a dialog where you can select why you want to ignore the issue and how long to ignore it.

![Ignore dialog from issue card](../../../.gitbook/assets/ignore-vulnerability-ui-updated.png)

If you select **Ignore temporarily,** then you can check the **Until fix is available** checkbox:

![Ignore temporarily](<../../../.gitbook/assets/image (19) (1) (1).png>)

This resurfaces the vulnerability as soon as Snyk has a fix for it, and optionally you can give additional details on why you are ignoring the issue. This is checked by default if there is currently no fix available for this issue.

{% hint style="info" %}
An issue is ignored until ANY of the conditions happen, the ignore period expires, OR the vulnerability becomes fixable.
{% endhint %}

When you ignore an issue in the Snyk Web UI, the issue shows who ignored it and allow you to edit or unignore it.

![Ignore set in the Snyk Web UI](<../../../.gitbook/assets/image (14) (1) (1).png>)

For more information see the training: [Ignoring issues](https://training.snyk.io/courses/ignore-strategies).

### Ignoring issues in the CLI

You can suppress issues through the CLI by using the `snyk ignore` command, for example:

`snyk ignore --id='npm:braces:20180219' --expiry='2018-04-01' --reason='testing'`

For more information see the [`ignore`](../../../snyk-cli/commands/ignore.md) command help and [Ignore vulnerabilities using Snyk CLI](../../../snyk-cli/test-for-vulnerabilities/ignore-vulnerabilities-using-snyk-cli.md).

When you use `snyk ignore`**,** the `.snyk` policy file is updated with the path and reason given, if one was provided. For example:

```
'npm:moment:20170905':
- moment:
reason: The reason given
expires: '2017-12-29T16:10:16.946Z'
```

### Scanning from the CLI or CI/CD, ignoring in the Web UI

Ignores between a CLI (or CI/CD run) and the Snyk UI are synchronized as follows:

1. A project is scanned and pushed to the UI using `snyk monitor`.
2. You see the results of the scan and choose to ignore an issue.
3. The issue is ignored when running `snyk test` or `snyk monitor` in the CI/CD or CLI.

Refer to the following example:

![Project imported from snyk monitor, ignore set in the Web UI](../../../.gitbook/assets/ignore-vulnerability-snyk-monitor-updated.png)

The following shows `snyk test` results before ignoring in the Web UI:

![Snyk test results before ignoring in the Web UI](<../../../.gitbook/assets/image (18) (1) (1) (1).png>)

The following shows `snyk test` results after ignoring in the Web UI:

![Snyk test results after ignoring in the Web UI](<../../../.gitbook/assets/image (20) (1) (1) (1).png>)

**Note**: The preceding example shows what happens if you ignore the project imported by `snyk monitor` from the CLI or CI/CD.

The same repository imported from the SCM is considered as a different project, and any ignore on an SCM project does not impact the results of a `snyk test` from the CLI or a CI/CD. SCM and CI projects behave as two stand alone projects.

### Ignoring issues with the .snyk file

For all projects, you can ignore the vulnerability by creating a `.snyk` YAML file.

![A .snyk file](../../../.gitbook/assets/screen+shot+2017-05-10+at+11.16.57+am.png)

For example, if you wanted to ignore the vulnerability with SNYK ID [SNYK-RUBY-FASTREADER-20085](https://snyk.io/vuln/SNYK-RUBY-FASTREADER-20085) in `fastreader`, with the reason “No fix available” until 01 Jan 2023, you would write the following:

```
version v1.5.0
ignore:
    'SNYK-RUBY-FASTREADER-20085':
     - '* > fastreader':
      reason: 'No fix available'
      expires '2023-01-01T00:00:00.000Z'
```

For more information, see [The .snyk file](https://docs.snyk.io/fixing-and-prioritizing-issues/policies/the-.snyk-file).

### Ignoring issues with policy actions

You can set [Security policies](https://docs.snyk.io/fixing-and-prioritizing-issues/security-policies) actions to ignore all vulnerabilities that match the conditions specified in a policy rule.

For more information, see [Security policies: Actions](https://docs.snyk.io/fixing-and-prioritizing-issues/security-policies/security-policies-actions).

## Snyk Code: ignoring issues

For [Snyk Code](https://docs.snyk.io/snyk-code), ignore functionality may capture a wider range of issues than other products.

Snyk Code's static code analysis transforms the input code into an **intermediate representation**, which captures the flow of code, but abstracts away some details. Snyk Code uses this representation to recognize the same issue even when you refactor your code or rename a variable.

Thus when you ignore an issue, Snyk Code can also ignore that issue if it occurs in multiple places in your code, even with minor code changes. This avoids generating multiple duplicate reports for pieces of code with the same ignored issue.

As an example, the following two code snippets (despite textual differences) denounce the same issue, as we only renamed the variables:

```
var fs = require('fs');
var logFileName = req.query.file || 'standard_log.log';
var logfile = fs.readFile(logFileName, "utf8", function(err, data) {...
```

```
var filesystem = require('fs');
var generalLogFileName = req.query.file || 'standard_log.log'; 
var handleLogFile = filesystem.readFile(generalLogFileName, "utf8", function(err, data) {...
```

For more information, see [Exploring the vulnerability issues discovered by Snyk Code](https://docs.snyk.io/products/snyk-code/exploring-and-working-with-the-snyk-code-results/exploring-the-vulnerability-issues-discovered-by-snyk-code).

## Snyk IaC: ignoring issues

When scanning your IaC configuration files using the Snyk CLI with `snyk iac test` you can ignore issues that are not relevant to you by using [The .snyk file](../../../snyk-cli/test-for-vulnerabilities/the-.snyk-file.md).

Snyk recommends storing and versioning the .snyk file in root of your working directory for where you store your IaC configuration files.

For more information, see [IaC ignores using the .snyk policy file](../../../products/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/iac-ignores-using-the-.snyk-policy-file.md).

## Configure ignore settings

Suppressing vulnerabilities carries a level of risk, so you can make this function available to admins only as follows

1. Go to your organization settings > **General**, then navigate in the **Ignores** section
2. Under **Ability to ignore an issue, or edit the ignore settings on an issue**, select **Admin users only**. Note that this also disables ignores from being added via the CLI.
3. Under **Require reason for each ignore**, you can also choose to set the **more details** field to be a required field when an issue is being ignored, ensuring the user enters a reason for each ignore.
4. Click **Update** to make the changes.

![Ignore settings](<../../../.gitbook/assets/Screenshot 2021-12-07 at 11.25.49.png>)

## Using ignores in reports

If you have access to the Snyk Reports feature, you can also see an overview of how many issues in your organization’s projects are ignored, along with an option to filter these so you can drill down into each one. If the issue was ignored in the Snyk Web UI, Snyk includes a credit for additional accountability, so you can see who initiated it.

For more information, see [Reports](../../snyk-reports/).
