---
description: How to configure the GitLab integration with Snyk, including generating a personal access token with the required permissions
nav_context: new
---

# GitLab

{% include "../../../../.gitbook/includes/pilot-guide-navigation.md" %}

Configure the GitLab integration at both the Group level and the Organization level. For more details about setting up the GitLab integration, contact your Snyk account team.

## Generate a GitLab PAT

Generate a GitLab PAT with the following permissions:

* `api`
* `read_api`
* `read_repository`

{% hint style="info" %}
Save the PAT details. You need the PAT for both the Group-level integration and the Organization-level integration.
{% endhint %}

## Configure the Group-level integration

1. In the Snyk Web UI, use the scope selector at the top of the page to switch to your Group.
2. Navigate to **Settings** > **Integrations** > **All integrations**, then click **Add integration**.

<figure><img src="../../../../.gitbook/assets/configure-group-level-integration.png" alt="Configuring the GitLab integration at the Group level"><figcaption></figcaption></figure>

3. Search for and select the GitLab integration.
4. Complete all mandatory fields, including the PAT details. For more details, visit [Integrate GitLab using Snyk Essentials](https://docs.snyk.io/developer-tools/integrations/scm-integrations/group-level-integrations/gitlab-for-snyk-essentials#gitlab-integrate-using-snyk-apprisk).
5. Optional: include the Backstage catalog. For details, visit [Backstage file for SCM integrations](https://docs.snyk.io/developer-tools/integrations/scm-integrations/application-context-for-scm-integrations#backstage-file-for-scm-integrations).

{% hint style="info" %}
After you configure the integration, the Group-level integration shows a **Partially connected** status. During the next synchronization, the integration moves to the connected state, and Snyk fills the Inventory with data from GitLab.
{% endhint %}

## Configure the Organization-level integration

1. In the scope selector, open the **Organization** dropdown and select your Organization.
2. Navigate to **Settings** > **Integrations** > **All integrations**.
3. Search for and select the GitLab integration.
4. Complete all mandatory fields, including the PAT details. For more details, visit [GitLab integration settings](https://docs.snyk.io/developer-tools/integrations/scm-integrations/organization-level-integrations/gitlab).

<figure><img src="../../../../.gitbook/assets/configure-organization-level-integration.png" alt="Configuring the GitLab integration at the Organization level"><figcaption></figcaption></figure>

The Organization-level integration is available immediately, so you can import repositories and start scanning.
