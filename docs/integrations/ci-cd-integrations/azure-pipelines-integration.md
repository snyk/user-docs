# Azure Pipelines integration

## Overview of Azure Pipelines integration

Snyk enables security across the Microsoft Azure ecosystem including for Azure Pipelines by automatically finding and fixing application and container vulnerabilities.

Ready-to-use tasks for Azure Pipelines can be inserted quickly and directly from the Azure interface, enabling you to customize and automate your pipelines with no extra coding. Among the tasks included is the Snyk task.

You can include the Snyk task in your pipeline to test for security vulnerabilities and licensing issues as part of your routine work. In this way you can test and monitor your application dependencies and container images for security vulnerabilities. When the testing is done you can review and work with results directly from the Azure Pipelines output, as well as from the Snyk interface.

The Snyk Security Scan task is available for all languages supported by Snyk and Azure DevOps.

## Snyk Security Scan task parameters and values

This section describes the Snyk task parameters for Azure Pipelines integration, their parallel configuration fields from the configuration panel in Azure Pipelines, and their valid values.

| <p><strong>Configuration field</strong><br>(Parameter)</p>                                               | **Description**                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | **Required** | **Default**   | **Type**                                                                          |
| -------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------ | ------------- | --------------------------------------------------------------------------------- |
| <p><strong>Snyk API token</strong><br><strong>service</strong><br>(ConnectionEndpoint)<br></p>           | <p>The Azure DevOps service connection endpoint where your Snyk API token is defined. Your admin defines this within your Azure DevOps project settings, assigning it with a unique string in order to differentiate between different connections.</p><p>The configuration panel displays all available Snyk service connections from a dropdown list like the following: <img src="../../.gitbook/assets/uuid-9c6a12b4-2c03-2248-ad0e-c7437a35e142-en.png" alt="image3.png"></p><p>If multiple Snyk service connections are available from the dropdown list, ask your administrator which to use for the pipeline you’re working with.</p> | Yes          | none          | String / Azure Service Connection Endpoint of type SnykAuth / Snyk Authentication |
| <p><strong>What do you want to test</strong> (testType)<br></p>                                          | Determines which dynamic fields to display as described in the rest of this table.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Yes          | "application" | string: "app" or "container"                                                      |
| **Container Image Name** (dockerImageName)                                                               | <p>The name of the container image to test.</p><p>This dynamic field appears when <strong>What do you want to test</strong> is set to <strong>Container Imager</strong></p><p>Set to <strong>Yes</strong> if container image test.</p>                                                                                                                                                                                                                                                                                                                                                                                                        | Yes          | none          | string                                                                            |
| **Path to Dockerfile** (dockerfilePath)                                                                  | <p>The path to the Dockerfile corresponding to the <code>dockerImageName</code></p><p>This dynamic field appears when <strong>What do you want to test</strong> is set to <strong>Container Imager</strong><br></p><p>Set to <strong>Yes</strong> if container image test.</p>                                                                                                                                                                                                                                                                                                                                                                | Yes          | none          | string                                                                            |
| **Custom path to manifest file to test** (targetFile)                                                    | <p>Applicable to application type tests only. The path to the manifest file to be used by Snyk. Should only be provided if non-standard.</p><p>This dynamic field appears when <strong>What do you want to test</strong> is set to <strong>Application</strong></p>                                                                                                                                                                                                                                                                                                                                                                           | No           | none          | string                                                                            |
| **Testing severity threshold** (severityThreshold)                                                       | <p>The severity-threshold to use when testing. By default, issues of all severity types are found.</p><p><strong>Note</strong>: if not configured, the default severity is set to <strong>Low</strong>.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                   | No           | "Low"         | string: "low" or "medium" or "high"                                               |
| **When to run Snyk Monitor** (monitorWhen)                                                               | When to run **snyk monitor** to capture the dependency tree of the application or container image and monitor it within Snyk.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Yes          | "always"      | string: "always", "onIssuesFound", or "never"                                     |
| **Fail build if Snyk finds issues** (failOnIssues)                                                       | This specifies if pipeline jobs should be failed or continued based on issues found by Snyk.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Yes          | true          | boolean                                                                           |
| **Project name in Snyk** (projectName)                                                                   | A custom name for the Snyk project to be created on snyk.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | No           | none          | string                                                                            |
| **Organization name (or ID) in Snyk** (organization)                                                     | Name of the Snyk organization under which this project should be tested and monitored                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | No           | none          | string                                                                            |
| **Test (Working) Directory** (testDirectory)                                                             | Alternate working directory. For example, if you want to test a manifest file in a directory other than the root of your repo, you would put in a relative path to that directory.                                                                                                                                                                                                                                                                                                                                                                                                                                                            | No           | none          | string                                                                            |
| <p><strong>Additional command-line args for Snyk CLI (advanced)</strong></p><p>(additionalArguments)</p> | <p>Additional Snyk CLI arguments to be passed in. See <a href="https://docs.snyk.io/snyk-cli/guides-for-our-cli/cli-reference">CLI reference</a> for details.</p><p><strong>Tip</strong>: add <strong>--all-projects</strong> as good practice (for example, for .NET), if no project has been found.</p>                                                                                                                                                                                                                                                                                                                                     | No           | none          | string                                                                            |

### Custom API Endpoints

By default, the task uses the [https://snyk.io/api](https://snyk.io/api) endpoint. It is possible to configure Snyk to use a different endpoint by set a `SNYK_API` environment variable in the pipeline:

![](<../../.gitbook/assets/Screenshot 2022-07-22 at 17.36.54.png>)

Please refer to the [Snyk documentation](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli#configuration-to-connect-to-the-snyk-api) for more information about environment configuration.

### **Example of a Snyk task to test a node.js (npm) based application**

This section shows examples of Snyk Security Scan task configurations and parameters for testing a Node.js (npm) application.

The configuration panel appears as follows:

![Snyk Security Scan configuration panel](../../.gitbook/assets/mceclip0-24-.png)

Click **add** and the task is added to your pipeline as follows:

![Snyk Security Scan task added to a pipeline](../../.gitbook/assets/mceclip1-15-.png)

### **Simple example of testing an application**

```
- task: SnykSecurityScan@0
  inputs:
    serviceConnectionEndpoint: 'snykToken'
    testType: 'app'
    monitorWhen: 'always'
    failOnIssues: true
```

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
