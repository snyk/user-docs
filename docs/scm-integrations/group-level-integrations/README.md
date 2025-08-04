# Group-level integrations

Group-level SCM integrations provide broader visibility into all the application assets for a given customer and pull in the additional application context and, or metadata, for example, information on developers, commits, and so on.

At the Group level, you can set up and customize your Snyk Essentials integrations from the **Integrations Hub,** where the following SCMs are available:

* [GitHub](github-for-snyk-essentials.md) and GitHub Enterprise
* [GitLab](gitlab-for-snyk-essentials.md)
* [Azure DevOps](azure-devops-for-snyk-essentials.md)
* [BitBucket](bitbucket-for-snyk-essentials.md)

{% hint style="info" %}
If your SCM instance is not publicly accessible, you must connect using Snyk Broker. For details, see [Snyk Broker - Snyk Essentials](../../enterprise-setup/snyk-broker/using-snyk-essentials-with-snyk-broker.md).
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
* Snyk Container&#x20;

Each connected integration enables you to:

* Pause data syncing
* Modify integration profiles and configurations
* Delete the integration
* Check when the integration was last synced and when the next sync is scheduled.

See the [Integration syncing time](../../integrate-with-snyk/#integrations-syncing-time) for more details about the time required to sync for each action.

### Wildcard SCM integration

The wildcard integration allows you to use a special character to detect and integrate multiple SCM organizations simultaneously.&#x20;

{% hint style="info" %}
The wildcard integration applies to the GitHub integration and offers support when you set it up using [Snyk Broker](../../enterprise-setup/snyk-broker/using-snyk-essentials-with-snyk-broker.md).&#x20;
{% endhint %}

You can use the wildcards while setting up your integration using the Integration Hub:

* Open the **Integration Hub** menu.&#x20;
* Select the **SCM** tag and search for GitHub or Azure DevOps.&#x20;
* Click the **Add** button.
* In the **Organizations** field, add the Organization details by using the `*` symbol. For example, using  `*snyk` integrates all SCM Organizations that have Snyk in their name.
* All the Organizations that match with the wildcard, `*` symbol will be added.&#x20;

{% hint style="info" %}
The wildcard, `*` symbol is considered a living command and will be applied every time you are rescanning your repositories.&#x20;
{% endhint %}

### Snyk Essentials and Snyk AppRisk integrations ecosystem

You can refer to the table below to verify the availability and compatibility of all integrations for Snyk Essentials and Snyk AppRisk. The integrations are categorized by type, listed by name, and indicated as available or not for both Snyk Essentials and Snyk AppRisk.

<table><thead><tr><th width="172">Integration type</th><th width="164">Integration name</th><th width="198">Snyk Essentials</th><th>Snyk AppRisk</th></tr></thead><tbody><tr><td>SCM</td><td><ul><li><a href="../organization-level-integrations/github.md#group-level-snyk-apprisk-integrations">GitHub</a></li><li><a href="../organization-level-integrations/bitbucket-cloud.md#bitbucket-setup-guide">BitBucket</a></li><li><a href="../organization-level-integrations/github.md#group-level-snyk-apprisk-integrations">GitLab</a></li><li><a href="../organization-level-integrations/azure-repositories-tfs.md#azure-devops-setup-guide">Azure DevOps</a></li></ul></td><td>                <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td>                   <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>Dev portals and Service catalogs</td><td><ul><li><a href="../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/">Backstage catalog</a></li><li><a href="../../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#servicenow-cmdb-setup-guide">ServiceNow CMDB</a></li><li><a href="../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#atlassian-compass">Atlassian Compass</a></li><li><a href="../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#harness">Harness</a></li><li><a href="../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#opslevel">OpsLevel</a></li><li><a href="../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#datadog-org-context-service-catalog">Datadog Org Context (Service Catalog)</a></li></ul></td><td>               <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td>                     <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>Risk management collaboration</td><td><ul><li><a href="../../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#jira-setup-guide">Jira</a></li><li><a href="../../integrate-with-snyk/jira-and-slack-integrations/slack-integration.md">Slack</a></li><li>Email</li></ul></td><td>                <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td>                    <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>AST</td><td><ul><li><a href="../../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#nightfall-setup-guide">NightFall</a></li><li><a href="../../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#gitguardian-setup-guide">GitGuardian</a></li></ul></td><td>               <span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td>                     <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>Runtime security and observability</td><td><ul><li><a href="../../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md">Snyk runtime sensor</a></li><li><a href="../../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#sysdig-setup-guide">Sysdig</a></li><li><a href="../../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#dynatrace-setup-guide">Dynatrace</a></li></ul></td><td>               <span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td>                     <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr></tbody></table>

### Using the Integration Hub

Use the Integration Hub page to onboard integrations and populate Snyk Essentials with data from SCM tools.

See the [Snyk Web UI](../../getting-started/snyk-web-ui.md#manage-your-integrations) page for step-by-step instructions on how to set up an integration.

After the integration is validated, a card is displayed on the Integrations page, allowing you to enable or disable the connection, edit the settings, or remove the connection from your configuration.

{% hint style="info" %}
If you modify the permissions and scopes after the initial configuration, it is essential to either initiate an import or implement a change within the repository. This action allows Snyk to acknowledge and incorporate the updates effectively.
{% endhint %}
