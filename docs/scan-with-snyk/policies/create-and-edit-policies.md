# Create and modify policies

The Policy Manager allows you to [create](create-and-edit-policies.md#create-a-policy), [edit](create-and-edit-policies.md#edit-a-policy), and [duplicate or delete a policy](create-and-edit-policies.md#duplicate-or-delete-a-policy).

## **Create a policy**

1. On the Policy Manager screen, select **Add new policy** and in response to the prompts, enter the details.
2. Enter a policy name and a description to help you quickly identify a policy.\
   Policies in the same category cannot have the same name. You cannot save a policy without a name.
3. Select whether to apply the policy to Organizations or Project attributes.
4. Select the [Organizations](assign-a-policy-to-an-organization.md) or [attributes](assign-policies-to-projects.md) to which you want to apply the policy.
5. Add rules to the policy. See [Create a license policy and rules](license-policies/create-a-license-policy-and-rules.md) or [Create a security policy and rules](security-policies/create-a-security-policy-and-rules.md).
6. Click **Submit** to create and save the policy.

<div align="left">

<figure><img src="../../.gitbook/assets/screenshot_2020-05-26_at_9.47.26_am.png" alt="Create a policy" width="563"><figcaption><p>Create a policy</p></figcaption></figure>

</div>

## Edit a policy

1. Click the name of an existing policy in the Policy Manager tab to make any changes.
2. Change the [Organizations](assign-a-policy-to-an-organization.md), [attributes](assign-policies-to-projects.md), and rules as you wish.
3. Click **Submit** to save your changes.

## **Duplicate or delete a policy**&#x20;

Click the three dots on the right-hand side to duplicate or delete a policy:

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2023-03-28 at 16.42.45.png" alt="Other policy actions"><figcaption><p>Other policy actions</p></figcaption></figure>

</div>

### Delete a policy

{% hint style="warning" %}
Deleting a policy cannot be undone. If you delete a policy that has Organizations assigned to it, those Organizations will return to the default policy.
{% endhint %}

### Duplicate a policy

Duplicating a policy copies the rules of a policy but not the assigned Organizations or Projects. The new policy is automatically called **Copy of (Policy Name)…** and can be edited as [explained in Edit a policy](create-and-edit-policies.md#edit-a-policy).
