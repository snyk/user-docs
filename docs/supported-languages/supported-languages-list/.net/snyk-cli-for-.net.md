# Snyk CLI for .NET

To analyze Open Source libraries, install your dependencies, then run `snyk test` using one of the following options:

* `--file=`: Targets a specific solution file (`.sln`) or Project file.
* `--all-projects`: Analyzes all Open Source Projects. Use this for Projects with multiple languages, package managers, or `.sln` files.

To perform source code analysis, run `snyk code test` from the root of the Project.

## NuGet

For NuGet-specific options, visit [Options for NuGet projects in the Test help](../../../developer-tools/snyk-cli/commands/test.md#options-for-nuget-projects) and [Options for NuGet projects in the Monitor help](../../../developer-tools/snyk-cli/commands/monitor.md#options-for-nuget-projects).

Snyk scans NuGet Projects using the `project.assets.json` file. Snyk supports the following Project files that resolve into `project.assets.json`:

* `*.csproj`
* `*.vbproj`
* `*.fsproj`

To scan a NuGet Project:

1. Run `dotnet restore`. This restores dependencies and creates the `obj/project.assets.json` file.
2. Run `snyk test`.

{% hint style="info" %}
You must run `dotnet restore` before `snyk test` to restore dependencies and ensure accurate scan results.
{% endhint %}

#### Support for project.json

Snyk supports `project.json` files, although Microsoft no longer recommends this format.

#### Support for packages.config

Snyk supports `packages.config` files. To scan these Projects:

1. Run `nuget install -OutputDirectory packages` to install dependencies into the `packages` folder.
2. Ensure the command created the `packages` directory.
3. Run `snyk test` as follows:
   1. Install the dependencies into the packages folder by running `nuget install -OutputDirectory packages`
   2. Ensure that the packages directory has been created by the previous command.
   3. Run `snyk test`.

## Paket

To use Paket with the Snyk CLI:

1. Ensure your project contains a `paket.lock` file and a `paket.dependencies` file.
2. Run `snyk test`.

## Options and plugins

* [snyk-to-html](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md): Generates reports locally or at build time.
* `--json` and `--sarif`: Generates output for programmatic access.
* [snyk-filter](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md): Provides advanced filtering options.
