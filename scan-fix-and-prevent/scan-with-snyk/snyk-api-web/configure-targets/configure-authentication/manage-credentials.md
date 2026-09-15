---
description: How to manage credentials for Snyk API and Web targets
nav_context: classic
---

# Manage credentials

Manage and rotate authentication credentials securely across your Snyk API & Web targets.

The Credentials Manager is a centralized hub for handling target authentication data across your account. Instead of re-entering sensitive information for every target, you can manage credentials in one secure location.

Using the Credentials Manager simplifies the password rotation process and improves security for sensitive credentials.

## Add a credential

1. Navigate to the **Credentials** section in your account settings.

<figure><img src="../../../../.gitbook/assets/manage-credentials-section.png" alt="Credentials section in account settings"><figcaption></figcaption></figure>

2. Click **Add Credential**.
3. Configure the credential:
   * **Sensitive**: Select this option to permanently hide credential values from all users. You can update sensitive credential values but never view them after saving.
   * **Scope**: Choose whether the credential is available account-wide or restricted to specific teams. Everyone in the account can use account-wide credentials.
   * **Description** (optional): Add relevant information to help your team understand when to use this credential.
4. Click **Save**.

<figure><img src="../../../../.gitbook/assets/manage-credentials-add-form.png" alt="Add credential form with sensitivity, scope, and description fields"><figcaption></figcaption></figure>

## Link a credential to a target

After creating a credential, link it to any target:

1. Navigate to the **Targets** page.
2. Locate your target and click the **gear icon** to access the target settings.
3. In the authentication or header configuration section, select the option to **Link from Credentials Manager**.
4. Choose the appropriate credential from the dropdown list.

<figure><img src="../../../../.gitbook/assets/manage-credentials-link-dropdown.png" alt="Credential dropdown showing available credentials to link"><figcaption></figcaption></figure>

## Where you can use credentials

You can securely store sensitive information in several areas across Snyk API & Web. Look for the **Add Credential** icon in the following locations:

**Target Settings > Authentication:**

* Login Form fields
* Custom Variables for Login Sequences
* Authentication Payloads
* Static Headers or Cookies
* Basic Auth username and password

**Target Settings > Scanner:**

* Custom Headers and Cookies
* API Parameter Custom Values
* Postman Environment Values

**Target Settings > Extra Hosts:**

* Custom Headers and Cookies

## Permissions and scope

Access to Credentials depends on your user role and assigned permissions:

* Update Target Configuration: Users with this permission can create credentials and use them within their specific targets.
* Manage Credentials: Users with this permission can create, view, update, and delete credentials across the account or team, even if someone else created the credential.
* Scoped Credentials: You can scope credentials to the entire account or restrict them to users from specific teams.

For more information about permissions, visit the Snyk API & Web permissions documentation.

## Transitioning from obfuscated values

The Credentials Manager replaces the Secret Obfuscation feature:

* Obfuscation toggle: The option to turn obfuscation on or off for account owners is hidden. Centralized management through the Credentials Manager is now the standard for sensitive data.
* Existing configurations: All existing configurations continue to work. You can keep them or replace them with Credentials Manager credentials. Snyk recommends replacing them.
* Obfuscated values: You cannot retrieve previously obfuscated values. To create new obfuscated values, use Sensitive Credentials in the Credentials Manager.
