# BitBucket

{% include "../../../../.gitbook/includes/pilot-guide-toc.md" %}

‌Review the steps below to configure the Bitbucket integration with Snyk. For more details about setting up the GitHub integration, contact your Snyk account team.\


## Generate a BitBucket PAT&#x20;

Generate a BitBucket PAT with the following permissions enabled:

* `Account` - Read
* `Workspace membership` - Read
* `Projects` - Read
* `Repositories` - Read
* `Pull request` - Read and Write
* `Issues` - Read
* `Webhooks` - Read and Write

{% hint style="info" %}
Save the PAT details since the PAT is required in two places, at the Group-level integration and the Organization-level integration.
{% endhint %}

## Configure the Group-level integration

Configure the Group-level integration by following these steps:

* Open the Snyk Web UI
* Navigate to the Group-level
* Open **Integrations**, then **Add integration**

<figure><img src="../../../../.gitbook/assets/image (32).png" alt=""><figcaption></figcaption></figure>

* Search and select the BitBucket integration
* Configure the integration and populate all mandatory fields, including the PAT details. For more details, see the [Integrate BitBucket using Snyk Essentials](../../../../scm-integrations/group-level-integrations/bitbucket-for-snyk-essentials.md) page.
* If relevant, you can also include the Backstage catalog. See the [Backstage file for SCM integrations](../../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#backstage-file-for-scm-integrations) page for more details.

{% hint style="info" %}
After the integration is configured, the Group-level integration shifts to a **Partially connected** status. During the next synchronization, it will transition to the connected state, and the Inventory view will be filled with data from the GitHub source.
{% endhint %}

## Configure the Organization-level integration

Configure the Organization-level integration by following these steps:

* Open the Snyk Web UI
* Navigate to the Organization-level
* Open **Integrations**
* Search and select the BitBucket integration
* Configure the integration and populate all mandatory fields, including the PAT details. For more details, see the [BitBucket integration settings](../../../../scm-integrations/organization-level-integrations/bitbucket-cloud.md) page.

<figure><img src="../../../../.gitbook/assets/image (30).png" alt=""><figcaption></figcaption></figure>

The Organization-level integration is immediately available to import repositories and begin scanning.
