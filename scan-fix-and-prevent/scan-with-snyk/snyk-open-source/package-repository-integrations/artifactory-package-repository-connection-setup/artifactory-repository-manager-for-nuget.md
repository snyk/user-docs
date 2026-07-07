# Artifactory repository manager for NuGet

{% hint style="info" %}
**Feature availability**\
Package repository integrations are available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

If you use a brokered Artifactory connection, you must use the Universal Broker.

This configuration applies only to Snyk Web UI integrations. The Snyk CLI supports private NuGet repositories automatically.

Snyk integrates with Artifactory Package Repositories for NuGet packages. This allows Snyk to resolve transitive dependencies and regenerate lockfiles with the correct URLs during pull requests.

Configure the URL where your private Artifactory Package Repositories host NuGet packages. This matches the information in your `nuget.config` file.

## Configure .NET language settings

If you use a brokered connection and have not configured the Universal Broker, [configure](https://docs.snyk.io/platform-administration/snyk-broker/universal-broker) it first and then [enable it for Open Source](https://docs.snyk.io/platform-administration/snyk-broker/universal-broker/basic-steps-to-install-and-configure-universal-broker#integrate-your-connection-with-an-organization-that-uses-universal-broker).

To do this:

1. In the Snyk web UI, navigate to **Settings** > **Languages** > **.NET** > **Brokered package registries**.
2. Select **Artifactory** for the **Registry type**.
3. Enter the **Registry URL** for the NuGet feed. The URL must be complete and refer to the `index.json` file for the feed. Specifically, it must end with `/v3/{repository-name}/index.json`. For example: `https://artifactory.example.com/artifactory/api/nuget/v3/nuget-local/index.json`
