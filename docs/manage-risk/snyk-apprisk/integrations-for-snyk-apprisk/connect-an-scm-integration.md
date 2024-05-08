# Connect an SCM integration

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that is automatically synced and provides access to the Integration Hub.

You can customize your AppRisk integrations from the **Integrations Hub** where the following SCMs are available:

* [GitHub](connect-an-scm-integration.md#github-setup-guide)
* [GitLab](connect-an-scm-integration.md#gitlab-setup-guide)
* [Azure DevOps](connect-an-scm-integration.md#azure-devops-setup-guide)
* [BitBucket](connect-an-scm-integration.md#bitbucket-setup-guide)

The following video explains some of the configuration capabilities available for you in the Integrations view:

{% embed url="https://www.youtube.com/watch?v=4W6B9lPU10c" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)!
{% endembed %}

If your SCM instance is not publicly accessible, you must connect using Snyk Broker. For details, see [Snyk Broker - AppRisk](../../../enterprise-configuration/snyk-broker/snyk-broker-apprisk.md).

## GitHub setup guide

### Pulled entities

* Repositories
* Builds - only when using GitHub Actions.
* Scans - only when using Code security.

### Integrate using Snyk AppRisk

1. Profile name (`mandatory`): Input your integration profile name.&#x20;
2. Organizations (`mandatory`): Input the names of all the relevant GitHub organizations.
3. Access Token (`mandatory`): Create your GitHub PAT from your GitHub organization. Follow the instructions in the [Generate a Personal access token from your GitHub settings](connect-an-scm-integration.md#generate-a-personal-access-token-from-your-github-settings) section. Authorize your personal access token if you have configured SAML SSO. See the [How to authorize your Personal Access Token and enable SSO](../../../integrate-with-snyk/snyk-scm-integrations/snyk-github-enterprise-integration.md#how-to-authorize-your-personal-access-token-and-enable-sso) page for more details.

{% hint style="info" %}
If you want to use the Broker Token follow the instructions from the [Snyk Broker AppRisk](../../../enterprise-configuration/snyk-broker/snyk-broker-apprisk.md) page.
{% endhint %}

4. API URL (`mandatory`) - Input the API URL. The default URL is `https://api.github.com`.
5. Pull personal repositories (`optional`): Enable the option if you only want to pull the repositories you own.
6. Add Backstage Catalog (`optional`): If you want to add your Backstage catalog, follow the instructions from the [Backstage file for SCM Integrations](backstage-file-for-scm-integrations.md) page.

#### Generate a Personal access token from your GitHub settings

1. Open GitHub and click the Settings menu for your profile.
2. Select Developer settings from the left sidebar.
3. Select Personal access tokens and then Tokens (classic).
4. Click Generate new token and, from the dropdown, select Generate new token (classic).
5. Add a description for your token in the Note field.
6. Select the required permissions:&#x20;
   * `repo`
   * `read:packages`
   * `read:org`
   * `read:user`
   * `user:email`.
7. Click Generate token.
8. Copy and store the displayed key.

{% hint style="info" %}
Fine-grained personal access token is not supported.
{% endhint %}

### API Version

You can use the[ GitHub REST API](https://docs.github.com/en/rest?apiVersion=2022-11-28) repository to access information about the API.

You can use as the host Address the IP/URL of the GitHub server. The default URL is [`https://api.github.com`](https://api.github.com).

The user associated with the token needs to have write permissions on relevant repositories to collect a breakdown of scan issues.

## GitLab setup guide

### Pulled entities

* Users
* Repositories&#x20;

### Integrate using Snyk AppRisk

1. Profile name (`mandatory`): Input your integration profile name.&#x20;
2. API Token (`mandatory`): Create your GitLab PAT from your GitLab organization. Follow the instructions in [Generate a Personal access token from your GitLab settings section](connect-an-scm-integration.md#generate-a-personal-access-token-from-your-gitlab-settings). Authorize your personal access token if you have configured SAML SSO.

{% hint style="info" %}
If you want to use the Broker Token follow the instructions from the [Snyk Broker AppRisk](../../../enterprise-configuration/snyk-broker/snyk-broker-apprisk.md) page.
{% endhint %}

3. Host URL (`mandatory`): The IP/URL of the GitLab server. The default URL is [`https://gitlab.com`](https://gitlab.com)
4. Verify SSL (`optional`): Enable the option if you want to verify the SSL.
5. Pull personal repositories (`optional`): Enable the option If you only want to pull the repositories you own.
6. Add Backstage Catalog (`optional`): If you want to add your Backstage catalog, follow the instructions from the [Backstage file for SCM Integrations](backstage-file-for-scm-integrations.md) page.

#### Generate a Personal access token from your GitLab settings

1. Navigate to your GitLab profile.
2. Select Edit Profile.
3. On the left sidebar, select Access Token.
4. Select Add New Token.
5. Enter a name and expiry date for the token.
6. Ensure to enable this permission:
   * `read_api` - Grants read access to the API, including all groups and projects, the container registry, and the package registry.
   * `read_repository` - Grants read-only access to repositories on private projects using Git-over-HTTP or the Repository Files API.
7. Click the Create personal access token button.
8. Copy and store the displayed key.

### API version

You can use the[ GitLab REST API v4](https://docs.gitlab.com/ee/api/index.html) repository to access information about the API.

## Azure DevOps setup guide

### Pulled entities

* Repository - the pulled entity retrieved by Snyk AppRisk.

### Integrate using Snyk AppRisk

1. Profile name (`mandatory`): Input your integration profile name.&#x20;
2. Organizations (`mandatory`): Input the names of all the relevant Azure DevOps organizations.
3. Access Token (`mandatory`): Create your Azure DevOps PAT from your Azure DevOps settings. Follow the instructions from the  [Generate a Personal access token from your Azure DevOps settings](connect-an-scm-integration.md#generate-a-personal-access-token-from-your-azure-devops-settings) section.

{% hint style="info" %}
If you want to use the Broker Token follow the instructions from the [Snyk Broker AppRisk](../../../enterprise-configuration/snyk-broker/snyk-broker-apprisk.md) page.
{% endhint %}

4. API URL (`mandatory`): The API URL, for example, [`https://dev.azure.com/`](https://dev.azure.com/). You can use a custom URL that is publicly accessible.
5. Add Backstage Catalog (`optional`): If you want to add your Backstage catalog, follow the instructions from the [Backstage file for SCM Integrations](backstage-file-for-scm-integrations.md) page.

#### Generate a Personal access token from your Azure DevOps settings

1. Open Azure DevOps and click the **Settings** menu for your profile.
2. Click **Personal access tokens** and then **New token**.
3. Select the following scopes:
   * Permissions
     * **Code** - read
     * **Project and Team** - read
     * **Graph** - read
     * **Analytics** - read
     * **Release** - read
     * **Security** - manage
   * Organization - Select **All accessible organizations** or a specific organization.
4. Set the expiration to 12 months.
5. Copy the generated personal access token and share it through a secured vault.

### API version

You can use the[ Azure DevOps REST API v6](https://learn.microsoft.com/en-us/rest/api/azure/devops/core/?view=azure-devops-rest-6.0) repository to access information about the API.

## BitBucket setup guide

{% hint style="info" %}
BitBucket Server and BitBucket Cloud do not support automatic language detection. \
If you use BitBucket Cloud you can manually add the language tags to a repository.&#x20;

Note that for BitBucket Server you are unable to manually add language tags to a repository.
{% endhint %}

### Pulled entities

* Users
* Repositories

### Integrate using Snyk AppRisk

1. Profile name (`mandatory`): Input your integration profile name.&#x20;
2. Access Token (`mandatory`): Create your BitBucket PAT from your BitBucket organization.

{% hint style="info" %}
If you want to use the Broker Token follow the instructions from the [Snyk Broker AppRisk](../../../enterprise-configuration/snyk-broker/snyk-broker-apprisk.md) page.
{% endhint %}

3. API URL (`mandatory`) - Input the API URL.
4. Username (`mandatory`): Input the BitBucket username of your organization.
5. App password (`mandatory`): Create an [API token](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#app-passwords) from your BitBucket account, with the following permissions:
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

6. Service type (`mandatory`): Select the service type, Cloud, or On-premises.
7. Add Backstage Catalog (`optional`): If you want to add your Backstage catalog, follow the instructions from the [Backstage file for SCM Integrations](backstage-file-for-scm-integrations.md) page.

### API version

You can use the [BitBucket REST API V2](https://developer.atlassian.com/bitbucket/api/2/reference/resource/) repository to access information about the API.\
\


\
