---
description: >-
  How to grant the Agentic Development Security scan Full Disk Access on managed
  Macs with a configuration profile in Jamf Pro or Microsoft Intune
last_reviewed: 2026-11-03
last_reviewed_source: reviewed
---

# Grant Full Disk Access on macOS

macOS protects some locations from background processes, and shows a permission pop-up when a process tries to read them. To let the Agentic Development Security (ADS) Machines (Agent Supply Chain Security) scan read these locations, deploy a configuration profile that grants Full Disk Access to the scan binary. Snyk provides the profile as a `.mobileconfig` file, and you deploy it through your mobile device management (MDM) tool.

The profile grants access to the scan binary in `/Library/Application Support/Snyk/ads-scan/bin/`, the directory that [managed mode](./#installation-mode) uses. The binary name depends on the processor, so each processor architecture has its own profile.

{% hint style="info" %}
This page applies to managed mode only. If you install ADS in user mode, you do not need a Full Disk Access profile.
{% endhint %}

## Prerequisites

You need the following:

* Macs enrolled in your MDM tool, with ADS installed in managed mode
* The profile for each processor architecture you manage:

| Processor     | Scan binary                | Profile                                                                                                       |
| ------------- | -------------------------- | ------------------------------------------------------------------------------------------------------------- |
| Apple silicon | `snyk-scan-macos-arm64`    | [Snyk-Agent-Scan-FDA-arm64.mobileconfig](../../../.gitbook/assets/Snyk-Agent-Scan-FDA-arm64.mobileconfig)     |
| Intel         | `snyk-scan-macos-x86_64`   | [Snyk-Agent-Scan-FDA-x86_64.mobileconfig](../../../.gitbook/assets/Snyk-Agent-Scan-FDA-x86_64.mobileconfig)   |

## Deploy the profile

Create one profile per processor architecture, and scope each profile to matching Macs only. An Intel Mac does not use the Apple silicon profile, and the reverse.

{% tabs %}
{% tab title="Jamf Pro" %}
1. Log in to Jamf Pro.
2. Navigate to **Computers** > **Configuration Profiles**.
3. Click **Upload**, and select the profile for your processor architecture, for example `Snyk-Agent-Scan-FDA-arm64.mobileconfig`.
4. In **General**, confirm that **Level** is **Computer Level** and that **Distribution Method** is **Install Automatically**.
5. Open **Scope**, and add the Macs that match the profile's architecture.
6. Click **Save**.
7. Repeat the steps for the other architecture.
{% endtab %}

{% tab title="Microsoft Intune" %}
{% hint style="warning" %}
**Deployment channel** must be **Device channel**. The profile applies to the whole Mac, not to one user. You cannot change the channel after you save the profile. To use another channel, create a new profile.
{% endhint %}

1. Log in to the Microsoft Intune admin center.
2. Navigate to **Devices** > **Manage devices** > **Configuration**, then click **Create** > **New policy**.
3. Set **Platform** to **macOS** and **Profile type** to **Templates**. Select the **Custom** template, then click **Create**.
4. In **Basics**, enter a **Name**, for example `Snyk Agent Scan - Full Disk Access (Apple silicon)` or `Snyk Agent Scan - Full Disk Access (Intel)`. Optionally, enter a **Description**. Click **Next**.
5.  In **Configuration settings**, set:
    * **Configuration profile name**: the name that devices and Intune status show, for example the same value as **Name**
    * **Deployment channel**: **Device channel**
    * **Configuration profile file**: select the profile for your processor architecture, for example `Snyk-Agent-Scan-FDA-arm64.mobileconfig`

    Click **Next**.
6. Optional: In **Scope tags**, assign a scope tag, then click **Next**.
7. In **Assignments**, assign the profile to a device group that contains the Macs that match the profile's architecture, then click **Next**.
8. In **Review + create**, review your settings, then click **Create**.
9. Repeat the steps for the other architecture.

Intune applies the profile at the next check-in of each assigned Mac. To apply it sooner, click **Sync** on the device page.
{% endtab %}
{% endtabs %}

## Verify the profile

To confirm that the profile works, visit [Verify an installation](verify-an-installation.md).

If the profile does not match the Mac's processor architecture, macOS does not apply it and the Full Disk Access pop-up still appears.
