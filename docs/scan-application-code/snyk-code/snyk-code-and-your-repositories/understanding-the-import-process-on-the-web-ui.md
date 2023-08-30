# Understanding the import process on the Web UI

After you select repositories on the Snyk Web UI and click the **Add selected repositories** button, the import starts, and a progress bar appears on the Projects page.

<figure><img src="../../../.gitbook/assets/importing projects.png" alt="Importing repositories progress bar"><figcaption><p>Importing repositories progress bar</p></figcaption></figure>

When the import is finished, a confirmation message appears on the Projects page. Your imported repositories appear as separate Target folders on the Projects page. Each Target folder has the name of your SCM account and the imported repository and contains the Snyk Projects that were created for it.

<figure><img src="../../../.gitbook/assets/target folders.png" alt="Target folders on Projects page"><figcaption><p>Target folders on Projects page</p></figcaption></figure>

{% hint style="info" %}
If some of the files in the selected repositories were not imported, you receive a notification about the Projects that were not imported. Click **View import log** to see more details.
{% endhint %}

<figure><img src="../../../.gitbook/assets/import_failed.png" alt="Project could not be created notification"><figcaption><p>Project could not be created notification</p></figcaption></figure>

Snyk Code analyzes the directories and files in the selected repositories during import. Then, Snyk Code creates a Project called **Code** **analysis** for each imported repository to contain the test results. The Snyk Code testing is already performed during the import process. After the import is completed, the Snyk Code results are displayed in the **Code** **analysis** Project without any additional action on your part.

<figure><img src="../../../.gitbook/assets/code analysis project.png" alt="Code analysis Project"><figcaption><p>Code analysis Project</p></figcaption></figure>

In addition, when the import is completed, you receive an email informing you that the import is completed. This confirmation email also contains links to the results of the import.

<figure><img src="../../../.gitbook/assets/Snyk Code - Imported Repository - Complete - Email Confirmation.png" alt="Snyk Code import confirmation email"><figcaption><p>Snyk Code import confirmation email</p></figcaption></figure>

After the import, Snyk Code continues to monitor the code of the imported repositories in search of vulnerabilities. When a repository is imported, Snyk automatically creates a webhook to the SCM, and this webhook is used after the import for regular scans. By default, Snyk Code analyzes and tests the pull requests in the imported repositories on a weekly basis, and it sends you a **New issues and remediations** summary via email when new vulnerabilities are detected.

{% hint style="info" %}
You can also set a daily test for your imported code or cancel the automatic testing via the **Code analysis** Project settings. In addition, you can retest the imported code on-demand using the **Retest now** option on the **Code Analysis** page.
{% endhint %}
