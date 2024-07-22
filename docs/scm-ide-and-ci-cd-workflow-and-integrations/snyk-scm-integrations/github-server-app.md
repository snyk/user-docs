# GitHub Server App

{% hint style="warning" %}
**Release status and feature availability**

The GitHub Server App is in [Closed Beta](../../getting-started/snyk-release-process.md). This feature must be enabled using a feature flag on the Group or Organization level. Contact [Snyk support](https://support.snyk.io/hc/en-us/requests/new) if you are interested in getting access.

The GitHub Server App is available to Snyk Enterprise plan customers. If you have a Legacy Business plan, contact [Snyk support](https://support.snyk.io/hc/en-us) for access. See the [Plans and pricing](https://snyk.io/plans/) page for details.

This feature is not yet supported for Snyk Broker.\
As Snyk does not have static IP addresses, this integration will not work with IP Whitelisting in GitHub.
{% endhint %}

When you want to add new integrations to your Snyk account you need to first decide the level type at which you want to install the integration.

* [Group level ](github-server-app.md#group-level-snyk-apprisk-integrations)- Add integrations to your Snyk application that will be available for your Snyk AppRisk Essentials or Snyk AppRisk Pro.&#x20;
* [Organization level](github-server-app.md#organization-level-snyk-integrations) - Add integrations for your Snyk application that will be available for all Snyk products, except Snyk AppRisk.

{% hint style="info" %}
If you want to set up integrations for Snyk AppRisk, use the Integrations menu at the Group level.
{% endhint %}

## Organization level - Snyk integrations

### Prerequisites for GitHub Server App

* A self-hosted instance of GitHub.
* Snyk Organization Admin user role.
* GitHub Organization Admin user role.
* A public or private GitHub repository.

{% hint style="info" %}
Users can install the app on GitHub Organizations they are Repository Admins on through the GitHub UI.
{% endhint %}

### GitHub Server App benefits

The Snyk GitHub Server App improves on many features compared to the Snyk GitHub Enterprise integration, including role-based granular access control, increased API rate limits, and the creation of an entry point for expanded and enhanced developer experiences.

* **RBAC (Role-Based Access Control) Compliance**: With the GitHub Server App, the access control mechanism is decoupled from individual user accounts. Instead, it is associated with the app entity itself. This separation allows for better management and enforcement of RBAC policies, as access control is handled at the application level rather than being tied to individual user accounts.
* **Granular access control**: The GitHub Server App allows for fine-grained control over access permissions at the repository level.&#x20;
* **Increased API rate limit**: The GitHub Server App provides higher rate limits, allowing Snyk to make a larger number of API requests. This increased limit will assist in handling large-scale use cases, such as monorepos with a large number of Projects, GitHub organizations with a large number of repositories, and more.
* **Enabler for an enhanced developer experience:**
  * Pull request checks: The Checks tab experience in GitHub is exclusively accessible through the GitHub Cloud App, enabling an SCM native experience as part of potential future PR check workflow improvements.
  * Fix and upgrade pull requests: Pull requests initiated by Snyk are performed directly by the GitHub App rather than a service account.

### How to set up the GitHub Server App

Log in to your Snyk account and navigate to the Integrations section in the Snyk Organization where you would like to set up the GitHub Server App.

Select the **GitHub Server App** tile.

<figure><img src="../../.gitbook/assets/Screenshot 2024-06-05 at 11.12.40.png" alt="Select the GitHub Server App tile in the Snyk Web UI"><figcaption><p>Select the GitHub Server App tile in the Snyk Web UI</p></figcaption></figure>

In the confirmation modal, select **Configure GitHub Server App.**\


<figure><img src="https://lh7-us.googleusercontent.com/docsz/AD_4nXceIkKngJie6IJ9QAdoTbHkJoNOFuJtzItSyvgawGixNoqJf_ITi4lgzM6UCO5DVLuN83-Ry7z9iagPUDzsplCt5MyZzREvrVYox0OoCep2me8gy4w4C4NUT_hv-HiDJSacXdxWpKqqXCghIo78WRydL0t5?key=zABDbDwEFVygIY1cldcfgQ" alt="Confirm configuration of the GitHub Server App" width="375"><figcaption><p>Confirm configuration of the GitHub Server App</p></figcaption></figure>

You'll be redirected to your GitHub instance in order to register the app. You can choose the name of the GitHub App that will be registered on your instance.

<figure><img src="../../.gitbook/assets/Screenshot 2024-06-19 at 10.02.55.png" alt="Registration of the app on your GitHub instance"><figcaption><p>Registration of the app on your GitHub instance</p></figcaption></figure>

You are then asked to authorize the app to act on your user’s behalf. The app uses this information to check which GitHub organizations you are authorized to install the app in.

<figure><img src="../../.gitbook/assets/Screenshot 2024-06-19 at 10.03.17 (1).png" alt="User authorization for the app" width="563"><figcaption><p>User authorization for the app</p></figcaption></figure>

When the install screen in GitHub opens, you can select the GitHub organization where you wish to install the app.

<figure><img src="../../.gitbook/assets/Screenshot 2024-06-19 at 10.03.39.png" alt="Selection of the GitHub organization to install the app into"><figcaption><p>Selection of the GitHub organization to install the app into</p></figcaption></figure>

If the GitHub Server App is already installed in a GitHub organization on your GitHub instance, you can select that same GitHub organization during the integration process for a different Snyk Organization.

<figure><img src="../../.gitbook/assets/Screenshot 2024-06-19 at 10.05.05.png" alt="Connect another GitHub organization into a Snyk Organization"><figcaption><p>Connect another GitHub organization into a Snyk Organization</p></figcaption></figure>

Specify whether you wish to install the app in all or a select number of the repositories belonging to the selected GitHub organization, then click **Install & Authorize**.

<figure><img src="../../.gitbook/assets/Screenshot 2024-06-19 at 10.04.09.png" alt="Install and authorize settings for the GitHub organization you are installing the GitHub Cloud App into" width="563"><figcaption><p>Install and authorize settings for the GitHub organization you are installing the GitHub Cloud App into</p></figcaption></figure>

{% hint style="warning" %}
The GitHub Server App will lose access to Snyk if it is uninstalled from the GitHub organization or if the repositories to which the app instance has access are edited.
{% endhint %}

### How to migrate to the GitHub Server App

If you are an Enterprise plan customer, you can migrate Snyk Targets to the GitHub Server App using the [snyk-migrate-to-github-app](https://github.com/snyk-labs/snyk-migrate-to-github-app) tool in the [tool repository](https://github.com/snyk-labs/snyk-migrate-to-github-app).

### Feedback on the GitHub Server App

Because this feature is in [Closed Beta](../../getting-started/snyk-release-process.md), the functionality will likely evolve based on your feedback. If you would like to provide any feedback, contact your Account Manager or [Snyk support](https://support.snyk.io/hc/en-us/requests/new).

## Group level - Snyk AppRisk integrations

Navigate to the [GitHub setup guide for Snyk AppRisk ](github-enterprise.md#github-setup-guide-for-snyk-apprisk)for all details on how to set up the GitHub integration for Snyk AppRisk.

