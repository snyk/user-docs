# Azure DevOps

{% include "../../../../../.gitbook/includes/pilot-guide-toc.md" %}

Review the steps below to configure the Azure DevOps integration with Snyk. For more details about setting up the GitHub integration, contact your Snyk account team.

## Generate an Azure DevOps PAT&#x20;

Generate an Azure DevOps PAT  with the following permissions enabled:\


* `Code` - Read and Write
* `Project and Team` - Read
* `Graph` - Read
* `Analytics` - Read
* `Release` - Read
* `Member Entitlement Management` - Read

{% hint style="info" %}
Save the PAT details since the PAT is required in two places, at the Group-level integration and the Organization-level integration.
{% endhint %}

## Configure the Group-level integration

Configure the Group-level integration by following these steps:

* Open the Snyk Web UI
* Navigate to the Group-level
* Open **Integrations**, then **Add integration**

<figure><img src="../../../../../.gitbook/assets/image (308).png" alt=""><figcaption></figcaption></figure>

* Search and select the Azure DevOps integration
* Configure the integration and populate all mandatory fields, including the PAT details. For more details, see the [Integrate Azure DevOps using Snyk Essentials](../../../../../developer-tools/scm-integrations/group-level-integrations/azure-devops-for-snyk-essentials.md#azure-devops-integrate-using-snyk-apprisk) page.
* If relevant, you can also include the Backstage catalog. See the [Backstage file for SCM integrations](../../../../../developer-tools/scm-integrations/application-context-for-scm-integrations/#backstage-file-for-scm-integrations) page for more details.

{% hint style="info" %}
After the integration is configured, the Group-level integration shifts to a **Partially connected** status. During the next synchronization, it will transition to the connected state, and the Inventory view will be filled with data from the GitHub source.
{% endhint %}

## Configure the Organization-level integration

Configure the Organization-level integration by following these steps:

* Open the Snyk Web UI
* Navigate to the Organization-level
* Open **Integrations**
* Search and select the Azure DevOps integration
* Configure the integration and populate all mandatory fields, including the PAT details. For more details, see the [Azure DevOps integration settings](../../../../../developer-tools/scm-integrations/organization-level-integrations/azure-repositories-tfs.md#integrate-using-the-snyk-web-ui) page.

<figure><img src="../../../../../.gitbook/assets/image (303).png" alt=""><figcaption></figcaption></figure>

The Organization-level integration is immediately available to import repositories and begin scanning.
