# Creating organizations in Snyk

This page has instructions for creating organizations (Orgs) in Snyk:

* Generating the data to create organizations in Snyk
  * [GitHub](creating-orgs-in-snyk.md#github.com-github-enterprise)
  * [GitLab](creating-orgs-in-snyk.md#gitlab.com-hosted-gitlab)
  * [Bitbucket Server](creating-orgs-in-snyk.md#bitbucket-server)
  * [Bitbucket Cloud](creating-orgs-in-snyk.md#bitbucket-cloud)
  * [Azure](creating-orgs-in-snyk.md#azure)
* [Methods of creating Orgs](creating-orgs-in-snyk.md#methods-of-creating-orgs)
  * [via the API](creating-orgs-in-snyk.md#via-api)
  * [via the `orgs:create` util](creating-orgs-in-snyk.md#via-orgs-create-util)
* [Recommendations](creating-orgs-in-snyk.md#recommendations)

Before an import can begin you must set up Snyk with the organizations you will populate with projects.

It is recommended to have as many organizations in Snyk as you have in the source you are importing from. For GitHub this means mirroring the GitHub organizations in Snyk. The `snyk-api-import` tool provides a utility to use to make this simpler when using Groups and Organizations in Snyk.

## Generating the data required to create organizations in Snyk with `orgs:data` util

This util helps generate data needed to mirror the GitHub.com, GitHub Enterprise, GitLab, Bitbucket Server, or Bitbucket Cloud organization structure in Snyk. This is an opinionated util and will assume every organization in GitHub.com, GitHub Enterprise, GitLab, Bitbucket Server, or Bitbucket Cloud should become an organization in Snyk. If this is not what you are looking for, consider using the [Organizations API](https://snyk.docs.apiary.io/#reference/organizations/create-organization/create-a-new-organization) directly to create the structure you need.

### Options

```
  --source             The source of the targets to be imported
                       (for example, Github, Github Enterprise, Gitlab,
                       Bitbucket Server) [required].
  --groupId            Public id of the group in Snyk (available
                       on group settings) [required].
  --sourceUrl          Custom base url for the source API that can
                       list organizations (for example, GitHub Enterprise url).
  --sourceOrgPublicId  Public id of the organization in Snyk that
                       can be used as a template to copy all
                       supported organization settings.
  --skipEmptyOrgs      Skip organizations that have no targets 
                       (for example, Github Organizations that have no repos).
```

### GitHub.com and GitHub Enterprise

1. set the [GitHub.com personal access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token) as an environment variable: `export GITHUB_TOKEN=your_personal_access_token`
2. Run the command to generate organization data:

* **GitHub.com:** `snyk-api-import orgs:data --source=github --groupId=<snyk_group_id>`
* **GitHub Enterprise:** `snyk-api-import orgs:data --source=github-enterprise --groupId=<snyk_group_id> -- sourceUrl=https://ghe.custom.github.com/`

This creates the organization data in a file `group-<snyk_group_id>-github-<com|enterprise>-orgs.json`

### GitLab.com and Hosted GitLab

1. set the [GitLab personal access token](https://docs.gitlab.com/ee/user/profile/personal\_access\_tokens.html) as an environment variable: `export GITLAB_TOKEN=your_personal_access_token`
2. Run the command to generate organization data:

* **GitLab:** `snyk-api-import orgs:data --source=gitlab --groupId=<snyk_group_id>`
* **Hosted GitLab:** `snyk-api-import orgs:data --source=gitlab --groupId=<snyk_group_id> -- sourceUrl=https://gitlab.custom.com`

This creates the organization data in a file `group-<snyk_group_id>-gitlab-orgs.json`

### Bitbucket Server

**Note that Bitbucket Server is a hosted environment and you must provide the custom URL for your Bitbucket Server instance in the command.**

1. set the [Bitbucket Server access token](https://www.jetbrains.com/help/youtrack/standalone/integration-with-bitbucket-server.html#enable-youtrack-integration-bbserver) as an environment variable: `export BITBUCKET_SERVER_TOKEN=your_personal_access_token`
2. Run the command to generate organization data:

* `snyk-api-import orgs:data --source=bitbucket-server --groupId=<snyk_group_id> --sourceUrl=https://bitbucket-server.custom.com`

This creates the organization data in a file `group-<snyk_group_id>-bitbucket-server-orgs.json`

### Bitbucket Cloud

**Note that the URL for Bitbucket Cloud is https://bitbucket.org/**

1. set the Bitbucket Cloud Username and Password as environment variables: `export BITBUCKET_CLOUD_USERNAME=your_bitbucket_cloud_username` and `export BITBUCKET_CLOUD_PASSWORD=your_bitbucket_cloud_password`
2. Run the command to generate organization data:

* `snyk-api-import orgs:data --source=bitbucket-cloud --groupId=<snyk_group_id>`

This creates the organization data in a file `group-<snyk_group_id>-bitbucket-cloud-orgs.json`

### Azure

**Note that for Azure, this step needs to be done manually** Since Azure has no API call for getting the Azure Organizations, the Orgs file must be created manually for the next commands to run:

The file should be formatted this way:

```
{
   "orgs":[
      {
         "name":"THE_NAME_OF_AN_AZURE_ORG",
         "groupId":"YOUR_SNYK_GROUP_ID",
         "sourceOrgId":"THE_SNYK_ORG_ID_FROM_WHICH_TO_COPY_THE_SETTINGS_FROM"   // **optional**
      },
      {
         "name":"THE_NAME_OF_ANOTHER_AZURE_ORG",
         "groupId":"YOUR_SNYK_GROUP_ID",
         "sourceOrgId":"THE_SNYK_ORG_ID_FROM_WHICH_TO_COPY_THE_SETTINGS_FROM"  // **optional**
      }
   ]
}
```

Once the file is created, you can feed it to the [orgs:create command](https://github.com/snyk-tech-services/snyk-api-import/blob/0e5162d29dec7f1d5acde247cc8da0553871db3f/docs/orgs.md#creating-organizations-in-snyk-1)

## Methods of creating Orgs

Use the generated data file to help create the organizations via API or use the provided util.

### via API

Use the generated data to feed into the Snyk [Orgs API](https://snyk.docs.apiary.io/#reference/groups/organizations-in-a-group/create-a-new-organization-in-a-group) to generate the organizations within a group.

### via `orgs:create` util

1. set the `SNYK_TOKEN` environment variable - your [Snyk api token](https://app.snyk.io/account)
2. Run the command to create Orgs: `snyk-api-import orgs:create --noDuplicateNames --includeExistingOrgsInOutput --file=group-<snyk_group_id>-github-<com|enterprise>-orgs.json`
3. Use the `noDuplicateNames` flag (optional) to skip creating an organization if the given name is already taken within the Group.
4. Use the `includeExistingOrgsInOutput` flag (optional, default is "true") to log information for existing organizations as well as newly created Orgs. To set this flag as false, use "--no-includeExistingOrgsInOutput" in the command as follows: `snyk-api-import orgs:create --no-includeExistingOrgsInOutput --file=group-<snyk_group_id>-github-<com|enterprise>-orgs.json`

The file format required for this command is as follows:

```
"orgs": [
  {
    "groupId": "<public_snyk_group_id>",
    "name": "<name_of_the_organization>",
    "sourceOrgId": "<public_snyk_organization_id>"
  }
]
```

* `groupId` - public id of the Snyk Group where the organization is to be created
* `name` - name to use when creating the organization
* `sourceOrgId` - **optional** public id of a Snyk organization to copy settings from

### Recommendations

* Have [notifications disabled](https://snyk.docs.apiary.io/#reference/organizations/notification-settings/set-notification-settings) for emails and so on to avoid receiving import notifications.
* Have the [fix PRs and PR checks disabled](https://snyk.docs.apiary.io/#reference/integrations/integration-settings/update) until import is complete to avoid sending extra requests to SCMs (GitHub, GitLab, Bitbucket, other).
