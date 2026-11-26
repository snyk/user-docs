---
last_reviewed: 2026-10-23T00:00:00.000Z
last_reviewed_source: focused-edit
description: >-
  How to deploy Agentic Development Security to Windows and macOS devices with
  Microsoft Intune
---

# Deploy with Microsoft Intune

Deploy Agentic Development Security (ADS) to Windows and macOS devices with Microsoft Intune. On Windows, Intune supports two methods. Assign each group of Windows devices to one method only. On macOS, use the [macOS managed install](deploy-with-microsoft-intune.md#macos-managed-install).

* Managed install: The `.msi` is a line-of-business app, and your Tenant credentials are command-line arguments. One install places the installer and onboards the device in managed mode. Machines, Agent Behavior Governance, and the managed hooks for Snyk Studio are installed to machine-wide paths and protected from removal by the people using the device. The managed configuration covers Claude, Cursor, Gemini CLI, and Codex. Assistants that exist only in a developer's own configuration are not installed.
* User-level install: The `.msi` is deployed with no command-line arguments, so it only places the installer. A remediation script then runs the installer as the signed-in user and configures Machines, Agent Behavior Governance, and Snyk Studio in that user's profile.

## Choose a method

| Feature                | User-level install                                                                    | Managed install                                                                                                                       |
| ---------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| **Products installed** | Machines, Agent Behavior Governance, and Snyk Studio, in the signed-in user's profile | Machines and Agent Behavior Governance for every local user, plus Snyk Studio managed hooks for Claude, Cursor, Gemini CLI, and Codex |
| **Applies to**         | The signed-in user                                                                    | Every user on the machine. The products are protected from removal.                                                                   |
| **How you deploy**     | A line-of-business app and a remediation script                                       | A line-of-business app with command-line arguments, and no remediation script                                                         |
| **Codex governance**   | Each developer must accept a prompt                                                   | No prompt                                                                                                                             |

{% hint style="warning" %}
**Codex in user mode requires each developer to accept a prompt**

The user-level install runs in user mode. After installation, Codex prompts each developer to allow hook configuration outside the managed path. If a developer does not accept, the hooks stay installed but do not run, and Agent Behavior Governance does not govern Codex on that device. If developers on these devices use Codex, use the managed install, which does not show the prompt.
{% endhint %}

{% hint style="info" %}
To grant Full Disk Access to Macs that Intune manages, visit [Grant Full Disk Access on macOS](grant-full-disk-access-on-macos.md). The profile applies to managed mode only.
{% endhint %}

## Prerequisites

You need the following:

* A published product selection: in **Settings**, select your products and click **Save & Publish**
* Your Tenant ID from **Settings**, and a push key. To create a push key, visit [Manage push keys](./#manage-push-keys).
* Your base URL, if you log in outside the default US region. The values are listed in [If you log in outside the United States](./#if-you-log-in-outside-the-united-states).
* The signed Windows package, `snyk-ads-installer-windows-x86_64.msi`, from **Settings**, where you also choose the installer version. If your automation cannot use the UI, download it from `downloads.snyk.io/ads/installer/latest/`.
* For macOS, the signed package that matches the Mac processor:
  * `snyk-ads-installer-macos-arm64.pkg` for Apple silicon
  * `snyk-ads-installer-macos-x86_64.pkg` for Intel
* For macOS, an Intune-enrolled Mac running macOS 15 or later, with the Microsoft Intune management agent version 2309.007 or later, so the post-install script can run
* For the user-level install, a Microsoft license that includes Intune Remediations. Check Microsoft's licensing requirements for Remediations before you start.

## Managed install

{% hint style="info" %}
The managed install needs no script. Do not create a detection script or a remediation, and do not run the installer from PowerShell.
{% endhint %}

Use the managed install when Machines and Agent Behavior Governance must cover every local user on a device, including shared machines and machines where nobody has signed in yet. When the Snyk Studio version for your Tenant is installed through the Snyk CLI, this install also writes Studio's managed hooks for Claude, Cursor, Gemini CLI, and Codex. Assistants that have no managed configuration are not installed. A legacy studio-recipes pin is not installed in managed mode.

Managed mode requires ADS installer v1.88.0 or later, with Machines and Agent Behavior Governance on version 0.6.4 or later. For how the two modes differ, visit [Installation mode](./#installation-mode).

{% stepper %}
{% step %}
### Log in to Intune

Log in to the [Microsoft Intune admin center](https://intune.microsoft.com).
{% endstep %}

{% step %}
### Create and assign the app

Navigate to **Apps** > **Windows** > **Create**, and upload the `.msi` as a **Line-of-business app**, for example `Snyk ADS Installer Windows`. On the **App information** page, set:

* **Publisher**: Snyk Ltd
* **App install context**: **Device**
*   **Command-line arguments**:

    ```
    TENANTID=<TENANT_ID> PUSHKEY=<PUSH_KEY> INSTALLMODE=managed /qn
    ```

    `/qn` keeps the `msiexec` run fully silent, with no interface and no prompts. If you log in outside the default US region, add `BASEURL=<BASE_URL>` before `/qn`.

    The `.msi` does not pass the push key to the installer's command line or write it to the verbose MSI log. The push key is still stored in the Intune app configuration and passed to `msiexec`, so restrict access to the app configuration. Do not enable the machine-wide Windows Installer `Debug=7` policy, which logs command-line properties even when an MSI marks them as hidden.
{% endstep %}
{% endstepper %}

Assign the app to the devices for managed mode. Use a separate group from any group assigned to the [user-level install](deploy-with-microsoft-intune.md#user-level-install).

{% hint style="warning" %}
Always include `INSTALLMODE=managed`.

With **App install context** set to **Device**, Intune runs the `.msi` as SYSTEM, and the installer runs as whoever starts `msiexec`, which is SYSTEM as well. Managed mode needs that elevated run, and it places the products in machine-wide paths that the machine's users cannot change. If you omit `INSTALLMODE`, or set it to `user`, on this device-context install, and your Tenant's installation mode is not set to managed, the installer runs a user-mode install for the person signed in to the device instead of a machine-wide install. If nobody is signed in, the install fails with `no user-scoped install target available`. If the installer cannot look up the account of the person signed in, the install fails with `lookup console user`. For a per-user install, leave the command-line arguments blank and use the [user-level install](deploy-with-microsoft-intune.md#user-level-install).
{% endhint %}

Onboarding runs the next time each targeted device checks in: within an hour for an enrolled device, or immediately when you click **Sync** on the device page. Each later run fetches your current published configuration and skips binaries that already match by checksum. To run the install again, visit [Apply a configuration change to a managed install](deploy-with-microsoft-intune.md#apply-a-configuration-change-to-a-managed-install).

### Verify the managed install

After Intune reports the app as installed, confirm the result from an elevated PowerShell session. Managed mode places the products in machine-wide paths owned by the system:

{% code overflow="wrap" expandable="true" %}
```powershell
& "C:\Program Files\Snyk\ADS Installer\snyk-ads-installer-windows-x86_64.exe" --version
Get-ChildItem "$env:ProgramFiles\Snyk\ADS\ads-scan\bin" -File | Select-Object Name, Length
Get-ChildItem "$env:ProgramData\Snyk\ads-scan" -File | Select-Object Name, Length
schtasks /Query /TN "\Snyk\ADS\" /V /FO LIST 2>$null | Select-String "TaskName|Run As User"
```
{% endcode %}

`Run As User` for the scheduled task shows `BUILTIN\Users`. That principal applies the scan to every local user.

When Snyk Studio is enabled with a CLI pin, the same install places the Snyk CLI at `C:\Program Files\Snyk\ads-scan\bin\snyk-cli-windows-x86_64.exe` and records the machine in `%ProgramData%\snyk-studio\install-state.json`. Hooks for Claude, Cursor, Gemini CLI, and Codex are in each assistant's managed configuration.

To confirm the machine in Evo, visit [Verify an installation](verify-an-installation.md#managed-mode-on-windows).

If the Tenant was not onboarded, for example because the push key was wrong, the install fails and Intune reports the whole app installation as failed. It does not report partial success.

### Apply a configuration change to a managed install

Publish the change in **Settings** first. After Intune installs the app on a device, it does not run the install command again on its own. Intune detects the app by the package product code, which does not change when you change your Tenant's configuration. **Sync** on the device requests a check-in. It does not reinstall an app that Intune already reports as installed.

To force a refresh, remove the device group from the app's **Required** assignment, then add that same group back:

1. Open the app in Intune and navigate to **Properties** > **Assignments** > **Edit**.
2. Under **Required**, remove the device group that covers the devices you want to refresh, then review and save.
3. Wait until those devices are no longer targeted by the assignment.
4. Edit **Assignments** again, add the same device group under **Required**, and save.

Devices run the install again on their next check-in after the group is added back. You can repeat this. Each run fetches your current published configuration and skips binaries that already match by checksum. The `.msi` command-line arguments are already the full onboarding command.

To upgrade the installer, replace the `.msi` on the app and redeploy. Intune detects the new package by its product code and reinstalls it.

## User-level install

The `.msi` only places the installer on the device. A remediation script completes onboarding with your Tenant credentials, running as the signed-in user, and installs Machines, Agent Behavior Governance, and Snyk Studio into that user's profile.

1. Log in to the [Microsoft Intune admin center](https://intune.microsoft.com).
2.  Navigate to **Apps** > **Windows** > **Create**, and upload the `.msi` as a **Line-of-business app**, for example `Snyk ADS Installer Windows`. On the **App information** page, set:

    * **Publisher**: Snyk Ltd
    * **App install context**: **Device**
    * **Command-line arguments**: leave blank

    Assign the app to your Windows devices. With no arguments, the package copies the installer to `C:\Program Files\Snyk\ADS Installer\` and does not onboard your Tenant.
3. Navigate to **Devices** > **Scripts and remediations** > **Remediations** > **Create script package**, and create a remediation:
   * **Name**: `Snyk ADS – onboard tenant`
   *   **Detection script**:

       <pre class="language-powershell" data-overflow="wrap" data-expandable="true"><code class="lang-powershell">$Path = "C:\Program Files\Snyk\ADS Installer\snyk-ads-installer-windows-x86_64.exe"

       if (Test-Path $Path) {
           Write-Output "Installer is present - ready to install."
           Exit 1
       }
       Exit 0
       </code></pre>

       Intune runs the remediation when the detection script exits with a non-zero code. This script exits 1 if the installer is present, so the remediation runs on every schedule. Each run fetches your current published configuration and skips binaries that already match by checksum.
   *   **Remediation script**:

       <pre class="language-powershell" data-overflow="wrap" data-expandable="true"><code class="lang-powershell">$exe = "C:\Program Files\Snyk\ADS Installer\snyk-ads-installer-windows-x86_64.exe"
       if (-not (Test-Path $exe)) {
         Write-Host "Snyk ADS installer not found at $exe"
         exit 1
       }

       $tenantId = "&#x3C;TENANT_ID>"   # replace with your Tenant ID
       $pushKey  = "&#x3C;PUSH_KEY>"    # replace with your push key
       $baseUrl  = ""              # leave empty for the default US region

       $installerArgs = @("--tenant-id", $tenantId, "--push-key", $pushKey, "--install-mode", "user")
       if ($baseUrl) { $installerArgs += @("--base-url", $baseUrl) }

       &#x26; $exe @installerArgs
       exit $LASTEXITCODE
       </code></pre>

       The script passes `--install-mode user`, which takes precedence over the installation mode in your Tenant's configuration, so the user-level install runs in user mode even when your Tenant is set to managed. The script runs as the signed-in user without elevation, and a managed install cannot run that way. If you log in outside the default US region, set `$baseUrl` to your region's base URL, as listed in [If you log in outside the United States](./#if-you-log-in-outside-the-united-states).
4.  On the remediation's **Settings** page, set:

    * **Run this script using the logged-on credentials**: **Yes**
    * **Enforce script signature check**: **No**
    * **Run script in 64-bit PowerShell**: **Yes**

    <div data-gb-custom-block data-tag="hint" data-style="warning" class="hint hint-warning"><p><strong>Run this script using the logged-on credentials</strong> must be <strong>Yes</strong>. With it set to <strong>No</strong>, the script runs as SYSTEM. The installer then has to find the person signed in to the device to install for. The install fails with <code>no user-scoped install target available</code> when nobody is signed in, and with <code>lookup console user</code> when the installer cannot look up the account of the person signed in.</p></div>
5. On **Assignments**, target the same devices as the app, choose how often the remediation runs, and save. A daily schedule keeps device configuration aligned with changes you make in Snyk.

Onboarding runs the next time each targeted device checks in: within an hour for an enrolled device, or immediately when you click **Sync** on the device page. The remediation runs only when a user is signed in. A device with no signed-in user is onboarded on a later run, after someone has signed in.

If you are not ready to move to packages, the previous command-line install is still supported, and you can run it from your existing deployment script. For the command per platform, visit [Install from the command line](install-on-a-single-machine.md#install-from-the-command-line). The command-line install always takes the latest installer, so you cannot pin a version that way.

### Verify the user-level install

After Intune reports the app and the remediation as succeeded, confirm the result:

{% code overflow="wrap" expandable="true" %}
```powershell
& "C:\Program Files\Snyk\ADS Installer\snyk-ads-installer-windows-x86_64.exe" --version
Get-ChildItem "$env:LOCALAPPDATA\Snyk\ads-scan\" -File | Select-Object Name, Length
schtasks /Query /TN "\Snyk\ADS\" 2>$null
```
{% endcode %}

To confirm the machine in Evo, visit [Verify an installation](verify-an-installation.md).

If the Tenant was not onboarded, for example because the push key was wrong, Intune reports the remediation as failed even though the app installed.

### Apply a configuration change to a user-level install

Publish the change in **Settings** first. The `Snyk ADS – onboard tenant` remediation runs on its own schedule and fetches your current published configuration each time, so changes apply on the next run. To apply a change sooner, click **Sync** on the device, or shorten the interval between remediation runs.

Each run skips binaries that already match by checksum, so repeated runs are cheap.

To upgrade the installer, replace the `.msi` on the app and redeploy. Intune detects the new package by its product code and reinstalls it. The remediation then runs the updated installer on its next schedule.

## macOS managed install

The `.pkg` only places the installer on the Mac. A post-install script completes onboarding with your Tenant credentials, running as root in managed mode. Create one app per processor architecture, and assign each app only to matching devices.

1. Log in to the [Microsoft Intune admin center](https://intune.microsoft.com).
2.  Navigate to **Apps** > **All apps** > **Create**, and select **macOS app (PKG)**. Upload the `.pkg` that matches the target devices, for example `Snyk ADS Installer macOS arm64`. On the **App information** page, set:

    * **Publisher**: Snyk Ltd

    Use **macOS app (PKG)** for this deployment. It supports packages installed outside `/Applications`, and it supports a post-install script. The macOS **Line-of-business app** does not include the command-line arguments field that the Windows `.msi` has. Microsoft documents the package type in [Add an unmanaged macOS PKG app to Microsoft Intune](https://learn.microsoft.com/en-us/intune/app-management/deployment/add-unmanaged-pkg-macos).
3.  On the **Program** page, upload the following as the **Post-install script**, saved as `snyk-ads-onboard.sh`:

    <pre class="language-zsh" data-overflow="wrap" data-expandable="true"><code class="lang-zsh">#!/bin/zsh
    set -u

    tenant_id="&#x3C;TENANT_ID>"   # replace with your Tenant ID
    push_key="&#x3C;PUSH_KEY>"     # replace with your push key
    base_url=""               # leave empty for the default US region
    # The package receipt records where this version installed the installer.
    location=$(pkgutil --pkg-info io.snyk.ads-installer 2>/dev/null | sed -n 's/^location: //p')
    installer="/${location#/}/snyk-ads-installer"

    if [[ -z "$tenant_id" || -z "$push_key" ]]; then
      echo "ADS onboarding failed: tenant ID and push key are required."
      exit 2
    fi

    if [[ -z "$location" ]]; then
      echo "ADS onboarding failed: package io.snyk.ads-installer is not installed."
      exit 3
    fi

    if [[ ! -x "$installer" ]]; then
      echo "ADS onboarding failed: installer not found at $installer"
      exit 3
    fi

    args=(
      --install-mode managed
      --tenant-id "$tenant_id"
      --push-key "$push_key"
      --scheduler launchd-agent
      --start-interval-second 900
    )
    if [[ -n "$base_url" ]]; then
      args+=(--base-url "$base_url")
    fi

    "$installer" "${args[@]}"
    exit $?
    </code></pre>

    The script runs the installer from the location in the `io.snyk.ads-installer` package receipt, so the same script works for any package version. Installer 1.93.0 and later use `/Library/Application Support/Snyk/ads-scan/bin`. Earlier packages use `/usr/local/bin/snyk-ads`, including after a downgrade to an older package.

    Replace the credentials in the script. This macOS app type has no separate parameter fields. Leave `base_url` empty when you log in at `evo.snyk.io`. Otherwise set it to your region's base URL, as listed in [If you log in outside the United States](./#if-you-log-in-outside-the-united-states).

    `--install-mode managed` selects the same mode as `INSTALLMODE=managed` on Windows. The scheduling flags configure Machines to run through launchd every 15 minutes when Machines is enabled. They do not enable Machines when it is disabled in your Tenant's configuration.
4. On **Requirements**, set **Minimum operating system** to macOS 15, or a later version your fleet requires. Review **Detection rules**. On **Assignments**, assign the app as **Required** to the device group for that processor architecture. Review the settings and create the app.
5. Repeat for the other architecture if you need it, using the same script and Tenant values.

Onboarding runs after Intune installs the package. To request a check-in, open Company Portal on the Mac, select the device, and choose **Check status**. **Check status** requests both an MDM check-in and an agent check-in. **Sync** in the Intune admin center requests only an MDM check-in. Microsoft documents the agent in [Microsoft Intune management agent for macOS](https://learn.microsoft.com/en-us/intune/app-management/deployment/management-agent-macos).

For Macs in managed mode, deploy a configuration profile that gives the scan Full Disk Access. For the profiles and the steps, visit [Grant Full Disk Access on macOS](grant-full-disk-access-on-macos.md).

### Verify the macOS managed install

After Intune reports the app as installed, confirm the result:

```bash
"/Library/Application Support/Snyk/ads-scan/bin/snyk-ads-installer" --version
pkgutil --pkg-info io.snyk.ads-installer
ls -l "/Library/Application Support/Snyk/ads-scan/bin"
find /Library/LaunchAgents -name 'io.snyk.ads.scan*.plist' -print
sudo ls -lt /Library/Logs/Snyk/ads-scan/
```

On packages before 1.93.0, run `/usr/local/bin/snyk-ads/snyk-ads-installer --version` instead.

The scan binary and LaunchAgent are present when Machines is enabled. Read the newest installer log under `/Library/Logs/Snyk/ads-scan` to confirm onboarding finished and which components were installed. A successful run can skip disabled components and not create their managed files.

If the Tenant was not onboarded, for example because the push key was wrong, Intune can still report the app as installed. A failed post-install script does not fail the package installation status. The installer log is what confirms that onboarding succeeded.

### Apply a configuration change on macOS

Publish the change in **Settings** first. After Intune installs the package, changing your Tenant's configuration or editing the post-install script does not guarantee another onboarding run. For a recurring refresh, deploy the same onboarding script as a separate macOS shell script. Run it as root: set **Run script as signed-in user** to **No**, and choose the frequency you need.

You can run the script repeatedly. Each run fetches your current published configuration and skips binaries that already match by checksum.

### Upgrade the macOS installer

Upload the new `.pkg` to the app for that processor architecture, and review its version detection settings.

If onboarding fails, visit [Troubleshooting](troubleshooting.md).

## Remove the installer

Removing the app in Intune deletes only the installer. It does not run `uninstall`, so the products and the scheduled scan stay on the device. They run the installer, so they fail at every interval after it is gone.

Run `uninstall` on the devices first, while the installer is still there, and then remove the app.

For the `uninstall` flags, visit [Remove an ADS product from a machine](./#remove-an-ads-product-from-a-machine).
