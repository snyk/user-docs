# Understand your vulnerabilities

{% hint style="info" %}
**Recap**\
You have [viewed and understood scanned Projects](view-your-first-snyk-projects.md); now you can look at the details of vulnerabilities in that Project.
{% endhint %}

## **See vulnerability details**

First, open a target to see your Snyk Projects:

<figure><img src="../../.gitbook/assets/image (354).png" alt="View imported Projects"><figcaption><p>View imported Projects</p></figcaption></figure>

Next, select a Project in that list, to see details of the vulnerabilities found in that Project.

For example, for a **Code analysis** project scanned by Snyk Code:

<figure><img src="../../.gitbook/assets/image (149) (1) (1) (2) (2).png" alt="Vulnerability example - Code analysis"><figcaption><p>Vulnerability example - Code analysis</p></figcaption></figure>

See [View project information](../../snyk-admin/snyk-projects/project-information.md) for more details.

## View Issue Cards

Now, look at the vulnerability information for each Snyk Project, provided in Issue Cards:

<figure><img src="../../.gitbook/assets/image (101) (2).png" alt="Vulnerability details Issue Card"><figcaption><p>Vulnerability details Issue Card</p></figcaption></figure>

Again, there's a lot of information for you to understand, so take the time to understand how all of this information relates to your vulnerability, to help you decide on what fix actions to take.

For details, see [Issue card information](../../snyk-admin/snyk-projects/issue-card-information.md).

## Access more vulnerability information

Snyk provides detailed resources for more information about vulnerabilities, accessible directly from the card:

* [**Snyk Vulnerability Database**](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/snyk-vulnerability-database.md): access details on a specific vulnerability.
* [**Snyk Learn**](../../snyk-learn/): access general information about that type of vulnerability.

### Access Snyk Vulnerability Database

For Open Source and Container vulnerabilities, click on the Snyk vulnerability Identifier (on the right of the Severity Level) to access detailed [Snyk Vulnerability Database](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/snyk-vulnerability-database.md) information for that vulnerability, as defined by Snyk. For example:

<figure><img src="../../.gitbook/assets/image (174) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1) (1).png" alt="Access Snyk Vulnerability Database"><figcaption><p>Access Snyk Vulnerability Database</p></figcaption></figure>

For this example, click on the Snyk vulnerability Identifier to see how Hibernate core and its libraries are vulnerable to SQL injection:

<figure><img src="../../.gitbook/assets/image (149) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (2).png" alt="Snyk Vulnerability Database example entry"><figcaption><p>Snyk Vulnerability Database example entry</p></figcaption></figure>

{% hint style="info" %}
[Snyk Code](../../scan-with-snyk/snyk-code/) and [Snyk IaC](../../scan-with-snyk/snyk-iac/scan-your-iac-source-code/) issue cards have separate information sets for these areas.
{% endhint %}

### Access Snyk Learn

Click **Learn about this type of vulnerability** to access [Snyk Learn](https://learn.snyk.io/) security educational materials:

<figure><img src="../../.gitbook/assets/image (119) (1).png" alt="Access Snyk Learn from a vulnerability card"><figcaption><p>Access Snyk Learn from a vulnerability card</p></figcaption></figure>

For example, see [Snyk Learn SQL injection](https://learn.snyk.io/lessons/sql-injection/javascript/) for more details about this type of vulnerability.

{% hint style="info" %}
Some cards may not have Snyk Learn lessons available - if so, no links are presented..
{% endhint %}

## Understand the Snyk Priority Score

The [Snyk Priority Score](../../manage-risk/prioritize-issues-for-fixing/priority-score.md), ranging from 0 - 1,000, is our evaluation of the seriousness of the vulnerability. The Snyk Priority Score includes [CVSS](https://www.first.org/cvss/calculator/3.1) (Common Vulnerability Scoring System) information, plus other factors such as attack complexity and known exploits. For example, this **Hibernate** vulnerability has no known exploit allowing attackers to take advantage of that vulnerability.

Other factors also affect the score. For example, SQL injections are easy to run (you just need a web browser and submit a form), so increasing the score, but it takes more work to understand and exploit the results for that attack, so decreasing the score.

## Open source vulnerabilities: fixes and dependency information

For open-source library scans by Snyk Open Source, you can also access fix and dependency information in the **Fixes** and **Dependencies** tabs of your Project results.

### Fixes tab

Snyk knowledge of the transitive dependencies in your Project makes it possible for Snyk to offer fix advice in the **Fixes** tab:

<figure><img src="../../.gitbook/assets/Screenshot 2021-10-19 at 11.57.07.png" alt="Fix advice for Open Source vulnerabilities"><figcaption><p>Fix advice for Open Source vulnerabilities</p></figcaption></figure>

See [Fix your first vulnerability](fix-your-first-vulnerability.md) for more details.

### Dependencies tab

Snyk uses the package manager of your application to build the dependency tree and display it in the **Dependencies** tab of the Project view:

<figure><img src="../../.gitbook/assets/image (119) (1) (1).png" alt="Dependencies for Open Source vulnerabilities"><figcaption><p>Dependencies for Open Source vulnerabilities</p></figcaption></figure>

Click the file tree icon (![](<../../.gitbook/assets/image (201) (1) (1) (1) (1) (1) (1) (2) (2).png>)) to build the dependency tree, showing which components introduce a vulnerability. This helps you understand how the dependency was introduced to the application:

<figure><img src="../../.gitbook/assets/image23 (1) (1) (1).png" alt="Dependency tree details"><figcaption><p>Dependency tree details</p></figcaption></figure>

For example, the above screenshot shows a vulnerability based on the transitive dependency **qs@2.2.4**, brought in from the direct dependency **body-parser@ 1.9.0**.

Now that you understand your vulnerability information, you can decide how to fix it.

Continue with [Fix your first vulnerability](fix-your-first-vulnerability.md).
