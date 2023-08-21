# Apply a policy to Projects

After you apply [Project attributes](../introduction-to-snyk-projects/project-attributes.md) to your Projects, you can create policies that apply to those attributes. Projects and policies are linked based on the attributes assigned to the policy.

{% hint style="info" %}
Policies assigned to Project attributes always take precedence over policies assigned to Organizations.
{% endhint %}

A policy can be applied to one or multiple Project attributes, but a set of attributes can only be assigned to one policy. For example, if there is already a policy applied to `Critical`**,** `Production`,`Frontend`, you cannot create another policy that matches _only_ these exact attributes.

{% hint style="info" %}
Policies assigned to Project attributes apply when you run `snyk monito`**r** in the CLI, assuming it runs on a CLI Project with Project attributes applied. Project attribute policies do not apply to `snyk test`.
{% endhint %}

## Apply a policy to Project attributes and remove attributes

To apply a policy to an attribute, in the attribute selector panel, check the box for the attribute you want to add to the policy,

To remove an attribute from a policy, uncheck the box next to the attribute you want to remove.

<div align="left">

<figure><img src="../../.gitbook/assets/screenshot_2021-03-11_at_1.20.42_pm.png" alt="Add attributes to policies"><figcaption><p>Add attributes to policies</p></figcaption></figure>

</div>

{% hint style="info" %}
You can create and save a policy where no attributes are selected, for example, if you have not yet decided which attributes should be associated with a policy. This policy does not apply to Projects where all attributes are left blank.
{% endhint %}

## Associating Projects with policies

To be associated with a policy, a Project must have all the attributes listed on the policy. The Project could also have more attributes that are not listed on the policy.

For example, if you have a policy applied to `Critical`, `External`, and `Frontend`, this policy applies to Projects which include those same attributes, but not to a Project with the attributes `Critical` and `External` only.

An example policy follows. It is applied to `Critical` **Business Criticality** and to `Frontend` and `External` **Environment**.

<div align="left">

<figure><img src="../../.gitbook/assets/screenshot_2021-03-11_at_11.54.33_am.png" alt="Example policy"><figcaption><p>Example policy</p></figcaption></figure>

</div>

The following Project will inherit the policy. The Project has the attributes `Frontend`, `External`, and `Critical`.

<div align="left">

<figure><img src="../../.gitbook/assets/screenshot_2021-03-11_at_12.26.02_pm.png" alt="Project inheriting a policy"><figcaption><p>Project inheriting a policy</p></figcaption></figure>

</div>

The following Project will not inherit the policy, because the Project lacks the `Frontend` environment attribute.

<div align="left">

<figure><img src="../../.gitbook/assets/screenshot_2021-03-11_at_12.29.03_pm.png" alt="Project not inheriting a policy"><figcaption><p>Project not inheriting a policy</p></figcaption></figure>

</div>

## Applying multiple policies to a Project

Multiple policies can apply to a Project. For example, you may have a policy applied to the attributes `Critical` and `External` and another policy applied to the attributes `Critical` and `Production`. If you have a Project with the attributes `Critical`, `External` and `Production`, either of these policies apply.

When multiple policies are applied to a Project, the order of the policies on the policy manager page determines precedence. The policy closest to the top of the list takes precedence over other applicable policies below it. To change the order of policies, either drag and drop the policies into the correct order or use the three dots on the right-hand side to move the policy up or down in the list.

<div align="left">

<figure><img src="../../.gitbook/assets/screenshot_2021-03-11_at_12.51.25_pm.png" alt="Change policy order"><figcaption><p>Change policy order</p></figcaption></figure>

</div>
