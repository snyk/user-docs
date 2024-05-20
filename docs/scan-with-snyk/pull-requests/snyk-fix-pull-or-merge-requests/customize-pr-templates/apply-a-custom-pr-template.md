# Apply a custom PR template

## Create and manage a custom PR template using the API

You can create a custom PR template using the API endpoint [Create or update pull request template for Group](https://apidocs.snyk.io/?#post-/groups/-group\_id-/settings/pull\_request\_template). Send an API request that contains a JSON payload with the custom properties. This request configures a Group-level pull request template that will be used on any Organization or Project within that Group. The pull request template created using the Snyk API can be updated at any time and all Projects in the Group are automatically updated with the latest changes.

API configuration of PR templates is available only at the Group level.

When a custom template is uploaded using an API request, all Snyk PRs in that Group adopt this format, effectively switching off the default Snyk template for the customizable properties. Strings are the only acceptable values; lists and numbers are not allowed.

If any customizable properties are missing from your custom template, Snyk reverts to the default values for these properties when opening a pull request.&#x20;

The following **properties are customizable using the API**:

* `title` - customize the PR title
* `commit_message` - customize the PR commit message
* `description` - customize the PR description&#x20;

You cannot customize the branch name for your PRs. The branch name of your PRs will use the Snyk default value.&#x20;

You can retrieve the custom PR template for your Group using the endpoint [Get pull request template for Group](https://apidocs.snyk.io/?#get-/groups/-group\_id-/settings/pull\_request\_template). This is useful if you want to consider changing your template and in troubleshooting.

To delete the template, use the endpoint [Delete pull request template for group](https://apidocs.snyk.io/?#delete-/groups/-group\_id-/settings/pull\_request\_template).

## Customize using a YAML PR template file

### Create the YAML file

Manually create the YAML template by using the [mustache](https://mustache.github.io) syntax for templating and add the file to your Project or repository.

When a custom template is uploaded to your Project, all PRs from Snyk for the Project adopt this format, effectively switching off the default Snyk template for the customized properties. Strings are the only acceptable values; lists and numbers are not allowed. If any customizable properties are missing from your template, Snyk reverts to the default values for these properties when opening a pull request.&#x20;

#### YAML multiline operators

You can use YAML multiline operators. You can create a detailed description that spans several lines by following this format:&#x20;

```yaml
description: |
  This pull request comes from Snyk
  For more information see the project page {{ snyk_project_url }}
  If you have more questions reach out to a member of the security team

```

The pipe operator preserves new line characters. Use greater than, `>` , to join all the lines by a space with a new line at the end. To use a colon, you can either use multiline operators, `|` or `>`, or enclose the line in double quotes:

```yaml
commitMessage: "snyk: this is a security pull request"
```

#### Customizable properties for YAML

The following properties are customizable:

* `title` - customize the PR title
* `commitMessage` - customize the PR commit message
* `description` - customize the PR description&#x20;

You cannot customize the branch name for your PRs. The branch name of your PRs will use the Snyk default value.&#x20;

### Use the YAML custom PR template

You can manually upload the YAML file with the name  `snyk_pull_request_template.yaml` to your [Project (repository)](#user-content-fn-1)[^1]. The method varies based on the type of integration.

* GitHub/ GitHub Enterprise - `/.github/snyk_pull_request_template.yaml`
* GitLab  - `/.gitlab/snyk_pull_request_template.yaml`
* Azure DevOps  - `/.azuredevops/snyk_pull_request_template.yaml`
* Other (such as BitBucket)  - `/.config/snyk_pull_request_template.yaml`

If you want to use a custom template for multiple repositories, add the YAML custom template file to each of these repositories.

## Broker configurations for fetching custom PR templates

If you use  [Snyk Broker](../../../../enterprise-configuration/snyk-broker/), you must allow access to these file locations in the `accept.json` configuration of your Broker client. The following describes the additional rules that should be added for each Git integration.

### GitHub and GitHub Enterprise

Under the list of `private` rules add the following input. For more information, see the Broker installation instructions for [GitHub](../../../../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/) and [GitHub Enterprise.](../../../../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/)

```json
{
  "//": "used to get custom pull request template",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/.github/snyk_pull_request_template.yaml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "used to get custom pull request template",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/.github%2Fsnyk_pull_request_template.yaml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
}
```

### Azure repositories

Under the list of `private` rules, add the following two elements to the existing `valid.values` array for file content. For more information, see the Broker installation instructions for [Azure repositories](../../../../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/).

```json
{
  "//": "get file content. restrict by file types",
  "method": "GET",
  "path": "/:owner/_apis/git/repositories/:repo/items",
  "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
  "valid": [
    {
      "queryParam": "path",
      "values": [
        "**/.azuredevops/snyk_pull_request_template.yaml",
        "**%2F.azuredevops%2Fsnyk_pull_request_template.yaml",
}
```

### BitBucket Server

Under the list of `private` rules, add the following input. For more information, see the Broker installation instructions for [Bitbucket Server/Data Center](../../../../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/).

```json
{
  "//": "used to get custom pull request template",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*/snyk_pull_request_template.yaml",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "used to get custom pull request template",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2Fsnyk_pull_request_template.yaml",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
}
```

### GitLab

Under the list of `private` rules, add the following input. For more information, the Broker installation instructions for [GitLab](../../../../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/).

```json
{
  "//": "used to get custom pull request template",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/snyk_pull_request_template.yaml",
  "origin": "https://${GITLAB}"
},
{
  "//": "used to get custom pull request template",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2Fsnyk_pull_request_template.yaml",
  "origin": "https://${GITLAB}"
}
```

If you use GitLab v3, add the following two elements to the existing `valid.values` array for file content:

```json
{
  "//": "used to determine the full dependency tree for v3 protocol",
  "method": "GET",
  "path": "/api/v3/projects/:project/repository/files",
  "origin": "https://${GITLAB}",
  "valid": [
    {
      "queryParam": "file_path",
      "values": [
        "**/.config/snyk_pull_request_template.yaml",
        "**%2F.config%2Fsnyk_pull_request_template.yaml",
}
```

[^1]: 
