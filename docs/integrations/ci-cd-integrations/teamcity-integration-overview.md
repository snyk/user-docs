# TeamCity integration overview

Integrate the Snyk Security plugin with JetBrains’ continuous integration \(CI\) tool, TeamCity, to embed open source vulnerability scanning directly into your automated build chain. The TeamCity project builds as an additional build step to check for vulnerabilities as part of your build, and thereafter you can easily push your project to Snyk for continuous monitoring.

By scanning as part of your build and then displaying those test results directly from the TeamCity UI, the Snyk plugin enables you to more quickly track, identify and remediate issues that risk your application’s security posture over time, as fixes are made available for vulnerabilities or new vulnerabilities are disclosed.

### Supported languages and repos

Snyk supports all TeamCity projects regardless of which Git repo is used.

All languages supported by both TeamCity and Snyk can be scanned for vulnerabilities by this plugin.

### TeamCity integration: how it works

Use the Snyk plugin with your TeamCity projects to test and monitor your code for vulnerabilities on an ongoing basis, breaking builds when newly disclosed vulnerabilities related to your project are announced and receiving relevant notifications—all based on your configurations.

1. The admin selects the Snyk plugin for installation in their TeamCity account.
2. TeamCity installs the plugin on the server in the Plugin directory.
3. The admin enables the plugin.
4. The user creates a project or updates an existing project, adding Snyk Security as a build step.
5. The user configures build, including the configuration of the Snyk Security step \(API token, policy changes, etc.\).
6. Snyk authenticates your account using the API token you configured in the build.
7. The user runs a build.
8. During the build, before scanning for vulnerabilities, your Snyk installation is verified and/or updated as necessary in the background \(if necessary, and as based on your policy configuration\).
9. Snyk then analyzes the manifest file of your project, automatically detecting project type to find direct and transitive dependencies and test your project against the Snyk vulnerability database for known vulnerabilities.
10. From TeamCity in the Build details, the tab Snyk Security Report displays the test results, indicating the number of known issues and the number of associated dependency paths identified.
11. Based on the Monitor project on build configuration setting for this project:
12. If the user did not choose the option when configuring the step, then Snyk displays all vulnerability results and details from the Snyk Security Report tab in TeamCity.
13. If the severity threshold was defined for a severity that is assigned to any vulnerability identified in your project, TeamCity breaks the build.
14. Otherwise, TeamCity continues to run the build to completion \(success or failure\) and Snyk activity ends.
15. If the user configured the Monitor project on build option, Snyk now runs the `snyk monitor` command and proceeds with the remainder of the steps as described here.
16. Snyk takes a snapshot of the project, analyzes the manifest file of your project to find its direct and transitive dependencies and tests your project against the Snyk vulnerability database for known vulnerabilities.
17. Snyk pushes the snapshot, displaying the project details and the dependency hierarchy from the Snyk UI as well as vulnerability results and remediation advice.
18. If the severity threshold was defined for a severity that is assigned to any vulnerability in your project, TeamCity breaks the build.
19. Once the snapshot is pushed to the Snyk UI, Snyk continues to monitor your project as new vulnerabilities are disclosed. Based on your configurations, if vulnerabilities are found, Snyk notifies you via email or Slack so that you can take immediate remediation action.

### TeamCity integration: use Snyk in your build

For any project, you can add Snyk to your build to scan the code while you build and to fail the build for vulnerabilities, based on your configurations.

We recommend running a build with the Snyk Security step before deployment, to ensure excellent security posture.

For additional information with TeamCity and its features, refer to their documentation.

**How to configure your build with a Snyk step**

1. Add the step to a new or existing project:
   * For new projects, after configuring the Git repo from which to create the build, activate the auto-detect feature to automatically identify relevant steps for your project build.
   * For existing projects, navigate to edit the project build steps. When complete, Snyk Security appears in the list of suggested steps and the current test policy appears in the Parameters Description column:

     ![image2.png](../../.gitbook/assets/uuid-97395df2-f141-6f77-4551-f19397ac0781-en.png)
2. Navigate to configure the Snyk Security step as follows:
   * Click anywhere on the Snyk Security row to access the configuration screen, or
   * for existing projects, click Add build step to access the configuration screen.

     ![image3.png](../../.gitbook/assets/uuid-88e38280-121e-a17b-cfd3-9fde89305b5c-en.png)
