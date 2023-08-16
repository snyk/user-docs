# View Project issues, fixes, and dependencies

## Introduction

The following Project information is available on the Snyk Web UI:

* [Issues](view-project-issues-remediations-and-dependencies.md#view-issues): the number of vulnerabilities and license issues
* [Fixes](view-project-issues-remediations-and-dependencies.md#view-fixes): fix advice
* [Dependencies](view-project-issues-remediations-and-dependencies.md#view-dependencies): the total number of direct and transitive (nested) dependencies

### View issues

The Project details page displays Issue cards on the **Issues** tab. The information provided includes vulnerabilities and license issues (for Open Source Projects).

<figure><img src="../../.gitbook/assets/Screenshot 2021-10-19 at 11.49.30.png" alt="Project details Issues tab and filters"><figcaption><p>Project details Issues tab and filters</p></figcaption></figure>

Use the filters in the panel to the left to narrow the search for issues. Select the checkboxes to filter issues by **Issue type**, **Severity**, **Exploit Maturity**, **Fixability**, and **Status**. You can also edit the **Priority Score** slider to change the range displayed; the default is 0 to 1000.

Issue details are shown on Issue cards in the main area, sorted by priority score. See [Issue card information](issue-card-information.md) for more details.

#### Fix issues

Snyk provides features to fix issues identified during scanning. See [Fix your vulnerabilities](../../scan-application-code/snyk-open-source/starting-to-fix-vulnerabilities/fix-your-vulnerabilities.md) for more details.

### View issue details

Click on an issue to view details, including its [priority score](../prioritizing-and-managing-issues/priority-score.md).

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2023-06-13 at 08.43.14.png" alt="View issue details"><figcaption><p>View issue details</p></figcaption></figure>

</div>

* Click **Show more detail** to [view detailed information](view-project-issues-remediations-and-dependencies.md#introduction).
* Click **Learn about this type of vulnerability** to access [Snyk Learn](../../more-info/snyk-learn.md) training information.

#### View detailed information

Click **Show more detail** to view detailed information about the vulnerability from [Snyk's vulnerability database](https://snyk.io/product/vulnerability-database/):

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2023-06-13 at 08.47.54.png" alt="More information from Snyk Vulnerability database"><figcaption><p>More information from Snyk Vulnerability database</p></figcaption></figure>

</div>

### View fixes

Snyk's knowledge of the transitive dependencies in your Project make it possible for Snyk to offer fix advice, in the **Fixes** tab:

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2021-10-19 at 11.57.07.png" alt="Issue details Fixes tab"><figcaption><p>Issue details Fixes tab</p></figcaption></figure>

</div>

See [Fix your vulnerabilities](../../scan-application-code/snyk-open-source/starting-to-fix-vulnerabilities/fix-your-vulnerabilities.md) for details.

### View dependencies

Snyk uses the package manager for your application to build the dependency tree and display it in the **Dependencies** tab of the Project issues detail page. This tab shows which components introduce a vulnerability, which indicates how the dependency was introduced to the application.

For example.

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2023-06-13 at 08.57.23.png" alt="Issues detail page dependencies tab"><figcaption><p>Issues detail page dependencies tab</p></figcaption></figure>

</div>
