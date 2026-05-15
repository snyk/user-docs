# Azure Pipelines integration using the Snyk Security Scan task

Snyk enables security across the Microsoft Azure ecosystem, including Azure Pipelines, by automatically finding and fixing application and container vulnerabilities.

The Snyk Security Scan task is available for all languages supported by Snyk and Azure DevOps.

{% hint style="info" %}
The Snyk Security Scan task supports Snyk Open Source, Snyk Container, and Snyk Code. If you plan to include other products in your pipeline, use the [Snyk CLI](../../snyk-cli/).
{% endhint %}

Ready-to-use tasks for Azure Pipelines can be inserted quickly and directly from the Azure interface, enabling you to customize and automate your pipelines with no extra coding. Among the tasks included is the Snyk task.

You can include the Snyk task in your pipeline to test for security vulnerabilities and open-source license issues as part of your routine work. In this way, you can test and monitor your application dependencies and container images for security vulnerabilities. When the testing is done you can review and work with results directly from the Azure Pipelines output, as well as from the Snyk interface.

For setup and use details, see the following pages:

* [How the Snyk Security Scan task works](how-the-snyk-security-scan-task-works.md)
* [Install the Snyk extension for your Azure pipelines](install-the-snyk-extension-for-your-azure-pipelines.md)
* [Add the Snyk Security Task to your pipelines](add-the-snyk-security-task-to-your-pipelines.md)
* [Snyk Security Scan task parameters and values](snyk-security-scan-task-parameters-and-values.md)
* [Custom API endpoints](regional-api-endpoints.md)
* [Example of a Snyk task to test a node.js (npm)-based application](example-of-a-snyk-task-to-test-a-node.js-npm-based-application.md)
* [Simple example of a Snyk task to test an application](simple-example-of-a-snyk-task-to-test-an-application.md)
* [Example of a Snyk task for a container image pipeline](example-of-a-snyk-task-for-a-container-image-pipeline.md)
* [Simple example of a Snyk task to test a container image](simple-example-of-a-snyk-task-to-test-a-container-image.md)
* [Example of a Snyk test to test application code](example-of-a-snyk-task-to-test-application-code.md)
* [Simple example of a Snyk task to run code test](simple-example-of-a-snyk-task-to-run-a-code-test.md)
