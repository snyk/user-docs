---
description: Snyk trial and Free plan limitations, including the 14-day trial and full Enterprise pilot options
nav_context: agnostic
---

# Trial limitations

You can try out Snyk functionalities in several ways:

* Using the limited Free plan
* Using a self-serve 14-day trial
* By piloting the full product with the Enterprise plan

The Snyk 14-day trial offers a sample of the features available in the paid Enterprise plan. However, certain features will have limited functionality or be entirely unavailable in order to provide a seamless experience when the trial concludes.

{% hint style="info" %}
After implementing Snyk Essentials, you must access the Inventory page to ensure it is populated with all the necessary information. Depending on the number of repositories you have imported, the update may take up to several hours.
{% endhint %}

If you consider purchasing the Team plan, remember that the trial offers features beyond those included in the Team plan.

The following Enterprise features are limited or unavailable during the trial:

* [**Single Sign-On (SSO)**](https://docs.snyk.io/platform-administration/user-management/single-sign-on-sso-for-authentication-to-snyk): Not available.
* [**Service Accounts**](https://docs.snyk.io/platform-administration/service-accounts/service-accounts): Not available.
* [**Group / Multiple Orgs**](https://docs.snyk.io/platform-administration/snyk-hierarchy/groups-and-organizations): Limited to one Group and one Organization
* [**Custom Project Tags**](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-projects/project-tags): Not available.
* [**Custom User Roles**](https://docs.snyk.io/platform-administration/user-management/user-role-management): Not available.
* [**Audit Logging**](https://docs.snyk.io/platform-administration/user-management/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group): Not available.
* [**Broker**](https://docs.snyk.io/platform-administration/snyk-broker/snyk-broker): Not available.
* [**Self-hosted Git**](https://docs.snyk.io/developer-tools/integrations/scm-integrations/organization-level-integrations/github-enterprise): Not available.
* [**Private Registry Integrations**](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/package-repository-integrations): Not available.
* [**Snyk Essentials Inventory**](https://docs.snyk.io/scan-fix-and-prevent/fix/manage-assets): Not available
* [**Snyk Essentials Issue page**](https://docs.snyk.io/scan-fix-and-prevent/fix/prioritize-issues-for-fixing/prioritization-for-snyk-essentials): Limited capabilities, the Funnel view, Risk factors, and evidence graph, will not be available.

## Inactive Free plan Tenant deletion

A Free plan Tenant that has had no activity for more than 60 days is slated for deletion. Snyk determines activity based on its own records. One login, CLI command, IDE session, PR check, or fix PR in that period keeps your Tenant active and resets the 60-day period.

The following are not subject to this evaluation for deletion:

* Paid plans and active subscriptions
* Active pilots and trials
* The [Secure Developer Program](https://snyk.io/open-source/)

If a pilot or trial ends without a purchase, this policy applies to the Tenant. If a paid subscription ends and your Tenant moves to the Free plan, this policy also applies to it. For a Tenant that moves from a paid plan to the Free plan, the 60-day period starts no earlier than the date of the plan change. Tenants created in the last 60 days are not deleted under this policy.

### What counts as activity

Any one of the following actions resets the 60-day period. Any member can perform the action, in any Organization of your Tenant, in any Snyk product, such as Snyk Open Source, Snyk Code, Snyk Container, or Snyk IaC.

| Activity | Description |
| --- | --- |
| [Web UI login](https://app.snyk.io/login) | Someone logs in to the Snyk Web UI. |
| [Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/getting-started-with-the-snyk-cli) | Someone runs any CLI command while authenticated to your Tenant, such as `snyk test`, `snyk monitor`, `snyk code test`, `snyk container test`, or `snyk iac test`. CI pipelines that run the CLI with your Tenant token also count. |
| [Snyk IDE plugin](https://docs.snyk.io/developer-tools/snyk-ide-plugins-and-extensions) | Someone uses the plugin while authenticated to your Tenant. |
| [PR checks and fix PRs](https://docs.snyk.io/scan-fix-and-prevent/fix/pull-requests) | Snyk runs a PR check or opens a fix PR through your GitHub or other SCM integration. |

Scheduled or recurring scans of imported Projects do not count as activity. After you import a Project, Snyk automatically retests it on a schedule, for example, daily or weekly. Snyk starts these tests itself, so they do not show that a person is using your Tenant. If automatic retests are all you use, log in or run a CLI command at least once every 60 days.

### After a Tenant is deleted

When a Free plan Tenant is slated for deletion, Snyk aims to permanently delete it. Integrations and pipelines that authenticate to the deleted Tenant stop working. A deleted Tenant cannot be restored.

Snyk is not responsible for any data lost through the deletion of an inactive Free plan Tenant under this policy.

To keep using Snyk after a deletion, [create a new Free account](https://app.snyk.io/signup), import your Projects again, and generate a new API token.

### Terms of Service

Your use of the Free plan is governed by the [Snyk Terms of Service](https://snyk.io/policies/terms-of-service/) or, where applicable, another governing agreement entered into between you and Snyk ("Terms of Service"). This section describes how Snyk handles inactive Free plan Tenants. It does not limit any right Snyk has under the Terms of Service, including the right to suspend, terminate, or delete a Free plan Tenant at any time, with or without notice.
