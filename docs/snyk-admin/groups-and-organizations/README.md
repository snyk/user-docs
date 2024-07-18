# Groups and Organizations

{% hint style="warning" %}
**Feature availability**\
Snyk Groups are only available on certain [pricing plans](https://snyk.io/plans/).
{% endhint %}

## Groups, Organizations, and Projects

Snyk has a hierarchy that allows you to control access to Snyk scanning and features. Key parts of the hierarchy are illustrated in this diagram:

<figure><img src="../../.gitbook/assets/image (1) (1) (1) (2).png" alt="Groups, Organizations and Projects"><figcaption><p>Groups, Organizations and Projects</p></figcaption></figure>

* [**Groups**](groups/): Often, a Group encompasses your entire base of Snyk users, although large companies may have multiple Groups. Groups can contain multiple Organizations.
* [**Organizations**](organizations/): An Organization represents a specific business area, such as a team. Organizations can contain multiple Projects.
* [**Projects**](../snyk-projects/)**:** A Project is established based on the item that Snyk scans for issues, such as a manifest file. Each Project shows the results of scans. You can configure your Projects to define how to scan for issues in that Project.

Snyk also has features to [manage users in Organizations](organizations/manage-users-in-organizations.md) and to [manage users in a Group](groups/manage-users-in-a-group.md). You can use the Snyk API v1 to [provision users to Orgs](../user-management-with-the-snyk-api/provision-users-to-organizations-using-the-v1-api.md) and [remove members from Groups and Orgs](../user-management-with-the-snyk-api/remove-members-from-groups-and-orgs-using-the-api.md).

You can [use Organization access requests](organizations/requests-for-access-to-an-organization.md) to add users and [configure session length for a Snyk Group](groups/configure-session-length-for-a-snyk-group.md).

When you want to find out when a new user was added or analyze unexpected activity, you can [retrieve audit logs of user-initiated activity](../user-management-with-the-snyk-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md) by Organization or Group through the Snyk REST API.

The documentation in this section covers [Groups](groups/) and [Organizations](organizations/).
