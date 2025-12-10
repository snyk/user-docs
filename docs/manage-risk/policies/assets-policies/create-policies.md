# Create policies

Snyk Essentials includes a powerful policy editor for creating and modifying policies.

There are two steps to building policies:

1. [Define filters](create-policies.md#define-filters) - Set filter conditions on asset properties.
2. [Set actions](create-policies.md#set-actions) - Define actions to be taken on filtered assets.

## New policy

You can create a new policy using the **Start from scratch** option or choose one of the available policy templates using the **Use a template** option.

<figure><img src="../../../.gitbook/assets/Policy - new UI.png" alt="Policy view, New policy button with the Start from scratch and Use a template options"><figcaption><p>Policy view, New policy options</p></figcaption></figure>

### Start from scratch - policy creation

To create a new policy, you have to click the **New Policy** option from the Policies/Assets view and select the **Start from scratch** option.

You must name your policy and, optionally, provide a description of the policy. After you complete these steps you have to [define the filters ](create-policies.md#define-filters)and [set the actions](create-policies.md#set-actions) of your policy.

### Use a template - policy creation

You can create a new policy by using one of the available templates. To select one of the policy templates, you have to click the **New Policy** option from the Policies/Assets view and select the **Use a template** option. You can select one of the templates from the templates library by clicking the **Use template** button from the policy template card.

Each policy template has a name, a description, and displays the graphic connections between filters and actions.

The following video explains how to use a policy template from the Policies view:

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737656948/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-5e_-_v1_-_Policy_Editor_Templates.mp4" %}
Overview of asset policy editor
{% endembed %}

You can customize the filters and actions or use the template as is. After finishing all the template changes, click the **Save** button to create the new policy.

<figure><img src="../../../.gitbook/assets/Policy template - new UI.png" alt="Policy templates accesible from Policy view, New policy button, the Use a template option"><figcaption><p>Policy templates accesible from Policy view, New policy button, the Use a template option</p></figcaption></figure>

## Define Filters

{% hint style="info" %}
**Release status**

The risk factors on assets are taking the release status of the applied [risk factor](/broken/pages/x5b9XFg7zGtut8eTWBQp).

The Runtime discovered and Runtime last seen filters are taking the release status of the used runtime [integration](broken-reference/).
{% endhint %}

The following video explains how to create a new policy:

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737656963/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-5b_-_v1_-_Policy_Editor.mp4" %}
Overview of asset policy editor
{% endembed %}

You must specify an asset property for each filter component. For all available properties for asset policies, see [Asset inventory filters](../../../manage-assets/assets-inventory-filters.md).

You can specify more than one filter component with an **And** or **Or** operator.

<figure><img src="../../../.gitbook/assets/Create policy New UI.png" alt="AppRisk - Create new policy"><figcaption><p>Asset policy - Create new policy</p></figcaption></figure>

The following video explains the use of filters and the use of the **And**, **Or** operator.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737656952/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-5c_-_v1_-_Policy_and_or_filters.mp4" %}
Using and/or operators in the asset policy editor
{% endembed %}

## Set actions

After defining filter components, you need to define the actions that the policy has to perform on the filtered assets. Asset policies support the following actions:

* **Send Email** - Receive an email every time there are asset updates. You can choose between daily emails or scheduling the checks. You can include a link to the relevant assets. Each notification lists all impacted assets. You can view them individually or see the aggregated view by clicking **Click Here**. The list of assets displayed in the email notification is automatically generated.
* **Send Slack Message** - Receive a Slack notification every time there are asset updates. You need to add your [Slack webhook URL](../../../integrations/jira-and-slack-integrations/slack-integration.md), then you can choose between daily emails or scheduling the checks. You can include a link to the relevant assets. Each notification lists all impacted assets. You can view them individually or see the aggregated view by clicking **Click Here**. The list of assets displayed in the email notification is automatically generated.
* **Set Asset Class** - Sets the class on the matched assets. Removing the policy or turning in off does not retroactively change the asset class back to default.
* **Set Asset Tag** - Sets a tag on the matched assets. Removing the policy or turning in off will remove the tags of this policy from the relevant assets.
* **Set Coverage Control Policy** - Sets a control on filtered assets that checks whether selected security products are scanning assets, optionally within a given timeframe. Assets that fail this control will be marked accordingly on inventory pages. This control applies the OR logic across products.

<figure><img src="../../../.gitbook/assets/Policy - Nwe UI.png" alt="AppRisk - Set a policy action"><figcaption><p>Asset policy - Set a policy action</p></figcaption></figure>

The editor supports multiple flows for the same policy. The flows can be independent or intersect.

<figure><img src="../../../.gitbook/assets/Multiple actions - New UI.png" alt="AppRisk - Set multiple policy actions"><figcaption><p>Asset Policy- Set multiple policy actions</p></figcaption></figure>
