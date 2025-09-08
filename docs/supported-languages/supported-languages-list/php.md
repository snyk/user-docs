# PHP

## Applicability and integration

{% hint style="info" %}
PHP is supported for Snyk Code and Snyk Open Source.
{% endhint %}

Available integrations:

* SCM import
* CLI and IDE: test or monitor your app

Available functions:

* Test your app's SBOM using `pkg:composer`
* Test your app's packages using `pkg:composer`

## Technical specifications

Snyk has tested support for PHP versions 5.2 through 8.0. You can also use later versions.

### Supported frameworks and libraries

For PHP, the following frameworks and libraries are supported:

* grpc-php - Comprehensive
* Laravel - Partial
* llphant - Comprehensive
* openai-php/client - Comprehensive
* orhanerday/open-ai - Comprehensive
* Pclzip - Comprehensive
* Symfony - Partial
* theodo-group/llphant - Comprehensive

### Supported package managers and package registries

* Supported package manager: [Composer](https://getcomposer.org)
* Supported package registry:  [packagist.org](https://packagist.org/)

## PHP for Snyk Code

For PHP with Snyk Code, the following file formats are supported: `.php`, `.phtml`, `.module`, `.inc`, `.install`

Available features:

* Reports
* Custom rules
* Interfile analysis

{% hint style="info" %}
The **Snyk Fix PR** feature is not available for Swift and Objective-C. This means that you will not be notified if the PR checks fail when the following conditions are met:

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available.**
* "**Fixed in" available** is set to **Yes.**
{% endhint %}

## PHP for Snyk Open Source

For PHP with Snyk Open Source, the following file formats are supported: `composer.json`, `composer.lock`

Available features:

* License scanning
* Reports

## CLI support for PHP

A [build is required](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md) to scan with the CLI if there is no `composer.lock` file present. There are no unique options for use when running Snyk for PHP.&#x20;

## SCM integration support for PHP

PHP Projects can be imported from any of the available Snyk SCM integrations. After Projects have been imported, Snyk analyzes your Projects based on their supported manifest files.

After you select a Project for import, Snyk builds the dependency tree based on these manifest files. Both of the following files are required:

* `composer.json`
* `composer.lock`

If the `composer.lock` file is not present in the repository, the import will not process the `composer.json` manifest.

By default, Snyk scans your production dependencies. Using the Snyk Web UI, you can configure whether or not to include your development dependencies, such as  `require_dev` \[...] in the scan for vulnerabilities.

To update language preferences:

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Select **Settings** > **Languages**.
3. Select **Edit settings** for PHP and select **Scan dev dependencies** to set your PHP projects in the specific Organization to include both development and production dependencies.
4. Select **Update settings**.

These settings are applied to all newly imported Projects and to all existing Projects when they are re-tested.
