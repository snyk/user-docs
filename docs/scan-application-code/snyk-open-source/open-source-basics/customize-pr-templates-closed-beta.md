# Customize PR Templates (Closed Beta)



{% hint style="warning" %}
This feature is currently in closed beta. The functionality is very likely to evolve as we take on feedback and there will be breaking changes. Please speak to your account manager about gaining access to this feature.
{% endhint %}

## Understanding Customized PRs

When you fix or upgrade Snyk Open Source and Container Projects imported using the [SCM integrations](../../../integrations/git-repository-scm-integrations/), Snyk raises pull requests (PRs) against your repository.&#x20;

Snyk has default templates for the title, description, commit message, and branch name. These indicate what packages are being changed, which issues are being fixed, and many other details.

You may have your own standards and practices for submitting pull requests. For instance, if a pull request comes from Snyk, you may require that the title begins with `SNYK:`. This page outlines the areas of pull requests that you can customize and provides instructions on how to do so.

## Using the Customize PRs feature

When enabling the Customize Snyk PR feature, you will need to add a custom template file to your repository. This allows you to configure the title, description, commit message, and branch name. Going forward, this will apply to all Projects associated with that repository.

{% hint style="warning" %}
We are initially looking for feedback on the variables and templating system, once we have validated our approach we will look into building more robust authoring workflows via our API and UI interfaces.
{% endhint %}

## Creating the PR template

