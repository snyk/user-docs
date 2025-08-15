# Project attributes

Project attributes are static and non-configurable fields that allow you to apply attribute values to a Project by selecting from a predefined list of values. After you have applied attributes, you can remove them from a Project as needed.

Use attributes to group, prioritize, and filter Projects. Use attributes such as lifecycle stage, business criticality, and environment to prioritize issues at a granular level. After you apply attributes to Snyk [policies](../../manage-risk/policies/), you can assign policies to Projects that have those attributes applied.

## Available attributes and their values

On the **Projects** listing, use **Group by none** (ungrouped) for optimal Project visibility and to apply [tags](project-tags.md) and filtering attributes at the Project level.

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-23 at 18.07.46.png" alt="Project level filtering attributes"><figcaption><p>Project level filtering attributes</p></figcaption></figure>

The available Project attributes are summarized in the following table.

| Attribute                                                                                                                                                                                                                                                                                                              | Attribute options                                                                                                                                           |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <p>Business criticality<br>Note that When Risk Score is enabled, the Business Criticality attribute automatically affects the score according to the highest attribute level. Learn more in the <a href="../../manage-risk/prioritize-issues-for-fixing/risk-score.md#business-criticality">Risk Score</a> docs.  </p> | <ul><li>Critical</li><li>High</li><li>Medium</li><li>Low</li></ul>                                                                                          |
| Environment                                                                                                                                                                                                                                                                                                            | <ul><li>Frontend</li><li>Backend</li><li>Internal</li><li>External</li><li>Mobile</li><li>SaaS</li><li>On-Prem</li><li>Hosted</li><li>Distributed</li></ul> |
| Lifecycle stage                                                                                                                                                                                                                                                                                                        | <ul><li>Production</li><li>Development</li><li>Sandbox</li></ul>                                                                                            |

{% hint style="info" %}
You can apply attributes to Projects and remove attributes using endpoint [Applying attributes](../../snyk-api/reference/projects-v1.md#org-orgid-project-projectid-attributes).

You can also apply and remove attributes using Snyk CLI options, `--project-business-criticality`, `--project-environment`, and `--project-lifecycle`. Refer to the [CLI commands and options summary](../../developer-tools/snyk-cli/cli-commands-and-options-summary.md) for the commands that support these options.
{% endhint %}

{% hint style="info" %}
Roles with the **edit project attributes** permission can add an attribute to a Project.
{% endhint %}

## **Apply an attribute value to Project**

1. On the Project detail page, click **+ Add a value** below the attribute for which you want to apply a value to the Project.
2. Select a value from the list.

After you have selected a value for the attribute, it appears on the Project detail page. You can apply multiple values for an attribute and multiple attributes to a Project.

<figure><img src="../../.gitbook/assets/gs2.png" alt="Project detail page showing attribute values applied"><figcaption><p>Project detail page showing attribute values applied</p></figcaption></figure>

## **Remove attribute value**

1. Select the attribute with an applied value that you want to remove from the Project.
2. Click the **x** icon for the value you want to remove.

## **Filter by attribute values on the Projects listing page**

On the left of the **Projects** listing page, when Projects are grouped by none (ungrouped), select the values for the attributes that you want to filter Projects by.

When you filter by multiple values on a single attribute, Snyk returns a list of Projects to which one or more of the values in the filter have been applied.

When you filter by multiple attributes, Snyk returns a list of Projects to which values of multiple attributes have been applied.
