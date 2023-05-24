# Azure Repositories integration

Snyk's Azure Repository integration lets you:

* Continuously perform security scanning across all the integrated repositories
* Detect vulnerabilities in your open source components
* Provide automated fixes and upgrades

## Setting up an Azure Repository integration

The process to connect Snyk with your Azure repositories includes the following steps:

1. Generate a unique Azure DevOps personal access token (PAT) for Snyk, based on a username and password combination, and configured with the specific permissions Snyk needs to access your Azure repositories.
2. Select the projects and repositories you want to associate with Snyk for testing and monitoring.\
   You can also enter custom file locations for any manifest files that aren't located in the root folders of your repositories.

Snyk then:

1. Evaluates the items you selected and imports the ones that have the relevant manifest files in their root folder and all the subfolders at any level.
2. Communicates directly with your repository for each test it runs (using the permissions you associated with your PAT), to determine exactly which code is currently pushed, and which dependencies are being used. Each dependency is tested against Snyk’s vulnerability database to see if it contains any known vulnerabilities.
3. Notifies you via email or a dedicated Slack channel if vulnerabilities are found (according to the preferences you configured), so that you can take immediate action to fix the issues.

## Add Projects to Snyk for Azure Repos

Snyk tests and monitors Azure Repos by evaluating root folders and custom file locations for the [languages that Snyk supports](https://docs.snyk.io/introducing-snyk/snyk-languages-and-integrations#supported-languages).

To add a default Project:

1. In Snyk, go to **Projects** > **Add projects**.
2. Choose the relevant repository or tool from which to import your Projects.\
   The available repositories for the integration you picked are displayed in a new window.
3. Select the repositories that you want Snyk to monitor for security and license issues. To import all the repos for a specific Organization, check the Organization.
4. Click **Add selected repositories**.\
   Snyk scans the entire file tree for dependency files and imports them to Snyk as Projects.

<div align="left">

<figure><img src="../../.gitbook/assets/uuid-cae3b5b8-6971-406c-3c00-91c9d1a570a2-en (1).png" alt="Import Projects"><figcaption><p>Import Projects</p></figcaption></figure>

</div>

## Adding custom file locations and excluding folders

### Add a custom file location

Use this procedure to add an Azure Repository dependency from a non-default path.

1. In Snyk, go to **Projects** > **Add projects > Azure repos > Settings** card.
2. Open the **Add custom file location (optional)** list and **select a repository...** to configure a custom path.
3. In the text field, enter the relative path for the manifest file location.

{% hint style="info" %}
The relative path field is case-sensitive.
{% endhint %}

<div align="left">

<figure><img src="../../.gitbook/assets/azure_custom_repo-11aug2022 (1).png" alt="Select Azure repos" width="563"><figcaption><p>Select Azure repos</p></figcaption></figure>

</div>

### Exclude folders from import

The Azure Repository integration works similar to the other Snyk Git integrations. To continue to monitor, fix, and manage your projects, see the related pages in the Snyk User Docs.

{% hint style="info" %}
The optional **Exclude folders** field is case-sensitive. The pattern you enter is applied to all the Azure repositories.
{% endhint %}

## **Confirming the repository import**

Once repositories are imported, a confirmation appears in green at the top of the screen. The selected files are indicated with a unique icon, they are named by organization/repo, and you can now also filter to view only those projects, as seen in the example below:

<div align="left">

<figure><img src="../../.gitbook/assets/image (22) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="View import results"><figcaption><p>View import results</p></figcaption></figure>

</div>

The Azure Repository integration works similarly to the other Snyk [Git integrations](./).

## Configure your Azure Repository integration

Snyk integrates with Microsoft Azure Repository to let you import your Projects and monitor the source code for your repositories. Snyk tests the Projects you’ve imported for known security vulnerabilities in the application’s dependencies, testing at a frequency you control.

{% hint style="info" %}
**Feature availability**\
Integration with Azure Repos Cloud is available for all of our pricing plans. Integration with Azure DevOps Server 2020 and above (also known as TFS) is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="warning" %}
Snyk only supports Git: Snyk does not currently support integration with Team Foundation Version Control (TFVC).
{% endhint %}

### Prerequisites

{% hint style="info" %}
Only a Snyk admin user can configure the integration within the UI.\
Collaborator users cannot complete this task.
{% endhint %}

To enable integration between Azure Repository and Snyk:

* Set up your Azure Repos account and your Snyk account: you must have an Azure project. If you don't have a project yet, create one in [Azure DevOps](https://docs.microsoft.com/en-us/azure/devops/user-guide/sign-up-invite-teammates?view=azure-devops) or set one up in an [on-premises Azure DevOps](https://docs.microsoft.com/en-us/azure/devops/organizations/projects/create-project?view=azure-devops).
* &#x20;Create a Personal Access Token (PAT). You must be a member of the [**Project Administrators** Group](https://docs.microsoft.com/en-us/azure/devops/organizations/security/change-project-level-permissions?view=azure-devops) so that the PAT has _**edit**_ _**subscriptions**_ _**permissions**_ required to enable webhooks.

### **In your Azure Repository**

1. Generate and copy a unique PAT to use for Snyk. For more information, see the [Azure DevOps documentation](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops).
2. When you're prompted in Azure, enable the following permissions for Snyk access:
   * **Expiry**: to avoid breaking the integration, we recommend that you choose a token expiration date that is far in the future
   * **Scopes**: _Custom defined_
   * **Code**: _Read & write_.

### In the Snyk Web UI

1. Log in to [your Snyk account](https://app.snyk.io) and go to **Integrations**.
2. In the **Azure Repos** tile, click <img src="../../.gitbook/assets/integration_settings_tile_cog-11aug2022.png" alt="" data-size="line"> to open **Organization Settings > Integrations >Azure Repos> Account credentials**.\
   <img src="../../.gitbook/assets/integrations -azure_repo_tile-11aug2022.png" alt="" data-size="original">
3.  Set the Azure DevOps Organization that you want to integrate with by entering the slug for your Organization (![](<../../.gitbook/assets/image (127) (1) (1) (1) (1) (1).png>)) and the PAT that you generated.

    Enterprise customers can also provide a custom URL for an Azure Repos Server private instance that is publicly reachable.
4. Click **Save**.\
   Snyk tests the connection values and the page reloads, displaying the Azure Repos integration information. A message to confirm that the details were updated is displayed at the top of the screen.

If the connection to Azure fails, a notification is displayed under the **Azure Repos** card title.\
Azure Repositories source control<img src="../../.gitbook/assets/azure-no-connect_31july2022.png" alt="" data-size="original">
