# GitHub for Snyk Essentials

{% hint style="info" %}
If you used GitHub Apps for your SCM integrations at the Snyk Organization level, Snyk Essentials requires an overview of your GitHub Organization. This means that the GitHub integration in Snyk Essentials uses an API token as an authentication method to onboard your GitHub Organization.
{% endhint %}

## Pulled entities <a href="#github-pulled-entities" id="github-pulled-entities"></a>

* Repositories
* Builds - only when using GitHub Actions.
* Scans - only when using Code security.

Ensure you meet all prerequisites listed on the [GitHub and GitHub Enterprise permission requirements](../user-permissions-and-access-scopes.md#github-and-github-enterprise-permissions-requirements) page.

To configure a Group-level integration, you must be a Group Admin or have a custom role that includes the `Edit Snyk Essentials` permissions under the [Group-level permissions](../../../snyk-platform-administration/user-roles/pre-defined-roles.md#group-level-permissions).

## Integrate GitHub using Snyk Essentials: <a href="#github-integrate-using-snyk-apprisk" id="github-integrate-using-snyk-apprisk"></a>

1. Profile name (`mandatory`): Input your integration profile name.
2. Organizations (`mandatory`): Input the names of all the relevant GitHub organizations.

{% hint style="info" %}
If you have changed the name of your GitHub organization, copy the new name from the GitHub URL and paste it into the **GitHub Organizations** field in the Snyk Essentials Integration Hub.
{% endhint %}

3. Access Token (`mandatory`): Create your GitHub PAT from your GitHub Organization.

* Generate your GitHub PAT by following the instructions in the [Generate a Personal access token from your GitHub settings](github-for-snyk-essentials.md#generate-a-personal-access-token-from-your-github-settings) section.
* Authorize your GitHub PAT if you have configured SAML SSO. See the [How to authorize your Personal Access Token and enable SSO](../organization-level-integrations/github-enterprise.md#how-to-authorize-your-personal-access-token-and-enable-sso) page for more details.

4. Broker Token (`mandatory`): Create and add your Broker token if you use Snyk Broker.
   * Generate your Broker token by following the instructions from the [Obtain your Broker token for Snyk Broker](../../../enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md) page.
   * Copy and paste the Broker token on the integration setup menu from the Integration Hub.
5. API URL (`mandatory`) - Input the API URL. The default URL is `https://api.github.com`.
6. Pull personal repositories (`optional`): Enable the option if you only want to pull the repositories you own.
7. Add Backstage Catalog (`optional`): If you want to add your Backstage catalog, follow the instructions from the [Backstage file for SCM Integrations](../application-context-for-scm-integrations/) page.

{% hint style="warning" %}
If you enabled the **Pull personal repositories** option, only your personal repositories are pulled, not the public ones.

If you want to pull data from both organization and personal repositories, then you must set up separate profiles.
{% endhint %}

## Generate a Personal access token from your GitHub settings

1. Open GitHub and click the Settings menu for your profile.
2. Select Developer settings from the left sidebar.
3. Select Personal access tokens and then Tokens (classic).
4. Click Generate new token and, from the dropdown, select Generate new token (classic).
5. Add a description for your token in the Note field.
6. Select the required permissions:
   * `repo`
   * `read:org`
   * `read:user`
   * `user:email`.
7. Click Generate token.
8. Copy and store the displayed key.

{% hint style="info" %}
Fine-grained personal access token is not supported.
{% endhint %}

## API version

You can use the[ GitHub REST API](https://docs.github.com/en/rest?apiVersion=2022-11-28) repository to access information about the API.

You can use as the host Address the IP/URL of the GitHub server. The default URL is [`https://api.github.com`](https://api.github.com).

The user associated with the token needs to have write permissions on relevant repositories to collect a breakdown of scan issues.
