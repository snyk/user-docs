# REST Issues experimental API to GA API migration guide

{% hint style="info" %}
**Important information about Experimental APIs**

An experimental endpoint should be considered unstable and regarded as a tech preview. Experimental versions may introduce breaking changes and may be discontinued at any time.
{% endhint %}

## What's new in the GA REST Issues API?

{% hint style="info" %}
GA REST Issues API documentation [`/groups/{group_id}/issues`](https://apidocs.snyk.io/#get-/groups/-group_id-/issues) and [`/orgs/{org_id}/issues`](https://apidocs.snyk.io/#get-/orgs/-org_id-/issues). Updated link: [Get issues by Org ID](../../reference/issues.md#get-orgs-org_id-issues)
{% endhint %}

This version of the API delivers:

* **Consistency:** Improved performance and reliability of the REST Issues API
* **Depth:** detailed representations for Open Source packages and fixes
* **Flexibility:** new filters for tailored API responses
* **Usability:** improved pagination and response management, simplifying the API interaction

Snyk understands that migrating to a new API can be a significant undertaking and wants to support you throughout the process. This comprehensive migration guide is intended to facilitate a seamless transition by providing step-by-step instructions, code examples, and best practices to help you smoothly integrate with the new API.

If you are using the deprecated endpoint, Snyk encourages you to review this migration guide and move all your automations over.

## Comparison of experimental vs GA API

{% hint style="info" %}
**Mapping experimental API issues to GA API issues**

One of the main differences you will see in the table below is that the format of the ID for an issue changes from URI format (consists of key & scan\_item.id) in the Experimental API to UUID in the GA API.  To match an issue in the experimental api response to the same issue in the GA API response you can use **key** and the **scan\_item.id**. Note that **scan\_item** is part of the **relationships** block and **key** is part of the **attributes** block.
{% endhint %}

<table><thead><tr><th width="314">Fields</th><th width="176">Experimental</th><th>GA</th></tr></thead><tbody><tr><td>classes</td><td>Present</td><td>No change</td></tr><tr><td>coordinates</td><td>Only available for cloud issues</td><td>Available for cloud and SCA issues and has new fixability fields.</td></tr><tr><td><p>coordinates.is_fixable_manually</p><p>coordinates.is_fixable_snyk</p><p>coordinates.is_fixable_upstream</p><p>coordinates.is_patchable</p><p>coordinates.is_pinnable</p><p>coordinates.is_upgradeable</p><p><br></p></td><td>Not present</td><td>Newly introduced fixability data</td></tr><tr><td>coordinates.reachability</td><td>Not present</td><td>Newly introduced</td></tr><tr><td>coordinates.remedies</td><td>Present</td><td>No change</td></tr><tr><td>representations</td><td>Present</td><td>New fields</td></tr><tr><td>representations.resourcePath</td><td>Present</td><td>No change</td></tr><tr><td>respresentations.dependencyChain</td><td>Present</td><td>Removed in favor of representations.dependency</td></tr><tr><td>representations.dependency</td><td>Not present</td><td><p>Newly introduced replaces representations.</p><p>dependencyChain</p></td></tr><tr><td><p></p><p>representations.dependency</p><p>.package_name</p><p></p><p>representations.dependency.</p><p>package_version</p><p><br></p></td><td>Not present</td><td>Newly introduced as part of represenations.dependency</td></tr><tr><td>cloud_resource</td><td>Present</td><td>No change</td></tr><tr><td>sourceLocation</td><td>Present</td><td>No change</td></tr><tr><td>created_at</td><td>Present</td><td>No change</td></tr><tr><td>description</td><td><br></td><td>No change</td></tr><tr><td>effective_severity_level</td><td>Present</td><td>No change</td></tr><tr><td>ignored</td><td>Present</td><td>No change</td></tr><tr><td>key</td><td>Present</td><td>No change</td></tr><tr><td>priority</td><td>Present</td><td>Removed and replaced with risk</td></tr><tr><td>priority.factors</td><td>Present</td><td>Replaced with risk.factors</td></tr><tr><td>priority.score</td><td>Present</td><td>Replaced with risk.score</td></tr><tr><td>risk</td><td>Not present</td><td>Newly introduced - replaces priority</td></tr><tr><td><a data-footnote-ref href="#user-content-fn-1"><mark style="background-color:orange;">risk.factors</mark></a></td><td>Not present</td><td>Newly introduced - replaces priority.factors</td></tr><tr><td><p>risk.factors[i].included_in_score</p><p>risk.factors[i].links</p><p>risk.factors[i].links.evidence</p><p>risk.factors[i].links.evidence.href</p><p>risk.factors[i].links.evidence.meta</p><p>risk.factors[i].name</p><p>risk.factors[i].updated_at</p><p>risk.factors[i].value</p></td><td>Not present</td><td>Newly introduced</td></tr><tr><td><a data-footnote-ref href="#user-content-fn-2"><mark style="background-color:orange;">risk.score</mark></a></td><td>Not present</td><td>Newly introduced replaces priority.score</td></tr><tr><td><p>risk.score.model</p><p>risk.score.updated_at</p><p>risk.score.value</p></td><td>Not present</td><td><br></td></tr><tr><td>problems</td><td>Present</td><td>No change</td></tr><tr><td>resolution</td><td>Only for cloud issues</td><td>For all issue types</td></tr><tr><td>severities</td><td>Present - Not populated</td><td>Removed as not populated and will not be</td></tr><tr><td>status</td><td>Present</td><td>No change</td></tr><tr><td>title</td><td>Present</td><td>No change</td></tr><tr><td>tool</td><td>Present</td><td>No change</td></tr><tr><td>type</td><td>Present</td><td>No change</td></tr><tr><td>updated_at</td><td>Present</td><td>No change</td></tr><tr><td><a data-footnote-ref href="#user-content-fn-3"><mark style="background-color:orange;">id</mark></a></td><td>URI format (consists of key &#x26; scan_item.id)</td><td>UUID format</td></tr><tr><td>relationships.ignore</td><td>Present</td><td>No change</td></tr><tr><td>relationships.organization</td><td>Present</td><td>No change</td></tr><tr><td>relationships.policies</td><td>Present - Not populated</td><td>Removed as not populated and will not be</td></tr><tr><td>relationships.previous</td><td>Present - Not populated</td><td>Removed as not populated and will not be</td></tr><tr><td>relationships.scan_item</td><td>Present</td><td>No change</td></tr><tr><td>relationships.test_execution</td><td>Present</td><td>No change</td></tr></tbody></table>

## Comparison of filters in experimental vs GA API

| Filter by                  | Purpose                                                   | Experimental | GA               |
| -------------------------- | --------------------------------------------------------- | ------------ | ---------------- |
| starting\_after            | return the page of results immediately after this cursor  | Present      | No change        |
| ending\_before             | return the page of results immediately before this cursor | Present      | No change        |
| limit                      | number of results to return per page                      | Present      | No change        |
| scan\_item.id              | filter issues through their scan item relationship        | Present      | No change        |
| scan\_item.type            | filter issues through their scan item relationship        | Present      | No change        |
| type                       | filter by issue type                                      | Present      | No change        |
| updated\_after             | filter issues updated after this date                     | Present      | No change        |
| updated\_before            | filter issues updated before this date                    | Not present  | Newly introduced |
| created\_before            | filter issues created before this date                    | Not present  | Newly introduced |
| created\_after             | filter issues created after this date                     | Not present  | Newly introduced |
| effective\_severity\_level | filter issues by one or more effective severity levels    | Not present  | Newly introduced |
| status                     | filter by an issues status                                | Not present  | Newly introduced |
| ignored                    | filter issues by their ignored status                     | Not present  | Newly introduced |

[^1]: The only risk factors presently in the issues API are insights factors.

[^2]: &#x20;For customers with risk score enabled this should be the risk score, but without any risk factors. For customers w/o the risk score enabled, this is the old priority score.

[^3]: As noted above,  to match an issue in the experimental api response to the same issue in the GA API response you can use key and the scan\_item.id. Note that scan\_item is part of the relationships block and key is part of the attributes block.
