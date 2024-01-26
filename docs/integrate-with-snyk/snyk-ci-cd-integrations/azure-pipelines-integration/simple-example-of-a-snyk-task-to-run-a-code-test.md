# Simple example of a Snyk task to run a code test

The following is a simple example of a Snyk task to run a Snyk Code test.

{% hint style="warning" %}
[Publishing Snyk Code results](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/publish-snyk-code-cli-results-and-ignore-issues.md) is not currently supported.
{% endhint %}

```
- task: SnykSecurityScan@1
  inputs:
    serviceConnectionEndpoint: 'snykToken'
    testType: 'code'
    codeSeverityThreshold: 'medium'
    failOnIssues: true
```
