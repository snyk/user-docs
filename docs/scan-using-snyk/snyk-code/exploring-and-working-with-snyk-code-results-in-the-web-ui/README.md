# Manage code vulnerabilities

After you import the repositories you want Snyk Code to analyze, you can start exploring and working with the Snyk Code results on the Web UI.

## Locating the Snyk Code results

During the import process, Snyk creates a Target folder for each imported repository. In this Target folder, Snyk creates different Snyk Projects for the files of the imported repository, according to the data type of the files. All the files that contain source code that can be analyzed by Snyk Code are aggregated in the **Code analysis** Project.

The name of the Target folder includes the name of the imported repository, together with the account name of the integrated SCM, the SCM icon, and the number of Snyk Projects that were created for the repository.

<figure><img src="../../../.gitbook/assets/Snyk code - 1.png" alt="Projects list showing Target folder with Code analysis Project and other Projects"><figcaption><p>Projects list showing Target folder with Code analysis Project and other Projects</p></figcaption></figure>

Unlike other Snyk products, which create a separate Snyk Project for each imported file, Snyk Code creates one Snyk Project for all the imported files of one repository. This way, all the vulnerabilities that were detected in the repository code are aggregated in one Project, and the Snyk Code results can present the data flow of a vulnerability issue across multiple files.

## Filtering existing Projects

The **Projects** page on the Web UI includes a filter pane that allows you to filter your Snyk Projects according to different criteria, and provides you with information about the number of existing Snyk Projects that match each criterion:

<figure><img src="../../../.gitbook/assets/Snyk code - 2.png" alt="Filter pane on Projects page"><figcaption><p>Filter pane on Projects page</p></figcaption></figure>

By default, only active Projects with discovered issues are displayed on the **Projects** page.

For Snyk Code, you can choose the following filter criteria on the filter pane:

* **SHOW**:
  * **With issues** and **Without issues** - Projects with or without issues that were discovered by Snyk.
  * **Active and Inactive** - Projects that are either in an **Active** or **Inactive** status in Snyk. Inactive Projects are Projects that were deactivated on the Web UI.\
    For more information, see [Removing imported repositories from Snyk Code testing](../snyk-code-and-your-repositories/removing-imported-repositories-from-snyk-code-testing.md).
* **INTEGRATIONS**: the integrated SCMs that store repositories that were imported to Snyk.

The other options in the filter pane are currently not applicable to Snyk Code.

## Viewing the discovered vulnerability issues in a repository

To view all the vulnerability issues that Snyk Code detected in your imported repository, on the **Projects** page, open the Target folder that contains the Projects of the repository. Then click the **Code analysis** Project:

The **Code Analysis** page opens, displaying the list and details of the vulnerability issues that were discovered in the selected repository. This issue list is organized by the severity level of the discovered issues, by default from the highest to the lowest:

<figure><img src="../../../.gitbook/assets/Snyk code - 4.png" alt="Code Analysis page showing vulnerabiity issues"><figcaption><p>Code Analysis page showing vulnerabiity issues</p></figcaption></figure>
