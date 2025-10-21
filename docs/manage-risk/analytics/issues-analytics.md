# Issues analytics

{% hint style="info" %}
**Release status**

Issues Analytics is only available for customers under Enterprise plans.&#x20;
{% endhint %}

Issues Analytics takes an opinionated approach, focusing teams on a relatively recent view of the most important metrics available for critical and high-severity issues in a 90-day lookback period. These static filters are displayed at the top of the page.

The Issues Analytics are focused on metrics for the following products:

* Snyk Open Source
* Snyk Code
* Snyk Container
* Snyk IaC

You can select the Groups and Organizations that will be included in the analysis.

Metrics are organized within the pillars of the AppSec performance framework. It is critical to understand both the current state of metrics and the trend to identify anomalies or patterns that contradict expectations.

Issues Analytics is designed to make a top-level metric and its associated trend for each pillar visible at all times to support a quick understanding of the state of the program. You can navigate to a more granular view of each tab as needed. Different pillars and metrics on these views may be more or less relevant at different times to different companies, business units, products, teams, and any other participants.&#x20;

Use the tooltips throughout the pages of the Issues Analytics application for explanations of metrics and terms.

You can filter Issues Analytics on Groups or Organizations manually, as needed.

{% hint style="info" %}
Data is refreshed in Issues Analytics daily, approximately between 13:00 and 14:00 UTC.
{% endhint %}

## AppSec performance framework

The AppSec performance dashboard uses the AppSec performance framework to show your performance for each pillar.

Many AppSec teams begin by trying to understand the most serious problems that must be resolved and reacting accordingly. As teams mature, they look for opportunities to proactively improve their programs in each pillar: Exposure, Manage, Prevention, and Coverage.

### Exposure

The Exposure pillar represents a view of the risk to the assets that make up your applications. The focus is on open issues that present real risk to your monitored assets.

### Manage

The Manage pillar represents how your business manages issues impacting monitored assets that make up your applications. Rather than looking at existing risk, this provides a look at the patterns of behavior of teams in addressing the resolution or acceptance of, or deferral of, action on risk that manifests itself in monitored assets.

### Prevention

The Prevention pillar represents a view of how effectively teams are optimizing the use of tools and processes to stop preventable issues from getting into assets that make up your applications.

### Coverage

The Coverage pillar represents a view of how completely the assets that make up your applications are covered from a security perspective.

## Delineation of how risk is introduced

When you are assessing exposure to risk only, you need not be concerned with how an issue was introduced into the product. When examining an issue or performing an audit, you must be aware that if an issue can be exploited, it may be exploited.

However, when you are evaluating the performance of an AppSec program, you must understand how an issue was introduced so you can stay ahead of attacks and be prepared for the next audit or attack defense.

Issues are categorized as follows:

### Baseline issues

This includes issues identified during the first day a Project began to be monitored. Baseline issues represent new visibility into the risk that existed when Snyk monitoring began.

Baseline issues can be identified when a company begins using a security tool, but can also be found over time. Baseline issues may be identified during the onboarding of a new business unit to Snyk, during the identification of legacy code from a previously onboarded team, or as the result of an acquisition.

While you should treat baseline issues in the same way as other issues from an exposure standpoint, an increase in baseline issues should not raise the same level of alarm as a spike in non-baseline issues.

{% hint style="info" %}
Any issues identified in a Project after the first day of monitoring are considered non-baseline and fall into one of the following categories.
{% endhint %}

### Preventable issues

Preventable issues could have been identified and remediated earlier in the SDLC if developers had taken greater advantage of available shift-left tooling and processes.

Issues can be prevented by developers taking advantage of Snyk Learn, leveraging the IDE plug-ins, activating PR checks, running `snyk test` locally in the CLI, breaking the build, or taking any other available actions for catching issues pre-production. If Snyk knows about an issue, a test can catch it. You can stop preventable issues from getting into production environments by increasing the threshold for what breaks a build. You can increase the threshold from only `critical` issues, to `critical` and `high` issues, or strictly refrain from approving PRs that fail a Snyk test.

Issues are categorized as Preventable if the problem was known to Snyk at least seven days before the detection of the issue. While it is possible that introducing an issue could have been prevented within the seven-day interval, this definition provides some buffer for scenarios where code takes longer to get through the deployment process or for weekly recurring tests.

{% hint style="success" %}
The following example is also a preventable issue.

A developer adds a vulnerable version of a package as part of development. Despite a strong shift-left culture where running the `snyk test` command found this known issue, there was no fix available. The developer opted to use the vulnerable library instead of selecting an alternative because the issue had been known for a long time, or the developer felt the exploit was not relevant to how the package was being used, or both.
{% endhint %}

### Non-preventable issues

Non-preventable issues are the result of an external factor, such as a new vulnerability being published or a new security rule being created, in contrast to developers not shifting left.

There may be a repository that has not been modified in months or years, but because of a newly published vulnerability, such as [log4j](https://snyk.io/blog/log4j-vulnerability-software-supply-chain-security-log4shell/), the asset is now vulnerable. With newly published critical or high-severity open-source vulnerabilities, there may be attention from senior leadership, partners, and customers, requiring the ability to identify issues in this category and measure their trends quickly.

Issues are categorized as Non-preventable if the issue is detected within seven days of becoming known to Snyk. This could include a new vulnerability in a dependency already in use or a vulnerability disclosed in the same time frame as the dependency was introduced to the Project. While it is possible that introducing an issue could have been prevented within the seven-day period, this definition provides some buffer for scenarios where code takes longer to get through the deployment process or for weekly recurring tests.

<figure><img src="../../.gitbook/assets/image (161).png" alt=""><figcaption></figcaption></figure>

### Other new issues

Not all issues can be easily categorized as Preventable or Non-preventable within Snyk today. The key inputs to this determination are the dates the issue and problem were identified. If either is in question, the issue will be categorized as **Other new**.

All Snyk Code and Snyk IaC issues will be labeled as **Other new**. Open Source license issues will also be categorized as **Other new**.

### Preventability analysis considerations

Snyk can confidently differentiate whether issues are preventable by measuring their preventability. However, this mechanism includes some assumptions in its calculation.&#x20;

First assumption is that any non-baseline issues detected within seven days of the underlying vulnerability being known to Snyk are non-preventable. In this scenario, a developer ran a test locally or in a pipeline, which detected the issue, and they deployed it right away. The seven day threshold provides reasonable flexibility to account for this scenario.&#x20;

The second assumption is that vulnerabilities may be captured within the Snyk Vulnerability Database before the SCA engine can detect them in scans. Snyk attempts to synchronize these changes closely, but in some cases, engine updates are made after the vulnerability is published in the Snyk Vulnerability Database.
