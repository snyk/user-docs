# Multi-tenant settings for Helm chart installation

To use the Helm chart in different multi-tenant regions, set the `brokerServerUrl` for the region you are using.

See [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls) for the list of regional URLs.

Use the following command:

```
--set brokerServerUrl=<broker-region-url>
```
