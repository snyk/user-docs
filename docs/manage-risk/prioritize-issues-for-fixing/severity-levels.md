# Severity levels

Use severity levels to help you with [vulnerability assessment](https://snyk.io/learn/vulnerability-assessment/) for your applications. Severity levels indicate the assessed level of risk, as **C**ritical, **H**igh, **M**edium, or **L**ow. Snyk reports the number of vulnerabilities at each level of severity in many places in the Snyk application.&#x20;

{% hint style="info" %}
Severity levels also apply to license issues. See [Licenses](../../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/open-source-license-compliance.md).
{% endhint %}

The severity levels are defined in the following table.

| Icon                                                                                                           | Level        | Description                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------- | ------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| <img src="../../.gitbook/assets/image (89).png" alt="C" data-size="line">                                      | **C**ritical | May allow attackers to access sensitive data and run code on your application                                                              |
| <img src="../../.gitbook/assets/image (103) (1) (1) (1) (1) (1) (1) (2) (1).png" alt="H" data-size="original"> | **High**     | May allow attackers to access sensitive data in your application                                                                           |
| ![M](<../../.gitbook/assets/image (17).png>)                                                                   | **M**edium   | Under some conditions, may allow attackers to access sensitive data on your application                                                    |
| ![L](<../../.gitbook/assets/image (60).png>)                                                                   | **L**ow      | Application may expose some data that allows vulnerability mapping, which can be used with other vulnerabilities to attack the application |

## Severity levels and Priority Score

Severity levels are one factor used in determining the Snyk Priority Score for each vulnerability. Other factors include [Snyk Exploit Maturity](https://snyk.io/blog/whats-so-wild-about-exploits-in-the-wild-and-how-can-we-prioritize-accordingly/) and [Reachable Vulnerabilities](https://snyk.io/blog/optimizing-prioritization-with-deep-application-level-context/) information.

See [Snyk Priority Score](priority-score.md) for details.

## How to view severity levels

Severity levels are displayed throughout Snyk, to keep this information visible at all times.

For example, the severity levels appear in the **Pending tasks** section of the Dashboard:

<img src="../../.gitbook/assets/image (36).png" alt="Severity levels with Pending tasks" data-size="original">

Severity levels are displayed in association with your [Snyk Projects](../../snyk-platform-administration/snyk-projects/):

<figure><img src="../../.gitbook/assets/image (43).png" alt="Severity levels assoicated with Projects"><figcaption><p>Severity levels associated with Projects</p></figcaption></figure>

The number of issues at each severity level is also displayed in the left sidebar of an issue card:

<figure><img src="../../.gitbook/assets/image (39) (1).png" alt="Issue card; severity levels in sidebar"><figcaption><p>Issue card; severity levels in sidebar</p></figcaption></figure>

## How Snyk determines severity levels

### Severity levels and CVSS

The Common Vulnerability Scoring System (CVSS) determines the severity level of a vulnerability.

Snyk supports the [CVSS framework version 4.0](https://www.first.org/cvss/v4-0/), along with the previous version, [CVSS framework version 3.1](https://www.first.org/cvss/v3-1/), to designate the characteristics and severity of vulnerabilities.

Vulnerabilities published prior to the support of CVSS v4.0, are based on the 3.1 version of CVSS to define severities.

| **Level** | **CVSS score** |
| --------- | -------------- |
| Critical  | 9.0 - 10.0     |
| High      | 7.0 - 8.9      |
| Medium    | 4.0 - 6.9      |
| Low       | 0.0 - 3.9      |

The severity level and score are determined based on the CVSS Base Score calculations using the [Base Metrics](https://www.first.org/cvss/v4.0/specification-document#Base-Metrics). The Temporal Score, based on the Temporal Metrics, affects the Priority Score.

See [Scoring security vulnerabilities 101: Introducing CVSS for CVEs](https://snyk.io/blog/scoring-security-vulnerabilities-101-introducing-cvss-for-cve/).

{% hint style="info" %}
Severity levels may not always align with CVSS scores. For example, Snyk Container severity scores for Linux vulnerabilities may vary depending on NVD severity rankings; see [Understanding Linux vulnerability severity](../../scan-with-snyk/snyk-container/how-snyk-container-works/severity-levels-of-detected-linux-vulnerabilities.md) for details.
{% endhint %}

### Why are there multiple CVSS Scores for the same vulnerability?

There are multiple CVSS Scores for the same vulnerability for several reasons:

* ​When evaluating the severity of a vulnerability, it is important to note that there is no single CVSS vector. There are multiple CVSS vectors defined by multiple vendors; the [National Vulnerability Database](https://nvd.nist.gov/) (NVD) is one.
* The majority of vulnerabilities published by Snyk originate from [proprietary research](https://security.snyk.io/disclosed-vulnerabilities), public information sources, or through third-party disclosures.
* For example, when Snyk discovered the Critical Severity [Spring4Shell vulnerability](https://security.snyk.io/vuln/SNYK-JAVA-ORGSPRINGFRAMEWORK-2436751), the advisory was published on March 30, 2022, with the CVSS vector analysis. This was before an official CVE was assigned and before NVD conducted its analysis, which was published nine days later on April 8, 2022.
* Having some differences in CVSS vectors is normal and expected. The likelihood of certain attack vectors will involve discrepancies and judgments made about them that make sense for the application and use cases of open source software users.
* The severity of a vulnerability is influenced by a variety of factors, including whether it comes from a "red team" angle or a "blue team" angle. To arrive at an objective and actionable rating, Snyk analysts examine the full range of data, from vendors to reporters to attackers.
* There are times when a vendor discovers additional information about a vulnerability that can affect its severity. Users can find all the relevant information used to determine the severity that Snyk curated in the description and references of the advisory.
*   Different vendors may use different versions of CVSS, resulting in differing scores and severity levels. For instance, Snyk adopted CVSS version 4.0 as its primary framework in 2024. When comparing a vulnerability published by Snyk using [CVSS v4.0 ](https://www.first.org/cvss/v4-0/)with one published by the NVD using CVSS v3.1, the scores and severity ratings may differ significantly. In addition, CVSS score calculation can be done based on different vectors:

    * Base (CVSS-B),
    * Base + Threat (CVSS-BT),
    * Base + Environmental (CVSS-BE),
    * Base + Threat + Environmental (CVSS-BTE).

    Snyk performs the calculation of the score and the severity based on the Base metrics (as advised in compliance frameworks like FedRAMP and PCI-DSS) and provides the Threat Metric (also known as Exploit Maturity) as a separate data point.

### Severity levels and CCSS

The Common Configuration Scoring System (CCSS), developed by the National Institute of Standards and Technology (NIST) and derived from CVSS, measures the severity of software security configuration issues.

Snyk uses the [CCSS](https://www.nist.gov/publications/common-configuration-scoring-system-ccss-metrics-software-security-configuration) to designate the characteristics and severity of IaC+ vulnerabilities and misconfigurations.

| **Level** | **CCSS score** |
| --------- | -------------- |
| Critical  | 9.0 - 10.0     |
| High      | 7.0 - 8.9      |
| Medium    | 4.0 - 6.9      |
| Low       | 0.0 - 3.9      |

The severity level and score are determined based on the CCSS Base Score calculations using the Base Metrics.
