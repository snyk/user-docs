# Create a security policy and rules

To create a new security policy, navigate to **Policies** in your Group menu, and in the Policies manager, expand the **Security policies** category and click **Add new policy**. For details, see [View policies](../view-create-and-modify-policies.md).

{% hint style="info" %}
Select **Snyk Default Security Policy** to change the conditions or actions for a security policy that applies to all Projects in all Organizations in the Group.

Security policies are applicable to Snyk Open Source and Snyk Container Projects.
{% endhint %}

## Rules, conditions, and actions

Security policy rules are in the “if, then” format with one or more conditions and an action. An example follows:

{% hint style="info" %}
When you create a new security policy, the first blank rule is created automatically.
{% endhint %}

Select the conditions and actions to complete a rule. See [Security policy conditions](security-policies-conditions.md) and [Security policy actions](security-policy-actions.md) for details.

* To add a new blank rule, click on the plus sign beneath the previous rule.
* To delete or duplicate a rule, click the three dots at the top right of the rule box.

{% hint style="info" %}
The order of the rules you create determines the order they will be applied. If multiple rules with the same action type match the same issue, the rule closest to the top takes precedence over any subsequent rules.

**Example:**

* Rule 1: Sets severity to HIGH for issues with mature exploit maturity.
* Rule 2: Sets severity to LOW for issue ID "CVE-2025-0001".

If an issue with "CVE-2025-0001" also has mature exploit maturity, both rules apply. Because Rule 1 is listed first, the final severity will be HIGH.
{% endhint %}
