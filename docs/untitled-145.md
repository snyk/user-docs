# Azure Pipelines integration

* [ Configure your Continuous Integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004002258-Configure-your-Continuous-Integration/README.md)
* [ Continuous Integration: language support](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004032157-Continuous-Integration-language-support/README.md)
* [ AWS CodePipeline integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/4402158184081-AWS-CodePipeline-integration/README.md)
* [ Azure Pipelines integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004127677-Azure-Pipelines-integration/README.md)
* [ Bitbucket Pipelines integration overview](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004032177-Bitbucket-Pipelines-integration-overview/README.md)
* [ CircleCI integration overview](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004002278-CircleCI-integration-overview/README.md)
* [ Configure your CircleCI integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004002298-Configure-your-CircleCI-integration/README.md)
* [ Getting Snyk Orb details from the CircleCI registry](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004032197-Getting-Snyk-Orb-details-from-the-CircleCI-registry/README.md)
* [ GitHub Actions integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360019718618-GitHub-Actions-integration/README.md)
* [ Bitbucket Pipelines integration](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360000921778-Bitbucket-Pipelines-integration/README.md)

  [See more](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/sections/360001152577-CI-CD-integrations/README.md)

## Azure Pipelines integration

Snyk enables security across the Microsoft Azure ecosystem, including for Azure Pipelines, automatically finding and fixing application and container vulnerabilities.

Ready-to-use tasks for Azure Pipelines can be quickly inserted directly from the Azure interface, enabling you to customize and automate your pipelines with no extra coding. Amongst the tasks included is the Snyk task.

You can include the Snyk task in your pipeline to test for security vulnerabilities and licensing issues as part of your routine work; in this way, you can test and monitor your application dependencies and container images for security vulnerabilities. Once tested, you can review and work with results directly from the Azure Pipelines output, as well as from the Snyk interface.

**Snyk support**

Our Snyk Security Scan task is available for all languages supported by Snyk and Azure DevOps.

After the Snyk Security Scan task is added to a pipeline, then each time the pipeline runs, the Snyk task does the following:

**Scan**

1. Scans application dependencies or container images for vulnerabilities and licensing issues and lists them.
2. If Snyk finds vulnerabilities, it does one of the following \(based on your configuration\):
   * Fails the pipeline
   * Lets the pipeline complete the build

**Monitor**

