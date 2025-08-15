# Tenant, Groups, and Organizations

{% hint style="info" %}
**Feature availability**\
Snyk Groups are available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

## The Snyk hierarchy

Snyk has a hierarchy that allows you to control access to Snyk scanning and features. This varies for Free, Team, and Enterprise plans.

<figure><img src="../../.gitbook/assets/Enterprise.png" alt=""><figcaption><p>The Snyk hierarchy for Enterprise plans</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/2024-12-13_10-54-22.png" alt=""><figcaption><p>The Snyk hierarchy for Free and Team plans.</p></figcaption></figure>

* **Account:** Users must log in to their Snyk account to scan and view or modify any settings and scan
* [**Tenants**](tenant/): A Tenant encompasses the entire Snyk workspace of your company, team, and individual users. You have one Tenant that encompasses all your Snyk work items: Groups, Organizations, Targets, Projects, and all their adjacent entities, for example, Snyk features, Tags, Collections, and so on.
* [**Groups**](groups/): A Group encompasses your entire base of Snyk users. Large companies may have multiple Groups with multiple Organizations. Enterprise customers have at least one Snyk Group.
* [**Organizations**](organizations/): An Organization represents a specific area, such as a team, in your business. Organizations can contain multiple Projects.
* [**Targets**:](../../discover-snyk/getting-started/glossary.md#target) A Target represents the external resource that Snyk scans, like a repository. One Target can relate to multiple Projects. For example, a Target `https://github.com/examplesnyk/example` contains the Projects `package.json` and `Dockerfile.`
* [**Projects**](../snyk-projects/)**:** A Project is established based on the item that Snyk scans for issues, such as a manifest file. Each Project shows the results of scans. You can configure your Projects to define how to scan for issues in that Project.

## Snyk features for user management

To manage users in your Tenant, Organizations, and Groups:

* You can use the Snyk API v1 to [provision users to Orgs](../user-management-with-the-api/provision-users-to-organizations-using-the-api.md) and [remove members from Groups and Orgs](../user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md).
* To find out when a new user was added or to analyze unexpected activity, you can [retrieve audit logs of user-initiated activity](../user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md) by Organization or Group through the Snyk REST API.
* You can [use Organization access requests](organizations/requests-for-access-to-an-organization.md) to add users and [configure session length for a Snyk Group](groups/configure-session-length-for-a-snyk-group.md).
