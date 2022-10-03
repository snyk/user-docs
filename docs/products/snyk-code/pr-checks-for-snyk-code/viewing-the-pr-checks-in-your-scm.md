# Viewing the PR Checks in your SCM

After you enabled the Automatic PR Checks feature on the Organization or Project level, you can view the status of your new PRs in your integrated SCM. In the SCM, the PR Check results for Snyk Code are grouped together in a single security check, called “**code/snyk**”.

{% hint style="info" %}
[PR Checks for Snyk Open Source](../../snyk-open-source/pr-checks-for-snyk-open-source/) are grouped and displayed in another row, called **security/snyk**.
{% endhint %}

The following statuses can appear on your Snyk Code checks in the integrated SCM:

* **Passed/successful** - no issues were discovered, and all the checks have passed.\
  **Note**: If you selected to manually pass failed checks via the Web UI, the checks will be displayed as “**passed**” in the SCM, but there will be an indication that the PR check was **Skipped**, and that a specific Snyk user forced this result change. For more information, see [Marking failed PR checks as successful](viewing-and-working-with-the-pr-check-results-on-the-snyk-web-ui.md#\_ref105582006).
* **Pending** - this status appears until the Snyk Code test is completed.
* **Failed** – security issues were identified in the PR. These issues must be fixed in order to pass the PR check.

**To view the PR checks in your SCM:**

**Note**: The instructions below use GitHub as the integrated SCM, but they apply to all supported SCMs with some small differences depending on the SCM workflow and platform. For more information, see [Exploring the display of the PR Checks on different integrated SCMs](viewing-the-pr-checks-in-your-scm.md#\_ref105582759).

1\. In your SCM, after you finished creating a new PR, click the **Create** **pull request** button. Snyk Code automatically checks the PR you created, and displays the results of the check – either **Failed** or **Passed** – in your SCM:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - GitHub - Checks Failed.png>)

2\. To view the details of the issues that were found in the PR on the Snyk Web UI, click the **Details** link on the left:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - GitHub - Checks Failed - Details link.png>)

The Snyk Web UI opens, displaying the details of the issues that were found in the PR:

<figure><img src="../../../.gitbook/assets/image (154).png" alt=""><figcaption></figcaption></figure>

You can now [view and work with the discovered PR Check issues on the Web UI](viewing-and-working-with-the-pr-check-results-on-the-snyk-web-ui.md).

### Exploring the display of the PR Checks on different integrated SCMs <a href="#_ref105582759" id="_ref105582759"></a>

The Automatic PR Checks feature works in a similar way on all supported SCMs. However, due to the differences between the various SCMs, the results of the PR Checks will be displayed differently on each SCM.

**Note**: The screenshots in this section show the aggregated results of the PR Checks for both Snyk Code and Snyk Open Source.

#### **GitHub and GitHub Enterprise**

**Notes**:

* The PR Check results appear in the same way on GitHub and GitHub Enterprise integration.
* The instructions in this entire PR Checks section include screenshots from GitHub integration.

The results of the PR Checks appear on GitHub and GitHub Enterprise integrations as follows:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - GitHub - Results .png>)

**To open the Snyk Web UI to view additional details on the PR Check results:**

* On the results area, click the **Details** link on the **code/snyk** **(organization name)** row.

#### **GitLab**

The results of the PR Checks appear on GitLab first as an indication of the Pipeline status:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - GitLab - Pipeline.png>)

When you click the pipeline link, the PR Checks results are displayed:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - GitLab (1).png>)

**To open the Snyk Web UI to view additional details on the PR Check results:**

* On the results area, click the **code/snyk** **(organization name)** link:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - GitLab - Link to Results.png>)

#### **Bitbucket Cloud**

The results of the PR Checks appear on Bitbucket Cloud as follows:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Bitbucket Cloud.png>)

**To open the Snyk Web UI to view additional details on the PR Check results:**

* On the results area, click the **code/snyk** **(organization name)** link:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Bitbucket Cloud - Results link.png>)

#### **Azure Repos**

The results of the PR Checks appear on Azure Repos as follows:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Azure Repos.png>)

**To view additional details on all the PR checks in Azure Repos platform:**

1\. Click the **View 2 checks** link:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Azure Repos - View Checks link.png>)

The details of the results appear in a separate pane:

![](<../../../.gitbook/assets/Snyk Code - PR Checks - Azure Repos - View Checks link - Pane.png>)

2\. To open the Snyk Web UI to view additional details on the PR Check results, click the **SAST issue** link on the **Checks** pane.
