# View analysis results from Visual Studio extension

## Issues display in the Visual Studio extension

You can filter vulnerabilities by name or by severity.

Filter by name by typing the name of the vulnerability in the search bar.

![Filter by name](../../../.gitbook/assets/readme_image_3_2_1.png)

Filter by severity by selecting one or more of the severities when you open the search bar filter.

![Filter by severity](../../../.gitbook/assets/readme_image_3_2_2.png)

Users can configure the Snyk extension using the **Solution settings** in the **Options**.

<figure><img src="../../../.gitbook/assets/image (41) (1).png" alt=""><figcaption><p>Add the -d parameter in the Solution settings</p></figcaption></figure>

## Net new Issues versus all issues

For Projects using Git repositories or when you specify a reference folder, Snyk can filter the displayed issues to show only issues introduced in the working branch.&#x20;

This functionality reduces noise and allows you to focus only on current changes. This helps prevent issues early, thus unblocking your CI/CD pipeline and speeding up your deliveries.

The logic uses your local Git repository or any folder to compare the current findings with those in a base branch or reference folder. Net new issues scanning (delta scanning) shows you the difference between the two branches or folders, highlighting only the new issues.

In version 2.1.0 and later, you can choose **any folder** as your base for scanning.&#x20;

To apply the filter and only see the new issues, use the **total/new** toggle in the summary panel.

<div align="center" data-full-width="false"><figure><img src="../../../.gitbook/assets/image (366).png" alt="" width="375"><figcaption><p>Toggle in summary panel to show the total number of issues <br> and the number of issues in the checked out branch or current folder</p></figcaption></figure></div>

You can also enable net new issues feature in the [scan settings](visual-studio-extension-configuration-environment-variables-and-proxy.md#scan-configuration) for the Visual Studio extension.&#x20;

For newly created feature branches, there will be no reported issues. That is an intended state that developers would aim for, as shown in the screen image that follows:

<figure><img src="../../../.gitbook/assets/image (367).png" alt="" width="481"><figcaption><p>Successful state, no net new issiues found</p></figcaption></figure>

## Changing the base branch

The base branch is usually determined automatically for each Git repository.&#x20;

You can change the base branch or base folder by following these steps, as illustrated in the screen image that follows:

1. Toggle the total/new filter in the summary panel
2. Click on the top-level node in the Issues tree to change the branch or directory.
3. Use the dropdown selection to choose any branch or reference folder.

<figure><img src="../../../.gitbook/assets/image (368).png" alt=""><figcaption><p>Change the reference branch or reference directory for calculation of new new issues.</p></figcaption></figure>
