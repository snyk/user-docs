# Creating import targets data for import

Use the `import:data` util to help generate the import json data needed by the import command. Follow the instructions on this page:

* Usage for the `import:data` util
  * [Github](creating-import-targets-data-for-import.md#github.com-github-enterprise)
  * [Gitlab](creating-import-targets-data-for-import.md#gitlab.com-hosted-gitlab)
  * [Azure Repos](creating-import-targets-data-for-import.md#dev.azure.com-hosted-azure)
  * [Bitbucket Server](creating-import-targets-data-for-import.md#bitbucket-server)
  * [Bitbucket Cloud](creating-import-targets-data-for-import.md#bitbucket-cloud)
* [Recommendations](creating-import-targets-data-for-import.md#recommendations)

In the section that applies to your SCM, review the information about creating the organizations data in json at the beginning and then follow the steps. **You must set your SCM personal access token or username and password credentials as environment variables.**

Note that archived repos are excluded by default.

## GitHub.com and GitHub Enterprise

You need the organizations data in json as an input to this command to help map Snyk organization IDs and Integration Ids that must be used during import against individual targets to be imported. The following format is required:

```
{
  "orgData": [
    {
      "name": "<org_name_in_gh_used_to_list_repos>",
      "orgId": "<snyk_org_id>",
      "integrations": {
        "github": "<snyk_org_integration_id>",
        "github-enterprise": "<snyk_org_integration_id>
      },
    },
    {...}
  ]
}
```

Note: the "name" of the GitHub or GitHub Enterprise organization is required in order to list all repos belonging to that organization via the GitHub API. The Snyk-specific data accompanying that organization name will be used as the information to generate import data assuming all repos in that organization will be imported into a given Snyk organization. This util is opinionated. If you want to customize the import data, create it manually as described on [Kicking off an import](kicking-off-an-import.md).

* GitHub and GitHub Enterprise organizations can be listed using the [/orgs API](https://docs.github.com/en/free-pro-team@latest/rest/reference/orgs)
* Integrations can be listed via [Snyk Integrations API](https://snyk.docs.apiary.io/#reference/integrations/integrations/list)
* All organization IDs can be found by listing all organizations a group admin belongs to via [Snyk Organizations API](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)

1. Set the [GitHub.com personal access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token) as an environment variable: `export GITHUB_TOKEN=your_personal_access_token`
2. Create the organizations data in json as described in the opening paragraphs.
3. Run the command to generate import data:\
   **GitHub.com:** `DEBUG=snyk* GITHUB_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=github --integrationType=github`\
   \`\`**GitHub Enterprise:** `DEBUG=snyk* GITHUB_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=github-enterprise --integrationType=github-enterprise --sourceUrl=https://ghe.custom.com`
4. Use the generated data to feed into the `import` command to [kick off the import](kicking-off-an-import.md).

## GitLab.com and Hosted GitLab

You need the organizations data in json as an input to this command to help map Snyk organization IDs and Integration Ids that must be used during import against individual targets to be imported. The following format is required:

```
{
  "orgData": [
    {
      "name": "<group_name_in_gitlab_used_to_list_repos>",
      "orgId": "<snyk_org_id>",
      "integrations": {
        "gitlab": "<snyk_org_integration_id>",
      },
    },
    {...}
  ]
}
```

Note: the "name" of the GitLab Group is required in order to list all projects belonging to that Group via the GitLab API. The Snyk-specific data accompanying that Group name will be used as the information to generate import data assuming all projects in that Group will be imported into a given Snyk organization. This is util opinionated. If you want to customize the import data, create it manually as described on [Kicking off an import](kicking-off-an-import.md).

* GitLab Groups can be listed using the [/groups API](https://docs.gitlab.com/ee/api/groups.html)
* Integrations can be listed via [Snyk Integrations API](https://snyk.docs.apiary.io/#reference/integrations/integrations/list)
* All organization IDs can be found by listing all organizations a group admin belongs to via [Snyk Organizations API](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)

1. Set the [GitLab personal access token](https://docs.gitlab.com/ee/user/profile/personal\_access\_tokens.html) as an environment variable: `export GITLAB_TOKEN=your_personal_access_token`
2. Create the organizations data in json as described in the opening paragraphs.
3. Run the command to generate import data:\
   **Gitlab.com:** `DEBUG=snyk* GITLAB_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=gitlab --integrationType=gitlab`\
   \`\`**Hosted Gitlab:** `DEBUG=snyk* GITLAB_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=gitlab --integrationType=gitlab --sourceUrl=https://gitlab.custom.com`
4. Use the generated data to feed into the `import` command to [kick off the import](kicking-off-an-import.md).

## dev.azure.com and Hosted Azure

**Note: This tool uses Azure API v1.**

You need the organizations data in json as an input to this command to help map Snyk organization IDs and Integration Ids that must be used during import against individual targets to be imported. The following format is required:

```
{
  "orgData": [
    {
      "name": "<org_name_in_azure_used_to_list_repos>",
      "orgId": "<snyk_org_id>",
      "integrations": {
        "azure-repos": "<snyk_org_integration_id>",
      },
    },
    {...}
  ]
}
```

Note: the "name" of the Azure organization is required in order to list all projects and repos belonging to that organization via the Azure API. The Snyk-specific data accompanying that organization name will be used as the information to generate import data assuming all projects in that organization will be imported into a given Snyk organization. This util opinionated. If you want to customize the import data, create it manually as described on [Kicking off an import](kicking-off-an-import.md).

* Your Org name in Azure is listed on the left pane in the [Azure-Devops-site](https://dev.azure.com/)
* Integrations can be listed via [Snyk Integrations API](https://snyk.docs.apiary.io/#reference/integrations/integrations/list)
* All organization IDs can be found by listing all organizations a group admin belongs to via [Snyk Organizations API](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)

1. Set the [Azure personal access token](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops\&tabs=preview-page) as an environment variable: `export AZURE_TOKEN=your_personal_access_token`
2. Create the organizations data in json as described in the opening paragraphs.
3. Run the command to generate import data:\
   **dev.azure.com:** `DEBUG=snyk* AZURE_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=azure-repos --integrationType=azure-repos`\
   \`\`**Hosted Azure:** `DEBUG=snyk* AZURE_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=azure-repos --integrationType=azure-repos --sourceUrl=https://azure.custom.com`
4. Use the generated data to feed into the `import` command to [kick off the import](kicking-off-an-import.md).

## Bitbucket Server

**Note: This tool uses Bitbucket server API version 1.0.**

You need the organizations data in json as an input to this command to help map Snyk organization IDs and Integration Ids that must be used during import against individual targets to be imported. The following format is required:

```
{
  "orgData": [
    {
      "name": "<project_name_in_bitbucket_server_used_to_list_repos>",
      "orgId": "<snyk_org_id>",
      "integrations": {
        "bitbucket-server": "<snyk_org_integration_id>",
      },
    },
    {...}
  ]
}
```

Note: the "name" of the Bitbucket server project is required in order to list all repos belonging to that project via the Bitbucket server API. The Snyk-specific data accompanying that project name will be used as the information to generate import data assuming all repos in that project will be imported into a given Snyk organization. If you want to customize the import data, create it manually as described on [Kicking off an import](kicking-off-an-import.md).

* Bitbucket Server Projects can be listed using the [/projects API](https://docs.atlassian.com/bitbucket-server/rest/7.19.2/bitbucket-rest.html)
* Integrations can be listed via [Snyk Integrations API](https://snyk.docs.apiary.io/#reference/integrations/integrations/list)
* All organization IDs can be found by listing all organizations a group admin belongs to via [Snyk Organizations API](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)

1. Set the [Bitbucket Server personal access token](https://www.jetbrains.com/help/youtrack/standalone/integration-with-bitbucket-server.html#enable-youtrack-integration-bbserver) as an environment variable: `export BITBUCKET_SERVER_TOKEN=your_personal_access_token`
2. Create the organizations data in json as described in the opening paragraphs.
3. Run the command to generate import data:\
   **Bitbucket Server:** `DEBUG=snyk* BITBUCKET_SERVER_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=bitbucket-server --integrationType=bitbucket-server --sourceUrl=https://bitbucket-server.dev.example.com`
4. Use the generated data to feed into the `import` command to [kick off the import](kicking-off-an-import.md).

## Bitbucket Cloud

**Note that this tool uses Bitbucket Cloud API version 2.0.**

You need the organizations data in json as an input to this command to help map Snyk organization IDs and Integration Ids that must be used during import against individual targets to be imported. The following format is required:

```
{
  "orgData": [
    {
      "name": "<workspace_name_in_bitbucket_cloud_used_to_list_repos>",
      "orgId": "<snyk_org_id>",
      "integrations": {
        "bitbucket-cloud": "<snyk_org_integration_id>",
      },
    },
    {...}
  ]
}
```

Note: the "name" of the Bitbucket Cloud workspace is required in order to list all repos belonging to that workspace via the Bitbucket Cloud API. The Snyk-specific data accompanying that workspace name will be used as the information to generate import data assuming all repos in that workspace will be imported into a given Snyk organization. This util opinionated. If you want to customize the import data, create it manually as described on [Kicking off an import](kicking-off-an-import.md).

* Bitbucket Cloud workspaces can be listed using the [/workspaces API](https://bitbucket.org/api/2.0/workspaces)
* Integrations can be listed via [Snyk Integrations API](https://snyk.docs.apiary.io/#reference/integrations/integrations/list)
* All organization IDs can be found by listing all organizations a group admin belongs to via [Snyk Organizations API](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)

1. Set the username and password credentials for your Bitbucket Cloud as environment variables:\
   `export BITBUCKET_CLOUD_USERNAME=your_bitbucket_cloud_username`\
   `export BITBUCKET_CLOUD_PASSWORD=your_bitbucket_cloud_password`
2. Create the organizations data in json as described in the opening paragraphs.
3. Run the command to generate import data:\
   **Bitbucket Cloud:** `DEBUG=snyk* BITBUCKET_CLOUD_USERNAME=*** BITBUCKET_CLOUD_PASSWORD=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=bitbucket-cloud --integrationType=bitbucket-cloud`
4. Use the generated data to feed into the `import` command to [kick off the import](kicking-off-an-import.md).

## Recommendations

* Have [notifications disabled](https://snyk.docs.apiary.io/#reference/organizations/notification-settings/set-notification-settings) for emails and so on to avoid receiving import notifications.
* Have the [fix PRs and PR checks disabled](https://snyk.docs.apiary.io/#reference/integrations/integration-settings/update) until import is complete to avoid sending extra requests to SCMs (GitHub, GitLab, Bitbucket, other).
