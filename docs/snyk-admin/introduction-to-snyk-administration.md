# Introduction to Snyk administration

{% hint style="info" %}
**Feature availability**\
Some functions, such as Groups, re only available on some [pricing plans](https://snyk.io/plans/).
{% endhint %}

## Accounts, Groups, Organizations, and Projects

Snyk has a hierarchy that controls access to scanning and other Snyk features.

* **Account:** Users must log in to their Snyk account to scan and to view or modify any settings and scan output.
* **Group:** Customers have at least one Snyk Group. The Group may correspond to the entire company. Large companies may have multiple Groups, corresponding to divisions, departments, or other parts of the company.
* **Organization:** Snyk Groups encompass one or more Snyk Organizations. An Organization may correspond to a team or types of Projects where multiple teams are involved.

Within an Organization, after you have scanned code, you will find one or more Snyk **Targets**. Each represents a repository imported into Snyk for scanning and re-testing. Targets contain one or more Snyk **Projects**. Each corresponds to the files scanned by Snyk.

Snyk administrators set up Groups and Organizations. See [Manage Groups and Organizations](manage-groups-and-organizations/) for details.

## Member (user) types

Snyk provides four different types of members (users):

* Collaborator
* Organization administrator
* Group member
* Group administrator

For definitions of the associated permissions per role, see [Managing permissions](manage-users-and-permissions/managing-permissions.md#permissions-per-role).

{% hint style="info" %}
**Feature availability**\
Group administrators and collaborators are available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.&#x20;
{% endhint %}

Snyk offers a wide range of tools to manage Groups, Organizations, and user

#### Manage users and permissions

You can manage users and permissions in your groups; See [Managing users & permissions](manage-users-and-permissions/).

<figure><img src="../.gitbook/assets/image (245) (1) (1) (1).png" alt="Manage users and permissions"><figcaption><p>Manage users and permissions</p></figcaption></figure>

#### Manage Groups and Organizations

Learn how Snyk groups and organizations help keep cross-team collaboration seamless; see [Managing groups & organizations](manage-groups-and-organizations/).

#### Define notifications

You can manage email notifications, for yourself and your organization. See [Managing notifications](manage-notifications.md).

<figure><img src="../.gitbook/assets/image (6) (2).png" alt="Manage email notifications"><figcaption><p>Manage email notifications</p></figcaption></figure>

#### Manage settings

Customize your Snyk account for your needs. See [Managing settings](manage-settings/).
