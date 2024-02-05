# Level 2: Policies

The [Snyk AppRisk policies](../policies-for-snyk-apprisk/) help you automate the process of adding business context and receiving notifications. You can set up policies to automatically identify the coverage control gaps. &#x20;

## Understand a policy

You need two components to create and understand a policy:

* &#x20;**Filters**: helps you create asset groups based on defined criteria such as tags or asset name.
* **Actions**: set actions that can be performed on the filtered assets, such as setting asset classification or sending a Slack message or email.

{% hint style="info" %}
A newly created policy may take up to 30 minutes to become visible and be applied.
{% endhint %}

## Create a policy

You can create a policy by navigating to the Policy view and using the New policy button in the top right corner. Name the policy and, optionally, describe it.&#x20;

The policy builder editor focuses on two main areas:

* [Define the filters](../policies-for-snyk-apprisk/create-policies.md#define-filters) - Set filter conditions on asset properties.
* [Set actions](../policies-for-snyk-apprisk/create-policies.md#set-actions) - Define actions to be taken on filtered assets.

### Key filter types <a href="#key-filter-types" id="key-filter-types"></a>

When creating a new policy, you can choose from several available filters. Here is a selection of the most used ones:

* **Name:** Many companies have a naming convention for their repositories and other assets. Use the naming convention to check if asset names contain or match a certain string.
* **Asset Type:** you can apply policies only to certain asset types, such as repositories or software packages.
* **Class:** you can reduce a lot of noise on actions by only targeting a specific Class. For example, only taking actions on Class A assets, as these are business-critical assets.
* **Tags:** You can use either out-of-the-box system tags or custom tags that Snyk has defined.

Here are the steps for creating a policy:

1. Navigate to the Policy view and click the **New policy** button in the top right corner.
2. Provide a name and, optionally, a description of your policy, and click **Next**.
3. Define the filters of your policy and click **Apply**.
4.  Set actions for your policy. Click the + icon to add an action. You can trigger an action based on the assets you have filtered or create a logic node to combine it with another filter node.&#x20;

    <figure><img src="../../../.gitbook/assets/image (2).png" alt="Create policy - multiple nodes"><figcaption><p>Create policy - Multiple nodes</p></figcaption></figure>
5. When you add another filter node and select the + icon next to it, you can flow into existing logic nodes.
6.  Trigger an action based on the assets you have filtered.&#x20;

    <figure><img src="../../../.gitbook/assets/image (3).png" alt=""><figcaption><p>Create policy - Set action</p></figcaption></figure>

### Policy freshness

All policies are automatically run in a maximum of 30 minutes after creation, then every 30 minutes. You can manually run a policy by clicking Run to apply the policy to your assets. Changes are applied automatically to your assets by implementing the actions you set on the policy.

### Use cases for policies

Familiarize yourself with the Snyk AppRisk policies by going through these use cases:

* [Coverage control](../policies-for-snyk-apprisk/use-cases-for-policies/coverage-control-policy-use-case.md) policy - identify and set coverage policies to allow your team to define where certain security controls need to be in place.
* [Classification](../policies-for-snyk-apprisk/use-cases-for-policies/classification-policy-use-case.md) policy - classify assets based on importance.
* [Tagging](../policies-for-snyk-apprisk/use-cases-for-policies/tagging-policy-use-case.md) policy - sets a tag on the matched assets.
* [Notification](../policies-for-snyk-apprisk/use-cases-for-policies/notification-policy-use-case.md) policy - get notifications about changes that take place on your assets.
* [Coverage and coverage gap](../policies-for-snyk-apprisk/use-cases-for-policies/coverage-and-coverage-gap-policies.md) policies - use the Coverage filter to verify if an asset has ever been tested by the product and the Coverage gap filter to verify if the asset meets the coverage requirements set in a Set coverage control policy.
