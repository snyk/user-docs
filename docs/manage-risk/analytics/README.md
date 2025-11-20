# Analytics

{% hint style="info" %}
Legacy Reports (in the Snyk Web UI) and the Reporting API v1 have been deprecated. They will be removed from the product on April 27, 2026. This extended timeline is designed to give your teams ample time to assess your usage and migrate to the new solution. View the API Migration guide for help transitioning to the newer [Export API](../../snyk-api/using-specific-snyk-apis/export-api-specifications-columns-and-filters.md).
{% endhint %}

Analytics provides executives, as well as Application Security (AppSec) leaders and practitioners a view into the performance of their AppSec program. Snyk customers can understand at a glance the strengths and weaknesses of their program, identify where successful practices can be discerned, and uncover the largest opportunities for improvement that warrant investment. Analytics is available at the tenant level.

{% hint style="info" %}
To access Analytics, you need to have one of the following [tenant roles](../../snyk-platform-administration/groups-and-organizations/tenant/manage-users-in-a-tenant.md): Tenant Admin, Tenant Viewer.
{% endhint %}

The Analytics view is structured as follows:

* [Issues Analytics](issues-analytics.md) - provides the exposure and performance details of Snyk issues in Groups and Organizations while focusing on the issue introduction method (baseline, preventable, or non-preventable).
* [Application Analytics](broken-reference) - provides data analytics for reviewing and comparing assets and issues metrics at the level of asset classes, applications, or code owners.

{% hint style="info" %}
**Feature availability**

Application Analytics is available only to Snyk AppRisk users.
{% endhint %}

The following table presents an overview of the features available for both Issues Analytics and Application Analytics.

| Issues Analytics                                                                                                                                                                                                                           | Application Analytics                                                                                                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| <ul><li>Data filtered by default on critical and high-severity issues.</li><li>Drill down to see the way that issues were introduced.</li><li>Issues framework: categorized based on Exposure, Manage, Prevention, and Coverage.</li></ul> | <ul><li>Data filtered based on assets, applications, and code owners (teams).</li><li>Helps you to identify and take action on risk, coverage gaps, and association gaps.</li><li>Asset class view</li><li>Application and owner view</li><li>Surface coverage gap</li><li>Comparison and prioritization</li></ul> |

{% hint style="info" %}
The specific features and availability of both products may vary as they continue to evolve. For the latest information, refer to the respective product documentation.
{% endhint %}
