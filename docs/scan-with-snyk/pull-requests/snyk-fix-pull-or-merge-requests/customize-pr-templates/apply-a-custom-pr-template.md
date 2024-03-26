# Apply a Custom PR template

## Create an API request with a Custom template

{% hint style="info" %}
Ensure that your account manager has turned on the feature flag and you have enabled the feature in the Snyk Preview settings.&#x20;
{% endhint %}

You can create a custom PR template by sending an [API request](https://apidocs.snyk.io/?version=2023-10-13%7Ebeta#tag--Pull-Request-Templates) containing a JSON payload with the customized properties. This request configures a Group level pull request template that will be used on any Organization or Project within that Group. The Snyk API pull request template can be updated at any time and all Projects in that Group are automatically updated with the latest changes.

{% hint style="info" %}
When the Customize Snyk PRs feature is enabled, all PRs from Snyk adopt this format, effectively switching off the default Snyk template. The only acceptable values are strings. Lists and numbers are not allowed.
{% endhint %}

{% hint style="info" %}
The API configuration is momentarily available only at Group level.
{% endhint %}

#### Customizable properties for API

The following properties are customizable:

* `title` - customize the PR title
* `commit_message` - customize the PR commit message
* `description` - customize the PR description&#x20;

{% hint style="info" %}
If any customizable properties are missing from your template, Snyk reverts to the default values for these properties when opening a pull request.&#x20;
{% endhint %}

## Customize and use a YAML PR template

### Customize the YAML file

[Manually create the YAML template](apply-a-custom-pr-template.md#manually-customize-the-pr-template) by using the [mustache](https://mustache.github.io) syntax for templating and add the file to your Project or repository.

{% hint style="info" %}
When the Customize Snyk PRs feature is enabled, all PRs from Snyk adopt this format, effectively switching off the default Snyk template. The only acceptable values are strings. Lists and numbers are not allowed.
{% endhint %}

#### YAML multiline operators

You can use YAML multiline operators. You can create a detailed description that spans several lines by following this format:&#x20;

```yaml
description: |
  This pull request comes from Snyk
  For more information see the project page {{ snyk_project_url }}
  If you have more questions reach out to a member of the security team

```

The pipe operator preserves new line characters. Use greater than, `>` , to join all the lines by a space with a new line on the end. To use a colon, you can either use multiline operators, `|` or `>`, or enclose the line in double quotes:

```yaml
commitMessage: "snyk: this is a security pull request"
```

#### Customizable properties for YAML

The following properties are customizable:

* `title` - customize the PR title
* `commitMessage` - customize the PR commit message
* `description` - customize the PR description&#x20;

{% hint style="info" %}
If any customizable properties are missing from your template, Snyk reverts to the default values for these properties when opening a pull request.&#x20;
{% endhint %}

### Use the YAML Custom PR template

{% hint style="info" %}
Ensure that your account manager has turned on the feature flag and you have enabled the feature in the Snyk Preview settings.&#x20;
{% endhint %}

You can manually upload the YAML file with the name  `snyk_pull_request_template.yaml` to your Project(repository)[^1]. The method varies based on the type of integration.

* Github/ GitHub Enterprise - `/.github/snyk_pull_request_template.yaml`
* GitLab  - `/.gitlab/snyk_pull_request_template.yaml`
* Azure  - `/.azuredevops/snyk_pull_request_template.yaml`
* Other (like BitBucket)  - `/.config/snyk_pull_request_template.yaml`

{% hint style="info" %}
If you want to use multiple repositories for a custom template, add the customized YAML template file to each of these repositories.
{% endhint %}

## Broker configurations for fetching PR templates

If you use  [Snyk Broker](../../../../getting-started-with-the-snyk-enterprise-plan/snyk-broker/), you must allow access to these file locations in the `accept.json` [configuration of your Broker client](../../../../getting-started-with-the-snyk-enterprise-plan/snyk-broker/snyk-broker-code-agent/install-snyk-broker-code-agent-using-docker/set-up-the-broker-client/run-the-broker-client-with-the-code-snippets-display.md). The following describes the additional rules that should be added for each git integration.

### Github and GitHub Enterprise

Under the list of `private` rules add the following input. For more information, see [Configure Broker to be used for GitHub Enterprise](../../../../getting-started-with-the-snyk-enterprise-plan/snyk-broker/install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-github-enterprise.md#configure-broker-to-be-used-for-github-enterprise) in the instructions for installing Broker for GitHub Enterprise.

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

### Azure repos

Under the list of `private` rules, add the following two elements to the existing `valid.values` array for file content. For more information, see [Configure Broker to be used with Azure Repos](../../../../getting-started-with-the-snyk-enterprise-plan/snyk-broker/install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md#configure-broker-to-be-used-with-azure-repos) in the instructions for installing Broker for Azure repos.

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

Under the list of `private` rules, add the following input. For more information, see [Bitbucket Server/Data Center - environment variables for Snyk Broker](../../../../getting-started-with-the-snyk-enterprise-plan/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker.md) in the instructions for installing Broker for BitBucket Server.

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

Under the list of `private` rules, add the following input. For more information, see [GitLab - environment variables for Snyk Broker](../../../../getting-started-with-the-snyk-enterprise-plan/snyk-broker/install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-environment-variables-for-snyk-broker.md) in the instructions for installing Broker for GitLab.

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
