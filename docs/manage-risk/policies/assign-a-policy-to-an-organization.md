# Assign a policy to an Organization

When you create a policy, you can apply it to one Organization. You cannot directly apply an Organization to or remove an Organization from the default policy using the Policy Manager.

{% hint style="info" %}
Policies applied to Organizations are in effect when you run the `snyk test` or `snyk monitor` CLI commands.
{% endhint %}

## Apply a policy to an Organization

To apply a policy to an Organization, in the Organization selector panel, check the box for the Organization to which you want to apply the policy.

If an Organization has another policy applied, you can see that policy from the selector, and the policy indicator next to the Organization name will be gray.

If the Organization already has the policy applied, the name of the policy is displayed in a yellow indicator next to the Organization name.

If you are applying a different policy to an Organization, in order to move that Organization from one policy to another, two indicators appear next to the Organization name in the selector. One shows, in yellow, the policy that is currently applied. The other shows the policy you will be applying to the Organization in gray.

## Remove a policy from an Organization

To remove a policy from an Organization, uncheck the box next to the Organization you want to remove from the policy that you are viewing.

The unchecked Organization will now automatically revert to the default policy.

## Apply the default policy to an Organization

Remove the policy currently applied to the Organization. The Organization will automatically revert to the default policy.

## Remove the default policy from an Organization

Apply a new policy to the Organization. The Organization will automatically be removed from the default policy and the new policy will be applied.
