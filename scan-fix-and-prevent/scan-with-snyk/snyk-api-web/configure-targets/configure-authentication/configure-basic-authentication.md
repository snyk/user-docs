---
nav_context: classic
description: How to configure basic authentication for Snyk API and Web targets
---

# Basic authentication

Configure basic authentication to scan targets protected by HTTP Basic Access Authentication.

Basic authentication is an authentication scheme built into the HTTP protocol. When you access a protected resource, the browser prompts you to log in.

This authentication process differs from your application's own authentication system and from form-based authentication methods. Basic authentication sends credentials in the HTTP header rather than through form submission.

## Set up basic authentication

1. From the **Targets** page, locate your target and click the **gear icon** to access the target settings.
2. Select the **Authentication** tab.
3. Scroll down to the **Basic Auth** section (for API targets) or to the **Basic Auth / NTLM** section (for Web targets).
4. Select the **Basic Auth** radio button as needed (for Web targets only).
5. Enter your credentials (username and password).
6. Click **Save and enable**.

{% hint style="info" %}
You can configure both Basic Auth and NTLM on the same target. The radio button controls which authentication method is active for scans. Selecting Basic Auth does NOT delete your NTLM configuration.
{% endhint %}

## Verify the configuration

After you save the configuration, Basic Auth is enabled. The next scan against this target automatically uses the configured credentials.

## Manage the configuration

You can manage these settings anytime from your target's **Authentication** tab:

* To temporarily disable the setting, use the **Off/On** toggle
* To permanently remove the configuration, use the **Delete** button
