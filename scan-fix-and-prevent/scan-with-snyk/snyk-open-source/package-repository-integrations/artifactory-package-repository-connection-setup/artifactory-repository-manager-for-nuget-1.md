---
description: How to connect Snyk to an Artifactory repository manager for Go
nav_context: classic
---

# Artifactory repository manager for Go

{% hint style="info" %}
**Feature availability**

Package repository integrations are available only with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk integrates with JFrog Artifactory so it can resolve your private Go modules and transitive dependencies during Open Source (SCA) scans.

You configure an Artifactory Go module proxy from a single page (**Settings > Languages > Go**) using one of two boxes, depending on how your Artifactory instance is reachable. The configuration steps are identical for both:

* **Private package registries** for an Artifactory instance that is publicly reachable and protected by credentials. Snyk connects to it directly. **Prerequisite:** configure your [Artifactory integration](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup) first.
* **Brokered package registries** for an Artifactory instance that is not publicly reachable. **Prerequisite:** configure the [Snyk Universal Broker](https://docs.snyk.io/platform-administration/snyk-broker/universal-broker) first and then [enable it for Open Source](https://docs.snyk.io/platform-administration/snyk-broker/universal-broker/basic-steps-to-install-and-configure-universal-broker#integrate-your-connection-with-an-organization-that-uses-universal-broker).

{% hint style="info" %}
This configuration applies to Snyk Web UI (SCM) integrations only. The Snyk CLI picks up your private Go proxy automatically from your `GOPROXY` environment and `.netrc` credentials.
{% endhint %}

### Configure an Artifactory Go registry

{% hint style="warning" %}
If a private or brokered package registry integration is not configured first, the steps below will not work — the registry settings field will not be visible on the page.
{% endhint %}

First, set up the Artifactory integration: go to **Organization Settings > Add Integration > Artifactory > Edit Settings**.

Then, on the **Settings > Languages > Go** page:

1. In the relevant box, **Private package registries** or **Brokered package registries,** select **Artifactory** as the registry type.
2. In the **Registry URL** field, enter the full URL of your Artifactory Go remote or virtual repository, using the Artifactory Go API path.
3. Save the integration.

Snyk authenticates to the registry using the credentials stored in your Artifactory integration and you should **not** enter credentials here.

#### Registry URL format

The URL must point at the Go API endpoint for the repository. For an Artifactory Go repository, the URL has the form:

```
https://<artifactory-host>/artifactory/api/go/<go-repo-name>
```

**Example:**

```
https://artifactory.example.com/artifactory/api/go/go-virtual
```

{% hint style="warning" %}
Make sure no credentials are present in the URL — provide only the path to the Go repository. Artifactory's **Set Me Up** dialog returns URLs with `username:password@` embedded in the host; you must remove that portion before pasting the URL. Snyk retrieves credentials from your Artifactory integration, not from this field.
{% endhint %}

This is the same base URL you set as `GOPROXY` when fetching modules locally, for example:

```
GOPROXY=https://artifactory.example.com/artifactory/api/go/go-virtual
```
