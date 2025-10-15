# Azure DevOps for Snyk Essentials

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that is automatically synced and provides access to the Integration Hub.

### Azure DevOps setup guide

#### Pulled entities <a href="#azure-devops-pulled-entities" id="azure-devops-pulled-entities"></a>

* Repository - the pulled entity retrieved by Snyk Essentials.

#### Prerequisites <a href="#azure-devops-integrate-using-snyk-apprisk" id="azure-devops-integrate-using-snyk-apprisk"></a>

To configure a Group-level integration, you must be a Group Admin or have a custom role that includes the `Edit Snyk Essentials` permissions under the [Group-level permissions](../../../snyk-platform-administration/user-roles/pre-defined-roles.md#group-level-permissions).

#### Integrate using Snyk Essentials <a href="#azure-devops-integrate-using-snyk-apprisk" id="azure-devops-integrate-using-snyk-apprisk"></a>

1. Profile name (`mandatory`): Input your integration profile name.&#x20;
2. Organizations (`mandatory`): Input the names of all the relevant Azure DevOps organizations.
3. Access Token (`mandatory`): Create your Azure DevOps PAT from your Azure DevOps settings.&#x20;
   * Access Token (`mandatory`): Create and add your Access token by following the instructions from the  [Generate a Personal access token from your Azure DevOps settings](azure-devops-for-snyk-essentials.md#generate-a-personal-access-token-from-your-azure-devops-settings) section.
   * API URL (`mandatory`): The API URL, for example, [`https://dev.azure.com/`](https://dev.azure.com/). You can use a custom URL that is publicly accessible.
4. Broker Token (`mandatory`): Create and add your Broker token if you use Snyk Broker.
   * Generate your Broker token by following the instructions from the [Obtain your Broker token for Snyk Broker](../../../enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md) page.&#x20;
   * Copy and paste the Broker token on the integration setup menu from the Integration Hub.
5. Add Backstage Catalog (`optional`): If you want to add your Backstage catalog, follow the instructions from the [Backstage file for SCM Integrations](../application-context-for-scm-integrations/) page.

{% hint style="warning" %}
The following PAT token permissions requirements are for Snyk Essentials integrations. For SCM integration, see the [Azure Respositories (TFS) permissions requirements](../user-permissions-and-access-scopes.md#azure-repositories-tfs-permission-requirements) on the Snyk SCM integrations pages.
{% endhint %}

#### Generate a Personal access token from your Azure DevOps settings

{% hint style="warning" %}
The user account that owns the PAT needs `Basic` access level on the Azure organisation.
{% endhint %}

1. Open Azure DevOps and click the **Settings** menu for your profile.
2. Click **Personal access tokens** and then **New token**.
3. Select the following scopes:
   * Permissions
     * **Code** - read
     * **Project and Team** - read
     * **Analytics** - read
     * **Member Entitlement Management** - read
   * Organization - Select **All accessible organizations** or a specific organization.
4. Set the expiration to 12 months.
5. Copy the generated personal access token and share it through a secured vault.

#### API version <a href="#azure-devops-api-version" id="azure-devops-api-version"></a>

You can use the[ Azure DevOps REST API v6](https://learn.microsoft.com/en-us/rest/api/azure/devops/core/?view=azure-devops-rest-6.0) repository to access information about the API.
