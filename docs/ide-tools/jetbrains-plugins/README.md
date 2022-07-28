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

##

##

##

### Support and contact information

{% hint style="info" %}
Need more help? [Contact Snyk support](https://support.snyk.io/hc/en-us/requests/new).
{% endhint %}

**Share your experience.**

Snyk continuously strives to improve the plugins experience. Would you like to share with us your feedback about the Snyk JetBrains Plugin? [Schedule a meeting](https://calendly.com/snyk-georgi/45min?month=2022-01).
