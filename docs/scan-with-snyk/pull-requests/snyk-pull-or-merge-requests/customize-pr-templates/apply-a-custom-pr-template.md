# Apply a custom PR template

## Create and manage a custom PR template using the API

You can create a custom PR template using the API endpoint [Create or update pull request template for Group](../../../../snyk-api/reference/pull-request-templates.md#post-groups-group_id-settings-pull_request_template). Send an API request that contains a JSON payload with the custom properties. This request configures a Group-level pull request template that will be used on any Organization or Project within that Group. The pull request template created using the Snyk API can be updated at any time, and all Projects in the Group are automatically updated with the latest changes.

API configuration of PR templates is available only at the Group level.

When a custom template is uploaded using an API request, all Snyk PRs in that Group adopt this format, effectively switching off the default Snyk template for the customizable properties. Strings are the only acceptable values; lists and numbers are not allowed.

If any customizable properties are missing from your custom template, Snyk reverts to the default values for these properties when opening a pull request.&#x20;

The following properties are customizable using the API:

* `title` - customize the PR title
* `commit_message` - customize the PR commit message
* `description` - customize the PR description&#x20;

You cannot customize the branch name for your PRs. The branch name of your PRs will use the Snyk default value.&#x20;

You can retrieve the custom PR template for your Group using the endpoint [Get pull request template for Group](https://apidocs.snyk.io/?#get-/groups/-group_id-/settings/pull_request_template). This is useful if you want to consider changing your template, and in troubleshooting.

To delete the template, use the endpoint [Delete pull request template for group](../../../../snyk-api/reference/pull-request-templates.md#delete-groups-group_id-settings-pull_request_template).

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

You can manually upload the YAML file with the name `snyk_pull_request_template.yaml` to your Project (repository). The method varies based on the type of integration.

* GitHub/ GitHub Enterprise - `/.github/snyk_pull_request_template.yaml`
* GitLab - `/.gitlab/snyk_pull_request_template.yaml`
* Azure DevOps - `/.azuredevops/snyk_pull_request_template.yaml`
* Other (such as BitBucket) - `/.config/snyk_pull_request_template.yaml`

If you want to use a custom template for multiple repositories, add the YAML custom template file to each of these repositories.

## Broker configuration for fetching custom PR templates

If you use [Snyk Broker](../../../../implementation-and-setup/enterprise-setup/snyk-broker/), you must use a Broker at version 4.188.0 or higher and enable the Broker to fetch the custom PR templates using the `ACCEPT_CUSTOM_PR_TEMPLATES` environment variable.

To do this, you must remove `ACCEPT=/path/to/custom.json` and add the following environment variable to your Broker container or deployment:

```
ACCEPT_CUSTOM_PR_TEMPLATES=true
```
