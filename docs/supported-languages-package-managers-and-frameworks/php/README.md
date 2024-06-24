# PHP



## Snyk for PHP support

Snyk supports [PHP for code analysis](php-for-code-analysis.md) and [PHP for open source](php-for-open-source.md).

## Snyk CLI for PHP

There are no unique options for use when running Snyk for PHP.

## Git repositories and PHP

PHP Projects can be imported from any of the Git services Snyk supports. After Projects have been imported, Snyk analyzes your Projects based on their supported manifest files.

Once you select a Project for import, Snyk builds the dependency tree based on these manifest files. Both of the following files are required:

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

