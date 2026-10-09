---
description: A simple example of a Snyk task to run a Snyk Code test in Azure Pipelines
nav_context: agnostic
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
