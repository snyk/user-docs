# Importing additional repositories to Snyk

If you already have Projects in your Snyk Account, you can add other repositories to be tested by Snyk Code.

{% hint style="warning" %}
If the **Snyk Code** feature was previously disabled, and you want to apply Snyk Code testing to your existing repositories as well, you must [re-import these repositories](re-importing-existing-repositories-for-snyk-code-testing.md).
{% endhint %}

Follow these steps to import additional repositories to Snyk:

1\. On the Snyk Web UI, click either **Dashboard** or **Projects**.

2\. On the **Dashboard** or **Projects** page, click the **Add project** button. Then, select from the pop-up list the SCM that contains the repositories you want to import:

<figure><img src="../../../.gitbook/assets/image (363) (1) (1).png" alt="Add project button"><figcaption><p>Add project button</p></figcaption></figure>

The **Personal and Organization repositories** page appears, displaying the repositories available to you in the selected SCM.

3\. On the **Personal and Organization repositories** page, select the checkboxes of the repositories you want to import to Snyk. If you also want Snyk Code to test repositories you previously imported, select these repositories as well for re-import.

{% hint style="info" %}
The repositories you previously imported are indicated with a checkmark <img src="../../../.gitbook/assets/Snyk Code - Add Repositories dialog box - Check Mark.png" alt="Checkmark" data-size="line">.

To import all the repositories of a specific SCM account, select the **account checkbox**.

The **Settings** options on this page are NOT applicable to Snyk Code, and you can only import entire repositories. However, you can [exclude directories and files from the import process using the .snyk file](excluding-directories-and-files-from-the-import-process.md).
{% endhint %}

<figure><img src="../../../.gitbook/assets/Snyk Code - Add Repositories dialog box - Entire Repositories.png" alt="Select the checkboxes of the repositories to import"><figcaption><p>Select the checkboxes of the repositories to import</p></figcaption></figure>

4\. After you select the required repositories, click the **Add selected repositories button** to import the repositories into your Snyk Account:

<figure><img src="../../../.gitbook/assets/Snyk Code - Add Repositories dialog box - Re-import - Add selected repositories button.png" alt="Add selected repositories button"><figcaption><p>Add selected repositories button</p></figcaption></figure>

The selected repositories are imported to Snyk Code, and a progress bar appears on the Projects page. When the import is completed, a confirmation message appears on the Projects page. Your imported repositories appear as Target folders, each containing the **Code analysis** Project that includes the findings of the Snyk Code test:

<figure><img src="../../../.gitbook/assets/image (453) (1).png" alt="Imported Target folder with Code analysis Project"><figcaption><p>Imported Target folder with Code analysis Project</p></figcaption></figure>
