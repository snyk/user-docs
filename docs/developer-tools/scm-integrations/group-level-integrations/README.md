# Group-level integrations

Group-level SCM integrations provide broader visibility into all the application assets for a given customer and pull in the additional application context and, or metadata, for example, information on developers, commits, and so on.

At the Group level, you can set up and customize your Snyk Essentials integrations from the **Integrations** page, where the following SCMs are available:

* [GitHub](github-for-snyk-essentials.md) and GitHub Enterprise
* [GitLab](gitlab-for-snyk-essentials.md)
* [Azure DevOps](azure-devops-for-snyk-essentials.md)
* [BitBucket](bitbucket-for-snyk-essentials.md)

{% hint style="info" %}
If your SCM instance is not publicly accessible, you must connect using Snyk Broker. For details, see [Snyk Broker - Snyk Essentials](../../../implementation-and-setup/enterprise-setup/snyk-broker/using-snyk-essentials-with-snyk-broker.md).
{% endhint %}

The Integrations page at the Group-level shows all active integrations, including any data automatically synced from your existing Snyk Organizations, and provides access to the Integration Hub.

{% hint style="warning" %}
The SCM integrations use an incremental approach to retrieve repositories. This means that when a sync is initiated, it checks the last update time of the repository and only transfers the repositories that have been modified since then.

If there have been changes to the scope of the user or PAT used for the integration, any repositories that are newly within scope will only be identified after either a change to trigger the incremental collection or a change to the integration configuration itself.
{% endhint %}

The following supported Snyk data are automatically synced:

* Snyk Open Source
* Snyk Code
* Snyk IaC
* Snyk Container

Each connected integration enables you to:

* Pause data syncing
* Modify integration profiles and configurations
* Delete the integration
* Check when the integration was last synced and when the next sync is scheduled.

See the [Integration syncing time](../../../integrations/integrate-with-snyk.md#integrations-syncing-time) for more details about the time required to sync for each action.

### Prerequisites

To configure a Group-level integration, you must be a Group Admin or have a custom role that includes the `Edit Snyk Essentials` permissions under the [Group-level permissions](../../../snyk-platform-administration/user-roles/pre-defined-roles.md#group-level-permissions).

### Wildcard SCM integration

The wildcard integration allows you to use a special character to detect and integrate multiple SCM organizations simultaneously.

{% hint style="info" %}
The wildcard integration applies to the GitHub integration and offers support when you set it up using [Snyk Broker](../../../implementation-and-setup/enterprise-setup/snyk-broker/using-snyk-essentials-with-snyk-broker.md).
{% endhint %}

You can use the wildcards while setting up your integration using the **Integrations** page:

* Navigate to **Integrations**.
* Select the **SCM** tag and search for GitHub or Azure DevOps.
* Click the **Add** button.
* In the **Organizations** field, add the Organization details by using the `*` symbol. For example, using `*snyk` integrates all SCM Organizations that have Snyk in their name.
* All the Organizations that match with the wildcard, `*` symbol will be added.

{% hint style="info" %}
The wildcard, `*` symbol is considered a living command and will be applied every time you are rescanning your repositories.
{% endhint %}

### Snyk Essentials integrations ecosystem

You can refer to the list below to view available and compatible integrations for Snyk Essentials.&#x20;

<table><thead><tr><th width="256.0667724609375">Integration type</th><th width="385.7491455078125">Integration name</th></tr></thead><tbody><tr><td>SCM</td><td><ul><li><a href="github-for-snyk-essentials.md">GitHub</a></li><li><a href="bitbucket-for-snyk-essentials.md">BitBucket</a></li><li><a href="gitlab-for-snyk-essentials.md">GitLab</a></li><li><a href="azure-devops-for-snyk-essentials.md">Azure DevOps</a></li></ul></td></tr><tr><td>Dev portals and Service catalogs</td><td><ul><li><a href="../application-context-for-scm-integrations/">Backstage catalog</a></li><li><a href="../application-context-for-scm-integrations/#servicenow-cmdb-for-scm-integrations">ServiceNow CMDB</a></li><li><a href="../application-context-for-scm-integrations/#atlassian-compass">Atlassian Compass</a></li><li><a href="../application-context-for-scm-integrations/#harness">Harness</a></li><li><a href="../application-context-for-scm-integrations/#opslevel">OpsLevel</a></li><li><a href="../application-context-for-scm-integrations/#datadog-org-context-service-catalog">Datadog Org Context (Service Catalog)</a></li></ul></td></tr><tr><td>Risk management collaboration</td><td><ul><li><a href="../../../integrations/jira-and-slack-integrations/jira-integration.md">Jira</a></li><li><a href="../../../integrations/jira-and-slack-integrations/slack-integration.md">Slack</a></li><li>Email</li></ul></td></tr></tbody></table>

### Using the Integrations page

Use the **Integrations** page to onboard integrations and populate Snyk Essentials with data from SCM tools.

After the integration is validated, a card is displayed on the Integrations page, allowing you to enable or disable the connection, edit the settings, or remove the connection from your configuration.

{% hint style="info" %}
If you modify the permissions and scopes after the initial configuration, it is essential to either initiate an import or implement a change within the repository. This action allows Snyk to acknowledge and incorporate the updates effectively.
{% endhint %}
