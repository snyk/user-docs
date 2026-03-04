# .NET (C# and VB.NET)

{% hint style="info" %}
.NET is supported for Snyk Code and Snyk Open Source.
{% endhint %}

## .NET for Snyk Code

{% hint style="info" %}
For .NET with Snyk Code, C# and VB.NET are supported.&#x20;
{% endhint %}

### Supported frameworks and libraries

For .NET with Snyk Code, the following frameworks and libraries are supported:

* .NET 6.0 - 9.0
* .NET Core
* .NET Framework 4.6-4.8.x
* Anthropic.SDK
* ASP.NET 6.x (C# only)
* Azure.AI.OpenAI
* Dapper
* fastJSON
* Google\_GenerativeAI
* grpc-dotnet (C# only)
* Microsoft.CodeAnalysis.VisualBasic
* Mistral.SDK
* System.CodeDom.Compiler
* Windows Forms

### Supported file formats

* For C#, Snyk supports the `.aspx` & `.cs` file formats.
* For VB.NET, Snyk supports the `.vb` file format.

### Available features

* Reports
* Custom rules
* Interfile analysis

## .NET for Snyk Open Source

{% hint style="info" %}
For .NET with Snyk Open Source, Snyk supports C# only.
{% endhint %}

### Supported versions

For .NET with Snyk Open Source, the following versions are supported:

* .NET versions 6, 7 - only the latest SDK version
* .NET versions 8, 9, and 10 - all SDK versions
* .NET Framework version 4 - using SDK-style projects
* All versions of .NET Standard

### Supported package managers and package registries

For .NET with Snyk Open Source, NuGet and Paket are supported as package managers and [nuget.org](https://www.nuget.org/) as a package registry.

### Supported file formats

* For NuGet: `project.assets.json`, `*.sln`, `packages.config,` `project.json`
* For Paket: `paket.dependencies` and `paket.lock`&#x20;

### Available integrations

* SCM import (not available for Paket)
* CLI and IDE: test or monitor your app

### Available features

For .NET with Snyk Open Source, the following features are available:

* Fix PRs (only for NuGet)
* License scanning
* Reports
* (Only for Paket) Test your app's SBOM and packages using `pkg:nuget` PURLs, using the [SBOM test](../../../developer-tools/snyk-cli/commands/sbom-test.md) command.

{% hint style="warning" %}
Snyk does not support `PackageReference` entries without a version attribute. If your Project is missing this attribute, Snyk cannot open a pull request. Ensure you add versions to all `PackageReference` entries.
{% endhint %}

### Dependency analysis using the CLI

{% hint style="info" %}
Snyk does not support `package-lock.json`. Snyk uses the build system to determine which dependencies are installed.
{% endhint %}

For dependency analysis using the CLI, Snyk must accurately resolve dependencies to identify vulnerabilities in .NET applications. Snyk scans and fixes build and development dependencies in your `*.proj`, `packages.config`, and `project.json` files.

If you manage Project dependencies using `PackageReference`, Snyk scans the `obj/project.assets.json` file.

Snyk resolves runtime dependencies (or meta-packages) more accurately when the host machine uses a runtime SDK similar to the one used by the server running the application.

### Dependency analysis using SCM integrations

In the .NET ecosystem, there are multiple levels of dependencies, some of which are obvious and some are completely hidden to a developer. To correctly identify the vulnerabilities for a given .NET application, these dependencies must be resolved accurately.

Snyk excludes developer dependencies by default because they are rarely deployed to production. This reduces noise in your scan results.

To include developer dependencies in NuGet SCM imports, navigate to **Settings** > **Languages** > **.NET**.

Snyk scans Projects differently depending on the format:

* Modern Projects (SDK-style): For modern `.csproj` formats, Snyk uses the .NET SDK to resolve dependencies. This improves accuracy and supports any Project the `dotnet` SDK can restore, including `Directory.Build.props`, `global.json`, and Central Package Management (`Directory.Packages.props`).
* Legacy Projects (non-SDK style): For Projects that use `packages.config`, `project.json`, or XML-style `*.csproj`, `*.vbproj`, or `*.fproj` files, Snyk uses static analysis. This approximates the dependency graph based on standard NuGet algorithms.

When scanning through an SCM integration, the methodology varies based on your Project format:

* Modern Projects (SDK-style): For modern `.csproj` formats, the scanner utilises the .NET SDK directly to resolve dependencies, resulting in higher accuracy. This also provides the capability of scanning any Project that can be successfully restored by the `dotnet` SDK itself including `Directory.Build.props` files, `global.json`, or Central Package Management (`Directory.Packages.props`).
* Legacy Projects (Non-SDK style): For Projects relying on `packages.config`, `project.json`, or XML-style `*.csproj` , `*.vbproj` or `*.fproj` files. The scanner uses static analysis to approximate the dependency graph based on standard NuGet algorithms.

To improve accuracy, migrate to the [Universal Broker.](../../../enterprise-setup/snyk-broker/universal-broker/)

### .NET framework package pruning&#x20;

For Snyk Projects targeting .NET 10.0 and later, the Snyk SCM and CLI scanners align with default Microsoft build behavior.

By default, Microsoft enables the `RestoreEnablePackagePruning` property for these frameworks. This setting prunes framework-provided package references during the restore operation.

Snyk respects your Project configuration. If you set `RestoreEnablePackagePruning` to `false` in your Project file or `Directory.Build.props`, Snyk does not prune these references during the scan.

For Projects targeting .NET 10.0 and later, Snyk SCM and CLI scanners align with the default Microsoft build behaviour.

By default, Microsoft enables the `RestoreEnablePackagePruning` property for these target frameworks. This setting prunes framework-provided package references during the restore operation.

The Improved .NET scanner respects your Project configuration. If you disable this feature by setting the `RestoreEnablePackagePruning` property to `false` either in your Project file or in a `Directory.Build.props` file, Snyk respects this setting and does not prune those references during the scan.
