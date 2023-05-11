# View policies

{% hint style="warning" %}
You must be a Group administrator to access policies for that Group.
{% endhint %}

Select the **Policies** menu option to see the policies in your Group, arranged by category ([License policies](license-policies/) and [Security policies](security-policies/)).

<div align="left">

<figure><img src="../../.gitbook/assets/Policies-menu.png" alt="View policies"><figcaption><p>View policies</p></figcaption></figure>

</div>

Expand a category to see policies in that category:

<figure><img src="../../.gitbook/assets/snyk-policy-manager.png" alt="Policy manager screen"><figcaption><p>Policy manager screen</p></figcaption></figure>

{% hint style="info" %}
This list includes the [default policy](view-policies.md#default-policies), which is automatically created for new groups for each policy category and cannot be removed.
{% endhint %}

### Policy manager details

The Policy manager screen appears similar to the following:

<figure><img src="../../.gitbook/assets/screenshot_2021-03-26_at_11.04.50_am.png" alt="Policy manager screen details"><figcaption><p>Policy manager screen details</p></figcaption></figure>

## Default policies

Each policy category has its own default policy. Default policies can only be applied to Organizations, not Projects.

When you create a new Organization, it will automatically be added to the default policy unless you have selected to copy an existing Organization's settings. Organizations can be moved to a different policy if desired.

The default policy cannot be deleted; however, the default policy name, description, and rules can be edited to match your preferences. A default policy can also contain no rules if you'd prefer.

See [Assign a policy to Organizations](apply-a-policy-to-organizations.md) for more details.
