# Re-importing existing repositories for Snyk Code testing

If your Snyk account includes repositories that were imported to it before the **Snyk Code** option was enabled, and you want Snyk Code to test these repositories, you must re-import these repositories to Snyk.

Follow these steps to re-import existing repositories:

1\. On the Snyk Web UI, click either the **Dashboard** or **Projects**.

2\. On the **Dashboard** or **Projects** page, click the **Add project** button. Then, select from the pop-up list the SCM that contains the repositories you want to re-import:

<figure><img src="../../../.gitbook/assets/image (99) (1).png" alt="Add project button"><figcaption><p>Add project button</p></figcaption></figure>

3\. On the **Personal and Organization repositories** page, select the checkboxes of the repositories you want to re-import to Snyk:

The repositories you already imported are indicated with a checkmark <img src="../../../.gitbook/assets/Snyk Code - Add Repositories dialog box - Check Mark.png" alt="Checkmark" data-size="line">.

<figure><img src="../../../.gitbook/assets/Snyk Code - Re-importing - Existing repositories.png" alt="Select the checkboxes of the repositories to import"><figcaption><p>Select the checkboxes of the repositories to import</p></figcaption></figure>

4\. Click the **Add selected repositories** button to re-import the selected repositories:

<figure><img src="../../../.gitbook/assets/Snyk Code - Re-importing - Add selected repositories button.png" alt="Add selected repositories buton"><figcaption><p>Add selected repositories buton</p></figcaption></figure>

Your selected repositories are re-imported to Snyk. During the re-import process, Snyk Code analyzes these repositories, and a **Code analysis** Project is added to each repository, containing the results of the Snyk Code test.

{% hint style="warning" %}
If you are re-importing repositories you deleted before the re-import, you must refresh the **Projects** page to view the re-import results.
{% endhint %}
