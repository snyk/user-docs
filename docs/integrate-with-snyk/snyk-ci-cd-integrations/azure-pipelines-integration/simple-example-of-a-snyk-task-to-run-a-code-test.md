# Simple example of a Snyk task to run a code test

The following is a simple example of a Snyk task to run a Snyk Code test.

{% hint style="warning" %}
[Publishing Snyk Code results](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/using-snyk-code-from-the-cli/publishing-cli-results-to-a-snyk-project-and-ignoring-cli-results.md) is not currently supported.
{% endhint %}

```
- task: SnykSecurityScan@1
  inputs:
    serviceConnectionEndpoint: 'snykToken'
    testType: 'code'
    codeSeverityThreshold: 'medium'
    failOnIssues: true
```
