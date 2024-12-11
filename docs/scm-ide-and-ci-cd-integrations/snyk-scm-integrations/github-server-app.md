# GitHub Server App

When you want to add new integrations to your Snyk account you need to first decide the level type at which you want to install the integration.

* [Group level ](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-cloud-app#group-level-snyk-apprisk-integrations)- Add integrations to your Snyk application that will be available for your Snyk AppRisk Essentials or Snyk AppRisk Pro. If you want to set up integrations for Snyk AppRisk, use the Integrations menu at the Group level.
* [Organization level](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-cloud-app#organization-level-snyk-integrations) - Add integrations for your Snyk application that will be available for all Snyk products, except Snyk AppRisk.

## Organization level - Snyk integrations

{% hint style="info" %}
**Release status**

The GitHub Server App is in Early Access and available only for Enterprise plans.

This feature supports self-hosted instances of GitHub and Snyk's [Universal Broker](../../enterprise-setup/snyk-broker/universal-broker/).
{% endhint %}

This page covers the following features:

* [GitHub Server App](github-server-app.md#prerequisites-for-github-server-app).
* [Migrate from an existing GitHub Enterprise integration](github-server-app.md#migrate-from-an-existing-github-enterprise-integration).
* [GitHub Server App for Universal Broker](github-server-app.md#set-up-the-github-server-app-for-universal-broker).

### Prerequisites for GitHub Server App

* A self-hosted instance of GitHub.
* Snyk Organization Admin user role.
* GitHub Organization Admin user role.
* A public or private GitHub repository.

### GitHub Server App benefits <a href="#github-server-app-benefits" id="github-server-app-benefits"></a>

The Snyk GitHub Server App improves on many features compared to the Snyk GitHub Enterprise integration, including role-based granular access control, increased API rate limits, and the creation of an entry point for expanded and enhanced developer experiences.

* **RBAC (Role-Based Access Control) Compliance**: With the GitHub Server App, the access control mechanism is decoupled from individual user accounts. Instead, it is associated with the app entity itself. This separation allows for better management and enforcement of RBAC policies, as access control is handled at the application level rather than being tied to individual user accounts.
* **Granular access control**: The GitHub Server App allows for fine-grained control over access permissions at the repository level.
* **Increased API rate limit**: The GitHub Server App provides higher rate limits, allowing Snyk to make a larger number of API requests. This increased limit will assist in handling large-scale use cases, such as monorepos with a large number of Projects, GitHub organizations with a large number of repositories, and more.
* **Enabler for an enhanced developer experience:**
  * Pull request checks: The Checks tab experience in GitHub is exclusively accessible through the GitHub Cloud App, enabling an SCM native experience as part of potential future PR check workflow improvements.
  * Fix and upgrade pull requests: Pull requests initiated by Snyk are performed directly by the GitHub App rather than a service account.

### Set up the GitHub Server App

{% hint style="warning" %}
When setting up the GitHub Server App, you can only implement one of the following scenarios:

* One GitHub organization connected to one Snyk Organization
* One GitHub organization connected to multiple Snyk Organizations
{% endhint %}

In Snyk's UI navigate to the integrations page and select the **GitHub Server App** tile.

<figure><img src="../../.gitbook/assets/image (594).png" alt=""><figcaption><p>Image with the GitHub Server App tile highlighted in Snyk's UI</p></figcaption></figure>

Clicking on the tile opens a modal that allows you to enter the URL of your GitHub Server. Entering the URL of your GitHub Server instance will then redirect you to your GitHub instance, where you'll be able to create the app.

<figure><img src="../../.gitbook/assets/image (596).png" alt=""><figcaption><p>Integration model prompting the user for their GitHub Server's URL</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (598).png" alt=""><figcaption><p>Registration of the app on your GitHub instance</p></figcaption></figure>

You are then asked to authorize the app to act on your user’s behalf. The app uses this information to check which GitHub organizations you are authorized to install the app in.

<figure><img src="../../.gitbook/assets/image (599).png" alt=""><figcaption><p>User authorization for the app</p></figcaption></figure>

When the install screen in GitHub opens, you can select the GitHub organization where you wish to install the app.

<figure><img src="../../.gitbook/assets/image (602).png" alt=""><figcaption><p>Selection of the GitHub organization to install the app into</p></figcaption></figure>

If the GitHub Server App is already installed in a GitHub organization on your GitHub instance, you can select that same GitHub organization during the integration process for a different Snyk Organization.

<figure><img src="../../.gitbook/assets/image (604).png" alt=""><figcaption><p>Connect another GitHub organization into a Snyk Organization</p></figcaption></figure>

Specify whether you wish to install the app in all or a select number of the repositories belonging to the selected GitHub organization, then click **Install & Authorize**.

<figure><img src="../../.gitbook/assets/image (608).png" alt=""><figcaption><p>Install and authorize settings for the GitHub organization you are installing the GitHub Cloud App into</p></figcaption></figure>

{% hint style="danger" %}
The GitHub Server App will lose access to Snyk if it is uninstalled from the GitHub organization. Should this happen you will be required to raise a support ticket before re-integrating with Snyk's GitHub Server App.&#x20;
{% endhint %}

### Migrate from an existing GitHub Enterprise integration

If you are an Enterprise plan customer, you can migrate Snyk Targets to the GitHub Server App using the [snyk-migrate-to-github-app](https://github.com/snyk-labs/snyk-migrate-to-github-app) tool in the [tool repository](https://github.com/snyk-labs/snyk-migrate-to-github-app).

### Set up the GitHub Server App for Universal Broker

The setup process for Universal Broker involves:

1. [Creating a GitHub App on your GitHub Server instance](github-server-app.md#create-a-github-app-for-universal-broker).
2. [Creating a base integration for the broker to use, created using the API](github-server-app.md#create-the-github-server-app-scm-integration).

#### Create a GitHub App for Universal Broker

To use the GitHub Server App with Universal Broker you need to create your own GitHub App on your GitHub Server instance. You can do this by using the `GITHUB-SERVER-URL` that pre-defines all the required permissions for Snyk services:

```
{{GITHUB-SERVER-URL}}/settings/apps/new?name=Snyk&description=Snyk%20helps%20you%20develop%20fast%20while%20staying%20secure%20by%20finding%20and%20automatically%20fixing%20security%20issues%20in%20your%20code%2C%20open%20source%20dependencies%2C%20containers%2C%20and%20infrastructure%20as%20code%20-%20all%20powered%20by%20Snyk%E2%80%99s%20security%20intelligence.&url=https%3A%2F%2Fgithub.com%2Fapps%2Fsnyk-io&public=false&webhook_active=true&webhook_url={{SNYK-ENV}}%2Fapi%2Fhidden%2Fscm-apps%2Fapi%2Fgithub-app%2Fwebhook&checks=write&statuses=write&contents=write&metadata=read&repository_projects=write&pull_requests=write&repository_hooks=write&members=read&events[]=repository 
```

Replace the following in the URL above:

* `{{GITHUB-SERVER-URL}}`: Replace this with the base URL of your GitHub Server instance.
* `{{SNYK-ENV}}`: Replace this with the tenacy of your Snyk account. This value needs to be URL encoded; the most common are listed below:
  * Snyk US: https%3A%2F%2Fapp.snyk.io
  * Snyk EU: https%3A%2F%2Fapp.eu.snyk.io
  * Snyk AU: https%3A%2F%2Fapp.au.snyk.io
  * Snyk AWS US: https%3A%2F%2Fapp.us.snyk.io

After the values are replaced navigate to that URL. This will take you to the app creation screen in your GitHub Server instance with all the required details pre-filled. Afterwards scroll to the end of the page and click on the **Create GitHub App** button, ensuring that the **Any account** radio button is selected.

<figure><img src="../../.gitbook/assets/image (611).png" alt=""><figcaption><p>Create GitHub App account setting selection</p></figcaption></figure>

{% hint style="warning" %}
On creation of your GitHub Server App you will be presented with a`ClientId` and `AppId` - store these safely as these are your app's credentials and should be treated as secrets.
{% endhint %}

After creating your GitHub Server App you'll be presented with a banner prompting you to create a private key. Please click on it and create a private key for you app.

<figure><img src="../../.gitbook/assets/image (618).png" alt=""><figcaption><p>Registration success message with a link for private key generation</p></figcaption></figure>

Generating a private key will initiate a download of a `.pem` file; this should also be treated as a secret and kept safe.

Your GitHub Server App is now ready to be installed against repositories in your Snyk organisation. You can scroll to the top of the same page and select **Install App** from the left-hand panel. Choose the newly created app and click on the **Install** button to the right.

<figure><img src="../../.gitbook/assets/image (624).png" alt=""><figcaption><p>Install the GitHub App in your repositories</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (626).png" alt=""><figcaption><p>Install the GitHub App in your selected repositories</p></figcaption></figure>

Choose where you want to install the app in your GitHub organisation. It can be installed against specific repositories, or all of them in your GitHub organisation.

{% hint style="info" %}
If you select to install the app on a subset of repositories in your GitHub organisation, the app will only work in those repositories. You can edit where the app is installed by returning to this screen at a later date if you want to add it to additional repositories
{% endhint %}

On installation of the app you will receive an `InstallationID`. These are the final numbers in the URL of the page. Make a note of that number, as you will need it to setup a broker connection.

#### Create the GitHub Server App SCM integration

Before the GitHub Server App can be used the integration must be setup within Snyk. In order to setup a brokered integration for the GitHub Server App the API must be used.&#x20;

To setup the integration you will need the following:

* The base API address for your Snyk environment&#x20;
  * For most users this is `https://api.snyk.io`&#x20;
  * A full list can be found [here](https://docs.snyk.io/snyk-api/v1-api#api-urls).
* The Snyk organisation ID you want to setup the integration in.
* A valid [Snyk API token](https://docs.snyk.io/snyk-api/v1-api#authorization).

With the above information make the following API call, replacing the required information inside the `{}` with your data.

```
curl --location 'https://{SNYK-BASE-API}/v1/org/{YOUR-ORG-ID}/integrations' \
--header 'Content-Type: application/json; charset=utf-8' \
--header 'Authorization: token ${REPLACE_WITH_API_TOKEN}' \
--data '{
    "type": "github-server-app",
    "broker": {
        "enabled": true
    }
}'
```

{% hint style="info" %}
On success of the above command you will receive an `id` value; this is the `integrationId` of your new GitHub Server App integration in Snyk. Keep a note of this, as you will need it for the Universal Broker.
{% endhint %}

Visiting the integrations page in Snyk will show that the integration has been configured.&#x20;

<figure><img src="../../.gitbook/assets/image (630).png" alt=""><figcaption><p>A successful GitHub Server App integration</p></figcaption></figure>

Upon finishing the setup and installation of your GitHub Server App, you will have all of the necessary credentials to setup a brokered connection. See the official documentation on Snyk's [Universal Broker](https://docs.snyk.io/enterprise-setup/snyk-broker/universal-broker) for details on how to setup a GitHub Server App connection with Broker.
