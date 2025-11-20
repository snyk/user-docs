# Prioritize issues for fixing

You can find prioritization within Snyk under several names and with different customizations depending on your Snyk plan. The following list presents an overview of all the places where you can find and use prioritization. The outputs can vary depending on the type of priority you choose to use.

* [Prioritization at the Project level](./#prioritization-at-the-project-level)
* [Prioritization within Reporting](./#prioritization-within-reporting)
* [Prioritization based on risk](./#prioritization-based-on-risk)
* [Prioritization based on coverage and security program](./#prioritization-based-on-coverage-and-security-program)

You can focus solely on one type of prioritization, or you can combine them for a more complex and targeted focus on the issues that require your immediate attention. You can use prioritization for your vulnerabilities, your security program, and your coverage (from a security perspective), or you can use it in the Reporting section of your Projects.\
\
For example, if you only prioritize the Projects from your repositories, then the list of issues is more general. However, if you choose to use prioritization with Issues, available for Snyk Essentials, the analyzed issues are filtered through several other factors (risk factors, assets) all contributing to the creation of an issues priority list based on more complex and targeted filters, allowing you to customize the prioritization based on your needs.

## Prioritization at the Project level

### Prioritize the Projects

[Prioritize the Projects](../../snyk-platform-administration/snyk-projects/project-information.md) from your repositories by the type of Projects you want to focus on, for example, prioritize your Projects depending on their severity.

With this type of prioritization, you can group, filter, and sort your Projects.

Prioritize Projects based on these filters:

* With issues: displays only Projects with issues.
* Without issues: displays only Projects without issues.
* Integrations: displays the integrated Git repositories imported to Snyk.

You can use this type of prioritization with any of the following Snyk Plans:

* Snyk Free/Team Plan
* Snyk Enterprise Plan
* Snyk Essentials
* Snyk AppRisk

You can prioritize your Projects at the Snyk Group level and Snyk Organization level.

### Prioritize the issues within a Project

[Prioritize the issues](../../snyk-platform-administration/snyk-projects/view-project-issues-fixes-and-dependencies.md) within a Project by using any of the following custom filters:

* Issue type
* Severity
* Fixability
* Exploit Maturity
* Status

You can use this type of prioritization with any of the following Snyk Plans:

* Snyk Free/Team Plan
* Snyk Enterprise Plan
* Snyk Essentials
* Snyk AppRisk

You can prioritize your Projects at the Snyk Organization level.

## Prioritization within Reporting

### Priority Score and Risk Score prioritization

You can use both Priority and Risk Scores as a prioritization tool when you are using the Reporting option from the Snyk Web UI and you select the desired Projects from the filters list.

The Priority Score focuses on urgency and ranks security issues based on various criteria, such as CVSS score, trending vulnerabilities, reachability, availability of exploits, and other factors. On the other hand, the Risk Score assesses the overall risk associated with vulnerabilities by considering both their severity and the context of the application, such as the likelihood of exploitation and the potential impact on the system. While the Priority Score assists teams in addressing the most urgent threats, the Risk Score offers a comprehensive view of the security posture.

#### Priority Score

[Priority Score](priority-score.md) helps teams quickly identify and address critical security vulnerabilities by ranking them based on urgency.

Priority Score is determined based on a number of industry-standard criteria:

* CVSS score
* trending vulnerabilities
* reachability
* availability of exploits
* other factors

You can use this type of prioritization with any of the following Snyk Plans:

* Snyk Free/Team Plan
* Snyk Enterprise Plan
* Snyk Essentials
* Snyk AppRisk

You can use the Priority Score at the Snyk Organization level.

#### Risk Score

[Risk Score](risk-score.md) assesses the potential impact of vulnerabilities, prioritizing those with severe consequences. You can use the Risk Score to perform automatic risk analysis for each security issue based on the potential impact and likelihood of exploitability.

You can use this type of prioritization with any of the following Snyk Plans:

* Snyk Free/Team Plan
* Snyk Enterprise Plan
* Snyk Essentials
* Snyk AppRisk

You can use the Risk Score at the Snyk Organization level.

### Prioritize issues when using reports

You can [prioritize issues when using reports](../reporting/available-snyk-reports.md) to generate reports across an Organization or Group.

You can use this type of prioritization with any of the following Snyk Plans:

* Snyk Enterprise Plan
* Snyk Essentials
* Snyk AppRisk

You can use the Risk Score at the Snyk Group or Snyk Organization level.

## Prioritization based on risk

[Prioritization with Issues](prioritization-for-snyk-essentials.md) for Snyk AppRisk uses holistic application intelligence to help you identify and prioritize container, code, and open source issues based on the risk they pose to your application. You can prioritize issues based on asset classification as defined in Snyk AppRisk policies.

You can use Prioritization with Insights to identify and prioritize your Container, Code, and Open Source issues based on the risk they pose to your application. These are the risk factors that you can use:

* Deployed
* Loaded package
* OS condition
* Public facing

You can use this type of prioritization with the following Snyk Plan:

* Snyk AppRisk

You can use the Prioritization with Insights at the Snyk Group or Snyk Organization level.

## Prioritization based on coverage and security program

Snyk Issues Analytics provides a comprehensive overview of your entire application security program, offering insights into strengths, weaknesses, and opportunities at the tenant level. On the other hand, Snyk Application Analytics focuses specifically on the security and performance of individual applications. This narrower analysis helps pinpoint specific app vulnerabilities and inefficiencies, enabling targeted improvements and more detailed insights at the Snyk Group or Organization level.

### Issue Analytics

[Issue Analytics](../analytics/issues-analytics.md) provides a view into the performance of your AppSec program. You can use Issues Analytics to better understand the strengths and weaknesses of your program, identify where successful practices can be discerned, and uncover the largest opportunities for improvement that warrant investment.

You can use this type of prioritization with the following Snyk Plan:

* Snyk Enterprise
* Snyk AppRisk

You can use Issues Analytics at the Snyk tenant level, including both the Snyk Group and Snyk Organization levels.

### Application Analytics

[Application Analytics](broken-reference) for Snyk AppRisk helps you assess your AppSec program from a top-down approach. It allows you to review applications, teams, and asset classes, and then focus on specific assets. You can improve security by identifying areas for enhancement, recognizing risks, and addressing blind spots. The tool retrieves data from all available groups for the tenant.

Prioritize the displayed data by using the available filters, dimension views, and specific timeframes.\
\
Filters:

* Groups
* Issue severity
* Asset-based filters

View by:

* Asset Class
* Application
* Owner

You can use this type of prioritization with the following Snyk Plan:

* Snyk AppRisk

## Prioritization strategies

Snyk has several features that help you determine which issues you discover are the most important for you to fix and the sequence in which to fix the issues.

{% hint style="info" %}
For information on how to ignore and exclude issues, see [The .snyk file](../policies/the-.snyk-file.md) and the [Policies ](../policies/)pages that explain how to create policies and assign them to Projects, as well as Security and License policies.
{% endhint %}

Some tools use only the single factor of severity to prioritize issues, but this can result in thousands of results, with no clear starting point for fixing these issues.

You can prioritize at the Project level when looking at a specific Project. Enterprise customers can prioritize across all Projects.

Snyk Priority Score and Risk Score rank the [severity](severity-levels.md) of an issue and the urgency of fixing it. For details, see [Priority Score vs Risk Score](priority-score-vs-risk-score.md), [Priority Score](priority-score.md), and [Risk Score](risk-score.md).

You can [ignore issues](ignore-issues/) and [triage issues](vulnerable-conditions.md) to establish your issue management strategy.

[View exploits](view-exploits.md) to see how vulnerabilities can be taken advantage of. You can then start evaluating and prioritizing vulnerabilities using guidance from the [Snyk Priority Score](priority-score.md) for each issue.

Consider [Malicious packages](malicious-packages.md) and how to address them in your Projects.

You can set up [reachable vulnerability analysis ](reachability-analysis.md)to identify vulnerabilities with a path to your code. This helps you asse are calculated as part of the priority score.

[Vulnerabilities with Social Trends](vulnerabilities-with-social-trends.md) are calculated as part of the Priority Score.

Based on your priorities, you can start [fixing vulnerabilities](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/).

See [Prioritize issues in the Snyk Web UI](https://learn.snyk.io/lesson/prioritize-issues-snyk) to learn about prioritization in action.

You can use many features of [Snyk Projects](../../snyk-platform-administration/snyk-projects/) to help you focus on priority issues:

* [View Project information](../../snyk-platform-administration/snyk-projects/project-information.md).
* Apply and remove [Project attributes ](../../snyk-platform-administration/snyk-projects/project-attributes.md)and [Project tags](../../snyk-platform-administration/snyk-projects/project-tags.md) to characterize Projects.
* Look at [Project collections groupings](../../snyk-platform-administration/snyk-projects/project-collections-groupings/).
* [View Project issues, fixes, and dependencies](../../snyk-platform-administration/snyk-projects/view-project-issues-fixes-and-dependencies.md).
* [View Project settings](../../snyk-platform-administration/snyk-projects/view-and-edit-project-settings.md).
