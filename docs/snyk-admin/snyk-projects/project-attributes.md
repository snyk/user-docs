# Project attributes

Project attributes are static and non-configurable fields that allow you to apply attribute values to a Project by selecting from a predefined list of values. After you have applied attributes, you can remove them from a Project as needed.

Use attributes to group, prioritize, and filter Projects. Use attributes such as lifecycle stage, business criticality, and environment to prioritize issues at a granular level. After you apply attributes to Snyk [policies](../../scan-with-snyk/policies/), you can assign policies to Projects that have those attributes applied.

## **Available attributes and their values**

On the **Projects** listing, use **Group by none** (ungrouped) for optimal Project visibility and to apply [tags](../introduction-to-snyk-projects/project-tags.md) and filtering attributes at the Project level.

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-23 at 18.07.46 (1) (1) (1) (1) (1) (1) (3).png" alt="Project level filtering attributes"><figcaption><p>Project level filtering attributes</p></figcaption></figure>

The available Project attributes are summarized in the following table.

| Attribute                                                                                                                                                                                                                                                                                    | Attribute options                                                                                                                                           |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <p>Business criticality<br>Note that When Risk Score is enabled, the Bussiness Criticality attribute automatically affects the score according to the highest attribute level. Learn more in the <a href="../../manage-issues/risk-score.md#business-criticality">Risk Score</a> docs.  </p> | <ul><li>Critical</li><li>High</li><li>Medium</li><li>Low</li></ul>                                                                                          |
| Environment                                                                                                                                                                                                                                                                                  | <ul><li>Frontend</li><li>Backend</li><li>Internal</li><li>External</li><li>Mobile</li><li>SaaS</li><li>On-Prem</li><li>Hosted</li><li>Distributed</li></ul> |
| Lifecycle stage                                                                                                                                                                                                                                                                              | <ul><li>Production</li><li>Development</li><li>Sandbox</li></ul>                                                                                            |

{% hint style="info" %}
You can apply attributes to Projects and remove attributes using the Snyk API v1 endpoint [Applying attributes](https://snyk.docs.apiary.io/#reference/projects/project-attributes/applying-attributes).

You can also apply and remove attributes using Snyk CLI options, `--project-business-criticality`, `--project-environment`, and -`-project-lifecycle`. Refer to the [CLI commands and options summary](../../snyk-cli/cli-commands-and-options-summary.md) for the commands that support these options.
{% endhint %}

{% hint style="info" %}
Organization admins can add an attribute to a Project. However, only Group admins can modify Project attributes in cases where attributes match a policy, because policies can only be managed by Group admins.
{% endhint %}

## **Apply an attribute value to Project**

1. On the Project detail page, click the **+** icon below the attribute for which you want to apply a value to the Project.
2. Select a value from the list.

<figure><img src="../../.gitbook/assets/gs1.png" alt="Select an attribute value to apply to the Project"><figcaption><p>Select an attribute value to apply to the Project</p></figcaption></figure>

After you have selected a value for the attribute, it appears on the Project detail page. You can apply multiple values for an attribute and multiple attributes to a Project.

<figure><img src="../../.gitbook/assets/gs2.png" alt="Project detail page showing attribute values applied"><figcaption><p>Project detail page showing attribute values applied</p></figcaption></figure>

## **Remove attribute value**

1. Select the attribute with an applied value that you want to remove from the Project.
2. Click the **x** icon for the value you want to remove.

<figure><img src="../../.gitbook/assets/gs3.png" alt="Removing Frontend attribute value"><figcaption><p>Removing Frontend attribute value</p></figcaption></figure>

The attribute value is removed from the Project.

<figure><img src="../../.gitbook/assets/gs4.png" alt="Project detail showing only Mobile attribute value"><figcaption><p>Project detail showing only Mobile attribute value</p></figcaption></figure>

## **Filter by attribute values on the Projects listing page**

On the left of the **Projects** listing page, when Projects are grouped by none (ungrouped), select the values for the attributes that you want to filter Projects by.

When you filter by multiple values on a single attribute, Snyk returns a list of Projects to which one or more of the values in the filter have been applied.

When you filter by multiple attributes, Snyk returns a list of Projects to which values of multiple attributes have been applied.
