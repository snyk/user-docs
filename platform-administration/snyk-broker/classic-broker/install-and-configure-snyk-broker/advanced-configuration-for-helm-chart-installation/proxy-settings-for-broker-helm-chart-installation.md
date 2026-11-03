---
description: How to run the Snyk Broker Helm chart behind a proxy using the httpProxy and httpsProxy values
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Proxy settings for Broker Helm chart installation

To use the Helm chart behind a proxy, set the `httpProxy` and `httpsProxy` values.

```
--set httpsProxy=<PROXY_URL>
```
