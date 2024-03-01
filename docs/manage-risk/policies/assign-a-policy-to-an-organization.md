# Assign a policy to an Organization

When you create a policy, you can apply it to one Organization. You cannot directly apply an Organization to or remove an Organization from the default policy using the Policy Manager.

{% hint style="info" %}
Policies applied to Organizations are in effect when you run the `snyk test` or `snyk monitor` CLI commands.
{% endhint %}

## Apply a policy to an Organization

To apply a policy to an Organization, in the Organization selector panel, check the box for the Organization to which you want to apply the policy.

If an Organization has another policy applied, you can see that policy from the selector, and the policy indicator next to the Organization name will be gray.

<div align="left">

<figure><img src="../../.gitbook/assets/mceclip3-2-.png" alt=".Gray indicator - Organization has another policy applied"><figcaption><p>Gray indicator - Organization has another policy applied</p></figcaption></figure>

</div>

If the Organization already has the policy applied, the name of the policy is displayed in a yellow indicator next to the Organization name.

<div align="left">

<figure><img src="../../.gitbook/assets/mceclip2-6-.png" alt="Yellow indicator - Organization already assigned to this policy"><figcaption><p>Yellow indicator - Organization already assigned to this policy</p></figcaption></figure>

</div>

If you are applying a different policy to an Organization, in order to move that Organization from one policy to another, two indicators appear next to the Organization name in the selector. One shows, in yellow, the policy that is currently applied. The other shows the policy you will be applying to the Organization in gray.

<div align="left">

<figure><img src="../../.gitbook/assets/mceclip1-16-.png" alt="Gray and Yellow indicators - Policies applied to the Organization and to be applied"><figcaption><p>Gray and Yellow indicators - Policies applied to the Organizaiton and to be applied</p></figcaption></figure>

</div>

## Remove a policy from an Organization

To remove a policy from an Organization, uncheck the box next to the Organization you want to remove from the policy that you are viewing.

<div align="left">

<figure><img src="../../.gitbook/assets/untitled-2-.png" alt="Remove an Organization from a policy"><figcaption><p>Remove an Organization from a policy</p></figcaption></figure>

</div>

The unchecked Organization will now automatically revert to the default policy.

## Apply the default policy to an Organization

Remove the policy currently applied to the Organization. The Organization will automatically revert to the default policy.

## Remove the default policy from an Organization

Apply a new policy to the Organization. The Organization will automatically be removed from the default policy and the new policy will be applied.
