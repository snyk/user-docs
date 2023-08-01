# Re-importing existing repositories for the Snyk Code test

If your Snyk Account includes repositories that were imported to it before the **Snyk Code** option was enabled, and you want Snyk Code to test these repositories, you need to re-import these repositories to Snyk.

**To re-import existing repositories:**

1\. On the Snyk Web UI, click either the **Dashboard** or **Projects** tabs on the top menu.

2\. On the **Dashboard** or **Projects** page, click the **Add project** button. Then, select from the pop-up list the SCM that contains the repositories you want to re-import:

![](<../../../.gitbook/assets/image (99) (1).png>)

3\. In the **Personal and Organization repositories** page, select the checkboxes of the repositories you want to re-import to Snyk:

**Note**: The repositories you already imported are indicated with a check mark <img src="../../../.gitbook/assets/Snyk Code - Add Repositories dialog box - Check Mark.png" alt="" data-size="line">.

![](<../../../.gitbook/assets/Snyk Code - Re-importing - Existing repositories.png>)

4\. Click the **Add selected repositories** button to re-import the selected repositories:

![](<../../../.gitbook/assets/Snyk Code - Re-importing - Add selected repositories button.png>)

Your selected repositories are re-imported to Snyk. During the re-import process, Snyk Code analyzes these repositories, and a **Code analysis** Project is added to each repository, containing the results of the Snyk Code test.

{% hint style="warning" %}
If you are re-importing repositories that you deleted before the re-import, you need to refresh the **Projects** page to view the re-import results.
{% endhint %}
