# Creating orgs in Snyk

## Creating Organizations in Snyk

### Table of Contents

* Generating the data to create Organizations in Snyk
  * [Github](creating-orgs-in-snyk.md#github.com-github-enterprise)
  * [Gitlab](creating-orgs-in-snyk.md#gitlab.com-hosted-gitlab)
  * [Bitbucket-Server](creating-orgs-in-snyk.md#bitbucket-server)
  * [Bitbucket-cloud](creating-orgs-in-snyk.md#bitbucket-cloud)
* [Creating the organizations](creating-orgs-in-snyk.md#creating-organizations-in-snyk-1)
  * [via the API](creating-orgs-in-snyk.md#via-api)
  * [via the `orgs:create` util](creating-orgs-in-snyk.md#via-orgs-create-util)
* [Recommendations](creating-orgs-in-snyk.md#recommendations)

Before an import can begin Snyk needs to be setup with the Organizations you will populate with projects.

It is recommended to have as many Organizations in Snyk as you have in the source you are importing from. So for Github this would mean mirroring the Github organizations in Snyk. The tool provides a utility that can be used to make this simpler when using Groups & Organizations in Snyk.

## Generating the data required to create Organizations in Snyk with `orgs:data` util

This util helps generate data needed to mirror the Github.com / Github Enterprise / Gitlab / Bitbucket Server / Bitbucket Cloud organization structure in Snyk. This is an opinionated util and will assume every organization in Github.com / Github Enterprise / Gitlab / Bitbucket Server / Bitbucket Cloud should become an organization in Snyk. If this is not what you are looking for, please look at using the [Organizations API](https://snyk.docs.apiary.io/#reference/organizations/create-organization/create-a-new-organization) directly to create the structure you need.

### Options

```
  --source             The source of the targets to be imported
                       (e.g. Github, Github Enterprise, Gitlab,
                       Bitbucket Server)                [required]
  --groupId            Public id of the group in Snyk (available
                       on group settings)               [required]
  --sourceUrl          Custom base url for the source API that can
                       list organizations (e.g. Github Enterprise url)
  --sourceOrgPublicId  Public id of the organization in Snyk that
                       can be used as a template to copy all
                       supported organization settings.
  --skipEmptyOrgs      Skip organizations that have no targets. 
                       (e.g. Github Organizations that have no repos)
```

### Github.com / Github Enterprise

1. set the [Github.com personal access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token) as an environment variable: `export GITHUB_TOKEN=your_personal_access_token`
2. Run the command to generate organization data:

* **Github.com:** `snyk-api-import orgs:data --source=github --groupId=<snyk_group_id>`
* **Github Enterprise:** `snyk-api-import orgs:data --source=github-enterprise --groupId=<snyk_group_id> -- sourceUrl=https://ghe.custom.github.com/`

This will create the organization data in a file `group-<snyk_group_id>-github-<com|enterprise>-orgs.json`

### Gitlab.com / Hosted Gitlab

1. set the [Gitlab personal access token](https://docs.gitlab.com/ee/user/profile/personal\_access\_tokens.html) as an environment variable: `export GITLAB_TOKEN=your_personal_access_token`
2. Run the command to generate organization data:

* **Gitlab:** `snyk-api-import orgs:data --source=gitlab --groupId=<snyk_group_id>`
* **Hosted Gitlab:** `snyk-api-import orgs:data --source=gitlab --groupId=<snyk_group_id> -- sourceUrl=https://gitlab.custom.com`

This will create the organization data in a file `group-<snyk_group_id>-gitlab-orgs.json`

### Bitbucket Server

**Please note that Bitbucket Server is a hosted environment and you must provide the custom URL for your Bitbucket Server instance in the command**

1. set the [Bitbucket Server access token](https://www.jetbrains.com/help/youtrack/standalone/integration-with-bitbucket-server.html#enable-youtrack-integration-bbserver) as an environment variable: `export BITBUCKET_SERVER_TOKEN=your_personal_access_token`
2. Run the command to generate organization data:

* `snyk-api-import orgs:data --source=bitbucket-server --groupId=<snyk_group_id> --sourceUrl=https://bitbucket-server.custom.com`

This will create the organization data in a file `group-<snyk_group_id>-bitbucket-server-orgs.json`

### Bitbucket Cloud

**Note that the URL for Bitbucket Cloud is https://bitbucket.org/**

1. set the Bitbucket Cloud Username and Password as an environment variables: `export BITBUCKET_CLOUD_USERNAME=your_bitbucket_cloud_username` and `export BITBUCKET_CLOUD_PASSWORD=your_bitbucket_cloud_password`
2. Run the command to generate organization data:

* `snyk-api-import orgs:data --source=bitbucket-cloud --groupId=<snyk_group_id>`

This will create the organization data in a file `group-<snyk_group_id>-bitbucket-cloud-orgs.json`

## Creating Organizations in Snyk

Use the generated data file to help create the organizations via API or use the provided util.

### via API

Use the generated data to feed into Snyk [Orgs API](https://snyk.docs.apiary.io/#reference/groups/organizations-in-a-group/create-a-new-organization-in-a-group) to generate the organizations within a group.

### via `orgs:create` util

1. set the `SNYK_TOKEN` environment variable - your [Snyk api token](https://app.snyk.io/account)
2. Run the command to create Orgs: `snyk-api-import orgs:create --noDuplicateNames --includeExistingOrgsInOutput --file=group-<snyk_group_id>-github-<com|enterprise>-orgs.json`

* Using the `noDuplicateNames` flag (optional) will Skip creating an organization if the given name is already taken within the Group.
* Using the `includeExistingOrgsInOutput` flag (optional, default is "true") will Log existing organization information as well as newly created. To set this flag as false, please use "--no-includeExistingOrgsInOutput" in the command, like so: `snyk-api-import orgs:create --no-includeExistingOrgsInOutput --file=group-<snyk_group_id>-github-<com|enterprise>-orgs.json`

The file format required for this looks like so:

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

* have [notifications disabled](https://snyk.docs.apiary.io/#reference/organizations/notification-settings/set-notification-settings) for emails etc to avoid receiving import notifications
* have the [fix PRs and PR checks disabled](https://snyk.docs.apiary.io/#reference/integrations/integration-settings/update) until import is complete to avoid sending extra requests to SCMs (Github/Gitlab/Bitbucket etc)
