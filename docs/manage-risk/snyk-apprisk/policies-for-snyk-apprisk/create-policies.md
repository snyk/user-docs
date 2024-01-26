# Create policies

Snyk AppRisk includes a powerful policy editor for creating and modifying policies.

There are two steps to building policies:

1. [Define filters](create-policies.md#define-filters) - Set filter conditions on asset properties.
2. [Set actions](create-policies.md#set-actions) - Define actions to be taken on filtered assets.

## **Define Filters**

Each filter component requires you to specify an asset property. Available properties for asset policies include:

* **Asset name** - the name of the repository asset or scanned artifact.
* **Asset type** - repository asset or scanned artifact.
* **Attribute** - asset attributes retrieved from the data source.
* **Class** - specify the class of the asset.
* **Coverage** - specify the product or products used to scan the asset. The Coverage filter identifies if at least one scan has been run by the specified product.
* **Coverage gap** - specify the products for which the asset has not been scanned and do not meet the Set Coverage Control Policy requirements.
* **Developers** - specify the developer or developers who contributed to the asset.
* **Discovered** - specify the period when the asset was discovered.
* **Last seen** - specify the repository freshness status.
* **Repository freshness** - the status of the repository and the date of the last commit.
  * **Active**: Had commits in the last 3 months.
  * **Inactive**: The last commits were made in the last 3 - 6 months.
  * **Dormant**: No commits in the last 6 months.
* **Source** - specify the asset source.
* **Tags** - information about the detected languages and repository update status.

{% hint style="info" %}
**Coverage** and **Coverage gap**  filter differences

* Use the Coverage filter to identify the assets scanned by the products at least once.
* Use the Coverage gap filter for assets that do not meet the requirements defined in the Set coverage control policy.&#x20;

The Coverage gap filter identifies assets that fall 'out of policy' and do not satisfy the coverage criteria you have specified, due to infrequent scanning or no scanning at all. On the other hand, the Coverage filter allows you to locate assets that have or have not been scanned, irrespective of any coverage requirements.
{% endhint %}



Each property contains different options for conditions and values:

| Property             | Conditions Values                                                                                                                         | Values                                                                                                                |
| -------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| Asset name           | <ul><li>is</li><li>is not</li><li>contains</li><li>does not contain</li><li>starts with</li><li>ends with</li></ul>                       | \[string]                                                                                                             |
| Asset type           | <ul><li>Is any of</li><li>Is not any of</li></ul>                                                                                         | <ul><li>Package</li><li>Repository</li></ul>                                                                          |
| Attribute            | <ul><li>is</li><li>is not</li><li>contains</li><li>does not contain</li><li>starts with</li><li>ends with</li></ul>                       | \[string]                                                                                                             |
| Class                | <ul><li>Is any of</li><li>Is not any of</li></ul>                                                                                         | A, B, C, D                                                                                                            |
| Coverage             | <ul><li>containing one or more of</li><li>containing all of</li><li>not containing one or more of</li><li>not containing all of</li></ul> | Snyk Code, Container, IaC, Open Source                                                                                |
| Coverage gap         | <ul><li>containing one or more of</li><li>containing all of</li><li>not containing one or more of</li><li>not containing all of</li></ul> | Snyk Code, Container, IaC, Open Source                                                                                |
| Developers           | <ul><li>is</li><li>is not</li><li>contains</li><li>does not contain</li><li>starts with</li><li>ends with</li></ul>                       | \[string]                                                                                                             |
| Discovered           | <ul><li>Is within</li><li>Is not within</li></ul>                                                                                         | <ul><li>Last 24 hours</li><li>Last 7 days</li><li>Last 30 days</li><li>Last 12 months</li><li>Year to date </li></ul> |
| Last seen            | <ul><li>Is within</li><li>Is not within</li></ul>                                                                                         | <ul><li>Last 24 hours</li><li>Last 7 days</li><li>Last 30 days</li><li>Last 12 months</li><li>Year to date</li></ul>  |
| Repository freshness | <ul><li>is one of</li><li>is not one of</li></ul>                                                                                         | <ul><li>Active</li><li>Inactive</li><li>Dormant</li></ul>                                                             |
| Source               | <ul><li>is one of</li><li>is not one of</li></ul>                                                                                         | <ul><li>azure-devops</li><li>GitHub</li><li>GitLab</li><li>Snyk</li></ul>                                             |
| Tags                 | <ul><li>containing one or more of</li><li>containing all of</li><li>not containing one or more of</li><li>not containing all of</li></ul> | all available tags you previously created                                                                             |



You can specify more than one filter component with an **And** or **Or** operator.\


<figure><img src="https://lh3.googleusercontent.com/OoGqaTS6E_XcIf6NB9q0r5wTxNHHoAI_oiD5aBlNIn4kqMenCdPcRsWrglS01DOppHudlg6X0aBED2SEYc6peiLJQarJ9pvRu6djZZ-1rb-7UuUFPljhawEO9lGobMHDSGzLDwHHpzqJdkQf2mCpfnw" alt="AppRisk - Create new policy"><figcaption><p>Snyk AppRisk - Create new policy</p></figcaption></figure>

## **Set actions**

After defining filter components, you need to define the actions that the policy has to perform on the filtered assets. Asset policies support the following actions:

* **Send Email** - Receive an email every time there are asset updates. You can choose between daily emails or scheduling the checks.
* **Send Slack Message** - Receive a Slack notification every time there are asset updates. You need to add your [Slack webhook URL](../../../integrate-with-snyk/notification-and-ticketing-systems-integrations/slack-integration.md), then you can choose between daily emails or scheduling the checks.
* **Set Asset Class** - Sets the class on the matched assets. Removing the policy or turning in off does not retroactively change the asset class back to default.
* **Set Asset Tag** - Sets a tag on the matched assets. Removing the policy or turning in off will remove the tags of this policy from the relevant assets.
* **Set Coverage Control Policy** - Sets a control on filtered assets that checks whether selected security products are scanning assets, optionally within a given timeframe. Assets that fail this control will be marked accordingly on inventory pages. This control applies the OR logic across products.

<figure><img src="../../../.gitbook/assets/image (1) (8).png" alt="AppRisk - Set a policy action"><figcaption><p>Snyk AppRisk - Set a policy action </p></figcaption></figure>

The editor supports multiple flows for the same policy. The flows can be independent or intersect.

<figure><img src="https://lh6.googleusercontent.com/YEy4S8gp_a2T8F02G_Wc4tY9571ZSyOXemao4v_Tb8SmWpGXEp7C-Eik1GX6gqE2hp-NQM6KNQ-EDx6xoHiyT-hL--znSsMSQoV0bQR9kKpNzP0p4ZGhoZG4mA8PjN1Hr-mO-o6NDmTg272rnbY9wYE" alt="AppRisk - Set multiple policy actions "><figcaption><p>Snyk AppRisk - Set multiple policy actions </p></figcaption></figure>

\
