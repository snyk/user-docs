# Snyk Azure Repositories (TFS) integration

{% hint style="info" %}
**Feature availability**\
Integration with Azure Repos Cloud is available with all Snyk pricing plans. Integration with Azure DevOps Server 2020 and above, also known as TFS, is available for Snyk Enterprise plan customers. For details, see [Plans and pricing](https://snyk.io/plans/).

Snyk supports only Git. Snyk does not currently support integration with Team Foundation Version Control (TFVC).
{% endhint %}

Snyk integrates with your Microsoft Azure Repository to let you import Projects and monitor the source code for your repositories. Snyk tests the Projects you have imported for known security vulnerabilities in the dependencies, testing at a frequency you control.

Snyk Azure Repository integration lets you:

* Continuously perform security scanning across all the integrated repositories
* Detect vulnerabilities in your open-source components
* Provide automated fixes and upgrades

## How Azure Repository integration works

The process to connect Snyk with your Azure repositories includes the following steps:

1. Generate a unique Azure DevOps personal access token (PAT) for Snyk, based on a username and password combination, and configured with the specific permissions Snyk needs to access your Azure repositories.\
   For more information, see [Configure your Azure Repository integration](snyk-azure-repositories-tfs-integration.md#configure-your-azure-repository-integration).
2. Select the Projects and repositories you want to associate with Snyk for testing and monitoring.\
   You can also enter custom file locations for any manifest files that are not located in the root folders of your repositories.

Snyk then does the following:

1. Evaluates the items you selected and imports the ones that have the relevant manifest files in their root folder and all the subfolders at any level.
2. Communicates directly with your repository for each test it runs using the permissions you associated with your PAT, to determine exactly which code is currently pushed by the Snyk application and which dependencies are being used. Each dependency is tested against the Snyk vulnerability database to see if it contains any known vulnerabilities.
3. Notifies you by email or a dedicated Slack channel if vulnerabilities are found according to the preferences you configured, so that you can take immediate action to fix the issues.

## Configure your Azure Repository integration

### Prerequisites for Azure Repository integration

{% hint style="info" %}
Only a Snyk admin user can configure the integration within the UI.\
Collaborator users cannot perform this task.
{% endhint %}

To enable integration between Azure Repository and Snyk:

* Set up your Azure Repos account and your Snyk account.
  * You must have an Azure project.
  * If you do not have a project yet, create one in [Azure DevOps](https://docs.microsoft.com/en-us/azure/devops/user-guide/sign-up-invite-teammates?view=azure-devops) or set one up in an [on-premise Azure DevOps](https://docs.microsoft.com/en-us/azure/devops/organizations/projects/create-project?view=azure-devops) instance.
* &#x20;Create a Personal Access Token (PAT). You must be a member of the [Project Administrators Group](https://docs.microsoft.com/en-us/azure/devops/organizations/security/change-project-level-permissions?view=azure-devops) so that the PAT has the `edit subscriptions permissions` required to enable webhooks.

### **Integrate with your Azure Repository**

1. Generate and copy a unique PAT to use for Snyk. For more information, see the [Azure DevOps documentation](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops).
2. When you are prompted in Azure, enable the following permissions for Snyk access:
   * **Expiry**: to avoid breaking the integration, Snyk recommends that you choose a token expiration date that is far in the future.
   * **Scopes**: Custom defined
   * **Code**: `Read & write`.

### Integrate using the Snyk Web UI

1. Log in to [your Snyk account](https://app.snyk.io) and navigate to **Integrations**.
2. On the **Azure Repos** tile, click the settings icon to open **Organization Settings** > **Integrations** > **Azure Repos** > **Account credentials**.\
   <img src="../../.gitbook/assets/integrations -azure_repo_tile-11aug2022.png" alt="Azure Repos integration tile" data-size="original">
3. Pay special attention to the instructions given on the **Account Credentials** page. Depending on your plan, you may need to enter just the Azure DevOps Organization, or you may need to enter the entire URL.
   *   **Set your organization**: Enter the slug for your Organization only. \
       For example, enter `your-azure-devops-org`

       <figure><img src="../../.gitbook/assets/image (345) (1).png" alt="If the Organization prompt appears, enter only the &#x27;org slug&#x27; found in your azure URL"><figcaption><p>If the Organization prompt appears, enter only the 'org slug' found in your azure URL</p></figcaption></figure>
   *   **Set your host**: enter the entire url. \
       For example, enter `https://dev.azure.com/your-azure-devops-org` \
       Alternatively, you may enter a custom url that is publicly reachable

       <figure><img src="../../.gitbook/assets/image (344).png" alt="Enterprise customers can provide a custom URL for an Azure Repos Server private instance that is publicly reachable"><figcaption><p>Enterprise customers can provide a custom URL for an Azure Repos Server private instance that is publicly reachable</p></figcaption></figure>
4. Click **Save**, and then enter the PAT that you generated.
5. Click **Save**.\
   Snyk tests the connection values and the page reloads, displaying the Azure Repos integration information. A message to confirm that the details were updated appears at the top of the screen.

If the connection to Azure fails, a notification appears under the **Azure Repos** card title.

<img src="../../.gitbook/assets/azure-no-connect_31july2022.png" alt="Could not connect to Azure Repos. Ensure you have entered your account credential correctly." data-size="original">



## Add Projects to Snyk for Azure Repos

Snyk tests and monitors Azure Repos by evaluating root folders and custom file locations for the [languages that Snyk supports](../../scan-applications/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/).

To add a default Project:

1. In Snyk, navigate to **Projects** > **Add projects**.
2. Choose the relevant repository or tool from which to import your projects.\
   The available repositories for the integration you chose are displayed in a new window.
3. Select the repositories that you want Snyk to monitor for security and license issues.
4. To import all the repos for a specific Organization, check the **Organization**.
5. Click **Add selected repositories**.\
   Snyk scans the entire file tree for dependency files and imports them to Snyk as Projects.

<div align="left">

<figure><img src="../../.gitbook/assets/uuid-cae3b5b8-6971-406c-3c00-91c9d1a570a2-en.png" alt="Import Projects"><figcaption><p>Import Projects</p></figcaption></figure>

</div>

## Adding custom file locations and excluding folders

### Add a custom file location

Use this procedure to add an Azure Repository dependency from a non-default path.

1. In Snyk, navigate to **Projects** > **Add projects > Azure repos > Settings**.
2. Open the **Add custom file location (optional)** list and **select a repository...** to configure a custom path.
3. In the text field, enter the `relative path for the manifest file location`.

{% hint style="info" %}
The relative path field is case-sensitive.
{% endhint %}

<div align="left">

<figure><img src="../../.gitbook/assets/azure_custom_repo-11aug2022.png" alt="Select Azure repos" width="563"><figcaption><p>Select Azure repos</p></figcaption></figure>

</div>

### Exclude folders from import

The optional **Exclude folders** field is case-sensitive. The pattern you enter is applied to all the Azure repositories.

## **Confirming the repository import**

After repositories are imported, a confirmation appears in green at the top of the screen. The selected files are marked with a unique icon and named by Organization and repository. You can filter to view only those Projects, as shown in the example that follows:

<div align="left">

<figure><img src="../../.gitbook/assets/image (37) (1) (1).png" alt="View import results"><figcaption><p>View import results</p></figcaption></figure>

</div>

{% hint style="info" %}
The Azure Repository integration works like the other Snyk [Git integrations](./). To continue to monitor, fix, and manage your Projects, see the [Projects](../../snyk-admin/snyk-projects/) documentation.
{% endhint %}
