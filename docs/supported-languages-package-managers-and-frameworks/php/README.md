# PHP

## Applicability

Snyk supports [PHP for code analysis](php-for-code-analysis.md) and [PHP for open source](php-for-open-source.md).

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.&#x20;

{% hint style="info" %}
**Supported PHP versions**

You can use PHP versions 5.2 up to 8.0.
{% endhint %}

Available functions:

* SCM import.&#x20;
* Test or monitor your app through CLI and IDE.
* Test your app's SBOM using `pkg:composer`&#x20;
* Test your app's packages using `pkg:composer`

## Package managers

This language supports Composer as a package manager and [packagist.org](https://packagist.org/) as a package registry.

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for PHP:&#x20;

* Laravel - Partial&#x20;
* llphant - Comprehensive&#x20;
* openai-php/client - Comprehensive&#x20;
* orhanerday/open-ai - Comprehensive&#x20;
* Pclzip - Comprehensive&#x20;
* Symfony - Partial&#x20;
* theodo-group/llphant - Comprehensive

## Features

The following features are supported in Snyk for PHP:

* License scanning&#x20;
* Reports
* Custom rules&#x20;
* Interfile analysis

## Snyk CLI for PHP

There are no unique options for use when running Snyk for PHP.

## SCM integrations and PHP

PHP Projects can be imported from any of the available Snyk SCM integrations. After Projects have been imported, Snyk analyzes your Projects based on their supported manifest files.

After you select a Project for import, Snyk builds the dependency tree based on these manifest files. Both of the following files are required:

* `composer.json`
* `composer.lock`

If the `composer.lock` file is not present in the repository, the import will not process the composer.json manifest.

By default, Snyk scans your production dependencies. Using the Snyk Web UI, you can configure whether or not to include your development dependencies (`require_dev`) in the scan for vulnerabilities.

To update language preferences:

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Select **Settings**, then **Languages**.
3. Select **Edit settings** for PHP and select **Scan dev dependencies** to set your PHP projects in the specific Organization to include both development and production dependencies.
4. Select **Update settings**.

These settings are applied to all newly imported Projects and to all existing Projects when they are re-tested.

## Troubleshooting Snyk for PHP

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;

