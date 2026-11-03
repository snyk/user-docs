---
description: A simple example of a Snyk task to run a Snyk Code test in Azure Pipelines
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Simple example of a Snyk task to run a code test

The following is a simple example of a Snyk task to run a Snyk Code test.

```
- task: SnykSecurityScan@1
  inputs:
    serviceConnectionEndpoint: 'snykToken'
    testType: 'code'
    codeSeverityThreshold: 'medium'
    failOnIssues: true
```
