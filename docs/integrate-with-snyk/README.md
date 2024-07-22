# Integrate with Snyk

## Integrations for Snyk

Many integrations are available for using third-party functionality within Snyk and using Snyk with other tools. See [SCM, IDE, and CI/CD workflow and integrations](../scm-ide-and-ci-cd-workflow-and-integrations/) for information on integrations and other methods of accomplishing that workflow.

This page identifies additional Snyk integrations and where to find them.

Snyk provides plugins for repository gatekeepers and integrations to connect with package repositories:

* [Gatekeeper plugins](gatekeeper-plugins/)
* [Package repository integrations](../scan-using-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/package-repository-integrations/)

There are integrations that support Snyk Container and Snyk Iac:

* [Snyk Container integrations](../scan-using-snyk/snyk-container/container-registry-integrations/)
* [Cloud platforms integrations](../scan-using-snyk/snyk-iac/cloud-platforms-integrations/)

Integrations for [event forwarding](event-forwarding/) allow you to push Snyk platform events directly to certain products on other platforms, enabling you to set up custom alerting, build your own reporting, trigger automation, and more.

[Notification and ticketing systems integrations](jira-and-slack-integrations/) help you work with Snyk in Jira and Slack.

Information is also provided on how Snyk can work with [vulnerability management tools](vulnerability-management-tools/).

Snyk provides alternative reporting tools. For more information, see [Reporting and BI integrations](../manage-risk/reporting/reporting-and-bi-integrations-snowflake-data-share.md).

## Integrations for Snyk AppRisk

The Integrations page shows all active integrations, including any data automatically synced from your existing Snyk Organizations, and provides access to the Integration Hub.

The following supported Snyk data are automatically synced: Snyk Open Source, Snyk Code, Snyk IaC, Snyk Container.&#x20;

{% hint style="info" %}
Navigate to the [Third-party integrations for Snyk AppRisk](third-party-integrations-for-snyk-apprisk.md) to check all available integrations and their setup details.&#x20;
{% endhint %}

Each connected integration enables you to pause data syncing, modify integration profiles and configurations, delete the integration, or check when the integration was last synced and when the next sync is scheduled.

### Snyk AppRisk integrations ecosystem

You can refer to the table below to verify the availability and compatibility of all integrations for Snyk AppRisk. The integrations are categorized by type, listed by name, and indicated as available or not for both Snyk AppRisk Essentials and Snyk AppRisk Pro.

<table><thead><tr><th width="172">Integration type</th><th width="164">Integration name</th><th width="198">Snyk AppRisk Essentials</th><th>Snyk AppRisk Pro</th></tr></thead><tbody><tr><td>SCM</td><td><ul><li><a href="../scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations/github.md#group-level-snyk-apprisk-integrations">GitHub</a></li><li><a href="../scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations/bitbucket-cloud.md#group-level-snyk-apprisk-integrations">BitBucket</a></li><li><a href="../scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations/gitlab.md#group-level-snyk-apprisk-integrations">GitLab</a></li><li><a href="../scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations/snyk-azure-repositories-tfs.md#group-level-snyk-apprisk-integrations">Azure DevOps</a></li></ul></td><td>                <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td>                   <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>Dev portals and Service catalogs</td><td><ul><li><a href="../scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations/application-context-for-scm-integrations.md">Backstage catalog</a></li><li><a href="third-party-integrations-for-snyk-apprisk.md#servicenow-cmdb-setup-guide">ServiceNow CMDB</a></li><li><a href="../scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations/application-context-for-scm-integrations.md#atlassian-compass">Atlassian Compass</a></li><li><a href="../scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations/application-context-for-scm-integrations.md#harness">Harness</a></li><li><a href="../scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations/application-context-for-scm-integrations.md#opslevel">OpsLevel</a></li><li><a href="../scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations/application-context-for-scm-integrations.md#datadog-org-context-service-catalog">Datadog Org Context (Service Catalog)</a></li></ul></td><td>               <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td>                     <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>Risk management collaboration</td><td><ul><li><a href="third-party-integrations-for-snyk-apprisk.md#jira-setup-guide">Jira</a></li><li><a href="jira-and-slack-integrations/slack-integration.md">Slack</a></li><li>Email</li></ul></td><td>                <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td>                    <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>AST</td><td><ul><li><a href="third-party-integrations-for-snyk-apprisk.md#nightfall-setup-guide">NightFall</a></li><li><a href="third-party-integrations-for-snyk-apprisk.md#gitguardian-setup-guide">GitGuardian</a></li></ul></td><td>               <span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td>                     <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>Runtime security and observability</td><td><ul><li><a href="snyk-runtime-sensor.md">Snyk runtime sensor</a></li><li><a href="third-party-integrations-for-snyk-apprisk.md#sysdig-setup-guide">Sysdig</a></li><li><a href="third-party-integrations-for-snyk-apprisk.md#dynatrace-setup-guide">Dynatrace</a></li></ul></td><td>               <span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td>                     <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr></tbody></table>

