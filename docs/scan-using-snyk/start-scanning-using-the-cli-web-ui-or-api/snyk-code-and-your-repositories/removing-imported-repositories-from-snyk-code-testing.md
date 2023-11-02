# Remove imported repository

If you do not want Snyk to continue testing one or more of your imported repositories, you can do one of the following:

* Remove the entire repository from your Snyk Account in one of the following ways:
  * Deactivate the repository.
  * Delete the repository from your Snyk Account.

{% hint style="warning" %}
If you remove the entire repository from your Account, your repository will no longer be analyzed by any of the Snyk products.
{% endhint %}

* Remove only the **Code analysis** Project from your Snyk Account in one of the following ways:
  * Deactivate the Project.
  * Delete the Project from your Snyk Account.

{% hint style="warning" %}
If you remove only the **Code analysis** Project, other Snyk products that are enabled in your account will continue to analyze the imported repository.
{% endhint %}

## **Remove imported repository methods**

To select the right method for you for removing repositories from Snyk testing, consider what will happen in each of the following actions:

* Deactivating an imported repository will:
  * Remove the webhook from Snyk to the SCM repository.
  * Disable pull request tests for new vulnerabilities.
  * Disable the Fix Pull Requests option from being opened for newly discovered vulnerabilities
  * Disable recurring tests; email alerts about newly discovered vulnerabilities will be turned off.
* Deleting a Snyk Project or an imported repository will:
  * Delete the entire Project or repository and all its historical snapshot data from Snyk.
  * Remove the webhook from the SCM repository.

{% hint style="info" %}
Deleting a Snyk Project or an imported repository will not have any effect on your source code.

If you want to remove specific directories or files from the Snyk Code test, use [the exclude option in the `.snyk` file](excluding-directories-and-files-from-the-import-process.md).
{% endhint %}

## **Deactivate and delete imported repositories**

To deactivate or delete an imported repository, you must select all the Projects in the repository, and then use a bulk action to either deactivate or delete all the Projects at once. If you want to entirely delete a repository from your Snyk account, after you delete its Projects, you must also delete its Target folder.

{% hint style="info" %}
If after deleting a repository you will want to re-import it to Snyk, you must refresh the **Projects** page to view the results from the re-import.
{% endhint %}

Follow these steps to deactivate or delete multiple Snyk Projects:

1\. On the Snyk Web UI, select <img src="../../../.gitbook/assets/Org Settings button - Icon (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (6) (3).png" alt="Settings icon" data-size="line"> **Settings** > **Usage**:

<figure><img src="../../../.gitbook/assets/image (295) (1).png" alt="Usage page"><figcaption><p>Usage page</p></figcaption></figure>

3\. Scroll down the **Usage** page until you reach the **Projects** section, where all your Snyk Projects are listed:

<figure><img src="../../../.gitbook/assets/Snyk Code - Org Settings - Usage - Projects section.png" alt="Projects section on Usage page"><figcaption><p>Projects section on Usage page</p></figcaption></figure>

4\. On the **Projects** list, select the checkboxes of all the Projects you want to remove:

<figure><img src="../../../.gitbook/assets/Snyk Code - Org Settings - Usage - Projects section - Selection.png" alt="Checkboxes selected for Projects to be removed"><figcaption><p>Checkboxes selected for Projects to be removed</p></figcaption></figure>

5\. From the **Bulk actions** dropdown, select either **Delete** or **Deactivate**:

<figure><img src="../../../.gitbook/assets/Snyk Code - Org Settings - Usage - Projects section - Bulk actions.png" alt="Bulk actions options"><figcaption><p>Bulk actions options</p></figcaption></figure>

Depending on your selected action, one of the following occurs:

* **Delete**: A confirmation message appears. In this case, continue to the next step.
* **Deactivate**: Your selected Projects are deactivated, and they will no longer be tested by Snyk. If later you will want to reactivate the Projects, repeat this procedure, and select **Activate** from the **Bulk actions** dropdown

6\. If you selected **Delete**, a confirmation message appears, asking you to confirm the deletion. Click **OK** to delete the Projects:

<figure><img src="../../../.gitbook/assets/Snyk Code - Org Settings - Usage - Delete Projects - Confirmation message (1).png" alt="Conformation message to delete Projects"><figcaption><p>Confirmation message to delete Projects</p></figcaption></figure>

7\. To remove from Snyk the Target folder of the repository whose Projects you deleted, return to the **Projects** page. Then open the Target folder, and click the **Delete** button:

<figure><img src="../../../.gitbook/assets/image (176) (1) (1) (1) (1) (1) (2) (1).png" alt="Selected Target folder and Delete button"><figcaption><p>Selected Target folder and Delete button</p></figcaption></figure>

The entire repository is now deleted from your Snyk account.

## **Deactivate and delete a Snyk Code Project**

To stop Snyk Code from testing an imported repository, you can either deactivate or delete the **Code analysis** Project in the repository. The **Code analysis** Project will no longer be active in the repository and Snyk Code will stop testing the repository, but other Snyk products will continue to scan the repository files.

Follow these steps to deactivate or delete the Code analysis Project:

1\. On the **Projects** page, locate the repository you want Snyk Code to stop testing. In its Target folder, locate the **Code analysis** Project, and click the **Settings** button at the end of its row:

<figure><img src="../../../.gitbook/assets/image (404) (1) (1).png" alt="Settings button for Code analysis Proejct"><figcaption><p>Settings button for Code analysis Proejct</p></figcaption></figure>

2\. On the **Settings** page of the **Code analysis** Project, click either the **Deactivate project** or **Delete project** button according to your needs:

<figure><img src="../../../.gitbook/assets/image (449).png" alt="Deactivate project and Delete project buttons on Code analysis Project Settings page"><figcaption><p>Deactivate project and Delete project buttons on Code analysis Project Settings page</p></figcaption></figure>

The **Code analysis** Project you selected is either deactivated or deleted, and its repository will no longer be tested by Snyk Code.

If you want Snyk Code to resume its testing after you delete or deactivate the **Code analysis** Project of a repository, do the following:

* After deleting the Code analysis Project, re-import the repository to Snyk and then refresh the **Projects** page to view the results of the re-import.
* **After deactivating the Code analysis Project,** re-activate the **Code analysis** Project via the **Settings** page of the Project. After you deactivate a Project, the **Deactivate project** button changes to **Activate project**, and a new **Activate** button appears at the top of the page. Click one of these buttons to re-activate the Project:

<figure><img src="../../../.gitbook/assets/image (86).png" alt="Activate project button on Code analysis Project Settings page"><figcaption><p>Activate project button on Code analysis Project Settings page</p></figcaption></figure>
