---
description: A simple example of a Snyk task to test the open source dependencies of an application in Azure Pipelines
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Simple example of a Snyk task to test an application

The following is a simple example of a Snyk task to test an application's open-source dependencies (SCA).

```
- task: SnykSecurityScan@1
  inputs:
    serviceConnectionEndpoint: 'snykToken'
    testType: 'app'
    monitorWhen: 'always'
    failOnIssues: true
```
