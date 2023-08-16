# Snyk for PHP

Snyk offers security scanning to test your projects for vulnerabilities, both through the [Snyk CLI ](../../../snyk-cli/)and from the Snyk Web UI through different [Snyk Integrations](../../../integrations/).

This page describes how to use Snyk to scan your PHP projects.

## Features of Snyk for PHP

{% hint style="info" %}
**Feature availability**\
Features might not be available, depending on your plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

| Package managers / Features         | <p>CLI</p><p>support</p> | <p>Git</p><p>support</p> | License scanning | Fix PRs |
| ----------------------------------- | ------------------------ | ------------------------ | ---------------- | ------- |
| [Composer](https://getcomposer.org) | ✔︎                       | ✔︎                       | ✔︎               |         |

## **How Snyk for PHP works**

After Snyk has built the tree, Snyk uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in the dependency tree.

In order to scan your dependencies, you must ensure you have first installed the relevant package manager, and that your Project contains the supported manifest files.

The way Snyk analyzes and builds the tree varies depending on the language and package manager of the Project, as well as the location of your project. See [Snyk CLI for PHP Projects](snyk-for-php.md#snyk-cli-for-php-projects) and [Git services for PHP Projects](snyk-for-php.md#git-services-for-php-projects).

## Snyk CLI for PHP Projects

In order to build the dependency tree Snyk analyzes the `composer.json` and `composer.lock` files that it finds to identify the dependencies and their versions.

### **Prerequisites for Snyk CLI with PHP**

* Ensure you have installed the relevant package manager before you begin using the Snyk CLI.
* Ensure you have included the relevant manifest files supported by Snyk before testing.
* [Install](../../../snyk-cli/install-or-update-the-snyk-cli/) and authenticate the Snyk CLI to start analyzing Projects from your local environment.

### **CLI options for use with PHP**

There are no unique options for use when running Snyk for PHP.

For more information about Snyk CLI see [Getting started with the CLI](../../../snyk-cli/getting-started-with-the-cli.md).

## Git services for PHP Projects

PHP Projects can be imported from any of the Git services Snyk supports. Once Projects have been imported, Snyk analyzes your Projects based on their supported manifest files.

Once you select a Project for import, Snyk builds the dependency tree based on these manifest files. Both of the following files are required:

* composer.json
* composer.lock

If the composer.lock file is not present in the repository, the import will not process the composer.json manifest.

## **Git settings for PHP**

By default Snyk scans your production dependencies. From the Snyk UI you can configure whether to include your development dependencies (`require_dev`) in the scan for vulnerabilities.

**To update language preferences:**

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Select **Settings** > **Languages**.
3. Select **Edit settings** for PHP and select **Scan dev dependencies** to set for your PHP Projects in the specific eOrganization to include both development and production dependencies.
4. Select **Update settings**.

These settings are applied to all newly imported Projects and to all existing projects when they are re-tested..

## Error messages for PHP Projects

The following error messages may appear when you are working with PHP Projects:

* composer.json or composer.lock not found in path
* Manifest file not found in path
* Lockfile missing packages property
* Lockfile or manifest file is not a valid JSON

## Support for Snyk for PHP

If you run across any of these, or other issues, send the following files to [support@snyk.io](mailto:support@snyk.io):

* `composer.json`
* `composer.lock`
