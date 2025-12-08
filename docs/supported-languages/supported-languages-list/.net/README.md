# .NET (C# and VB.NET)

## Applicability

Snyk supports [.NET for code analysis ](.net-for-code-analysis.md)(C# & VB.NET) and [.NET for open source](.net-for-open-source.md) (C# only).&#x20;

There is [guidance for Snyk for .NET](guidance-for-snyk-for-.net.md). If you need help, see [Troubleshooting Snyk for .NET](troubleshooting-snyk-for-.net.md).

Snyk has introduced a new Early Access feature with significantly enhanced scanning capabilities for NuGet applications. For detailed information and access to these features, see the [Improved .NET scanning](improved-.net-scanning.md) page.

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.

Available functions:

* SCM import available for Snyk Open Source and Snyk Code. If you use .NET for Snyk Open Source, then the SCM import is only available when used with NuGet.
* Test or monitor your app through CLI and IDE, available for both Snyk Open Source and Snyk Code.
* Test your app's SBOM using `pkg:nuget`
* Test your app's packages using `pkg:nuget`

## Package managers and supported file extensions

Snyk for .NET supports NuGet and Paket as package managers and [nuget.org](https://www.nuget.org/) as a package registry and supports the following file formats:

* Snyk Open Source:
  * For NuGet: `project.assets.json`, `*.sln`, `packages.config,` `project.json`
  * For Paket: `paket.dependencies` and `paket.lock`&#x20;
* Snyk Code:
  * For C#, Snyk supports the `.aspx` & `.cs` file formats.
  * For VB.NET, Snyk supports the `.vb` file format.

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk Code for .NET:

* .NET 6.0 - 9.0
* .[^1]NET Core
* .NET Framework 4.6-4.8.x
* Anthropic.SDK
* ASP.NET 6.x **(C# only)**
* Azure.AI.OpenAI
* Dapper
* fastJSON
* Google\_GenerativeAI
* grpc-dotnet **(C# only)**
* Microsoft.CodeAnalysis.VisualBasic
* Mistral.SDK
* System.CodeDom.Compiler
* Windows Forms

## Features

The following features are supported in Snyk for .NET:

| Snyk Open Source                                                           | Snyk Code                                                                 |
| -------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| <ul><li>Fix PRs (NuGet)</li><li>License scanning</li><li>Reports</li></ul> | <ul><li>Reports</li><li>Custom rules</li><li>Interfile analysis</li></ul> |

[^1]: 
