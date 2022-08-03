# Understand your vulnerabilities

{% hint style="info" %}
**Recap**\
You have [viewed and understood scanned Projects](view-your-first-snyk-projects.md); now you can look at the details of vulnerabilities in that Project.
{% endhint %}

### See your vulnerabilities

First, open a target to see your Snyk Projects:

![](<../../.gitbook/assets/image (143).png>)

Next, click on a Snyk Project in that list to see details of the vulnerabilities found in that Project.&#x20;

For example, for a code analysis scanned by Snyk Code:

![](<../../.gitbook/assets/image (93).png>)

See [View project information](../../snyk-web-ui/introduction-to-snyk-projects/view-project-information.md) for more details.&#x20;

### View Issue Cards

Now, look at the vulnerability information for each Snyk Project, provided in Issue Cards:

![](<../../.gitbook/assets/image (55).png>)

Again, there's a lot of information for you to understand, so take the time to understand how all of this information relates to your vulnerability, to help you decide on what fix actions to take.

For details, see [Issue card information](https://docs.snyk.io/introducing-snyk/introduction-to-snyk-projects/issue-card-information) (docs) and [Understand issue components](https://training.snyk.io/learn/course/introduction-to-the-snyk-ui/scan-results/understand-issue-components?page=1) (training).

### Access more vulnerability information

Snyk provides detailed resources for more information about vulnerabilities, accessible directly from the card:

* **Snyk Vulnerability Database**: access details on a specific vulnerability.
* **Snyk Learn**: access general information about that type of vulnerability.

#### Access Snyk Vulnerability Database

For Open Source and Container vulnerabilities, click on the Snyk vulnerability Identifier (on the right of the Severity Level) to access detailed [Snyk Vulnerability Database](../../features/fixing-and-prioritizing-issues/starting-to-fix-vulnerabilities/using-the-snyk-vulnerability-database.md) information for that vulnerability, as defined by Snyk. For example:

![](<../../.gitbook/assets/image (73).png>)

For this example, click on the Snyk vulnerability Identifier to see how Hibernate core and its libraries are vulnerable to SQL injection:

![](<../../.gitbook/assets/image (112) (2).png>)

{% hint style="info" %}
[Snyk Code](../../products/snyk-code/) and [Snyk IaC](../../products/snyk-infrastructure-as-code/) issue cards have separate information sets for these areas.
{% endhint %}

#### Access Snyk Learn&#x20;

To research more about a vulnerability, click **Learn about this type of vulnerability** to access [Snyk Learn](https://learn.snyk.io/) security educational materials:

![](<../../.gitbook/assets/image (36).png>)

For example, see [Snyk Learn SQL injection](https://learn.snyk.io/lessons/sql-injection/javascript/) for more details about this type of vulnerability.

### Understand the Snyk Priority Score

The [Snyk Priority Score](../../features/fixing-and-prioritizing-issues/issue-management/snyk-priority-score.md), ranging from 0 - 1,000, is our evaluation of the seriousness of the vulnerability. The Snyk Priority Score includes [CVSS](https://www.first.org/cvss/calculator/3.1) (Common Vulnerability Scoring System) information, and other factors such as attack complexity and known exploits. For example, this Hibernate vulnerability has no known exploit allowing attackers to take advantage of that vulnerability.

Other factors also affect the score. For example, SQL injections are easy to run (you just need a web browser and submit a form), so increasing the score, but it takes more work to understand and exploit the results for that attack, so decreasing the score.

### Open source vulnerabilities: fixes and dependency information

For open source library scans by Snyk Open Source, you can also access fix and dependency information., in the **Fixes** and **Dependencies** tabs of you Project results.

#### Fixes tab

Snyk's knowledge of the transitive dependencies in your project make it possible for Snyk to offer fix advice, in the **Fixes** tab:

![](<../../.gitbook/assets/Screenshot 2021-10-19 at 11.57.07.png>)

See [Fixing vulnerabilities](../../products/snyk-open-source/open-source-basics/fixing-vulnerabilities.md) for more details

#### Dependencies tab

Snyk uses the package manager of your application to build the dependency tree and display it in the **Dependencies** tab of the project view:

![](<../../.gitbook/assets/image (96).png>)

Click the file tree icon (![](<../../.gitbook/assets/image (31) (1).png>)) to build the dependency tree, showing which components introduce a vulnerability. This helps you understand how the dependency was introduced to the application:

![](<../../.gitbook/assets/image23 (1).png>)

For example, the above screenshot shows a vulnerability based on the transitive dependency **qs@2.2.4**, brought in from the direct dependency **body-parser@ 1.9.0**.

### What's next?

Now you understand your vulnerability information, you can decide how to fix it.

See [Fix your first vulnerability](fix-your-first-vulnerability.md).
