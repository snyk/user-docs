---
description: How to set the broker server URL for your region in a Snyk Broker Helm chart installation
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Multi-tenant settings for Helm chart installation

To use the Helm chart in different multi-tenant regions, set the `brokerServerUrl` for the region you are using.

See [Broker URLs](https://docs.snyk.io/snyk-data-and-governance/regional-hosting-and-data-residency#broker-server-urls) for the list of regional URLs.

Use the following command:

```
--set brokerServerUrl=<broker-region-url>
```
