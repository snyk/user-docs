---
nav_context: classic
description: >-
  Prerequisites and steps to install and configure Snyk Broker for Nexus
  Repository Manager, on Enterprise plans
---

{% include "../../../../../../.gitbook/includes/new-navigation-banner.md" %}

# Nexus Repository - prerequisites and steps to install and configure Broker

{% hint style="info" %}
**Feature availability**

Integration with Nexus Repository Manager is available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Before installing, review the general instructions for the installation method you plan to use, [Helm](../../../../../../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-helm.md) or [Docker](../../../../../../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-docker.md).

Before installing the Snyk Nexus Repository Broker, ask your Snyk account team for a Broker token, or generate one by following the steps in [Obtain Broker token for Nexus integration](./#obtain-broker-token-for-nexus-integration).

You must have Docker or another way to run Linux containers.\
Some Docker deployments for Windows run only Windows containers. Ensure that your deployment can run Linux containers.

If you are installing Classic Broker, use the [Docker](nexus-repository-install-and-configure-using-docker.md) or [Helm](nexus-repository-install-and-configure-using-helm.md) instructions to obtain a token.\
If you are installing Universal Broker, skip the token steps and follow the Universal Broker setup instructions instead.&#x20;

## Obtain Broker token for Nexus integration

The environment variable configuration in this section also applies to Universal Broker. When you configure Nexus using a Broker (Classic or Universal), you manage all configuration through Broker environment variables. The Snyk Web UI does not display Nexus settings under **Languages** or **Brokered package registries**.

If you are installing Universal Broker, enter your Nexus Base URL in this format when prompted: `https://username:password@your.nexus.hostname.`

1. Navigate to **Settings** > **Integrations** > **Package Repositories** > **Nexus**.
2. Ensure that the Nexus configuration screen is visible. If the Snyk Broker switch is missing, you do not have the necessary permissions and can only add a publicly accessible instance. Contact [Snyk Support](https://support.snyk.io) to request access to add a private registry.
3. Move the Snyk Broker switch to **On** to display the Broker token generation form.
4. Select **Generate and Save**.
5. Copy the generated token to use when you set up the Broker Client.
