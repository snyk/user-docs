---
nav_context: agnostic
description: >-
  How to test .NET Projects with the Snyk CLI, including open source analysis
  with solution and project files and source code scanning
last_reviewed: 2026-09-16
last_reviewed_source: bulk-sync
---

# CLI support for .NET

To analyze Open Source libraries, install your dependencies, then run `snyk test` using one of the following options:

* `--file=`: Targets a specific solution file (`.sln`) or Project file.
* `--all-projects`: Analyzes all Open Source Projects. Use this for Projects with multiple languages, package managers, or `.sln` files.

To perform source code analysis, run `snyk code test` from the root of the Project.

## Dependency analysis

Snyk resolves .NET dependencies differently in the Snyk CLI and in SCM integrations, such as GitHub and GitLab. As a result, a CLI scan and an SCM import of the same Project can report different dependencies and different vulnerabilities.

### Dependency analysis in the CLI

In the CLI, Snyk resolves dependencies in your local build environment, using either the .NET SDK or the files that the restore step generates:

* For SDK-style Projects, such as .NET Core, .NET Standard, and .NET 6 and later, Snyk uses runtime resolution by default when the .NET SDK (`dotnet`) is on your PATH. Runtime resolution uses the .NET SDK to resolve the dependency graph for each target framework.
* For Projects that manage dependencies with `PackageReference`, Snyk scans the `obj/project.assets.json` file when runtime resolution is turned off with `--dotnet-runtime-resolution=false`, or when `dotnet` is not on your PATH.
* For Projects that manage dependencies with `packages.config`, Snyk scans the `packages` directory.

When Snyk scans `obj/project.assets.json` or the `packages` directory, restore your dependencies before you run `snyk test`. For the commands, visit [NuGet](#nuget). For details about runtime resolution, visit [`--dotnet-runtime-resolution`](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/test#dotnet-runtime-resolution) in the Test help.

{% hint style="info" %}
Snyk resolves runtime dependencies, also known as meta-packages, such as `System.Net.Http`, more accurately in the CLI when the host machine uses a runtime SDK similar to the one on the server that runs the application.
{% endhint %}

### Dependency analysis in SCM integrations

In SCM integrations, the files that the restore step generates, such as `obj/project.assets.json` and the `packages` directory, are not available. Snyk builds the dependency tree from your manifest files instead, and the method depends on the Project format:

* SDK-style Projects, which manage dependencies with `PackageReference`: Snyk resolves dependencies directly with the .NET SDK.
* Legacy (non-SDK-style) Projects, such as Projects that use `packages.config`: Snyk uses static analysis to approximate the dependency graph, based on the [NuGet dependency resolution algorithm](https://learn.microsoft.com/en-us/nuget/concepts/dependency-resolution).

Because of these differences, the results of an SCM import can differ from the results of a CLI scan.

For more details about how Snyk builds the dependency tree in SCM integrations, visit [SCM integrations for .NET](scm-integrations-for-.net.md) and [Dependency analysis using SCM integrations](README.md#dependency-analysis-using-scm-integrations).

## NuGet

For NuGet-specific options, visit [Options for NuGet projects in the Test help](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/test#options-for-nuget-projects) and [Options for NuGet projects in the Monitor help](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/commands/monitor#options-for-nuget-projects).

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

### Support for project.json

Snyk supports `project.json` files, although Microsoft no longer recommends this format.

### Support for packages.config

Snyk supports `packages.config` files. To scan these Projects:

1. Run `nuget install -OutputDirectory packages` to install dependencies into the `packages` folder.
2. Ensure the command created the `packages` directory.
3. Run `snyk test`.

## Paket

To use Paket with the Snyk CLI:

1. Ensure your project contains a `paket.lock` file and a `paket.dependencies` file.
2. Run `snyk test`.

## Options and plugins

* [snyk-to-html](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html): Generates reports locally or at build time.
* `--json` and `--sarif`: Generates output for programmatic access.
* [snyk-filter](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter): Provides advanced filtering options.
