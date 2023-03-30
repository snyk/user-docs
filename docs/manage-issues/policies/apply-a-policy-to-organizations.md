# Apply a policy to Organizations

A policy can be applied to multiple Organizations; however, an Organization can only be assigned to one policy. Organizations cannot be added to or removed from the default policy directly from the default policy modal.

{% hint style="info" %}
Policies assigned to Organizations also apply when running **snyk test** or **snyk monitor** CLI commands.
{% endhint %}

## To add an Organization to a policy

Click on the associated check box from the Organization selector panel.

If this Organization is already assigned to another policy, you will be able to see which policy it is applied to from this selector and the pill next to the organization name will be grey

<figure><img src="../../.gitbook/assets/mceclip3-2-.png" alt="Grey pill - Organization already assigned to another policy"><figcaption><p>Grey pill - Organization already assigned to another policy</p></figcaption></figure>

If an Organization is assigned to the policy you are currently viewing, that policy name will appear in a yellow pill next to the Organization name

<figure><img src="../../.gitbook/assets/mceclip2-6-.png" alt="Yellow pill - Organization already assigned to this policy"><figcaption><p>Yellow pill - Organization already assigned to this policy</p></figcaption></figure>

If you are moving an Organization from one policy to another, 2 pills will appear to show which policy is currently applied (yellow) and which policy you will be applying (grey)

<figure><img src="../../.gitbook/assets/mceclip1-16-.png" alt="Grey and Yellow pill - Organization already assigned and will move policy"><figcaption><p>Grey and Yellow pill - Organization already assigned and will move policy</p></figcaption></figure>

## To remove an Organization from a policy

Uncheck the desired Organization check box from the organization selector panel.

![Remove an Organization from a policy](../../.gitbook/assets/untitled-2-.png)

The unchecked Organization will now revert back to the default policy.

## Add an Organization to the default policy

1. Remove the Organization from the policy it is currently assigned to
2. The Organization will automatically move to the default policy

## Remove an Organization from the default policy

1. Assign the relevant Organization to a new policy
2. The Organization will automatically disassociate from the default policy and move to the newly selected policy
