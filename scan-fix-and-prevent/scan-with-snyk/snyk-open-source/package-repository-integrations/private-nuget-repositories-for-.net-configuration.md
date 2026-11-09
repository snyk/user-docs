---
description: How to configure private NuGet repositories for .NET with Snyk
nav_context: classic
last_reviewed: 2026-11-05
last_reviewed_source: reviewed
---

{% include "../../../.gitbook/includes/new-navigation-banner.md" %}

# Private NuGet repositories for .NET configuration

{% hint style="info" %}
**Feature availability**\
This feature is available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

This configuration applies only to Snyk Web UI integrations. The Snyk CLI automatically supports private NuGet repositories.

To resolve transitive dependencies, Snyk must access relevant private NuGet repositories.

Snyk recommends using [`nuget.config`](https://learn.microsoft.com/en-us/nuget/reference/nuget-config-file) files and registering the credentials in Snyk NuGet private package repository integration. To do this, navigate to Organization **Settings** > **Integrations** > **NuGet Repositories**.

{% hint style="warning" %}
Name your NuGet config file `nuget.config`, `NuGet.config`, or `NuGet.Config`. Snyk restores .NET Projects on Linux, where file names are case-sensitive and NuGet reads only these three names.

NuGet ignores any other casing, such as `Nuget.Config`, without an error. Snyk then cannot use your private package sources, and the scan fails with `NU1101` or `NU1102`. This happens even if the same file works on Windows or macOS, where file names are not case-sensitive.
{% endhint %}

{% hint style="info" %}
If you instruct the .NET ecosystem of where to look for private packages in a different way than using `nuget.config`, Snyk tries to add all private NuGet repository credentials that are defined in the private package repository integration as a `dotnet nuget` source before restoring the Project.
{% endhint %}

Fill in the fields in the **Your tokens** section: by adding a **Username**, the **Personal access token**, and the repository **URL**, which supports only HTTPS sources.
