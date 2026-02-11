# Overview tab

Analytics provides executives, as well as Application Security (AppSec) leaders and practitioners, a view into the performance of their AppSec program. Snyk customers can understand at a glance the strengths and weaknesses of their program, identify where successful practices can be discerned, and uncover the largest opportunities for improvement that warrant investment.&#x20;

Analytics data is organized in multiple sections, allowing you to focus on specific metrics, from key performance indicators to coverage, exposure, remediation, or prevention. You can select and remove metrics from each section to customize your view.

## Key Performance Indicators section

{% hint style="info" %}
To customize the **Key Performance Indicators** section, click **Edit section** and select from the dropdown the indicators to display.
{% endhint %}

* **Open issues (Featured zero-day):** tracks open issues related to featured zero-day vulnerabilities to highlight exposure to critical emerging threats.
* **Ignored issues**: tracks issues that your team has intentionally ignored or "muted' instead of fixing them.
* **New issues**: represents the volume of freshly identified vulnerabilities or tickets discovered in the selected timeframe.
* **Resolved issues**: tracks the number of vulnerabilities or tickets that were successfully closed or fixed in the selected timeframe.
* **Unique vulnerabilities:** shows the distinct count of unresolved vulnerabilities to highlight the breadth of different security problems.
* **Tested repositories in CI/CD:** tracks the number of repositories tested in the continuous integration/continuous delivery (CI/CD) pipeline.

## Coverage section

* **Repository coverage by Snyk product**: Indicates the percentage of repositories covered by each Snyk product.
* **Repository coverage percentage over time, by Snyk product**: Shows the trend of repository coverage by Snyk products over time.
* **Assets discovered over time, by type**: Tracks the trend of assets discovered, categorized by their type.

## Exposure section

{% hint style="info" %}
To customize the **Exposure** section, click **Edit section** and select from the dropdown the indicators to display.
{% endhint %}

* **Organizations with the most open issues**: identifies high-risk areas to help you focus remediation resources.
* **Open issues over time, by severity**: displays the trend of open issues by severity.
* **Open issues by Snyk product**: lists open issues by the identified Snyk product.
* **Open issues by severity**: categorizes open issues by severity level.
* **Open issues by repository freshness**: identifies risks based on the last commit (freshness) to highlight blind spots. You can configure the definitions for:
  * Active (default: less than 3 months)
  * Inactive (default: 3–6 months)
  * Dormant (default: over 6 months)
  * N/A indicates that the Group SCM integration detected no commits
* **Open issues (OWASP Top 10, 2021) by severity**: groups open issues by OWASP Top 10 (2021) category and severity.
* **Open issues by asset class**: enables prioritization of remediation efforts based on the business criticality of different assets.
* **SCA CVEs with the most open issues**: highlights high-impact open-source vulnerabilities (CVEs) with the most open issues identified through Software Composition Analysis (SCA).

## Remediation section&#x20;

* **Issues identified and resolved over time**: assesses your ability to address new vulnerabilities and reduce the security backlog. Use this metric to prioritize remediation.
* **Resolved issues by MTTR duration**: displays typical resolution timeframes and indicates remediation efficiency. Mean time to resolution (MTTR) is the average time required to fix an issue.
* **Organizations with the highest MTTR**: identifies Organizations with the longest average fix times. High MTTR values highlight teams that need improved processes or resources.
* **Snyk-generated PRs merged vs. open**: compares the number of merged Snyk-generated pull requests to the number of open pull requests.

## Prevention section

* **Snyk IDE & CLI usage over time (developer count)**: shows the number of developers using Snyk IDE plugins and the CLI for local development.
* **Organizations introducing most new SCA preventable issues**: identifies organizations with the highest rate of new preventable Software Composition Analysis (SCA) issues. Use this data to prioritize prevention efforts.
* **New issues by introduction category**: identifies the primary sources of new vulnerabilities to help you strengthen prevention strategies.
* **New SCA preventable issues introduced over time**: displays the number of new open-source vulnerabilities you could have avoided. This metric helps you assess the effectiveness of prevention strategies for open-source risks.

