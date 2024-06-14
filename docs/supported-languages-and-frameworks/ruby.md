# Ruby

## Supported frameworks and package managers

### Code analysis

{% hint style="info" %}
Interfile is currently not supported. The data flow is monitored within a single file, not between multiple files.
{% endhint %}

Snyk Code supports the following frameworks:

* Ruby On Rails

### Open source and licensing

{% hint style="info" %}
**Feature availability**\
Features may not be available, depending on your plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

Snyk supports testing, monitoring, and fixing Ruby Projects in the [CLI ](ruby.md#snyk-cli)and Git [integrations](ruby.md#snyk-integrations) that have their dependencies managed by [Bundler](https://bundler.io/) and comparing the specific dependency versions against the [Ruby vulnerability database](https://snyk.io/vuln?type=rubygems).

Snyk tests all Bundler groups. Currently, it is not possible to exclude certain groups, such as test or development groups.

If your Gemfile needs access to private Gem sources, see [Private Gem sources](ruby.md#private-gem-sources).

{% hint style="warning" %}
Platform-specific packages are currently not supported. If these are present in your `Gemfile.lock`, this can cause an invalid **Fix PR** to be created. If possible, use the non-platform-specific variant of a package.
{% endhint %}

#### Open source policy

To manage licenses from your developer workflows through policy, see the following topics:

* [Defining a secure open source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

#### Open source license compliance

To check compliance for open source licenses, see [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/scan-application-code/snyk-open-source/licenses/getting-started-snyk-licensing-compliance).

## Getting started with Snyk for Ruby across environments

### Snyk CLI&#x20;

#### Prerequisites

1. [Create a Snyk account](../getting-started/quickstart/create-or-log-in-to-a-snyk-account.md)
2. [Install Snyk CLI and authenticate your machine](../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine)
3. [Set the default Organization for all Snyk tests](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/set-the-snyk-organization-for-the-cli-tests.md) (code analysis)

#### Code analysis

To start testing your code using Snyk Code open your repository in a terminal and run the following  command:

```javascript
snyk code test
```

To customize test options, run other commands, exclude directories and files, and explore the results in different formats, see the following:

* [Snyk CLI commands](../snyk-cli/commands/#available-commands)
* [Exclude directories and files from the Snyk tests](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)
* [Explore test results in a JSON or SARIF format in the terminal](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#output-test-results)
* [Exporting the test results to a JSON or SARIF file](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results)

#### Open source and licensing

The following sections list the steps to start scanning your dependencies. The basic commands are covered, such as `snyk test` and `snyk monitor`. To check the full list, see [CLI commands and options summary](../snyk-cli/cli-commands-and-options-summary.md).

{% hint style="info" %}
To scan your dependencies, ensure you install the relevant package manager and that your Project contains the supported manifest files.
{% endhint %}

#### Manifest files supported with Snyk for Ruby

The following manifest files are supported:

* `Gemfile`
* `Gemfile.lock`

{% hint style="info" %}
Snyk requires both files to be present to correctly test, monitor, and fix Ruby projects.
{% endhint %}

#### Supported Ruby versions

| Ruby main version           | Ruby specific version                                         |
| --------------------------- | ------------------------------------------------------------- |
| <h4><code>2.3.X</code></h4> | `2.3.1`, `2.3.6`                                              |
| <h4><code>2.4.X</code></h4> | `2.4.0`, `2.4.1`, `2.4.2`, `2.4.5`, `2.4.6`, `2.4.9`          |
| <h4><code>2.5.X</code></h4> | `2.5.0`, `2.5.1`, `2.5.3`                                     |
| <h4><code>2.6.X</code></h4> | `2.6.1`, `2.6.3`, `2.6.5`, `2.6.6`                            |
| <h4><code>2.7.X</code></h4> | `2.7.2`, `2.7.3`, `2.7.4`, `2.7.5`, `2.7.6`, `2.7.7`, `2.7.8` |
| <h4><code>3.0.X</code></h4> | `3.0.0`                                                       |
| <h4><code>3.1.X</code></h4> | `3.1.0`, `3.1.1`, `3.1.2`, `3.1.3`                            |
| <h4><code>3.2.X</code></h4> | `3.2.0`, `3.2.1`                                              |

### Snyk Web UI (Git repository integration)

You can test your Ruby projects using the Snyk Web UI.&#x20;

#### Fixing vulnerabilities in your Ruby Projects

Snyk can fix vulnerabilities by updating vulnerable gems using `bundle update`after modifying your Gemfile, adhering to the rules you have specified there as far as possible.

In some scenarios, Snyk cannot upgrade all dependencies to non-vulnerable versions. In this case, consider updating the rules in your Gemfile.

{% hint style="info" %}
For Ruby versions < 3.2, Snyk does not currently support pinning a specific version of Ruby in the Gemfile, for example, `ruby "2.7.7".` You must use a more permissive version range that encapsulates all point versions, such as`ruby "~> 2.7.x"`
{% endhint %}

#### **Private Gem sources**

Using private Gem sources should work normally when you are using the Snyk CLI.

When creating Fix PRs for Ruby Projects using private Gem sources, Snyk may need access to the service hosting the Gems to update the file correctly.

### Snyk integrations&#x20;

:link: For integrated development environments, see [Use Snyk in your IDE](../scm-ide-and-ci-cd-workflow-and-integrations/use-snyk-in-your-ide/).

:link: If you prefer continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software (see [Snyk CI/CD](../scm-ide-and-ci-cd-workflow-and-integrations/snyk-ci-cd-integrations/) and [Snyk API](../snyk-api/)).

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;
