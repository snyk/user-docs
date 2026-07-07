# Integrate with Jenkins

Configure a Jenkins CI/CD pipeline to scan your application for vulnerabilities.

## Overview

With the Snyk API & Web plugin, you can automatically start a scan every time your Jenkins pipeline runs.

Jenkins lets you have an arbitrary number of build and test scenarios, but a common pattern is as follows:

* **Build step**: Compiles the application or creates the containers.
* **Deploy step**: Sends the compiled code or the containers to a test server and runs them.
* **Test step**: Runs tests on the running application.

The Snyk plugin is a build task that runs after Jenkins builds and deploys the application. Snyk recommends that the security tests run after the integration or functional tests pass, to ensure the application is working properly. A broken application can cause security tests to miss vulnerabilities, because a particular feature is not working.

## Installation and setup

Installing and setting up the plugin takes less than five minutes.

1. Open Jenkins and click **Manage Jenkins**.
2. Click **Manage Plugins**.
3. Click the **Available** tab.
4. In the **Filter** search box, enter **probely**.
5. Select the **Probely Security Scanner** plugin.
6. Click **Download now and install after restart**.
7. After Jenkins restarts, the plugin is installed. Continue reading to set up the required API key from Snyk.

## Generate an API key

Before you use the plugin, generate an API key so that Jenkins can start a scan with Snyk.

After you create the API key, take note of its value, because you need it to configure the plugin credentials later, and Snyk does not display it again. You also need the ID of the target you want to scan. You can obtain this ID from the target page.

## Configure the plugin

You can use the plugin in both Freestyle and Pipeline projects. This article provides an example for each one. You can learn more about these two project types and their differences in the Jenkins documentation.

### Configure credentials

1. Click **Credentials**.
2. Click the down arrow near **(global)** to open the dropdown menu and select **Add credentials**.
3. In the **Kind** dropdown menu, select **Secret text**.
4. Enter the API key in the **Secret** text box.
5. Enter a value for the credentials in the **ID** text box (for example, **probely-test-site**).
6. Enter an optional description and click **OK**.

### Use the plugin in a Freestyle project

Creating a Freestyle project is the simplest way to have a repeatable process to build and test your application, especially for simple applications with a few jobs.

If you already have a Freestyle project, you only need to configure the plugin. On the project listing page, select **Configure** in the dropdown menu next to the project name. To create a new project, follow these steps:

1. Click **New Item**.
2. Enter your project name, select **Freestyle Project**, and click **OK**.
3. In the **Build** section, add **Probely scan step**.

This assumes you have configured the other project options properly, such as checking out from your SCM, building the code, and deploying it.

In the **Probely Security Scanner** section you just added:

1. Add the **Target ID** of the target you want to scan.
2. Select the credentials you configured in Configure credentials. If the connection to the Snyk API is working correctly and the credentials are valid, the message **Credentials verified successfully** appears.
3. After you configure all steps properly, click **Save**.

The next time the build job for this project runs, Snyk tests the security of the configured target and sends you an email with the scan results at the end.

### Use the plugin in a Pipeline project

Pipeline projects are the most flexible and powerful way to create CI/CD pipelines with Jenkins.

The projects need a configuration file, a **Jenkinsfile**. The one in this example uses the more modern declarative syntax instead of the imperative one.

1. Click **New Item**.
2. Enter your project name, choose **Pipeline Project** and click **OK**.
3. Create a Jenkinsfile:

```groovy
pipeline {
    agent {
        docker {
            image 'maven:3-alpine' 
        }
    }
    stages {
        stage('Unit tests') { 
            steps {
                sh './gradlew check'
            }
        }
        stage('Scan with Probely') {
            steps {
                probelyScan targetId: '9nl6yy0TWWKv', credentialsId: 'probely-test-site'  
            }
         }
    }
}
```

As with the Freestyle project, the security tests run after the functional tests, in this case after the Unit tests stage, to ensure the application is working properly.

1. Configure Jenkins to use the Jenkinsfile in your repository.

If your Jenkinsfile is stored in a repository already configured here (Definition **Pipeline script from SCM**), you only need to commit the updated file to the repository.

Click **Save**. Your pipeline now has a stage that scans your target for vulnerabilities.

On the target settings page, you have several options for how Snyk scans your target. You can choose different scan profiles, configure authentication to scan behind login pages, add custom headers, and enable automatic synchronization with your project in Jira.
