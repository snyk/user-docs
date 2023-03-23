# Introduction to policies

**Prerequisites**

You must be a group administrator of the group in order to update **Policy** settings

{% hint style="info" %}
**Feature availability**\
This feature is available to Enterprise customers. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

## Navigate to the Policy manager

1. Log in to Snyk
2. Navigate to your group
3. Click on the Policies tab in the navigation bar to see all the policies that exist within your group, broken out by category. This list will include the [default policy](shared-policies-overview.md#default-policies), which is automatically created for new groups for each policy category and cannot be removed.

<figure><img src="../../.gitbook/assets/snyk-policy-manager.png" alt="Policy manager screen"><figcaption><p>Policy manager screen</p></figcaption></figure>

The Policy manager appears similar to the following:

<figure><img src="../../.gitbook/assets/screenshot_2021-03-26_at_11.04.50_am.png" alt="Policy manager screen details"><figcaption><p>Policy manager screen details</p></figcaption></figure>

## Default policies

Each policy category has its own default policy. Default policies can only be applied to Organizations, not Projects.

When you create a new Organization, it will automatically be added to the default policy unless you have selected to copy an existing Organization's settings. Organizations can be moved to a different policy if desired.

The default policy cannot be deleted; however, the default policy name, description, and rules can be edited to match your preferences. A default policy can also contain no rules if you'd prefer.

See [Assign a policy to Organizations](assign-a-policy-to-organizations.md) for more details.
