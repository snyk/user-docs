# Issue card information

Issue cards appear on the details page for a Project. You can use available options to do the following:

* [View issue card information](issue-card-information.md#view-issue-card-information).
* [Expand an issue card to show more details](issue-card-information.md#expand-an-issue-card-to-show-more-details).
* [Filter and sort issue cards](issue-card-information.md#filter-and-sort-issue-cards).
* [Perform additional card actions](issue-card-information.md#perform-additional-card-actions).

## View Issue card information

Issue cards show details for a specific vulnerability or license issue and actions you can take.

<figure><img src="../../.gitbook/assets/image (289) (1).png" alt="Issue card for the npmconf vulnerability"><figcaption><p>Issue card for the npmconf vulnerability</p></figcaption></figure>

The issue card provides a [Header section](issue-card-information.md#header-section) and [Body section](issue-card-information.md#body-section) with information as explained in the next sections of this documentation.

### Header section

<figure><img src="../../.gitbook/assets/issue-card-header_8dec2022.png" alt="npmconf issue card header"><figcaption><p>npmconf issue card header</p></figcaption></figure>

* [Severity level](../priorities-for-fixing-issues/severity-levels.md): for example, **High**
* Issue name: for example, **Uninitialized Memory Exposure**, with a link that can be copied
* **Score**: [Priority score](../priorities-for-fixing-issues/priority-score.md): 0 to 1,000
* Type: **VULNERABILITY** or LICENSE ISSUE
* Links to [CWE](https://cwe.mitre.org/index.html) (Common Weakness Enumeration), [CVSS](https://www.first.org/cvss/calculator/3.1) (Common Vulnerability Scoring System), and [Snyk Vulnerability Database](https://snyk.io/vuln) information for the issue. You can use these links to view more information about the CWE, CVE, and CVSS scores or navigate to the Snyk Vulnerability Database information for a specific vulnerability from its issue card.

### Body section

<figure><img src="../../.gitbook/assets/issue-card-body-eg1_8dec2022.png" alt="npmconf issue card body details"><figcaption><p>npmconf issue card body details</p></figcaption></figure>

* **Introduced through**: The path through which the vulnerability or license was introduced
* **Fixed in:** The file the vulnerability is fixed in
* [**Exploit maturity**](../priorities-for-fixing-issues/view-exploits.md): for example, **Mature** or **Proof Of Concept**
* **Reachability**: for example, **Reachable**. For information and an example, see [Reachable vulnerabilities](../priorities-for-fixing-issues/reachable-vulnerabilities.md)
* **Social Trends**: Snyk shows a [Trending](../priorities-for-fixing-issues/vulnerabilities-with-social-trends.md) banner for issues that are being actively discussed on Twitter.

## Expand an issue card to show more details

When you expand the issue card to show more information, you can see details for the vulnerability, including:

* Detailed path information
* Fix advice
* Overview of the vulnerability
* Any vulnerable functions within the vulnerability

<figure><img src="../../.gitbook/assets/image-card-expanded_8dec2022.png" alt="Details for the npmconf vulnerability"><figcaption><p>Details for the npmconf vulnerability</p></figcaption></figure>

## Filter and sort issue cards

You can apply multiple filters to a Project to show a set of issues based on specific criteria:

* Vulnerability or license issue
* Issues with a specific severity
* Issues within a range of the priority score
* Issues that have an exploit and how mature the exploit is
* Issues that are open or have been patched or ignored

You can sort the issue cards in a Project based on their priority score or severity.

## Perform additional card actions

You can perform the following actions on the issue card:

* [Ignore the issue](../priorities-for-fixing-issues/ignore-issues.md): if you do not need to take action on an issue, or it does not need to appear on your reports, you can ignore it.
* [Create a Jira ticket](../../integrations/notifications-ticketing-system-integrations/jira.md): if you have the Jira integration, you can link your issue boards to Snyk and create Jira tickets directly from the Project details page to fix vulnerabilities.
* [Fix the vulnerability](../../scan-application-code/snyk-open-source/starting-to-fix-vulnerabilities/fix-your-vulnerabilities.md): if a fix is available, you can fix individual vulnerabilities.
