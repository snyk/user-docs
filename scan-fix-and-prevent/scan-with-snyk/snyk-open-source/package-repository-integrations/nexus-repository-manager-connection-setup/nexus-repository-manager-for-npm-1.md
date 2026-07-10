---
description: How to connect Snyk to a Nexus Repository Manager for NuGet
---

# Nexus repository manager for NuGet

{% hint style="info" %}
**Feature availability**\
Package repository integrations are available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

If you use a brokered Nexus connection, you must use the Universal Broker.

This configuration applies only to Snyk Web UI integrations. The Snyk CLI supports private NuGet repositories automatically.

Snyk can use Nexus Repository Manager with NuGet packages. This allows Snyk to resolve transitive dependencies and regenerate lockfiles with the correct URLs during pull requests.

Configure the URL where your private Nexus Repository Manager hosts NuGet packages. This matches the information in your `nuget.config` file.

## Configure .NET language settings

If you use a brokered connection and have not configured the Universal Broker, configure it first and then enable it for Open Source.

To configure a brokered NuGet registry:

1. In the Snyk web UI, navigate to **Settings** > **Languages** > **.NET** > **Brokered package registries**.
2. Select **Nexus** for the **Registry type**.
3. Enter the **Registry URL** for the NuGet feed. The URL must be complete and refer to the `index.json` file for the feed. Specifically, it must end with `/index.json`. For example: `https://nexus.example.com/repository/nuget-hosted/index.json`
