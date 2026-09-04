---
nav_context: classic
---

# Activation and deployment

You configure and deploy Agentic Development Security (ADS) from **Settings** in Evo. Choose which products to roll out, then install on a single machine or across your company through your mobile device management (MDM) tool.

{% hint style="info" %}
The **Settings** page requires a Tenant role with full Evo access. A user restricted to specific Organizations cannot open it. To learn more, visit [Access and authentication](../access-and-authentication.md).
{% endhint %}

## Authenticate

ADS uses a push key to bind installed agents to your Snyk Tenant, so each machine's data lands in your Tenant. You can rotate the key at any time.

## Choose products

{% hint style="info" %}
**Feature availability**

Agent Behavior Governance is in open preview. It does not connect to platform surfaces (Inventory, Policies and issues, Reports, and Evo chat) and has limitations described here.
{% endhint %}

Under **Products**, select the products to roll out:

* **Machines** (Agent Supply Chain Security)
* **Agent Behavior Governance**
* **Snyk Studio** (Trusted Output Assurance)

By default, each product stays on the latest version and updates automatically.

## Install

Follow the on-screen instructions to install on a local machine or through your MDM tool.

### What the installer does

When the installer runs, it sets up each selected product on the machine:

* **Agent Supply Chain Security** discovers the skills and MCP servers in the known directories, performs a risk assessment, and sends the results to the Evo Tenant associated with the push key.
* **Agent Behavior Governance** checks whether the supported agents have the required hooks configured. If they do not, it writes them with the push key, so subsequent agent activity is pushed to Evo and evaluated against your Tenant's policies.
* **Trusted Output Assurance** checks whether the required configuration is in place: the Snyk CLI and its MCP server, the Secure at inception hooks, and the package health check, fix commands, and skills. It adds anything missing, so subsequent agent activity runs through the Secure-at-inception loop. Each developer still authenticates the Snyk CLI and MCP server individually.

## Uninstall

There are two ways to uninstall an ADS product.

| What you want to do                              | Use this                                |
| ------------------------------------------------ | --------------------------------------- |
| Stop a product from being used across your fleet | Unselect it in ADS settings             |
| Fully remove a product from a machine            | Run the installer's `uninstall` command |

### Disabling a product across your fleet

Recommended for MDM-managed deployments.

The ADS installer reconciles each machine with the products you select in ADS settings. Unselect a product, and the installer removes it on the next run, whether you trigger it manually or your MDM does.

1. In ADS settings, unselect the product.
2. Run the installer again. There is no separate apply step: re-running the install is what applies the change.
   * On a single machine: run the install command shown in **Settings** page in Evo
   * Across a fleet: re-push the installer through your MDM tool, using the same command or package you deploy with today.
   * If you have a scheduled installer run, you can wait for the next scheduled run.

This is the recommended path for MDM deployments; it requires no change to your MDM scripts or policies.

{% hint style="warning" %}
The installer removes any unselected products on the next run, including a run you start for an unrelated reason.

Disabling a product stops it from running but leaves its files on the machine. Use `uninstall` to remove the files.
{% endhint %}

### Removing an ADS product from a machine

Use the `uninstall` command when you are working on a single machine rather than a fleet.

{% hint style="info" %}
The `uninstall` command lives in the ADS installer, which is not left on the machine if you used the install command on the **Settings** page. The install command on the **Settings** page downloads the installer to a temporary location and removes it once the install finishes, and the installer is not copied into the ADS install directory. The ADS installer is needed to use the `uninstall` command
{% endhint %}

Run the installer with `uninstall` and pass a component flag to specify what to remove:

```
<PATH_TO_INSTALLER> uninstall --tenant-id <TENANT_ID> --push-key <PUSH_KEY> --<FLAG>
```

Replace \<PATH\_TO\_INSTALLER> with the install command from the **Settings** page. It contains the correct installer binary for your operating system and architecture. **Remove** the step at the end of the command that deletes the installer. For Mac OS/ Linux remove `&& rm -f /tmp/snyk-ads-installer-macos-arm64` (this command removal applies to Mac OS/ Linux only)

Replace `<FLAG>` with `--scan`, `--guard`, or `--studio` as shown in the table below.

| ADS product                                | Flag       |
| ------------------------------------------ | ---------- |
| **Machines** (Agent Supply Chain Security) | `--scan`   |
| **Agent Behavior Governance**              | `--guard`  |
| **Snyk Studio** (Trusted Output Assurance) | `--studio` |

Omitting the component flags removes all three ADS products.

`uninstall` with no `--scan`, `--guard`, or `--studio` flag removes Machines, Agent Behavior Governance, and Snyk Studio.

`sudo` is not required.
