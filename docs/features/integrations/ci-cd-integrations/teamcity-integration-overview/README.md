# TeamCity integration

Integrate the Snyk Security plugin with the JetBrains continuous integration (CI) tool, TeamCity, to embed open source vulnerability scanning directly into your automated build chain. The TeamCity integration adds an additional build step to scan for vulnerabilities in your project as part of your build. After the scan you can push your project to Snyk for continuous monitoring.

The Snyk plugin scans as part of your build and then displays test results directly from the TeamCity UI. This enables you to track, identify, and fix issues that risk the security of your application more quickly and over time as fixes are made available for vulnerabilities or new vulnerabilities are disclosed.

## Supported languages and repos

Snyk supports all TeamCity projects regardless of which Git repo is used.

All languages supported by both TeamCity and Snyk can be scanned for vulnerabilities by this plugin.

## How TeamCity integration works

Use the Snyk plugin with your TeamCity projects to test and monitor your code for vulnerabilities on an ongoing basis, breaking builds when newly disclosed vulnerabilities related to your project are announced and receiving relevant notifications, all based on your configurations.

1. The TeamCity account admin selects the Snyk plugin for installation.
2. TeamCity installs the plugin on the server in the Plugin directory.
3. The admin enables the plugin.
4. The user creates a project or updates an existing project, adding Snyk Security as a build step.
5. The user configures the build, including the configuration of the Snyk Security step (API token, policy changes, and so on).
6. Snyk authenticates your account using the API token you configured in the build.
7. The user runs a build.
8. During the build, before scanning for vulnerabilities, your Snyk installation is verified and updated as necessary in the background, based on your policy configuration.
9. Snyk then analyzes the manifest file of your project, automatically detecting the project type to find direct and transitive dependencies and test your project against the Snyk vulnerability database for known vulnerabilities.
10. The **Snyk Security Report** tab displays the test results from TeamCity in the Build details. The results indicate the number of known issues and the number of associated dependency paths identified.
11. If the user did not choose the the **Monitor project on build** configuration setting for this project:
    1. Snyk displays all vulnerability results and details from the **Snyk Security Report tab** in TeamCity.
    2. If the severity threshold was specified for a severity that is assigned to any vulnerability identified in your project, TeamCity breaks the build. Otherwise, TeamCity continues to run the build to completion (success or failure) and Snyk activity ends.
12. If the user configured the **Monitor project on build** option, Snyk now runs the `snyk monitor` command and proceeds with the remainder of the steps as described here.
    1. Snyk takes a snapshot of the project, analyzes the manifest file of your project to find its direct and transitive dependencies, and tests your project against the Snyk vulnerability database for known vulnerabilities.
    2. Snyk pushes the snapshot to the Snyk UI. The snapshot displays the project details and the dependency hierarchy as well as vulnerability results and fix advice.
    3. If the severity threshold was specified for a severity that is assigned to any vulnerability in your project, TeamCity breaks the build.
    4. Once the snapshot is pushed to the Snyk UI, Snyk continues to monitor your project as new vulnerabilities are disclosed. Based on your configuration, if vulnerabilities are found, Snyk notifies you by email or Slack so that you can take immediate fix action.

For information on how to configure your build with a Snyk step, see [Team City integration: use Snyk in your build](teamcity-integration-use-snyk-in-your-build/).

## Install the Snyk plugin

Install or upgrade the Snyk Security plugin by following these steps. When the installation is complete, you can add a Snyk step to your projects.

**Note:** Before you begin, sign up for a Snyk account.

1. Log in to your TeamCity instance to install the Snyk Security plugin. Configure the **Plugins list** to **Periodically check for plugin updates**, in order to ensure regular automatic upgrades in the background.
2. Navigate to the [JetBrains Plugins Repository](https://plugins.jetbrains.com/plugin/12227-snyk-security), search for Snyk, and from the **Get** dropdown list, select the plugin for your TeamCity installation.
3. In response to the prompt, click **Install**.
4. When the installation ends, and the **Administration Plugins List** loads notifying you that the plugin has been uploaded, ensure the plugin is enabled.

![Install plugin from the JetBrains Plugins Repository](../../../../.gitbook/assets/uuid-fe65f4bc-9578-016c-00dd-6ddb97d2ead7-en.png)

To configure the integration, see [TeamCity configuration parameters](teamcity-configuration-parameters.md). For information on how to configure your build with a Snyk step, see [Team City integration: use Snyk in your build](teamcity-integration-use-snyk-in-your-build/).
