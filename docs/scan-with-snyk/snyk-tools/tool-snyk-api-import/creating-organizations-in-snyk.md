# Creating Organizations in Snyk

This page has instructions for creating Organizations (Orgs) in Snyk:

* Generating the data to create organizations in Snyk
  * [GitHub](creating-organizations-in-snyk.md#github.com-and-github-enterprise)
  * [GitLab](creating-organizations-in-snyk.md#gitlab.com-and-hosted-gitlab)
  * [Bitbucket Server](creating-organizations-in-snyk.md#bitbucket-server)
  * [Bitbucket Cloud](creating-organizations-in-snyk.md#bitbucket-cloud)
  * [Azure](creating-organizations-in-snyk.md#azure)
* [Methods of creating Organizations](creating-organizations-in-snyk.md#methods-of-creating-organizations)
  * [Using the AP](creating-organizations-in-snyk.md#using-the-api)I
  * [Using the `orgs:create` utility](creating-organizations-in-snyk.md#using-the-orgs-create-utility)
* [Recommendations](creating-organizations-in-snyk.md#recommendations)

Before an import can begin you must set up Snyk with the Organizations you will populate with Projects.

Snyk recommends that you have as many Organizations in Snyk as you have in the source you are importing from. For GitHub, this means mirroring the GitHub organizations in Snyk. The `snyk-api-import` tool provides a utility to use to make this simpler when using Groups and Organizations in Snyk.

## Generating the data required to create Organizations in Snyk with the `orgs:data` utility

This utility helps generate data needed to mirror the GitHub.com, GitHub Enterprise, GitLab, Bitbucket Server, or Bitbucket Cloud organization structure in Snyk. This opinionated utility will assume every organization in GitHub.com, GitHub Enterprise, GitLab, Bitbucket Server, or Bitbucket Cloud should become an Organization in Snyk. If this is not what you are looking for, consider using the API endpoint [Create a new organization](../../../snyk-api/reference/organizations-v1.md#org) directly to create the structure you need.

### Options

```
  --source             The source of the targets to be imported
                       (for example, Github, Github Enterprise, Gitlab,
                       Bitbucket Server) [required].
  --groupId            Public id of the group in Snyk (available
                       on Group settings) [required].
  --sourceUrl          Custom base url for the source API that can
                       list organizations (for example, GitHub Enterprise url).
  --sourceOrgPublicId  Public id of the Organization in Snyk that
                       can be used as a template to copy all
                       supported Organization settings.
  --skipEmptyOrgs      Skip organizations that have no targets 
                       (for example, Github organizations that have no repos).
```

### GitHub.com and GitHub Enterprise

1. Set the [GitHub.com personal access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token) as an environment variable: `export GITHUB_TOKEN=your_personal_access_token`
2. Run the command to generate organization data:

* GitHub.com: `snyk-api-import orgs:data --source=github --groupId=<snyk_group_id>`
* GitHub Enterprise: `snyk-api-import orgs:data --source=github-enterprise --groupId=<snyk_group_id> -- sourceUrl=https://ghe.custom.github.com/`

This creates the organization data in a file `group-<snyk_group_id>-github-<com|enterprise>-orgs.json`

### GitLab.com and Hosted GitLab

1. Set the [GitLab personal access token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html) as an environment variable: `export GITLAB_TOKEN=your_personal_access_token`
2. Run the command to generate organization data:

* **GitLab:** `snyk-api-import orgs:data --source=gitlab --groupId=<snyk_group_id>`
* Hosted GitLab: `snyk-api-import orgs:data --source=gitlab --groupId=<snyk_group_id> -- sourceUrl=https://gitlab.custom.com`

This creates the organization data in a file `group-<snyk_group_id>-gitlab-orgs.json`

### Bitbucket Server

{% hint style="info" %}
Bitbucket Server is a hosted environment, and you must provide the custom URL for your Bitbucket Server instance in the command.
{% endhint %}

1. Set the [Bitbucket Server access token](https://www.jetbrains.com/help/youtrack/standalone/integration-with-bitbucket-server.html#enable-youtrack-integration-bbserver) as an environment variable: `export BITBUCKET_SERVER_TOKEN=your_personal_access_token`
2. Run the command to generate organization data:\
   `snyk-api-import orgs:data --source=bitbucket-server --groupId=<snyk_group_id> --sourceUrl=https://bitbucket-server.custom.com`

This creates the organization data in a file `group-<snyk_group_id>-bitbucket-server-orgs.json`

### Bitbucket Cloud

{% hint style="info" %}
The URL for Bitbucket Cloud is https://bitbucket.org/
{% endhint %}

1. Set the Bitbucket Cloud Username and Password as environment variables: `export BITBUCKET_CLOUD_USERNAME=your_bitbucket_cloud_username` and `export BITBUCKET_CLOUD_PASSWORD=your_bitbucket_cloud_password`
2. Run the command to generate organization data:\
   `snyk-api-import orgs:data --source=bitbucket-cloud --groupId=<snyk_group_id>`

This creates the organization data in a file `group-<snyk_group_id>-bitbucket-cloud-orgs.json`

### Azure

For Azure, this step must be done manually. Since Azure has no API call for getting the Azure Organizations, the Orgs file must be created manually for the next commands to run. The file should be formatted this way:

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

After the file is created, you can feed it to the [`orgs:create` command](https://github.com/snyk/snyk-api-import/blob/0e5162d29dec7f1d5acde247cc8da0553871db3f/docs/orgs.md#creating-organizations-in-snyk-1).

## Methods of creating Organizations

Use the generated data file to help create the Organizations using the API endpoint [Create a new Organization](../../../snyk-api/reference/organizations-v1.md#org) or use the provided utility.

### Using the API

Use the generated data to feed into the endpoint [Create a new Organization](../../../snyk-api/reference/organizations-v1.md#org) to generate the Organizations within a Group.

### Using the `orgs:create` utility

1. Set the `SNYK_TOKEN` environment variable, your [Snyk API token](https://app.snyk.io/account).
2. Run the command to create Organizations: `snyk-api-import orgs:create --noDuplicateNames --includeExistingOrgsInOutput --file=group-<snyk_group_id>-github-<com|enterprise>-orgs.json`.
3. Use the `noDuplicateNames` flag (optional) to skip creating an Organization if the given name is already taken within the Group.
4. Use the `includeExistingOrgsInOutput` flag (optional, default is `true`) to log information for existing Organizations as well as newly created Organizations. To set this flag as `false`, use --no-includeExistingOrgsInOutput in the command as follows: `snyk-api-import orgs:create --no-includeExistingOrgsInOutput --file=group-<snyk_group_id>-github-<com|enterprise>-orgs.json`

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

* `groupId` - public id of the Snyk Group where the Organization is to be created
* `name` - name to use when creating the Organization
* `sourceOrgId` - **optional** public id of a Snyk Organization to copy settings from

### Recommendations

* Use the endpoint [Set notification settings](../../../snyk-api/reference/organizations-v1.md#org-orgid-notification-settings) to disable notifications for emails and so on to avoid receiving import notifications.
* Use the [Update](../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-integrationid-settings) (integration settings) endpoint to disable the fix PRs and PR checks until import is complete to avoid sending extra requests to SCMs (GitHub, GitLab, Bitbucket, and so on).
