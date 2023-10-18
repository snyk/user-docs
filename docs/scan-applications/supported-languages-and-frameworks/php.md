# PHP

## Supported frameworks and package managers

{% hint style="warning" %}
You might encounter false positives or false negatives for partially supported frameworks and package managers.
{% endhint %}

### Code analysis

{% hint style="info" %}
Interfile is currently not supported. The data flow is monitored within a single file, not between multiple files.
{% endhint %}

Snyk Code supports the following frameworks:

* Symfony
* Laravel
* PHP Standards

### Open source and licensing

#### Open source policy

To manage licenses from your developer workflows through policy, see the following topics:

* [Defining a secure open source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

#### Open source license compliance

To check compliance for open source licenses, see [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/scan-application-code/snyk-open-source/licenses/getting-started-snyk-licensing-compliance).

#### Open source supported features

{% hint style="info" %}
**Feature availability**\
Features might not be available, depending on your plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

| Package managers / Features         | <p>CLI</p><p>support</p> | <p>Git</p><p>support</p> | License scanning | Fix PRs |
| ----------------------------------- | ------------------------ | ------------------------ | ---------------- | ------- |
| [Composer](https://getcomposer.org) | ✔︎                       | ✔︎                       | ✔︎               |         |

After Snyk has built the dependencies tree, Snyk uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the packages anywhere in the dependency tree.

To scan your dependencies, you must ensure you have first installed the relevant package manager, and that your Project contains the supported manifest files.

The way Snyk analyzes and builds the dependencies tree varies depending on the language and package manager of the Project, as well as the location of your Project.&#x20;

## Getting started with Snyk for PHP across environments

### Snyk CLI&#x20;

#### Prerequisites

1. [Create a Snyk account](../../getting-started/quickstart/create-a-snyk-account/).
2. [Install Snyk CLI and authenticate your machine](../../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine).
3. [Set the default Organization for all Snyk tests](../../scan-application-code/snyk-code/cli-for-snyk-code/set-the-snyk-organization-for-the-cli-tests/setting-the-default-organization-for-all-cli-tests.md) (code analysis).
4. Ensure you have installed the relevant package manager before you begin using the Snyk CLI (open source).
5. Ensure you have included the relevant manifest files supported by Snyk before testing.

#### Code analysis

To start testing your code using Snyk Code open your repository in a terminal and run the following  command:

```javascript
snyk code test
```

To customize test options, run other commands, exclude directories and files, and explore the results in different formats, see the following:

* [Snyk CLI commands](../../snyk-cli/commands/#available-commands)
* [Exclude directories and files from the Snyk tests](../../scan-application-code/snyk-code/cli-for-snyk-code/excluding-directories-and-files-from-the-snyk-code-cli-test.md)
* [Explore test results in a JSON or SARIF format in the terminal ](../../scan-application-code/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/outputting-the-test-results-to-json-or-sarif-format-in-the-terminal.md)
* [Exporting the test results to a JSON or SARIF file](../../scan-application-code/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/exporting-the-test-results-to-a-json-or-sarif-file.md)

#### Open source and licensing

To build the dependency tree, Snyk analyzes the `composer.json` and `composer.lock` files to identify the dependencies and their versions.

There are no unique options for use when running Snyk for PHP.

For more information about Snyk CLI see [Getting started with the CLI](../../snyk-cli/getting-started-with-the-snyk-cli.md).

### Snyk Web UI (Git repository integration)

PHP Projects can be imported from any of the Git services Snyk supports. After Projects have been imported, Snyk analyzes your Projects based on their supported manifest files.

Once you select a Project for import, Snyk builds the dependency tree based on these manifest files. Both of the following files are required:

* `composer.json`
* `composer.lock`

If the `composer.lock` file is not present in the repository, the import will not process the composer.json manifest.

By default, Snyk scans your production dependencies. From the Snyk Web UI you can configure whether to include your development dependencies (`require_dev`) in the scan for vulnerabilities.

**To update language preferences:**

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Select **Settings**, then **Languages**.
3. Select **Edit settings** for PHP and select **Scan dev dependencies** to set your PHP projects in the specific Organization to include both development and production dependencies.
4. Select **Update settings**.

These settings are applied to all newly imported Projects and to all existing Projects when they are re-tested.

#### What's next?

* [Open a Fix PR](php.md#open-a-fix-pr)&#x20;
* [Configure PR Checks](../../scan-application-code/run-pr-checks/configure-pr-checks.md)

### Snyk integrations&#x20;

:link: For integrated development environments, see [Use Snyk in your IDE](../../integrations/ide-tools/).

:link: If you prefer continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software (see [Snyk CI/CD](../../integrations/snyk-ci-cd-integrations/) and [Snyk API](../../snyk-api/)).

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;

