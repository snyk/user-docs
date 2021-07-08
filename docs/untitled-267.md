# View project issues, remediation and dependencies

* [ Introduction to projects](/hc/en-us/articles/360019058297-Introduction-to-projects)
* [ View project information](/hc/en-us/articles/360011450838-View-project-information)
* [ View project issues, remediation and dependencies](/hc/en-us/articles/360016910877-View-project-issues-remediation-and-dependencies)
* [ View project settings](/hc/en-us/articles/360017002718-View-project-settings)
* [ View project history](/hc/en-us/articles/360016910977-View-project-history)
* [ Issue card information](/hc/en-us/articles/360018049037-Issue-card-information)
* [ Project Tags](/hc/en-us/articles/360013865038-Project-Tags)
* [ Project Attributes](/hc/en-us/articles/360012703537-Project-Attributes)

##  View project issues, remediation and dependencies

For your project, you can view the following displays:

* [Issues](): the number of vulnerabilities and license issues.
* [Remediation](): remediation advice.
* [Dependencies](): the total number of direct and transitive \(nested\) dependencies.

Below the project summary details, you can see vulnerabilities and licensing issues, in the **Issues** tab:

Use the left-hand area to filter and search issues. Click the checkboxes to filter issues by **Issue type**, **Severity**, **Exploit Maturity**, and **Status**. You can also edit the **Priority Score** slider to change the range displayed \(by default this is set from 0 - 1000\).

Issues details appear in the main area, sorted by priority score. See **View Issue Details.**

### Fix issues \(Fix PR\)

Snyk provides powerful features to fix issues identified during scanning, as shown in the **Open a fix PR** section of the **Issues** and **Remediation** displays:

You can also select to fix a specific issue, clicking **Fix this vulnerability** on a specific issue: 

See [Fixing vulnerabilities](https://support.snyk.io/hc/en-us/articles/360011484018-Fixing-vulnerabilities) for an overview.  
See [Fixing and prioritizing issues](https://support.snyk.io/hc/en-us/categories/360001328418-Fixing-and-prioritizing-issues) for more details.

### View issue details

For each issue, this display shows details about the vulnerability including its priority score \(see [Prioritizing Snyk issues](https://support.snyk.io/hc/en-us/articles/360009884837-Prioritizing-Snyk-issues)\):

Click **More about this issue** to view detailed information about the vulnerability using [Snyk's vulnerability database](https://snyk.io/product/vulnerability-database/), giving you a deeper insight into the issue, including its CVSS score:

Snyk's knowledge of the transitive dependencies in your project make it possible for Snyk to offer remediation advice, in the **Remediation** tab:

See [Remediate your vulnerabilities](https://support.snyk.io/hc/en-us/articles/360006113798-Remediate-your-vulnerabilities) for details.

Synk uses the package manager of your application to build the dependency tree and display it in the **dependency** tab of the project view. This shows which components introduce a vulnerability, to show how the dependency was introduced to the application:

For example, the above screenshot shows a vulnerability based on the transitive dependency **qs@2.2.4**, brought in from the direct dependency **body-parser@ 1.9.0**.

