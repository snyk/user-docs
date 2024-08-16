# Snyk Admin

Snyk administration encompasses the following functions:

* [Manage Groups and Organizations](groups-and-organizations/)
* [Manage and use Snyk Projects](snyk-projects/)
* [Manage users in Organizations](groups-and-organizations/organizations/manage-users-in-organizations.md) and [Groups](groups-and-organizations/groups/manage-users-in-a-group.md)
* [Manage user roles](user-roles/)
* [Manage notifications](manage-notifications.md)
* [Manage settings](groups-and-organizations/group-and-organization-settings.md)

This page covers the following topics:

* [Accounts, Groups, Organizations, Targets, and Projects](./#accounts-groups-organizations-targets-and-projects)
* [User types](./#user-types)
* [Snyk Admin tools](./#snyk-admin-tools)

{% hint style="warning" %}
**Feature availability**

Some functions, such as Groups, are available only on some [pricing plans](https://snyk.io/plans/).\
See [Enterprise setup](../enterprise-configuration/) for information about set up for Enterprise plan customers.
{% endhint %}

## Accounts, Groups, Organizations, Targets, and Projects

Snyk has a hierarchy that controls access to scanning and other Snyk features.

* **Account:** Users must log in to their Snyk account to scan and view or modify any settings and scan output.
* [**Groups**](groups-and-organizations/groups/)**:** Customers have at least one Snyk Group. The Group may correspond to the entire company. Large companies may have multiple Groups, corresponding to divisions, departments, or other parts of the company. Groups can contain multiple Organizations.
* [**Organizations**](groups-and-organizations/organizations/)**:** Snyk Groups encompass one or more Snyk Organizations. Organizations represent specific business areas, such as teams. Organizations can contain multiple Projects.
* [**Targets**](snyk-projects/#target)**:** Each Target represents a repository imported into Snyk for scanning and re-testing.
* [**Projects**](snyk-projects/)**:** A Project is established based on the items that Snyk scans for issues, such as manifest files, and shows the results of scans. You can configure your Projects to define how to scan for issues in that Project. See the [Quickstart](../getting-started/quickstart/) for the basic steps.

Snyk Admins set up Groups and Organizations. See [Manage Groups and Organizations](groups-and-organizations/) for details. Targets and [Projects](snyk-projects/) are created when Snyk users import development Projects into Snyk for scanning.

## User types

Snyk has the following types of pre-defined users:

* Organization Admin
* Organization Collaborator
* Group Admin
* Group Viewer
* Group Member

See [Pre-defined roles](user-roles/pre-defined-roles.md) for more details, including the permissions associated with each role.

{% hint style="warning" %}
**Feature availability**

Group-level roles are available only for Enterprise plans.

See [Pricing plans](https://snyk.io/plans).
{% endhint %}

## Snyk Admin tools

Snyk provides tools to manage Groups, Organizations, and user roles and permissions, as well as notifications and settings.

### Manage users and permissions

You can manage users and permissions in your Groups. For details, see [Manage users and permissions](user-roles/user-role-management.md).

<figure><img src="../.gitbook/assets/image (245) (1) (1) (1).png" alt="Manage members interface"><figcaption><p>Manage members interface</p></figcaption></figure>

### Manage Groups and Organizations

Snyk groups and organizations help to maintain collaboration across teams. For details, see [Manage Groups and Organizations](groups-and-organizations/).

### Define notifications

You can manage email notifications for yourself and your Organization. For details, see [Manage notifications](manage-notifications.md).

<figure><img src="../.gitbook/assets/image (6) (2).png" alt="Manage email notifications interface"><figcaption><p>Manage email notifications interface</p></figcaption></figure>

### Manage settings

You can customize your Snyk account to suit your work process. For details, see [Manage settings](groups-and-organizations/group-and-organization-settings.md).
