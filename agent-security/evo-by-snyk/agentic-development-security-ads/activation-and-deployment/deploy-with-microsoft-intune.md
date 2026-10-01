---
description: >-
  How to deploy Agentic Development Security to Windows devices with Microsoft
  Intune
---

# Deploy with Microsoft Intune

Deploy Agentic Development Security (ADS) to Windows devices with Microsoft Intune. The package delivers the installer only. A remediation script then runs the installer with your Tenant credentials, which installs and configures the products you selected.

## Prerequisites

You need the following:

* a published product selection: in **Settings**, select your products and click **Save & Publish**
* your Tenant ID and push key, both from **Settings**
* the signed Windows package, `snyk-ads-installer-windows-x86_64.msi`, from **Settings**, where you also choose the installer version. If your automation cannot use the UI, download it from `downloads.snyk.io/ads/installer/latest/`.
* licensing that includes Intune Remediations. Microsoft requires Windows Enterprise E3 or E5, Windows Education A3 or A5, or an equivalent license for Remediations.

## Deploy

1. Log in to the Microsoft Intune admin center.
2.  Navigate to **Apps** > **Windows** > **Add**, and upload the `.msi` as a line-of-business app, for example `Snyk ADS Installer Windows`. On the **App information** page, set:

    * **Publisher**: Snyk Ltd
    * **App install context**: **Device**
    * **Command-line arguments**: leave blank

    Assign the app to your Windows devices.
3.  Navigate to **Devices** > **Scripts and remediations** > **Remediations** > **Create script package**, and create a remediation named `Snyk ADS – onboard tenant`.

    Detection script:

    ```powershell
    exit 1
    ```

    The detection script always reports the device as non-compliant on purpose, so the remediation runs on every schedule and picks up your current published configuration.

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
4. On the remediation's **Settings** page, set:
   * **Run this script using the logged-on credentials**: **Yes**
   * **Enforce script signature check**: **No**
   * **Run script in 64-bit PowerShell**: **Yes**
5. On **Assignments**, target the same devices as your app, choose how often the remediation runs, and save.

Onboarding runs the next time each targeted device checks in: within an hour for an enrolled device, or immediately when you click **Sync** on the device page.

{% hint style="warning" %}
**Run this script using the logged-on credentials** must be **Yes**. On Windows, the installer runs as the account that invokes it, so ADS installs for the account the remediation runs under. If you set it to **No**, the remediation runs as SYSTEM and installs ADS for SYSTEM rather than for the person using the machine. The signed-in user then has nothing installed.
{% endhint %}

If you are not ready to move to packages, the previous command-line install is still supported, and you can run it from your existing deployment script. For the command per platform, visit [Install from the command line](install-on-a-single-machine.md#install-from-the-command-line). The command-line install always takes the latest installer, so you cannot pin a version that way.

## If you log in outside the United States

Set `$baseUrl` in the remediation script to the base URL for your region, as listed in [If you log in outside the United States](./#if-you-log-in-outside-the-united-states). Leave it empty for the default US region.

## Verify the deployment

To confirm the deployment, visit [Verify an installation](verify-an-installation.md).

If the Tenant was not onboarded, for example because the push key was wrong, Intune reports the remediation as failed even though the app installed.

## Apply a configuration change

Publish the change in **Settings** first. The `Snyk ADS – onboard tenant` remediation runs on its own schedule and fetches your current published configuration each time, so changes apply on the next run. To apply a change sooner, click **Sync** on the device, or shorten the interval between remediation runs.

Each run skips binaries that already match by checksum, so repeated runs are cheap.

## Upgrade the installer

Replace the `.msi` on the app and redeploy. Intune detects the new package and reinstalls it automatically.

## Managed mode

Managed mode is an opt-in for companies that need the products to apply to every user on a machine and to be protected from removal by the people using it. It requires ADS installer v1.88.0 or later, with Machines and Agent Behavior Governance on version 0.6.4 or later. For how the two modes differ, visit [Installation mode](./#installation-mode).

To select managed mode, pass the credentials and the mode as command-line arguments on the line-of-business app, and do not create the remediation:

```
TENANTID=<TENANT_ID> PUSHKEY=<PUSH_KEY> INSTALLMODE=managed /qn
```

If you log in outside the default US region, add `BASEURL=<BASE_URL>` before `/qn`. **App install context** must be **Device**. Scope this deployment to a dedicated device group, not to the assignment you use for the default flow.
