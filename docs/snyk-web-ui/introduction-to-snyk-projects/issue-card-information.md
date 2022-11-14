# Issue card information

Issue cards appear in the details page for your project - see [View project information](https://docs.snyk.io/getting-started/introduction-to-snyk-projects/view-project-information).

Issue cards show details for a specific vulnerability or license issue, and actions you can take with it.

![](../../.gitbook/assets/issue-card\_trending-tag\_14sept2022.png)

Header section information:

![](../../.gitbook/assets/issue-card\_header-section\_14sept2022.png)

* [Severity level](https://docs.snyk.io/introducing-snyk/snyks-core-concepts/severity-levels): for example, **H** (High) or **C** (Critical)
* **Issue name**: for example, **Expression Language Injection**
* **Score**: [Priority score](https://docs.snyk.io/fixing-and-prioritizing-issues/starting-to-fix-vulnerabilities/snyk-priority-score): 0 - 1,000
* **Type**: VULNERABILITY or LICENSE ISSUE
* Links to [CWE](https://cwe.mitre.org/index.html) (Common Weakness Evaluation), [CVSS](https://www.first.org/cvss/calculator/3.1) (Common Vulnerability Scoring System), and Snyk [Intel Vulnerability DB](https://snyk.io/vuln) information for that issue

Body section information:

![](../../.gitbook/assets/issue-card\_body-section\_14sept2022.png)

* **Introduced through**: The path the vulnerability or license was introduced through
* **Fixed in:** The file the vulnerability is fixed in
* [Exploit maturity](https://docs.snyk.io/fixing-and-prioritizing-issues/issue-management/evaluating-and-prioritizing-vulnerabilities): for example, **Proof Of Concept**.
* [Reachability](https://support.snyk.io/hc/en-us/articles/360010554837-Reachable-Vulnerabilities-): for example, **Reachable**.
* **Social Trends**: Snyk shows a [Trending banner](https://docs.snyk.io/fixing-and-prioritizing-issues/prioritizing-issues/prioritize-by-social-trends) on issues that are being actively discussed on Twitter.

When you expand the issue card for more information, you can see additional information.

![](../../.gitbook/assets/issue\_card-expanded\_14sept2022.png)

This information includes:

* Detailed path information
* Fix advice
* Overview of the vulnerability
* Any vulnerable functions within the vulnerability

You can apply multiple filters to a project to show a set of issues:

* If they are a vulnerability or a license issue
* With a specific severity
* Within a range of the priority score
* Based on whether it has an exploit, and how mature the exploit is
* That are open or have been patched or ignored

The issue cards in a project can be sorted based on their priority score or severity.

## Card actions

You can perform the following actions on the issue card:

* ****[**Ignore the issue**](../../features/fixing-and-prioritizing-issues/issue-management/ignore-issues.md): if you don't need to take action on an issue, or it doesn't need to appear on your reports, you can ignore it.
* [**Create a Jira ticket**](https://docs.snyk.io/integrations/untitled-3/jira): if you have the [Jira integration](https://docs.snyk.io/integrations/untitled-3/jira), you can link your issue boards to Snyk and create Jira tickets directly from the project details page to fix vulnerabilities.
* [**Fix the vulnerability**](https://docs.snyk.io/snyk-open-source/open-source-basics/fixing-vulnerabilities): if a fix is available, you can fix individual vulnerabilities.
* **View more information about the CWE, CVE, and CVSS scores**: navigate from the issue card to further information about these scores.
* **View the Snyk vulnerability database:** navigate to the Snyk vulnerability database about a specific vulnerability from its issue card.
