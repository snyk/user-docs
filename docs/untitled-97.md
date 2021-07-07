# Azure Repos integration

* [ GitHub integration](/hc/en-us/articles/360004032117-GitHub-integration)
* [ GitHub Enterprise Integration](/hc/en-us/articles/360015951318-GitHub-Enterprise-Integration)
* [ Bitbucket Cloud integration](/hc/en-us/articles/360004032097-Bitbucket-Cloud-integration)
* [ Bitbucket Data Center / Server integration](/hc/en-us/articles/360004002218-Bitbucket-Data-Center-Server-integration)
* [ GitLab integration](/hc/en-us/articles/360004002238-GitLab-integration)
* [ Azure Repos integration](/hc/en-us/articles/360004002198-Azure-Repos-integration)
* [ GitHub Read-Only Projects](/hc/en-us/articles/360010766557-GitHub-Read-Only-Projects)
* [ Opening fix and upgrade pull requests from a fixed GitHub account](/hc/en-us/articles/360010843797-Opening-fix-and-upgrade-pull-requests-from-a-fixed-GitHub-account-)
* [ Test your PRs for vulnerabilities before merging](/hc/en-us/articles/360006528057-Test-your-PRs-for-vulnerabilities-before-merging)
* [ Snyk checks on pull requests](/hc/en-us/articles/360006581938-Snyk-checks-on-pull-requests)

 [See more](/hc/en-us/sections/360001138098-Git-repository-SCM-integrations)

##  Azure Repos integration

The user generates a unique Azure DevOps personal access token \(PAT\) generated for Snyk specifically. Together the username and password constitute a token that Snyk uses. The token authorizes Snyk to access the user’s repos for only the specific permissions that the user indicates to Azure Repos when generating it.

**Note:** The account creating the PAT must be a member of the Project Administrators Group or have Manage permissions set to Allow for Git repositories \(see [Azure's doc](https://docs.microsoft.com/en-us/azure/devops/repos/git/set-git-repository-permissions)\).

1. The user selects projects and repositories for import to Snyk \(for testing and monitoring\). The user can also enter custom file locations for any manifest files that are not located in the root folders of their repositories.
2. Snyk evaluates the items that the user selected and imports any that have relevant manifest files in their root folder and all the subfolders at any level.
3. Snyk communicates directly with your repository for each test it runs to determine exactly what code is currently pushed and what dependencies are being used. Each dependency is tested against Snyk’s vulnerability database to see if it contains any known vulnerabilities.
4. Based on your configurations, if vulnerabilities are found, Snyk notifies you via email or Slack so that you can take immediate remediation action.

### Add projects to Snyk for Azure Repos

Snyk tests and monitors Azure Repos that are in any of our supported languages by evaluating root folders and custom file locations.

#### Adding custom file location

#### Excluding folders from import

This integration works similar to our other integrations. To continue to monitor, remediate and manage your projects, see the relevant pages in our Docs.

### Configure your integration for Azure Repos

**Feature availability**  
Integration with Azure Repos Cloud is available for all of our pricing plans. Integration with Azure Repos Server v2018 Update 2 and above \(also known as TFS\) is available with Enterprise and Business plans. See [Pricing plans](https://snyk.io/plans/) for more details.

Snyk integrates with Microsoft Azure Repos to enable you to import your projects and monitor the source code for your repositories. Snyk tests the projects you’ve imported for any known security vulnerabilities found in the application’s dependencies, testing at a frequency you control.

#### How to configure your integration

Enable integration between Azure Repos and Snyk, and start managing your vulnerabilities.

**Prerequisites**

Ensure you have set up your Azure Repos account and your Snyk account.

**Steps**

1. Access your Azure Repos account and retrieve a unique personal access token for use by Snyk. For help doing this, see the [Azure DevOps documentation](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops).
2. When prompted in Azure, enable the following permissions for Snyk access as follows:
   * Expiry—We recommend choosing an expiration date for this token that is far in the future to avoid breaking integration.
   * Scopes—Custom defined
   * Code—Read and write
3. [Log in](https://app.snyk.io/) to your Snyk account.
4. Navigate to **Integrations** from the menu bar at the top.
5. From the **Integrations** page under the Azure Repos logo, click the **Connect to Azure Repos button:**
6. From the **Settings** page in the **Integrations** area, enter the Azure DevOps organization that you want to integrate with \(i.e. https://dev.azure.com/{org-name}\) and the personal access token that you just generated.  


   \* Enterprise customers can also provide a custom URL for Azure Repos Server private instance which is publicly reachable.

7. Click **Save**.
8. Snyk tests the connection values and the page reloads, now displaying Azure Repos integration information. A confirmation message that the details were saved also appears in green at the top of the screen. In addition, if the connection to Azure failed, a notification appears under the Connected to Azure Repos section.

