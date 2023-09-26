# .NET

## Supported frameworks and package managers

{% hint style="warning" %}
You might encounter false positives or false negatives for partially supported frameworks and package managers.
{% endhint %}

### Code analysis

Snyk Code supports the following frameworks:

* NET Framework
* ASP.NET (version 6.x)
* .NET Core:

### Open source and licensing

Snyk Open Source provides support for both NuGet and Paket, as outlined below.

{% hint style="info" %}
**Feature availability**\
Features may not be available, depending on your plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

| Package managers / Features                            | CLI support | Git support | License scanning | Fix PRs |
| ------------------------------------------------------ | ----------- | ----------- | ---------------- | ------- |
| [NuGet](https://www.nuget.org)                         | ✔︎          | ✔︎          | ✔︎               | ✔︎      |
| [Paket](https://fsprojects.github.io/Paket/index.html) | ✔︎          |             |                  |         |

{% hint style="warning" %}
Snyk does not currently support PackageReference without a version attribute. If your Project lacks this, Snyk may fail to open a PR for your Project.\
\
The current workaround is to add versions to all PackageReferences.
{% endhint %}

#### Open source policy

To manage licenses from your developer workflows through policy, see the following topics:

* [Defining a secure open source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

#### Open source license compliance

To check compliance for open source licenses, see [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/scan-application-code/snyk-open-source/licenses/getting-started-snyk-licensing-compliance).

## Getting started with Snyk for .NET across environments

### Snyk CLI&#x20;

#### Prerequisites

1. [Create a Snyk account](../../../getting-started/quickstart/create-a-snyk-account/)
2. [Install Snyk CLI and authenticate your machine](../../../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine)
3. [Set the default Organization for all Snyk tests](../../../scan-application-code/snyk-code/cli-for-snyk-code/set-the-snyk-organization-for-the-cli-tests/setting-the-default-organization-for-all-cli-tests.md) (code analysis)

#### Code analysis

To start testing your code using Snyk Code open your repository in a terminal and run the following  command:

```javascript
snyk code test
```

To customize test options, run other commands, exclude directories and files, and explore the results in different formats, see the following:

* [Snyk CLI commands](../../../snyk-cli/commands/#available-commands)
* [Exclude directories and files from the Snyk tests](../../../scan-application-code/snyk-code/cli-for-snyk-code/excluding-directories-and-files-from-the-snyk-code-cli-test.md)
* [Explore test results in a JSON or SARIF format in the terminal ](../../../scan-application-code/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/outputting-the-test-results-to-json-or-sarif-format-in-the-terminal.md)
* [Exporting the test results to a JSON or SARIF file](../../../scan-application-code/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/exporting-the-test-results-to-a-json-or-sarif-file.md)

#### Open source and licensing

The following sections list the steps to start scanning your dependencies. The basic commands are covered, such as `snyk test` and `snyk monitor`. To check the full list, see [CLI commands and options summary](../../../snyk-cli/cli-commands-and-options-summary.md).

{% hint style="info" %}
To scan your dependencies, ensure you install the relevant package manager and that your Project contains the supported manifest files.
{% endhint %}

#### Nuget

#### Dependencies managed by PackageReference

Restore dependencies in the .NET project by running `dotnet restore` and make sure that **obj/project.assets.json** has been created by the previous command. Then run `snyk test`. For more information see [Getting started with the CLI](../../../snyk-cli/getting-started-with-the-snyk-cli.md).

Examples of supported Project files that resolve into **project.assets.json** include:

* **\*.csproj**
* **\*.vbproj**
* **\*.fsproj**

The **project.assets.json** file is required for scanning.

{% hint style="info" %}
Project files can be combined with [lock files](https://docs.microsoft.com/en-us/nuget/consume-packages/package-references-in-project-files#locking-dependencies) for a more deterministic **project.assets.json** resolution.
{% endhint %}

#### Dependencies managed by packages.config

While there are two approaches for dependencies managed by **packages.config**., the following is the recommended approach because it will yield the most accurate results:

First, install the dependencies into the **packages** folder by running `nuget install -OutputDirectory packages` and make sure the **packages** directory has been created by the previous command. Then run `snyk test`.

Examples of supported project files that resolve into **packages** include: **packages.config**

{% hint style="info" %}
While you should also be able to run `snyk test` without previously installing dependencies this will result in less accurate vulnerability results.
{% endhint %}

#### **Nuget CLI options**&#x20;

* For information about the `snyk test` options available for use with NuGet, see [Options for NuGet projects in the Test help](../../../snyk-cli/commands/test.md#options-for-nuget-projects).&#x20;
* For the available `snyk monitor` options, see [Options for NuGet projects in the Monitor help](../../../snyk-cli/commands/monitor.md#options-for-nuget-projects).

#### Paket **CLI options**

* To use Paket, be sure a **paket.lock** file is present in combination with a **paket.dependencies** file.&#x20;
* Run `snyk test`.

#### **CLI options for use with** other dependencies

Other support includes **project.json** (no longer recommended, refer to [Microsoft documentation](https://docs.microsoft.com/en-us/nuget/archive/project-json)).

To build the dependency tree, Snyk analyzes the **paket.dependencies** and **paket.lock** files.

### Snyk Web UI (Git repository integration)

Import .NET Projects from any of the Git services Snyk supports.

When your Projects have been imported, Snyk analyzes your Projects based on their supported manifest files and then builds the dependency tree and displays it in the Snyk Web UI, similar to the following:

<figure><img src="../../../.gitbook/assets/dotNet Project.png" alt=".NET dependency tree in Snyk Web UI"><figcaption><p>.NET dependency tree in Snyk Web UI</p></figcaption></figure>

#### **Nuget**

After you select a Project for import, Snyk builds the dependency tree based on these manifest files:

* For .NET Core, the **\*.proj** files
* For .NET Framework, the **\*.proj** file, and **packages.config**

Examples of supported Project files include:

* **\*.csproj**
* \***.vbproj**
* **\*.fsproj**

A .NET Project can target multiple target frameworks. Snyk creates a separate dependency tree for each target framework, displaying each as a separate Snyk Project from the interface. This makes it easier to understand why a dependency is being used and also to assess the fix strategy.

#### **Paket**

No import support currently.

#### **Git settings for .NET**

From the Snyk Web UI, you can configure Snyk to scan your entire Project, including the build dependencies, or skip the build dependencies.

You can also **update your language preferences.**

1. Log in to your account and navigate to the relevant Group and Organization you want to manage.
2. Go to **Settings** and select settings for **.NET**. To scan all development dependencies, be sure that **Scan build dependencies** are checked.
3.

    <figure><img src="../../../.gitbook/assets/DotNet languages.png" alt=""><figcaption><p>Update your language preferences</p></figcaption></figure>

#### Fixing vulnerabilities for .NET

For a general understanding of how Snyk helps you fix Open Source vulnerabilities within your Projects, see [Fix your vulnerabilities](../../snyk-open-source/manage-vulnerabilities/fix-your-vulnerabilities.md).

{% hint style="info" %}
**Feature availability**\
The Fix PR feature is _only_ available across Snyk [SCM](../../../integrations/git-repository-and-ci-cd-integrations-comparisons.md) integrations.
{% endhint %}

#### Fix PR supported manifest files

If you are currently managing your Project dependencies with NuGet and leveraging [`PackageReference`](https://docs.microsoft.com/en-us/nuget/consume-packages/package-references-in-project-files) or [`packages.config`](https://docs.microsoft.com/en-us/nuget/reference/packages-config) Snyk can automatically update the dependency version in your manifest file, provided there is an actual fix for it. You can then review and merge your fixes.

### Snyk integrations&#x20;

:link: For integrated development environments, see [Use Snyk in your IDE](../../../integrations/ide-tools/).

:link: If you prefer continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software (see [Snyk CI/CD](../../../integrations/snyk-ci-cd-integrations/) and [Snyk API](../../../snyk-api/)).

## Help

For best practices and troubleshooting, see [Help](help-.net.md).

