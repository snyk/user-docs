# Assign a policy to project attributes

After applying [project attributes](https://docs.snyk.io/fixing-and-prioritizing-issues/policies/assign-a-policy-to-project-attributes) to your projects, you can create policies that apply to those attributes. Projects and policies are linked based on the attributes assigned to the policy.

{% hint style="info" %}
Policies assigned to attributes always take precedence over policies assigned to organizations.
{% endhint %}

A policy can be applied to one or multiple project attributes; but a set of attributes can only be assigned to one policy. For example, if there is already a policy applied to `Critical`**,** `Production`,`Frontend`, you cannot create another policy that matches _only_ these exact attributes.

Reminder: Policies assigned to project attributes apply when running **snyk monitor** in the CLI, assuming it runs on a CLI project with project attributes applied. Project attribute policies do not apply to **snyk test**.

## Add / remove an attribute to a policy

To add an attribute, click on the desired attribute checkbox\(es\) from the attribute selector panel.

To remove an attribute from a policy, uncheck the desired attribute checkbox\(es\) from the attribute selector panel.

![](../../../.gitbook/assets/screenshot_2021-03-11_at_1.20.42_pm.png)

{% hint style="info" %}
You can create and save a policy where no attributes are selected, for example, if you have not yet decided which attributes should be associated with that policy. This policy does not apply to projects where all attributes are left blank.
{% endhint %}

## Matching projects and policies

To be associated with a policy, a project must have all the attributes listed on the policy \(the project could also have more attributes that are not listed on the policy\).

For example, if you have a policy assigned to `Critical`, `External`, and `Frontend`, this policy applies to projects which includes those same attributes, but not to a project with the attributes `Critical` and `External`.  
Here is our sample policy:

![](../../../.gitbook/assets/screenshot_2021-03-11_at_11.54.33_am.png)

Here is a project that will inherit the policy:

![](../../../.gitbook/assets/screenshot_2021-03-11_at_12.26.02_pm.png)

Here is a project that will not inherit the policy:

![](../../../.gitbook/assets/screenshot_2021-03-11_at_12.29.03_pm.png)

## Applying multiple policies to a project

It is possible that more than one policy can be apply for a project. For example, if you have a policy assigned to `Critical` and `External` and another policy assigned to `Critical` and `Production`. If you have a project that has the attributes `Critical`, `External` and `Production`, it could apply to either of these policies!

If more than one policy can be associated with a project, the order of the policies on the policy manager page determines precedence. The policy closest to the top of the list takes precedence over other applicable policies below it. To change the order of policies, either drag and drop the policies into the right order, or use the **...** button on the right hand side to move the policy up or down in the list.

![](../../../.gitbook/assets/screenshot_2021-03-11_at_12.51.25_pm.png)

