---
description: How to use an existing Kubernetes service account with the Snyk Broker Helm chart
nav_context: agnostic
---

# Service accounts for Helm Chart installation

To use an existing service account, add the following parameters to the install command:

```
--set serviceAccount.create=false \
--set serviceAccount.name=<ENTER_EXISTING_SERVICE_ACCOUNT> \
```
