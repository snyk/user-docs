---
description: A simple example of a Snyk task to test a container image in Azure Pipelines
nav_context: agnostic
---

# Simple example of a Snyk task to test a container image

The following is a simple example of a Snyk task to test a container image.

```
- task: SnykSecurityScan@1
  inputs:
    serviceConnectionEndpoint: 'snykToken'
    testType: 'container'
    dockerImageName: 'goof'
    dockerfilePath: 'Dockerfile'
    monitorWhen: 'always'
    failOnIssues: true
```
