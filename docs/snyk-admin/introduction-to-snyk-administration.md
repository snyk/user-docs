# Introduction to Snyk administration

{% hint style="info" %}
**Feature availability**\
Some functions, such as Groups, are available only on some [pricing plans](https://snyk.io/plans/).
{% endhint %}

## Accounts, Groups, Organizations, Targets, and Projects

Snyk has a hierarchy that controls access to scanning and other Snyk features.

* **Account:** Users must log in to their Snyk account to scan and view or modify any settings and scan output.
* [**Groups**](manage-groups-and-organizations/introduction-to-groups.md)**:** Customers have at least one Snyk Group. The Group may correspond to the entire company. Large companies may have multiple Groups, corresponding to divisions, departments, or other parts of the company. Groups can contain multiple Organizations.
* [**Organizations**](manage-groups-and-organizations/whats-a-snyk-organization.md)**:** Snyk Groups encompass one or more Snyk Organizations. Organizations represent specific business areas, such as teams. Organizations can contain multiple Projects.
* **Targets**. Each Target represents a repository imported into Snyk for scanning and re-testing.
* [**Projects**](snyk-projects/)**:** A Project is established based on the items that Snyk scans for issues, such as manifest files, and shows the results of scans. You can configure your Projects to define how to scan for issues in that Project. See the [Quickstart](../getting-started/quickstart/) for the basic steps.

Snyk administrators set up Groups and Organizations. See [Manage Groups and Organizations](manage-groups-and-organizations/) for details. Targets and [Projects](snyk-projects/) are created when Snyk users import development Projects into Snyk for scanning.

## Member (user) types

Snyk has the following types of members (users), defined by their roles:

* Collaborator
* Organization administrator
* Group member
* Group administrator

{% hint style="info" %}
**Feature availability**\
Group administrators and Collaborators are available with Snyk Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.&#x20;
{% endhint %}

For details of the permissions associated with each role, see [Permissions associated with each role ](manage-users-and-permissions/managing-permissions.md#permissions-associated-with-each-role)on the Managing permissions page.

## Snyk Admin tools

Snyk provides tools to manage Groups, Organizations, and user roles and permissions, as well as notifications and settings.

### Manage users and permissions

You can manage users and permissions in your groups. For details, see [Manage users and permissions](manage-users-and-permissions/).

<figure><img src="../.gitbook/assets/image (245) (1) (1).png" alt="Manage members interface"><figcaption><p>Manage members interface</p></figcaption></figure>

### Manage Groups and Organizations

Snyk groups and organizations help to maintain collaboration across teams. For details, see [Manage Groups and Organizations](manage-groups-and-organizations/).

### Define notifications

You can manage email notifications for yourself and your Organization. For details, see [Manage notifications](manage-notifications.md).

<figure><img src="../.gitbook/assets/image (6) (2).png" alt="Manage email notifications interface"><figcaption><p>Manage email notifications interface</p></figcaption></figure>

### Manage settings

You can customize your Snyk account to suit your work process. For details, see [Manage settings](manage-settings/).
