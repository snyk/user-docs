# Custom rules (beta)

Create custom rules to run queries against the code stack as part of your investigation workflow.

You can save successful queries and implement them in your regular static analysis scans to enable custom rules to be triggered every time Snyk Code analyzes your code.

{% hint style="info" %}
**Feature availability**

Snyk Code custom rules is in [Open Beta](../../../more-info/snyk-feature-release-process.md#open-beta) and only available for Enterprise plans through Snyk Preview (see [Enable custom rules](getting-started-with-custom-rules.md)).
{% endhint %}

{% hint style="warning" %}
Custom rules beta is not currently compatible with [SCM integrations via Broker](../../../enterprise-setup/snyk-broker/#integrations-with-snyk-broker).
{% endhint %}

## When to use custom rules

Currently, you can create net new rules not currently supported by Snyk's coverage.

Some typical use cases that custom rules can be used for consist of but are not limited to the following:

* Defining a custom vulnerable method ([Sink](how-custom-rules-work.md#sink)) that security teams may be worried about.
* Creating regular expression scans to check for use of secrets and credentials that should not be part of the code.
* Determining whether certain unwanted methods are being called within the code base that is deemed unsafe by the security teams.
* After creating a query, you can test it against a [code snippet](run-query.md#run-query-on-a-code-snippet) or [repository](run-query.md#run-query-on-a-repository) you've previously imported to Snyk. This way, you can see the results of your query before you run regular scans. You can use this feature to validate a rule before pushing it to production and ensure that it provides the expected results.

