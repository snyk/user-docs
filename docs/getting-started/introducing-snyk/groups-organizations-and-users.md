# Groups, organizations, and users

{% hint style="info" %}
**Feature availability**\
Some functions (such as Groups) are only available on certain plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk has a hierarchy that allows you to control access to features such as reports. This hierarchy is as follows:

* **Group:** the highest level; for example, the entire company.
* **Organization:** the second level of grouping; for example, your team.
* **Projects:** the lowest level for individual projects; for example, a container image.

### Snyk groups

Typically, a Snyk group represents the entire company or business division.

Groups can contain multiple organizations, allowing you to collaborate with multiple teams.

See [What’s a Snyk group?](../../features/user-and-group-management/managing-groups-and-organizations/whats-a-snyk-group.md)

### Snyk organizations

Organizations are contained in groups. Based on your company requirements, you can define organizations to represent business areas such as teams, products or environments.

Organizations can contain multiple projects. For example, if an organization represents an engineering team, this allows each team to see the applications they are working on.

{% hint style="info" %}
When you sign up to Snyk using a social login, you have a default organization. Any projects you add appear in this organization by default.
{% endhint %}

See [What’s a Snyk organization?](../../features/user-and-group-management/managing-groups-and-organizations/whats-a-snyk-organization.md)

{% hint style="info" %}
[Projects](projects.md) are contained in Organizations.
{% endhint %}

## User access: member types

Snyk provides four different types of members or users:

* Collaborator
* Organization administrator
* Group member
* Group administrator

For definitions of the associated permissions, see this [information on permissions per role](https://docs.snyk.io/features/user-and-group-management/managing-users-and-permissions/managing-permissions#permissions-per-role).

{% hint style="info" %}
**Feature availability**\
Group administrators and collaborators are available with Business and Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

More on [user and group management](../../features/user-and-group-management/).
