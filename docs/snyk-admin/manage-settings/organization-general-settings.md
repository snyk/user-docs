# Organization general settings

To view and manage general settings for your Organization: in your Organization, navigate to **Settings > ORGANIZATION SETTINGS > General:**

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-02 at 15.39.28.png" alt="View Organization general settings"><figcaption><p>View Organization general settings</p></figcaption></figure>

</div>

* [Organization name](organization-general-settings.md#organization-name): set the display name for your Organization.
* [Organization API key](organization-general-settings.md#organization-api-key): used for Service accounts.
* [Organization ID:](organization-general-settings.md#organization-id) auto-generated internal ID for that Organization.
* [Ignores](organization-general-settings.md#ignores): define Organization-wide ignore settings.
* [Requesting access](organization-general-settings.md#requesting-access): allow external users to request access.
* [Leave Organization](organization-general-settings.md#leave-organization): leave this Organization.
* [Delete Organization](organization-general-settings.md#delete-organization): delete this Organization.

{% hint style="info" %}
Also see [Manage Organizations](../manage-groups-and-organizations/manage-organizations.md).
{% endhint %}

### Organization name

Set the display name for this Organization For example, you may want to change the default initial Organization display name, to reflect better the work that Organization relates to.

{% hint style="warning" %}
Changing this Organization display name does not change the URL displayed in the Web UI browser bar, or the internal name used by the CLI. To fully change the name, [create a new Organization](https://app.snyk.io/create-organisation) with the desired name.
{% endhint %}

### Organization API key

[Service accounts](../../enterprise-setup/service-accounts.md) require an API key (token) to substitute for standard user credentials. Click **Manage Service accounts** to set up details for these accounts.

### Organization ID

This ID uniquely identifies this Organization; required if you use the [Snyk API](../../snyk-api/). These Organization identification details are generated automatically by Snyk for each Organization when it is created&#x20;

### Ignores

Use these functions to define Organization-wide ignore settings.

<div align="left">

<figure><img src="../../.gitbook/assets/image (1) (2) (3).png" alt="Configure ignore settings for an Organization"><figcaption><p>Configure ignore settings for an Organization</p></figcaption></figure>

</div>

See **Configure Ignore settings** in the [Ignore issues](../../manage-issues/priorities-for-fixing-issues/ignore-issues.md#configure-ignore-settings) documentation for more details.

### Requesting Access

If enabled, users without access to this Snyk Organization can request access.

<div align="left">

<figure><img src="../../.gitbook/assets/image (11) (3).png" alt="Enable / disable request access"><figcaption><p>Enable / disable request access</p></figcaption></figure>

</div>

See **Settings** in the [Organization access requests](../manage-users-and-permissions/organization-access-requests.md#settings) documentation for more details_._

### Leave Organization

Click **Leave Organization** to leave an Organization, losing access to the Projects and notifications for that Organization.

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-02 at 16.55.50.png" alt="Leave Organization"><figcaption><p>Leave Organization</p></figcaption></figure>

</div>

### Delete Organization

Deleting an Organization will remove it entirely from Snyk, including all of its Projects and their historical data.

See **Delete an Organization** in the [Manage Organizations](../manage-groups-and-organizations/manage-organizations.md#delete-an-organization) documentation for details.
