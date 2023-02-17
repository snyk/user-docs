# Snyk admin

## Introduction

{% hint style="info" %}
**Feature availability**\
Some functions (such as Groups) are only available on certain plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk has a hierarchy that allows you to control access to features such as reports. This hierarchy is as follows:

* **Group:** the highest level; for example, the entire company.
* **Organization:** the second level of grouping; for example, your team.
* **Projects:** the lowest level for individual projects; for example, a container image.

#### Snyk Groups

Typically, a Snyk group represents the entire company or business division.

Groups can contain multiple organizations, allowing you to collaborate with multiple teams.

See [What’s a Snyk group?](managing-groups-and-organizations/whats-a-snyk-group.md)

#### Snyk Organizations

Organizations are contained in Groups. Based on your company's requirements, you can define Organizations to represent business areas such as teams, products or environments.

Organizations can contain multiple Projects. For example, if an Organization represents an engineering team, this allows each team to see the applications they are working on.

{% hint style="info" %}
When you sign up to Snyk using a social login, you have a default organization. Any projects you add appear in this organization by default.
{% endhint %}

See [What’s a Snyk organization?](managing-groups-and-organizations/whats-a-snyk-organization.md)

#### Snyk Projects

Snyk Projects are contained in Organizations.

{% hint style="info" %}
[Projects](../features/user-and-group-management/broken-reference/) are contained in Organizations.
{% endhint %}

### Snyk users: member types

Snyk provides four different types of members or users:

* Collaborator
* Organization administrator
* Group member
* Group administrator

For definitions of the associated permissions, see this [information on permissions per role](https://docs.snyk.io/features/user-and-group-management/managing-users-and-permissions/managing-permissions#permissions-per-role).

{% hint style="info" %}
**Feature availability**\
Group administrators and collaborators are available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

### Management tools

Snyk offers a wide range of tools to manage Groups, Organizations, and users.

#### Enable authentication

Snyk offers API tokens to enable authentication to service accounts or third party tools; see [authentication](authentication/ "mention").

#### Set up Single Sign-On (SSO)

SSO makes authentication and provisioning simple; see [setting-up-sso-for-authentication](setting-up-sso-for-authentication/ "mention").

![](<../.gitbook/assets/image (167) (1) (1) (1) (1) (1) (1) (1) (2).png>)

#### Manage users and permissions

You can manage users and permissions in your groups; See [managing-users-and-permissions](managing-users-and-permissions/ "mention").

![](<../.gitbook/assets/image (245) (1).png>)

#### Manage Groups and Organizations

Learn how Snyk groups and organizations help keep cross-team collaboration seamless; see [managing-groups-and-organizations](managing-groups-and-organizations/ "mention")

#### Define notifications

You can manage email notifications, for yourself and your organization. See [notifications.md](notifications.md "mention").

![](<../.gitbook/assets/image (200) (1) (1).png>)

#### Manage settings

Customize your Snyk account for your needs. See [managing-settings](managing-settings/ "mention").

![](<../.gitbook/assets/image (118) (1) (1) (1) (1) (2) (1) (1) (1) (1) (1).png>)
