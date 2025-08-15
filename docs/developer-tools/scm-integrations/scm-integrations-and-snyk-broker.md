# SCM integrations and Snyk Broker

If your SCM instance is not publicly accessible, you need Snyk Broker. You can install and configure your Snyk Broker using Docker or Helm. For more information about Snyk Broker, see the [Snyk Broker](../../implementation-and-setup/enterprise-setup/snyk-broker/) documentation, including [Using Snyk Essentials wtih Snyk Broker](../../implementation-and-setup/enterprise-setup/snyk-broker/using-snyk-essentials-with-snyk-broker.md).

{% hint style="warning" %}
Enable the Snyk Essentials flag in your Snyk Broker deployment environment before running the commands.
{% endhint %}

You can find on [GitHub](https://github.com/snyk/broker/tree/565242baf003f06f445489dd96cc68c8386ede38/defaultFilters/apprisk) all the updated `.json` files that include the allowed list of accessible endpoints for the integrations.

## Integrated SCM tokens for Classic Broker

An integrated SCM token is required for [Broker client setup](../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/). It is used in the `-e <SCM>_TOKEN` parameter, for example, `-e GITHUB_TOKEN=xxx…`, to enable access to the SCM. These meet certain permissions needed for the operation of Broker and Snyk Code.

An integrated SCM token can be generated for the following SCM integrations:

* [GitHub and GitHub Enterprise](scm-integrations-and-snyk-broker.md#github-and-github-enterprise-scm-token)
* [GitLab](scm-integrations-and-snyk-broker.md#gitlab-scm-token)
* [Azure Repositories (TFS)](scm-integrations-and-snyk-broker.md#azure-repositories-tfs-scm-token)
* [Bitbucket Server/Data Center](scm-integrations-and-snyk-broker.md#bitbucket-server-data-center-scm-token)

### GitHub and GitHub Enterprise SCM token

**Format**: `GITHUB_TOKEN=` - a GitHub personal access token. \
**Scopes:** `repo, read:org` and `admin:repo_hook.`

### GitLab SCM token

**Format**: `GITLAB_TOKEN=` - a GitLab personal access token.\
**Scopes**: `api`.

GitLab account with `Maintainer` permission.

### Azure Repositories (TFS) SCM token

**Format**: `AZURE_REPOS_TOKEN=` - an Azure Repos personal access token. \
**Scopes**: `Custom defined`,  `Code:`  `Read & write`_._

### Bitbucket Server/Data Center SCM token

**Format**: `BITBUCKET_USERNAME=`, `BITBUCKET_PASSWORD=` – the Bitbucket Server username and password or a Bitbucket Server personal access token.\
**Scope**: `Repository admin`**.**

## GitHub Server App for Universal Broker

If your GithHub server is not publicly available, you can provide access to it through the Universal Broker, a proxy deployed in your internal network to facilitate outbound connections and communication with Snyk.

The setup process for Universal Broker involves:

1. [Creating a GitHub App on your GitHub Server instance](scm-integrations-and-snyk-broker.md#create-a-github-app-for-universal-broker)
2. [Creating the Universal Broker connection using the `snyk-broker-config` command](scm-integrations-and-snyk-broker.md#create-the-universal-broker-connection-for-your-github-server-app)

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

<figure><img src="../../.gitbook/assets/image (290).png" alt=""><figcaption><p>Create GitHub App account setting selection</p></figcaption></figure>

{% hint style="warning" %}
On creation of your GitHub Server App you will see a`ClientId` and `AppId` - store these safely as these are your app's credentials and should be treated as secrets.
{% endhint %}

After creating your GitHub Server App you will see a banner at the top of the page prompting you to create a private key. Click on it and create a private key for your app.

<figure><img src="../../.gitbook/assets/image (298).png" alt=""><figcaption><p>Registration success message with a link for private key generation</p></figcaption></figure>

Generating a private key will initiate a download of a `.pem` file; this should be treated as a secret and kept safe.

Your GitHub Server App is now ready to be installed against repositories in your Snyk Organization. You can scroll to the top of the same page and select **Install App** from the left-hand panel. Choose the newly created app and click on the **Install** button to the right.

<figure><img src="../../.gitbook/assets/image (304).png" alt=""><figcaption><p>Install the GitHub App in your repositories</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (334).png" alt="" width="563"><figcaption><p>Install the GitHub App in your selected repositories</p></figcaption></figure>

Choose where you want to install the app in your GitHub organization. It can be installed against specific repositories, or all of them in your GitHub organization.

{% hint style="info" %}
If you select to install the app on a subset of repositories in your GitHub organization, the app will work only in those repositories. You can edit where the app is installed by returning to this screen at a later date if you want to add it to additional repositories
{% endhint %}

On installation of the app you will be assigned an `InstallationID`. These are the final numbers in the URL of the page. Make a note of that number, as you will need it to set up a Universal Broker connection.

### Create the Universal Broker connection for your GitHub Server App

Before the GitHub Server App can be used with the Universal Broker, you must create a connection of the `github-server-app` type using the `snyk-broker-config` tool. For more details, see the  [Universal Broker](../../implementation-and-setup/enterprise-setup/snyk-broker/universal-broker/) documentation. After the connection is created, it can be integrated with one or more Organization(s) of your choice.

To set up the integration you will need the following:

* The base API address for your Snyk region; refer to the list of [API URLs](../../snyk-data-and-governance/regional-hosting-and-data-residency.md#api-urls) for Snyk regional hosting.
* The Snyk Organization ID you want to set up the integration in.
* A valid [Snyk API token](../../snyk-api/authentication-for-api/#how-to-obtain-your-personal-token).
* `snyk-broker-config` tool installed
* Tenant Admin role

When you have these prerequisites, use the following commands

* `snyk-broker-config workflows connections create,`choosing the `github-server-app` option and providing the information you are prompted for in the workflow.
* `snyk-broker-config workflows connections integrate`to integrate the newly created connection to the Organization of your choice, pasting the org ID when you are prompted.

Visit the integrations page in Snyk to see that the integration has been configured.&#x20;

<figure><img src="../../.gitbook/assets/image (306).png" alt=""><figcaption><p>A successful GitHub Server App integration</p></figcaption></figure>

See the [Universal Broker](../../implementation-and-setup/enterprise-setup/snyk-broker/universal-broker/) documentation for more details.
