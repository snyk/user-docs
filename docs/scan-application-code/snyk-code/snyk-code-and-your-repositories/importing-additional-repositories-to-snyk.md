# Importing additional repositories to Snyk

If you already have Projects in your Snyk Account, you can add other repositories to be tested by Snyk Code.

{% hint style="warning" %}
If the **Snyk Code** feature was previously disabled, and you want to apply the Snyk Code testing to your existing repositories as well, you need to [re-import these repositories](re-importing-existing-repositories-for-the-snyk-code-test.md).
{% endhint %}

**To import additional repositories to Snyk:**

1\. Open the Snyk Web UI, and click either the **Dashboard** or **Projects** tabs on the top menu.

2\. On the **Dashboard** or **Projects** page, click the **Add project** button. Then, select from the pop-up list the SCM that contains the repositories you want to import:

![](<../../../.gitbook/assets/image (363) (1) (1).png>)

The **Personal and Organization repositories** page appears, displaying the repositories available to you in the selected SCM.

3\. In the **Personal and Organization repositories** page, select the checkboxes of the repositories you want to import to Snyk. If you want Snyk Code to also test repositories you previously imported, select these repositories as well for re-import:

**Notes**:

* The repositories you previously imported are indicated with a check mark ![](<../../../.gitbook/assets/Snyk Code - Add Repositories dialog box - Check Mark.png>).
* To import all the repositories of a specific SCM account, select the account checkbox.
* The **Settings** options on this page are NOT applicable to Snyk Code, and you can only import entire repositories. However, you can [exclude directories and files from the import process using the .snyk file](excluding-directories-and-files-from-the-import-process.md).

![](<../../../.gitbook/assets/Snyk Code - Add Repositories dialog box - Entire Repositories.png>)

4\. After you selected the required repositories, click the **Add selected repositories button** to import the repositories into your Snyk Account:

![](<../../../.gitbook/assets/Snyk Code - Add Repositories dialog box - Re-import - Add selected repositories button.png>)

The selected repositories are imported to Snyk Code, and a progress bar appears on the **Projects** page. When the import is completed, a confirmation message appears on the **Projects** page, informing you of the success of the import. Your imported repositories appear as Target folders, each containing the **Code analysis** Project that includes the findings of the Snyk Code test:

<figure><img src="../../../.gitbook/assets/image (453) (1).png" alt=""><figcaption></figcaption></figure>
