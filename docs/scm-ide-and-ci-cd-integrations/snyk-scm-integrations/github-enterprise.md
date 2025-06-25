# GitHub Enterprise

When you want to add new integrations to your  Snyk account you need to first decide the level type at which you want to install the integration.

* [Group level](github-enterprise.md#group-level-snyk-essentials-integrations) - Add integrations to your Snyk application that will be available for Snyk Essentials. To set up integrations for Snyk Essentials, use the Integrations menu at the Group level.
* [Organization level](github-enterprise.md#organization-level-snyk-integrations) - Add integrations for your Snyk application that will be available for all Snyk products, except Snyk Essentials and Snyk AppRisk.

## Organization level - Snyk integrations

{% hint style="info" %}
**Feature availability**

GitHub Enterprise integration is available only with Enterprise plans. For more information, see [Plans and pricing](https://snyk.io/plans/).
{% endhint %}

If you are a Snyk Enterprise plan customer, Snyk recommends that you use the GitHub Enterprise integration. If you use the self-hosted GitHub Enterprise product, you must use the GitHub Enterprise integration.

### Prerequisites for GitHub Enterprise integration

* Internet-accessible repositories.\
  If your repositories are not internet-accessible, you must use [Snyk Broker](../../enterprise-setup/snyk-broker/). This requires creating a startup script. For the script and instructions, see [GitHub Enterprise - install and configure using Docker](../../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-docker.md).
* A public or private GitHub project.
* The required [PAT](github-enterprise.md#generate-a-personal-access-token) and GitHub repository access scope permissions. For more information, see [GitHub and GitHub Enterprise permission requirements](./#github-and-github-enterprise-permissions-requirements).

{% hint style="info" %}
You do not need to be on a GitHub Enterprise level plan to use the Snyk GitHub Enterprise integration.
{% endhint %}

### GitHub Enterprise integration features

The GitHub Enterprise integration lets you:

* Perform [periodic security scans](github-enterprise.md#obtain-project-level-security-reports) across all integrated repositories.
* [Detect vulnerabilities](github-enterprise.md#monitor-projects-and-generate-automatic-fix-pull-requests) in your Open Source components.
* Provide [automated fixes](github-enterprise.md#test-new-pull-requests) and upgrades through status checks in GitHub.

### How to set up the GitHub Enterprise integration

Follow these steps to connect Snyk with your GitHub repositories:

1. Create a dedicated service account in GitHub Enterprise with a write level or higher scope for the repos you want to monitor with Snyk permissions.\
   See [Types of GitHub accounts](https://docs.github.com/en/get-started/learning-about-github/types-of-github-accounts) and [Required access scopes for the GitHub integration ](github-enterprise.md#required-access-scopes-for-snyk-github-enterprise-integration)for details.\
   Note that to create webhooks, which is required for PR checks, the repo permission for the account must be `Admin`. GitHub custom roles are not supported.&#x20;
2. [Generate a personal access token](github-enterprise.md#generate-a-personal-access-token) for that account.
3. [Authorize your personal access token and enable SSO.](github-enterprise.md#how-to-authorize-your-personal-access-token-and-enable-sso)
4. [Import your GitHub repositories](github-enterprise.md#how-to-import-github-repositories).

#### Generate a Personal Access Token

A personal access token (PAT) or fine-grained PAT must be generated in GitHub Enterprise under **Developer settings**.

See [GitHub and GitHub Enterprise permissions requirements](./#github-and-github-enterprise-permissions-requirements) for detailed information on the access scope requirements you must adhere to.

#### How to authorize your Personal Access Token and enable SSO

1. In Snyk, navigate to the **Integrations** page and click the **GitHub Enterprise** card.
2. Enter your GitHub Enterprise URL and the personal access token (PAT) for the service account you created, and **Save** your changes. After Snyk has successfully connected to the GitHub instance, the list of available repositories displays for your selection.
3. If your GitHub Enterprise organization enforces SAML/SSO, select **Configure SSO** next to the PAT in GitHub after the PAT has been created.\
   Occasionally, SSO is enforced in your GitHub Enterprise organizations after a PAT and Integration are configured. If this happens, any Projects that have already been imported show in Snyk, but retests, PR Checks, and so on, will not be performed. To fix this, check the **Configure SSO** settings to ensure the GitHub Enterprise organization is **Authorized**.\
   If the organization is showing as **Authorized**, but the issue still persists, try de-authorizing the organization and then re-authorizing.

{% hint style="info" %}
To use the integration with GitHub Enterprise Cloud, add the URL `https://api.github.com`. To integrate with a self-hosted GitHub Enterprise, add the URL `https://your.github-enterprise.host` in step two of PAT authorization.

Ensure that there are no trailing characters such as `/` following the url. An integration with trailing characters in the URL may connect successfully but provide incorrect links back to the GitHub files.
{% endhint %}

{% hint style="warning" %}
If the PAT token changes or expires in GitHub, the integration with Snyk will not function. To resolve this, update the token in the Snyk **GitHub Enterprise Integration settings.**
{% endhint %}

#### How to import GitHub repositories

Select the repositories you want to import to Snyk and click **Add selected repositories**.

Snyk starts scanning the selected repositories for dependency files, such as `package.json`, in the entire directory tree and imports the repositories to Snyk as Projects.

The imported Projects appear on your **Projects** page and are continuously checked for vulnerabilities.

<figure><img src="../../.gitbook/assets/github_integration-fix_15dec2022 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (16) (26) (8).jpeg" alt="Importing selected repositories to Snyk"><figcaption><p>Importing selected repositories to Snyk</p></figcaption></figure>

### Uses of the GitHub Enterprise integration

#### Obtain Project-level security reports

Snyk produces advanced security reports, allowing you to explore the vulnerabilities found in your repositories and fix them by opening a fix pull request directly to your repository with the required upgrades or patches.

The example that follows shows a Project-level security report.

<figure><img src="../../.gitbook/assets/project_lvl_security_rpt-18july2022.png" alt="Project-level security report"><figcaption><p>Project-level security report</p></figcaption></figure>

#### Monitor Projects and generate automatic fix pull requests

Snyk scans your Projects on either a daily or a weekly basis. When new vulnerabilities are found, Snyk notifies you by email and opens an automated pull request with fixes for your repositories.

The example that follows shows a fix pull request opened by Snyk.

<figure><img src="../../.gitbook/assets/github_fix_pr_cropped-14july2022 (1) (1).png" alt="Fix pull request created by Snyk"><figcaption><p>Fix pull request created by Snyk</p></figcaption></figure>

To review and update the automatic fix pull request settings:

1. In Snyk, navigate to **Settings** > **Integrations** > **Source control** > **GitHub Enterprise** > **Edit Settings**.
2. Scroll to the **Automatic fix pull requests** section, then select options as required:

<figure><img src="../../.gitbook/assets/Screenshot 2023-04-28 at 15.41.56.png" alt="Automatic pull requests settings"><figcaption><p>Automatic pull request settings</p></figcaption></figure>

#### Test new pull requests

The [PR Checks](../../scan-with-snyk/pull-requests/pull-request-checks/) feature enables Snyk to test any newly-created pull requests in your repositories for security vulnerabilities and sends a status check to GitHub. This allows you to see, directly from GitHub, whether the pull request introduces new security issues.

The following example shows how Snyk pull request checks appear on the pull requests page in GitHub.

<figure><img src="../../.gitbook/assets/pr_testing-14july2022.png" alt="Pull request checks shown in GitHub Enterprise"><figcaption><p>Pull request checks shown in GitHub Enterprise</p></figcaption></figure>

To review and adjust the pull request test settings: In Snyk, navigate to Organization **Settings** > **Integrations** > **Source control** > **GitHub Enterprise**, and select **Edit Settings**.

1. Scroll to **Snyk PR status checks**; see [Configure PR Checks](../../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md) for details.

<figure><img src="../../.gitbook/assets/Screenshot 2023-04-28 at 15.43.34.png" alt="Default Snyk test for pull requests setting enabled"><figcaption><p>Default Snyk test for pull requests setting enabled</p></figcaption></figure>

### **How to assign pull requests to users** <a href="#pr-assignment" id="pr-assignment"></a>

{% hint style="info" %}
**Feature availability**

The Auto-assign PRs feature is supported only for private repositories.
{% endhint %}

Snyk can automatically assign the pull requests it creates to ensure that they are handled by the right team members.

Auto-assign for PRs can be enabled for the GitHub and GitHub Enterprise integration and all Projects imported via GitHub, or on a per-Project basis.

Users can either be manually specified, and all will be assigned, or automatically selected based on the last commit user account.

#### **Enable Auto-assign for all Projects in the GitHub Enterprise integration**

To configure the Auto-assign settings for all the Projects from an imported private repository, go to the Github integration settings using Organization **Settings** > **Integrations** > **Source control** > **GitHub** and select **Enable pull request assignees**.&#x20;

You can then choose to assign PRs to the last user to change the manifest file or specified contributors.

<div align="left"><figure><img src="../../.gitbook/assets/image (7) (2) (1).png" alt="Auto-assign PRs in private repos"><figcaption><p>Auto-assign PRs in private repos</p></figcaption></figure></div>

{% hint style="info" %}
For pull request assignees, the option **The last user to change the manifest file** is based on blame data, not Git commits.
{% endhint %}

#### **Enable Auto-assign for a single Project**

To configure the Auto-assign settings for a specific Project from an imported private repository, follow these steps:

1. In the **Projects** tab for your Organization, select and expand the relevant private repository, select a Target, and click the **Settings** cog. This opens the Project page.
2. On the Project page, apply unique settings for that specific Project.\
   Select the **Settings** tab in the upper right and the **Github integration** \_\_ option in the left sidebar.
3. Go to the **Pull request assignees for private repos** section at the bottom of the page and choose to **Inherit from integration settings** or **Customize only for this Project**.
4. Ensure **Auto-assign PRs for this private Project** is enabled.
5. Choose to assign PRs to the last user to change the manifest file or named contributors.

<div align="center"><figure><img src="../../.gitbook/assets/image (5) (2) (1).png" alt="Auto-assign PRs for this private Project" width="375"><figcaption><p>Auto-assign PRs for this private Project</p></figcaption></figure></div>

### How to disconnect the GitHub Enterprise integration

{% hint style="warning" %}
Disconnecting the Snyk GitHub Enterprise integration halts all scans for imported repositories, PR checks cannot be executed, and Projects are deactivated in the Snyk Web UI.
{% endhint %}

1. Navigate to the Snyk GitHub Enterprise integration **Settings**.
2. At the bottom of the page, select **Remove GitHub Enterprise.**
3.  A confirmation modal opens. To proceed, select **Disconnect GitHub Enterprise**. \


    <figure><img src="../../.gitbook/assets/2023-11-09_17-38-28.png" alt="Confirm disconnecting from GitHub Enterprise" width="375"><figcaption><p>Confirm disconnecting from GitHub Enterprise</p></figcaption></figure>

After GitHub Enterprise is disconnected, imported Snyk Projects will be set to inactive, and you will no longer get alerts, pull requests, or Snyk tests on pull requests.

You can re-connect anytime; however, re-initiating GitHub Enterprise projects for monitoring requires setting up the integration again.

## Group level - Snyk Essentials integrations

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that is automatically synced and provides access to the [Integration Hub](../../getting-started/snyk-web-ui.md#manage-integrations-for-asset-discovery-asset-coverage-and-issues-from-third-party-vendors).

### GitHub setup guide for Snyk Essentials

{% hint style="info" %}
If you used GitHub Apps for your SCM integrations at the Snyk Organization level, Snyk Essentials requires an overview of your GitHub Organization. This means that the GitHub integration in Snyk Essentials uses an API token as an authentication method to onboard your GitHub Organization.&#x20;
{% endhint %}

#### Pulled entities by Snyk Essentials from GitHub <a href="#github-pulled-entities" id="github-pulled-entities"></a>

* Repositories
* Builds - only when using GitHub Actions.
* Scans - only when using Code security.

#### Prerequisites&#x20;

Ensure you meet all prerequisites listed on the [GitHub and GitHub Enterprise permission requirements](./#github-and-github-enterprise-permissions-requirements) page.

#### Integrate GitHub using Snyk Essentials <a href="#github-integrate-using-snyk-apprisk" id="github-integrate-using-snyk-apprisk"></a>

1. Profile name (`mandatory`): Input your integration profile name.&#x20;
2. Organizations (`mandatory`): Input the names of all the relevant GitHub organizations.

{% hint style="info" %}
If you have changed the name of your GitHub organization, copy the new name from the GitHub URL and paste it into the **GitHub Organizations** field in the Snyk Essentials Integration Hub.
{% endhint %}

3. Access Token (`mandatory`): Create your GitHub PAT from your GitHub Organization.&#x20;

* Generate your GitHub PAT by following the instructions in the [Generate a Personal access token from your GitHub settings](github-enterprise.md#generate-a-personal-access-token-from-your-github-settings) section.
* Authorize your GitHub PAT if you have configured SAML SSO. See the [How to authorize your Personal Access Token and enable SSO](github-enterprise.md#how-to-authorize-your-personal-access-token-and-enable-sso) page for more details.

4. Broker Token (`if your SCM is not publicly accessible`): Create and add your Broker token if you use Snyk Broker.
   * Generate your Broker token by following the instructions from the [Obtain your Broker token for Snyk Broker](../../enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md#obtain-your-broker-token-for-snyk-broker-code-agent) page.
   * Copy and paste the Broker token on the integration setup menu from the Integration Hub.
5. API URL (`mandatory`) - Input the API URL. The default URL is `https://api.github.com`.
6. Pull personal repositories (`optional`): Enable the option if you only want to pull the repositories you own.
7. Add Backstage Catalog (`optional`): If you want to add your Backstage catalog, follow the instructions from the [Backstage file for SCM Integrations](application-context-for-scm-integrations/) page.

{% hint style="warning" %}
If you enabled the **Pull personal repositories** option, only your personal repositories are pulled, not the public ones.&#x20;

If you want to pull data from both organization and personal repositories, then you must set up separate profiles.
{% endhint %}

#### Generate a Personal access token from your GitHub settings

1. Open GitHub and click the Settings menu for your profile.
2. Select Developer settings from the left sidebar.
3. Select Personal access tokens and then Tokens (classic).
4. Click Generate new token and, from the dropdown, select Generate new token (classic).
5. Add a description for your token in the Note field.
6. Select the required permissions:&#x20;
   * `repo`
   * `read:org`
   * `read:user`
   * `user:email`.
7. Click Generate token.
8. Copy and store the displayed key.

{% hint style="info" %}
Fine-grained personal access token is not supported.
{% endhint %}

#### API Version <a href="#github-api-version" id="github-api-version"></a>

You can use the[ GitHub REST API](https://docs.github.com/en/rest?apiVersion=2022-11-28) repository to access information about the API.

You can use as the host Address the IP/URL of the GitHub server. The default URL is [`https://api.github.com`](https://api.github.com).

The user associated with the token needs to have write permissions on relevant repositories to collect a breakdown of scan issues.

