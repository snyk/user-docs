# Vulnerabilities

Each vulnerability issue discovered by Snyk Code contains the following details:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Issues - diagram.png" alt="Vulnerability issues discovered by Snyk Code"><figcaption><p>Vulnerability issues discovered by Snyk Code</p></figcaption></figure>

The following explains each of the areas depicted in a box in the image:

* **Vulnerability type** - the type of security vulnerability or security rule discovered in the source code, in this example, **SQL Injection**.\
  For the full list of vulnerability types and security rules that Snyk Code applies to your source code, see [Security rules used by Snyk Code](../../../../scan-application-code/snyk-code/exploring-and-working-with-snyk-code-results-in-the-web-ui/broken-reference/).
* **CWE reference** - the CWE (Common Weakness Enumeration) ID of the specific vulnerability type and a link to the CWE website, where this vulnerability type is described:
* **Priority score** - a number calculated and assigned by Snyk Code to each discovered issue based on the severity level, risk, frequency, and ease of fix of the issue.\
  For more information, see [Understanding the priority score of the Snyk Code issues](priority-score-in-snyk-code.md).
* **Severity level** - can be one of the following: **H**igh, **M**edium, or **L**ow.\
  The Severity level of an issue in Snyk Code analysis is defined mainly by the risk of the sink of the potential vulnerability.\
  For more information about the source and sink of an issue, see [Exploring the Data flow page](broken-reference).

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Issues - CWE link.png" alt="CWE link"><figcaption><p>CWE link</p></figcaption></figure>

By clicking the **CWE** link, you can open the CWE web page that provides additional details about the vulnerability type of the discovered issue:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Issues - CWE web page.png" alt="CWE page for a vulnerability type"><figcaption><p>CWE page for a vulnerability type</p></figcaption></figure>

The explanation continues for each area depicted in a box in the image at the beginning of the page:

* **Code snippet** - the sink area in the analyzed code where the discovered vulnerability may be executed. To view the entire code snippet, click the **Full details** button to open the **Data flow** page.
* **Vulnerability description** - an overview of the discovered vulnerability. To view the full description of the vulnerability, click the **Full details** button to open the **Data flow** page.
* **Link to the vulnerable file in the SCM** - a link to the tainted source code in the integrated SCM.

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Issues - Repo link.png" alt="Link to the vulnerable file"><figcaption><p>Link to the vulnerable file</p></figcaption></figure>

Click the link to open the source code file that includes the sink area of the issue. The vulnerable sink area is highlighted in the code:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Issues - Repo link - in SCM.png" alt="Highlighted vulnerable sink area"><figcaption><p>Highlighted vulnerable sink area</p></figcaption></figure>

The explanation continues for each area depicted in a box in the image at the beginning of the page:

* **Snyk Learn link** - if available, a link to [Snyk Learn](https://learn.snyk.io/) for interactive lessons on understanding, fixing, and avoiding the discovered vulnerability.
* **Ignore** button - enables you to ignore this vulnerability issue.\
  For more information, see [Ignore issue](../../../find-and-manage-priority-issues/ignore-issues/)s.
* **Full details** button - enables you to open the [Data flow](broken-reference) and [Fix analysis](fix-analysis.md) pages for more information on the discovered issue and fix recommendations and examples.

## Vulnerabilities in Code Analysis

The Vulnerability Issues area on the **Code Analysis** page displays all the issues that were discovered in the imported repository by Snyk Code. This issue list is organized by the severity level of the discovered issues, by default from the highest to the lowest:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Issues Area .png" alt="Vulnerabiltiy issues on Code Analysis page"><figcaption><p>Vulnerabiltiy issues on Code Analysis page</p></figcaption></figure>

You can change the display of the issues on the **Code Analysis** page using the following options:

* [Grouping](vulnerabilities.md#grouping-the-discovered-issues-according-to-their-file-or-vulnerability-type) - by **File** or **Vulnerability Type**.
* [Sorting](vulnerabilities.md#sorting-the-discovered-issues-according-to-their-severity-level) - from the **Highest** or the **Lowest** severity level.
* [Filtering](vulnerabilities.md#filtering-the-discovered-issues-according-to-pre-defined-criteria-or-a-search-word) - according to pre-defined criteria or a search word.

{% hint style="info" %}
For more information on the available details and options for each discovered issue, see [Exploring the vulnerability issues discovered by Snyk Code](broken-reference).
{% endhint %}

## **Group vulnerabilities**

You can group the discovered vulnerabilities issues according to the **file** or **vulnerability type** criterion, using the **Group by** dropdown list. This grouping option can help you identify problematic files with multiple issues or address prevalent vulnerability types:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Main UI Features - Group.png" alt="Group by option on Code Analysis page"><figcaption><p>Group by option on Code Analysis page</p></figcaption></figure>

## **Sort by severity level**

You can sort the discovered vulnerabilities issues according to their severity level, **highest** or **lowest,** using the **Sort by** dropdown list:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Main UI Features - Sort.png" alt="ort by option on Code Analysis page"><figcaption><p>Sort by option on Code Analysis page</p></figcaption></figure>

## **Filter vulnerabilities**

You can filter the discovered vulnerabilities issues by the following:

* Pre-defined criteria available in the filter pane - **Severity**, **Priority Score**, **Status**, **Languages**, and **Vulnerability Types**.\
  For more information, see [The Code Analysis page - the Filter pane](broken-reference).
* A searched word entered in the **Search** box:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Issues Area - Search box.png" alt="earch box on Code Analysis page"><figcaption><p>Search box on Code Analysis page</p></figcaption></figure>
