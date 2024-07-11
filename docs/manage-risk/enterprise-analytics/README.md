# Analytics

{% hint style="warning" %}
**Release status**\
Analytics is in [Closed Beta](../../getting-started/snyk-release-process.md#closed-beta) and available only for Enterprise plans. For more information, see [Plans and pricing](https://snyk.io/plans).

During Closed Beta, access is available at the user level. When tenant membership is available, access to Analytics will be more easily managed at scale.

Group admins will be able to see all relevant data their users can access in the  Analytics report during the Closed Beta period. You can use filters to include relevant Groups or Organizations only. This is especially useful for sharing a link with another user.
{% endhint %}

{% hint style="info" %}
Data is refreshed in Analytics daily, approximately between 13:00 and 14:00 UTC.
{% endhint %}

Analytics provides executives, as well as Application Security (AppSec) leaders and practitioners a view into the performance of their AppSec program. Snyk customers can understand at a glance the strengths and weaknesses of their program, identify where successful practices can be discerned, and uncover the largest opportunities for improvement that warrant investment.

The following is an example of the Enterprise Analytics dashboard showing AppSec performance for your Groups and Organizations.&#x20;

Issues can be prevented by developers' taking advantage of Snyk Learn, leveraging the IDE plug-ins, activating PR checks, running `snyk test` locally in the CLI, breaking the build, or taking any other available actions for catching issues pre-production. If Snyk knows about an issue, a test can catch it. Other actions to stop preventable issues from getting into production environments include increasing the threshold for what breaks a build from `critical` to `critical` and `high` severity issues, or more strictly refraining from approving PRs that fail a Snyk test.

{% hint style="info" %}
The following example is also a preventable issue.

A developer adds a vulnerable version of a package as part of development. Despite a strong shift-left culture where running the `snyk test` command found this known issue, there was no fix available. The developer opted to use the vulnerable library instead of selecting an alternative because the issue had been known for a long time, or the developer felt the exploit was not relevant to how the package was being used, or both.
{% endhint %}

Enterprise Analytics is designed to make a top-level metric and its associated trend for each pillar visible at all times to support a quick understanding of the state of the program. You can navigate to a more granular view of each tab as needed. Different pillars and metrics on these views may be more or less relevant at different times to different companies, business units, products, teams, and any other participants.&#x20;

The Analytics view is structured as follows:

* [Enterprise Analytics](enterprise-analytics.md) - provides the exposure and performance details of Snyk issues in Groups and Organizations while focusing on the issue introduction method (baseline, preventable, or non-preventable).
* [ASPM Analytics](aspm-analytics/) - provides data analytics for reviewing and comparing assets and issues metrics at the level of asset classes, applications, or code owners.

The following table presents an overview of the features available for both Enterprise Analytics and Application Analytics.

| Enterprise Analytics                                                                                                                                                                                                                       | ASPM Analytics                                                                                                                                                                                                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| <ul><li>Data filtered by default on critical and high-severity issues.</li><li>Drill down to see the way that issues were introduced.</li><li>Issues framework: categorized based on Exposure, Manage, Prevention, and Coverage.</li></ul> | <ul><li>Data filtered based on assets, applications, and code owners (teams).</li><li>Helps you to identify and take action on risk, coverage gaps, and association gaps.</li><li>Asset class view</li><li>Application and owner view</li><li>Surface coverage gap</li><li>Comparison and prioritization</li></ul> |

{% hint style="info" %}
The specific features and availability of both products may vary as they continue to evolve. For the latest information, refer to the respective product documentation.
{% endhint %}
