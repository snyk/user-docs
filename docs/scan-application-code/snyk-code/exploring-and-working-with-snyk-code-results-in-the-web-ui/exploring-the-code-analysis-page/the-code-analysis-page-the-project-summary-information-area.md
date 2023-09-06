# Code Analysis page Project summary Information

The Project summary information on the **Code Analysis** page includes the following:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Information Area.png" alt="Project summary information"><figcaption><p>Project summary information</p></figcaption></figure>

* **Created** – the date of the repository import.
* **Snapshot taken** – the last time the Project was tested. The results that are displayed on this page are the ones that were found in the last test.\
  By hovering over this option, you can display the exact date of the last test:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Information Area - Last test date - tooltip.png" alt="Date of last test and Retest now option"><figcaption><p>Date of last test and Retest now option</p></figcaption></figure>

Each time a Project is tested, Snyk Code takes a snapshot of the repository in its current state and analyzes it to find vulnerabilities. The tests for which results are displayed on the **Code Analysis** page could be initiated either by Snyk Code or by you:\
Snyk Code performs a test during the initial import process.\
Snyk Code automatically performs recurring tests when a schedule is set for them.\
You can perform a test on demand by clicking the **Retest now** option.

* **Retest now** – this option enables you to perform a manual test of the repository on demand to get its most up-to-date vulnerability results. When you click the **Retest now** option, Snyk Code takes a snapshot of the repository, analyzes its source code files, and displays the new results on the **Code Analysis** page.

You can also use the **Retest now** option to apply the exclusion rules of the `.snyk` file to an imported repository. For more information, see [Excluding directories and files from the import process](../../snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process.md).

Take into consideration that Snyk counts a manual test as a new test. For more information, see [What counts as a test?](https://support.snyk.io/hc/en-us/articles/360000925418-What-counts-as-a-test-)

* **IMPORTED BY** – the SCM username of the user who imported the analyzed repository.
* **PROJECT OWNER** \[Optional] - the name of the user responsible for the Project in your Organization. This feature is designed for the administrative needs of your Organization, and it does not affect the Snyk Project. By default, this field is empty. A Project Owner can only be a user who belongs to the Snyk Organization of the Project.\
  **To set a Project Owner:**\
  Click the **Add a project owner** option. Then, select the required Project Owner from the list of users that are defined in your Snyk Organization:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Information Area - Project Owner - 2.png" alt="Add a project owner option"><figcaption><p>Add a project owner option</p></figcaption></figure>

* **ANALYSIS SUMMARY**:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Information Area - Analysis Summary - 2.png" alt="Analysis summary detail"><figcaption><p>Analysis summary detail</p></figcaption></figure>

* **Analyzed files** - the number of code files analyzed by Snyk Code in the repository and the percentage of the analyzed files out of the total repository code files.
* **Repo breakdown** - the number of analyzed and not analyzed files in the repository, including the types of the files. The **UNANALYZED FILES** category includes text files that were not analyzed by Snyk Code since their language or extension is currently not supported.\
  Note that unsupported **Unknown** files are any files in the repository that are not present in the list of extensions. Unknown files could be multimedia data like pictures and videos, binaries, files with proprietary formats, or any niche format that is not of interest.

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Information Area - Analysis Summary - Repo breakdown - 2.png" alt="Repo breakdown detail"><figcaption><p>Repo breakdown detail</p></figcaption></figure>
