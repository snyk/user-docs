# Organization reports

{% hint style="info" %}
**Feature availability**

Snyk Learn organization reports are available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

By default, only Snyk Org or Group admins can view and export reports. Group admins can create custom roles by using the standard Snyk workflow. Learn more about the access controls for reports at [Snyk Learn Access Controls](../snyk-learn-access-controls.md).&#x20;

{% hint style="info" %}
Snyk Learn organization reports support Organizations with up to 5000 members.
{% endhint %}

## Organization Overview report

The Overview report provides a high-level view of progress across your Organization. This report shows how many users in your Organization have started and completed each lesson or learning path.

<figure><img src="../../../.gitbook/assets/image (260).png" alt=""><figcaption><p>Snyk Learn Overview report</p></figcaption></figure>

### Organization Detailed report

The Detailed report provides individual user-level progress tracking within your Organization. This report includes the following progress data for specific lessons or learning paths:

* completion status
* when the lesson or learning path was completed
* when the lesson was previously completed, if lesson progress was reset after the user completed it&#x20;

<figure><img src="../../../.gitbook/assets/image (261).png" alt=""><figcaption><p>Snyk Learn Detailed report</p></figcaption></figure>

### Exporting Organization reports&#x20;

All reports are available as interactive table views and downloadable CSV reports.

User progress is associated with individual users, meaning that if they are members of multiple Organizations, their lesson progress is the same across all Organizations.

The report CSV contains the historical completions and current progress. The learning path overview report CSV contains the current learning path progress and the last completion date.

Progress reporting is also available through the Snyk Learn API (beta), offering two endpoints:

* [Org catalog progress](https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-/learn/progress/catalog): progress mapped to the Snyk Learn catalog
* [Org user progress](https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-/learn/progress/users): progress mapped to the Snyk user&#x20;
