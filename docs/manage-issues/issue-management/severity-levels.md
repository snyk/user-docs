# Severity levels

Use severity levels to help you with [vulnerability assessment](https://snyk.io/learn/vulnerability-assessment/) for your applications.

<img src="../../.gitbook/assets/Screenshot 2022-08-16 at 09.52.22.png" alt="" data-size="original">

* [Introduction to Snyk severity levels](severity-levels.md#introduction-to-snyk-severity-levels)
* [Viewing severity levels](severity-levels.md#viewing-severity-levels)
* [Determining severity levels](severity-levels.md#determining-severity-levels)

## Introduction to Snyk severity levels

Severity levels indicate the assessed level of risk, as one of **C**ritical / **H**igh / **M**edium / **L**ow:

| Icon                                                                                                      | Level        | Description                                                                                                                                |
| --------------------------------------------------------------------------------------------------------- | ------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| <img src="../../.gitbook/assets/image (131) (1) (1) (1).png" alt="" data-size="line">                     | **Critical** | This may allow attackers to access sensitive data and run code on your application                                                         |
| <img src="../../.gitbook/assets/image (103) (1) (1) (1) (1) (1) (1) (2).png" alt="" data-size="original"> | **High**     | This may allow attackers to access sensitive data in your application                                                                      |
| ![](<../../.gitbook/assets/image (133) (1).png>)                                                          | **Medium**   | Under some conditions, this may allow attackers to access sensitive data on your application                                               |
| ![](<../../.gitbook/assets/image (422).png>)                                                              | **Low**      | Application may expose some data that allows vulnerability mapping, which can be used with other vulnerabilities to attack the application |

{% hint style="info" %}
Severity levels also apply to license issues. See [Licenses](../../scan-application-code/snyk-open-source/licenses/).
{% endhint %}

## Severity levels and Priority Scores

Severity levels are one factor feeding into Snyk's Priority Score for each vulnerability, along with factors such as [Snyk’s Exploit Maturity](https://snyk.io/blog/whats-so-wild-about-exploits-in-the-wild-and-how-can-we-prioritize-accordingly/) and [Reachable Vulnerabilities](https://snyk.io/blog/optimizing-prioritization-with-deep-application-level-context/) information.

See [Snyk Priority Score](../prioritizing-issues/priority-score.md) for details.

## Viewing severity levels

Severity levels are displayed throughout Snyk, to show this information at all times.

For example, in the **Pending tasks** section of the Dashboard:

<img src="../../.gitbook/assets/image (158) (1) (1) (1) (1) (1) (1) (1).png" alt="" data-size="original">

Associated with your [Snyk Projects](../snyk-projects/):

<figure><img src="../../.gitbook/assets/image (43) (2).png" alt="Severity levels in Projects"><figcaption><p>Severity levels in Projects</p></figcaption></figure>

And for each vulnerability in a project:

<figure><img src="../../.gitbook/assets/image (39) (1).png" alt="Severity levels in vulnerabilities"><figcaption><p>Severity levels in vulnerabilities</p></figcaption></figure>

## Determining severity levels

### Severity levels and CVSS

The **Common Vulnerability Scoring System** (**CVSS**) determines the severity level of a vulnerability.

At Snyk, we use [CVSS framework version 3.1](https://www.first.org/cvss/v3-1/) to communicate the characteristics and severity of vulnerabilities.

| **Level** | **CVSS score** |
| --------- | -------------- |
| Critical  | 9.0 - 10.0     |
| High      | 7.0 - 8.9      |
| Medium    | 4.0 - 6.9      |
| Low       | 0.0 - 3.9      |

The severity level and score are determined based on the [CVSS Base Score](https://www.first.org/cvss/specification-document) calculations using the Base Metrics. The Temporal Score, based on the Temporal Metrics, affects the Priority Score.

See [Scoring security vulnerabilities 101: Introducing CVSS for CVEs](https://snyk.io/blog/scoring-security-vulnerabilities-101-introducing-cvss-for-cve/).

{% hint style="info" %}
Severity levels may not always align with CVSS scores. For example, Snyk Container severity scores for Linux vulnerabilities may vary depending on NVD severity rankings; see [Understanding Linux vulnerability severity](../../scan-containers/how-snyk-container-works/understanding-linux-vulnerability-severity.md) for more details.
{% endhint %}

### **Why are there multiple CVSS Scores for the same vulnerability?**

* ​When evaluating the severity of a vulnerability, it's important to note that there is no single CVSS vector - there are multiple CVSS vectors defined by multiple vendors, with the [National Vulnerability Database](https://nvd.nist.gov/) (NVD) being one of them.
* The majority of vulnerabilities published by Snyk originate from [proprietary research](https://security.snyk.io/disclosed-vulnerabilities), public information sources, or through 3rd party disclosures.
* For example, when Snyk discovered the Critical Severity [Spring4Shell vulnerability](https://security.snyk.io/vuln/SNYK-JAVA-ORGSPRINGFRAMEWORK-2436751), the advisory published on March 30th, 2022, with the CVSS vector analysis, before an official CVE was assigned, and before NVD conducted their analysis, which was published 9 days later, on April 8th, 2022.
* Having some differences in CVSS vectors is normal and expected. The likelihood of certain attack vectors will raise discrepancies, and judgments will need to be made about them in a way that makes sense for the application and use cases of open-source software users.
* A vulnerability's severity is influenced by a variety of factors, including whether it comes from a "red team" angle or a "blue team" angle. To arrive at an objective and actionable rating, Snyk analysts examine the full range of data - from vendors to reporters to attackers.
* There are times when a vendor discovers additional information about a vulnerability that can affect its severity. Users can find all the relevant information used to determine the severity that Snyk curated in the advisory's description and references.
