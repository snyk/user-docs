---
description: Use this documentation to get started with the JetBrains plugin.
---

# JetBrains plugins

Snyk has a plugin for JetBrains IDEs, which support all Snyk products: [Snyk Open Source](https://docs.snyk.io/snyk-open-source), [Snyk Code](https://docs.snyk.io/snyk-code), [Snyk Container](https://docs.snyk.io/products/snyk-container) and [Snyk Infrastructure as Code](https://docs.snyk.io/products/snyk-infrastructure-as-code). The Snyk JetBrains plugin touches on all aspects of securing your application including:

* Security vulnerabilities in open source dependencies (Snyk Open Source).
* Security vulnerabilities and code quality issues in first party code (Snyk Code).
* Configuration issues in your Infrastructure as Code such as Terraform, AWS CloudFormation, Kubernetes, and Azure Resource Manager (ARM) (Snyk IaC)
* Security vulnerabilities in your container images found in Kubernetes workload files (Snyk Container)

{% hint style="info" %}
The Snyk JetBrains plugin is available for installation on the JetBrains marketplace: [https://plugins.jetbrains.com/plugin/10972-snyk-vulnerability-scanner](https://plugins.jetbrains.com/plugin/10972-snyk-vulnerability-scanner).
{% endhint %}

## Supported JetBrains IDEs

Snyk supports JetBrains plugin versions from version 2020.2 on the following IDEs:

* Android Studio
* AppCode
* GoLand
* IntelliJ
* PhpStorm
* PyCharm
* Rider
* RubyMine
* WebStorm

## **How the Snyk JetBrains plugin works**

* The plugin is based on Snyk CLI, but not on the CLI only. The plugin supports product features in the CLI for Snyk Open Source and Snyk Container as well as for Snyk Code and Snyk IaC with some limitations.
* The plugin automatically downloads the CLI in the background; you will be asked to [authenticate](./#authentication).
* Snyk supports all the [languages supported by Snyk Code](https://docs.snyk.io/products/snyk-code/snyk-code-language-and-framework-support#language-support-with-snyk-code-ai-engine). You can install the plugin on any of the IDEs (such as RubyMine). Once the plugin is installed, Snyk analyzes all the language files that it finds.
* If the CLI is already installed on the machine, the plugin uses the token provided to it. Otherwise, you must provide the authentication token through the plugin [authentication mechanism](./#authentication).

## **Install the plugin**

Install using the IDE plugins library:

1. Open the **Preferences** window from the IDE.
2. Navigate to the **Plugins** tab.
3. In the **Plugins** tab, search for **Snyk**.
4. Select the **Snyk vulnerability scanning** plugin.
5. Click on the **Install** button.
6. When the installation is complete, restart the IDE.

![Select the Snyk vulnerability scanning plugin](<../../.gitbook/assets/Screen Shot 2022-03-09 at 5.06.13 PM (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (3).png>)

## Analysis results: Snyk Configuration

Snyk Configuration analysis shows issues in your Terraform, Kubernetes, AWS CloudFormation, and Azure Resource Manager (ARM) code with every scan. Based on the Snyk CLI, the scan is fast and friendly for local development. For more detailed information select an issue.

![Display more information for Snyk Configuration analysis](../../.gitbook/assets/intellij\_iac\_issues.png)

The Snyk plugin provides information so you can quickly understand and fix the underlying issue:

* **Description:** what the misconfiguration is
* **Impact:** how the misconfiguration could potentially be exploited
* **Path:** which path in the tree the issue occurs
* **Remediation:** how to fix the issue
* **References:** where you can investigate deeper from a variety of sources
* **Ignore:** a button to ignore the issue if applicable

## Analysis results: Snyk Container

The plugin scans Kubernetes configuration files and searches for container images. Vulnerabilities are found fast using the extracted container images and comparative analysis against the latest information from the [Snyk Intel Vulnerability Database](https://security.snyk.io).

Snyk Container analysis shows each of the security vulnerabilities to which your image might be vulnerable. For more detailed information select a vulnerability.

A comparison table is displayed with various severity levels such as critical or high. This shows the difference in vulnerabilities between the current image and the image recommended by Snyk, with the same characteristics sorted by severity. This helps you decide if you want to upgrade your image to the recommended one and increase the level of confidence in the image you are running in production.

![Display more information for Snyk Container analysis](../../.gitbook/assets/intellij\_container\_vulnerabilites.png)

## How Snyk Container and Kubernetes integration works

The plugin scans your Kubernetes workload files and collects the images used. To troubleshoot whether a plugin is correctly scanning a container image, you can verify:

* Whether the image definition is in the Kubernetes YAML file in the project. Make sure you have the image is specified with an image name mapped in the format `imageValue:imageKey` for the image yaml attribute, for example, \`image: nginx:1.17.1
* Whether the container image has been successfully built locally and/or pushed to a container registry. It is also a good practice to verify this before referring to the container image in the Kubernetes YAML file.

If you encounter an error [contact support](https://snyk.zendesk.com/agent/dashboard).

For each image found, perform a test with the Snyk CLI.

* Refer to the [doc](https://docs.snyk.io/products/snyk-container/snyk-cli-for-container-security#testing-an-image) for more information about how Snyk Container performs a test on the image.
* While testing the image the CLI downloads the image if it is not already available locally in your Docker daemon.
* Snyk plans to expand the scope of Container scanning, so if there are more files (like Dockerfiles) or workflows that you want to be supported, submit a feature request [to Snyk support](https://support.snyk.io/hc/en-us/requests/new).

## Filter results

### Filter by severity

Snyk reports critical, high, medium and low severities. You can filter for the severity level you need by selecting the value from the dropdown as shown in the screenshot that follows. By default all levels are selected. You must select at least one.

![Select severity level to report](../../.gitbook/assets/filter-severity.png)

### Filter by issue type

Snyk reports the following types of issues:

* **Open Source Vulnerabilities**: found in open source dependencies
* **Security Vulnerabilities**: found in your application’s source code
* **Quality Issues**: found in your application source code
* **Configuration Issues**: found in infrastructure as code files
* **Container Vulnerabilities**: found in images sourced from Kubernetes workload files

You can filter for each one of them by selecting the value from the dropdown as shown in the screenshot that follows. By default all issue types shown are selected.

![Select issue type to support](../../.gitbook/assets/JetBrains-filter-issue-type.png)

##

### Support and contact information

{% hint style="info" %}
Need more help? [Contact Snyk support](https://support.snyk.io/hc/en-us/requests/new).
{% endhint %}

**Share your experience.**

Snyk continuously strives to improve the plugins experience. Would you like to share with us your feedback about the Snyk JetBrains Plugin? [Schedule a meeting](https://calendly.com/snyk-georgi/45min?month=2022-01).
