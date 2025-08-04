# Guidance for Snyk for .NET

## Dependency analysis

In the .NET ecosystem, there are multiple levels of dependencies, some of which are obvious and some completely hidden to a developer. To correctly identify the vulnerabilities for a given .NET application, these dependencies must be resolved accurately.

Snyk resolves dependencies differently in the Snyk CLI and the Source Code Management (SCM) systems such as GitHub.

In the CLI, if you manage your Project dependencies using `PackageReference`, Snyk scans your `obj/project.assets.json`. if you manage your dependencies using `packages.config`, Snyk scans the `packages` directory. This approach contributes to accuracy of scan results.

{% hint style="info" %}
Runtime dependencies (provided by the runtime environment also known as "meta-packages") are resolved more accurately in the CLI if the host machine uses a similar runtime SDK to the server running the app.
{% endhint %}

In an SCM integration, scanning uses a different process, as the generated files mentioned above are not available. To overcome this, Snyk follows the NuGet dependency [resolution algorithm](https://docs.microsoft.com/en-us/nuget/concepts/dependency-resolution) to construct a dependency tree.

## Build-time versus runtime dependencies

* Build-time dependency: Snyk understands build-time dependency to be resolved during build time and not susceptible to change at runtime.
* Runtime dependency: Snyk understands runtime dependency to be resolved against the installed runtime, for example, packages coming from the .NET framework (<=4) / .NET [runtime](https://docs.microsoft.com/en-us/dotnet/core/versions/selection?WT.mc_id=DOP-MVP-5001511&) (for Core and .NET 5+) such as [`System.Net.Http`](https://www.nuget.org/packages/System.Net.Http). Snyk sometimes refers to runtime dependencies as meta-packages.

You can choose one of the following actions to address **vulnerabilities from runtime dependencies**. These vary between the SCM and the CLI.

### Vulnerabilities from runtime dependencies in SCM

If you believe you have found false positives because the application runs on a system that always has the latest patches from Microsoft installed, which _may_ mean the vulnerability is no longer relevant to your Project, you may choose to [ignore](../../../manage-risk/prioritize-issues-for-fixing/ignore-issues/) it.

### Vulnerabilities from runtime dependencies in CLI

If you believe you have found false positives because when the application runs in production, you always pull the latest/explicit patches from Microsoft, which may mean the vulnerability is no longer relevant to your Project, you may [ignore](../../../manage-risk/prioritize-issues-for-fixing/ignore-issues/) them and do the following:

* If, in production your application always runs on the latest SDK patch version, you can set `TargetLatestRuntimePatch` to `true` in the Project file. Make sure to upgrade your environments (for example, dev, prod) to the latest runtime version.

```
<TargetLatestRuntimePatch>true</TargetLatestRuntimePatch>
```

* You may choose to publish a [self-contained](https://docs.microsoft.com/en-us/dotnet/core/deploying/#publish-self-contained) app that includes the runtime. Then set `RuntimeFrameworkVersion` to the specific patch version in the Project file. You may choose to [ignore](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/ignore-vulnerabilities-using-the-snyk-cli.md) vulnerabilities that you believe are no longer relevant.

```
<PropertyGroup>
  <RuntimeFrameworkVersion>5.0.7</RuntimeFrameworkVersion>
</PropertyGroup>
```

Use this guide to apply Snyk effectively in your technology stack.

## C\#

Snyk Code can analyze your C# code using IDE, CLI, and Git integration.

For framework support, see [Snyk Code - Supported languages and frameworks](../../supported-languages-package-managers-and-frameworks.md).

## Nuget

**Target Frameworks**: Snyk identifies the target frameworks and presents results against each identified version using the git integration.

**Development dependencies**: Snyk generally does not scan developer dependencies, as they are not usually pushed to production and are seen as "noise."\
Enable visibility in Nuget git import using the **Settings** > **Languages** > **.Net** settings (see [Git settings for .NET](./#git-settings-for-.net)).\
Snyk scans and fixes the build and `development Dependency` sections of your `*.proj`, `packages.config` and `project.json` files

**Lock files**: Currently, **packages-lock.json** is not supported. Snyk interacts with the build system to determine the installed dependencies.

**PackageReference**: Snyk currently requires a version attribute. If your Project lacks this, Snyk may fail to open a PR for your Project.

**Git analysis**

How dependency trees are created:

* For .NET Core, using the `*.proj` files
* For .NET Framework, using the `*.proj` file, and `packages.config`

SCM integrations support the following:

* `*.csproj`
* `*.fsproj`
* `*.vbproj`
* `packages.config`

Fix Pull Requests

* If you currently manage your Project dependencies with NuGet and leverage [`PackageReference`](https://docs.microsoft.com/en-us/nuget/consume-packages/package-references-in-project-files) or [`packages.config`](https://docs.microsoft.com/en-us/nuget/reference/packages-config) Snyk can automatically update the dependency version in your manifest file, provided there is an actual fix for it. You can then review and merge your fixes.

**CLI Analysis**

The CLI supports the following config files:

**project.assets.json**

Snyk can scan project.assets.json to determine dependencies, but the file must be generated. Similarly, if you point to the solution file (.sln), you must generate the file first.

Run `dotnet restore` to generate the necessary `project.assets.json` before running the `snyk test` command.

The solution file contains pointers to the files necessary to perform the analysis. Note that the projects themselves must have `project.assets.json` files to be scanned. If you want Snyk to use the solution file as an entry point for scanning, you can point the Snyk CLI to the solution file by using `--file=<filename>.sln`.

Where multiple target frameworks are used in the same Project, the CLI scan returns results for the first target framework declared in the Project.

**packages.config**

Run `nuget install -OutputDirectory packages` before running the **`snyk test`** command.

{% hint style="info" %}
Runtime dependencies (provided by the runtime environment known as "meta-packages") are resolved more accurately in the CLI if the host machine uses a similar runtime SDK to the server running the app.
{% endhint %}

## Paket

Snyk can analyze dependencies managed by Paket through the CLI. Keep in mind that the paket.dependencies and paket.lock must be present.

For more information on Paket support, see [Snyk for .NET](./).

## Other

Snyk provides custom test APIs for your unique dependency management strategies. See the endpoint [List issues for a package](../../../snyk-api/reference/issues.md#get-orgs-org_id-packages-purl-issues).

## Build-time versus runtime dependencies

Navigate to the [.NET](./) page for more information.

## Snyk CLI tips and tricks

Navigate to the [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/)​ page for more details.

## What to test

Use the `--help` option in the CLI for details of Snyk CLI commands. Navigate to the [CLI commands and options summary](../../../developer-tools/snyk-cli/cli-commands-and-options-summary.md) page for more details.

### **Open Source libraries**

For open source analysis in the CLI, first, install the dependencies. After installing the dependencies, use one of the following options to run `snyk test`:

* `--file=`: Choose this option to target the solution file (.sln) or a specific file.
* &#x20;`--all-projects`: Choose this option to analyze your open source projects, especially if multiple languages, package managers, and .sln files are involved.

### **Codebase**

Use the `snyk code test` command from the root of the Project to perform source code analysis.

### **Containers**

[Supported Operating System Distributions](../../../scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md)

[Snyk CLI for container security](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/)

### Infrastructure as Code

Navigate to the [Infrastructure as Code security](https://snyk.io/product/infrastructure-as-code-security/) page for more details.

### **Helpful options and plugins**

* To help generate reports locally or at build time, see [snyk-to-html](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md).
* See the `--json` and `--sarif` options for generating output that can be programmatically accessed.
* For advanced filtering options, see[ snyk-filter](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md).

## Additional security topics for .Net developers

The following is a collection of articles from the Snyk Security team and Developer Relations related to this ecosystem.

* [Snyk Blog](https://snyk.io/blog/)
* [Snyk for .Net](./)
* [Best Practices for Containerizing .NET applications](https://snyk.io/blog/best-practices-for-containerizing-net-applications/)
