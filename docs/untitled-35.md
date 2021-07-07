# How does Snyk aggregate .NET Projects?

##  How does Snyk aggregate .NET Projects?

For an updated doc on .NET support, please see here: [https://snyk.io/docs/snyk-for-dotnet/](https://snyk.io/docs/snyk-for-dotnet/)

Snyk supports testing and monitoring [.NET](https://docs.microsoft.com/en-us/dotnet/) projects that have their dependencies managed by [NuGet](https://www.nuget.org/) or [Paket](https://fsprojects.github.io/Paket/index.html).

### **Source Control Management\*** <a id="source-control-management*"></a>

Snyk's source control integrations support testing and monitoring of .NET Core and .NET Framework projects which use [Nuget](https://www.nuget.org/) as their package manager.

  
In order to import a .NET project to Snyk, we require that one of the following files will be part of the imported repository:

*  packages.config \(.NET Framework\)
*  .proj file \(.NET Core\)
*  project.json \(.NET Core\)

Snyk will analyze each manifest in the context of a target framework in order to understand all the dependencies brought in by the projects. Snyk uses Nuget API \(3+\) to grab all transitive dependencies.

![](https://res.cloudinary.com/snyk/image/upload/v1547730543/docs/dotNet_scm.png)

Once the project is imported, we create a dependency tree for each target framework and compares the specific versions of every [direct and deep dependency](https://support.snyk.io/hc/en-us/articles/360000905138-What-are-direct-and-deep-dependencies-) in your project against our [NuGet vulnerability database](https://snyk.io/vuln?type=nuget). In case we find a vulnerability, we will notify you so you will be able to decide which action to take.

### **CLI** <a id="cli"></a>

Snyk's CLI supports testing and monitoring .NET Core and .NET Framework projects who use [NuGet](https://www.nuget.org/) or [Paket](https://fsprojects.github.io/Paket/index.html) as their package manager.  
In order to test and monitor .NET projects we require one of the following:

* packages.config \(.NET Framework\)
*  .proj file or project.json or project.assets.json \(.NET Core\)
*  paket.dependencies and paket.lock

Once the project is imported, we create a dependency tree for each project and compares the specific versions of every [direct and deep dependency](https://support.snyk.io/frequently-asked-questions/snyk-basics/what-are-direct-and-deep-dependencies) in your project against our [NuGet vulnerability database](https://snyk.io/vuln?type=nuget). In case we find a vulnerability, we will notify you so you will be able to decide which action to take.

