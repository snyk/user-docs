# Snyk Code - User Experience

Here are the main features available in the Snyk Code Web UI.

### Issue filtering, sorting, and grouping

You can filter issues by their severity, language, priority score, and more. Along with sorting and grouping issues by the vulnerability type or file to identify very prevalent problems, you can sort by fix in several locations or problematic files that contain multiple issues.

### Priority Score

Priority Score allows you to sort by and prioritize the more important issues by incorporating factors such as issue prevalence, ease of fix, and risk factor into a single risk score.

### Issue data flow

See exactly how the issue flows though your application with a step-by-step visualization from the source to the sink.

### Vulnerability overview

Learn more about the vulnerability with delivered to you, curated content that explains how the vulnerability was created, what the risk factors are, and popular mitigation strategies for it.

### Fix examples

Several examples with links to actual code that show how others in the world are fixing the same issues in similar data flows to gain insight and context on how to fix the problem.

### Create Jira ticket

Export issue information from Snyk directly to your project in Jira to make sure issues are tracked and followed up.

### Ignore issues

Use the **ignore** button to suppress specific warnings and ignore suggested fixes for that issue. For example, you may have deliberately used hard coded passwords to test your routines in test code, or you are aware of an issue but have decided not to fix it.

After you click **ignore**, you are prompted to provide more details:

![](../../../.gitbook/assets/snykcode-ignore-pic2.png)

* Select whether this is to be marked **not vulnerable**, **ignore temporarily** (to suppress the message for now), or **ignore permanently** (to not fix ever)
* Add any notes to yourself or your colleagues in the comment box; Snyk recommends writing a quick explanation for your decision.
* Set a timer for how long to ignore the issue (14, 30, 60, or 90 days, or **ignore forever**).

Click **save** to ignore this issue with the parameters selected. After you ignore an issue, it will not appear in scan results.

{% hint style="info" %}
There is a status selector for ignored issues on the left side and you can include those (and remove or edit the ignore flag) if you want to review ignored results.
{% endhint %}

See [Ignoring issues not prioritized for your project](https://docs.snyk.io/fixing-and-prioritizing-issues/issue-management/ignore-issues) for full details of using the ignore function.

See [Ignoring issues in Snyk Code](https://docs.snyk.io/fixing-and-prioritizing-issues/issue-management/ignore-issues#Ignore-Snyk-Code) for specific details of how Snyk Code processes the ignore function.

#### Excluding files

1. Checks and reads for DeepCode/Snyk ignore specific files `.gitignore` `.dcignore` (if they exist)
2. Using the information obtained in step 1, we are filtering to get only [the following source code files](../snyk-code-language-and-framework-support.md#supported-extensions):
   * We are accessing only the files in the project directory
   * We do not go above the current project directory
3. Files which size is less than 4 MB found in step 2 are bundled and sent to Snyk