Optionally, if the pipeline continues \(the Snyk task completes successfully\) and **monitorOnBuild** is set to **true** in the task, then Snyk saves a snapshot of the project dependencies on [snyk.io](https://snyk.io/), where you can see the dependency tree with all of the issues, and be alerted if and when new issues are found in the dependencies.

### Install the Snyk extension for your Azure pipelines

To start using our task as part of your pipeline build, first install the extension on your Azure DevOps instance per organization, from the Marketplace.

**Prerequisites:**

* Create a Snyk account.
* Ensure you are an owner of or an administrator for this account.

**Steps:**

1. Access your Snyk account.
2. For free plans, go to your General Account Settings and find, copy and save your personal API authentication token on the side.
3. For paid plans, navigate to the organization you’d like to integrate with, then go to **Settings** to create a new service account token. Copy and save it on the side.
4. Access your Azure DevOps account and navigate to **Extensions -&gt; Browse marketplace.**
5. Search for the **Snyk Security Scan** extension, click **Get it free** .
6. Create a new _Service Connection_ in your project via **Project Settings** —&gt; **Pipelines** —&gt; **Service Connections**
7. Select "**Snyk Authentication**" service connection: 
8. In the Snyk Authentication service connection form, enter the **Server URL** and the **Snyk API Token** along with a **Service connection name**: 
9. Click on **Save**, ensuring the new service connection appears in your list of service connections.

### Add the Snyk Security Task to your pipelines

**Prerequisites**

* Ensure you have a pipeline within the repository for the code you’d like to test.
* If you created a pipeline with the Azure Repos wizard, this file is called `azure-pipelines.yml`.
* If this repository has multiple service connections, check with your Snyk admin which to use for your pipeline.
* If you want to add your Dockerfile for additional base image data when testing your container, ensure the image has first been built.

**Requirements**

This extension requires that Node.js and npm be installed on the build agent. These are available by default on all Microsoft-hosted build agents. However, if you are using a self-hosted build agent, you may need to explicitly activate Node.js and npm and ensure they are in your [PATH](https://en.wikipedia.org/wiki/PATH_%28variable%29). This can be done using the [NodeTool task from Microsoft](https://docs.microsoft.com/en-us/azure/devops/pipelines/tasks/tool/node-js?view=azure-devops) prior to the SnykSecurityScan task in your pipeline.

**Steps:**

1. Add the Snyk Security Scan task when you create your pipeline or while editing an existing one. See the [Azure Pipelines documentation](https://docs.microsoft.com/en-us/azure/devops/pipelines/?view=azure-devops)
2. From Azure, access the pipeline that you’d like to scan for vulnerabilities, open it for editing and check that the Build step is included just before the point at which you’d like to insert the Snyk task \(this is not required but is considered best practice for consistency across projects\).
3. Open the **assistant**, search for the Snyk Security Scan task and click it. The configuration panel opens on top of the assistant. 
4. Complete the fields in the configuration. Find full details about the parameters in this article: [Snyk Security Scan task parameters and values](untitled-145.md).

   If the **Fail build if Snyk finds issue** option is checked, then if the build fails, the pipeline will not be pushed to Snyk even if the **Run Snyk monitor after test?**

   If you remove the checkmark from the **Fail build if Snyk finds issue** option, the Snyk task tests for vulnerabilities, but does not cause the pipeline job to fail.

   When testing a container image, you can specify the path to the Dockerfile with the dockerfilePath property in order to receive additional information about issues in your base image.

   To add your Dockerfile for additional base image data when testing your container, ensure the image has first been built.

5. Place your cursor inside the pipeline, ensuring you place it before a deployment step, such as **npm publish** or **docker push**.

   You can have multiple instances of the Snyk Security Scan task within your pipeline. This might be useful, for example, if you have multiple project manifest files you want to test or if you want to test both the application and the container images.

6. From the configuration panel, click **Add**. The task is inserted to your pipeline where your cursor was placed, appearing similar to the following:

   ```text
   - task: SnykSecurityScan@0
     inputs:
       testType: 'app'
       monitorOnBuild: true
       failOnIssues: true
   ```

7. Once included in your pipeline, the task runs each time the pipeline runs, and the results appear in the Azure Pipelines output view:

   If the Snyk task fails the build, an error message appears in the results indicating that the build failed due to `snyk test`.

### Snyk Security Scan task parameters and values

This section describes the Snyk task parameters for Azure Pipelines integration, their parallel configuration fields \(from the configuration panel in Azure Pipelines\) and their valid values:

<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>Configuration field</b>
        <br />(Parameter)</th>
      <th style="text-align:left"><b>Description</b>
      </th>
      <th style="text-align:left"><b>Required</b>
      </th>
      <th style="text-align:left"><b>Default</b>
      </th>
      <th style="text-align:left"><b>Type</b>
      </th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>Snyk API token<br />service<br /></b>(ConnectionEndpoint)<b><br /></b>
      </td>
      <td style="text-align:left">
        <p>The Azure DevOps service connection endpoint where your Snyk API token
          is defined. Your admin defines this within your Azure DevOps project settings,
          assigning it with a unique string in order to differentiate between different
          connections.</p>
        <p>The configuration panel displays all available Snyk service connections
          from a dropdown list like the following:</p>
        <p>If multiple Snyk service connections are available from the dropdown list,
          ask your administrator which to use for the pipeline you&#x2019;re working
          with.</p>
      </td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">none</td>
      <td style="text-align:left">String / Azure Service Connection Endpoint of type SnykAuth / Snyk Authentication</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>What do you want to test?</b>(testType)<b><br /></b>
      </td>
      <td style="text-align:left">Determines which dynamic fields to display as described in the rest of
        this table.</td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">&quot;application&quot;</td>
      <td style="text-align:left">string: &quot;app&quot; or &quot;container&quot;</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Container Image Name</b>(dockerImageName)</td>
      <td style="text-align:left">
        <p>The name of the container image to test.</p>
        <p>This dynamic field appears when <b>What do you want to test</b> is set to <b>Container Imager</b>
        </p>
        <p>Set to <b>Yes</b> if container image test.</p>
      </td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">none</td>
      <td style="text-align:left">string</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Path to Dockerfile</b>(dockerfilePath)</td>
      <td style="text-align:left">
        <p>The path to the Dockerfile corresponding to the <code>dockerImageName</code>
        </p>
        <p>This dynamic field appears when <b>What do you want to test</b> is set to <b>Container Imager<br /></b>
        </p>
        <p>Set to <b>Yes</b> if container image test.</p>
      </td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">none</td>
      <td style="text-align:left">string</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Custom path to manifest file to test</b>(targetFile)</td>
      <td style="text-align:left">
        <p>Applicable to application type tests only. The path to the manifest file
          to be used by Snyk. Should only be provided if non-standard.</p>
        <p>This dynamic field appears when <b>What do you want to test</b> is set to <b>Application</b>
        </p>
      </td>
      <td style="text-align:left">No</td>
      <td style="text-align:left">none</td>
      <td style="text-align:left">string</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Testing severity threshold</b>(severityThreshold)</td>
      <td style="text-align:left">
        <p>The severity-threshold to use when testing. By default, issues of all
          severity types will be found.</p>
        <p><b>Note</b>: if not cofigured, the default severity is set to <b>Low</b>.</p>
      </td>
      <td style="text-align:left">No</td>
      <td style="text-align:left">&quot;low&quot;</td>
      <td style="text-align:left">string: &quot;low&quot; or &quot;medium&quot; or &quot;high&quot;</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Run Snyk monitor after test?</b>(monitorOnBuild)</td>
      <td style="text-align:left">Whether or not to capture the dependencies of the application / container
        image and monitor them within Snyk.</td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">true</td>
      <td style="text-align:left">boolean</td>
    </tr>
    <tr>
      <td style="text-align:left">
        <p><b>Fail build if Snyk finds issues</b>
        </p>
        <p>(failOnIssues)</p>
      </td>
      <td style="text-align:left">This specifies if builds should be failed or continued based on issues
        found by Snyk.</td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">true</td>
      <td style="text-align:left">boolean</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Project name in Snyk</b>(projectName)</td>
      <td style="text-align:left">A custom name for the Snyk project to be created on snyk.io</td>
      <td style="text-align:left">No</td>
      <td style="text-align:left">none</td>
      <td style="text-align:left">string</td>
    </tr>
    <tr>
      <td style="text-align:left">
        <p><b>Organization name (or ID) in Snyk</b>
        </p>
        <p>(organization)</p>
      </td>
      <td style="text-align:left">Name of the Snyk organisation name, under which this project should be
        tested and monitored</td>
      <td style="text-align:left">No</td>
      <td style="text-align:left">none</td>
      <td style="text-align:left">string</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Test (Working) Directory</b>(testDirectory)</td>
      <td style="text-align:left">Alternate working directory. For example, if you want to test a manifest
        file in a directory other than the root of your repo, you would put in
        relative path to that directory.</td>
      <td style="text-align:left">No</td>
      <td style="text-align:left">none</td>
      <td style="text-align:left">string</td>
    </tr>
    <tr>
      <td style="text-align:left">
        <p><b>Additional command-line args for Snyk CLI (advanced)</b>
        </p>
        <p>(additionalArguments)</p>
      </td>
      <td style="text-align:left">
        <p>Additional Snyk CLI arguments to be passed in. See <a href="https://support.snyk.io/hc/en-us/articles/360003812578-Our-full-CLI-reference#UUID-c88e66cf-431c-9ab1-d388-a8f82991c6e0_section-5dc1ab1f12700-idm46079268377440">CLI reference</a> for
          details.</p>
        <p><b>Tip</b>: add <b>--all-projects</b> as good practice (for example, for
          .NET), if no project has been found.</p>
      </td>
      <td style="text-align:left">No</td>
      <td style="text-align:left">none</td>
      <td style="text-align:left">string</td>
    </tr>
  </tbody>
</table>

**Example of a Snyk task to test a node.js \(npm\) based application**

This section displays examples of Snyk Security Scan task configurations and \[parameters when testing a Node.js \(npm\) application.

The configuration panel appears as follows:

Click **add** and it is added to your pipeline as follows:

**Simple Application Testing Example**

```text
- task: SnykSecurityScan@0
  inputs:
    serviceConnectionEndpoint: 'mySnykToken'
    testType: 'app'
    monitorOnBuild: true
    failOnIssues: true
```

**Example of a Snyktask for a container image pipeline**

The following is an example of the Snyk Security Scan task within the script for a container image pipeline.

When populated with the most common settings, the configuration panel in Azure appears similar to the following:

Following is an example of the same configuration once you've added it to your pipeline.

**Simple Container Image Testing Example**

```text
- task: SnykSecurityScan@0
  inputs:
    serviceConnectionEndpoint: 'mySnykAuth'
    testType: 'container'
    dockerImageName: 'my-container-image-name'
    dockerfilePath: 'Dockerfile'
    monitorOnBuild: true
    failOnIssues: true
```

