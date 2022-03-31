# AWS CodePipeline integration

Snyk integrates seamlessly with AWS CodePipeline to scan your application for open source security vulnerabilities and help you deliver secure applications with continuous delivery service. This integration allows CodePipeline users to make security an automated part of their build, test, and deploy phases.

{% hint style="info" %}
Snyk integration is currently available in AWS `sa-east-1` | `ca-central-1` | ap-`southeast-1` | `ap-southeast-2` | `ap-south-1` | `ap-northeast-2` | `ap-northeast-1` | `eu-west-3` | `eu-west-1` | `eu-north-1` | `us-east-1` | `us-west-2` | `eu-west-2` | `eu-central-1` regions. Snyk is actively working on expanding to additional regions.
{% endhint %}

## Language support

Snyk integration with AWS CodePipeline is supported for the following languages:

* JavaScript
* Java
* .NET
* Python
* Ruby
* PHP
* Scala
* Swift/Objective-C
* Go

## Setup

You can initiate Snyk AWS CodePipeline integration directly from the AWS CodePipeline console.\
Add Snyk to a new or existing pipeline using the steps that follow.

### Requirements

Check if your project must be built before the scan in the CodePipeline. If the project needs to be built, you must add a CodeBuild step before the Snyk Step.

|      Language     | Project Type | Build Required |                                            Notes                                           |
| :---------------: | :----------: | -------------- | :----------------------------------------------------------------------------------------: |
|     Javascript    |      npm     | No\*           |   Build only required if no `package-lock.json` file present; run npm install to generate  |
|     Javascript    |     Yarn     | No\*           |      Build only required if no `yarn.lock` file present; run yarn install to generate      |
|        Java       |     Maven    | Yes            |                              Run `mvn install` before testing                              |
|        Java       |    Gradle    | No             |                                                                                            |
|        .NET       |     Nuget    | No\*           |                  Build only required if no `packages.config` file present                  |
|       Python      |      Pip     | No\*           |     Build only required if missing a Snyk config file with the language-settings param     |
|       Python      |   Setup.py   | Yes            |                            Run `pip install -e .` before testing                           |
|       Python      |    Poetry    | No\*           |     Build only required if no `poetry.lock` file present; run `poetry lock` to generate    |
|        Ruby       |    Bundler   | No\*           |   Build only required if no `Gemfile.lock` file present; run `bundle install` to generate  |
|        PHP        |   Composer   | No\*           | Build only required if no `composer.lock` file present; run `composer install` to generate |
|       Scala       |      SBT     | No             |                                                                                            |
|         Go        |  Go Modules  | No             |                                                                                            |
| Swift/Objective-C |   Cocoapods  | No\*           |     Build only required if no `Podfile.lock` file present; run pod install to generate     |

### CodeBuild Step Example

Note the Scan's input artifact must be provided with the build's output artifact as shown in the configuration.

Example of Javascript CodeBuild (`buildspec.yml`):

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

Example of Maven CodeBuild (`buildspec.yml`):

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

### Step 1: Add stage

At any point after the Source stage, you can add a Snyk scan stage, allowing you to test your application at different stages of the CI/CD pipeline.

Click **Edit**, and **Add a Scan Stage**.

![Add scan stage](../../../.gitbook/assets/aws-cp-add-stage.png)

### Step 2: Add action group

Click **Add an Action Group** to open the **Edit Action** window:

![Edit action window](../../../.gitbook/assets/aws-cp-edit-action.png)

Name the action, then select **Snyk** as the **Action Provider**.

Click **Connect with Snyk** to begin the connection process.

### Step 3: Connect to Snyk

Select how you would like to authenticate with Snyk to give AWS CodePipeline permission to begin scanning your open source code.

![Snyk log in screen](../../../.gitbook/assets/snyk-cp-int-config.png)

### Step 4: Configure settings

The following options are available for configuration:

