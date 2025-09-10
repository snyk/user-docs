# How ignores work for Projects imported using an SCM and the CLI

When you ignore an issue, you must consider the following factors:

* How was the Project imported: Through an SCM? Through the CLI?
* What Snyk product scanned the Project: Snyk Open Source, Snyk Code, or another?
* How was the ignore created: In a policy? Through the UI from the Project page? Using an API? In a  `.snyk` file?
  * For ignores created in a policy, ignoring an issue on the Organization level will ensure the issues are ignored for Projects imported using an SCM and Projects imported using the CLI.
  * Ignoring for a particular Project or attribute will ensure the issue is ignored only for the specific Project and only for Projects imported using an SCM.
  * For more information, see [Ignore issues in the Snyk Web UI](./#ignore-issues-in-the-snyk-web-ui), [Security policies](../../policies/security-policies/), and [The `.snyk` file](../../policies/the-.snyk-file.md).

Depending on these factors, the ignore is respected for testing in different places. Where will the ignore be respected testing?

* In the UI and an SCM PR test?
* In a CLI test, either local or in a build pipeline, and in an IDE test?

For information about deciding to ignore issues and setting ignores, see I[gnore Issues](./).

{% hint style="info" %}
Test results are the same for an ignore set by policy and an ignore set by the `.snyk` file, whether you have imported the Project through an SCM or the CLI.

If you added an ignore in the UI, support for ignores and the test results differ from support and results for Projects imported through an SCM and through the CLI.
{% endhint %}

## Import through an SCM

The following table summarizes how an ignore will be respected for testing depending on the way you set the ignore for Projects imported through an SCM.

| **Import through an SCM and set ignore by available methods**                                                                                                 | **Ignore respected in UI for testing**                 | **Ignore respected in CLI and IDE tests**              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------ | ------------------------------------------------------ |
| Ignore by policy (add ignore commands to your security policies)                                                                                              | <p>Open Source: ✅<br>Container: ✅</p><p>Code: ❌</p>    | <p>Open Source: ✅</p><p>Container: ✅</p><p>Code: ❌</p> |
| Ignore by UI or API                                                                                                                                           | <p>Open Source: ✅</p><p>Container: ✅</p><p>Code: ✅</p> | <p>Open Source: ❌</p><p>Container: ❌</p><p>Code: ✅</p> |
| Ignore by `.snyk` file (add the issue to a `.snyk` file in the repository; for Open Source, the `.snyk` file must be in the same folder as the manifest file) | <p>Open Source: ✅</p><p>Container: ✅</p><p>Code: ❌</p> | <p>Open Source: ✅</p><p>Container: ✅</p><p>Code: ❌</p> |

## Import through the CLI

The following table summarizes how an ignore will be respected for testing depending on the way you set the ignore for Projects imported through the CLI.

| **Import through the CLI and set ignore by available methods**                                                                                                                                                        | **Ignore respected in UI for testing**                                | **Ignore respected in CLI and IDE tests**                                           |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| Ignore by policy (add ignore commands to your security policies) For an IDE and the CLI the testing must be done in the relevant Organization in order for the policies to be used.                                   | <p>Open Source: ✅</p><p>Container: ✅</p><p>Code: ❌</p>                | <p>Open Source: ✅</p><p>Container: ✅</p><p>Code: ❌</p>                              |
| Ignore by UI or API                                                                                                                                                                                                   | <p>Open Source: ✅</p><p>Container: ✅</p><p>[Early Access] Code: ✅</p> | <p>Open Source: ✅</p><p>Container: ✅</p><p>[Early Access] Code: ✅ CLI    ❌ IDE </p> |
| Ignore by `.snyk` file (add the issue to a `.snyk` file in the repository; for Open Source; the `.snyk` file must be in the same folder as the manifest file; applies to the `snyk test` and `snyk monitor` commands) | <p>Open Source: ✅</p><p>Container: ✅</p><p>Code: ❌</p>                | <p>Open Source: ✅</p><p>Container: ✅</p><p>Code: ❌</p>                              |
