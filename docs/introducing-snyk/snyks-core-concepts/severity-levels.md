# Severity levels

A severity level is applied to a vulnerability, to indicate the risk for that vulnerability in an application. Severity levels are key factors in [vulnerability assessment](https://snyk.io/learn/vulnerability-assessment/), and can be:

* **Low:** the application may expose some data allowing vulnerability mapping, which can be used with other vulnerabilities to attack the application. 
* **Medium:** may allow attackers under some conditions to access sensitive data on your application.
* **High:** may allow attackers to access sensitive data on your application.
* **Critical:** may allow attackers to access sensitive data and run code on your application.

{% hint style="info" %}
Severity levels also apply to license issues. See [Licenses overview](snyk-open-source/licenses).
{% endhint %}

### Determining severity levels

The **Common Vulnerability Scoring System** \(**CVSS**\) determines the severity level of a vulnerability.

At Snyk, we use [CVSS framework version 3.1](https://www.first.org/cvss/v3-1/) to communicate the characteristics and severity of vulnerabilities.

| **Severity level** | **CVSS score** |
| :--- | :--- |
| Low | 0.0 - 3.9 |
| Medium | 4.0 - 6.9 |
| High | 7.0 - 8.9 |
| Critical | 9.0 - 10.10 |

See [Scoring security vulnerabilities 101: Introducing CVSS for CVEs](https://snyk.io/blog/scoring-security-vulnerabilities-101-introducing-cvss-for-cve/).

### Severity and priority scoring

Severity levels are one factor feeding into Snyk's Priority Score for each vulnerability, along with factors such as [Snyk’s Exploit Maturity](https://snyk.io/blog/whats-so-wild-about-exploits-in-the-wild-and-how-can-we-prioritize-accordingly/) and [Reachable Vulnerabilities](https://snyk.io/blog/optimizing-prioritization-with-deep-application-level-context/) information. Together, this scoring helps developers determine which vulnerabilities should be addressed first.

See [Snyk Priority Score](../../fixing-and-prioritizing-issues/starting-to-fix-vulnerabilities/snyk-priority-score.md) for details of how severity levels are used in Snyk's priority scores.

## Viewing severity levels in Snyk

Severity levels are displayed throughout Snyk, to show this information at all times.

For example, in the initial dashboard:

![](../../.gitbook/assets/image%20%2846%29.png)

For your projects:

![](../../.gitbook/assets/image%20%2843%29.png)

And for each vulnerability in a project:

![](../../.gitbook/assets/image%20%2839%29.png)

See [Getting started](getting-started) documentation for more details of using Snyk.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

