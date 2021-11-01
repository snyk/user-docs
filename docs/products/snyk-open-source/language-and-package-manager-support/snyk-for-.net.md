# Snyk for .NET

Snyk offers security scanning to test your projects for vulnerabilities, both through your CLI and through different integrations from our UI.

The following describes how to use Snyk to scan your .NET projects:

{% hint style="info" %}
**Note**\
Features might not be available, depending on your subscription plan.
{% endhint %}

| Package managers / Features                            | CLI support | Git support | License scanning | Fixing | Runtime monitoring |
| ------------------------------------------------------ | ----------- | ----------- | ---------------- | ------ | ------------------ |
| [NuGet](https://www.nuget.org)                         | ✔︎          | ✔︎          | ✔︎               | ✔︎     |                    |
| [Paket](https://fsprojects.github.io/Paket/index.html) | ✔︎          | ✔︎          |                  | ✔︎     |                    |

## **How it works**

Once we’ve built the tree, we can use our [vulnerability database ](https://snyk.io/vuln)to find vulnerabilities in any of the packages anywhere in the dependency tree.

{% hint style="info" %}
**Note**\
In order to scan your dependencies, you must ensure you have first installed the relevant package manager, and that your project contains the supported manifest files
{% endhint %}

The way by which Snyk analyzes and builds the tree varies depending on the language and package manager of the project, as well as the location of your project:

* [Snyk CLI tool for .NET projects](snyk-for-.net.md#snyk-cli-tool-for-net-projects)
* [Git services for .NET projects](snyk-for-.net.md#git-services-for-net-projects)

## Snyk CLI tool for .NET projects

### Nuget

#### Dependencies managed by PackageReference

First, restore dependencies in the .NET project by running `dotnet restore` and make sure **obj/project.assets.json** has been created by the previous command, run `snyk test`. For more information on building projects, check out [Getting started with the CLI](https://docs.snyk.io/snyk-cli/guides-for-our-cli/getting-started-with-the-cli).

Examples of supported project files that resolve into **project.assets.json** include:

* \*.csproj&#x20;
* \*.vbproj
* \*.fsproj

**Note:** Project files can be combined with [lock files](https://docs.microsoft.com/en-us/nuget/consume-packages/package-references-in-project-files#locking-dependencies) for a more deterministic **project.assets.json** resolution

#### Dependencies managed by packages.config

Whilst there are two approaches for dependencies managed by **packages.config**, the following is the recommended approach as this will yield the most accurate results:

First, install the dependencies into the **packages** folder by running `nuget install -OutputDirectory packages` and make sure the **packages** dir has been created by the previous command, run `snyk test`.

Examples of supported project files that resolve into **packages** include:

* packages.config

**Note:** While you should also be able to run `snyk test` without previously installing dependencies this will result in less accurate vulnerability results

#### **CLI parameters**

This section describes the unique CLI options available when working with NuGet managed projects.

| Option                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--file=.sln`            | Test all .NET projects included in the given `.sln` file. For example `snyk test --file=myApp.sln`                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `--file=packages.config` | Test an individual .NET project.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `--packages-folder`      | <p>This is the folder in which your dependencies are installed, provided you are using <code>packages.config</code>. If you’ve assigned a unique name to this folder, then Snyk can only find it if you enter a custom path.</p><p>Use the absolute or relative path, including the name of the folder where your dependencies reside. </p><p></p><p>For example: <code>snyk test --packages-folder=../location/to/packages</code> for Unix OS <code>snyk test --packages-folder=..\location\to\packages</code> for Windows.</p> |
| `--assets-project-name`  | When monitoring a .NET project using NuGet, the `PackageReference` key uses the project name that is indicated in the project.assets.json.                                                                                                                                                                                                                                                                                                                                                                                       |

### Paket

#### Dependencies managed by Paket

To use Paket a **paket.lock** file is required in combination with a **paket.dependencies** file. Run `snyk test`

Other support includes: **project.json** (no longer recommended, please refer to [Microsoft documentation](https://docs.microsoft.com/en-us/nuget/archive/project-json))

In order to build the dependency tree, Snyk analyzes the **paket.dependencies** and **paket.lock** files.

## Git services for .NET projects

.NET projects can be imported from any of the Git services we support.

Once imported, Snyk analyzes your projects based on their supported manifest files and then builds the dependency tree and displays it from our app, similar to the following:

![](../../../.gitbook/assets/uuid-c995621c-85c8-c79f-accd-f014e2293921-en.png)

### **Nuget**

Once you select a project for import, we build the dependency tree based on these manifest files:

* For .NET Core, the **\*.proj** files&#x20;
* For .NET Framework, the **\*.proj** file, and **packages.config**&#x20;

Examples of supported project files include:

* \*.csproj&#x20;
* \*.vbproj
* \*.fsproj

A .NET project can target multiple target frameworks. Snyk creates a separate dependency tree for each target framework, displaying each as a separate Snyk project from the interface. In this way, it’s easier to understand why a dependency is being used and also to assess the fix strategy.

### **Paket**

No import support currently.

### **Git settings for .NET**

From the Snyk UI, you can configure whether Snyk should scan your entire project, including the build dependencies, or if the build dependencies should be skipped.

****

**Update language preferences**

1. Log in to your account and navigate to the relevant group and organization that you want to manage.
2. Go to settings ![](../../../.gitbook/assets/cog\_icon.png) > and click for .NET Scan build dependencies - \_\*\*\_If checked, Snyk scans all development dependencies.

## Fixing vulnerabilities



actionable fix advice

runtime vs buildtime&#x20;

* CLI \*\*
* SCM \*\* X&#x20;

fix prs &#x20;

* \-supported manifest files
* SCM support only



&#x20;

