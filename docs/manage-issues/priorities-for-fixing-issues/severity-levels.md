# Severity levels

Use severity levels to help you with [vulnerability assessment](https://snyk.io/learn/vulnerability-assessment/) for your applications. Severity levels indicate the assessed level of risk, as **C**ritical, **H**igh, **M**edium, or **L**ow. Snyk reports the number of vulnerabilities at each level of severity in many places in the Snyk application. The display varies; a typical example follows.

<img src="../../.gitbook/assets/Screenshot 2022-08-16 at 09.52.22.png" alt="Issues at each level of severity, C, H, M, and L" data-size="original">

{% hint style="info" %}
Severity levels also apply to license issues. See [Licenses](../../scan-application-code/snyk-open-source/licenses/).
{% endhint %}

The severity levels are defined in the following table.

| Icon                                                                                                       | Level        | Description                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| <img src="../../.gitbook/assets/image (131) (1) (1) (1).png" alt="C" data-size="line">                     | **C**ritical | May allow attackers to access sensitive data and run code on your application                                                              |
| <img src="../../.gitbook/assets/image (103) (1) (1) (1) (1) (1) (1) (2).png" alt="H" data-size="original"> | **High**     | May allow attackers to access sensitive data in your application                                                                           |
| ![M](<../../.gitbook/assets/image (133) (1).png>)                                                          | **M**edium   | Under some conditions, may allow attackers to access sensitive data on your application                                                    |
| ![L](<../../.gitbook/assets/image (422).png>)                                                              | **L**ow      | Application may expose some data that allows vulnerability mapping, which can be used with other vulnerabilities to attack the application |

## Severity levels and Priority Score

Severity levels are one factor used in determining the Snyk Priority Score for each vulnerability. Other factors include [Snyk Exploit Maturity](https://snyk.io/blog/whats-so-wild-about-exploits-in-the-wild-and-how-can-we-prioritize-accordingly/) and [Reachable Vulnerabilities](https://snyk.io/blog/optimizing-prioritization-with-deep-application-level-context/) information.

See [Snyk Priority Score](priority-score.md) for details.

## How to view severity levels

Severity levels are displayed throughout Snyk, to keep this information visible at all times.

For example, the severity levels appear in the **Pending tasks** section of the Dashboard:

<img src="../../.gitbook/assets/image (158) (1) (1) (1) (1) (1) (1) (1).png" alt="Severity levels with Pending tasks" data-size="original">

Severity levels are displayed in association with your [Snyk Projects](../snyk-projects/):

<figure><img src="../../.gitbook/assets/image (43) (2).png" alt="Severity levels assoicated with Projects"><figcaption><p>Severity levels associated with Projects</p></figcaption></figure>

The number of issues at each severity level is also displayed in the left sidebar of an issue card:

<figure><img src="../../.gitbook/assets/image (39) (1).png" alt="Issue card; severity levels in sidebar"><figcaption><p>Issue card; severity levels in sidebar</p></figcaption></figure>

## How Snyk determines severity levels

### Severity levels and CVSS

The Common Vulnerability Scoring System (CVSS) determines the severity level of a vulnerability.

Snyk uses the [CVSS framework version 3.1](https://www.first.org/cvss/v3-1/) to designate the characteristics and severity of vulnerabilities.

| **Level** | **CVSS score** |
| --------- | -------------- |
| Critical  | 9.0 - 10.0     |
| High      | 7.0 - 8.9      |
| Medium    | 4.0 - 6.9      |
| Low       | 0.0 - 3.9      |

The severity level and score are determined based on the [CVSS Base Score](https://www.first.org/cvss/specification-document) calculations using the Base Metrics. The Temporal Score, based on the Temporal Metrics, affects the Priority Score.

See [Scoring security vulnerabilities 101: Introducing CVSS for CVEs](https://snyk.io/blog/scoring-security-vulnerabilities-101-introducing-cvss-for-cve/).

{% hint style="info" %}
Severity levels may not always align with CVSS scores. For example, Snyk Container severity scores for Linux vulnerabilities may vary depending on NVD severity rankings; see [Understanding Linux vulnerability severity](../../scan-containers/how-snyk-container-works/understanding-linux-vulnerability-severity.md) for details.
{% endhint %}

### **Why are there multiple CVSS Scores for the same vulnerability?**

There are multiple CVSS Scores for the same vulnerability for several reasons:

* ​When evaluating the severity of a vulnerability, it is important to note that there is no single CVSS vector. There are multiple CVSS vectors defined by multiple vendors; the [National Vulnerability Database](https://nvd.nist.gov/) (NVD) is one.
* The majority of vulnerabilities published by Snyk originate from [proprietary research](https://security.snyk.io/disclosed-vulnerabilities), public information sources, or through third-party disclosures.\
  For example, when Snyk discovered the Critical Severity [Spring4Shell vulnerability](https://security.snyk.io/vuln/SNYK-JAVA-ORGSPRINGFRAMEWORK-2436751), the advisory was published on March 30, 2022, with the CVSS vector analysis. This was before an official CVE was assigned and before NVD conducted its analysis, which was published nine days later on April 8, 2022.
* Having some differences in CVSS vectors is normal and expected. The likelihood of certain attack vectors will involve discrepancies and judgments made about them that make sense for the application and use cases of open source software users.
* The severity of a vulnerability is influenced by a variety of factors, including whether it comes from a "red team" angle or a "blue team" angle. To arrive at an objective and actionable rating, Snyk analysts examine the full range of data, from vendors to reporters to attackers.
* There are times when a vendor discovers additional information about a vulnerability that can affect its severity. Users can find all the relevant information used to determine the severity that Snyk curated in the description and references of the advisory.
