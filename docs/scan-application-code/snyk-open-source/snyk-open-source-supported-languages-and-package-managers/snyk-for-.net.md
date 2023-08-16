# Snyk for .NET

Snyk offers security scanning to test your projects for vulnerabilities, both through the [Snyk CLI ](../../../snyk-cli/)and from the Snyk Web UI through different [Snyk Integrations](../../../integrations/).

This page describes how to use Snyk to scan your .NET Projects.

## Features of Snyk for .NET

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

## **How Snyk for .NET works**

Once Snyk has built the tree, Snyk uses its [vulnerability database](https://security.snyk.io) to find vulnerabilities in any of the packages anywhere in the dependency tree.

{% hint style="info" %}
To scan your dependencies, you must first install the relevant package manager and ensure that your project contains the supported manifest files.
{% endhint %}

The way Snyk analyzes and builds the tree varies depending on the language and package manager of the project, as well as the location of your Project:

* [Snyk CLI for .NET projects](snyk-for-.net.md#snyk-cli-for-.net-projects)
* [Git services for .NET projects](snyk-for-.net.md#git-services-for-.net-projects)

## Snyk CLI for .NET projects

### Nuget

#### Dependencies managed by PackageReference

First, restore dependencies in the .NET project by running `dotnet restore` and make sure that **obj/project.assets.json** has been created by the previous command. Then run `snyk test`. For more information see [Getting started with the CLI](../../../snyk-cli/getting-started-with-the-cli.md).

Examples of supported Project files that resolve into **project.assets.json** include:

* **\*.csproj**
* **\*.vbproj**
* **\*.fsproj**

The **project.assets.json** file is required for scanning

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

#### **CLI options for use with Nuget**

For information about the `snyk test` options available for use with NuGet, see [Options for NuGet projects in the Test help](../../../snyk-cli/commands/test.md#options-for-nuget-projects). For the available `snyk monitor` options, see [Options for NuGet projects in the Monitor help](../../../snyk-cli/commands/monitor.md#options-for-nuget-projects).

### Paket: Dependencies managed by Paket

To use Paket, be sure a **paket.lock** file is present in combination with a **paket.dependencies** file. Then run `snyk test`

Other support includes **project.json** (no longer recommended, refer to [Microsoft documentation](https://docs.microsoft.com/en-us/nuget/archive/project-json)).

In order to build the dependency tree, Snyk analyzes the **paket.dependencies** and **paket.lock** files.

## Git services for .NET projects

.NET projects can be imported from any of the Git services Snyk supports.

When your Projects have been imported, Snyk analyzes your Projects based on their supported manifest files and then builds the dependency tree and displays it in the Snyk Web UI, similar to the following:

<figure><img src="../../../.gitbook/assets/dotNet Project.png" alt=".NET dependency tree in Snyk Web UI"><figcaption><p>.NET dependency tree in Snyk Web UI</p></figcaption></figure>

### **Nuget**

Once you select a Project for import, Snyk builds the dependency tree based on these manifest files:

* For .NET Core, the **\*.proj** files
* For .NET Framework, the **\*.proj** file, and **packages.config**

Examples of supported project files include:

* **\*.csproj**
* \***.vbproj**
* **\*.fsproj**

A .NET project can target multiple target frameworks. Snyk creates a separate dependency tree for each target framework, displaying each as a separate Snyk Project from the interface. This makes it easier to understand why a dependency is being used and also to assess the fix strategy.

### **Paket**

No import support currently.

### **Git settings for .NET**

From the Snyk Web UI, you can configure Snyk to scan your entire Project, including the build dependencies, or skip the build dependencies.

You can also **update your language preferences.**

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Go to **Settings** and select settings for **.NET**. To scan all development dependencies, be sure that **Scan build dependencies** is checked.

## Fixing vulnerabilities for .NET

For a general understanding of how Snyk helps you fix Open Source vulnerabilities within your Projects, see [Fix your vulnerabilities](../starting-to-fix-vulnerabilities/fix-your-vulnerabilities.md).

{% hint style="info" %}
**Feature availability**\
The Fix PR feature is _only_ available across Snyk [SCM](../../../integrations/git-repository-and-ci-cd-integrations-comparisons.md) integrations.
{% endhint %}

### Fix PR supported manifest files

If you are currently managing your Project dependencies with NuGet and leveraging [`PackageReference`](https://docs.microsoft.com/en-us/nuget/consume-packages/package-references-in-project-files) or [`packages.config`](https://docs.microsoft.com/en-us/nuget/reference/packages-config) Snyk can automatically update the dependency version in your manifest file, provided there is an actual fix for it. You can then review and merge your fixes.

### Dependency analysis

In the .NET ecosystem, there are multiple levels of dependencies, some of which are obvious and some completely hidden to a developer. To correctly identify the vulnerabilities for a given .NET application, these dependencies must be resolved accurately.

Snyk resolves dependencies differently in the Snyk CLI and the Source Code Management (SCM) systems such as GitHub.

**In the CLI,** if you manage your Project dependencies using `PackageReference`, Snyk scans your `obj/project.assets.json`. if you manage your dependencies using `packages.config`, Snyk scans the `packages` directory. This approach contributes to accuracy of scan results.

{% hint style="info" %}
Runtime dependencies (provided by the runtime environment also known as "meta-packages") are resolved more accurately in the CLI if the host machine uses a similar runtime SDK to the server running the app.
{% endhint %}

**In an SCM integration**, scanning uses a different process, as the generated files mentioned above are not available. To overcome this, Snyk follows the NuGet dependency [resolution algorithm](https://docs.microsoft.com/en-us/nuget/concepts/dependency-resolution) to construct a dependency tree.

For further information on .NET automated fixes, see the [Snyk blog](https://snyk.io/blog/automated-vulnerability-fixes-dot-net-dependencies).

### Build-time versus runtime dependencies

* **Build-time dependency**: Snyk understands build-time dependency to be resolved during build time and not susceptible to change at runtime.
* **Runtime dependency**: Snyk understands runtime dependency to be resolved against the installed runtime, for example, packages coming from the .NET framework (<=4) / .NET [runtime](https://docs.microsoft.com/en-us/dotnet/core/versions/selection?WT.mc\_id=DOP-MVP-5001511&) (for Core and .NET 5+) such as [`System.Net.Http`](https://www.nuget.org/packages/System.Net.Http) . Snyk sometimes refers to runtime dependencies as meta-packages.

### Tackling vulnerabilities from runtime dependencies

There are a couple of actions you can choose to take in order to address these type of vulnerable dependencies. These vary between the SCM and the CLI.

#### **Vulnerabilities from runtime dependences in SCM**

If you believe you have found false positives because the application runs on a system that always has the latest patches from Microsoft installed, which _may_ mean the vulnerability is no longer relevant to your Project, you may choose to [ignore](../../../manage-issues/prioritizing-issues/ignore-issues.md) it.

#### **Vulnerabilities from runtime dependencies in CLI**

If you believe you have found false positives because when the application runs in production you always pull the latest/explicit patches from Microsoft, which may mean the vulnerability is no longer relevant to your Project, you may [ignore](../../../manage-issues/prioritizing-issues/ignore-issues.md) them and do the following:

* If in production your application always runs on the latest SDK patch version, you can set `TargetLatestRuntimePatch` to `true` in the Project file. Make sure to upgrade your environments (for example, dev, prod) to the latest runtime version.

```
<TargetLatestRuntimePatch>true</TargetLatestRuntimePatch>
```

* You may choose to publish a [self-contained](https://docs.microsoft.com/en-us/dotnet/core/deploying/#publish-self-contained) app that includes the runtime. Then set `RuntimeFrameworkVersion`to the specific patch version in the Project file. You may choose to [ignore](../../../snyk-cli/test-for-vulnerabilities/ignore-vulnerabilities-using-snyk-cli.md) vulnerabilities that you believe are no longer relevant.

```
<PropertyGroup>
  <RuntimeFrameworkVersion>5.0.7</RuntimeFrameworkVersion>
</PropertyGroup>
```

## Not supported in Snyk for .NET

* [`Directory.Build.props`](https://docs.microsoft.com/en-us/visualstudio/msbuild/customize-your-build?view=vs-2022#directorybuildprops-and-directorybuildtargets) and [`Directory.Build.targets`](https://docs.microsoft.com/en-us/visualstudio/msbuild/customize-your-build?view=vs-2022#directorybuildprops-and-directorybuildtargets) are not currently supported via the SCM integration. You can scan private dependencies using central package management using the Snyk CLI. You have to do `dotnet restore` and then run `snyk` with `-all-projects`, as each sub folder will contain their own `project.assets.json` file.
* `<ProjectReference>`elements are not currently supported.
* Private dependency scanning is not supported for the SCM integration. You can scan private dependencies using the Snyk CLI.
