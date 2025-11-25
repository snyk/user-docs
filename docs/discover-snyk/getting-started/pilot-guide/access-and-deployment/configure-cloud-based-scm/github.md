# GitHub

{% include "../../../../../.gitbook/includes/pilot-guide-toc.md" %}

‌Review the steps below to configure the GitHub integration with Snyk. For more details about setting up the GitHub integration, contact your Snyk account team.

## Generate a GitHub PAT

Generate a GitHub PAT with the following permissions enabled:

* `repo`
* `read:org`
* `read:user`
* `user:email`
* `admin:repo_hook`

{% hint style="info" %}
Save the PAT details since the PAT is required in two places, at the Group-level integration and the Organization-level integration.
{% endhint %}

## Configure the Group-level integration

Configure the Group-level integration by following these steps:

* Open the Snyk Web UI
* Navigate to the Group-level
* Open **Integrations**, then **Add integration**

<figure><img src="broken-reference" alt=""><figcaption></figcaption></figure>

* Search and select the GitHub integration
* Configure the integration and populate all mandatory fields, including the PAT details. For more details, see the [Integrate GitHub using Snyk Essentials](../../../../../developer-tools/scm-integrations/group-level-integrations/github-for-snyk-essentials.md#github-integrate-using-snyk-apprisk) page.

{% hint style="info" %}
After the integration is configured, the Group-level integration shifts to a **Partially connected** status. During the next synchronization, it will transition to the connected state, and the Inventory view will be filled with data from the GitHub source.
{% endhint %}

## Configure the Organization-level integration

Configure the Organization-level integration by following these steps:

* Open the Snyk Web UI
* Navigate to the Organization-level
* Open **Integrations**
* Search and select the GitHub integration
* Configure the integration and populate all mandatory fields, including the PAT details. For more details, see the [GitHub integration settings](../../../../../developer-tools/scm-integrations/organization-level-integrations/github.md#github-integration-settings) page.

<figure><img src="../../../../../.gitbook/assets/image (303).png" alt=""><figcaption></figcaption></figure>

The Organization-level integration is immediately available to import repositories and begin scanning.

\\
