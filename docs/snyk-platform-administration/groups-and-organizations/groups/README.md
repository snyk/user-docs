# Groups

Snyk Groups make it easier for you to work in teams. Groups can contain many Organizations, and each Organization can contain many collaborators and Projects.

As part of your onboarding, Snyk sets you up with a Group for your company. You can then add your current Organizations to this Group. Enterprise plan customers may have more than one Group. If you think your company will need multiple Groups, submit a request to [Snyk Support](https://support.snyk.io).

{% hint style="info" %}
Visit [Structure your account for high application performance](../../structure-your-account-for-high-application-performance.md) for details concerning multiple Groups.
{% endhint %}

## Group-level options

Use Group-level options to view [Organizations](./#group-organizations), [reports](./#group-reports), [dependencies](./#group-dependencies), and [policies](./#group-policies) across all of the Organizations in your Group, configure your [Group settings](./#group-settings), and view all the [users](./#group-members) in a Group.

### Group Organizations

Select **Organizations** to view all the Organizations you have access to and your assigned role inside each one:

If your Group is set up to let its users join Organizations, you also see a list of all the Organizations in the Group and options to join the Organizations where you are not a member.

{% hint style="info" %}
For a detailed breakdown of user roles and their associated access permissions, see [User roles](../../user-roles/).
{% endhint %}

### Group reports

Select [**Reports**](../../../manage-risk/analytics/reports-tab/) to view the vulnerability status of the Organizations in your Group in one place as a report.

### Group inventory

Select [**Inventory**](../../../manage-assets/manage-assets.md#inventory-menu) to view, filter, and manage your assets.

### Group issues

Select [**Issues**](../../../manage-risk/prioritize-issues-for-fixing/) to better identify and prioritize your Container, Code, and Open Source issues based on the risk they pose to your application. **Issues** offer a centralized view of all the issues identified by Snyk with additional asset context.

### Group dependencies

You can [view dependencies](../../../manage-risk/dependencies-and-licenses/view-dependencies.md) and [license information](https://docs.snyk.io/manage-risk/dependencies-and-licenses/view-licenses) for all Projects in your Group or Organization using the **Dependencies** option in your Group or Organization menu.

### Group Policies

With Policies, you can easily automate the process of adding business context and receiving notifications.

After a policy is created, it is run in a maximum of 3 hours after creation, then once every 3 hours.

If your policy is set to run daily, then the policy is run 3 hours after the 24-hour period ends. You can always manually run a policy by using the Run button.

### Group integrations

Select **Integrations** to view and configure the available integrations for [SCM integrations](../../../developer-tools/scm-integrations/group-level-integrations/).

### Group members

Select **Members** to view users in the Group.

Group members are users who have access to all Organizations in the Group. Users of the Organizations are managed in the Settings of each Organization.

<figure><img src="../../../.gitbook/assets/2024-04-02_09-41-48.png" alt=""><figcaption><p>Members view for Group users</p></figcaption></figure>

{% hint style="info" %}
See [Manage users in a Group](manage-users-in-a-group.md) for details.
{% endhint %}

### Group Settings

Select **Settings** to view and manage Group settings.

See [Manage settings](../group-and-organization-settings.md) for details.
