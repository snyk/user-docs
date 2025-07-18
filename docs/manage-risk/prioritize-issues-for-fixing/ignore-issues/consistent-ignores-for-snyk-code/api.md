# Consistent Ignores for Snyk Code API

You can manage ignores individually through the [Snyk Policies API (REST)](https://apidocs.snyk.io/version=2024-10-14~experimental?version=2024-10-15#get-/orgs/-org_id-/policies).&#x20;

The SARIF output from Snyk CLI contains the `snyk/asset/finding/v1` identifier used to manage ignores at the start of the Early Access program.&#x20;

This API leverages the `snyk/asset/finding/v1` identifier and not the `issueId` used by the legacy ignores API. Consider migrating any scripts or automation that rely on the legacy ignores API to the new policy API.
