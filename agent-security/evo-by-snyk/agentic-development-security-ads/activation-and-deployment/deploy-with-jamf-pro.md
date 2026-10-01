---
description: How to deploy Agentic Development Security to macOS devices with Jamf Pro
---

# Deploy with Jamf Pro

Deploy Agentic Development Security (ADS) to managed Macs with Jamf Pro. The package delivers the installer only. A policy script then runs the installer with your Tenant credentials, which installs and configures the products you selected.

## Prerequisites

You need the following:

* a published product selection: in **Settings**, select your products and click **Save & Publish**
* your Tenant ID and push key, both from **Settings**
* the signed macOS packages, one per processor architecture, from **Settings**, where you also choose the installer version. If your automation cannot use the UI, download them from `downloads.snyk.io/ads/installer/latest/`:
  * `snyk-ads-installer-macos-arm64.pkg` for Apple silicon Macs
  * `snyk-ads-installer-macos-x86_64.pkg` for Intel Macs
* macOS 12.0 or later on managed Macs

## Deploy

Create one policy per processor architecture, and scope each policy to matching devices only.

1. Log in to Jamf Pro.
2. Navigate to **Computer Management** > **Packages**, and upload each `.pkg` as a separate package record, for example `Snyk ADS Installer macOS arm64` and `Snyk ADS Installer macOS x86_64`.
3.  Navigate to **Computer Management** > **Scripts** > **New**, and create a script named `Snyk ADS – install and onboard`:

    ```bash
    #!/bin/zsh
    set -u

    tenant_id="${4:-}"
    push_key="${5:-}"
    base_url="${6:-}"
    installer="/usr/local/bin/snyk-ads/snyk-ads-installer"

    if [[ -z "$tenant_id" || -z "$push_key" ]]; then
      echo "ADS onboarding failed: tenant ID and push key are required."
      exit 2
    fi

    if [[ ! -x "$installer" ]]; then
      echo "ADS onboarding failed: installer not found at $installer"
      exit 3
    fi

    args=(--tenant-id "$tenant_id" --push-key "$push_key")
    if [[ -n "$base_url" ]]; then
      args+=(--base-url "$base_url")
    fi

    "$installer" "${args[@]}"
    exit $?
    ```

    Set the script parameter labels: **Parameter 4** to Tenant ID, **Parameter 5** to Push Key, and **Parameter 6** to Base URL (optional).
4. Navigate to **Computers** > **Policies** > **New**, and configure the Apple silicon policy:
   * **Display Name**: `Snyk ADS – macOS arm64 – install and onboard`
   * **Trigger**: **Recurring Check-in**
   * **Execution Frequency**: **Once per computer**
   * **Packages**: the arm64 package only, action **Install**
   * **Scripts**: `Snyk ADS – install and onboard`, **Priority After**. Set **Parameter 4** to your Tenant ID, **Parameter 5** to your push key, and **Parameter 6** to your base URL, or leave **Parameter 6** blank for the default US service.
   * **Scope**: your Apple silicon Macs only
5. Save the policy, then duplicate it for Intel. Change only the display name, the package, and the scope. Keep the same script and parameter values.

Jamf runs policy scripts as root. In user mode, the installer detects this and sets up the products for the user who is logged in to the Mac, not for root. A user must be logged in when the policy runs: a check-in at the login window, with nobody logged in, fails. If your Macs often check in with nobody logged in, use the **Login** trigger instead.

If you are not ready to move to packages, the previous command-line install is still supported, and you can run it from your existing deployment script. For the command per platform, visit [Install from the command line](install-on-a-single-machine.md#install-from-the-command-line). The command-line install always takes the latest installer, so you cannot pin a version that way.

## If you log in outside the United States

Set **Parameter 6** to the base URL for your region, as listed in [If you log in outside the United States](./#if-you-log-in-outside-the-united-states). Leave it blank for the default US region.

## Verify the deployment

To confirm the deployment, visit [Verify an installation](verify-an-installation.md).

If the Tenant was not onboarded, for example because the push key was wrong, the package installs and the policy script fails.

## Apply a configuration change

Publish the change in **Settings** first. A **Once per computer** policy does not repeat on a machine, so create a second policy that runs the installer without a package:

1. Navigate to **Computers** > **Policies** > **New**.
2. Under **General**, set **Trigger** to **Recurring Check-in**, and choose how often to refresh the configuration, for example daily.
3. Under **Scripts**, add `Snyk ADS – install and onboard` with the same parameter values. Do not add a package.
4. Under **Scope**, select the Macs where the installer is already present.

You can run this policy repeatedly. Each run fetches your current published configuration and skips binaries that already match by checksum.

## Upgrade the installer

Replace the package on the policy for the matching architecture, and redeploy.

## Managed mode

Managed mode is an opt-in for companies that need the products to apply to every user on a machine and to be protected from removal by the people using it. It requires ADS installer v1.88.0 or later, with Machines and Agent Behavior Governance on version 0.6.4 or later. For how the two modes differ, visit [Installation mode](./#installation-mode).

To select managed mode, add `--install-mode managed` to the arguments in the policy script:

```bash
args=(--tenant-id "$tenant_id" --push-key "$push_key" --install-mode managed)
```

Jamf policy scripts run as root, which is the elevated invocation managed mode requires. You do not need any other change to the policy.
