# Code Analysis results history

You can view the results of previous Snyk Code tests performed on an imported repository. The result history can include the following Snyk Code tests:

* The first test performed when the repository was initially imported to Snyk.
* Recurring tests that were automatically performed according to a pre-defined schedule.
* Tests that were performed on demand with the **Retest** **now** option.

The result history is shown on the **History** page of the **Code Analysis** Project. This page displays the snapshots taken when a test was performed.

Only two unique snapshots can be shown on the **History** page. A unique snapshot is made when the repository or its vulnerability results have changed since the last test, and the snapshot taken for the new test reflects these changes and is different from the previous one. If the repository or its results have not changed since the last test, the new snapshot will be identical to the previous snapshot, and it will be shown as an additional test on the **History** page. Therefore, you may have multiple past tests displayed on the **History** page, but at most only two of them will show different results.

## View Code Analysis history

Follow these steps to view the result history for a repository:

1\. On the **Code Analysis** page for the repository, in the header area, click the **History** tab:

<figure><img src="../../../.gitbook/assets/Snyk Code - Results - History tab.png" alt="History tab on Code Analysis page"><figcaption><p>History tab on Code Analysis page</p></figcaption></figure>

The **History** page opens, displaying the list of tests performed by Snyk Code on the selected repository:

<figure><img src="../../../.gitbook/assets/Snyk Code - Results - History page.png" alt="Code Analysis history page"><figcaption><p>Code Analysis history page</p></figcaption></figure>

2\. To display the snapshot of a previous test, click that test on the **History** page:

<figure><img src="../../../.gitbook/assets/Snyk Code - Results - History page - Clicking a test.png" alt="Display snapshot of previous test"><figcaption><p>Display snapshot of previous test</p></figcaption></figure>

The snapshot of the selected test appears on the **Code Analysis** page, with the caption **HISTORIC SNAPSHOT** and the date of the test:

<figure><img src="../../../.gitbook/assets/Snyk Code - Results - History page - Previous test snapshot.png" alt="Historic snapshot"><figcaption><p>Historic snapshot</p></figcaption></figure>

3\. To return to your most recent test results, click **View most recent snapshot**:

<figure><img src="../../../.gitbook/assets/Snyk Code - Results - History page - Previous snapshot - Return to recent option.png" alt="View most recent snapshot"><figcaption><p>View most recent snapshot</p></figcaption></figure>
