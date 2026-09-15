---
description: Configure NTLM to scan targets protected by NTLM v2 authentication.
---

# NTLM

NTLM (NT LAN Manager) authentication allows you to scan enterprise applications that use Windows-based authentication. This feature is designed for security engineers who need to scan internal applications protected by NTLM v2 authentication.

When you configure NTLM credentials on a Web target, Snyk API & Web automatically authenticates to your application during scans, enabling comprehensive vulnerability detection in authenticated areas.

## Prerequisites

* You must have the **change target settings** permission.
* You need valid NTLM credentials. A username and password are required; the domain and workstation are optional.
* Snyk supports NTLM for Web targets on the Enterprise and Trial plans. API targets do not support NTLM authentication.

## Configure the authentication settings

1. From the **Targets** page, locate your target and click the **gear icon** to access the target settings.
2. Select the **Authentication** tab and locate the **Basic Auth / NTLM** section.
3. Select the **NTLM** radio button.
4. Fill out the input fields as applicable:

* **Username** (required)
* **Password** (required)
* **Domain** (optional): Your Windows domain name (for example, `ACME`)
* **Workstation** (optional): Workstation name (for example, `WS01`)

6. Click **Save**.

{% hint style="info" %}
You can configure both Basic Auth and NTLM on the same target. The radio button controls which authentication method is active for scans. Selecting NTLM does NOT delete your Basic Auth configuration.
{% endhint %}

## Verify the configuration

After you save the configuration, NTLM authentication is enabled. The next scan against this target automatically uses the configured credentials.

## Manage the configuration

You can manage these settings anytime from your target's **Authentication** tab:

* To temporarily disable the setting, use the **Off/On** toggle
* To permanently remove the configuration, use the **Delete** button
