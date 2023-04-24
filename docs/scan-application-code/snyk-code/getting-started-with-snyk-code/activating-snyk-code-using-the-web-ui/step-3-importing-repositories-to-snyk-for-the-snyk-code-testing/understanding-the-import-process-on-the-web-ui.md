# Understanding the import process on the Web UI

The process of importing your selected repositories from your SCM to your Snyk Account, starts after you select the required repositories on the Snyk Web UI, and clicking the **Add selected repositories** button to import them. When the import starts, a progress bar appears on the **Projects** page:

<figure><img src="../../../../../.gitbook/assets/image (215).png" alt=""><figcaption></figcaption></figure>

When the import is completed, a confirmation message appears on the **Projects** page, informing you of the success of the import. Your imported repositories appear as separate Target folders on the **Projects** page. Each Target folder bears the name of your SCM account and the imported repository, and it contains the Snyk Projects that were created for it:

<figure><img src="../../../../../.gitbook/assets/image (172).png" alt=""><figcaption></figcaption></figure>

**Note**: If some of the files in the selected repositories were not imported, you receive a message informing you that “**X projects failed**”. To view more details, click the **last import** **log** link on the message:

![](<../../../../../.gitbook/assets/Snyk Code - Imported Repositories - Failed import.png>)

During the import process, Snyk Code analyzes the directories and files in the selected repositories. Then, Snyk Code creates one Snyk Project, called **Code** **analysis**, for each imported repository, to contain the test results. The Snyk Code testing is already performed during the import process, and once the import is completed, the Snyk Code results are displayed in the **Code** **analysis** Project without any additional action on your part:

<figure><img src="../../../../../.gitbook/assets/image (178) (1).png" alt=""><figcaption></figcaption></figure>

In addition, when the import is completed, you receive an email informing you on the import completion. This confirmation email also contains links to the results of the import:

![](<../../../../../.gitbook/assets/Snyk Code - Imported Repository - Complete - Email Confirmation.png>)

After the import, Snyk Code will continue to monitor the code of the imported repositories in search of vulnerabilities. When a repository is imported, Snyk automatically creates a webhook to its SCM, and this webhook is used after the import for regular scans. By default, Snyk Code analyzes and tests the Pull Requests in the imported repositories on a weekly basis, and it sends you a **New issues and remediations** summary via email, when new vulnerabilities are detected.

{% hint style="info" %}
You can also set a daily test for your imported code or cancel the automatic testing via the **Code analysis** Project Settings. In addition, you can retest the imported code on-demand, using the **Retest now** option on the **Code Analysis** page.
{% endhint %}
