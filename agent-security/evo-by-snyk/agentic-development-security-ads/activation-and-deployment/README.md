---
nav_context: classic
description: >-
  How to configure Agentic Development Security and install it on one machine or
  across your company
---

# Activation and deployment

You configure and deploy Agentic Development Security (ADS) from **Settings** in Evo. Choose which products to roll out, then install on a single machine or across your company through your mobile device management (MDM) tool.

{% hint style="info" %}
The **Settings** page requires a Tenant role with full Evo access. A user restricted to specific Organizations cannot open it. To learn more, visit [Access and authentication](../../access-and-authentication.md).
{% endhint %}

## Authenticate

ADS uses a push key to bind installed agents to your Snyk Tenant, so each machine's data lands in your Tenant. You can rotate the key at any time.

## Choose products

Under **Capabilities**, select the products to roll out:

* **Machines** (Agent Supply Chain Security)
* **Agent Behavior Governance**
* **Snyk Studio** (Trusted Output Assurance)

Each product defaults to the latest published version. You can pin a specific version instead, as described in [Choose a version](./#choose-a-version).

After you select the products, click **Save & Publish**. Publishing writes the configuration manifest that the installer reads when it runs.

## Install

The manifest is the published record of which products your Tenant has selected. The installer fetches it each time it runs.

Before you install, ensure that you have published the configuration. If you changed the product selection since then, click **Save & Publish** again before you install.

### Choose a deployment method

| Method                                                          | Use it for                                                                                 |
| --------------------------------------------------------------- | ------------------------------------------------------------------------------------------ |
| [Install on a single machine](install-on-a-single-machine.md)   | One machine, for testing or a demo. Also the install route on Linux, which has no package. |
| [Deploy with Jamf Pro](deploy-with-jamf-pro.md)                 | macOS devices managed with Jamf Pro                                                        |
| [Deploy with Microsoft Intune](deploy-with-microsoft-intune.md) | Windows devices managed with Microsoft Intune                                              |

After you install, confirm the result with [Verify an installation](verify-an-installation.md). If the install or uninstall returns an error, visit [Troubleshooting](troubleshooting.md).

### Get the installer

In **Settings**, under **Install on a local machine or your MDM tool of choice**, choose your **Operating system** and **Architecture**, then click **Download ADS Installer**. You get a `.pkg` on macOS and an `.msi` on Windows. The command shown below the button is the one you run after installing, pointing at the installed binary.

Linux has no package. On Linux, use the command-line install shown in **Settings**.

### Choose a version

Two versions are involved in an install:

* The **installer version** is the version of the ADS installer itself. You choose it when you download the package.
* The **product versions** decide which build of each product the installer sets up. Under **Capabilities**, each product has a version selector that defaults to the latest published version. Machines and Agent Behavior Governance share one version. Snyk Studio is versioned separately.

Pin a product version when every machine in a rollout needs the same build. A pinned version does not change until you change it, and the installer applies it the next time it runs.

{% hint style="warning" %}
Scheduled scanning requires version 0.6.8 or later of Machines. Earlier versions run one scan each time the installer runs, not on a schedule.
{% endhint %}

### Installation mode

The installer runs in one of two modes. If you do not pass `--install-mode`, the installer uses the mode set in your Tenant's configuration, and uses user mode when none is set.

**User mode** needs no elevated privileges, and the installer sets up the products for the user who runs it.

{% hint style="warning" %}
**Codex in user mode**

After installation in user mode, Codex prompts the developer to allow hook configuration outside the managed path. If the developer does not accept, Agent Behavior Governance does not govern Codex on that machine. Managed mode does not show this prompt, so use managed mode for fleets where Codex is in use.
{% endhint %}

**Managed mode** is for companies that need the products to apply to every user on a machine and to be protected from removal by the people using them. Add `--install-mode managed` and run the installer with elevated privileges. Managed mode requires ADS installer v1.88.0 or later, with Machines and Agent Behavior Governance on version 0.6.4 or later. Existing user-mode installs are not migrated.

The mode applies to the whole installer run. Agent Supply Chain Security and Agent Behavior Governance both take the mode you pass. There is no per-component setting, so you cannot install one in user mode and the other in managed mode.

For Agent Supply Chain Security, the mode changes where the installer registers the scheduled task. In user mode, the installer registers it for the user who ran the installer. In managed mode, the installer registers it at the system level, so a non-privileged user cannot remove or disable it.

The scan always runs at user level, with that user's permissions, in both modes. Managed mode protects the scheduled task from removal. It does not give the scan elevated permissions.

### Scheduled scans

Agent Supply Chain Security scans run every 24 hours by default on macOS and Windows. The installer registers a scheduled job on the machine, and that job runs the scan and exits. Your MDM tool does not need to run the installer again to trigger a scan. Scheduled scans are not supported on Linux.

The first scan runs at install time rather than waiting for the job's first run. Running the installer again within 24 hours of the last scan does not trigger another scan or disturb the schedule, so you can safely run the installer on a frequent MDM policy.

The scheduled job runs the scan only. It does not fetch changes to your product selection or versions. Those apply the next time the installer itself runs on the machine.

#### Change the schedule

Your Tenant's configuration sets whether scans are scheduled and how often. To override it on a machine, add these flags when you run the installer:

| Flag                                | What it does                                                                                                                                          |
| ----------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--start-interval-second <SECONDS>` | Sets the time between scheduled scans, in seconds. For example, `43200` scans every 12 hours. `0` uses the interval from your Tenant's configuration. |
| `--scheduler launchd-agent`         | Registers the scheduled job. Despite its name, this value works on both macOS and Windows.                                                            |
| `--scheduler none`                  | Registers no scheduled job. The machine scans only when the installer runs, so your MDM policy sets how often scans happen.                           |

For example, to scan every 12 hours:

```bash
<PATH_TO_INSTALLER> --tenant-id <TENANT_ID> --push-key <PUSH_KEY> --start-interval-second 43200
```

{% hint style="info" %}
The Windows `.msi` accepts only `TENANTID`, `PUSHKEY`, `INSTALLMODE`, and `BASEURL`. A `.msi` install that onboards in one command gets the schedule from your Tenant's configuration. To override the schedule on Windows, run the installer executable directly with the flags above, for example from the Intune remediation script.
{% endhint %}

### What the installer does

When the installer runs, it sets up each selected product on the machine:

* **Agent Supply Chain Security** discovers the skills and MCP servers in the known directories, performs a risk assessment, and sends the results to the Evo Tenant associated with the push key.
* **Agent Behavior Governance** checks whether the supported agents have the required hooks configured. If they do not, it writes them with the push key, so subsequent agent activity is pushed to Evo and evaluated against your Tenant's policies.
* **Trusted Output Assurance** checks whether the required configuration is in place: the Snyk CLI and its MCP server, the Secure at inception hooks, and the package health check, fix commands, and skills. It adds anything missing, so subsequent agent activity runs through the Secure-at-inception loop. Each developer still authenticates the Snyk CLI and MCP server individually.

## If you log in outside the United States

Every deployment method needs your region's API host when you log in outside the default US region. The host mirrors the URL you log in with, so `evo.eu.snyk.io` uses `https://api.eu.snyk.io`.

| Region               | Base URL                 |
| -------------------- | ------------------------ |
| SNYK-US-01 (default) | not needed               |
| SNYK-US-02           | `https://api.us.snyk.io` |
| SNYK-EU-01           | `https://api.eu.snyk.io` |
| SNYK-AU-01           | `https://api.au.snyk.io` |

Each deployment page shows where to set the base URL for that method.

{% hint style="warning" %}
A push key works only against its own region. If you use the wrong base URL, authentication fails. The installer does not fall back to another region.
{% endhint %}

## Uninstall

Unselecting a product in ADS settings and running `uninstall` do different things.

| What you want to do                                                | Use this                                |
| ------------------------------------------------------------------ | --------------------------------------- |
| Stop a product from being rolled out to new machines in your fleet | Unselect it in ADS settings             |
| Remove a product from a machine that already has it                | Run the installer's `uninstall` command |

### Stop rolling out a product to new machines

Unselecting a product in ADS settings stops the installer from setting it up on machines that do not have it yet.

It does not change machines that already have the product. The installer does not reconcile existing machines and removes nothing on later runs. The scheduled scan keeps running, and Agent Behavior Governance hooks stay in place. To remove a product from those machines, run the `uninstall` command described in [Remove an ADS product from a machine](./#remove-an-ads-product-from-a-machine).

1. In ADS settings, unselect the product.
2. Click **Save & Publish**. Publishing updates the manifest, and the installer reads it the next time it runs on a new machine.

### Remove an ADS product from a machine

Use `uninstall` to remove a product from a machine. Unselecting the product in ADS settings does not remove it. To remove a product across a fleet, run `uninstall` on each machine, for example through your MDM tool.

{% hint style="warning" %}
**Run `uninstall` before you remove the package.**

Removing the package removes the installer, not the products. The products and the scheduled scan stay on the machine, and the tool you need to remove them is gone. This applies on Windows, where `msiexec /x` and an Intune uninstall assignment remove only the installer, and on macOS, where a `.pkg` has no uninstaller.
{% endhint %}

#### Find the installer

The `uninstall` command is part of the ADS installer, so the installer must be on the machine. Its location depends on how you installed.

If you installed from a package, the installer is already on the machine:

* macOS: `/usr/local/bin/snyk-ads/snyk-ads-installer`
* Windows: `C:\Program Files\Snyk\ADS Installer\snyk-ads-installer-windows-x86_64.exe`

If you installed from the command line on macOS or Linux, the command deleted the installer after the install finished. Download the installer again from **Settings**, or run your original command again without the step that deletes it, for example `&& rm -f /tmp/snyk-ads-installer-macos-arm64`.

If the install used a custom installation directory, pass the same directory to `uninstall` with `--install-dir <DIRECTORY>`.

#### Remove a managed install

A user-mode install does not need `sudo` to uninstall. A managed install does: `uninstall` does not read the installation mode from your Tenant's configuration and assumes user mode, so to remove managed products, run it with elevated privileges and add `--install-mode managed`.

```bash
sudo <PATH_TO_INSTALLER> uninstall --all --install-mode managed
```

On Windows, run the same command from an elevated PowerShell, without `sudo`.

#### Remove all products

`--all` removes all three ADS products from the machine: Machines, Agent Behavior Governance, and Snyk Studio. It removes them whether or not they are still selected in ADS settings.

```bash
<PATH_TO_INSTALLER> uninstall --all
```

#### Remove specific products

To remove one or two products rather than all three, pass a component flag for each product you want to remove. You cannot combine `--all` with component flags. You must pass one or the other, because the installer rejects `uninstall` with no flags.

| ADS product                                | Flag       |
| ------------------------------------------ | ---------- |
| **Machines** (Agent Supply Chain Security) | `--scan`   |
| **Agent Behavior Governance**              | `--guard`  |
| **Snyk Studio** (Trusted Output Assurance) | `--studio` |

```bash
<PATH_TO_INSTALLER> uninstall --tenant-id <TENANT_ID> --scan
```

`--scan` also requires `--tenant-id`, so that the installer removes only that Tenant's scheduled scan and configuration.

If you pass both `--tenant-id` and `--push-key`, the installer reports the uninstall to Evo. Without them, the installer still removes the products.
