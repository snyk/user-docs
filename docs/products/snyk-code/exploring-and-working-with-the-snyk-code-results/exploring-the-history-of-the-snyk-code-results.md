# Exploring the History of the Snyk Code Results

You can view the results of previous Snyk Code tests that were performed on an imported repository. The result history can include the following Snyk Code tests:

* The first test that was performed when the repository was initially imported to Snyk.
* Recurring tests that were automatically performed according to a pre-defined schedule.
* Tests that were performed on-demand using the **Retest** **now** option.

The result history is shown in the **History** page of the **Code Analysis** Project. This page displays the snapshots that were taken when a test was performed.

Only two unique snapshots can be shown on the **History** page. A unique snapshot means that the repository or its vulnerability results have changed since the last test, and the snapshot that was taken for the new test reflects these changes, and it is different from the previous one. If the repository or its results have not changed since the last test, the new snapshot will be identical to the previous snapshot, and it will be shown as an additional test on the **History** page. Therefore, you may have multiple past tests displayed on the **History** page, but at the most only two of them will show different results.

**To view the result history of a repository:**

1\. On the **Code Analysis** Project page of the required repository > Header area, click the **History** tab:

![](<../../../.gitbook/assets/Snyk Code - Results - History tab.png>)

The **History** page appears, displaying the list of tests that were performed by Snyk Code on the selected repository:

![](<../../../.gitbook/assets/Snyk Code - Results - History page.png>)

2\. To display the snapshot of a previous test, click the required test on the **History** page:

![](<../../../.gitbook/assets/Snyk Code - Results - History page - Clicking a test.png>)

The snapshot of the selected test appears in the **Code Analysis** page, with the caption **HISTORIC SNAPSHOT** and the date of the test:

![](<../../../.gitbook/assets/Snyk Code - Results - History page - Previous test snapshot.png>)

3\. To return to your most recent test results, click the **View most recent snapshot** link below the Header:

![](<../../../.gitbook/assets/Snyk Code - Results - History page - Previous snapshot - Return to recent option.png>)
