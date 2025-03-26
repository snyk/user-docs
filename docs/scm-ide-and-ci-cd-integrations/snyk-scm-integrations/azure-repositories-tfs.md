# Azure Repositories (TFS)

When you want to add new integrations to your Snyk account you need to first decide the level type at which you want to install the integration.

* [Group level ](azure-repositories-tfs.md#group-level-snyk-apprisk-integrations)- Add integrations to your Snyk application that will be available for your Snyk Essentials.&#x20;
* [Organization level](azure-repositories-tfs.md#organization-level-snyk-integrations) - Add integrations for your Snyk application that will be available for all Snyk products, except Snyk Essentials or Snyk AppRisk.

{% hint style="info" %}
If you want to set up integrations for Snyk Essentials, use the Integrations menu at the Group level.
{% endhint %}

## Organization level - Snyk integrations

{% hint style="info" %}
**Feature availability**\
Integration with Azure DevOps Server 2020 and above, also known as TFS, is available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).

Snyk supports only Git. Snyk does not currently support integration with Team Foundation Version Control (TFVC).
{% endhint %}

### Prerequisites for Azure Repositories (TFS) integration

* [Snyk Organization Admin](../../snyk-admin/user-roles/pre-defined-roles.md) user role.
* An Azure project. If you do not have a project yet, create one in [Azure DevOps](https://docs.microsoft.com/en-us/azure/devops/user-guide/sign-up-invite-teammates?view=azure-devops) or set one up in an [on-premise Azure DevOps](https://docs.microsoft.com/en-us/azure/devops/organizations/projects/create-project?view=azure-devops) instance.
* The required Personal Access Token (PAT) access scopes. For details on the permissions required, see [Azure Repositories (TFS) permission requirements](./#azure-repositories-tfs-permission-requirements).

### Azure Repositories (TFS) integration features

Snyk integrates with your Microsoft Azure Repository to let you import Projects and monitor the source code for your repositories. Snyk tests the Projects you have imported for known security vulnerabilities in the dependencies, testing at a frequency you control.

The Azure Repository integration lets you:

* Continuously perform security scanning across all the integrated repositories
* Detect vulnerabilities in your open-source components
* Provide automated fixes and upgrades

After the integration is configured, Snyk does the following:

1. Evaluates the items you selected and imports the ones that have the relevant manifest files in their root folder and all the subfolders at any level.
2. Communicates directly with your repository for each test it runs using the permissions you associated with your PAT, to determine exactly which code is currently pushed by the Snyk application and which dependencies are being used. Each dependency is tested against the Snyk vulnerability database to see if it contains any known vulnerabilities.
3. Notifies you by email or a dedicated Slack channel if vulnerabilities are found according to the preferences you configured, so that you can take immediate action to fix the issues.

### How to set up the Azure Repositories (TFS) integration

The process to connect Snyk with your Azure repositories includes the following steps:

1. Generate a unique Azure DevOps personal access token (PAT) for Snyk, based on a username and password combination, and configured with the specific permissions Snyk needs to access your Azure repositories. For more information, see [Configure a Personal Access Token (PAT)](azure-repositories-tfs.md#configure-a-personal-access-token-pat).
2. [Enable the integration through the Snyk Web UI](azure-repositories-tfs.md#integrate-using-the-snyk-web-ui).
3. [Select the Projects and repositories](azure-repositories-tfs.md#add-projects-to-snyk-for-azure-repos) you want to associate with Snyk for testing and monitoring.\
   You can also enter custom file locations for any manifest files that are not located in the root folders of your repositories.

### **Configure a Personal Access Token (PAT)**

Generate and copy a unique PAT to use for Snyk. For more information on the PAT access scope requirements to enable in Azure, see [Azure Repositories (TFS) permission requirements](./#azure-repositories-tfs-permission-requirements).

### Integrate using the Snyk Web UI

1. Log in to [your Snyk account](https://app.snyk.io) and navigate to **Integrations**.
2. On the **Azure Repos** tile, click the settings icon to open **Organization Settings** > **Integrations** > **Azure Repos** > **Account credentials**.
3. Pay special attention to the instructions given on the **Account Credentials** page. Depending on your plan, you may need to enter just the Azure DevOps Organization, or you may need to enter the entire URL.
   * **Set your organization**: Enter the slug for your Organization only. \
     For example, enter `your-azure-devops-org`
   * **Set your host**: enter the entire url. \
     For example, enter `https://dev.azure.com/your-azure-devops-org` \
     Alternatively, you may enter a custom url that is publicly reachable
4. Click **Save**, and then enter the PAT that you generated.
5. Click **Save**.\
   Snyk tests the connection values and the page reloads, displaying the Azure Repos integration information. A message to confirm that the details were updated appears at the top of the screen.

{% hint style="info" %}
If the connection to Azure fails, a notification will appear under the **Azure Repos** card title.
{% endhint %}

### Add Projects to Snyk for Azure Repos

Snyk tests and monitors Azure Repos by evaluating root folders and custom file locations for the [languages that Snyk supports](../../supported-languages-package-managers-and-frameworks/).

To add a default Project:

1. In Snyk, navigate to **Projects** > **Add projects**.
2. Choose the relevant repository or tool from which to import your projects.\
   The available repositories for the integration you chose are displayed in a new window.
3. Select the repositories that you want Snyk to monitor for security and license issues.
4. To import all the repos for a specific Organization, check the **Organization**.
5. Click **Add selected repositories**.\
   Snyk scans the entire file tree for dependency files and imports them to Snyk as Projects.

### Adding custom file locations and excluding folders

#### Add a custom file location

Use this procedure to add an Azure Repository dependency from a non-default path.

1. In Snyk, navigate to **Projects** > **Add projects > Azure repos > Settings**.
2. Open the **Add custom file location (optional)** list and **select a repository...** to configure a custom path.
3. In the text field, enter the `relative path for the manifest file location`.

{% hint style="warning" %}
The relative path field is case-sensitive.
{% endhint %}

#### Exclude folders from import

The optional **Exclude folders** field is case-sensitive. The pattern you enter is applied to all the Azure repositories.

### **Confirming the repository import**

After repositories are imported, a confirmation appears in green at the top of the screen. The selected files are marked with a unique icon and named by Organization and repository. You can filter to view only those Projects by selecting the Azure Repos filter option.

{% hint style="info" %}
The Azure Repository integration works like the other [Snyk SCM integrations](./). To continue to monitor, fix, and manage your Projects, see the [Projects](../../snyk-admin/snyk-projects/) documentation.
{% endhint %}

## Group level - Snyk Essentials integrations

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that is automatically synced and provides access to the Integration Hub.

### Azure DevOps setup guide

#### Pulled entities <a href="#azure-devops-pulled-entities" id="azure-devops-pulled-entities"></a>

* Repository - the pulled entity retrieved by Snyk Essentials.

#### Integrate using Snyk Essentials <a href="#azure-devops-integrate-using-snyk-apprisk" id="azure-devops-integrate-using-snyk-apprisk"></a>

1. Profile name (`mandatory`): Input your integration profile name.&#x20;
2. Organizations (`mandatory`): Input the names of all the relevant Azure DevOps organizations.
3. Access Token (`mandatory`): Create your Azure DevOps PAT from your Azure DevOps settings.&#x20;
   * Access Token (`mandatory`): Create and add your Access token by following the instructions from the  [Generate a Personal access token from your Azure DevOps settings](azure-repositories-tfs.md#generate-a-personal-access-token-from-your-azure-devops-settings) section.
   * API URL (`mandatory`): The API URL, for example, [`https://dev.azure.com/`](https://dev.azure.com/). You can use a custom URL that is publicly accessible.
4. Broker Token (`mandatory`): Create and add your Broker token if you use Snyk Broker.
   * Generate your Broker token by following the instructions from the [Obtain your Broker token for Snyk Broker ](../../enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md#obtain-your-broker-token-for-snyk-broker-code-agent)page.&#x20;
   * Copy and paste the Broker token on the integration setup menu from the Integration Hub.
5. Add Backstage Catalog (`optional`): If you want to add your Backstage catalog, follow the instructions from the [Backstage file for SCM Integrations](application-context-for-scm-integrations/) page.

{% hint style="warning" %}
The following PAT token permissions requirements are for Snyk Essentials integrations. For SCM integration, see the [Azure Respositories (TFS) permissions requirements](./#azure-repositories-tfs-permission-requirements) on the Snyk SCM integrations pages.
{% endhint %}

#### Generate a Personal access token from your Azure DevOps settings

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
