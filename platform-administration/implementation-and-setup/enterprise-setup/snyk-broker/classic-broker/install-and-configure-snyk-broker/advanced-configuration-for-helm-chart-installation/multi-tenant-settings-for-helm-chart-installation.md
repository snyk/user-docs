# Multi-tenant settings for Helm chart installation

To use the Helm chart in different multi-tenant regions, set the `brokerServerUrl` for the region you are using.

See [Broker URLs](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/ELvljsaLKPkSpffOkmsQ/snyk-data-and-governance/regional-hosting-and-data-residency#broker-server-urls) for the list of regional URLs.

Use the following command:

```
--set brokerServerUrl=<broker-region-url>
```
