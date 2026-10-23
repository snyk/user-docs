---
description: >-
  How to deploy Agentic Development Security to Windows devices with Microsoft
  Intune
---

# Deploy with Microsoft Intune

Deploy Agentic Development Security (ADS) to Windows devices with Microsoft Intune. In the standard install, the package delivers the installer only, and a remediation script then runs the installer with your Tenant credentials, which installs and configures the products you selected. In managed mode, the package installs and onboards in one step, and you deploy no script.

## Choose a method

Intune supports two deployment methods. Assign each group of devices to one method only.

|                       | Standard install                                              | Managed mode                                                                         |
| --------------------- | ------------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| **Products installed**| The products you selected, including Studio                   | Scan and Guard. Studio is not installed.                                             |
| **Applies to**        | The signed-in user                                            | Every user on the machine, and protected from removal by the people using them       |
| **How you deploy**    | A line-of-business app and a remediation script               | A line-of-business app with command-line arguments. No remediation or other script is needed. |
| **Codex governance**  | Each developer must accept a prompt in Codex                  | No prompt                                                                            |

The steps below describe the standard install. For managed mode, visit [Managed mode](deploy-with-microsoft-intune.md#managed-mode).

{% hint style="warning" %}
**Codex in user mode requires each developer to accept a prompt**

The standard install runs in user mode. After installation, Codex prompts each developer to allow hook configuration outside the managed path. If a developer does not accept, the hooks stay installed but do not run, and Agent Behavior Governance does not govern Codex on that device. If developers on these devices use Codex, use managed mode instead, which does not show the prompt. Managed mode does not install Studio.
{% endhint %}

## Prerequisites

You need the following:

* a published product selection: in **Settings**, select your products and click **Save & Publish**
* your Tenant ID from **Settings**, and a push key. To create a push key, visit [Manage push keys](./#manage-push-keys).
* your base URL, if you log in outside the default US region
* the signed Windows package, `snyk-ads-installer-windows-x86_64.msi`, from **Settings**, where you also choose the installer version. If your automation cannot use the UI, download it from `downloads.snyk.io/ads/installer/latest/`.
* for the standard install, a Microsoft license that includes Intune Remediations. Check Microsoft's licensing requirements for Remediations before you start.

## Deploy

1. Log in to the Microsoft Intune admin center.
2.  Navigate to **Apps** > **Windows** > **Add**, and upload the `.msi` as a line-of-business app, for example `Snyk ADS Installer Windows`. On the **App information** page, set:

    * **Publisher**: Snyk Ltd
    * **App install context**: **Device**
    * **Command-line arguments**: leave blank

    Assign the app to your Windows devices. With no arguments, the package copies the installer to `C:\Program Files\Snyk\ADS Installer\` and does not onboard your Tenant.
3.  Navigate to **Devices** > **Scripts and remediations** > **Remediations** > **Create script package**, and create a remediation named `Snyk ADS – onboard tenant`.

    Detection script:

    ```powershell
    $Path = "C:\Program Files\Snyk\ADS Installer\snyk-ads-installer-windows-x86_64.exe"

    if (Test-Path $Path) {
        Write-Output "Installer is present - ready to install."
        Exit 1
    }
    Exit 0
    ```

    Intune runs the remediation when the detection script exits with a non-zero code. This script exits 1 when the installer is present, so the remediation runs on every schedule. Each run fetches your current published configuration and skips binaries that already match by checksum.

    Remediation script:

    ```powershell
    $exe = "C:\Program Files\Snyk\ADS Installer\snyk-ads-installer-windows-x86_64.exe"
    if (-not (Test-Path $exe)) {
      Write-Host "Snyk ADS installer not found at $exe"
      exit 1
    }

    $tenantId = "<TENANT_ID>"   # replace with your Tenant ID
    $pushKey  = "<PUSH_KEY>"    # replace with your push key
    $baseUrl  = ""              # leave empty for the US default

    $installerArgs = @("--tenant-id", $tenantId, "--push-key", $pushKey)
    if ($baseUrl) { $installerArgs += @("--base-url", $baseUrl) }

    & $exe @installerArgs
    exit $LASTEXITCODE
    ```

    Leave `--install-mode` unset. The installer then takes the mode from your Tenant's configuration, and uses user mode when none is set. The remediation runs as the signed-in user without elevation, so keep your Tenant configuration on user mode for devices that use the standard install.
4. On the remediation's **Settings** page, set:
   * **Run this script using the logged-on credentials**: **Yes**
   * **Enforce script signature check**: **No**
   * **Run script in 64-bit PowerShell**: **Yes**
5. On **Assignments**, target the same devices as your app, choose how often the remediation runs, and save. A daily schedule keeps device configuration aligned with changes you make in Snyk.

Onboarding runs the next time each targeted device checks in: within an hour for an enrolled device, or immediately when you click **Sync** on the device page. The remediation runs only when a user is signed in. A device with no signed-in user is onboarded on a later run, after someone signs in.

{% hint style="warning" %}
**Run this script using the logged-on credentials** must be **Yes**. On Windows, the installer runs as the account that invokes it, so ADS installs for the account the remediation runs under. If you set it to **No**, the remediation runs as SYSTEM and installs ADS for SYSTEM rather than for the person using the machine. The signed-in user then has nothing installed.
{% endhint %}

If you are not ready to move to packages, the previous command-line install is still supported, and you can run it from your existing deployment script. For the command per platform, visit [Install from the command line](install-on-a-single-machine.md#install-from-the-command-line). The command-line install always takes the latest installer, so you cannot pin a version that way.

## If you log in outside the United States

The base URL mirrors the URL you use to log in to Evo. If you log in at `evo.snyk.io`, leave it out. Otherwise, use your region's base URL, as listed in [If you log in outside the United States](./#if-you-log-in-outside-the-united-states). In the standard install, set `$baseUrl` in the remediation script. In managed mode, add `BASEURL=<BASE_URL>` to the command-line arguments.

## Verify the deployment

To confirm the deployment, visit [Verify an installation](verify-an-installation.md).

If the Tenant was not onboarded, for example because the push key was wrong, Intune reports the remediation as failed even though the app installed.

## Apply a configuration change

Publish the change in **Settings** first. The `Snyk ADS – onboard tenant` remediation runs on its own schedule and fetches your current published configuration each time, so changes apply on the next run. To apply a change sooner, click **Sync** on the device, or shorten the interval between remediation runs.

Each run skips binaries that already match by checksum, so repeated runs are cheap. For managed mode, visit [Apply a configuration change in managed mode](deploy-with-microsoft-intune.md#apply-a-configuration-change-in-managed-mode).

## Upgrade the installer

Replace the `.msi` on the app and redeploy. Intune detects the new package by its product code and reinstalls it automatically. In the standard install, the remediation then runs the updated installer on its next schedule.

## Managed mode

Managed mode is an opt-in for companies that need the products to apply to every user on a machine and to be protected from removal by the people using it. It requires ADS installer v1.88.0 or later, with Machines and Agent Behavior Governance on version 0.6.4 or later. For how the two modes differ, visit [Installation mode](./#installation-mode).

Use managed mode when Scan and Guard should cover every local user on a device, including shared machines and machines where nobody has signed in yet. Studio is not installed in managed mode. To include Studio, use the standard install.

Managed mode needs no script. Do not create a detection script or a remediation, and do not run the installer from PowerShell. To select managed mode, pass the credentials and the mode as command-line arguments on the line-of-business app. The `.msi` installs and onboards in one step:

```
TENANTID=<TENANT_ID> PUSHKEY=<PUSH_KEY> INSTALLMODE=managed /qn
```

`/qn` keeps the `msiexec` run fully silent, with no interface and no prompts. If you log in outside the default US region, add `BASEURL=<BASE_URL>` before `/qn`. **App install context** must be **Device**. Assign the app to a dedicated device group, separate from any group that uses the standard install.

{% hint style="warning" %}
Always include `INSTALLMODE=managed`. With **App install context** set to **Device**, Intune runs the `.msi` as SYSTEM, and the installer runs as whoever starts `msiexec`, which is SYSTEM as well. Managed mode needs that elevated run, and it places the products in machine-wide paths that the machine's users cannot change. If you leave `INSTALLMODE` out, the installer writes the products into SYSTEM's profile and the signed-in user is not onboarded. For a per-user install, leave the command-line arguments blank and use the standard install.
{% endhint %}

If the Tenant was not onboarded, for example because the push key was wrong, the install fails and Intune reports the whole app installation as failed. It does not report partial success.

### Apply a configuration change in managed mode

After Intune installs the app on a device, it does not run the install command again on its own. Intune detects the app by the package's product code, which does not change when you change your Tenant's configuration. To apply a configuration change, publish it in **Settings**, then do one of the following:

* In Intune, open the app's device install status, select the device, and use **Retry**.
* Click **Sync** on the device. This is enough when the app was left in a non-compliant state, for example after a failed onboarding.

Repeating the install is safe. The installer fetches your current published configuration each time and skips binaries that already match by checksum. To apply changes on a recurring schedule instead of on demand, automate periodic **Retry** actions, for example through the Microsoft Graph API.

To confirm a managed mode deployment, visit [Verify an installation](verify-an-installation.md).
