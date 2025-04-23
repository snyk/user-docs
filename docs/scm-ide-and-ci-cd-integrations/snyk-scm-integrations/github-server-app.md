# GitHub Server App

When you want to add new integrations to your Snyk account, you must first decide the level at which you want to install the integration.

* [Group level ](github-cloud-app.md#group-level-snyk-essentials-integrations)- Add integrations to your Snyk application that will be available for your Snyk Essentials. If you want to set up integrations for Snyk Essentials, use the Integrations menu at the Group level.
* [Organization level](github-server-app.md#organization-level-snyk-integration) - Add integrations for your Snyk application that will be available for all Snyk products, except Snyk Essentials or Snyk AppRisk.

## Organization level - Snyk integration

{% hint style="info" %}
**Feature availability**

The GitHub Server App is available only for Enterprise plans.

This feature supports self-hosted instances of GitHub and Snyk [Universal Broker](../../enterprise-setup/snyk-broker/universal-broker/).
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

## Set up the GitHub Server App for Universal Broker

If your GithHub server is not publicly available, you can provide access to it through the Universal Broker, a proxy deployed in your internal network to facilitate outbound connections and communication with Snyk.

The setup process for Universal Broker involves:

1. [Creating a GitHub App on your GitHub Server instance](github-server-app.md#create-a-github-app-for-universal-broker)
2. [Creating the Universal Broker connection using the `snyk-broker-config` command](github-server-app.md#create-the-universal-broker-connection-for-your-github-server-app)

### Create a GitHub App for Universal Broker

To use the GitHub Server App with Universal Broker you must create your own GitHub App on your GitHub Server instance. You can do this by using the predefined `GITHUB-SERVER-URL` that follows and includes all the required permissions for Snyk services:

```
{{GITHUB-SERVER-URL}}/settings/apps/new?name=Snyk&description=Snyk%20helps%20you%20develop%20fast%20while%20staying%20secure%20by%20finding%20and%20automatically%20fixing%20security%20issues%20in%20your%20code%2C%20open%20source%20dependencies%2C%20containers%2C%20and%20infrastructure%20as%20code%20-%20all%20powered%20by%20Snyk%E2%80%99s%20security%20intelligence.&url=https%3A%2F%2Fgithub.com%2Fapps%2Fsnyk-io&public=false&webhook_active=true&webhook_url={{SNYK-ENV}}%2Fapi%2Fhidden%2Fscm-apps%2Fapi%2Fgithub-app%2Fwebhook&checks=write&statuses=write&contents=write&metadata=read&pull_requests=write&repository_hooks=write&members=read&events[]=repository 
```

Replace the following in the URL above:

* `{{GITHUB-SERVER-URL}}`: Replace this with the base URL of your GitHub Server instance.
* `{{SNYK-ENV}}`: Replace this with the tenacy of your Snyk account. This value needs to be URL encoded; the most common are listed below:
  * Snyk US-01: https%3A%2F%2Fapp.snyk.io
  * Snyk EU: https%3A%2F%2Fapp.eu.snyk.io
  * Snyk AU: https%3A%2F%2Fapp.au.snyk.io
  * Snyk US-02: https%3A%2F%2Fapp.us.snyk.io

After the values are replaced, navigate to that URL. This will take you to the app creation screen in your GitHub Server instance with all the required details pre-filled. Afterward, scroll to the end of the page and click on the **Create GitHub App** button, ensuring that the **Any account** radio button is selected.

<figure><img src="../../.gitbook/assets/image (611).png" alt=""><figcaption><p>Create GitHub App account setting selection</p></figcaption></figure>

{% hint style="warning" %}
On creation of your GitHub Server App you will see a`ClientId` and `AppId` - store these safely as these are your app's credentials and should be treated as secrets.
{% endhint %}

After creating your GitHub Server App you will see a banner at the top of the page prompting you to create a private key. Click on it and create a private key for your app.

<figure><img src="../../.gitbook/assets/image (618).png" alt=""><figcaption><p>Registration success message with a link for private key generation</p></figcaption></figure>

Generating a private key will initiate a download of a `.pem` file; this should be treated as a secret and kept safe.

Your GitHub Server App is now ready to be installed against repositories in your Snyk Organization. You can scroll to the top of the same page and select **Install App** from the left-hand panel. Choose the newly created app and click on the **Install** button to the right.

<figure><img src="../../.gitbook/assets/image (624).png" alt=""><figcaption><p>Install the GitHub App in your repositories</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (658).png" alt="" width="563"><figcaption><p>Install the GitHub App in your selected repositories</p></figcaption></figure>

Choose where you want to install the app in your GitHub organization. It can be installed against specific repositories, or all of them in your GitHub organization.

{% hint style="info" %}
If you select to install the app on a subset of repositories in your GitHub organization, the app will work only in those repositories. You can edit where the app is installed by returning to this screen at a later date if you want to add it to additional repositories
{% endhint %}

On installation of the app you will be assigned an `InstallationID`. These are the final numbers in the URL of the page. Make a note of that number, as you will need it to set up a Universal Broker connection.

### Create the Universal Broker connection for your GitHub Server App

Before the GitHub Server App can be used with the Universal Broker, you must create a connection of the `github-server-app` type using the `snyk-broker-config` tool. For more details, see the  [Universal Broker](../../enterprise-setup/snyk-broker/universal-broker/) documentation. After the connection is created, it can be integrated with one or more Organization(s) of your choice.

To set up the integration you will need the following:

* The base API address for your Snyk region; refer to the list of [API URLs](../../working-with-snyk/regional-hosting-and-data-residency.md#api-urls) for Snyk regional hosting.
* The Snyk Organization ID you want to set up the integration in.
* A valid [Snyk API token](../../snyk-api/authentication-for-api/#how-to-obtain-your-personal-token).
* `snyk-broker-config` tool installed
* Tenant Admin role

When you have these prerequisites, use the following commands

* `snyk-broker-config workflows connections create,`choosing the `github-server-app` option and providing the information you are prompted for in the workflow.
* `snyk-broker-config workflows connections integrate`to integrate the newly created connection to the Organization of your choice, pasting the org ID when you are prompted.

Visit the integrations page in Snyk to see that the integration has been configured.&#x20;

<figure><img src="../../.gitbook/assets/image (630).png" alt=""><figcaption><p>A successful GitHub Server App integration</p></figcaption></figure>

See the [Universal Broker](../../enterprise-setup/snyk-broker/universal-broker/) documentation for more details.

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
