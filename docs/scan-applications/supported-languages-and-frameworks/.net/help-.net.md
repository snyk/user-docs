# Best practices for .NET

## Dependency analysis

In the .NET ecosystem, there are multiple levels of dependencies, some of which are obvious and some completely hidden to a developer. To correctly identify the vulnerabilities for a given .NET application, these dependencies must be resolved accurately.

Snyk resolves dependencies differently in the Snyk CLI and the Source Code Management (SCM) systems such as GitHub.

**In the CLI,** if you manage your Project dependencies using `PackageReference`, Snyk scans your `obj/project.assets.json`. if you manage your dependencies using `packages.config`, Snyk scans the `packages` directory. This approach contributes to accuracy of scan results.

{% hint style="info" %}
Runtime dependencies (provided by the runtime environment also known as "meta-packages") are resolved more accurately in the CLI if the host machine uses a similar runtime SDK to the server running the app.
{% endhint %}

**In an SCM integration**, scanning uses a different process, as the generated files mentioned above are not available. To overcome this, Snyk follows the NuGet dependency [resolution algorithm](https://docs.microsoft.com/en-us/nuget/concepts/dependency-resolution) to construct a dependency tree.

For further information on .NET automated fixes, see the [Snyk blog](https://snyk.io/blog/automated-vulnerability-fixes-dot-net-dependencies).

## Build-time versus runtime dependencies

* **Build-time dependency**: Snyk understands build-time dependency to be resolved during build time and not susceptible to change at runtime.
* **Runtime dependency**: Snyk understands runtime dependency to be resolved against the installed runtime, for example, packages coming from the .NET framework (<=4) / .NET [runtime](https://docs.microsoft.com/en-us/dotnet/core/versions/selection?WT.mc\_id=DOP-MVP-5001511&) (for Core and .NET 5+) such as [`System.Net.Http`](https://www.nuget.org/packages/System.Net.Http) . Snyk sometimes refers to runtime dependencies as meta-packages.

### Tackling vulnerabilities from runtime dependencies

There are a couple of actions you can choose to take to address these types of vulnerable dependencies. These vary between the SCM and the CLI.

#### **Vulnerabilities from runtime dependencies in SCM**

If you believe you have found false positives because the application runs on a system that always has the latest patches from Microsoft installed, which _may_ mean the vulnerability is no longer relevant to your Project, you may choose to [ignore](../../../manage-issues/priorities-for-fixing-issues/ignore-issues.md) it.

#### **Vulnerabilities from runtime dependencies in CLI**

If you believe you have found false positives because when the application runs in production, you always pull the latest/explicit patches from Microsoft, which may mean the vulnerability is no longer relevant to your Project, you may [ignore](../../../manage-issues/priorities-for-fixing-issues/ignore-issues.md) them and do the following:

* If, in production your application always runs on the latest SDK patch version, you can set `TargetLatestRuntimePatch` to `true` in the Project file. Make sure to upgrade your environments (for example, dev, prod) to the latest runtime version.

```
<TargetLatestRuntimePatch>true</TargetLatestRuntimePatch>
```

* You may choose to publish a [self-contained](https://docs.microsoft.com/en-us/dotnet/core/deploying/#publish-self-contained) app that includes the runtime. Then set `RuntimeFrameworkVersion`to the specific patch version in the Project file. You may choose to [ignore](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/ignore-vulnerabilities-using-the-snyk-cli.md) vulnerabilities that you believe are no longer relevant.

```
<PropertyGroup>
  <RuntimeFrameworkVersion>5.0.7</RuntimeFrameworkVersion>
</PropertyGroup>
```

Use this guide to apply Snyk effectively in your technology stack.

## C\#

Snyk Code can analyze your C# code using IDE, CLI, and Git integration.&#x20;

For framework support, see [Snyk Code - Supported languages and frameworks](broken-reference).

## Nuget

* **Target Frameworks**: Snyk identifies the target frameworks and presents results against each identified version using the git integration.
* **Development dependencies**: Snyk generally does not scan developer dependencies, as they are not usually pushed to production and are seen as "noise." \
  Enable visibility in Nuget git import using the **Settings > Languages > .Net** settings (see [Git settings for .NET](./#git-settings-for-.net)). \
  Snyk scans and fixes the build and `development Dependency` sections of your [`*.proj`](#user-content-fn-1)[^1], `packages.config` and `project.json` files
* **Lock files**: Currently, **packages-lock.json** is not supported. Snyk interacts with the build system to determine the installed dependencies.
* **PackageReference:** Snyk currently requires a version attribute. If your Project lacks this, Snyk may fail to open a PR for your Project.
*   **Git analysis**

    How dependency trees are created:

    * For .NET Core, using the \*.proj files&#x20;
    * For .NET Framework, using the \*.proj file, and packages.config

    Git integrations support the following:&#x20;

    * \*.csproj&#x20;
    * \*.fsproj
    * \*.vbproj
    * packages.config

    Fix Pull Requests

    * If you currently manage your Project dependencies with NuGet and leverage [`PackageReference`](https://docs.microsoft.com/en-us/nuget/consume-packages/package-references-in-project-files) or [`packages.config`](https://docs.microsoft.com/en-us/nuget/reference/packages-config) Snyk can automatically update the dependency version in your manifest file, provided there is an actual fix for it. You can then review and merge your fixes.
*   **CLI Analysis**

    The CLI supports the following config files:

| project.assets.json                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | packages.config                                                                             |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| <p></p><p>Snyk can scan project.assets.json to determine dependencies, but the file must be generated. Similarly, if you point to the solution file (.sln), you must generate the file first.</p><p></p><p>Run "<strong>dotnet restore"</strong> to generate the necessary <code>project.assets.json</code> before running the "<strong>snyk test</strong>" command.</p><p></p><p>The solution file contains pointers to the files necessary to perform the analysis. Note that the projects themselves must have <code>project.assets.json</code> files to be scanned. If you want Snyk to the solution file as an entry point for scanning, you can point the Snyk CLI to the solution file by using <code>--file=&#x3C;filename>.sln</code>.</p><p></p><p>Where multiple target frameworks are used in the same Project, the CLI scan returns results for the first target framework declared in the Project.</p> | Run "**nuget install -OutputDirectory packages**" before running the **snyk test** command. |

{% hint style="info" %}
Runtime dependencies (provided by the runtime environment known as "meta-packages") are resolved more accurately in the CLI if the host machine uses a similar runtime SDK to the server running the app.
{% endhint %}

## Paket

Snyk can analyze dependencies managed by Paket through the CLI. Keep in mind that the paket.dependencies and paket.lock must be present.

For more information on Paket support, see [Snyk for .net](broken-reference).

## Other

Snyk provides custom test APIs for your unique dependency management strategies.

:link: [Purl issues](https://apidocs.snyk.io/?version=2022-11-14#get-/orgs/-org\_id-/packages/-purl-/issues)&#x20;

## Build-time versus runtime dependencies

See [Snyk for .NET](broken-reference) for more information

## Snyk CLI tips and tricks

:link: [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/)​

## What to test&#x20;

Use the `--help` option in the CLI for details of Snyk CLI commands.

:link: [CLI commands and options summary](../../../snyk-cli/cli-commands-and-options-summary.md)

### **Open Source libraries**

For open source analysis in the CLI, first, install the dependencies. After installing the dependencies, use one of the following strategies when you go to run a test

* `--file=`: This option targets the solution file (.sln) or a specific file.&#x20;
* &#x20;`--all-projects`: This option to analyze your open source, especially if multiple languages/package managers/.sln files are involved.

### **Codebase**

Use the `snyk code test` command from the root of the Project to perform source code analysis.

### **Containers**

:link: [Supported Operating System Distributions](../../snyk-container/how-snyk-container-works/supported-operating-system-distributions.md)

:link: [Snyk CLI for container security](../../snyk-container/use-snyk-container-from-the-cli/)&#x20;

### Infrastructure as Code

:link: [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/)

### **Helpful Options/Plugins**

* To help generate reports locally or at build time, see [snyk-to-html plugin](../../snyk-code/using-snyk-code-from-the-cli/displaying-the-cli-results-in-an-html-format-using-the-snyk-to-html-feature/).
* See `--json` and `--sarif` options for generating output that can be programmatically accessed.
* For advanced filtering options, see[ snyk-filter](../../../snyk-api-info/other-tools/tool-snyk-filter.md).

Additional security topics for .Net developers

The following is a collection of articles from the Snyk Security team and Developer Relations related to this ecosystem.&#x20;

* [Snyk Blog](https://snyk.io/blog/)
* [Snyk for .Net](broken-reference)
* [Best Practices for Containerizing .NET applications](https://snyk.io/blog/best-practices-for-containerizing-net-applications/)



[^1]: 
