# Apply a policy to Projects

After applying [Project attributes](../introduction-to-snyk-projects/project-attributes.md) to your Projects, you can create policies that apply to those attributes. Projects and policies are linked based on the attributes assigned to the policy.

{% hint style="info" %}
Policies assigned to Project attributes always take precedence over policies assigned to Organizations.
{% endhint %}

A policy can be applied to one or multiple Project attributes; but a set of attributes can only be assigned to one policy. For example, if there is already a policy applied to `Critical`**,** `Production`,`Frontend`, you cannot create another policy that matches _only_ these exact attributes.

{% hint style="info" %}
Policies assigned to project attributes apply when running **snyk monitor** in the CLI, assuming it runs on a CLI project with project attributes applied. Project attribute policies do not apply to **snyk test**
{% endhint %}

## Add / remove an attribute to a policy

To add an attribute, click on the desired attribute checkbox(es) from the attribute selector panel.

To remove an attribute from a policy, uncheck the desired attribute checkbox(es) from the attribute selector panel.

<div align="left">

<figure><img src="../../.gitbook/assets/screenshot_2021-03-11_at_1.20.42_pm.png" alt="Add a policy attribute"><figcaption><p>Add a policy attribute</p></figcaption></figure>

</div>

{% hint style="info" %}
You can create and save a policy where no attributes are selected, for example, if you have not yet decided which attributes should be associated with that policy. This policy does not apply to projects where all attributes are left blank.
{% endhint %}

## Matching projects and policies

To be associated with a policy, a Project must have all the attributes listed on the policy (the Pincluderoject could also have more attributes that are not listed on the policy).

For example, if you have a policy assigned to `Critical`, `External`, and `Frontend`, this policy applies to projects which include those same attributes, but not to a project with the attributes `Critical` and `External`.

Here is our sample policy:

<div align="left">

<figure><img src="../../.gitbook/assets/screenshot_2021-03-11_at_11.54.33_am.png" alt="Example policy"><figcaption><p>Example policy</p></figcaption></figure>

</div>

Here is a Project that will inherit the policy:

<div align="left">

<figure><img src="../../.gitbook/assets/screenshot_2021-03-11_at_12.26.02_pm.png" alt="Project inheriting a policy"><figcaption><p>Project inheriting a policy</p></figcaption></figure>

</div>

Here is a project that will not inherit the policy:

<div align="left">

<figure><img src="../../.gitbook/assets/screenshot_2021-03-11_at_12.29.03_pm.png" alt="Project not inheriting a policy"><figcaption><p>Project not inheriting a policy</p></figcaption></figure>

</div>

## Applying multiple policies to a Project

Multiple policies can apply to a Project. For example, if you have a policy assigned to `Critical` and `External` and another policy assigned to `Critical` and `Production`. If you then have a Project with attributes `Critical`, `External` and `Production`, it could have either of these policies applying.

If multiple policies can be applied to a project, the order of the policies on the policy manager page determines precedence. The policy closest to the top of the list takes precedence over other applicable policies below it. To change the order of policies, either drag and drop the policies into the right order, or use the **...** button on the right hand side to move the policy up or down in the list.

<div align="left">

<figure><img src="../../.gitbook/assets/screenshot_2021-03-11_at_12.51.25_pm.png" alt="Change policy order"><figcaption><p>Change policy order</p></figcaption></figure>

</div>
