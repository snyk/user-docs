# View, create, and modify policies

## View policies

{% hint style="info" %}
You must be a Group administrator to view, create, and modify policies for that Group.
{% endhint %}

Select the **Policies** menu option to see the policies in your Group, arranged by category, [License policies](license-policies/), and [Security policies](security-policies/).

Expand a category to see a list of the policies in that category.

{% hint style="info" %}
This list includes the [default policy](view-create-and-modify-policies.md#default-policies), which is automatically created for new Groups for each policy category and cannot be removed.
{% endhint %}

### Policy details

When you expand a category, the screen shows the policies applied to **Project attributes** and applied to **Organizations**. You can click to **Learn which policies take precedence** in each category. You can also search for a particular policy.

### Default policies

Each policy category has a default policy. Default policies can be applied only to Organizations, not Project attributes.

When you create a new Organization, it will automatically be added to the default policy unless you have copied the settings of an existing Organization. You can assign an Organization to a different policy if desired.

The default policy cannot be deleted; however, the default policy name, description, and rules can be edited to match your preferences. A default policy can also contain no rules if you'd prefer.

See [Assign a policy to an Organization](assign-a-policy-to-an-organization.md) for details.

The Policy Manager allows you to [create](view-create-and-modify-policies.md#create-a-policy), [edit](view-create-and-modify-policies.md#edit-a-policy), and [duplicate or delete a policy](view-create-and-modify-policies.md#duplicate-or-delete-a-policy).

## Create a policy

1. On the Policy Manager screen, select **Add new policy** and in response to the prompts, enter the details.
2. Enter a policy name and a description to help you quickly identify a policy.\
   Policies in the same category cannot have the same name. You cannot save a policy without a name.
3. Select whether to apply the policy to Organizations or Project attributes.
4. Select the [Organizations](assign-a-policy-to-an-organization.md) or [attributes](assign-policies-to-projects.md) to which you want to apply the policy.
5. Add rules to the policy. See [Create a license policy and rules](license-policies/create-a-license-policy-and-rules.md) or [Create a security policy and rules](security-policies/create-a-security-policy-and-rules.md).
6. Click **Submit** to create and save the policy.

## Edit a policy

1. Click the name of an existing policy in the Policy Manager tab to make any changes.
2. Change the [Organizations](assign-a-policy-to-an-organization.md), [attributes](assign-policies-to-projects.md), and rules as you wish.
3. Click **Submit** to save your changes.

## Duplicate or delete a policy

To duplicate or delete a policy, click the three dots on the right-hand side:

Duplicating a policy copies the rules of a policy but not the assigned Organizations or Projects. The new policy is automatically called **Copy of (Policy Name)…** and can be edited as [explained in Edit a policy](view-create-and-modify-policies.md#edit-a-policy).

{% hint style="info" %}
Deleting a policy cannot be undone. If you delete a policy with Organizations assigned to it, those Organizations will return to the default policy.
{% endhint %}
