# Manage your AppSec program with Enterprise Analytics

{% hint style="info" %}
**Feature availability**\
Enterprise Analytics is available to Snyk Enterprise plan customers. For more information, see [Plans and pricing](https://snyk.io/plans/).

Enterprise Analytics is in Closed Beta. For access, contact your Snyk Account team or Snyk support. During Closed Beta, access at the user level is available. When tenant membership is available, access to Enterprise Analytics will be managed at scale more easily.

Data for all Snyk Groups and Organizations that a user has access to will be displayed in the Enterprise Analytics report during the Closed Beta period. You can use filters to include relevant Groups or Organizations only. This is especially useful for sharing a link with another user.
{% endhint %}

{% hint style="info" %}
Data is refreshed in Enterprise Analytics daily approximately between 13:00 and 14:00 UTC.
{% endhint %}

Enterprise Analytics provides executives as well as Application Security (AppSec) leaders and practitioners a view into the performance of their AppSec program. Snyk customers can understand at a glance the strengths and weaknesses of their program, identify where successful practices can be discerned, and uncover the largest opportunities for improvement that warrant investment.

The following is an example of the Enterprise Analytics dashboard showing AppSec performance for your Groups and Organizations.&#x20;

<figure><img src="https://lh3.googleusercontent.com/7ZCFLTUzQ0r--P354i0p14zVLdE7YjhjvkgODlCfbUho8UMtDUh-EqdFzXmLV8PrvPLIbvE1bDE-qfl1ccYRkNqRksLDhXKr7nvldehocZ89Xa8YQ99nnqt8SmJ-lLTkGO_U05Rl_yrLvXVlOUpmpKg" alt="AppSec performance display"><figcaption><p>AppSec performance dashboard</p></figcaption></figure>

## AppSec performance framework

The AppSec performance dashboard uses the AppSec performance framework to show your performance for each pillar.

Many AppSec teams begin by trying to understand the most serious problems that must be resolved and reacting accordingly. As teams mature, they look for opportunities to proactively improve their programs in each pillar: Exposure, Manage, Prevention, and Coverage.

### Exposure

The Exposure pillar represents a view of the risk to the assets that make up your applications. The focus is on open issues that present real risk to your monitored assets.

### Manage

The Manage pillar represents a view of how your business is managing issues that impact monitored assets that make up your applications. Rather than looking at existing risk, this provides a look at the patterns of behavior of teams in addressing the resolution or acceptance of, or deferral of, action on risk that manifests itself in monitored assets.

### Prevention

The Prevention pillar represents a view of how effectively teams are optimizing the use of tools and processes to stop preventable issues from getting into assets that make up your applications.

### Coverage

The Coverage pillar represents a view of how completely the assets that make up your applications are covered from a security perspective.

## Delineation of how risk is introduced

When you are assessing exposure to risk only, you need not be concerned with how an issue was introduced into the product. When examining an issue or performing an audit, you must be aware that if an issue can be exploited, it may be exploited.

However, when you are evaluating the performance of an AppSec program, you must understand how an issue was introduced so you can stay ahead of attacks and be prepared for the next audit or attack defense.

Issues are categorized as follows.

### Baseline issues

This includes issues identified during the first day a Project began to be monitored. Baseline issues represent new visibility into the risk that existed when Snyk monitoring began.

Baseline issues can be identified when a company begins using a security tool, but can also be found over time. Baseline issues may be identified during the onboarding of a new business unit to Snyk, during the identification of legacy code from a previously onboarded team, or as the result of an acquisition.

While you should treat baseline issues in the same way as other issues from an exposure standpoint, an increase in baseline issues should not raise the same level of alarm as a spike in non-baseline issues.

{% hint style="info" %}
Any issues identified in a Project after the first day of monitoring are considered non-baseline and fall into one of the following categories.
{% endhint %}

### Preventable issues

Preventable issues could have been identified and remediated earlier in the SDLC if developers had taken greater advantage of available shift-left tooling and processes.

Issues can be prevented by developers' taking advantage of Snyk Learn, leveraging the IDE plug-ins, activating PR checks, running `snyk test` locally in the CLI, breaking the build, or taking any other available actions for catching issues pre-production. If Snyk knows about an issue, a test can catch it. Other actions to stop preventable issues from getting into production environments include increasing the threshold for what breaks a build from `critical` to `critical` and `high` severity issues, or more strictly refraining from approving PRs that fail a Snyk test.

Issues are categorized as Preventable if the problem was known to Snyk at least seven days before the detection of the issue. While it is possible that introducing an issue could have been prevented within the seven-day interval, this definition provides some buffer for scenarios where code takes longer to get through the deployment process or for weekly recurring tests.

{% hint style="info" %}
The following example is also a preventable issue.

A developer adds a vulnerable version of a package as part of development. Despite a strong shift-left culture where running the `snyk test` command found this known issue, there was no fix available. The developer opted to use the vulnerable library instead of selecting an alternative because the issue had been known for a long time, or the developer felt the exploit was not relevant to how the package was being used, or both.
{% endhint %}

### Non-preventable issues

Non-preventable issues are the result of an external factor, such as a new vulnerability being published or a new security rule being created, in contrast to developers not shifting left.

There may be a repository that has not been modified in months or years, but because of a newly published vulnerability, such as [log4j](https://snyk.io/blog/log4j-vulnerability-software-supply-chain-security-log4shell/), the asset is now vulnerable. With newly published critical or high-severity open-source vulnerabilities, there may be attention from senior leadership, partners, and customers, requiring the ability to identify issues in this category and measure their trends quickly.

Issues are categorized as Non-preventable if the issue is detected within seven days of becoming known to Snyk. This could include a new vulnerability in a dependency already in use or a vulnerability disclosed in the same time frame as the dependency was introduced to the Project. While it is possible that introducing an issue could have been prevented within the seven-day period, this definition provides some buffer for scenarios where code takes longer to get through the deployment process or for weekly recurring tests.

<figure><img src="https://lh6.googleusercontent.com/ykfbiHzdDVB2X3va4iSzYpSfZ6Ca5yBmeplYMz95wj2Gq6i-xcW4ZZdVOX6Vsl3B1bOaL-gbpTHyksYMjeeAaCHbjAz2QNx3vJ_6h3Oz5ykjqXg2oRWbA5_U-DyoOjGSjSInP_XJv6hIyLsKxmjjFUo" alt="Example of dates an issue was identified, categorized as Non-preventable and Preventable"><figcaption><p>Example of dates an issue was identified, categorized as Non-preventable and Preventable</p></figcaption></figure>

### Other new issues

Not all issues can be easily categorized as Preventable or Non-preventable within Snyk today. The key inputs to this determination are the dates the issue and problem were identified. If either is in question, the issue will be categorized as **Other new**.

Currently, all Snyk Code and Snyk IaC issues will be labeled as **Other new**. Open Source license issues will also be categorized as **Other new**.

## Enterprise Analytics metrics and filters

Enterprise Analytics takes an opinionated approach, focusing teams on a relatively recent view of the most important metrics available for critical and high-severity issues in a 90-day lookback period. These static filters are displayed at the top of the page.

Metrics are organized within the pillars of the AppSec performance framework. It is critical to understand both the current state of metrics and the trend so you can identify anomalies or patterns that contradict expectations.

Enterprise Analytics is designed to make a top-level metric and its associated trend for each pillar visible at all times to support a quick understanding of the state of the program. You can navigate to a more granular view of each tab as needed. Different pillars and metrics on these views may be more or less relevant at different times to different companies, business units, products, teams, and any other participants.&#x20;

Use the tooltips throughout the pages of the Enterprise Analytics application for explanations of metrics and terms.

You can filter Enterprise Analytics on Groups or Organizations manually as needed.



\
