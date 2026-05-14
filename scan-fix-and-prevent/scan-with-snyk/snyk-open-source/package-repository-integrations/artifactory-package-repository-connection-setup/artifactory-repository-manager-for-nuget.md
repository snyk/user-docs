# Artifactory repository manager for NuGet

{% hint style="info" %}
**Feature availability**\
Package repository integrations are available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

If you use a brokered Artifactory connection, you must use the Universal Broker.

This configuration applies only to Snyk Web UI integrations. The Snyk CLI supports private NuGet repositories automatically.

&#x20;Snyk integrates with Artifactory Package Repositories for NuGet packages. This allows Snyk to resolve transitive dependencies and regenerate lockfiles with the correct URLs during pull requests.

Configure the URL where your private Artifactory Package Repositories host NuGet packages. This matches the information in your `nuget.config` file.

## Configure .NET language settings

To do this:

1. In the Snyk web UI, navigate to **Settings** > **Languages** > **.NET** > **Brokered package registries**.&#x20;

If you use a brokered connection and have not configured the Universal Broker, follow the Universal Broker instructions to set it up.

2. Select **Artifactory** for the **Registry type**.
3. Enter the **Registry URL** for the NuGet feed. The URL must be complete and refer to the `index.json` file for the feed. Specifically, it must end with `/v3/{repository-name}/index.json`. For example: `https://artifactory.example.com/artifactory/api/nuget/v3/nuget-local/index.json`
