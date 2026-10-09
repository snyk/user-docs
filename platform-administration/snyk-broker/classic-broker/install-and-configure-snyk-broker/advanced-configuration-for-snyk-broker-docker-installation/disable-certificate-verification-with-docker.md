---
description: How to disable certificate verification for a Snyk Broker Docker installation, for self-signed certificates
nav_context: agnostic
---

# Disable certificate verification with Docker

To disable certificate verification, for example, in the case of self-signed certificates, add the following to the docker run command:

```
-e NODE_TLS_REJECT_UNAUTHORIZED=0
```
