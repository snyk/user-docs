---
description: >-
  How to install Agentic Development Security on one machine, without a device
  management tool
---

# Install on a single machine

Install Agentic Development Security (ADS) on your own machine, for testing or a demo. To deploy across a fleet, use your device management tool instead: visit [Deploy with Jamf Pro](deploy-with-jamf-pro.md) or [Deploy with Microsoft Intune](deploy-with-microsoft-intune.md).

Linux has no package. On Linux, use [Install from the command line](install-on-a-single-machine.md#install-from-the-command-line). The package sections of this page apply to macOS and Windows.

## Prerequisites

You need the following:

* a published product selection: in **Settings**, select your products and click **Save & Publish**
* your Tenant ID and push key, both from **Settings**
* the installer version you want: **Settings** defaults to the latest published version
* administrator access on the machine, to install the package
* macOS 12.0 or later, or Windows

{% hint style="info" %}
On macOS, the package delivers the installer only. Installing it does not set up the products or start a scan, so you run the installer with your credentials as a second step. On Windows, the `.msi` installs and onboards in one step when you pass your Tenant ID and push key.
{% endhint %}

## Install

{% tabs %}
{% tab title="macOS" %}
1. In **Settings**, under **Install on a local machine or your MDM tool of choice**, set **Operating system** to macOS and **Architecture** to match your processor, then click **Download ADS Installer (.pkg)**. The package that matches your selection is:
   * `snyk-ads-installer-macos-arm64.pkg` for Apple silicon
   * `snyk-ads-installer-macos-x86_64.pkg` for Intel
2.  Install the package, either by opening the `.pkg` and following the prompts, or from Terminal. Replace `<PACKAGE_FILE>` with the file name for your processor:

    ```bash
    sudo installer -pkg <PACKAGE_FILE> -target /
    ```
3.  Run the installer with your credentials, as yourself and without `sudo`. In user mode, the installer refuses to set up Agent Behavior Governance when it runs with elevated privileges.

    ```bash
    /usr/local/bin/snyk-ads/snyk-ads-installer \
      --tenant-id "<TENANT_ID>" \
      --push-key "<PUSH_KEY>"
    ```

    The first run can take several minutes while the installer downloads the products your Tenant has selected.
{% endtab %}

{% tab title="Windows" %}
1. In **Settings**, under **Install on a local machine or your MDM tool of choice**, set **Operating system** to Windows, then click **Download ADS Installer (.msi)**. This downloads `snyk-ads-installer-windows-x86_64.msi`.
2.  From an elevated PowerShell, install and onboard in one command. The products are set up for the account that runs the command, so run it as the person who uses the machine. The `-Wait` flag and the log file let you see whether the install succeeded:

    ```powershell
    Start-Process msiexec.exe -Wait -ArgumentList '/i snyk-ads-installer-windows-x86_64.msi TENANTID=<TENANT_ID> PUSHKEY=<PUSH_KEY> /qn /l*v ads-install.log'
    ```
{% endtab %}
{% endtabs %}

Replace `<TENANT_ID>` and `<PUSH_KEY>` with your own values before you run a command. The command copied from **Settings** fills in only the Tenant ID.

## Install from the command line

The command-line install is still supported. Use it when you are not ready to move to packages, when your tooling cannot use a downloaded package, or on Linux, where there is no package. The command downloads the installer, runs it with your credentials, and, on macOS and Linux, deletes it. It installs in user mode, so it does not need elevated privileges.

Copy the command for your platform from **Settings**, which fills in your Tenant ID for you, or use the following commands.

{% tabs %}
{% tab title="macOS" %}
Apple silicon:

```bash
curl -fsSL https://downloads.snyk.io/ads/installer/latest/snyk-ads-installer-macos-arm64 -o /tmp/snyk-ads-installer-macos-arm64 && chmod +x /tmp/snyk-ads-installer-macos-arm64 && /tmp/snyk-ads-installer-macos-arm64 --tenant-id <TENANT_ID> --push-key <PUSH_KEY> && rm -f /tmp/snyk-ads-installer-macos-arm64
```

Intel:

```bash
curl -fsSL https://downloads.snyk.io/ads/installer/latest/snyk-ads-installer-macos-x86_64 -o /tmp/snyk-ads-installer-macos-x86_64 && chmod +x /tmp/snyk-ads-installer-macos-x86_64 && /tmp/snyk-ads-installer-macos-x86_64 --tenant-id <TENANT_ID> --push-key <PUSH_KEY> && rm -f /tmp/snyk-ads-installer-macos-x86_64
```
{% endtab %}

{% tab title="Linux" %}
arm64:

```bash
curl -fsSL https://downloads.snyk.io/ads/installer/latest/snyk-ads-installer-linux-arm64 -o /tmp/snyk-ads-installer-linux-arm64 && chmod +x /tmp/snyk-ads-installer-linux-arm64 && /tmp/snyk-ads-installer-linux-arm64 --tenant-id <TENANT_ID> --push-key <PUSH_KEY> && rm -f /tmp/snyk-ads-installer-linux-arm64
```

x86\_64:

```bash
curl -fsSL https://downloads.snyk.io/ads/installer/latest/snyk-ads-installer-linux-x86_64 -o /tmp/snyk-ads-installer-linux-x86_64 && chmod +x /tmp/snyk-ads-installer-linux-x86_64 && /tmp/snyk-ads-installer-linux-x86_64 --tenant-id <TENANT_ID> --push-key <PUSH_KEY> && rm -f /tmp/snyk-ads-installer-linux-x86_64
```
{% endtab %}

{% tab title="Windows" %}
From PowerShell:

```powershell
$ProgressPreference = 'SilentlyContinue'; Invoke-WebRequest -UseBasicParsing -Uri "https://downloads.snyk.io/ads/installer/latest/snyk-ads-installer-windows-x86_64.exe" -OutFile "$env:TEMP\snyk-ads-installer-windows-x86_64.exe"; & "$env:TEMP\snyk-ads-installer-windows-x86_64.exe" --tenant-id <TENANT_ID> --push-key <PUSH_KEY>
```
{% endtab %}
{% endtabs %}

{% hint style="warning" %}
On macOS and Linux, the command deletes the installer when it finishes. You need the installer again to uninstall, so either download it again or run the command without the final `rm -f` step. On Windows, the installer stays in `%TEMP%`.
{% endhint %}

### Differences from the package install

The command-line install always uses the latest published installer. It has no version selector, so you cannot pin a build this way. It also leaves nothing at a predictable path on macOS and Linux, which makes a later uninstall harder. Packages solve both problems, which is why they are the default.

## If you log in outside the United States

Add your region's base URL, as listed in [If you log in outside the United States](./#if-you-log-in-outside-the-united-states). For a package install, add `--base-url "<BASE_URL>"` on macOS, or `BASEURL=<BASE_URL>` before `/qn` on Windows. For a command-line install, add `--base-url "<BASE_URL>"` on every platform.

## What gets installed

Either method sets up the products for the signed-in user.

## Apply a configuration change

Publish your changes in **Settings**, then run the same command again. The installer fetches your current published configuration each time and skips anything already up to date, so you can safely repeat it.

## Upgrade the installer

After a package install, install the newer package over the existing one, then run the installer again. After a command-line install, run the same command again. It always fetches the latest installer.

## Remove ADS

For the `uninstall` command and its flags, visit [Remove an ADS product from a machine](./#remove-an-ads-product-from-a-machine).
