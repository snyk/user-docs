# Jenkins plugin integration with Snyk

Snyk offers a native plugin for Jenkins that is based on the [Snyk CLI](https://docs.snyk.io/snyk-cli/cli-reference), to test and monitor Projects for vulnerabilities in your pipelines.

{% hint style="warning" %}
The Snyk Jenkins plugin supports Snyk Open Source. If you plan to include Snyk Code, Snyk Container, and Snyk IaC scans in your pipeline, use the generic [Snyk CLI](https://docs.snyk.io/snyk-cli/cli-reference).
{% endhint %}

For more information, [see the Snyk Jenkins Plugin repository](https://github.com/jenkinsci/snyk-security-scanner-plugin).

Follow the steps in each section of this document to use the Snyk Jenkins plugin:

1. Install the Snyk Security Jenkins Plugin.
2. Configure a Snyk installation.
3. Configure a Snyk API token credential.
4. Add Snyk Security to your Project.
5. Run a build and view your Snyk report.

## 1. Install the Snyk Security Jenkins Plugin

* From your Jenkins dashboard, go to **Manage Jenkins**, **Plugins** and select the **Available plugins** tab.
* Search for **Snyk Security**.
* Install the plugin.

## 2. Configure a Snyk installation

* Go to **Manage Jenkins**, **Tools**.
* Add a **Snyk Installation**.
* Configure the Installation.
* Remember the **Name** for use when configuring the build step.

### Automatic installations

The plugin can download the latest version of Snyk binaries and keep them up-to-date for you.

<figure><img src="../../.gitbook/assets/snyk_config_auto-update_v2 (1).png" alt="Snyk Jenkins plugin automatic installation"><figcaption><p>Snyk Jenkins plugin automatic installation</p></figcaption></figure>

### Manual installations

* Download the following binaries. Choose the binary suitable for your agent's operating system:
  * [Snyk CLI](https://github.com/snyk/snyk/releases/latest)
  * [snyk-to-html](https://github.com/snyk/snyk-to-html/releases/latest)
* Place the binaries in a single directory on your agent.
  * Do not change the filename of the binaries.
  * Ensure you have the correct permissions to execute the binaries.
* Provide the absolute path to the directory under **Installation directory**.

<figure><img src="../../.gitbook/assets/snyk_config_manual_v2.png" alt="Snyk Jenkins plugin manual installation"><figcaption><p>Snyk Jenkins plugin manual installation</p></figcaption></figure>

### Custom API endpoints

By default, Snyk uses the https://api.snyk.io endpoint. It is possible to configure Snyk to use a different endpoint by changing the `SNYK_API` environment variable:

* Go to **Manage Jenkins**, **System**.
* Under **Global Properties** check the **Environment** variables option.
* Click **Add**.
* Set the name to SNYK\_API and the value to the custom endpoint.

For more information see [Configure Snyk CLI to connect to Snyk API](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli/configure-snyk-cli-to-connect-to-snyk-api).

## 3. Configure a Snyk API token credential

* Get your [Snyk API Token](https://docs.snyk.io/snyk-api-info/authentication-for-api).
* Go to **Manage Jenkins**, **Credentials**.
* Choose a Store.
* Choose a Domain.
* Go to **Add Credentials**.
* Select **Snyk API Token**.
* Configure the Credentials.
* Remember the **ID** for use when configuring the build step.

<figure><img src="../../.gitbook/assets/snyk_configuration_token_v2.png" alt="Configure Snyk API token for Jenkins plugin"><figcaption><p>Configure Snyk API token for Jenkins plugin</p></figcaption></figure>

## 4. Add Snyk Security to your Project

This step depends on whether you are using Freestyle Projects or Pipeline Projects.

### Freestyle Projects

* Select a Project.
* Go to **Configure**.
* Under **Build**, select **Add build step** and **Invoke Snyk Security Task**.
* Configure as needed. Click the **?** icons for more information about each option.

<figure><img src="../../.gitbook/assets/snyk_buildstep_freestyle.png" alt="Configure for Freestyle Project"><figcaption><p>Configure for Freestyle Project</p></figcaption></figure>

### Pipeline Projects

Use the `snykSecurity` step as part of your pipeline script. You can use the **Snippet Generator** to generate the code from a web form and copy it into your pipeline. Refer to the following example.

```
pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        echo 'Building...'
      }
    }
    stage('Test') {
      steps {
        echo 'Testing...'
        snykSecurity(
          snykInstallation: '<Your Snyk Installation Name>',
          snykTokenId: '<Your Snyk API Token ID>',
          // place other parameters here
        )
      }
    }
    stage('Deploy') {
      steps {
        echo 'Deploying...'
      }
    }
  }
}
```

You can pass the following parameters to your`snykSecurity` step.

**`snykInstallation` (required)**

Snyk Installation Name, as configured in step 2. Configure a Snyk installation.

**`snykTokenId` (optional, default: \_none**\_**)**

Snyk API Token Credential ID., as configured in step 3. Configure a Snyk API token credential.

If you prefer to provide the Snyk API token another way, such as using alternative credential bindings, you must provide a `SNYK_TOKEN` build environment variable.

**`failOnIssues` (optional, default: `true`)**

Whether the step should fail if issues and vulnerabilities are found.

**`failOnError` (optional, default: `true`)**

Whether the step should fail if Snyk fails to scan the project due to an error. Errors include scenarios like: failing to download Snyk's binaries, improper Jenkins setup, bad configuration, and server errors.

**`monitorProjectOnBuild` (optional, default: \_none\_)**

Whether to monitor the Project on every build by taking a snapshot of its current dependencies on Snyk.io. Selecting this option will keep you notified about newly disclosed vulnerabilities and remediation options in the Project.

**`organization` (optional, default: \_automatic**\_**)**

The Snyk Organization in which this Project should be tested and monitored. See `--org` in the [Snyk CLI commands and options summary](https://docs.snyk.io/snyk-cli/cli-reference) for the default behavior.

**`projectName` (optional, default: \_automatic**\_**)**

A custom name for the Snyk Project created for this Jenkins project on every build. See `--project-name` in the [CLI commands and options summary](https://docs.snyk.io/snyk-cli/cli-reference) for default behavior.

**`targetFile` (optional, default: \_automatic**\_**)**

The path to the manifest file to be used by Snyk. See `--file` in the [CLI commands and options summary](https://docs.snyk.io/snyk-cli/cli-reference) for default behavior

**`severity` (optional, default: \_automatic**\_**)**

The minimum severity to detect. Can be one of the following: `low`, `medium`, `high`, `critical`. See `--severity-threshold` in the [CLI commands and options summary](https://docs.snyk.io/snyk-cli/cli-reference) for default behavior.

**`additionalArguments` (optional, default: \_none**\_**)**

See the [CLI commands and options summary](https://docs.snyk.io/snyk-cli/cli-reference) for information on additional CLI options.

## 5. View your Snyk Security Report

* Complete a new build of your Project.
* Navigate to the build's page.
* Select **Snyk Security Report** in the sidebar to see the results.

<figure><img src="../../.gitbook/assets/snyk_build_report.png" alt="Snyk Security Report for Jenkins plugin"><figcaption><p>Snyk Security Report for Jenkins plugin</p></figcaption></figure>

If there are any errors, you may not see the report. Refer to the Troubleshooting section that follows.

## Troubleshooting

### Increase Logging

To see more information on your steps, you can increase logging and re-run your steps.

* View the **Console Output** for a specific build.
* Add a logger to capture all `io.snyk.jenkins` logs. See [How do I create a logger in Jenkins for troubleshooting and diagnostic information?](https://docs.cloudbees.com/docs/cloudbees-ci-kb/latest/client-and-managed-masters/how-do-i-create-a-logger-in-jenkins-for-troubleshooting-and-diagnostic-information)
* Add `--debug` to **Additional Arguments** to capture all Snyk CLI logs. Debug output is available under **Console Output** for your build.

### Failed Installations

By default, Snyk Installations will download Snyk binaries over the network from `downloads.snyk.io` and use `static.snyk.io` as a fallback. If this fails, there may be a network or proxy issue. If you cannot fix the issue, you can do a manual installation instead. See step 2. Configure a Snyk installation.
