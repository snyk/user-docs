# GitHub Server App

{% hint style="info" %}
**Feature availability**

The GitHub Server App is available only for Enterprise plans.

This feature supports self-hosted instances of GitHub and Snyk [Universal Broker](../../enterprise-setup/snyk-broker/universal-broker/). For instructions, see [GitHub Server App for Universal Broker](../scm-integrations-and-snyk-broker.md#github-server-app-for-universal-broker).
{% endhint %}

### Prerequisites for the GitHub Server App

* A self-hosted instance of GitHub
* Snyk Organization Admin user role
* GitHub organization Admin user role
* A public or private GitHub repository

### GitHub Server App benefits <a href="#github-server-app-benefits" id="github-server-app-benefits"></a>

The Snyk GitHub Server App improves on many features compared to the Snyk GitHub Enterprise integration, including role-based granular access control, increased API rate limits, and the creation of an entry point for expanded and enhanced developer experiences.

* RBAC (Role-Based Access Control) Compliance:&#x20;
  * With the GitHub Server App, the access control mechanism is decoupled from individual user accounts. Instead, it is associated with the app entity itself. This separation allows for better management and enforcement of RBAC policies, as access control is handled at the application level rather than being tied to individual user accounts.
* Granular access control:&#x20;
  * The GitHub Server App allows for fine-grained control over access permissions at the repository level.
* Increased API rate limit:&#x20;
  * The GitHub Server App provides higher rate limits, allowing Snyk to make a larger number of API requests. This increased limit will assist in handling large-scale use cases, such as monorepos with a large number of Projects, GitHub organizations with a large number of repositories, and more.
* Enabler for an enhanced developer experience:
  * Pull request checks: The Checks tab experience in GitHub is exclusively accessible through the GitHub Cloud App, enabling an SCM native experience as part of potential future PR check workflow improvements.
  * Fix and upgrade pull requests: Pull requests initiated by Snyk are performed directly by the GitHub App rather than a service account.

### Set up the GitHub Server App

{% hint style="warning" %}
When setting up the GitHub Server App, you can implement only one of the following scenarios:

* One GitHub organization connected to one Snyk Organization
* One GitHub organization connected to multiple Snyk Organizations
{% endhint %}

In the Snyk UI navigate to the integrations page and select the **GitHub Server App** tile.

<figure><img src="../../.gitbook/assets/image (594).png" alt=""><figcaption><p>GitHub Server App tile highlighted in the Snyk UI</p></figcaption></figure>

Clicking on the tile opens a modal that allows you to enter the URL of your GitHub Server. Entering the URL of your GitHub Server instance will redirect you to your GitHub instance, where you will be able to create the app.

<figure><img src="../../.gitbook/assets/image (596).png" alt=""><figcaption><p>Integration model prompting you for your GitHub Server's URL</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (598).png" alt=""><figcaption><p>Registration of the app on your GitHub instance</p></figcaption></figure>

You are then asked to authorize the app to act on your users' behalf. The app uses this information to check which GitHub organizations you are authorized to install the app in.

<figure><img src="../../.gitbook/assets/image (599).png" alt=""><figcaption><p>User authorization for the app</p></figcaption></figure>

When the install screen in GitHub opens, you can select the GitHub organization where you wish to install the app.

<figure><img src="../../.gitbook/assets/image (602).png" alt=""><figcaption><p>Selection of the GitHub organization to install the app into</p></figcaption></figure>

If the GitHub Server App is already installed in a GitHub organization on your GitHub instance, you can select that same GitHub organization during the integration process for a different Snyk Organization.

<figure><img src="../../.gitbook/assets/image (604).png" alt=""><figcaption><p>Connect another GitHub organization into a Snyk Organization</p></figcaption></figure>

Specify whether you wish to install the app in all or a select number of the repositories belonging to the selected GitHub organization, then click **Install & Authorize**.

<figure><img src="../../.gitbook/assets/image (656).png" alt=""><figcaption><p>Install and authorize settings for the GitHub organization you are installing the GitHub Cloud App into</p></figcaption></figure>

{% hint style="danger" %}
The GitHub Server App will lose access to Snyk if it is uninstalled from the GitHub organization. If this happens, you can create a fresh integration in Snyk to regain access.
{% endhint %}

### Migrate from an existing GitHub Enterprise integration

If you are an Enterprise plan customer, you can migrate Snyk Targets to the GitHub Server App using the [snyk-migrate-to-github-app](https://github.com/snyk-labs/snyk-migrate-to-github-app) tool in the [tool repository](https://github.com/snyk-labs/snyk-migrate-to-github-app).

### How to disconnect a non-brokered GitHub Server App integration

{% hint style="warning" %}
Disconnecting the Snyk GitHub Server App integration halts all scans for imported repositories. PR checks cannot be executed and Projects are deactivated in the Snyk Web UI.

Note that the GitHub App will remain listed on your GitHub organization until removed manually.
{% endhint %}

1. Navigate to the Snyk GitHub Server App integration **Settings**.
2. At the bottom of the page, select **Remove GitHub Server App.**
3. When the confirmation modal opens, select **Disconnect GitHub Server App**.

<figure><img src="../../.gitbook/assets/image (689).png" alt="" width="375"><figcaption><p>Confirm disconnecting from GitHub Server App</p></figcaption></figure>

After the integration is disconnected, imported Snyk Projects will be set to inactive, and you will no longer get alerts, pull requests, or Snyk tests on pull requests.

You can re-connect anytime; however, re-initiating the Snyk Projects for monitoring requires setting up the integration again.

### How to disconnect a brokered GitHub Server App integration

{% hint style="warning" %}
Disconnecting the Snyk GitHub Server App integration halts all scans for imported repositories. PR checks cannot be executed and Projects are deactivated in the Snyk Web UI.

Note that the GitHub App will remain listed on your GitHub organization until the app is removed manually.
{% endhint %}

Run `snyk-broker-config workflows connections disconnect`and select the connection you want to disconnect.

After the integration is disconnected, imported Snyk Projects will be set to inactive, and you will no longer get alerts, pull requests, or Snyk tests on pull requests.

You can re-connect anytime; however, re-initiating the Snyk Projects for monitoring requires setting up the integration again.
