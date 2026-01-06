# Phase 4: Create a fix strategy

After you set up your integrations, create your Organizations, and import your Projects, view reports for visibility into the current vulnerability backlog for your business.

{% hint style="info" %}
For more information on how to use reports, see [Reporting](../../manage-risk/reporting/) and [Available Snyk Reports](../../manage-risk/reporting/available-snyk-reports.md).
{% endhint %}

In most cases, when you look across all of your Organizations and vulnerabilities, there will be many issues, and it can be difficult to know where to start.

## Decide important focus areas

Before diving into specific vulnerabilities, consider your Organizations and repositories (Targets in Snyk) to identify any areas of specific importance. The discovery of business-critical applications you did in phase 1 can provide information that helps you prioritize during this phase. For example, if your Organizations match different products, you can initially focus on the one Organization (product) with the most users or where security matters most.

In that Organization, you can consider the repositories that make up the application's different parts. Areas that handle sensitive data, or are public-facing, may be more important to secure, so considering the repositories could be another way to narrow down your initial list of Projects to review.

{% hint style="info" %}
If you have used **Attributes** or **Tags** to add metadata to your Projects, you can filter and reduce the number of Projects that you are considering based on these attributes and tags.
{% endhint %}

## Group work by development teams

After you have your reduced set of Projects to prioritize, you may want to split these between different development teams. For example, you may have one development team fixing issues with your open-source and first-party code, and a separate DevOps team responsible for containers and base image vulnerabilities.

## Fix based on prioritization methods

Filters are available to help prioritize what issues need to be fixed urgently. The following search criteria are most commonly used when building a prioritization plan and can be used iteratively or in combination as you analyze results.

* If you wish to filter based on Priority or Risk Score versus Severity:
  * Snyk uses a wide variety of information to assign a score to an issue. You could filter on 900+ for the first round, then 800+, then 700+, and so on to work through triaging by Priority Score. If you are using Risk Score, follow a similar process starting from the highest risk items and working through them. Alternatively, if you prefer to filter on your own metrics, or a combination of Score and other key metrics, then your process can be filtering individually on attributes.
  * Severity: start with **High** and **Critical**. It is common to filter by Critical severity; however, if you are using Snyk Code and Snyk Open Source:
    * Filtering by Critical and High ensures Snyk Code issues are not filtered out.
    * Alternatively, filtering on Critical severity for Snyk Open Source, then filtering on severity=high, product=Snyk Code for Snyk Code independently, ensures you identify the most severe issues for each scan type.
* If you wish to focus on Open Source during triage:
  * By choosing the Exploit Maturity filter, you implicitly filter only open-source results. [Exploit Maturity](https://snyk.io/blog/whats-so-wild-about-exploits-in-the-wild-and-how-can-we-prioritize-accordingly/) addresses issues that are **Mature** or have **Proof of Concept** and thus are more exploitable.
  * By choosing the Fixable filter, you can find issues where there is a fix available by upgrading a package. This makes fixing faster than without an upgrade.
  * You can choose CVSS Score for Open Source Vulnerabilities.
* When you filter by [Priority Score](../../manage-risk/prioritize-issues-for-fixing/priority-score.md), CVSS Score is included in the calculation. One strategy is to eliminate the vulnerabilities with a score of 900-1000, and then move to vulnerabilities with a score of 800-900, and so on.

When planning your fix strategy, decide which metrics will be used on a specific timeline. If you choose to fix by severity, for example, estimate the time it will take to resolve vulnerability per severity. It is recommended to be specific with your fix strategy.

{% hint style="info" %}
**Example fix strategy**

If there are fifty Critical severity issues, and one hundred High severity issues, you may plan on two weeks to fix critical vulnerabilities and then four weeks to fix high severity issues, based on the size of your team and workload.
{% endhint %}

## Fix by issue type

Alternatively, you can fix based on issue type. It is common to focus initially on specific issue types, for example, open-source vulnerabilities, as then you can more easily compare issues across different Projects.

Examples of processes follow for fixes led by different types of teams.

### Example: Developer-led priorities

Developer-led implementation, with instruction from executives to minimize license risk, follows this process:

1. Resolve all License issues identified by Snyk Open Source.
2. Address all fixable Critical or High vulnerabilities in Snyk Open Source.
3. Shift focus to Code Analysis Projects using Snyk Code, starting with High severity issues.
4. Scan the containers and IaC files being used to run applications and environments.

### **Example: DevSecOps-led priorities**

DevSecOps-led implementation, focusing on securing your custom images and environment, follows this process:

1. Scan and secure any custom base images that your development teams pull from.
2. Integrate with your container registry and scan the images you provide to your development teams.
   * Scan the image you have chosen and built after adding common internal tools and standardizing image parameters.
   * Ensure developers also scan their containers after adding their custom tools and packages, to ensure that the container remains secure before deployment. This scan will also detect application vulnerabilities.
3. After you have secured your containers, start actively scanning your IaC files and cloud environments for misconfigurations that could lead to security breaches.
4. Introduce Snyk to development teams to scan their open source and code to reduce application vulnerabilities.

## Targeted vulnerabilities campaigns

When you operationalize security testing in your development process, another option for your fix strategy is to have campaigns to eliminate vulnerability types, for example, SQL injection. Using CWE filters can be very helpful in reporting to identify and log issues.

## Update your timeline

After you have created your fix strategy, update the timeline for [Phase 7: Triages, ignores, and fixes](phase-7-triages-ignores-and-fixes.md).
