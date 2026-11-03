---
nav_context: agnostic
description: >-
  Snyk support for PHP with Snyk Code and Snyk Open Source, including supported
  versions, frameworks, and interfile analysis
last_reviewed: 2026-09-14
last_reviewed_source: focused-edit
---

# PHP

{% hint style="info" %}
Snyk supports PHP for Snyk Code and Snyk Open Source.
{% endhint %}

## PHP for Snyk Code

Snyk Code supports PHP versions 5.2 through 8.0 and is designed to process code from newer PHP versions where feasible.

For an overview of the supported security rules, visit [PHP rules](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-code/snyk-code-security-rules/php-rules).

### Supported frameworks and libraries

For PHP, Snyk supports the following frameworks and libraries:

* grpc-php
* Laravel
* llphant
* openai-php/client
* orhanerday/open-ai
* Pclzip
* Symfony
* theodo-group/llphant

### Supported file formats

Snyk supports the following file formats: `.php`, `.phtml`, `.module`, `.inc`, `.install`, `.theme`, `.profile`.

### Available features

* Reports
* Interfile analysis

### Template file analysis

Snyk Code parses template files and follows data from your application code into the template, so it reports a cross-site scripting vulnerability that becomes exploitable only where the template renders its output.

For PHP, Snyk Code supports the following template engines:

* Twig

For file extensions and the full set of supported language and template engine combinations, visit [Template file analysis](../technical-specifications-and-guidance.md#template-file-analysis).

## PHP for Snyk Open Source

For PHP with Snyk Open Source, Snyk supports PHP versions 5.2 through 8.5.

For PHP with Snyk Open Source, Snyk supports the following file formats: `composer.json` and `composer.lock`

### Available integrations

* SCM import
* CLI and IDE: test or monitor your app

### Supported package managers and package registries

* Supported package manager: [Composer](https://getcomposer.org)
* Supported package registry: [packagist.org](https://packagist.org/)

### Available features

* License scanning
* Reports
* Test your app's SBOM and packages using `pkg:composer` PURLs through the [SBOM test](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/sbom-test) CLI command

{% hint style="info" %}
The **Snyk Fix PR** feature is not available for PHP. This means that Snyk does not notify you if the PR checks fail when the following conditions are met:

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available**.
* **"Fixed in" available** is set to **Yes**.
{% endhint %}

## CLI support for PHP

If no `composer.lock` file is present, a [build is required](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli) to scan with the CLI. Snyk for PHP has no unique CLI options.

## SCM integrations for PHP

You can import PHP Projects from any of the available Snyk SCM integrations. After you import Projects, Snyk analyzes your Projects based on their supported manifest files.

After you select a Project for import, Snyk builds the dependency tree based on these manifest files. Both of the following files are required:

* `composer.json`
* `composer.lock`

If the `composer.lock` file is not present in the repository, the import does not process the `composer.json` manifest.

By default, Snyk scans your production dependencies. Using the Snyk Web UI, you can configure whether to include your development dependencies, such as `require_dev` \[...] in the scan for vulnerabilities.

To update language preferences:

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Select **Settings** > **Languages**.
3. Select **Edit settings** for PHP and select **Scan dev dependencies** to set your PHP Projects in the specific Organization to include both development and production dependencies.
4. Select **Update settings**.

Snyk applies these settings to all newly imported Projects and to all existing Projects when they are retested.
