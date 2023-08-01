# Removing imported repositories from the Snyk Code test

If you do not want Snyk Code to continue testing one or more of your imported repositories, you can:

* Remove the entire repository from your Snyk Account, by one of the following ways:\
  **Note**: If you remove the entire repository from your Account, your repository will no longer be analyzed be any of the Snyk products.
  * Deactivate the repository.
  * Delete the repository from your Snyk Account.
* Remove only the **Code analysis** Project from your Snyk Account, by one of the following ways:\
  **Note**: If you remove only the **Code analysis** Project, other Snyk products that are enabled in your Account will continue to analyze the imported repository.
  * Deactivate the Project.
  * Delete the Project from your Snyk Account.

### **Exploring the Repository Removal Methods**

To select the right method for you for removing repositories from the Snyk Code test, take into consideration what will happen in each of the following actions:

* **Deactivating an imported repository** **will**:
  * Remove the webhook from Snyk to the SCM repository.
  * Disable Pull Request tests for new vulnerabilities.
  * Disable the Fix Pull Requests option from being opened for newly discovered vulnerabilities.
  * Disable recurring tests - email alerts about newly discovered vulnerabilities will be turned off.
* **Deleting a Snyk Project or an imported repository will:**
  * Delete the entire Project or repository and all its historical snapshot data from Snyk.
  * Remove the webhook from the SCM repository.\
    **Note**: Deleting a Snyk Project or an imported repository will not have any effect on your source code.

**Note**: If you want to remove specific directories or files from the Snyk Code test, use [the Exclude option in .snyk file](excluding-directories-and-files-from-the-import-process.md).

### **Deactivating and Deleting Imported Repositories**

To deactivate or delete an imported repository, you need to select all the Projects of the repository, and then use a bulk action to either deactivate or delete all the Projects at once. If you want to entirely delete a repository from your Snyk Account, after you delete its Projects, you also need to delete its Target folder.

**Note**: If after the deletion of a repository you will want to re-import it to Snyk, you will need to refresh the **Projects** page to view the re-import results.

**To deactivate or delete multiple Snyk Projects:**

1\. On the Snyk Web UI, select <img src="../../../.gitbook/assets/Org Settings button - Icon (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (6).png" alt="" data-size="line"> **Settings** > **Usage**:

<figure><img src="../../../.gitbook/assets/image (295) (1).png" alt=""><figcaption></figcaption></figure>

3\. Scroll down the **Usage** page until you reach the **Projects** section, where all your Snyk Projects are listed:

![](<../../../.gitbook/assets/Snyk Code - Org Settings - Usage - Projects section.png>)

4\. On the **Projects** list, select the checkboxes of all the Projects you want to remove:

![](<../../../.gitbook/assets/Snyk Code - Org Settings - Usage - Projects section - Selection.png>)

5\. Click the **Bulk actions** button above the Project list, and select either **Delete** or **Deactivate**:

![](<../../../.gitbook/assets/Snyk Code - Org Settings - Usage - Projects section - Bulk actions.png>)

Depending on your selected action, one of the following occurs:

* **Delete** – a confirmation message appears. In this case, continue to the next step.
* **Deactivate** – your selected Projects are deactivated, and they will no longer be tested by Snyk. If later you will want to reactivate the Projects, repeat this procedure, and select **Activate** from the **Bulk actions** menu.

6\. If you selected to delete Projects, a confirmation message appears, asking you to confirm the deletion. Click **OK** to delete the Projects:

![](<../../../.gitbook/assets/Snyk Code - Org Settings - Usage - Delete Projects - Confirmation message.png>)

7\. To remove from Snyk the Target folder of the repository whose Projects you deleted, return to the **Projects** page. Then, open the Target folder, and click the **Delete** button on the right corner:

<figure><img src="../../../.gitbook/assets/image (176) (1) (1) (1) (1) (1) (2).png" alt=""><figcaption></figcaption></figure>

The entire repository is now deleted from your Snyk Account.

### **Deactivating and deleting the Snyk Code Project**

To stop Snyk Code from testing an imported repository, you can either deactivate or delete the **Code analysis** Project of the repository. After the **Code analysis** Project will no longer be active in the repository, Snyk Code will stop testing the repository, but other Snyk products will continue to scan the repository files.

**To deactivate or delete the Code analysis Project:**

1\. On the **Projects** page, locate the repository you want Snyk Code to stop testing. In its Target folder, locate the **Code analysis** Project, and click the **Settings** button on the right side of its row:

<figure><img src="../../../.gitbook/assets/image (404) (1).png" alt=""><figcaption></figcaption></figure>

2\. On the **Settings** page of the **Code analysis** Project, click either the **Deactivate project** or **Delete project** button according to your needs:

<figure><img src="../../../.gitbook/assets/image (449).png" alt=""><figcaption></figcaption></figure>

The **Code analysis** Project you selected is either deactivated or deleted according to the action you selected, and its repository will no longer be tested by Snyk Code.

**Notes**: If you want Snyk Code to resume its testing after you deleted or deactivated the **Code analysis** Project of a repository, perform the following:

* **Code analysis deletion** - re-import the repository to Snyk, and then refresh the **Projects** page to view the re-import results.
* **Code analysis deactivation** – re-activate the **Code analysis** Project via the **Settings** page of the Project. After you deactivate a Project, the **Deactivate project** button changes to **Activate project**, and a new **Activate** button appears at the top of the page. Click one of these buttons to re-activate the Project:

<figure><img src="../../../.gitbook/assets/image (86).png" alt=""><figcaption></figcaption></figure>
