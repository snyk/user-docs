# Integrate with Jenkins

Configure Jenkins CI/CD pipeline to scan your application for security issues.

## Overview

With the Snyk API & Web plugin, you can automatically start a scan every time your Jenkins pipeline is executed.

Jenkins allows you to have an arbitrary number of build and test scenarios, but a common pattern is as follows:

* **Build step**: Compiles the application or creates the Docker containers.
* **Deploy step**: Sends the compiled code or the containers to test server and executes them.
* **Test step**: Executes tests on the running application.

The Snyk API & Web plugin is a build task that should run after the application is built and deployed. It is recommended that the security tests run after the integration or functional tests pass, to ensure the application is working properly. A broken application may cause security tests to miss vulnerabilities because a particular feature is not working.

## Installation and setup

Installing and setting up the plugin takes less than five minutes.

1. Open Jenkins and click **Manage Jenkins**.
2. Click **Manage Plugins**.
3. Click the **Available** tab.
4. On the **Filter** search box, enter **probely**.
5. Select the **Probely Security Scanner** plugin.
6. Click **Download now and install after restart**.
7. After Jenkins restarts, the plugin is installed. Continue reading to set up the required API key from Snyk API & Web.

## Generate an API key

Before using the plugin, you must generate an API Key for Jenkins to be able to start a scan with Snyk API & Web.

Once the API key is created, take note of its value, as it is required to configure the Plugin credentials later on, and it is not displayed again. You also need the ID of the target you want to scan. You can obtain this ID from the target page.

## Configure the plugin

The plugin can be used in both Freestyle and Pipeline projects, and this article provides an example for each one. You can learn more about these two project types and their differences in the Jenkins documentation.

### Configure credentials

1. Click **Credentials**.
2. Click the down arrow near **(global)** to enable the dropdown menu and choose **Add credentials**.
3. On the Kind dropdown menu, choose **Secret text**.
4. Enter the API key in the **Secret** textbox.
5. Enter a value for the credentials in the **ID** textbox (for example, **probely-test-site**).
6. Enter an optional Description and click **OK**.

### Use the plugin in a Freestyle project

Creating a freestyle is the simplest way to have a repeatable process to build and test your application, especially for simple applications with just a few jobs.

If you already have a freestyle project, you only need to configure the plugin: in the project listing page, click the Configure in the drop-down menu next to the project name. If you are creating a new project, follow the next steps:

1. Click **New Item**.
2. Enter your project name, choose **Freestyle Project** and click **OK**.
3. In the **Build** section add **Probely scan step**.

This assumes you have configured the other project options properly, such as checking out from your SCM, building the code and deploying it.

In the just added **Probely Security Scanner** section:

1. Add the **Target ID** of the target you want to scan.
2. Select the right credentials, which were configured in Configure credentials. If the connection to Snyk API & Web's API is working correctly, and the credentials are valid, you should see the message "Credentials verified successfully".
3. When all steps are properly configured, click **Save**.

The next time the build job for this project runs, Snyk API & Web tests the security of the configured target and sends you an email with the scan results at the end.

### Use the plugin in a Pipeline project

Pipeline projects are the most flexible and powerful way of creating CI/CD pipelines with Jenkins.

The projects need a configuration file, a **Jenkinsfile**. The one in this example uses the more modern declarative syntax, instead of the imperative one.

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

As with the Freestyle project, the security tests are executed after the functional tests, in this case after the Unit tests stage, to ensure the application is working properly.

1. Configure Jenkins to use the Jenkins file on your repository.

If your Jenkinsfile is stored in a repository that was already configured here (Definition **Pipeline script from SCM**), you only need to commit the updated file to the repository.

Click **Save**. Your pipeline now has a stage that scans your target for vulnerabilities.

You have some options on how Snyk API & Web scans your target, at the target settings page: you can choose different scan profiles, configure authentication to scan behind login pages, add custom headers, and enable automatic synchronization with your project at Jira.
