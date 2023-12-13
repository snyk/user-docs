# View policies

{% hint style="warning" %}
You must be a Group administrator to view, create, and modify policies for that Group.
{% endhint %}

Select the **Policies** menu option to see the policies in your Group, arranged by category, [License policies](license-policies/), and [Security policies](security-policies/).

<div align="left">

<figure><img src="../../.gitbook/assets/Policies-menu.png" alt="View policies"><figcaption><p>View policies</p></figcaption></figure>

</div>

Expand a category to see a list of the policies in that category:

<figure><img src="../../.gitbook/assets/snyk-policy-manager.png" alt="License policies list expanded"><figcaption><p>License policies list expanded</p></figcaption></figure>

{% hint style="info" %}
This list includes the [default policy](view-policies.md#default-policies), which is automatically created for new Groups for each policy category and cannot be removed.
{% endhint %}

## Policy details

When you expand a category, the screen shows the policies applied to **Project attributes** and applied to **Organizations**. You can click to **Learn which policies take precedence** in each category. You can also **search** for a particular policy.

<figure><img src="../../.gitbook/assets/screenshot_2021-03-26_at_11.04.50_am.png" alt="Policy manager screen including attributes and Organizations to which each policy is applied"><figcaption><p>Policy manager screen including attributes and Organizations to which each policy is applied</p></figcaption></figure>

## Default policies

Each policy category has a default policy. Default policies can be applied only to Organizations, not Project attributes.

When you create a new Organization, it will automatically be added to the default policy unless you have copied the settings of an existing Organization. You can assign an Organization to a different policy if desired.

The default policy cannot be deleted; however, the default policy name, description, and rules can be edited to match your preferences. A default policy can also contain no rules if you'd prefer.

See [Assign a policy to an Organization](apply-a-policy-to-organizations.md) for details.
