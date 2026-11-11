---
description: >-
  How to distribute Snyk Studio as a managed utility across your development
  teams
nav_context: classic
last_reviewed: 2026-09-16
last_reviewed_source: bulk-sync
---

# Distribution at scale

Deploy Snyk Studio across your development teams with the Agentic Development Security (ADS) installer. In Evo, select Snyk Studio under **Capabilities**, publish the configuration, and deploy the ADS installer through your mobile device management (MDM) tool.

{% hint style="info" %}
You must have Evo enabled for your Snyk Tenant and a Tenant role with full Evo access to open **Settings** and select products. Users restricted to specific Organizations cannot open **Settings**. For access requirements and how to enable Evo, visit [Access and authentication](../evo-by-snyk/access-and-authentication.md).
{% endhint %}

For a developer installing Studio on their own machine, use [local setup](getting-started-with-snyk-studio.md#install-on-your-machine).

Use this rollout sequence:

1. [Choose the products and assistants](#choose-products-and-assistants).
2. [Configure versions and authentication](#configure-versions-and-authentication).
3. [Test the deployment](#test-the-deployment).
4. [Roll out through MDM](#roll-out-through-mdm).
5. [Update or remove the deployment](#update-or-remove-the-deployment).

## Choose products and assistants

Choose the scope of your rollout before publishing the ADS configuration:

| Rollout | Choose it when | Product selection |
| --- | --- | --- |
| Studio alone | Your rollout covers Trusted Output Assurance | Select **Snyk Studio**. |
| Studio with other ADS products | Your rollout also covers Agent Supply Chain Security or Agent Behavior Governance | Select **Snyk Studio** and the additional capabilities in your rollout. |

Apply your selection under **Capabilities** in Evo **Settings**.

Review the [Studio assistant capability matrix](getting-started-with-snyk-studio.md#coding-assistant-support) when planning developer workflows. Distinguish assistants with Secure at inception hooks from assistants using MCP tools and directives.

For the product selection and publication procedure, visit [Choose products](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md#choose-products).

## Configure versions and authentication

Choose how to control product versions before testing the deployment:

| Version choice | Choose it when | Effect on the rollout |
| --- | --- | --- |
| Latest published version | You want subsequent installer runs to use the latest published build | The selected version follows new releases. Machines receive it when the ADS installer runs. |
| Pinned version | You need the same build across the rollout or want to test a release before adopting it | The selected version stays fixed until you change and publish it. Machines receive the change when the ADS installer runs. |

The ADS installer version and the Studio product version are separate choices. For the version selection procedure, visit [Choose a version](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md#choose-a-version).

Use a push key for the Evo Tenant that receives deployment data, and configure the region's API host when required. Each developer still authenticates the Snyk CLI and MCP server individually. A deployment push key does not replace developer scanning authentication.

Choose the ADS installation mode for your deployment. Use the [ADS installation mode guidance](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md#installation-mode) for the scope and requirements of each mode. Studio CLI installation scopes and ADS installation modes are separate configuration mechanisms.

## Test the deployment

Start with a small group that represents your operating systems, coding assistants, and repository types.

1. Publish the ADS configuration in Evo **Settings**.
2. Follow [Install on a single machine](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/install-on-a-single-machine.md) with the intended product selection and version.
3. Complete [Verify an installation](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/verify-an-installation.md).
4. Have each developer authenticate with Snyk and start a new coding-assistant session in a repository.
5. Ask the assistant to scan the repository through MCP. For assistants with hooks, also run a coding task and check that background scans run.
6. Review installation and scanning failures before expanding the rollout. Use [Studio troubleshooting](troubleshooting.md) and [ADS troubleshooting](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/troubleshooting.md) for the corresponding installation route.

A connected MCP server verifies only the MCP integration. It does not confirm that the Secure at inception hooks ran. Include both checks when testing assistants with hook support.

## Roll out through MDM

Choose the deployment guide for your platform:

| Platform | Deployment guide |
| --- | --- |
| macOS | [Deploy with Jamf Pro](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/deploy-with-jamf-pro.md) |
| Windows | [Deploy with Microsoft Intune](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/deploy-with-microsoft-intune.md) |
| Linux | [Install on a single machine](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/install-on-a-single-machine.md) |

Use the ADS installer and deployment command from Evo **Settings**. Apply the product selection and version choices from your pilot, then expand the device scope through your MDM tool.

## Update or remove the deployment

To change the product selection or pinned versions, update and publish the configuration in Evo **Settings**, then run the ADS installer on the affected machines. Publishing a configuration alone does not apply it to machines.

Unselecting a product stops rollout to new machines. It does not remove an existing installation. To remove Studio, follow [Remove an ADS product from a machine](../evo-by-snyk/agentic-development-security-ads/activation-and-deployment/README.md#remove-an-ads-product-from-a-machine).

Keep the installation and removal procedures together in your MDM playbook so administrators use the procedure that matches the deployment.
