# PHP

## Applicability

Snyk supports [PHP for code analysis](php-for-code-analysis.md) and [PHP for open source](php-for-open-source.md).

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.

{% hint style="info" %}
**Supported PHP versions**

Snyk has tested support for PHP versions 5.2 through 8.0. You can use later versions as well.
{% endhint %}

Available functions:

* SCM import, available for Snyk Open Source and Snyk Code.
* Test or monitor your app through CLI and IDE, available for Snyk Open Source and Snyk Code.
* Test your app's SBOM using `pkg:composer`
* Test your app's packages using `pkg:composer`

## Package managers and supported file extensions

Snyk for PHP supports Composer as a package manager and [packagist.org](https://packagist.org/) as a package registry and supports the following file formats:

* Snyk Open Source: `composer.json`, `composer.lock`
* Snyk Code: `.php`, `.phtml`, `.module`, `.inc`, `.install`,

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for PHP:

* grpc-php - Comprehensive
* Laravel - Partial
* llphant - Comprehensive
* openai-php/client - Comprehensive
* orhanerday/open-ai - Comprehensive
* Pclzip - Comprehensive
* Symfony - Partial
* theodo-group/llphant - Comprehensive

## Features

The following features are supported in Snyk for PHP:

| Snyk Open Source                                   | Snyk Code                                                                 |
| -------------------------------------------------- | ------------------------------------------------------------------------- |
| <ul><li>License scanning</li><li>Reports</li></ul> | <ul><li>Reports</li><li>Custom rules</li><li>Interfile analysis</li></ul> |

PR Checks that are configured to "Only fail when the issues found have a fix available" rely on Snyk FixPR support and, therefore, will not alert for PHP and other languages that do not support FixPRs.

## Snyk CLI for PHP

A [build is required](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md) to scan with the CLI if there is no `composer.lock` file present. There are no unique options for use when running Snyk for PHP.&#x20;

## SCM integrations and PHP

PHP Projects can be imported from any of the available Snyk SCM integrations. After Projects have been imported, Snyk analyzes your Projects based on their supported manifest files.

After you select a Project for import, Snyk builds the dependency tree based on these manifest files. Both of the following files are required:

* `composer.json`
* `composer.lock`

If the `composer.lock` file is not present in the repository, the import will not process the composer.json manifest.

By default, Snyk scans your production dependencies. Using the Snyk Web UI, you can configure whether or not to include your development dependencies (`require_dev`) in the scan for vulnerabilities.

To update language preferences:

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Select **Settings** > **Languages**.
3. Select **Edit settings** for PHP and select **Scan dev dependencies** to set your PHP projects in the specific Organization to include both development and production dependencies.
4. Select **Update settings**.

These settings are applied to all newly imported Projects and to all existing Projects when they are re-tested.
