# User Experience

Here are the main features available in Snyk Code's user interface:

### Issue filtering, sorting and grouping

You can filter issues by their severity, language, priority score and more, as well as sorting or grouping issues by the vulnerability type or file to identify very prevalent problems that you can fix in several locations or problematic files which contain multiple issues.

### Priority Score

Priority score allows you to prioritize the more important issues by incorporating factors such as issue prevalence, ease of fix and risk into a single risk score.

### Issue data flow

See exactly how the issue flows though your application with a step-by-step visualization from the source to the sink.

### Vulnerability overview

Learn more about the vulnerability with curated content about how is it created, what's the risk in it and what are some popular mitigation strategies for it.

### Fix examples

See several examples with links to actual code of how others in the world are fixing the same issues in similar data flows to get inspiration for how to remediate the problem.

### Create Jira ticket

Export issue information from Snyk directly to your project in Jira to make sure issues are tracked and followed up on.

### Ignore issues

Use the **Ignore** button to suppress specific warnings, ignoring suggestions for fixes for that issue. For example, you may have deliberately used hard coded passwords to test your routines in test code, or you are aware of an issue but have decided not to fix it.

After you click **Ignore**, you are prompted to provide more details:

![](../../../.gitbook/assets/snykcode-ignore-pic2.png)

Select:

* Whether this is to be marked **Not vulnerable**, **ignore temporarily** (to suppress the message for now), or **ignore permanently** (won’t fix ever)
* Add any notes to yourself or your colleagues in the comment box; we recommend writing a quick explanation for your decision.
* Set a timer for how long to ignore the issue (14, 30, 60, or 90 days, or click **Ignore forever**).

Click **Save** to ignore this issue with the parameters selected. Afer you select to ignore an issue, it does not appear in scan results.

{% hint style="info" %}
There is a status selector for ignored issues on the left side, and you can include those (and remove or edit the ignore flag) if you want to review ignored results.
{% endhint %}

See [Ignoring issues not prioritized for your project](https://docs.snyk.io/fixing-and-prioritizing-issues/issue-management/ignore-issues) for full details of using the Ignore function.

See [Ignoring issues in Snyk Code](https://docs.snyk.io/fixing-and-prioritizing-issues/issue-management/ignore-issues#Ignore-Snyk-Code) for specific details of how Snyk Code processes the Ignore function.

#### Excluding files

1. Checks and reads for DeepCode/Snyk ignore specific files `.gitignore` `.dcignore` (if they exist).
2. Using the information obtained in step 1, we are filtering to get only [the following source code files](../snyk-code-language-and-framework-support.md#supported-extensions):
   * We are accessing only the files in the project directory. We do not go above the current project directory.
3. Files which size is less than 4 MB found in step 2 are bundled and sent to Snyk.
