# Fixing vulnerabilities

Snyk provides actionable remediation advice to fix vulnerabilities. For more details, see [Remediate your vulnerabilities](https://docs.snyk.io/fixing-and-prioritizing-issues/issue-management/remediate-your-vulnerabilities).

Snyk supports workflows to fix vulnerabilities using:

* Automatic pull / merge requests \(PRs / MRs\).
* Manual pull / merge requests.

## **Automatic pull / merge requests**

For projects imported via an SCM \(Source Code Manager\), Snyk offers the following types of automated pull / merge requests:

* [Fix pull requests for new vulnerabilities](https://support.snyk.io/hc/en-us/articles/360017186498-Fix-PRs-for-new-vulnerabilities)
* [Fix pull requests to clear the backlog of vulnerabilities in priority order](https://support.snyk.io/hc/en-us/articles/360017186958-Fix-PRs-to-clear-the-backlog-of-vulnerabilities-in-priority-order)
* [Dependency upgrade pull requests](https://support.snyk.io/hc/en-us/articles/360006581898-Upgrading-dependencies-with-automatic-PRs)

## Manual pull / merge requests for a project code

To generate a PR / MR directly from your project, using the Snyk UI:

1. Navigate to your project from the project list.
2. Select the file.  ![image22.png](../../.gitbook/assets/image22.png)
3. Select **Open a Fix PR/MR** or **Fix this vulnerability**:  ![image21.png](../../.gitbook/assets/image21.png)
4. A preview screen appears, showing you what fixes will be applied:

![image18.png](../../.gitbook/assets/image18.png) 5. Click **Open a Fix PR** on this screen to generate the pull request.