![Snyk AWS CodePipeline configuration options](../../../.gitbook/assets/Snyk\_AWS\_CodePipeline\_Config\_y\_CodePipeline\_-\_AWS\_Developer\_Tools\_png.png)

* **Snyk organization:** Select the Snyk organization where reports of findings are saved.
* **Vulnerability handling**: Define the pipeline behavior if a vulnerability is found. If the `Block deployment when Snyk finds an error` checkbox is checked, the pipeline fails and does not proceed to the next stage in the CodePipeline.
* **Block deployment for vulnerabilities with a minimum severity of**: **low**|**medium**|**high**|**critical**: Report only vulnerabilities of the specified level or higher.
*   **Monitoring behavior on build**: Set the criteria to monitor projects from the AWS CodePipeline. The available options are:

    * **Always monitor**: The project snapshot is created independent of the test results.
    * **When test fails**: The project snapshot is created only when the test fails.
    * **When test passes**: The project snapshot is created only when the test is successful.
    * **Never monitor**: The project snapshot is never created.

    Unless the **Never monitor** option is selected, the **Project to monitor** field is mandatory. This is to prevent any unintentional project overrides due to naming conflicts. The report is created associated with the selected **Snyk organization**.
* **Project to monitor**: Specify the project group name for your projects. This is the same as using the [remote-repo-url](https://support.snyk.io/hc/en-us/articles/360000910677-Snyk-CLI-monitored-projects-are-created-with-IDs-in-the-project-name) option in the CLI. The field does not allow any spaces in the names. This field is mandatory except when the **Never monitor** option has been selected.
* **Auto-detect all projects in the working directory**: Check this checkbox to auto-detect all projects in the AWS CodePipeline. If this option is not selected the plugin tests the first project it finds because it is using the `--all-projects` option to detect all projects.
* **Advanced options** (all of them are optional):
  * **Excluded directories**: This option only appears when **Auto-detect all projects** is checked. Specify sub-directories to exclude; the directories must be comma-separated.
  * **Custom path to manifest file to test**: This option only appears when **Auto-detect all projects** is not checked. You can specify the file path to the manifest file that Snyk is to scan. When this option is omitted Snyk tries to auto-detect the manifest file for your project.
  * **Additional arguments:** A number of additional options are allowed, which can be applied to `test` and `monitor`. The options are as follows: `--dev`, `--detection-depth`, `--prune-repeated-subdependencies`, `--strict-out-of-sync`, `--yarn-workspaces`, `--skip-unresolved`. For more information on these options, see the [CLI reference](https://docs.snyk.io/features/snyk-cli/cli-reference).

{% hint style="info" %}
To change the configuration settings of a previously-configured stage, click the **Snyk** link.
{% endhint %}

Confirm the connection to Snyk when prompted.

![Confirm connection with OAuth](../../../.gitbook/assets/aws-cp-confirm-oauth.png)

Save the pipeline after successfully connecting to Snyk.

This configures your Snyk stage in the CodePipeline, so you can test your application. Release the latest changes through the CodePipeline options for the latest changes to take effect.

## View scan results

You can view scan results in the AWS CodePipeline console, by clicking **Details** in the Scan stage:

![Details in the Scan stage](../../../.gitbook/assets/aws-cp-findings-report.png)

Click the **link to execution details** to view your detailed vulnerability report.

![Link to execution details](../../../.gitbook/assets/image4-2-.png)

## Test report details

Snyk analyzes the manifest file of the application and correlates the list of dependencies with the Snyk vulnerability database. [Snyk provides detailed reports for your open source code](../../reports/reports-overview.md). By analyzing the manifest file, Snyk builds a full dependency tree, accurately identifying both direct and transitive dependencies (transitive dependencies account for 78% of the vulnerabilities detected by Snyk). This enables Snyk to show exactly how a vulnerability was introduced into the application.

![Snyk test report](../../../.gitbook/assets/prototype.png)

{% hint style="info" %}
Reports are stored for 14 days before they expire. Subsequent pipeline runs update the report and reset the retention period.
{% endhint %}
