# Create a security policy and rules

To create a new security policy, navigate to **Policies** in your Group menu, and in the Policies manager, expand the **Security policies** category and click **Add new policy**. For details, see [View policies](../view-policies.md).

<figure><img src="../../../.gitbook/assets/screenshot_2020-10-20_at_10.01.49_am.png" alt="Security policies category expanded"><figcaption><p>Security policies category expanded</p></figcaption></figure>

{% hint style="info" %}
Select **Snyk Default Security Policy** to change the conditions or actions for a security policy that applies to all Projects in all Organizations in the Group.
{% endhint %}

## Rules, conditions, and actions

Security policy rules are in the “if, then” format with one or more conditions and an action. An example follows:

<div align="left">

<figure><img src="../../../.gitbook/assets/screenshot_2020-07-06_at_11.38.07.png" alt="Security policy rule format with plus sign to add a rule and three dots top right"><figcaption><p>Security policy rule format with plus sign to add a rule and three dots top right</p></figcaption></figure>

</div>

{% hint style="info" %}
When you create a new security policy, the first blank rule is created automatically.
{% endhint %}

Select the conditions and actions to complete a rule. See [Security policy conditions](security-policies-conditions.md) and [Security policy actions](security-policy-actions.md) for details.

* To add a new blank rule, click on the plus sign beneath the previous rule.
* To delete or duplicate a rule, click the three dots at the top right of the rule box.

{% hint style="info" %}
The order of the rules you create sets the precedence. If there is a conflict, the rule closest to the top takes precedence over any subsequent rules.
{% endhint %}