* Create the YAML template in an [appropriate location ](customize-pr-templates-closed-beta.md#file-locations)and use the [handlebars](https://handlebarsjs.com) syntax for templating.
* The YAML file can define several string properties, including:
  * `title`
  * `branchName`
  * `commitMessage`
  * `description`

{% hint style="info" %}
When the Customize Snyk PRs feature is enabled, all PRs from Snyk adopt this format, effectively switching off Snyk’s default template. Only strings are acceptable values. Lists or numbers are not allowed.
{% endhint %}

* You can use YAML multiline operators. You can create a detailed description that spans several lines by following this format:&#x20;

```yaml
description: |
  This pull request comes from Snyk
  For more information see the project page {{ snyk_project_url }}
  If you have more questions reach out to a member of the security team

```

:link:The pipe operator preserves new line characters. Use greater than `>` to join all the lines by a space with a new line on the end. To use a colon, for example, you can either use multiline operators, `|` or `>`, or enclose the line in double quotes:

```yaml
commitMessage: "snyk: this is a security pull request"
```

### Customizable properties

* `title` - customize the PR title
* `commitMessage` - customize the PR commit message
* `branch`- customize the PR branch name
* `description` - customize the PR description&#x20;

{% hint style="info" %}
Snyk will use the default content if any of the customizable properties are missing.
{% endhint %}

### File locations

To enable Snyk to access your customized template, you need to add a file called `snyk_pull_request_template.yaml` to your Project. The storage method varies based on the type of integration.

* Github - `./github/snyk_pull_request_template.yaml`
* GitLab  - `./gitlab/snyk_pull_request_template.yaml`
* Azure  - `./.azuredevops/snyk_pull_request_template.yaml`
* Other (like BitBucket)  - `./.config/snyk_pull_request_template.yaml`

## Variables list and description

You can use the following variables in your template.

### <mark style="color:purple;">`jira_ids: string[]`</mark>

A list of Jira tickets associated with the issues contained within the pull request. Make sure that the Snyk Jira integration is enabled on your Project and that you have linked Snyk issues to JIRA tickets.

To automatically link Jira to the relevant pull requests, include a list of associated Jira tickets in the commit message.&#x20;

#### Input

```yaml
commitMessage: |
  This pull request is from Snyk and relates to {{ jira_ids }}
```

#### Output

```yaml
This pull request is from Snyk and relates to JIRA-1,JIRA-2,JIRA-3
```

This outcome indicates that the suggested solution successfully resolved three problems, and it also includes links to every Jira ticket.

### <mark style="color:purple;">`snyk_project_url: string`</mark>

This is the Snyk Project URL and can be used to link to the Snyk Project page.&#x20;

Use the following example to link to the Snyk Project page from the pull request description.

#### Input

```yaml
description: |
  To find more details, see the Snyk project {{ snyk_project_url }}
```

#### Output

```yaml
To find more details, see the Snyk project https://app.snyk.io/org/my-org/project/xx-xxx-xx-xx
```

In this outcome, `my-org` is your Snyk Organisation name and `xx-xxx-xx-xx-xxxx` is the public ID of your Project.&#x20;

### <mark style="color:purple;">`snyk_project_name: string`</mark>

This is the Snyk Project name. You can add the Snyk Project name to your description.

#### Input

```yaml
description: |
  Fix applied to project {{ snyk_project_name }}
```

#### Output

```yaml
Fix applied to project my-org/project:filename
```

### <mark style="color:purple;">`snyk_org_name: string`</mark>

This is the Snyk Organization name. You can add the Snyk Organization name to your description.

#### Input

```yaml
description: |
  Fix applied by {{ snyk_org_name }}
```

#### Output

```yaml
Fix applied by my-org
```

### <mark style="color:purple;">`package_name: string`</mark>

This is the name of the package being fixed or upgraded. In cases where more than one package is changed this variable will default to the first one.

Follow the below example to display in the branch name the number of issues in your Project.&#x20;

#### Input

```yaml
description: |
  Fixes {{ package_name }}
```

#### Output

```yaml
Fixes adm-zip
```

### <mark style="color:purple;">`package_from: string`</mark>

This is the previous version of the package that is being fixed or upgraded. In cases where more than one package is changed this variable will default to the `from` version of the first one.

#### Input

```yaml
description: |
  Fix is applied by moving from {{ package_from}}

```

#### Output

```yaml
Fix is applied by moving from 0.4.7
```

### <mark style="color:purple;">`package_to: string`</mark>

The package is transitioning to this particular version. In cases where more than one package is changed this variable will default to the `from` version of the first one.

#### Input

```yaml
description: |
  Fix is applied by moving to {{ package_to}}

```

#### Output

```yaml
Fix is applied by moving to 0.5.2
```

### <mark style="color:purple;">`issue_count: number`</mark>

This is the number of issues in your Project that are covered by the PR.&#x20;

#### Input

```yaml
branchName: fix/{{ issue_count }}-issues
```

#### Output

```yaml
fix/98-issues-4e25b18624124a1b6f4dd00e3caa4f6c

```

### <mark style="color:purple;">`is_fix_pr: boolean`</mark>

This checks if the pull request is a fix PR, for example, opened to fix new vulnerabilities introduced to the Project in the latest scan.

#### Input

```yaml
description: |
  Is this pr a fix pr? {{ is_fix_pr }}

```

#### Output

```yaml
Is this a fix pr? true
```

### <mark style="color:purple;">`is_backlog_pr: boolean`</mark>

This checks if the pull request is a backlog PR, for example, opened to fix known vulnerabilities already in the Project.

#### Input

```yaml
description: |
  Is this pr a backlog pr? {{ is_backlog_pr }}
```

#### Output

```yaml
Is this a backlog pr? false
```

### <mark style="color:purple;">`is_upgrade_pr: boolean`</mark>

This checks if the pull request is an upgrade PR, for example, opened to upgrade dependencies to newer versions regardless of vulnerabilities.

#### Input

```yaml
description: |
  Is this pr an upgrade pr? {{ is_upgrade_pr }}
```

#### Output

```yaml
Is this an upgrade pr? false
```

### <mark style="color:purple;">`snyk_pull_request_type: prType (fix, upgrade, backlog, unknown)`</mark>

This is the prType of your Project. You can use it to display the PR type from the pull request description.

#### Input

```yaml
description: |
  This is a {{ snyk_pull_request_type }} pull request
```

#### Output

```yaml
This is a fix pull request
```

## Troubleshooting

If you encounter an error, double-check your template and apply the below instructions:

### Check template syntax

Ensure that the template is written in valid handlebars syntax. Any syntax errors can lead to unexpected behavior. Double-check the syntax and make any necessary changes. &#x20;

### Verify variable names

Review the variables used in the template and ensure they match the supported variables provided in the [Variables list](customize-pr-templates-closed-beta.md#variables-list-and-description). If a variable is misspelled or not recognized, it may result in unexpected output or fallback to default values.&#x20;

{% hint style="info" %}
As the feature progresses, we will provide more comprehensive error reporting and validation of the template.
{% endhint %}

## Limitations

#### Branch names&#x20;

Although you can still customize your branch name, to avoid any potential issues, we automatically append a hash to the end of the branch name you choose. This hash is derived from the PR contents and prevents opening duplicate branches, as well as naming conflicts with existing or future branches.

#### Jira

You can add Jira tickets to your PR, but we are limited to those created via Snyk workflows (either UI or API).

#### Fail to open PR

Currently, we do not surface errors when we fail to open a PR, this will be resolved in an upcoming update.

#### No line between variables

It is possible that a new line is not added between variables when two lines start with a variable `{{}}`.

## Example

The following template example shows how to use the variables inside a PR template.

```yaml
title: This PR fixes {{ issue_count }} issues
branchName: fix/{{ issue_count }}-issues
commitMessage: "fix: {{ issue_count }} Snyk issues"
description: |
  Fixes {{ issue_count }} issues
  
  **Tickets**
  {{ #each jira_ids }}
  - Fixes {{ self }}
  {{ /each }}
  
  To find more details, see the Snyk project [{{ snyk_project_name }}]({{ snyk_project_url }})
```

### Validating the template

You can validate the correctness of your template by:

1. Adding the template file to your Project.&#x20;
2. Opening a PR and verifying that your customized inputs are being used.&#x20;

{% hint style="info" %}
The Customize PRs feature is in closed beta. We welcome any feedback you have to improve its functionality.
{% endhint %}

\
