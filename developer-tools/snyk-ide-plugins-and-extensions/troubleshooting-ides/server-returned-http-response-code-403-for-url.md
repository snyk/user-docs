---
description: How to resolve an HTTP 403 error by checking the Snyk endpoint URL and token
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Server returned HTTP response code 403 for URL

Check the endpoint URL and the authentication information.

Most likely the API URL or the token is not ok. They need to use [https://api.snyk.io](https://api.snyk.io) without `/api` . For other tenants, the situation is similar: the path `/api` is not supported on the `api` subdomain.

Even if the URL may have worked before, the URL is invalid and must be changed.
