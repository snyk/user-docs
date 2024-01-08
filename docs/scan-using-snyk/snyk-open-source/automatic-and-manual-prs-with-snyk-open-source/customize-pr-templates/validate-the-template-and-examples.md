# Examples and template validation

## Example of a Custom PR template

The following template example shows how to use the variables with a PR template.

### API Custom PR template examples

```json
{
    "data": {
        "attributes": {
            "title": "[{{ snyk_pull_request_type }}] for {{ package_name }}",
            "commit_message": "{{ snyk_pull_request_type}}: for {{ package_name }}",
            "branch_name": "fix/{{ issue_count }}-issues",
            "description": "Moving package {{ package_name }} from {{ package_from }} to {{ package_to }}\nFixes {{ issue_count }} issues\nFor more details see {{ snyk_project_url }}\nProject {{ snyk_project_name }}\nOrg {{ snyk_org_name }}"
        },
        "type": "pull_request_template"
    }
}
```

### YAML Custom PR template examples

```yaml
title: This PR fixes {{ issue_count }} issues
branchName: fix/{{ issue_count }}-issues
commitMessage: "fix: {{ issue_count }} Snyk issues"
description: |
  {{ #is_upgrade_pr }}
  This PR has been opened to make sure our repositories are kept up-to-date.
  It updates {{ package_name }} from version {{ package_from }} to version {{ package_to }}.
  Review relevant docs for possible breaking changes.
  {{ /is_upgrade_pr }}
  
  **Tickets**
  {{ #jira_ids }}
  - Fixes {{ . }}
  {{ /jira_ids }}
  
  To find more details, see the Snyk project [{{ snyk_project_name }}]({{ snyk_project_url }})
```

## Validate the Custom PR template

You can validate the correctness of your template by following the steps described for each section.

### API template validation

For the API configuration, you can validate the correctness of your template by following these steps:

1. An error response is returned if there are any issues with the customized content. Fix the issues and wait to receive an API success response.&#x20;
2. Open a PR and verify that your customized inputs are being used.&#x20;

{% hint style="info" %}
The Customize PRs feature is in closed beta. Any feedback on improving the functionality is welcome.
{% endhint %}

### YAMl template validation

You can validate the correctness of your template by following these steps:

1. Adding the template file to your Project(repository).&#x20;
2. Opening a PR and verifying that your customized inputs are being used.&#x20;
