# Azure DevOps Configurational Changes

## AWS CodePipeline integration

## About the integration

Snyk integrates seamlessly with AWS CodePipeline to scan your application for open source security vulnerabilities and help you deliver secure applications via continuous delivery service, allowing CodePipeline users to make security an automated part of their build, test, and deploy phases.

## Setup

Snyk’s AWS CodePipeline integration can be initiated directly within the AWS CodePipeline console. Connect your Snyk account with just a few clicks. From the AWS CodePipeline console, you can add Snyk to a new or existing pipeline with a few simple steps.

### Step 1: Add stage

At any point after the Source stage, you can add a Scan stage. Click **Edit**, and **add a scan stage**.

### Step 2: Add action group

Then, click Add an Action Group to open the Edit Action window. Here you can name the action, click the Action Provider pulldown arrow, select Snyk, and click **Connect with Snyk** to begin the connection process.

With the Snyk scan stage, you can test your application at different stages of the CI/CD pipeline depending on your stage configuration.

### Step 3: Connect to Snyk

Select how you would like to authenticate with Snyk to give AWS CodePipeline permission to begin scanning your open-source code.

### Step 4: Configure settings

Configure the integration with severity thresholds for issues to be reported. For example, you may want to fail builds only for high severity vulnerabilities. You can also choose to scan your full application code or focus on code manifests.

Confirm the connection to Snyk when prompted.

Save the pipeline after successfully connecting to Snyk.

This should configure your Snyk stage in the CodePipeline and you should be ready to test your application. Users should release the latest changes through the CodePipeline options for the latest changes to take effect.

## Settings

The following options are available for configuration:

1. **Snyk Organization:** Select the Snyk organization where findings reports are saved.
2. **Vulnerability handling**: Select to fail a pipeline if a vulnerability is found. If Fail on issues is selected, the pipeline will fail depending on the sub-options selected. The sub-options available are:
   1. **Block deployment for error types**
      1. **All**: Selecting all fails when there is at least one vulnerability that can be either upgraded or patched.
      2. **Upgradable**: Selecting upgradable fails when there is at least one vulnerability that can be upgraded.
      3. **Patchable**: Selecting patchable fails when there is at least one vulnerability that can be patched.
   2. **Block deployment for error severities**: \(**low**\|**medium**\|**high**\|**critical**\) Only report vulnerabilities of provided level or higher.

## Findings Report

Scan results are accessible from the AWS CodePipeline console. These may be accessed by clicking on the **Details** link within the Scan stage:

**Note:** Changes to the Configuration settings can also be made to a previously configured stage by clicking on the **Snyk** link.

Snyk analyzes the application’s manifest file and correlates the list of dependencies with Snyk’s vulnerability database. Snyk provides [detailed reports](https://support.snyk.io/hc/en-us/categories/360000598418-Reports-and-remediation) for your open-source code. You can navigate to Projects and choose View Report to set the frequency with which the project is checked for vulnerabilities.

From **View Report**, you can also choose the Dependencies tab to see which open source dependencies are being used to build the application. By analyzing the application’s manifest file, Snyk builds a full dependency tree, accurately identifying both direct and transitive dependencies \(transitive dependencies account for 78% of the vulnerabilities detected by Snyk\). This enables Snyk to show exactly how a vulnerability was introduced into the application.

