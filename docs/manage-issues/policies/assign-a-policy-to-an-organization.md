# Apply a policy to Organizations

You can apply a policy to multiple Organizations. However, you can assign an Organization to only one policy. You cannot add an Organization to or remove an Organization from the default policy directly from the default policy dialog.

{% hint style="info" %}
Policies assigned to Organizations apply when your run the `snyk test` or `snyk monitor` CLI commands.
{% endhint %}

## Apply an Organization to a policy

To apply an Organization to a policy, in the Organization selector panel, check the box for the Organization you want to add.

If this Organization is already applied to another policy, you can see that policy from the selector, and the indicator next to the Organization name will be gray.

<div align="left">

<figure><img src="../../.gitbook/assets/mceclip3-2-.png" alt="Gray indicator - Organization already applied to another policy"><figcaption><p>Gray indicator - Organization already applied to another policy</p></figcaption></figure>

</div>

If an Organization is applied to the policy you are currently viewing, the name of the policy is displayed in a yellow indicator next to the Organization name.

<div align="left">

<figure><img src="../../.gitbook/assets/mceclip2-6-.png" alt="Yellow indicator - Organization already assigned to this policy"><figcaption><p>Yellow indicator - Organization already assigned to this policy</p></figcaption></figure>

</div>

If you are moving an Organization from one policy to another, two indicators appear to show where the Organization is currently applied, yellow, and the policy where you will be applying the Organization, gray.

<div align="left">

<figure><img src="../../.gitbook/assets/mceclip1-16-.png" alt="Gray and Yellow indicators - Policies where the Organization is applied and will be applied"><figcaption><p>Gray and Yellow indicators - Policies where the Organization is applied and will be applied</p></figcaption></figure>

</div>

## Remove an Organization from a policy

To remove an Organization from a policy, uncheck the box next to the Organization you want to remove.

<div align="left">

<figure><img src="../../.gitbook/assets/untitled-2-.png" alt="Remove an Organization from a policy"><figcaption><p>Remove an Organization from a policy</p></figcaption></figure>

</div>

The unchecked Organization will now revert to the default policy.

## Add an Organization to the default policy

Remove the Organization from the policy it is currently applied to. The Organization will automatically move to the default policy.

## Remove an Organization from the default policy

Apply the Organization to a new policy. The Organization will automatically be removed from the default policy and move to the newly selected policy.
