# Validate the template and examples

## Example of a Custom PR template

The following template example shows how to use the variables inside a PR template.

<pre class="language-yaml"><code class="lang-yaml">title: This PR fixes {{ issue_count }} issues
branchName: fix/{{ issue_count }}-issues
commitMessage: "fix: {{ issue_count }} Snyk issues"
description: |
  {{ #is_upgrade_pr }}
  This PR has been opened to make sure our repositories are kept up-to-date.
  It updates {{ package_name }} from version {{ package_from }} to version {{ package_to }}.
  Review relevant docs for possible breaking changes.
<strong>  {{ /is_upgrade_pr }}
</strong>  
  **Tickets**
  {{ #jira_ids }}
  - Fixes {{ . }}
  {{ /jira_ids }}
  
  To find more details, see the Snyk project [{{ snyk_project_name }}]({{ snyk_project_url }})
</code></pre>

## Validate the Custom PR template

You can validate the correctness of your template by following these steps:

1. For the API configuration, an error response is returned if there are any issues with the customized content. Fix the issues and wait to receive an API success response.&#x20;
2. For the manual customization, add the template file to your Project or repository and follow the next step.&#x20;
3. For both manual and API customization, open a PR and verify that your customized inputs are being used.&#x20;

{% hint style="info" %}
The Customize PRs feature is in closed beta. Any feedback on improving the functionality is welcome.
{% endhint %}
