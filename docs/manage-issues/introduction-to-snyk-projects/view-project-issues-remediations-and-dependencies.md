# View Project issues, fixes, and dependencies

## Introduction

The following Project information is available on the Snyk Web UI:

* [Issues](view-project-issues-remediations-and-dependencies.md#view-issues): the number of vulnerabilities and license issues
* [Fixes](view-project-issues-remediations-and-dependencies.md#view-fixes): fix advice
* [Dependencies](view-project-issues-remediations-and-dependencies.md#view-dependencies): the total number of direct and transitive (nested) dependencies

### View issues

The Project details page displays Issue cards on the Issues tab. The information provided includes vulnerabilities and, for Open Source Project, license issues.

<figure><img src="../../.gitbook/assets/Screenshot 2021-10-19 at 11.49.30.png" alt="Project details Issues tab and filters"><figcaption><p>Project details Issues tab and filters</p></figcaption></figure>

Use the filters in the panel to the left to narrow the search for issues. Select the checkboxes to filter issues by **Issue type**, **Severity**, **Exploit Maturity**, **Fixability**, and **Status**. You can also edit the **Priority Score** slider to change the range displayed; the default is 0 to 1000.

Issue details are shown on Issue cards in the main area, sorted by priority score. For details see [Issue card information](issue-card-information.md).

#### Fix issues (Fix PR)

Snyk provides powerful features to fix issues identified during scanning. Use the Fix option on the Issues or Fixes tab.

<figure><img src="../../.gitbook/assets/image27.png" alt="ix option on Fixes tab"><figcaption><p>Fix option on Fixes tab</p></figcaption></figure>

You can also select **Fix this vulnerability** for a specific issue:

<figure><img src="../../.gitbook/assets/image26.png" alt="Fix this vulnerability for an issue"><figcaption><p>Fix this vulnerability for an issue</p></figcaption></figure>

See [Fix vulnerabilities with Snyk Open Source](../../scan-application-code/snyk-open-source/open-source-basics/) for an overview.\
See [Manage issues](../) for more details.

### View issue details

For each issue, details about the vulnerability are displayed including its [priority score](../issue-management/priority-score.md).

<figure><img src="../../.gitbook/assets/image12.png" alt="View issue details"><figcaption><p>View issue details</p></figcaption></figure>

#### More information - Snyk Vulnerability Database

Click **More about this issue** to view detailed information about the vulnerability from [Snyk's vulnerability database](https://snyk.io/product/vulnerability-database/), which provides a deeper insight into the issue, including its CVSS score:

<figure><img src="../../.gitbook/assets/image15.png" alt="More information from Snyk Vulnerability database"><figcaption><p>More information from Snyk Vulnerability database</p></figcaption></figure>

### View fixes

Snyk's knowledge of the transitive dependencies in your Project make it possible for Snyk to offer fix advice, in the **Fixes** tab:

<figure><img src="../../.gitbook/assets/Screenshot 2021-10-19 at 11.57.07.png" alt="Issue details Fixes tab"><figcaption><p>Issue details Fixes tab</p></figcaption></figure>

See [Fix your vulnerabilities](../../scan-application-code/snyk-open-source/starting-to-fix-vulnerabilities/fix-your-vulnerabilities.md) for details.

### View dependencies

Snyk uses the package manager for your application to build the dependency tree and display it in the **Dependencies** tab of the Project issues detail page. This tab shows which components introduce a vulnerability, which indicates how the dependency was introduced to the application.

The example that follows shows a vulnerability based on the transitive dependency **qs@2.2.4**, brought in from the direct dependency **body-parser@ 1.9.0**.

<figure><img src="../../.gitbook/assets/image23.png" alt="Issues detail page dependencies tab"><figcaption><p>Issues detail page dependencies tab</p></figcaption></figure>
