# Connect an SCM integration

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that is automatically synced, and provides access to the Integration Hub.

You can customize your AppRisk integrations from the **Integrations Hub** where the following SCMs are available:

* [GitHub](connect-an-scm-integration.md#github-setup-guide)
* [GitLab](connect-an-scm-integration.md#gitlab-setup-guide)
* [Azure DevOps](connect-an-scm-integration.md#azure-devops-setup-guide)
* [BitBucket](connect-an-scm-integration.md#bitbucket-setup-guide)

## GitHub setup guide

### Pulled entities

* Repositories
* Builds - only when using GitHub Actions.
* Scans - only when using Code security.

### Required parameters&#x20;

* Organization Name - The GitHub organization name.
* Profile name - The GitHub user used to read repositories.
* API Token - API access token with the following permissions:
  * `repo`
  * `read:packages`
  * `read:org`
  * `read:user`
  * `user:email`

Authorize your personal access token if you have configured SAML SSO.

If you only want to pull the repositories you own, select the **Pull personal repositories** checkbox on the Integration Hub page.

{% hint style="info" %}
Fine-grained personal access token is not supported.
{% endhint %}

### GitHub enterprise additional parameters

You can use as the host Address the IP/URL of the GitHub server. The default URL is [`https://api.github.com`](https://api.github.com).

### GitHub code security additional parameters

The user associated with the token needs to have write permissions on relevant repositories to collect a breakdown of scan issues.

### API Version

You can use the[ GitHub REST API](https://docs.github.com/en/rest?apiVersion=2022-11-28) repository to access information about the API.

### API token setup

1. Open GitHub and click the Settings menu for your profile.
2. Select Developer settings from the left sidebar.
3. Select Personal access tokens and then Tokens (classic).
4. Click Generate new token and, from the dropdown, select Generate new token (classic).
5. Add a description for your token in the Note field.
6. Select the required permissions: `repo`, `read:packages`, `read:org`, `read:user`, `user:email`.
7. Click Generate token.
8. Copy and store the displayed key.

{% hint style="info" %}
Fine-grained personal access token is not supported.
{% endhint %}

## GitLab setup guide

### Pulled entities

* Users
* Repositories&#x20;

### Required parameters&#x20;

* Host address - The IP/URL of the GitLab server.
* Personal API token - Associated to a user account with permissions to fetch users and repositories. The URL should follow this example: `https://gitlab.com/-/profile/personal_access_tokens`.
* Permissions:
  * `read_api` - Grants read access to the API, including all groups and projects, the container registry, and the package registry.
  * `read_repository` - Grants read-only access to repositories on private projects using Git-over-HTTP or the Repository Files API.

Authorize your personal access token if you have configured SAML SSO.

If you only want to pull the repositories you own, select the **Pull personal repositories** checkbox on the Integration Hub page.

### API version

You can use the[ GitLab REST API v4](https://docs.gitlab.com/ee/api/index.html) repository to access information about the API.

## Azure DevOps setup guide

### Pulled entities

* Repositories

### Required parameters

* Azure DevOps PAT - Generate a personal access token
  * Permissions
    * **Code** - read
      * **Project and Team** - read
      * **Graph** - read
      * **Analytics** - read
      * **Build** - read
      * **Release** - read
      * **Security** - manage
  * Organization - Select **All accessible organizations** or a specific organization.
* Azure DevOps organization names - The names of all the relevant Azure DevOps organizations.
* API URL - The API URL, for example, [`https://dev.azure.com/`](https://dev.azure.com/). You can use a custom URL that is publicly accessible.

### API version

You can use the[ Azure DevOps REST API v6](https://learn.microsoft.com/en-us/rest/api/azure/devops/core/?view=azure-devops-rest-6.0) repository to access information about the API.

### Generate a Personal access token

1. Open Azure DevOps and click the **Settings** menu for your profile.
2. Click **Personal access tokens** and then **New token**.
3. Select the **Code**, **Project** **and** **Team**, **Graph**, and **Analytics** read scopes.
4. Set the expiration to 12 months.
5. Copy the generated personal access token and share it through a secured vault.

### Permissions

Enable the following permissions:

* **Graph** - Read
* **Security** - Manage
* **Scopes** - Custom defined
* **Analytics** - Read
* **Build** - Read
* **Code** - Read
* **Project and Team** - Read
* **Release** - Read

## BitBucket setup guide

{% hint style="info" %}
BitBucket Server and BitBucket Cloud do not support automatic language detection. \
If you use BitBucket Cloud you can manually add the language tags to a repository.&#x20;

Note that for BitBucket Server you are unable to manually add language tags to a repository.
{% endhint %}

### Pulled entities

* Users
* Repositories

### Required parameters

* API URL - The URL of the Bitbucket API
* Username - The [Bitbucket username](https://bitbucket.org/account/settings/)&#x20;
* App password - Create an [API token](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#app-passwords) from your BitBucket account, with the following permissions:
  * **Workspace membership** - Read
  * **Account** - Read
  * **Projects** - Read
  * **Repositories** - Read
  * **Issues** - Read

{% hint style="info" %}
Create a BitBucket app password by following these steps:

1. Open your BitBucket account&#x20;
2. Click the Settings option
3. Click the Personal BitBucket settings option&#x20;
4. Navigate to the App passwords sub-section from the ACCESS MANAGEMENT section.
{% endhint %}

* Workspaces - The Bitbucket workspace names, accessible after authentication).

### API version

You can use the [BitBucket REST API V2](https://developer.atlassian.com/bitbucket/api/2/reference/resource/) repository to access information about the API.\
\


\
