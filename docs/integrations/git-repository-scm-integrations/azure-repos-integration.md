# Azure Repository integration

Snyk's Azure Repository integration lets you:

* Continuously perform security scanning across all the integrated repositories
* Detect vulnerabilities in your open source components
* Provide automated fixes and upgrades

## Setting up an Azure Repository integration

The process to connect Snyk with your Azure repositories includes the following steps:&#x20;

1. Generate a unique Azure DevOps personal access token (PAT) for Snyk, based on a username and password combination, and configured with the specific permissions Snyk needs to access your Azure repositories.&#x20;
2. Select the projects and repositories you want to associate with Snyk for testing and monitoring. \
   You can also enter custom file locations for any manifest files that aren't located in the root folders of your repositories.

Snyk then:&#x20;

1. Evaluates the items you selected and imports the ones that have the relevant manifest files in their root folder and all the subfolders at any level.
2. Communicates directly with your repository for each test it runs (using the permissions you associated with your PAT), to determine exactly which code is currently pushed, and which dependencies are being used. Each dependency is tested against Snyk’s vulnerability database to see if it contains any known vulnerabilities.
3. Notifies you via email or a dedicated Slack channel if vulnerabilities are found (according to the preferences you configured), so that you can take immediate action to fix the issues.

## Add projects to Snyk for Azure Repos

Snyk tests and monitors Azure Repos by evaluating root folders and custom file locations for any of the languages Snyk supports.

To add a default project:&#x20;

1. In Snyk, go to **Projects** > **Add projects**.&#x20;
2. Choose the tool from which to import your projects.\
   The available repositories for the selected integration are displayed in a popup.
3. Select the repos that you want Snyk to monitor for security and license issues. To import all repos for a specific organization, check the organization.
4. Click **Add selected repositories**. \
   Snyk scans the entire file tree for dependency files and imports them to Snyk **** as **** Projects**.**

![](<../../.gitbook/assets/uuid-cae3b5b8-6971-406c-3c00-91c9d1a570a2-en (1).png>)

## Add a custom file location

Use this procedure to add a dependency from a non-default path.

1. From the **Add custom file location** dropdown list, select the repository to configure a custom path. The repo must first be selected from the **Add Projects** view, as described in the previous section.
2. In the text field, enter the relative path in which the manifest file is located, as demonstrated in the image above.

{% hint style="info" %}
**Note**\
This field is case-sensitive.
{% endhint %}

## Exclude folders from import

This integration works similar to our other integrations. To continue to monitor, fix and manage your projects, see the relevant pages in our Docs.

{% hint style="info" %}
**Note**\
This field is case-sensitive and the pattern applies for all repos.
{% endhint %}

**Next steps**

Once repositories are imported, a confirmation appears in green at the top of the screen. The selected files are indicated with a unique icon, they are named by organization/repo, and you can now also filter to view only those projects, as seen in the example below:

![](../../.gitbook/assets/screen-shot-2021-09-16-at-9.12.12-am.png)

This integration works similar to our other integrations. To continue to monitor, fix and manage your projects, see the relevant pages in our Docs.

## Configure your integration for Azure Repos

Snyk integrates with Microsoft Azure Repos to enable you to import your projects and monitor the source code for your repositories. Snyk tests the projects you’ve imported for any known security vulnerabilities found in the application’s dependencies, testing at a frequency you control.

{% hint style="info" %}
**Feature availability**\
Integration with Azure Repos Cloud is available for all of our pricing plans. Integration with Azure DevOps Server 2020 and above (also known as TFS) is available with Enterprise and Business plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="warning" %}
Snyk only supports Git: Snyk does _**not**_ currently support integration with Team Foundation Version Control (TFVC).
{% endhint %}

## How to configure your integration

Enable integration between Azure Repos and Snyk, and start managing your vulnerabilities.

**Prerequisites**

* Ensure that you have set up your Azure Repos account and your Snyk account: you must have an Azure project. If you don't have a project yet, create one in [Azure DevOps](https://docs.microsoft.com/en-us/azure/devops/user-guide/sign-up-invite-teammates?view=azure-devops) or set one up in an [on-premises Azure DevOps](https://docs.microsoft.com/en-us/azure/devops/organizations/projects/create-project?view=azure-devops).
* To create a personal access token (PAT), you must be a member of the [**Project Administrators** Group](https://docs.microsoft.com/en-us/azure/devops/organizations/security/change-project-level-permissions?view=azure-devops) or have your **Manage permissions** set to **Allow** for Git repositories. For more information, see the [Azure DevOps documentation](https://docs.microsoft.com/en-us/azure/devops/repos/git/set-git-repository-permissions).

{% hint style="info" %}
Have a Snyk admin user configure the integration within the UI. \
Collaborator users cannot complete this task.
{% endhint %}

**Steps**

1. Access your Azure Repos account, generate, and copy a unique personal access token to use for Snyk. For more information, see the [Azure DevOps documentation](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops).
2. When prompted in Azure, enable the following permissions for Snyk access:
   * **Expiry**: to avoid breaking the integration, we recommend that you choose a token expiration date that is far in the future&#x20;
   * **Scopes**: _Custom defined_
   * **Code**: _Read & write_ (enable Project Administrators group permissions if the user creating the personal access token is not an admin of the Repositories)
3. Log in to [your Snyk account](https://app.snyk.io) and go to **Integrations**.
4. Click the **Azure Repos** card **** to go to **Organization Settings > Integrations >Azure Repos> Account credentials**.&#x20;
5.  Set the Azure DevOps organization that you want to integrate with by entering the slug for your organization (![](<../../.gitbook/assets/image (153).png>)) and the personal access token that you generated. \


    Enterprise customers can also provide a custom URL for an Azure Repos Server private instance that is publicly reachable.
6. Click **Save**. \
   Snyk tests the connection values and the page reloads, displaying the Azure Repos integration information. A message to confirm that the details were updated is displayed at the top of the screen.&#x20;

![](../../.gitbook/assets/azure-connected\_31july2022.png)

If the connection to Azure fails, a notification is displayed under the **Azure Repos** card title.\
&#x20;<img src="../../.gitbook/assets/azure-no-connect_31july2022.png" alt="" data-size="original">
