# Severity levels

A severity level is applied to a vulnerability, to indicate the risk for that vulnerability in an application.

<img src="../../../.gitbook/assets/Screenshot 2022-08-16 at 09.52.22.png" alt="" data-size="original">

Severity levels are key factors in [vulnerability assessment](https://snyk.io/learn/vulnerability-assessment/), and can be:

* ![](<../../../.gitbook/assets/image (86) (1).png>) **Low:** application may expose some data allowing vulnerability mapping, which can be used with other vulnerabilities to attack the application.
* ![](<../../../.gitbook/assets/image (84) (1) (1).png>) **Medium:** may allow attackers under some conditions to access sensitive data on your application.
* <img src="../../../.gitbook/assets/image (131) (1) (1) (1) (1).png" alt="" data-size="original"> **High:** may allow attackers to access sensitive data on your application.
* <img src="../../../.gitbook/assets/image (83) (1) (2).png" alt="" data-size="line">**Critical:** may allow attackers to access sensitive data and run code on your application.

{% hint style="info" %}
Severity levels also apply to license issues. See [Licenses overview](https://docs.snyk.io/snyk-open-source/licenses).
{% endhint %}

### Determining severity levels

The **Common Vulnerability Scoring System** (**CVSS**) determines the severity level of a vulnerability.

At Snyk, we use [CVSS framework version 3.1](https://www.first.org/cvss/v3-1/) to communicate the characteristics and severity of vulnerabilities.

| **Severity level** | **CVSS score** |
| ------------------ | -------------- |
| Low                | 0.0 - 3.9      |
| Medium             | 4.0 - 6.9      |
| High               | 7.0 - 8.9      |
| Critical           | 9.0 - 10.0     |

The severity level and score are determined based on the CVSS Base Score calculations using the [Base Metrics](https://www.first.org/cvss/specification-document#Temporal-Metrics). The Temporal Score, which is based on the [Temporal Metrics](https://www.first.org/cvss/specification-document#Temporal-Metrics) is affecting the Priority Score.

See [Scoring security vulnerabilities 101: Introducing CVSS for CVEs](https://snyk.io/blog/scoring-security-vulnerabilities-101-introducing-cvss-for-cve/).

{% hint style="info" %}
Severity levels may not always align with CVSS scores. For example, Snyk Container severity scores for Linux vulnerabilities may vary depending on NVD severity rankings; see [Understanding Linux vulnerability severity](../../../products/snyk-container/getting-started-snyk-container/understanding-linux-vulnerability-severity.md) for more details.
{% endhint %}

### Understanding Snyk's Vulnerability Analysis

**Why are there multiple CVSS Scores for the same vulnerability?**

* ​When evaluating the severity of a vulnerability, it's important to note that there is no single CVSS vector - there are multiple CVSS vectors defined by multiple vendors, with the [National Vulnerability Database](https://nvd.nist.gov/) (NVD) being one of them.
* The majority of vulnerabilities published by Snyk originate from [proprietary research](https://security.snyk.io/disclosed-vulnerabilities), public information sources, or through 3rd party disclosures.
* For example, when Snyk discovered the Critical Severity [Spring4Shell vulnerability](https://security.snyk.io/vuln/SNYK-JAVA-ORGSPRINGFRAMEWORK-2436751), the advisory published on March 30th, 2022, with the CVSS vector analysis, before an official CVE was assigned, and before NVD conducted their analysis, which was published 9 days later, on April 8th, 2022.
* Having some differences in CVSS vectors is normal and expected. The likelihood of certain attack vectors will raise discrepancies, and judgments will need to be made about them in a way that makes sense for the application and use cases of open-source software users.
* A vulnerability's severity is influenced by a variety of factors, including whether it comes from a "red team" angle or a "blue team" angle. To arrive at an objective and actionable rating, Snyk analysts examine the full range of data - from vendors to reporters to attackers.
* There are times when a vendor discovers additional information about a vulnerability that can affect its severity. Users can find all the relevant information used to determine the severity that Snyk curated in the advisory's description and references.

### Severity and priority scoring

Severity levels are one factor feeding into Snyk's Priority Score for each vulnerability, along with factors such as [Snyk’s Exploit Maturity](https://snyk.io/blog/whats-so-wild-about-exploits-in-the-wild-and-how-can-we-prioritize-accordingly/) and [Reachable Vulnerabilities](https://snyk.io/blog/optimizing-prioritization-with-deep-application-level-context/) information. Together, this scoring helps developers determine which vulnerabilities should be addressed first.

See [Snyk Priority Score](priority-score.md) for details of how severity levels are used in Snyk's priority scores.

## Viewing severity levels in Snyk

Severity levels are displayed throughout Snyk, to show this information at all times.

For example, in the **Pending tasks** section of the Dashboard:

<img src="../../../.gitbook/assets/image (159) (1).png" alt="" data-size="original">

Associated with your [Projects](../../../introducing-snyk/projects.md):

![](<../../../.gitbook/assets/image (43) (2).png>)

And for each vulnerability in a project:

![](<../../../.gitbook/assets/image (39) (1).png>)
