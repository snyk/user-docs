# Viewing and working with the PR Check results on the Snyk Web UI

After Snyk Code scanned your PR, and you clicked the link of the PR Check results in your SCM, you move directly to the PR checks page on the Snyk Web UI. From this Web UI page, you can do the following:

* [View the full details of each of the discovered issues in your PR](viewing-and-working-with-the-pr-check-results-on-the-snyk-web-ui.md#viewing-the-full-details-of-each-of-the-discovered-issues).
* [Mark failed PR checks as successful](viewing-and-working-with-the-pr-check-results-on-the-snyk-web-ui.md#\_ref105582006).
* [Return to your SCM from the Web UI](viewing-and-working-with-the-pr-check-results-on-the-snyk-web-ui.md#returning-to-your-scm).

### Viewing the full details of each of the discovered issues

After the Automatic PR Check discovers security issues in a PR, you can view the full details and fix example Snyk Code provides for these issues on the Web UI.

**To view the full details of issues discovered in PR Checks:**

**Note**: The **SAST check** page appears after you click the PR Checks results in the SCM.

* On the **SAST check** page, locate the required issue, and click the **Full details** button on the bottom right corner:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Web UI - Full details button.png>)

The full details of the discovered issue are displayed:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Web UI - Full details button - Results (1).png>)

### Marking failed PR checks as successful <a href="#_ref105582006" id="_ref105582006"></a>

If you want to merge PRs that have security issues, and therefore were failed by Snyk Code, you can mark them as Successful. When these issues will be marked as Successful, Snyk Code will pass them in your SCM and will allow you to merge them.

The **Mark as successful** action does NOT ignore the issue, but only allows it to pass in the current PR branch. If this issue will not be fixed, it will show up again in future commits and in the repository tests, once it will be merged into the target branch.

**Note:** The **Mark as successful** option is not supported in Bitbucket Cloud integration.

**To mark failed PR checks as successful:**

1\. On the **SAST check** page header, click the **Mark as successful in SCM** button on the right:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Web UI - Mark as successful (1).png>)

2\. On the confirmation message that appears, click **OK**:

!\[Graphical user interface, text, application, email

Description automatically generated]\(<../../../.gitbook/assets/Snyk Code - PR Checks - Web UI - Mark as successful - message (1) (1) (1) (1) (1) (2).png>)

A message appears on the **SAST check** page header, informing you that the failed PRs are now marked as successful:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Web UI - Mark as successful - Confirmation on the UI.png>)

On your SCM, the previous PR Check “**failed**” message is now changed to “**passed**”. In addition, there is an indication that the PR check was **Skipped**, and that a specific Snyk user forced this result change:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Mark as successful - On GitHub (1).png>)

### Returning to your SCM

You can return to your SCM directly from the Snyk Web UI.

**To return to your SCM from the Web UI:**

1\. On the **SAST check** page -> Project Summary Information area, click the **PR Check Group** option:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Web UI - PR Check Group option.png>)

2\. On the **Security check** page, click the repository link in the **REPOSITORY** field:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Web UI - Second page.png>)

You are moved back to your SCM.
