# Project attributes

Project attributes are static and non-configurable fields that allow you to add additional metadata (also known as _values_) to a project. Attributes have a predefined list of values you can select from. Values can be assigned and removed from attributes and can be used to group, prioritize, and filter projects.

## **Available attributes and their values**

In the **Projects** menu, use **Group by none** (ungrouped) for better Project visibility and to apply [tags](project-tags.md) and filtering attributes at the Project level. \*\*\*\*

<figure><img src="../../.gitbook/assets/project-attributes_20sept2022.png" alt=""><figcaption><p>Project level filtering attributes</p></figcaption></figure>

The available Project attributes are summarized in the following table.

| Attribute            | Attribute options                                                                                                                                           |
| -------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Business criticality | <ul><li>Critical</li><li>High</li><li>Medium</li><li>Low</li></ul>                                                                                          |
| Environment          | <ul><li>Frontend</li><li>Backend</li><li>Internal</li><li>External</li><li>Mobile</li><li>SaaS</li><li>On-Prem</li><li>Hosted</li><li>Distributed</li></ul> |
| Lifecycle stage      | <ul><li>Production</li><li>Development</li><li>Sandbox</li></ul>                                                                                            |

{% hint style="info" %}
Filtering is available in projects by their values from [Snyk's REST API](../../snyk-api-info/).

Assigning and removing attribute values can be done in [Snyk's CLI](../../snyk-cli/) or the API.
{% endhint %}

## **Assigning values to attributes**

1. On the project page, click the **+** icon below the attribute you want to assign the value
2. Select a value from the list available

![](../../.gitbook/assets/gs1.png)

After you have assigned a value to the attribute, it appears on the project listing page. Each attribute can have multiple values assigned to it, and you can assign values to multiple attributes.

![](../../.gitbook/assets/gs2.png)

## **Removing values from attributes**

1. Select the attribute you want to remove a value from.
2. Click the **x** for the value.

![](../../.gitbook/assets/gs3.png)

The value is removed from the attribute.

![](../../.gitbook/assets/gs4.png)

## **Filtering values in the project listing page**

1. On the left of the project listing page, select the values for the attributes that you want to filter projects by.
2. When you filter by multiple values on a single attribute, you return projects that have been assigned one or more of the values in the filter.
3. When you filter by multiple attributes, you return projects which have been assigned values of both attributes in the filter.
4. Different filtering options are available depending on the grouping option you choose.
   * Filters for **Group by target**:
     * Show - with/without issues
     * Integrations
   * Filters for **Group by none** (ungrouped):
     * Show - with/without issues, active/inactive projects
     * Integrations, Tags, Business Criticality, Environment, Lifecycle

<figure><img src="../../.gitbook/assets/Project attributes.png" alt="Screenshot of filters in the Snyk Projects listing page"><figcaption></figcaption></figure>
