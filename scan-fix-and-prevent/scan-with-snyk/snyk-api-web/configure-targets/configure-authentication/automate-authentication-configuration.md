---
description: Configure authentication for web targets using AI-driven setup.
---

# Automate authentication configuration

Snyk API & Web uses a large language model (LLM) to automatically map a web target's login structure and generate a complete [login sequence](configure-login-sequence.md), including two-factor authentication (2FA). This entirely eliminates the need to manually record browser sessions or configure CSS selectors.

## Prerequisites

* Enable AI capabilities for your account. For details, visit [Enhancing your security with AI](../../managing-account/enhancing-your-security-with-ai.md).
* Ensure the web target uses a login sequence.

## Configure authentication automatically

In Snyk API & Web, navigate to **Target** > **Settings** > **Authentication**. In the **Login Sequence** section, click **Configure Authentication Automatically**.

{% stepper %}
{% step %}
### Enter credentials

Enter the authentication values in the corresponding fields:

* **Username or email**: the login identifier.
* **Password**: the account password.
* **OTP seed**: the time-based one-time password seed, when required.
* **Other login credentials**: any value the login flow needs, such as a company name or phone number.

{% hint style="info" %}
Add reusable credentials with the [Credential Manager](manage-credentials.md).
{% endhint %}
{% endstep %}

{% step %}
### Add instructions (optional)

For a non-standard or complex login flow, provide natural-language instructions to guide the LLM:

* Multiple roles: "Sign in as a student."
* Single sign-on: "Select SSO, then enter the company name."
* Regional server: "Select the EU server before signing in."

{% hint style="warning" %}
Do not enter credentials into the instructions field. Use the dedicated credential fields so Snyk does not send sensitive information to the LLM.
{% endhint %}
{% endstep %}
{% endstepper %}

## Validate generated output

LLM-generated output is sometimes incomplete or inaccurate, depending on application complexity. Validate the generated login sequence before you scan:

* Confirm that authentication succeeds. Navigate to [Test target configuration](../test-target-configuration.md).
* If the sequence is incorrect, add additional instructions or manually configure login steps. Navigate to [Configure the login sequence](configure-login-sequence.md).
