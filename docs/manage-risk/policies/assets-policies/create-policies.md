# Create policies

Snyk Essentials or Snyk AppRisk includes a powerful policy editor for creating and modifying policies.

There are two steps to building policies:

1. [Define filters](create-policies.md#define-filters) - Set filter conditions on asset properties.
2. [Set actions](create-policies.md#set-actions) - Define actions to be taken on filtered assets.

## New policy

You can create a new policy using the **Start from scratch** option or choose one of the available policy templates using the **Use a template** option.

<figure><img src="../../../.gitbook/assets/Policy - new UI.png" alt="Policy view, New policy button with the Start from scratch and Use a template options "><figcaption><p>Policy view, New policy options </p></figcaption></figure>

### Start from scratch - policy creation

To create a new policy, you have to click the **New Policy** option from the Policies/Assets view and select the **Start from scratch** option.

You must name your policy and, optionally, provide a description of the policy. After you complete these steps you have to [define the filters ](create-policies.md#define-filters)and [set the actions](create-policies.md#set-actions) of your policy.&#x20;

### Use a template - policy creation

You can create a new policy by using one of the available templates. To select one of the policy templates, you have to click the **New Policy** option from the Policies/Assets view and select the **Use a template** option. You can select one of the templates from the templates library by clicking the **Use template** button from the policy template card.&#x20;

Each policy template has a name, a description, and displays the graphic connections between filters and actions.

The following video explains how to use a policy template from the Policies view:

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737656948/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-5e_-_v1_-_Policy_Editor_Templates.mp4" %}
Overview of asset policy editor
{% endembed %}

You can customize the filters and actions or use the template as is. After finishing all the template changes, click the **Save** button to create the new policy.

<figure><img src="../../../.gitbook/assets/Policy template - new UI.png" alt="Policy templates accesible from Policy view, New policy button, the Use a template option "><figcaption><p>Policy templates accesible from Policy view, New policy button, the Use a template option </p></figcaption></figure>

## **Define Filters**

{% hint style="info" %}
**Release status**

The risk factors on assets are taking the release status of the applied [risk factor](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md).

The Runtime discovered and Runtime last seen filters are taking the release status of the used runtime [integration](../../snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md).
{% endhint %}

Each filter component requires you to specify an asset property. Navigate to the [Filters capabilities](../../../manage-assets/assets-inventory-features.md#filters-capabilities) page to view all available properties for asset policies.

The following video explains how to create a new policy:

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737656963/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-5b_-_v1_-_Policy_Editor.mp4" %}
Overview of asset policy editor
{% endembed %}

Each property contains different options for conditions and values:

| Property                     | Conditions Values                                                                                                                         | Values                                                                                                                |
| ---------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| Application\*                | <ul><li>Is one of</li><li>Is not one of</li></ul>                                                                                         | all available applications for which you have configured the application context in Snyk Essentials or Snyk AppRisk.  |
| Asset ID                     | <ul><li>is</li><li>is not</li><li>contains</li><li>does not contain</li><li>starts with</li><li>ends with</li></ul>                       | \[string]                                                                                                             |
| Asset name                   | <ul><li>is</li><li>is not</li><li>contains</li><li>does not contain</li><li>starts with</li><li>ends with</li></ul>                       | \[string]                                                                                                             |
| Asset type                   | <ul><li>Is one of</li><li>Is not one of</li></ul>                                                                                         | <ul><li>Package</li><li>Repository</li><li>Scanned artifact</li></ul>                                                 |
| Attribute                    | <ul><li>is</li><li>is not</li><li>contains</li><li>does not contain</li><li>starts with</li><li>ends with</li></ul>                       | \[string]                                                                                                             |
| Catalog name\*               | <ul><li>Is one of</li><li>Is not one of</li></ul>                                                                                         | the list of names of your application context.                                                                        |
| Category                     | <ul><li>Is one of</li><li>Is not one of</li></ul>                                                                                         | the list of the available categories of a repository asset                                                            |
| Class                        | <ul><li>Is one of</li><li>Is not one of</li></ul>                                                                                         | A, B, C, D                                                                                                            |
| Coverage                     | <ul><li>containing one or more of</li><li>containing all of</li><li>not containing one or more of</li><li>not containing all of</li></ul> | Snyk Code, Container, IaC, Open Source                                                                                |
| Coverage gap                 | <ul><li>containing one or more of</li><li>containing all of</li><li>not containing one or more of</li><li>not containing all of</li></ul> | Snyk Code, Container, IaC, Open Source                                                                                |
| Developers                   | <ul><li>is</li><li>is not</li><li>contains</li><li>does not contain</li><li>starts with</li><li>ends with</li></ul>                       | \[string]                                                                                                             |
| Discovered                   | <ul><li>Is within</li><li>Is not within</li></ul>                                                                                         | <ul><li>Last 24 hours</li><li>Last 7 days</li><li>Last 30 days</li><li>Last 12 months</li><li>Year to date </li></ul> |
| <p></p><p>Issue severity</p> | <ul><li>containing one or more of</li><li>containing all of</li><li>not containing one or more of</li><li>not containing all of</li></ul> | <ul><li>Critical </li><li>High </li><li>Medium </li><li>Low</li></ul>                                                 |
| Issue source                 | <ul><li>containing one or more of</li><li>containing all of</li><li>not containing one or more of</li><li>not containing all of</li></ul> | Snyk Code, Container, IaC, Open Source, Nightfall                                                                     |
| Last seen                    | <ul><li>Is within</li><li>Is not within</li></ul>                                                                                         | <ul><li>Last 24 hours</li><li>Last 7 days</li><li>Last 30 days</li><li>Last 12 months</li><li>Year to date</li></ul>  |
| Lifecycle\*                  | <ul><li>Is one of</li><li>Is not one of</li></ul>                                                                                         | the available list of the lifecycle states of the application context component                                       |
| Locked attributes            | <ul><li>is one of</li><li>is not one of</li></ul>                                                                                         | <ul><li>Class</li></ul>                                                                                               |
| Owner\*                      | <ul><li>is one of</li><li>is not one of</li></ul>                                                                                         | the list of teams owning the repository for which the application context was configured.                             |
| Risk factors                 | <ul><li>containing one or more of</li><li>containing all of</li><li>not containing one or more of</li><li>not containing all of</li></ul> | the list of available risk factors                                                                                    |
| Runtime discovered           | <ul><li>Is within</li><li>Is not within</li></ul>                                                                                         | <ul><li>Last 24 hours</li><li>Last 7 days</li><li>Last 30 days</li><li>Last 12 months</li><li>Year to date</li></ul>  |
| Runtime last seen            | <ul><li>Is within</li><li>Is not within</li></ul>                                                                                         | <ul><li>Last 24 hours</li><li>Last 7 days</li><li>Last 30 days</li><li>Last 12 months</li><li>Year to date</li></ul>  |
| SCM Repository freshness     | <ul><li>is one of</li><li>is not one of</li></ul>                                                                                         | <ul><li>Active</li><li>Inactive</li><li>Dormant</li></ul>                                                             |
| Source                       | <ul><li>is one of</li><li>is not one of</li></ul>                                                                                         | <ul><li>azure-devops</li><li>GitHub</li><li>GitLab</li><li>Snyk</li></ul>                                             |
| Tags                         | <ul><li>containing one or more of</li><li>containing all of</li><li>not containing one or more of</li><li>not containing all of</li></ul> | all available tags you previously created                                                                             |
| Title\*                      | <ul><li>is one of</li><li>is not one of</li></ul>                                                                                         | the list with all the names of the component for which the application context was configured                         |

**\***&#x41;ll filters marked with `*` are visible only to the users who configured the application context for their SCM integrations.

You can specify more than one filter component with an **And** or **Or** operator.\


<figure><img src="../../../.gitbook/assets/Create policy New UI.png" alt="AppRisk - Create new policy"><figcaption><p>Snyk Essentials or Snyk AppRisk - Create new policy</p></figcaption></figure>

The following video explains the use of filters and the use of the **And**, **Or** operator.&#x20;

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737656952/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-5c_-_v1_-_Policy_and_or_filters.mp4" %}
Using and/or operators in the asset policy editor
{% endembed %}

## **Set actions**

After defining filter components, you need to define the actions that the policy has to perform on the filtered assets. Asset policies support the following actions:

* **Send Email** - Receive an email every time there are asset updates. You can choose between daily emails or scheduling the checks. You can include a link to the relevant assets. Each notification lists all impacted assets. You can view them individually or see the aggregated view by clicking **Click Here**. The list of assets displayed in the email notification is automatically generated.
* **Send Slack Message** - Receive a Slack notification every time there are asset updates. You need to add your [Slack webhook URL](../../../integrate-with-snyk/jira-and-slack-integrations/slack-integration.md), then you can choose between daily emails or scheduling the checks. You can include a link to the relevant assets. Each notification lists all impacted assets. You can view them individually or see the aggregated view by clicking **Click Here**. The list of assets displayed in the email notification is automatically generated.
* **Set Asset Class** - Sets the class on the matched assets. Removing the policy or turning in off does not retroactively change the asset class back to default.
* **Set Asset Tag** - Sets a tag on the matched assets. Removing the policy or turning in off will remove the tags of this policy from the relevant assets.
* **Set Coverage Control Policy** - Sets a control on filtered assets that checks whether selected security products are scanning assets, optionally within a given timeframe. Assets that fail this control will be marked accordingly on inventory pages. This control applies the OR logic across products.

<figure><img src="../../../.gitbook/assets/Policy - Nwe UI.png" alt="AppRisk - Set a policy action"><figcaption><p>Snyk Essentials or Snyk AppRisk - Set a policy action </p></figcaption></figure>

The editor supports multiple flows for the same policy. The flows can be independent or intersect.

<figure><img src="../../../.gitbook/assets/Multiple actions - New UI.png" alt="AppRisk - Set multiple policy actions "><figcaption><p>Snyk Essentials or Snyk AppRisk - Set multiple policy actions </p></figcaption></figure>
