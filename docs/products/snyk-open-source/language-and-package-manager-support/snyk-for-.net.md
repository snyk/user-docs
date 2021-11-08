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
| [Paket](https://fsprojects.github.io/Paket/index.html) | ✔︎          |             |                  |        |                    |

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
2.  Go to settings ![](../../../.gitbook/assets/cog\_icon.png) > and click for .NET&#x20;

    Scan build dependencies - If checked, Snyk scans all development dependencies.

## Fixing vulnerabilities

For a general understanding of how Snyk helps you fix Open Source vulnerabilities within your projects, please visit the following document [Fix your vulnerabilities](https://docs.snyk.io/features/fixing-and-prioritizing-issues/issue-management/remediate-your-vulnerabilities).&#x20;

{% hint style="info" %}
Please note the Fix PR feature is currently in open beta and only available across our [SCM](https://docs.snyk.io/getting-started/scm-git-and-ci-cd-integration-deployment-intro) integrations. If you would like to be part of the early beta program we recommend you get in touch with your Customer Success Manager.
{% endhint %}

### Fix PR supported manifest files

If you are currently managing your project dependencies with NuGet and leveraging [`PackageReference`](https://docs.microsoft.com/en-us/nuget/consume-packages/package-references-in-project-files) or [`packages.config`](https://docs.microsoft.com/en-us/nuget/reference/packages-config) we will be able to automatically update the dependency version in your manifest file, provided there is an actual fix for it. You should then be able to easily review and merge your fixes.&#x20;

### Dependency analysis

In the .NET ecosystem, there are multiple levels of dependencies, some of which are obvious and some completely hidden to a developer. To correctly identify the vulnerabilities for a given .NET application, these dependencies must be resolved accurately.

We resolve dependencies differently in the Snyk CLI, and the Source Code Management (SCM) systems such as GitHub:

* In the CLI, if you manage your project dependencies using `PackageReference`, we scan your `obj/project.assets.json`; if you manage your dependencies using `packages.config`, we scan the `packages` directory. This approach allows us to be very accurate.
*   In the SCM integration, scanning uses a different process, as the generated files mentioned above are not available. To overcome this we follow the nuget dependency [resolution algorithm](https://docs.microsoft.com/en-us/nuget/concepts/dependency-resolution) to construct a dependency tree.&#x20;

    **Note**: runtime dependencies (provided by the runtime environment also known as "meta-packages") are resolved more accurately in the CLI if the host machine uses a similar runtime SDK to the server running the app.

#### Build-time vs Runtime dependencies

* **Build-time dependency**: We understand build time dependency to be a dependency that is resolved during build time and is not susceptible to change at runtime.
* **Runtime dependency**: We understand build time dependency to be a dependency that is resolved during runtime. For example, packages coming from standard .NET SDK such as [`System.Net.Http`](https://www.nuget.org/packages/System.Net.Http) . We sometimes refer to runtime dependencies as meta-packages.

### Tackling vulnerabilities from runtime dependencies

There are a couple of actions you can choose to take in order to address these type of vulnerable dependencies. These vary if you are using the SCM or CLI:

**SCM**

If you believe you have found false positives because when the application runs in production you always pull the latest/explicit patches from Microsoft, which may mean the vulnerability is no longer relevant to your project, you may choose to [ignore](https://docs.snyk.io/features/fixing-and-prioritizing-issues/issue-management/ignore-issues#ignoring-issues-in-the-ui) it.&#x20;

**CLI**

If you believe you have found false positives because when the application runs in production you always pull the latest/explicit patches from Microsoft, which may mean the vulnerability is no longer relevant to your project, you may do the following:

* If in production your application always runs on the latest SDK patch version, you can set `TargetLatestRuntimePatch` to `true` in the project file.

```
<TargetLatestRuntimePatch>true</TargetLatestRuntimePatch>
```

* You may choose to publish a [self-contained](https://docs.microsoft.com/en-us/dotnet/core/deploying/#publish-self-contained) app that includes runtime. Then set `RuntimeFrameworkVersion` to the specific patch version in the project file.

```
<PropertyGroup>
  <RuntimeFrameworkVersion>5.0.7</RuntimeFrameworkVersion>
</PropertyGroup>
```

