# SCM integrations for .NET

{% hint style="info" %}
Snyk does not support Projects using Visual Studio Build Tools.

Snyk supports Windows-specific frameworks (WPF, WCF) only for .NET 10 and later.
{% endhint %}

Snyk supports NuGet as the only package manager for SCM integrations. For brokered SCMs and brokered private package repositories, Snyk supports only the Universal Broker.

Considerations:

* Snyk operates on a case-sensitive file system. Manifest definitions, such as `<ProjectReference>` strings, must match the case of the files and folders.
* `Directory.*.props`, `global.json`, and other .NET-specific manifest files must use the exact casing Microsoft describes.
* Snyk supports a subset of the major, minor, and patch versions supported by Microsoft.

Snyk uses the following files to create dependency trees:

* .NET Core: `*.proj` files
* .NET Framework: `*.proj` files and `packages.config`

## Import and scan .NET Projects

You can import .NET Projects from any supported SCM integration. Snyk analyzes your Projects using supported manifest files, builds a dependency tree, and displays the results in the Snyk Web UI.

To configure Snyk to scan build and development dependencies or skip them, navigate to the relevant Group and Organization **Settings** > **.NET**. To scan all development dependencies, ensure that the **Scan build dependencies** option is checked.

## Support manifest files

Snyk builds the dependency tree using the following manifest files:

* `*.csproj`
* `*.vbproj`
* `*.fsproj`&#x20;
* `global.json`
* `Directory.Build.props`&#x20;
* `Directory.Packages.props`&#x20;

{% hint style="info" %}
A .NET Project can target different frameworks. Snyk creates a separate list of dependencies for each framework and displays each as a distinct Snyk Project. This isolates dependencies and simplifies remediation planning.
{% endhint %}

## Fix vulnerabilities in .NET Projects

If you manage Project dependencies with NuGet using `PackageReference` or `packages.config`, Snyk automatically updates the dependency version in your manifest file if a fix is available. You can then review and merge the fixes.
