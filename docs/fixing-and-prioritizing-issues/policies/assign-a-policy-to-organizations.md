# Assign a policy to organizations

A policy can be applied to multiple organizations; however, an organization can only be assigned to one policy. Organizations cannot be added to or removed from the default policy directly from the default policy modal.

Policies assigned to organizations will also apply when running snyk test or snyk monitor in the CLI.

## To add an organization to a policy:

Click on the associated check box from the organization selector panel.

If this organization is already assigned to another policy, you will be able to see which policy it is applied to from this selector and the pill next to the organization name will be grey

![mceclip3.png](https://support.snyk.io/hc/article_attachments/360007934458/mceclip3.png)

If an organization is assigned to the policy you are currently viewing, that policy name will appear in a yellow pill next to the organization name

![mceclip2.png](https://support.snyk.io/hc/article_attachments/360007850817/mceclip2.png)

If you are moving an organization from one policy to another, 2 pills will appear to show which policy is currently applied \(yellow\) and which policy you will be applying \(grey\)

## ![mceclip1.png](https://support.snyk.io/hc/article_attachments/360007934258/mceclip1.png)

## To remove an organization from a policy:

Uncheck the desired organization check box from the organization selector panel.

![Untitled.png](https://support.snyk.io/hc/article_attachments/360010813917/Untitled.png)

The unchecked organization\(s\) will now revert back to the default policy.

## Add an organization to the default policy:

1. Remove the organization from the policy it is currently assigned to
2. The organization will automatically move to the [default policy](https://support.snyk.io/hc/en-us/articles/360007476397)

## Remove an organization from the default policy

1. Assign the relevant organization to a new policy 
2. The organization will automatically disassociate from the [default policy](https://support.snyk.io/hc/en-us/articles/360007476397) and move to the newly selected policy

