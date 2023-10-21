# Customize PR templates (closed beta)

{% hint style="info" %}
**Feature availability**\
This feature is currently in closed beta. The functionality is likely to evolve based on feedback, and there will be breaking changes. Contact your account manager to get access to this feature.&#x20;
{% endhint %}

{% hint style="warning" %}
After your account manager turns on the feature flag, you can activate this feature in the [Snyk Preview](../../../../snyk-admin/manage-settings/snyk-preview.md) settings.&#x20;
{% endhint %}

{% hint style="info" %}
**API support**\
Customize your PR template automatically by making an API request; see [Create an API request to customize the PR template](./#create-an-api-request-to-customize-the-pr-template). For more information, see the [Snyk API pull request tempates endpoints](https://apidocs.snyk.io/?version=2023-10-13%7Ebeta#tag--Pull-Request-Templates).
{% endhint %}

## Understand Customized PRs

When you fix or upgrade Snyk Open Source and Container Projects imported using the [SCM integrations](../../../../integrations/git-repository-scm-integrations/), Snyk raises pull requests (PRs) against your repository.&#x20;

Snyk has default templates for the title, description, commit message, and branch name. These indicate what packages are being changed, which issues are being fixed, and many other details.

You may have your own standards and practices for submitting pull requests. For instance, if a pull request comes from Snyk, you may require the title to begin with `SNYK:`. This page identifies the areas of pull requests that you can customize and provides instructions on how to do so.

On the [supported languages, frameworks, and feature availability overview](../../../supported-languages-and-frameworks/) page, you can find all the languages for which the Fix PR functionality is supported.

{% hint style="warning" %}
Snyk is looking initially for feedback on the variables and templating system. After the approach is validated, Snyk will look into building more robust authoring workflows using the API and UI interfaces.
{% endhint %}

## Enable the Customize PR feature

Follow these steps to enable the Configure Snyk Pull Requests feature:

1. Log in to your Snyk account.
2. Select **Settings**, then **Snyk Preview**.
3. Enable the **Configure Snyk Pull Requests** feature.\
   You can enable this feature at the Group level. See [Configure Automatic fix PRs](../../../../scan-application-code/snyk-open-source/open-source-basics/fix-pull-requests-for-known-vulnerabilities-backlog.md) for more configuration details.

<figure><img src="../../../../.gitbook/assets/Enable config pull request.png" alt="Enable the Configure Snyk Pull Requests feature"><figcaption><p>Enable the Configure Snyk Pull Requests feature</p></figcaption></figure>

## Create the PR template

Snyk provides the following options for creating a PR template:

1. [Manually create the YAML template](./#manually-customize-the-pr-template) by using the [mustache](https://mustache.github.io) syntax for templating and add the file to your Project or repository.
2. Customize the content of the PR template (`title`, `branchName`, `commitMessage`, `description`) and [send the content as a payload to the Snyk API](./#create-an-api-request-to-customize-the-pr-template), instead of creating it as a YAML file.&#x20;

The YAML file can define several string properties, including:

* `title`
* `branchName`
* `commitMessage`
* `description`

{% hint style="info" %}
When the Customize Snyk PRs feature is enabled, all PRs from Snyk adopt this format, effectively switching off the default Snyk template. The only acceptable values are strings. Lists and numbers are not allowed.
{% endhint %}

### YAML multiline operators

You can use YAML multiline operators. You can create a detailed description that spans several lines by following this format:&#x20;

```yaml
description: |
  This pull request comes from Snyk
  For more information see the project page {{ snyk_project_url }}
  If you have more questions reach out to a member of the security team

```

:link:The pipe operator preserves new line characters. Use greater than, `>` , to join all the lines by a space with a new line on the end. To use a colon, you can either use multiline operators, `|` or `>`, or enclose the line in double quotes:

```yaml
commitMessage: "snyk: this is a security pull request"
```

### Customizable properties

The following properties are customizable:

* `title` - customize the PR title
* `commitMessage` - customize the PR commit message
* `branch`- customize the PR branch name
* `description` - customize the PR description&#x20;

{% hint style="info" %}
If any customizable properties are missing from your template, Snyk reverts to the default values for these properties when opening a pull request.&#x20;
{% endhint %}

## Use the Custom PR template

{% hint style="info" %}
Ensure that your account manager has turned on the feature fla and you enabled the feature in the Snyk Preview settings.&#x20;
{% endhint %}

You can choose to customize the PR template using an API request or by manually uploading the customized PR template as a YAML file to your Snyk Project.

* [**Customize the PR template using the API**](./#create-an-api-request-to-customize-the-pr-template): the template you configured using an API request is propagated to all Projects included at the Group level.
* &#x20;[**Customize the PR template manually**](./#manually-customize-the-pr-template): the template you configured is applied only to the Project where you added the customized template file.

{% hint style="info" %}
The API configuration is momentarily available only at Group level.
{% endhint %}

### **Create an API request to customize the PR template**&#x20;

You can customize your the `title`, `commit message,` `branch name`, and `description`  for your Group by using the [Snyk API pull request template](https://apidocs.snyk.io/?version=2023-10-13%7Ebeta#get-/groups/-group\_id-/settings/pull\_request\_template). After the template is customized on the Group level, all Projects in that Group are automatically updated with the latest changes.

### **Manually customize the PR template**

You can manually upload the YAML file with the name  `snyk_pull_request_template.yaml` to your Project(repository)[^1]. The method varies based on the type of integration.

* Github/ GitHub Enterprise - `/.github/snyk_pull_request_template.yaml`
* GitLab  - `/.gitlab/snyk_pull_request_template.yaml`
* Azure  - `/.azuredevops/snyk_pull_request_template.yaml`
* Other (like BitBucket)  - `/.config/snyk_pull_request_template.yaml`

{% hint style="info" %}
If you want to use multiple repositories for a custom template, add the customized YAML template file to each of these repositories.
{% endhint %}

If you use  [Snyk Broker](../../../../enterprise-setup/snyk-broker/), you must allow access to these file locations in the `accept.json` [configuration of your Broker client](../../../../enterprise-setup/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.2b-running-the-broker-client-with-the-code-snippets-display.md). The following describes the additional rules that should be added for each git integration.

#### Custom PR template for GitHub and GitHub Enterprise

Under the list of `private` rules add the following input. For more information, see [Configure Broker to be used for GitHub Enterprise](../../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-enterprise-install-and-configure-broker/setup-broker-with-github-enterprise.md#configure-broker-to-be-used-for-github-enterprise) in the instructions for installing Broker for GitHub Enterprise.

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

#### Custom PR template for Azure repos

Under the list of `private` rules, add the following two elements to the existing `valid.values` array for file content. For more information, see [Configure Broker to be used with Azure Repos](../../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/azure-repos-install-and-configure-broker/setup-broker-with-azure-repos.md#configure-broker-to-be-used-with-azure-repos) in the instructions for installing Broker for Azure repos.

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

#### Custom PR template for BitBucket Server

Under the list of `private` rules, add the following input. For more information, see [Bitbucket Server/Data Center - environment variables for Snyk Broker](../../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker.md) in the instructions for installing Broker for BitBucket Server.

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

#### Custom PR template for GitLab

Under the list of `private` rules, add the following input. For more information, see [GitLab - environment variables for Snyk Broker](../../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/gitlab-install-and-configure-broker/gitlab-environment-variables-for-snyk-broker.md) in the instructions for installing Broker for GitLab.

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
