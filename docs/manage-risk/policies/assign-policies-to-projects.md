# Assign policies to Projects

After you apply [Project attributes](../../snyk-platform-administration/snyk-projects/project-attributes.md) to your Projects, you can create policies that apply to those attributes. Projects and policies are linked based on the attributes that have the policy applied.

{% hint style="info" %}
Policies applied to Project attributes always take precedence over policies applied to Organizations.
{% endhint %}

A policy can be applied to one or multiple Project attributes, but only one policy can be applied to a set of attributes. For example, if there is a policy applied to `Critical`**,** `Production`,`Frontend`, you cannot create another policy that is applied _only_ to these exact attributes.

{% hint style="info" %}
Policies applied to Project attributes affect the CLI command `snyk monitor`, assuming it runs on a CLI Project that has Project attributes assigned. Project attributes applied to policies do not affect `snyk test`.
{% endhint %}

## Apply a policy to Project attributes and remove a policy

To apply a policy to an attribute, in the attribute selector panel, check the box for the attribute to which you want to apply the policy.

You can also search for tags that have already been created in Projects in your Group. You can select more than one tag for the policy.

To remove a policy from an attribute, uncheck the box next to the attribute from which you want to remove the policy.

To remove a tag, click the **x** next to the tag.

{% hint style="info" %}
You can create and save a policy where no attributes are selected, for example, if you have not yet decided the attributes to which the policy should be applied. A policy cannot apply to Projects if all attributes are left blank.
{% endhint %}

## Assign Projects to policies

To have a policy assigned, a Project must have all the attributes listed on the policy applied to the Project. The Project can also have attributes that are not listed on the policy.

{% hint style="info" %}
If a policy applies to a Project based on the attributes, then role with edit project attribute permission can edit the Project attributes.
{% endhint %}

If multiple tags are added to a policy, the Project needs to match with only one of the Project tags. However, if other attributes are also listed on the policy, the Project would need to have all the attributes and at least one of the listed tags.

For example, if you have a policy applied to `Critical`, `External`, and `Frontend`, this policy is assigned to Projects that have the same attributes, but not to a Project with the attributes `Critical` and `External` only.

An example policy follows. It is applied to an attribute in the **Business Criticality** section, `Critical`, and to attributes in the **Environment** section, `Frontend` and `External`. The policy also has two Project tags. The first tag has the key `PCI`, with the value of `Compliant`. The second tag has the key `owner`, with the value of `fred`.

The following Project has the attributes `Frontend`, `External`, and `Critical`, and has at least one matching tag, `PCI:Compliant`. Thus the Project will inherit the policy, that is, the policy is assigned to this Project.

The following Project will not inherit the policy, because the Project lacks the `External` environment attribute.

## Assign multiple policies to a Project

Multiple policies can be assigned to a Project. For example, you may have a policy applied to the attributes `Critical` and `External` and another policy applied to the attributes `Critical` and `Production`. If you have a Project with the attributes `Critical`, `External` and `Production`, both policies are assigned.

When multiple policies are assigned to a Project, the order of the policies on the policy manager page determines precedence. The policy closest to the top of the list takes precedence over other assigned policies after it. To change the order of policies, either drag and drop the policies into the order you want or use the three dots on the right-hand side to move the policy up or down in the list.
