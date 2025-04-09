# API

{% hint style="info" %}
**Release status**

Snyk Code Consistent Ignores is in Early Access and available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).

To make sure Snyk Code Consistent Ignores Early Access meets your needs and requirements, review [Known limitations](known-limitations.md) and [FAQ](consistent-ignores-for-snyk-code-faqs.md) sections.
{% endhint %}

You can manage ignores individually through the [Snyk Policies API (REST)](https://apidocs.snyk.io/version=2024-10-14~experimental?version=2024-10-15#get-/orgs/-org_id-/policies).&#x20;

The SARIF output from Snyk CLI contains the `snyk/asset/finding/v1` identifier used to manage ignores at the start of the Early Access program.&#x20;

This API leverages the `snyk/asset/finding/v1` identifier and not the `issueId` used by the legacy ignores API. You should prepare to migrate any scripts or automation that rely on the legacy ignores API to the new policy API.
