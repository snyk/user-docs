# How TeamCity integration works

Use the Snyk plugin with your TeamCity Projects to test and monitor your code for vulnerabilities on an ongoing basis, for breaking builds when newly disclosed vulnerabilities related to your Project are announced, and to receive relevant notifications, all based on your configurations.

The steps in the process follow:

1. The TeamCity account admin selects the Snyk plugin for installation.
2. TeamCity installs the plugin on the server in the Plugin directory.
3. The admin enables the plugin.
4. The user creates a Project or updates an existing Project, adding Snyk Security as a build step.
5. The user configures the build, including the configuration of the Snyk Security step, API token, policy changes, and so on.
6. Snyk authenticates your account using the API token you configured in the build.
7. The user runs a build.
8. During the build, before scanning for vulnerabilities, your Snyk installation is verified and updated as necessary in the background, based on your policy configuration.
9. Snyk then analyzes the manifest file of your Project, automatically detecting the Project type to find direct and transitive dependencies, and tests your Project against the Snyk Vulnerability Database for known vulnerabilities.
10. On the **Snyk Security Report** tab, view the test results from TeamCity in the build details. The results indicate the number of known issues and the number of associated dependency paths identified.
11. If you did not choose the **Monitor project on build** configuration setting for this Project:
    1. Snyk displays all vulnerability results and details on the **Snyk Security Report tab** in TeamCity.
    2. If you specified the severity threshold for a severity assigned to any vulnerability identified in your Project, TeamCity breaks the build. Otherwise, TeamCity continues to run the build to completion (success or failure), and Snyk activity ends.
12. If you configured the **Monitor project on build** option, Snyk now runs the `snyk monitor` command and proceeds with the remainder of the steps as described here.
    1. Snyk takes a snapshot of the Project, analyzes the manifest file of your Project to find its direct and transitive dependencies, and tests your Project against the Snyk vulnerability database for known vulnerabilities.
    2. Snyk pushes the snapshot to the Snyk UI. The snapshot displays the Project details and the dependency hierarchy, as well as vulnerability results and fix advice.
    3. If you specified the severity threshold for a severity that is assigned to any vulnerability in your Project, TeamCity breaks the build.
    4. After the snapshot is pushed to the Snyk UI, Snyk continues to monitor your Project as new vulnerabilities are disclosed. Based on your configuration, if vulnerabilities are found, Snyk notifies you by email or Slack so that you can take immediate fix action.

For information on how to configure your build with a Snyk step, see [Team City integration: use Snyk in your build](teamcity-integration-use-snyk-in-your-build.md).