3. Configure the TeamCity fields \(Runner type, Step name and Execute Step\).
4. Optionally, click Show advanced options. Additional Snyk parameters are revealed:

   ![image4.png](../../.gitbook/assets/uuid-8f294e8d-ca5e-123b-2992-a98c1e62fd6f-en.png)

5. Configure Snyk Settings and Snyk Tool Settings. For more information see TeamCity configuration parameters.
6. Once configured, run the build. When the Snyk Security step ends successfully, you can navigate to the Snyk Security Report tab to view results within TeamCity and to navigate seamlessly to the Snyk UI for further action:

   ![image5.png](../../.gitbook/assets/uuid-e8b1fd6f-3b49-069c-c9fe-c0948931b141-en.png)

7. From the top of the report, click View on Snyk.io to view the snapshot and vulnerability information directly from our app.

### TeamCity integration: install the Snyk plugin

Install or upgrade the Snyk Security plugin with these steps. Once complete, you’re all set to add a Snyk step to your projects.

## Warning

You must sign up for an account with Snyk before you begin.

**How to install the Snyk plugin**

1. Log in to your TeamCity instance to install the Snyk Security plugin. Configure the Plugins list to Periodically check for plugin updates, in order to ensure regular automatic upgrades in the background.
2. Navigate to the [JetBrains Plugins Repository](https://plugins.jetbrains.com/plugin/12227-snyk-security), search for Snyk and from the Get dropdown list, select to install the plugin for your TeamCity installation.
3. When the following prompt appears, click Install.

   ![image1.png](../../.gitbook/assets/uuid-fe65f4bc-9578-016c-00dd-6ddb97d2ead7-en.png)

4. When the installation ends, the Administration Plugins List loads, notifying the plugin has been uploaded.
5. Ensure the plugin is enabled.

### TeamCity configuration parameters

<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>Parameters</b>
      </th>
      <th style="text-align:left"><b>Description and values</b>
      </th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>Snyk settings</b>
      </td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Severity threshold</td>
      <td style="text-align:left">
        <p>Default: low</p>
        <p>For the first vulnerability found in your build with the threshold as
          configured, the build fails.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Monitor project on build</td>
      <td style="text-align:left">
        <p>Default: ON</p>
        <p>Snyk runs the snyk monitor command during the build, sending a project
          snapshot to the Snyk app and continuing to monitor the project for vulnerabilities
          even after this build.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">File</td>
      <td style="text-align:left">
        <p>Optional.</p>
        <p>If the manifest file is not on the root level, enter the relative path
          to that file.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Organization</td>
      <td style="text-align:left">
        <p>Optional.</p>
        <p>The ID of the Snyk organization to which this project should be associated
          when imported to the UI.</p>
        <p>Copy the Organization ID from the Snyk UI in the Settings area.</p>
        <p>
          <img src="../../.gitbook/assets/uuid-dfede20b-acb5-fc08-8d1d-59e8476240a5-en.png"
          alt="image6.png" />
        </p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Project name</td>
      <td style="text-align:left">
        <p>Optional.</p>
        <p>Enter any unique name for this project to recognize it when viewing from
          the Snyk UI.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Additional parameters</td>
      <td style="text-align:left">
        <p>Optional.</p>
        <p>Enter additional CLI arguments as necessary. See our CLI documentation
          and cheat sheet for additional information.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Snyk tool settings</b>
      </td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Snyk API token</td>
      <td style="text-align:left">
        <p>From the Settings area in the Snyk UI, copy the Org or Personal API token
          or create a service account. This is the token used to authenticate your
          Snyk account when connecting to TeamCity.</p>
        <p>
          <img src="../../.gitbook/assets/uuid-c27d25fc-00a7-f0f4-261c-d0d9f8653d1d-en.png"
          alt="image7.png" />
        </p>
        <p>
          <img src="../../.gitbook/assets/uuid-be0e9602-023b-99a4-f08c-eded5ea77dac-en.png"
          alt="image8.png" />
        </p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Snyk version</td>
      <td style="text-align:left">
        <p>Default: the most recent version</p>
        <p>Select the plugin version to be used in your build if you would like an
          older Snyk CLI version to support the plugin.</p>
        <p>We recommend configuring automatic upgrades and using the most recent
          version.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Use custom build tool path</td>
      <td style="text-align:left">
        <p>Specify which tool instance in your local environment is to be used for
          this build by Snyk.</p>
        <p>Otherwise Snyk auto-detects the tool and locates it in your environment
          based on project type.</p>
      </td>
    </tr>
  </tbody>
</table>

