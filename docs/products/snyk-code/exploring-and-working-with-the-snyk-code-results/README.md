# Exploring and working with the Snyk Code results on the Web UI

After you imported to Snyk the repositories you want Snyk Code to analyze, you can start exploring and working with the Snyk Code results on the Web UI.

### Locating the Snyk Code results

During the import process, Snyk creates a Target folder for each imported repository. In this Target folder, Snyk creates different Snyk Projects for the files of the imported repository, according to the data type of the files. All the files that contain source code that can be analyzed by Snyk Code, are aggregated in the **Code analysis** Project:

**Note**: The name of the Target folder includes the name of the imported repository, together with the account name of the integrated SCM, the SCM icon, and the number of Snyk Projects that were created for the repository.

![ ](<../../../.gitbook/assets/Snyk Code - Results - General diagram.png>)

**Note**: Unlike other Snyk products, which create a separate Snyk Project for each imported file, Snyk Code creates one Snyk Project for all the imported files of one repository. This way, all the vulnerabilities that were detected in the repository code are aggregated in one Project, and the Snyk Code results can present the data flow of a vulnerability issue across multiple files.&#x20;

### Filtering Existing Projects

The **Projects** page on the Web UI includes a Filter pane. This Filter pane allows you to filter your Snyk Projects according to different criteria, and it provides you with information about the number of existing Snyk Projects that match each criterion:

![](<../../../.gitbook/assets/Snyk Code - Results - Projects page - Filter pane.png>)

By default, only active Projects with discovered issues are displayed on the **Projects** page.

For Snyk Code, you can use the following Filter criteria:

* **SHOW**:
  * **With issues** and **Without issues** - Projects with or without issues that were discovered by Snyk.
  * **Active and Inactive** - **** Projects that are either in an **Active** or **Inactive** status in Snyk. Inactive Projects are Projects that were deactivated on the Web UI.\
    **Note**: For more information, see [Deactivating and deleting the Snyk Code Project](https://docs.snyk.io/products/snyk-code/getting-started-with-snyk-code/activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/removing-imported-repositories-from-the-snyk-code-test).
* **INTEGRATIONS**: the integrated SCMs that store repositories that were imported to Snyk.

**Note**: The other options in the Filter pane are currently not applicable to Snyk Code. &#x20;

### Viewing the discovered vulnerability issues in a repository

**To view all the vulnerability issues that Snyk Code detected in your imported repository:**

* On the **Projects** page, open the Target folder that contains the Projects of the required repository. Then, click the **Code analysis** Project:

![](<../../../.gitbook/assets/Snyk Code - Results - Selecting the Code analysis Project.png>)

The **Code Analysis** page appears, displaying the list and details of the vulnerability issues that were discovered in the selected repository. This issue list is organized by the severity level of the discovered issues, by default from the highest to the lowest:

![ ](<../../../.gitbook/assets/Snyk Code - Getting Started - Code Analysis Project (1) (1).png>)

&#x20;
