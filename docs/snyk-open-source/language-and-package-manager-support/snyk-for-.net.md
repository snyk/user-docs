# Snyk for .NET

Snyk offers security scanning to test your projects for vulnerabilities, both through your CLI and through different integrations from our UI.

The following describes how to use Snyk to scan your .NET projects:

## Note

Features might not be available, depending on your subscription plan.

|  | Package managers/Features | CLI support | Git support | License scanning | Remediation | Runtime monitoring |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| ![i\_icon\_nuget.png](https://support.snyk.io/hc/article_attachments/360007259058/uuid-b997ca27-61ff-f00b-941c-16bf3aa4a0e0-en.png) | [NuGet](https://www.nuget.org/) | ✔︎ | ✔︎ | ✔︎ |  |  |
| ![i\_icon\_paket.png](https://support.snyk.io/hc/article_attachments/360007259078/uuid-d8e44fe4-c0ea-e3ea-de3b-1e15e4a6b391-en.png) | [Paket](https://fsprojects.github.io/Paket/index.html) | ✔︎ |  | ✔︎ |  |  |

### Snyk CLI tool for .NET projects <a id="h_01ED93MRAVCCCSMJTFJM4595T5"></a>

### Dependencies managed by PackageReference

First, restore dependencies in the .NET project by running `dotnet restore` and make sure **obj/project.assets.json** has been created by the previous command, run `snyk test`. For more information on building projects, check out [Getting started with the CLI](https://support.snyk.io/hc/en-us/articles/360003812458).

Examples of supported project files that resolve into **project.assets.json** include:

* \*.csproj 
* \*.vbproj
* \*.fsproj

**Note:** Project files can be combined with [lock files](https://docs.microsoft.com/en-us/nuget/consume-packages/package-references-in-project-files#locking-dependencies) for a more deterministic **project.assets.json** resolution

### Dependencies managed by packages.config

Whilst there are two approaches for dependencies managed by **packages.config**, the following is the recommended approach as this will yield the most accurate results:

First, install the dependencies into the **packages** folder by running `nuget install -OutputDirectory packages` and make sure the **packages** dir has been created by the previous command, run `snyk test`.

Examples of supported project files that resolve into **packages** include:

* packages.config

  **Note**: whilst you should also be able to run `snyk test` without previously installing dependencies this will result in less accurate vulnerability results

### Dependencies managed by Paket

To use Paket a **paket.lock** file is required in combination with a **paket.dependencies** file. Run `snyk test`

### Other support includes: **project.json** \(no longer recommended, please refer to [Microsoft documentation](https://docs.microsoft.com/en-us/nuget/archive/project-json)\)

#### Nuget

Follow the same instructions as [Snyk CLI tool for .NET projects](snyk-for-.net.md)

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
      <td style="text-align:left"><code>--file=.sln</code>
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

### Git services for .NET projects

.NET projects can be imported from any of the Git services we support.

Once imported, Snyk analyzes your projects based on their supported manifest files and then builds the dependency tree and displays it from our app, similar to the following:

![dependency\_tree.png](https://support.snyk.io/hc/article_attachments/360007259098/uuid-c995621c-85c8-c79f-accd-f014e2293921-en.png)

No import support currently.

From the Snyk UI, you can configure whether Snyk should scan your entire project, including the build dependencies, or if the build dependencies should be skipped.

**Update language preferences**

1. 1. Go to settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; and click for .NET
      * If checked, Snyk scans all development dependencies.

