# Custom rules (beta)

{% hint style="info" %}
**Feature availability**

The Snyk Code custom rules feature is in [Open Beta](../../../more-info/snyk-feature-release-process.md#open-beta) and available only for Enterprise plans through Snyk Preview. See [Enable custom rules](introducing-custom-rules.md).
{% endhint %}

{% hint style="warning" %}
Custom rules beta is not currently compatible with [SCM integrations via Broker](../../../enterprise-setup/snyk-broker/#integrations-with-snyk-broker).
{% endhint %}

This feature allows you to create custom rules to run queries against the code stack as part of your investigation workflow.

You can save successful queries and implement them in your regular static analysis scans to enable custom rules to be triggered every time Snyk Code analyzes your code.

Currently, you can create net new rules not currently supported by Snyk's coverage.

Custom rules can be used to accomplish the following. This list provides examples only; uses are not limited to the ones identified here.

* Define a custom vulnerable method ([Sink](how-custom-rules-work.md#sink)) that security teams may be worried about.
* Create regular expression scans to check for use of secrets and credentials that should not be part of the code.
* Determine whether certain unwanted methods are being called within the code base that is deemed unsafe by the security teams.
* After creating a query, test it against a [code snippet](run-query.md#run-query-on-a-code-snippet) or [repository](run-query.md#run-query-on-a-repository) you have previously imported to Snyk. This way, you can see the results of your query before you run regular scans. You can use this feature to validate a rule before pushing it to production and ensure that it provides the expected results.

