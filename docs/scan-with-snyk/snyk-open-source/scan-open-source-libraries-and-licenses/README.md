# Scan open-source libraries and licenses

You can scan your open-source libraries using Snyk Open Source:&#x20;

* In the [Snyk Web UI](../../../getting-started/snyk-web-ui.md)
* With your [IDE](https://docs.snyk.io/integrations/ide-tools)
* With a [CI/CD integration](../../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/)
* Through the [Snyk CLI](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/)
* Through the [Snyk API](../../../snyk-api/reference/test-v1.md)

## View vulnerabilities in your open-source libraries

You can view vulnerability results for imported Projects. The **Projects** page appears by default after import, showing vulnerability information for the Snyk Projects you have imported, grouped into **Targets**, that is, the repositories you have scanned.

You can expand a **Target** to see vulnerability information for Projects, including the number of issues found, grouped by severity level:

Click a Project to open the issues page for that Project, where you will see the issue cards, showing the module where each issue was introduced, how to fix it, and more details about the vulnerability itself.

For more details, see [View Project information](../../../snyk-admin/snyk-projects/project-information.md).

## Fix vulnerabilities in your open-source libraries

For some languages, Snyk can fix vulnerabilities using fix pull/merge requests. For more information, see [Automatic and manual PRs with Snyk Open Source](../../pull-requests/snyk-pull-or-merge-requests/).

Navigate to the **Issues** card for a Project.

<figure><img src="../../../.gitbook/assets/os_project_issues_fix_vuln.png" alt=""><figcaption><p>Issues tab in an Open Source Project</p></figcaption></figure>

To fix vulnerabilities:

1. Click **Fix this vulnerability** to open a fix PR for this issue. Click **Fix these vulnerabilities** to fix multiple issues.
2. The **Open a Fix PR** screen opens, displaying the selected vulnerabilities.
3. Check or uncheck the issues you want to fix or remove from this fix.
4. Scroll to the bottom of the screen and click **Open a Fix PR**.
5. Snyk acts on the PR and displays a results screen.
6. Optionally, select the **Files changed** tab to see details of the changes made.

<figure><img src="../../../.gitbook/assets/screenshot_2021-04-09_at_17.46.22.png" alt=".Files changed tab in GitHub after triggering Fix PR for an open source project"><figcaption><p>Files changed tab in GitHub after triggering Fix PR for an open source project</p></figcaption></figure>

For more details, see [Fix your vulnerabilities](../manage-vulnerabilities/fix-your-vulnerabilities.md).
