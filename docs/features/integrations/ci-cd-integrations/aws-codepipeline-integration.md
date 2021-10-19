# AWS CodePipeline integration

Snyk integrates seamlessly with AWS CodePipeline to scan your application for open source security vulnerabilities and help you deliver secure applications via continuous delivery service. This integration allows CodePipeline users to make security an automated part of their build, test, and deploy phases.

{% hint style="info" %}
This integration is currently available in AWS’s `us-east-1` | `us-west-2` | `eu-west-2` | `eu-central-1` regions. We are actively working on expanding to additional regions soon!
{% endhint %}

## Language Support

Snyk integration with AWS CodePipeline is supported for the following languages:

* JavaScript
* Java
* .NET
* Python
* Ruby
* PHP

## Setup

You can initiate Snyk’s AWS CodePipeline integration directly from the AWS CodePipeline console.\
Add Snyk to a new or existing pipeline using the following steps.

## Requirements

Check if your project must be built before the scan in the CodePipeline. If the project is needed to be built, you will need to add a CodeBuild step before the Snyk Step.

|  Language  | Project Type | Build Required |                                            Notes                                           |
| :--------: | :----------: | -------------- | :----------------------------------------------------------------------------------------: |
| Javascript |      NPM     | No\*           |   Build only required if no `package-lock.json` file present, run npm install to generate  |
| Javascript |     Yarn     | No\*           |      Build only required if no `yarn.lock` file present, run yarn install to generate      |
|    Java    |     Maven    | Yes            |                              Run `mvn install` before testing                              |
|    Java    |    Gradle    | No             |                                                                                            |
|    .NET    |     Nuget    | No\*           |                  Build only required if no `packages.config` file present                  |
|   Python   |      Pip     | No\*           |     Build only required if missing a Snyk config file with the language-settings param     |
|   Python   |   Setup.py   | Yes            |                            Run `pip install -e .` before testing                           |
|   Python   |    Poetry    | No\*           |     Build only required if no `poetry.lock` file present, run `poetry lock` to generate    |
|    Ruby    |    Bundler   | No\*           |   Build only required if no `Gemfile.lock` file present, run `bundle install` to generate  |
|     PHP    |   Composer   | No\*           | Build only required if no `composer.lock` file present, run `composer install` to generate |

### CodeBuil Step Example

Please note the Scan's input artifact must be provided with the build's output artifact as shown in the configuration

Example of Javascript's CodeBuild (Buildspec.yml):

```
version: 0.2
phases:
  build:
    commands:
      - npm install
artifacts:
  files:
    - '**/*'
```

Example of Maven CodeBuild (Buildspec.yml):

```
version: 0.2
phases:
  build:
    commands:
      - mvn install
artifacts:
  files:
    - '**/*'
```

## Step 1: Add stage

At any point after the Source stage, you can add a Snyk scan stage, allowing you to test your application at different stages of the CI/CD pipeline.

Click **Edit**, and **add a scan stage**.

![](../../../.gitbook/assets/aws-cp-add-stage.png)

## Step 2: Add action group

Click **Add an Action Group** to open the **Edit Action** window:

![](../../../.gitbook/assets/aws-cp-edit-action.png)

Name the action, then select **Snyk** as the **Action Provider**.

Click **Connect with Snyk** to begin the connection process.

## Step 3: Connect to Snyk

Select how you would like to authenticate with Snyk to give AWS CodePipeline permission to begin scanning your open source code.

![](../../../.gitbook/assets/snyk-cp-int-config.png)

## Step 4: Configure settings

The following options are available for configuration:

![](../../../.gitbook/assets/configure.png)

![](<../../../.gitbook/assets/image8 (1).png>)

1. **Snyk Organization:** Select the Snyk organization where findings reports are saved.&#x20;
2. **Vulnerability handling**: Select to fail a pipeline if a vulnerability is found. If Fail on issues is selected, the pipeline will fail depending on the sub-options selected. The sub options available are:&#x20;
   1. **All**: Selecting all fails when there is at least one vulnerability that can be either upgraded or patched.&#x20;
   2. **Upgradable**: Selecting upgradable fails when there is at least one vulnerability that can be upgraded.&#x20;
   3. **Patchable**: Selecting patchable fails when there is at least one vulnerability that can be patched.&#x20;
3. **Block deployment for vulnerabilities of type**&#x20;
4. **Block deployment for vulnerabilities with a minimum severity of**: (**low**|**medium**|**high**|**critical**) Only report vulnerabilities of provided level or higher.
5. **Snyk Organization:** Select the Snyk organization where findings reports are saved.
6. **Vulnerability handling**: Select to fail a pipeline if a vulnerability is found. If Fail on issues is selected, the pipeline will fail depending on the sub-options selected. The sub options available are:&#x20;
   * **Block deployment for vulnerabilities of type:**
     * **All**: Selecting all fails when there is at least one vulnerability that can be either upgraded or patched.&#x20;
     * **Upgradable**: Selecting upgradable fails when there is at least one vulnerability that can be upgraded.&#x20;
     * **Patchable**: Selecting patchable fails when there is at least one vulnerability that can be patched.&#x20;
   * **Block deployment for vulnerabilities with a minimum severity of**: (**low**|**medium**|**high**|**critical**) Only report vulnerabilities of provided level or higher.
7. **Monitor project**: Select the monitor project checkbox to monitor projects from the AWS CodePipeline. The project snapshot will be created under the Snyk organization selected. Whenever you select the Monitor Project option please note that the Project Group Name will be required. This is to prevent any unintentional project overrides due to naming conflicts.
8. **Project Group Name:** Enter the project group name for your projects.. This is the same as using [--remote-repo-url](https://support.snyk.io/hc/en-us/articles/360000910677-Snyk-CLI-monitored-projects-are-created-with-IDs-in-the-project-name) when using the CLI. The field does not allow any spaces in the names.

{% hint style="info" %}
You can change the Configuration settings of a previously-configured stage, by clicking on the **Snyk** link.
{% endhint %}

Confirm the connection to Snyk when prompted.

![](../../../.gitbook/assets/aws-cp-confirm-oauth.png)

Save the pipeline after successfully connecting to Snyk.

This configures your Snyk stage in the CodePipeline, so you can test your application. Users should release the latest changes through the CodePipeline options for the latest changes to take effect.

## View scan results

You can view scan results in the AWS CodePipeline console, by clicking **Details** in the Scan stage:

![](../../../.gitbook/assets/aws-cp-findings-report.png)

Click **Link to execution details** to view your detailed vulnerability report.

![](../../../.gitbook/assets/image4-2-.png)

## Test report details

Snyk analyzes the application’s manifest file and correlates the list of dependencies with Snyk’s vulnerability database. Snyk provides [detailed reports](https://support.snyk.io/hc/en-us/categories/360000598418-Reports-and-remediation) for your open source code. By analyzing the application’s manifest file, Snyk builds a full dependency tree, accurately identifying both direct and transitive dependencies (transitives account for 78% of the vulnerabilities detected by Snyk). This enables Snyk to show exactly how a vulnerability was introduced into the application.

![](../../../.gitbook/assets/prototype.png)

{% hint style="info" %}
Reports are stored for 14 days before they expire. Subsequent pipeline runs will update the report and reset the retention period.
{% endhint %}