### Using the Integration Hub

Use the [Integration Hub](../getting-started/snyk-web-ui.md#manage-integrations-for-asset-discovery-asset-coverage-and-issues-from-third-party-vendors) page to onboard integrations and populate Snyk AppRisk with data from SCM tools.

You can add an integration by following these steps:

1. Open AppRisk and navigate to the Integrations page.
2. Click **Integration Hub**.
3. Click **Add** on the integration you want to connect.
4. Configure your connection and click **Done**.

See the [Snyk SCM integrations](../scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) page or the [Third-party integrations for Snyk AppRisk](third-party-integrations-for-snyk-apprisk.md) page for step-by-step details about how to set up an integration.

After the integration is validated, a card is displayed on the Integrations page, allowing you to enable or disable the connection, edit the settings, or remove the connection from your configuration.

<figure><img src="../.gitbook/assets/image (11) (4).png" alt="AppRisk - Integration status" width="375"><figcaption><p>Snyk AppRisk - Integration status</p></figcaption></figure>

### Using Snyk Broker

If your SCM instance is not publicly accessible, you need Snyk Broker. You can install and configure your Snyk Broker using Docker or Helm. For more information about Snyk Broker, see the Snyk Broker documentation, including [Snyk Broker - AppRisk](../enterprise-configuration/snyk-broker/snyk-broker-apprisk.md).

{% hint style="info" %}
Enable the Snyk AppRisk flag in your Snyk Broker deployment environment before running the commands.
{% endhint %}

* GitHub - install and configure Snyk Broker
  * [using Docker](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/broker-example-set-up-snyk-broker-with-github.md#docker-run-command-to-set-up-a-broker-client-for-github)
  * [using Helm](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/githhub.com-install-and-configure-using-helm.md)
  * [environment variables](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-environment-variables-for-snyk-broker.md)
* GitHub Enterprise - install and configure Snyk Broker:
  * [using Docker](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-github-enterprise.md#docker-run-command-to-set-up-a-broker-client-for-github-enterprise)
  * [using Helm](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-helm.md)
  * [environment variables](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-environment-variables-for-snyk-broker.md)
* BitBucket - install and configure Snyk Broker:
  * [using Docker](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md#docker-run-command-to-set-up-a-broker-client-for-bitbucket)
  * [using Helm](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-helm.md)
  * [environment variables](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker-basic-auth.md)
* GitLab - install and configure Snyk Broker:
  * [using Docker](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-gitlab.md#docker-run-command-to-set-up-a-broker-client-for-gitlab)
  * [using Helm](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-install-and-configure-using-helm.md)
  * [environment variables](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-environment-variables-for-snyk-broker.md)
* Azure - install and configure Snyk Broker:
  * [using Docker](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md#docker-run-command-to-set-up-a-broker-client-for-azure-repos)
  * [using Helm](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-install-and-configure-and-configure-using-helm.md)
  * [environment variables](../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-environment-variables-for-snyk-broker.md)

You can find on [GitHub](https://github.com/snyk/broker/tree/565242baf003f06f445489dd96cc68c8386ede38/defaultFilters/apprisk) all the updated `.json` files that include the allowed list of accessible endpoints for the integrations.
