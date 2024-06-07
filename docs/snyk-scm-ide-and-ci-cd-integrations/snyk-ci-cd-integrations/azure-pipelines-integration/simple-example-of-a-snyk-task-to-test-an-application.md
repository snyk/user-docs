# Simple example of a Snyk task to test an application

The following is a simple example of a Snyk task to test an application's open-source dependencies (SCA).

```
- task: SnykSecurityScan@0
  inputs:
    serviceConnectionEndpoint: 'snykToken'
    testType: 'app'
    monitorWhen: 'always'
    failOnIssues: true
```
