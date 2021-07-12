# Snyk for .NET

* [ Language support summary](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360020352437-Language-support-summary/README.md)
* [ Snyk for JavaScript](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004712477-Snyk-for-JavaScript/README.md)
* [ Snyk for Java \(Gradle, Maven\)](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003817357-Snyk-for-Java-Gradle-Maven-/README.md)
* [ Snyk for .NET](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004519138-Snyk-for-NET/README.md)
* [ Snyk for Python](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004699377-Snyk-for-Python/README.md)
* [ Snyk for Golang](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003817417-Snyk-for-Golang/README.md)
* [ Snyk for Swift and Objective-C \(CocoaPods\)](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004701658-Snyk-for-Swift-and-Objective-C-CocoaPods-/README.md)
* [ Snyk for Scala](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003781318-Snyk-for-Scala/README.md)
* [ Snyk for Ruby](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003781298-Snyk-for-Ruby/README.md)
* [ Snyk for PHP](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003817397-Snyk-for-PHP/README.md)

  [See more](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/sections/360001087857-Language-package-manager-support/README.md)

## Snyk for .NET

Snyk offers security scanning to test your projects for vulnerabilities, both through your CLI and through different integrations from our UI.

The following describes how to use Snyk to scan your .NET projects:

### Note

Features might not be available, depending on your subscription plan.

|  | Package managers/Features | CLI support | Git support | License scanning | Remediation | Runtime monitoring |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
|  | [NuGet](https://www.nuget.org/) | ✔︎ | ✔︎ | ✔︎ |  |  |
|  | [Paket](https://fsprojects.github.io/Paket/index.html) | ✔︎ |  | ✔︎ |  |  |

#### Snyk CLI tool for .NET projects <a id="h_01ED93MRAVCCCSMJTFJM4595T5"></a>

First, build the .NET project by running `dotnet restore` prior to running `snyk test`. For more information on building projects, check out [Getting started with the CLI](https://support.snyk.io/hc/en-us/articles/360003812458).

The following manifest files are supported:

**Nuget**

Follow the same instructions as [Snyk CLI tool for .NET projects](untitled-153.md)

In order to build the dependency tree, Snyk analyzes the **paket.dependencies** and **paket.lock** files.

This section describes the unique CLI options available when working with .NET-based projects.

<table>
  <thead>
    <tr>
      <th style="text-align:left">Option</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><code>--assets-project-name</code>
      </td>
      <td style="text-align:left">When monitoring a .NET project using NuGet, the <code>PackageReference</code> key
        uses the project name that is indicated in the project.assets.json.</td>
    </tr>
    <tr>
      <td style="text-align:left"><code>--packages-folder</code>
      </td>
      <td style="text-align:left">
        <p>This is the folder in which your dependencies are installed. If you&#x2019;ve
          assigned a unique name to this folder, then Snyk can only find it if you
          enter a custom path.</p>
        <p>Use the absolute or relative path, including the name of the folder where
          your dependencies reside.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><code>--file=&lt;your_solution&gt;.sln</code>
      </td>
      <td style="text-align:left">Test all .NET projects included in the given <code>.sln</code> file</td>
    </tr>
    <tr>
      <td style="text-align:left"><code>--file=packages.config</code>
      </td>
      <td style="text-align:left">Test an individual .NET project.</td>
    </tr>
  </tbody>
</table>

#### Git services for .NET projects

.NET projects can be imported from any of the Git services we support.

Once imported, Snyk analyzes your projects based on their supported manifest files and then builds the dependency tree and displays it from our app, similar to the following:

No import support currently.

From the Snyk UI, you can configure whether Snyk should scan your entire project, including the build dependencies, or if the build dependencies should be skipped.

**Update language preferences**

1. 1. Go to settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; and click for .NET
      * If checked, Snyk scans all development dependencies.

