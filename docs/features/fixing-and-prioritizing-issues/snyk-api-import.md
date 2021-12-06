---
description: https://github.com/snyk-tech-services/snyk-api-import
---

# Snyk-API-Import

Snyk API project importer. This script is intended to help import projects into Snyk with a controlled pace utilizing available [Snyk APIs](https://snyk.docs.apiary.io) to avoid rate limiting from Github/Gitlab/Bitbucket etc and to provide a stable import. The script will kick off an import in batches, wait for completion and then keep going. Any failed requests will be retried before they are considered a failure and logged.
