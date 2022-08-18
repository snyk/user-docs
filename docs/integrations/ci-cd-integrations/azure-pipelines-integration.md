# Azure Pipelines integration

## Overview of Azure Pipelines integration

Snyk enables security across the Microsoft Azure ecosystem including for Azure Pipelines by automatically finding and fixing application and container vulnerabilities.

Ready-to-use tasks for Azure Pipelines can be inserted quickly and directly from the Azure interface, enabling you to customize and automate your pipelines with no extra coding. Among the tasks included is the Snyk task.

You can include the Snyk task in your pipeline to test for security vulnerabilities and licensing issues as part of your routine work. In this way you can test and monitor your application dependencies and container images for security vulnerabilities. When the testing is done you can review and work with results directly from the Azure Pipelines output, as well as from the Snyk interface.

The Snyk Security Scan task is available for all languages supported by Snyk and Azure DevOps.

### **Example of a Snyk task for a container image pipeline**

The following is an example of the Snyk Security Scan task within the script for a container image pipeline.

When populated with the most common settings, the configuration panel in Azure looks much like the following:

![](../../.gitbook/assets/mceclip2-5-.png)

The following is an example of the same configuration once you've added it to your pipeline.

![](../../.gitbook/assets/mceclip3-1-.png)

### **Simple example of testing a container image**

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
