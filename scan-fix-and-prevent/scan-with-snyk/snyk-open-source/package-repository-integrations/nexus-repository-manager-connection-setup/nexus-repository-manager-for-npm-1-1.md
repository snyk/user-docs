---
description: How to connect Snyk to a Nexus Repository Manager for Go
---

# Nexus repository manager for Go



{% hint style="info" %}
**Feature availability**

Package repository integrations are available only with Enterprise plans. Visit [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk integrates with Nexus Repository Manager to resolve your private Go modules and transitive dependencies during Open Source (SCA) scans.

You can configure a Nexus Go module proxy from the Snyk web UI, by navigating to  **Settings** > **Languages** > **Go**.  Use one of two options, depending on how your Nexus instance is reachable. The configuration steps are identical for both options:

* **Private package registries**: Use this option for a publicly reachable Nexus instance protected by credentials. Snyk connects directly. As a prerequisite, ensure to configure your [Nexus integration](./) first.
* **Brokered package registries**: Use this option for a Nexus instance that is not publicly reachable. As a prerequisite, ensure to configure the [Snyk Universal Broker](https://docs.snyk.io/platform-administration/snyk-broker/universal-broker) first and then enable it for Open Source. For details, visit [Integrate your connection with an Organization that uses Universal Broker](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/universal-broker/basic-steps-to-install-and-configure-universal-broker).

{% hint style="info" %}
This configuration applies only to Snyk web UI (SCM) integrations. The Snyk CLI automatically uses your private Go proxy from your `GOPROXY` environment and `.netrc` credentials.
{% endhint %}

### Configure a Nexus Go registry

Ensure you configure a private or brokered package registry integration first. Without it, the following steps do not work, and the registry settings field is not visible on the page.

First, set up the Nexus integration by navigating to **Organization Settings > Add Integration > Nexus > Edit Settings**.

Then, on the **Settings** > **Languages** > **Go** page, complete the following steps:

1. In the **Private package registries** or **Brokered package registries** section, select **Nexus** as the registry type.
2. n the **Registry URL** field, enter the full URL of your Nexus Go (`golang`) proxy repository.
3. Save the integration.

{% hint style="info" %}
Snyk authenticates to the registry with credentials from your Nexus integration. Do not enter credentials here.
{% endhint %}

#### Registry URL format

The URL must point to the Go proxy repository in Nexus. For a Nexus `golang` proxy repository, the URL format is:

```
https://<nexus-host>/repository/<go-repo-name>/
```

Example:

```
https://nexus.example.com/repository/go-proxy/
```

{% hint style="info" %}
Ensure the URL contains no credentials. Provide only the path to the Go repository. Some Nexus copy-URL helpers embed `username:password@` in the host. Remove that portion before you paste the URL. Snyk retrieves credentials from your Nexus integration, not from this field.
{% endhint %}

This is the same base URL you set as `GOPROXY` when fetching modules locally. For example:

```
GOPROXY=https://nexus.example.com/repository/go-proxy/
```
