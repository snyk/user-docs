---
description: How to connect Snyk to an Artifactory repository manager for NuGet
nav_context: classic
last_reviewed: 2026-11-05
last_reviewed_source: reviewed
---

{% include "../../../../.gitbook/includes/new-navigation-banner.md" %}

# Artifactory repository manager for NuGet

{% hint style="info" %}
**Feature availability**\
Package repository integrations are available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

If you use a brokered Artifactory connection, you must use the Universal Broker.

This configuration applies only to Snyk Web UI integrations. The Snyk CLI supports private NuGet repositories automatically.

Snyk integrates with Artifactory Package Repositories for NuGet packages. Snyk uses this integration to scan. When Snyk runs `dotnet restore`, NuGet uses the feed to resolve your private packages and their transitive dependencies. Fix pull requests update package versions in your project files and do not use the integration.

Configure the URL where your private Artifactory Package Repositories host NuGet packages. Use the same URL as the matching package source in your `nuget.config` file.

{% hint style="warning" %}
Name your NuGet config file `nuget.config`, `NuGet.config`, or `NuGet.Config`. Snyk restores .NET Projects on Linux, where file names are case-sensitive and NuGet reads only these three names.

NuGet ignores any other casing, such as `Nuget.Config`, without an error. Snyk then cannot use your private package sources, and the scan fails with `NU1101` or `NU1102`. This happens even if the same file works on Windows or macOS, where file names are not case-sensitive.
{% endhint %}

## Configure .NET language settings

If you use a brokered connection and have not configured the Universal Broker, [configure](https://docs.snyk.io/platform-administration/snyk-broker/universal-broker) it first and then [enable it for Open Source](https://docs.snyk.io/platform-administration/snyk-broker/universal-broker/basic-steps-to-install-and-configure-universal-broker#integrate-your-connection-with-an-organization-that-uses-universal-broker).

To do this:

1. In the Snyk web UI, navigate to **Settings** > **Languages** > **.NET** > **Brokered package registries**.
2. Select **Artifactory** for the **Registry type**.
3. Enter the **Registry URL** for the NuGet feed. The URL must be complete and refer to the `index.json` file for the feed. Specifically, it must end with `/v3/{repository-name}/index.json`. For example: `https://artifactory.example.com/artifactory/api/nuget/v3/nuget-local/index.json`
