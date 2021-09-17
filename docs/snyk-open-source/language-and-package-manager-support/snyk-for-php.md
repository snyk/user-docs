# Snyk for PHP

Snyk offers security scanning to test your projects for vulnerabilities, both through your CLI and through different integrations from our UI.

The following describes how to use Snyk to scan your PHP projects:

## Features

{% hint style="info" %}
**Note**  
Features might not be available, depending on your subscription plan.
{% endhint %}

|  | Package managers/Features | CLI support | Git support | License scanning | Remediation | Runtime monitoring |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| [![composer-logo.svg](../../.gitbook/assets/uuid-3415efde-9bfb-2b07-3781-8169e25a2fcd-en.png)](https://support.snyk.io/hc/article_attachments/360007258958/uuid-3415efde-9bfb-2b07-3781-8169e25a2fcd-en.png/) | [Composer](https://getcomposer.org/) | ✔︎ | ✔︎ | ✔︎ |  |  |

## **How it works**

Once we’ve built the tree, we can use our [vulnerability database](https://snyk.io/vuln/) to find vulnerabilities in any of the packages anywhere in the dependency tree.

{% hint style="info" %}
**Note**  
In order to scan your dependencies, you must ensure you have first installed the relevant package manager, and that your project contains the supported manifest files.
{% endhint %}

The way by which Snyk analyzes and builds the tree varies depending on the language and package manager of the project, as well as the location of your project:

* [Snyk CLI tool for PHP projects](snyk-open-source/language-and-package-manager-support/snyk-for-php/)
* [Git services for PHP projects](https://support.snyk.io/hc/en-us/articles/360003817397-Snyk-for-PHP#UUID-8d36f8e1-8835-5fa6-8e56-8d9f091d54ed/)

## Snyk CLI tool for PHP projects

The way by which Snyk analyzes and builds the tree varies depending on the language and package manager of the project.

In order to build the dependency tree Snyk analyzes the `composer.json` and `composer.lock` files that it finds to analyze the dependencies and their versions.

## **CLI parameters for PHP**

### **Prerequisites**

* Ensure you've installed the relevant package manager before you begin using the Snyk CLI tool.
* Ensure you've included the relevant manifest files supported by Snyk before testing.
* Install and authenticate the Snyk CLI to start analyzing projects from your local environment. Read more about our CLI in [Getting started with the CLI](https://support.snyk.io/hc/articles/360003812458#UUID-6d3e2b39-daa0-f2f1-19d2-b9107b678c81/) as well.

### **Parameters**

There are no unique parameters when running Snyk for PHP.

Read more about our CLI in [Getting started with the CLI](https://support.snyk.io/hc/articles/360003812458#UUID-6d3e2b39-daa0-f2f1-19d2-b9107b678c81).

## Git services for PHP projects

PHP projects can be imported from any of the Git services we support. Once imported, Snyk analyzes your projects based on their supported manifest files.

Once you select a project for import, we build the dependency tree based on these manifest files:

* Composer.json
* composer.lock

## **Git settings for PHP**

By default, Snyk scans your production dependencies. From the Snyk UI you can configure whether to include your development dependencies \(`require_dev`\) in the scan for vulnerabilities.

### **To update language preferences:**

1. Log in to your account and navigate to the relevant group and organization that you want to manage
2. Click on settings ![](../../.gitbook/assets/cog_icon.png)&gt; **Languages**. 
3. Click **Edit settings** for PHP and select **Scan dev dependencies** to set for your PHP projects in the specific organization to include both development and production dependencies. 
4. Click **Update settings**.

These settings will then be applied to all newly imported projects, and once re-tested, to all existing projects.

## Troubleshooting for your PHP projects

## Error messages

The following error messages may appear for you when working with your PHP projects:

* composer.json or composer.lock not found in path
* Manifest file not found in path
* Lockfile missing packages property
* Lockfile or manifest file is not a valid JSON

## Support

If you run across any of these, or other issues, please send the following files to us at [support@snyk.io](mailto:support@snyk.io/) and we'll help you out:

* `composer.json`
* `composer.lock`

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

