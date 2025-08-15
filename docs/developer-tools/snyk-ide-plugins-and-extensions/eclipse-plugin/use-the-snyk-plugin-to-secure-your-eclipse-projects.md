# Use the Snyk plugin to secure your Eclipse projects

After the Eclipse plugin is downloaded and authentication is complete, the plugin starts the workspace scan. You may notice a confirmation that a workspace scan is starting. Alternatively, you can trigger a workspace scan from the context menu of your Project, or from the Snyk View.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-10-19 at 09.02.25.png" alt=""><figcaption><p>Starting workspace scan</p></figcaption></figure>

## Issues are displayed in the Eclipse plugin

All of the issues found by Snyk are now integrated natively with Eclipse flows. Issues are shown in the **Problems** tab, as illustrated in the following screen image. There is a squiggly line indicating the issue while you code, along with gutter icons to indicate where the issue is.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-05-13 at 12.20.26.png" alt=""><figcaption><p>Eclipse Problems tab</p></figcaption></figure>

In addition, starting with version 3+, Snyk provides a custom UI in the Snyk Tab, that displays issue details:

<figure><img src="../../../.gitbook/assets/image (322).png" alt=""><figcaption><p>Issue details in Eclipse plugin</p></figcaption></figure>

## Net new issues versus all issues

Starting with version 3.1.0, it is possible to see only newly introduced issues.

This functionality reduces noise and allows you to focus only on current changes. This helps prevent issues early, thus unblocking your CI/CD pipeline and speeding up your deliveries.

The logic uses your local Git repository or any folder to compare the current findings with those in a base branch or reference folder. Net new issues scanning (delta scanning) shows you the difference between the two branches or folders, highlighting only the new issues.

To apply the filter and see only the new issues, use the **total**/**new** toggle in the summary panel, or apply the **Show only Net New Issues filter** from the **View** menu (under **Issues Status**).\


<figure><img src="../../../.gitbook/assets/image (371).png" alt=""><figcaption><p>Net new issues filter enabled after clicking on the total/new issues toggle</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (327).png" alt=""><figcaption><p>Activate Net new issues in the dot menu of the Snyk View</p></figcaption></figure>

For newly created feature branches, there will be no reported issues. That is an intended state that developers would aim for, as illustrated in the screen image that follows:

<figure><img src="../../../.gitbook/assets/image (373).png" alt=""><figcaption><p>No new issues introduced in a newly created branch</p></figcaption></figure>

The base branch is usually automatically determined for each Git repository.&#x20;

You may change the base branch or base folder by following these steps, as illustrated in the screen image that follows:

1. Toggle the **total**/**new** filter in the summary panel.
2. Click on the top-level node in the Issues tree to change the branch or directory.
3. Use the dropdown selection to choose any branch or reference folder.

<figure><img src="../../../.gitbook/assets/image (374).png" alt=""><figcaption><p>Changing reference branch or reference directory for calculation net new issues</p></figcaption></figure>

\
Continue by following the instructions on the page for the type of scan you are doing:

* [SAST scanning results (SAST, Snyk Code)](sast-scanning-results-sast-snyk-code.md)
* [Misconfiguration scanning results (Snyk Infrastructure as Code)](misconfiguration-scanning-results-snyk-infrastructure-as-code.md)
* [Third-party dependency scanning (SCA, Snyk Open Source)](third-party-dependency-scanning-sca-snyk-open-source.md)
