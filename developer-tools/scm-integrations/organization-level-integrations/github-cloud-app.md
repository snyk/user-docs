---
nav_context: classic
description: >-
  How to set up the Snyk GitHub Cloud App integration, including use with
  Universal Broker
---

# GitHub Cloud App

{% include "../../.gitbook/includes/new-navigation-banner.md" %}

## GitHub Cloud App

{% hint style="info" %}
**Feature availability**

If you are using Broker, visit the [GitHub Cloud App for Universal Broker](../scm-integrations-and-snyk-broker.md#github-cloud-app-for-universal-broker) documentation.

If you want to use your own GitHub App, contact your Snyk account team.
{% endhint %}

### Prerequisites for GitHub Cloud App

* Snyk Organization Admin user role.
* GitHub Organization Admin user role.
* A public, internal, or private GitHub repository.
* The required app permissions. For more information, visit [GitHub Cloud App permission requirements](../user-permissions-and-access-scopes.md#github-cloud-app-permission-requirements).

{% hint style="info" %}
Users can install the app on GitHub Organizations they are Repository Admins on through the GitHub UI.

The GitHub Cloud App integration is available at the Organization level within the Snyk Web UI.
{% endhint %}

### GitHub Cloud App benefits

The GitHub Cloud App improves on many features as compared to the current GitHub integration, including role-based, granular access control, increased API rate limits, and \
a foundation for an enhanced developer experience.

* Role-Based Access Control (RBAC) Compliance: The GitHub Cloud App decouples access control from user accounts, linking it to the app. This separation improves the management and enforcement of RBAC policies at the application level instead of individual accounts.
* Granular access control: The GitHub Cloud App allows fine-grained control over repository-level access permissions.
* The GitHub Cloud App provides a higher rate limit, enabling Snyk to make more requests and helping manage large-scale cases such as monorepos and organizations with many repositories.
* Enabler for an enhanced developer experience:
  * Pull request checks: The Checks tab in GitHub is only accessible through the Cloud App, providing an SCM-native experience for potential future PR check workflow improvements.
  * Fix and upgrade pull requests: The GitHub App performs Snyk-initiated pull requests directly rather than going through a service account.

### How to set up the GitHub Cloud App

{% hint style="warning" %}
When setting up the GitHub Cloud App, you can only implement one of the following scenarios:

* One GitHub organization connected to one Snyk Organization
* One GitHub organization connected to multiple Snyk Organizations
{% endhint %}

Log in to your Snyk account and navigate to the Integrations section in the Snyk Organization where you would like to set up the GitHub Cloud App.

Select the **GitHub Cloud App** tile.

<figure><img src="../../.gitbook/assets/01-github-cloud-app-tile-on-integrations-page.png" alt="GitHub Cloud App tile on the Integrations page"><figcaption><p>GitHub Cloud App tile on the Integrations page</p></figcaption></figure>

In the confirmation modal, click **Configure GitHub Cloud App**.

<figure><img src="../../.gitbook/assets/02-configuration-notice-for-github-cloud-app copy_2.png" alt="Configuration notice for the GitHub Cloud App" width="375"><figcaption><p>Configuration notice for the GitHub Cloud App</p></figcaption></figure>

The app then asks you to authorize it to act on your behalf, so it can check which GitHub organizations you can install it in.

<figure><img src="../../.gitbook/assets/03-user-authorization-for-the-app.png" alt="User authorization for the app" width="375"><figcaption><p>User authorization for the app</p></figcaption></figure>

When the install screen in GitHub opens, you can select the GitHub organization where you wish to install the app.

<figure><img src="../../.gitbook/assets/04-selection-of-github-organization-to-install-app-into.png" alt="Selection of the GitHub organization to install the app into" width="375"><figcaption><p>Selection of the GitHub organization to install the app into</p></figcaption></figure>

If the GitHub Cloud App is already installed in a GitHub organization, you can select that same GitHub organization during the integration process for a different Snyk Organization.

<figure><img src="../../.gitbook/assets/GitHub Cloud App Installation_22.png" alt="Connect another GitHub organization into a Snyk Organization" width="563"><figcaption><p>Connect another GitHub organization into a Snyk Organization</p></figcaption></figure>

Specify whether to install the app in all repositories in the selected GitHub organization or in only some of them. Then click **Install & Authorize**.

<figure><img src="../../.gitbook/assets/06-install-and-authorize-settings-for-github-org.png" alt="Install and authorize settings for the GitHub organization you are installing the GitHub Cloud App into" width="375"><figcaption><p>Install and authorize settings for the GitHub organization <br>you are installing the GitHub Cloud App into</p></figcaption></figure>

{% hint style="warning" %}
The GitHub Cloud App loses access to Snyk if you uninstall it from the GitHub organization or edit the repositories that the app instance can access.
{% endhint %}

## Configuring IP allowlists

{% hint style="info" %}
The GitHub Cloud App has no pre-configured IP addresses. Contact your Snyk account team for the Snyk IP addresses to add manually.
{% endhint %}

If an IP allowlist protects access to your GitHub Cloud or Enterprise Cloud repositories, add the Snyk IP addresses to enable communication with your SCM tool. IP allowlists at the Enterprise level override those at the Organization level in GitHub Enterprise Cloud, so you must add the Snyk IP addresses to both levels.

### Configuring the allowlist at the Organization level

{% hint style="warning" %}
You must be a GitHub Organization admin to perform this.
{% endhint %}

Follow GitHub's instructions for [adding an allowed IP address](https://docs.github.com/en/enterprise-cloud@latest/organizations/keeping-your-organization-secure/managing-security-settings-for-your-organization/managing-allowed-ip-addresses-for-your-organization#adding-an-allowed-ip-address), and enter the Snyk IP addresses in the form fields.

### Configuring the allowlist at the Enterprise level

{% hint style="warning" %}
You must be a GitHub Enterprise admin to perform this.
{% endhint %}

Follow GitHub's instructions for [adding an allowed IP address](https://docs.github.com/en/enterprise-cloud@latest/organizations/keeping-your-organization-secure/managing-security-settings-for-your-organization/managing-allowed-ip-addresses-for-your-organization#adding-an-allowed-ip-address), and enter the Snyk IP addresses in the form fields.

## How to migrate to the GitHub Cloud App

If you are an Enterprise plan customer, you can migrate Snyk Targets to the GitHub Cloud App using the [snyk-migrate-to-github-app](https://github.com/snyk-labs/snyk-migrate-to-github-app) tool in the [tool repository](https://github.com/snyk-labs/snyk-migrate-to-github-app).
