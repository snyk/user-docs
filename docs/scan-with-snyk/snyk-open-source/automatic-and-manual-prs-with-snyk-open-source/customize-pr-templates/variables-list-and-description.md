# Variables list and description

{% tabs %}
{% tab title="API custom PR templates" %}


You can use the following variables in your template.

## <mark style="color:purple;">`jira_ids: string[]`</mark>

A list of Jira tickets associated with the issues contained within the pull request. Ensure[ ](#user-content-fn-1)[^1]that the Snyk Jira integration is enabled on the Project or repository that contains the Project and that you have linked Snyk issues to JIRA tickets.

To automatically link Jira to the relevant pull requests, include a list of associated Jira tickets in the commit message.&#x20;

### Input

```json
{
    "data": {
        "attributes": {
            "commit_message": "This pull request is from Snyk and relates to {{ jira_ids }}" 
        },
        "type": "pull_request_template"
    }
}
```

### Output

The commit message of your PR will be:&#x20;

```json
This pull request is from Snyk and relates to JIRA-1,JIRA-2,JIRA-3
```

This output indicates that the suggested solution successfully resolved three problems. It also includes links to every Jira ticket.

## <mark style="color:purple;">`snyk_project_url: string`</mark>

This is the Snyk Project URL and can be used to link to the Snyk Project page.&#x20;

### Input

```json
{
    "data": {
        "attributes": {
            "description": "To find more details, see the Snyk project {{ snyk_project_url }}"
            
        },
        "type": "pull_request_template"
    }
}
```

### Output

The description of your PR will be:&#x20;

```yaml
To find more details, see the Snyk project https://app.snyk.io/org/my-org/project/xx-xxx-xx-xx
```

In this output, `my-org` is your Snyk Organization name and `xx-xxx-xx-xx-xxxx` is the public ID of your Project or repository.&#x20;

##

## <mark style="color:purple;">`snyk_project_name: string`</mark>

This is the Snyk Project name. You can add the Snyk Project name to your description.

### Input

```json
{
    "data": {
        "attributes": {
            "description": "Fix applied to project {{ snyk_project_name }}"
            
        },
        "type": "pull_request_template"
    }
}
```

### Output

The description of your PR will be:&#x20;

```yaml
Fix applied to project my-org/project:filename
```

## <mark style="color:purple;">`snyk_org_name: string`</mark>

This is the Snyk Organization name. You can add the Snyk Organization name to your description.

### Input

```json
{
    "data": {
        "attributes": {
            "description": "Fix applied by {{ snyk_org_name }}"
            
        },
        "type": "pull_request_template"
    }
}
```

### Output

The description of your PR will be:&#x20;

```yaml
Fix applied by my-org
```

## <mark style="color:purple;">`package_name: string`</mark>

This is the name of the package being fixed or upgraded. When more than one package is changed, this variable will default to the first one.

Follow this example to display in the description the package name of the first dependency being fixed in the PR.&#x20;

### Input

```json
{
    "data": {
        "attributes": {
            "description": "Fixes {{ package_name }}"
            
        },
        "type": "pull_request_template"
    }
}
```

### Output

The description of your PR will be:&#x20;

```yaml
Fixes adm-zip
```

## <mark style="color:purple;">`package_from: string`</mark>

This is the version of the package that is being fixed or upgraded. In cases where more than one package is changed, this variable will default to the `from` version of the first one.

### Input

```json
{
    "data": {
        "attributes": {
            "description": "Fix is applied by moving from {{ package_from}}"
            
        },
        "type": "pull_request_template"
    }
}
```

### Output

The description of your PR will be:&#x20;

```yaml
Fix is applied by moving from 0.4.7
```

## <mark style="color:purple;">`package_to: string`</mark>

The package is transitioning to this particular version. In cases where more than one package is changed, this variable will default to the `to` version of the first one.

### Input

```json
{
    "data": {
        "attributes": {
            "description": "Fix is applied by moving to {{ package_to}}"
            
        },
        "type": "pull_request_template"
    }
}
```

### Output

The description of your PR will be:&#x20;

```yaml
Fix is applied by moving to 0.5.2
```

## <mark style="color:purple;">`issue_count: number`</mark>

This is the number of issues in your Project or repository that are covered by the PR.&#x20;

### Input

```json
{
    "data": {
        "attributes": {
            "branch_name": "fix/{{ issue_count }}-issues"
            
        },
        "type": "pull_request_template"
    }
}
```

### Output

The branch name of your PR will be:&#x20;

```yaml
fix/98-issues-4e25b18624124a1b6f4dd00e3caa4f6c
```

## <mark style="color:purple;">`is_fix_pr: boolean`</mark>

This variable can be used to customize attributes based on whether the PR is a backlog PR,  for example, opened to fix new vulnerabilities introduced to the Project or repository in the latest scan. In the example below you can see that the description of the PR will only show if it is a fix PR.

### Input

```json
{
    "data": {
        "attributes": {
            "description": "{{ #is_fix_pr }} This PR has been opened to fix vulnerabilities in your project. {{ /is_fix_pr }}"
            
        },
        "type": "pull_request_template"
    }
}
```

### Output

If your PR is a fix PR then the description of your PR will be:&#x20;

```yaml
This PR has been opened to fix vulnerabilities in your project.
```

## <mark style="color:purple;">`is_backlog_pr: boolean`</mark>

This variable can be used to customize attributes based on whether the PR is a backlog PR, for example, opened to fix known vulnerabilities already in the Project or repository. In the example below you can see that the description of the PR will only show if it is a backlog PR.&#x20;

### Input

```json
{
    "data": {
        "attributes": {
            "description": "{{ #is_backlog_pr }} This PR has been opened to fix known vulnerbilities. These vulnerabilities are retrieved from the Project's backlog. {{ /is_backlog_pr }}"
            
        },
        "type": "pull_request_template"
    }
}
```

### Output

If your PR is a backlog PR then the description of your PR will be:&#x20;

```yaml
This PR has been opened to fix known vulnerabilities. These vulnerabilities are retrieved from the Project's backlog.
```

## <mark style="color:purple;">`is_upgrade_pr: boolean`</mark>

This variable can be used to customize attributes based on whether the PR is an Upgrade PR, or to upgrade dependencies to newer versions regardless of vulnerabilities. In the example below you can see that the description of the PR will only show if it is an upgrade PR.&#x20;

### Input

```json
{
    "data": {
        "attributes": {
            "description": "{{ #is_upgrade_pr }} This PR has been opened to make sure our repositories are kept up-to-date. It updates {{ package_name }} from version {{ package_from }} to version {{ package_to }}. Review relevant docs for possible breaking changes. {{ /is_upgrade_pr }} "
            
        },
        "type": "pull_request_template"
    }
}
```

### Output

If your PR is an upgrade PR then the description of your PR will be:&#x20;

```json
This PR has been opened to make sure our repositories are kept up-to-date. It updates package-x from version 1.0.0 to version 2.0.0. Review relevant docs for possible breaking changes.
```

## <mark style="color:purple;">`snyk_pull_request_type: prType (fix, upgrade, backlog, unknown)`</mark>

This is the prType of your Project or repository. You can use it to display the PR type from the pull request description.

### Input

```json
{
    "data": {
        "attributes": {
            "commit_message": "{{ snyk_pull_request_type}}: for {{ package_name }}"
            
        },
        "type": "pull_request_template"
    }
}
```

### Output

If you have opened a Fix PR then the commit message of your PR will be:&#x20;

```yaml
fix: for package-x
```
{% endtab %}

{% tab title="YAML file custom PR templates" %}


You can use the following variables in your template. These variables can be used in any of the customizable PR properties.&#x20;

## <mark style="color:purple;">`jira_ids: string[]`</mark>

A list of Jira tickets associated with the issues contained within the pull request. Ensure[ ](#user-content-fn-2)[^2]that the Snyk Jira integration is enabled on the Project or repository that contains the Project and that you have linked Snyk issues to JIRA tickets.

To automatically link Jira to the relevant pull requests, include a list of associated Jira tickets in the commit message.&#x20;

### Input

```yaml
commitMessage: |
  This pull request is from Snyk and relates to {{ jira_ids }}
```

### Output

The commit message of your PR will be:&#x20;

```yaml
This pull request is from Snyk and relates to JIRA-1,JIRA-2,JIRA-3
```

This output indicates that the suggested solution successfully resolved three problems. It also includes links to every Jira ticket.

## <mark style="color:purple;">`snyk_project_url: string`</mark>

This is the Snyk Project URL and can be used to link to the Snyk Project page.&#x20;

### Input

```yaml
description: |
  To find more details, see the Snyk project {{ snyk_project_url }}
```

### Output

The description of your PR will be:&#x20;

```yaml
To find more details, see the Snyk project https://app.snyk.io/org/my-org/project/xx-xxx-xx-xx
```

In this output, `my-org` is your Snyk Organization name and `xx-xxx-xx-xx-xxxx` is the public ID of your Project or repository.&#x20;

## <mark style="color:purple;">`snyk_project_name: string`</mark>

This is the Snyk Project name. You can add the Snyk Project name to your description.

### Input

```yaml
description: |
  Fix applied to project {{ snyk_project_name }}
```

### Output

The description of your PR will be:&#x20;

```yaml
Fix applied to project my-org/project:filename
```

## <mark style="color:purple;">`snyk_org_name: string`</mark>

This is the Snyk Organization name. You can add the Snyk Organization name to your description.

### Input

```yaml
description: |
  Fix applied by {{ snyk_org_name }}
```

### Output

The description of your PR will be:&#x20;

```yaml
Fix applied by my-org
```

## <mark style="color:purple;">`package_name: string`</mark>

This is the name of the package being fixed or upgraded. When more than one package is changed, this variable will default to the first one.

Follow this example to display in the description the package name of the first dependency being fixed in the PR.&#x20;

### Input

```yaml
description: |
  Fixes {{ package_name }}
```

### Output

The description of your PR will be:&#x20;

```yaml
Fixes adm-zip
```

## <mark style="color:purple;">`package_from: string`</mark>

This is the version of the package that is being fixed or upgraded. In cases where more than one package is changed, this variable will default to the `from` version of the first one.

### Input

```yaml
description: |
  Fix is applied by moving from {{ package_from}}

```

### Output

The description of your PR will be:&#x20;

```yaml
Fix is applied by moving from 0.4.7
```

## <mark style="color:purple;">`package_to: string`</mark>

The package is transitioning to this particular version. In cases where more than one package is changed, this variable will default to the `to` version of the first one.

### Input

```yaml
description: |
  Fix is applied by moving to {{ package_to}}

```

### Output

The description of your PR will be:&#x20;

```yaml
Fix is applied by moving to 0.5.2
```

## <mark style="color:purple;">`issue_count: number`</mark>

This is the number of issues in your Project or repository that are covered by the PR.&#x20;

### Input

```yaml
branchName: fix/{{ issue_count }}-issues
```

### Output

The branch name of your PR will be:&#x20;

```yaml
fix/98-issues-4e25b18624124a1b6f4dd00e3caa4f6c

```

## <mark style="color:purple;">`is_fix_pr: boolean`</mark>

This checks to determine whether the pull request is a fix PR, for example, opened to fix new vulnerabilities introduced to the Project or repository in the latest scan.

### Input

```yaml
description: |
  Is this pr a fix pr? {{ is_fix_pr }}

```

### Output

The description of your PR will be:&#x20;

```yaml
Is this a fix pr? true
```

## <mark style="color:purple;">`is_backlog_pr: boolean`</mark>

This checks to determine whether the pull request is a backlog PR, for example, opened to fix known vulnerabilities already in the Project or repository.

### Input

```yaml
description: |
  Is this pr a backlog pr? {{ is_backlog_pr }}
```

### Output

The description of your PR will be:&#x20;

```yaml
Is this a backlog pr? false
```

## <mark style="color:purple;">`is_upgrade_pr: boolean`</mark>

This checks whether the pull request is an upgrade PR, for example, opened to upgrade dependencies to newer versions regardless of vulnerabilities.

### Input

```yaml
description: |
  Is this pr an upgrade pr? {{ is_upgrade_pr }}
```

### Output

The description of your PR will be:&#x20;

```yaml
Is this an upgrade pr? false
```

## <mark style="color:purple;">`snyk_pull_request_type: prType (fix, upgrade, backlog, unknown)`</mark>

This is the prType of your Project or repository. You can use it to display the PR type from the pull request description.

### Input

```yaml
description: |
  This is a {{ snyk_pull_request_type }} pull request
```

### Output

The description of your PR will be:&#x20;

```yaml
This is a fix pull request
```
{% endtab %}
{% endtabs %}

[^1]: Ensure



[^2]: Ensure

