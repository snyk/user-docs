# GitLab for Snyk Essentials

The Integrations page shows all active integrations, including data from your existing Snyk Organizations that is automatically synced and provides access to the Integration Hub.

### GitLab setup guide

#### Pulled entities <a href="#gitlab-pulled-entities" id="gitlab-pulled-entities"></a>

* Users
* Repositories

#### Prerequisites <a href="#gitlab-integrate-using-snyk-apprisk" id="gitlab-integrate-using-snyk-apprisk"></a>

To configure a Group-level integration, you must be a Group Admin or have a custom role that includes the `Edit Snyk Essentials` permissions under the [Group-level permissions](../../../snyk-platform-administration/user-roles/pre-defined-roles.md#group-level-permissions).

#### Integrate using Snyk Essentials <a href="#gitlab-integrate-using-snyk-apprisk" id="gitlab-integrate-using-snyk-apprisk"></a>

1. Profile name (`mandatory`): Input your integration profile name.
2. Access Token (`mandatory`):
   * API Token (`mandatory`): Create your GitLab PAT from your GitLab organization. Follow the instructions in [Generate a Personal access token from your GitLab settings section](gitlab-for-snyk-essentials.md#generate-a-personal-access-token-from-your-gitlab-settings). Authorize your personal access token if you have configured SAML SSO.
   * Host URL (`mandatory`): The IP/URL of the GitLab server. The default URL is [`https://gitlab.com`](https://gitlab.com)
3. Broker Token (`mandatory`): Create and add your Broker token if you use Snyk Broker.
   * Generate your Broker token by following the instructions from the [Obtain your Broker token for Snyk Broker](../../../implementation-and-setup/enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md) page.
   * Copy and paste the Broker token on the integration setup menu from the Integration Hub.
4. Pull personal repositories (`optional`): Enable the option If you only want to pull the repositories you own.
5. Add Backstage Catalog (`optional`): If you want to add your Backstage catalog, follow the instructions from the [Backstage file for SCM Integrations](../application-context-for-scm-integrations/) page.

#### Generate a [Personal access token](../organization-level-integrations/gitlab.md#gitlab-access-tokens) from your GitLab settings

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

#### API version <a href="#gitlab-api-version" id="gitlab-api-version"></a>

You can use the[ GitLab REST API v4](https://docs.gitlab.com/ee/api/index.html) repository to access information about the API.
