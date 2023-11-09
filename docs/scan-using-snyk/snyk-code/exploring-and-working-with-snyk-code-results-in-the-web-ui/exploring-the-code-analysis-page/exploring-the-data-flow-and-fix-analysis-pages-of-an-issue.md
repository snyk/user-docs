# Data flow

The **Data flow** page shows the exact location of the discovered issue in your source code and how it flows throughout your application. This page displays the taint flow of the issue in the code, with a step-by-step visualization from the source to the sink, presenting the code lines of all the steps in the flow.

{% hint style="info" %}
The source is the input point of the potential problem. This is a point in the application where a user or an external device can enter data, which will potentially violate the security of the application. For example, in a SQL Injection issue, the source will be a form or any other data input area filled by a user.

The sink is the operation in the code where the problem is executed by the application. This point must receive clean input, or it can be exploited. For example, in a SQL Injection issue, the sink will be the internal operation that instructs the DB to perform certain actions according to the received input.
{% endhint %}

For example, in the following **Path Traversal** issue, the developer has not sanitized the input. This allows an attacker to perform a pass traversal attack to access any file in the file system, including sensitive data such as password files:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Issues - Data flow page - Example.png" alt="Path Traversal issue"><figcaption><p>Path Traversal issue</p></figcaption></figure>

Every issue discovered by Snyk Code has a data flow. If an issue has only one step, for example, in the case of hardcoded secrets, the source of the issue will be displayed on the **Data flow** page.

The **Data flow** page enables you to do the following:

* [View the taint flow of an issue in your code from source to sink](exploring-the-data-flow-and-fix-analysis-pages-of-an-issue.md#viewing-the-taint-flow-of-an-issue-in-your-code).
* [Open the tainted source code in the integrated SCM](exploring-the-data-flow-and-fix-analysis-pages-of-an-issue.md#opening-the-tainted-source-code-in-the-integrated-scm).
* Ignore the open vulnerability issue using the **Ignore** button.\
  For more information, see [Ignore issues](https://docs.snyk.io/features/fixing-and-prioritizing-issues/issue-management/ignore-issues).

## **Data flow analysis example**

The **Data flow** page shows the taint flow of an issue in your code from source to sink, including the exact code lines where the taint flow occurs.

To view the exact code lines of the taint flow, select the required step on the left pane. The corresponding code snippet will appear in the right pane:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Issues - Data flow page - Example - Selecting step.png" alt="Code lines of the taint flow"><figcaption><p>Code lines of the taint flow</p></figcaption></figure>

## **Open source code in data flow**

To open the displayed source code on the SCM, click the file name above the right pane:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Issues - Data flow page - Source code link.png" alt="Filename above the rigth pane"><figcaption><p>Filename above the rigth pane</p></figcaption></figure>

The source code appears in the integrated SCM, showing you exactly where to fix the vulnerability. Here, you can make the required fix to address the vulnerability in your code:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Issues - Data flow page - Source code - in SCM.png" alt="HIghlight showoing shere to fix the vulnerability"><figcaption><p>HIghlight showoing shere to fix the vulnerability</p></figcaption></figure>
