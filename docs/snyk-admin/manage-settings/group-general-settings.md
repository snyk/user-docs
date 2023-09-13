# Group general settings

To view and manage general settings for your Group, navigate to **Settings > GROUP SETTINGS > General:**

<figure><img src="../../.gitbook/assets/image (1) (3) (1).png" alt=""><figcaption><p>Group general settings</p></figcaption></figure>

* [Group name](group-general-settings.md#group-name)
* [Group ID](group-general-settings.md#group-id)
* [Session expiration](group-general-settings.md#session-expiration)
* [Requesting access](group-general-settings.md#requesting-access)
* [Project test frequency](group-general-settings.md#project-test-frequency)

### Group name

The name of the group as displayed in the Snyk Web UI.

### Group ID

The unique ID for this group, this is required if you use the [Snyk API](../../snyk-api/).

### Session expiration

By default, users are logged out of Snyk after 24 hours of inactivity. You can change the session expiration time for your group here. See [Session length](../manage-users-and-permissions/session-length.md) for more details.

### Requesting access

If enabled, users without access to a Snyk Organization can request access. This is only possible for users who have registered with Snyk, and when trying to reach the URL of a specific Project in the Organization, such as from a pull request. This restriction reduces the risk of someone guessing the URL of your Organization.

The value set is used as the default for any new Organizations, but does not override the **Requesting access** setting for existing [Snyk Organization general settings](organization-general-settings.md#requesting-access).&#x20;

See **Settings** in the [Organization access requests](../manage-users-and-permissions/organization-access-requests.md#settings) documentation for more details_._

### Project test frequency

Set the default test frequency for any new Projects created in this Group.

{% hint style="warning" %}
Changing this setting will not affect the default test frequency of [Snyk Infrastructure as Code](../../scan-infrastructure/snyk-infrastructure-as-code/) or [Snyk Code](../../scan-applications/snyk-code/) Projects. The default for both of these is weekly.
{% endhint %}



\
