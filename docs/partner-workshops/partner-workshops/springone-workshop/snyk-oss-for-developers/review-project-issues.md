# Review Project Issues

## Reviewing application issues

The Snyk UI makes it easy to view your application vulnerabilities and licensing issues per project. Start by finding your application on the project tab, expand the project, and select the **package manager file** associated with your project. In our sample application SPC, select the **pom.xml** file.

![](../../../.gitbook/assets/screen-shot-2020-08-21-at-4.43.05-pm.png)

Snyk displays detailed information about your project. The following are some of the critical pieces of information Snyk knows about your application.

* **Vulnerabilities** - This is the total number of know vulnerabilities for your application.
* **Dependencies** - The total number of dependencies for your application. This includes direct and transitive dependencies from the open-source. 
* **Source** - Where the project is sourced in the developer toolchain. In our example, it is GitHub.
* **Branch** - Which branch the project is viewing. This is specific to projects sourced from SCM.

![](../../../.gitbook/assets/screen-shot-2020-08-21-at-4.48.13-pm.png)

After reviewing your project details, Snyk users start to review their known issues which include application vulnerabilities and licensing issues. The issues section is configurable to identify the most important vulnerabilities and licensing issues first. Prioritization is important and the users can select **Issue Type**, **Severity**, **Exploit Maturity**, and **Status** to filter the most critical information first.

As you review each vulnerability you will notice a _**more about this issue**_ ****option at the bottom of the vulnerability listing. This provides detailed information about the vulnerability using [Snyk's vulnerability database](https://snyk.io/product/vulnerability-database/) and gives you insight on the severity and exploit maturity. Select that option to view a detailed explanation and CVSS scoring information.

![](../../../.gitbook/assets/screen-shot-2020-08-21-at-4.47.15-pm.png)

