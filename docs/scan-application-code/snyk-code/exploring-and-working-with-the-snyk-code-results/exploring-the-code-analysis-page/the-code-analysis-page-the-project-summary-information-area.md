# The Code Analysis page - the Project Summary Information area

The Project Summary Information area of the **Code Analysis** page includes the following:

![](<../../../../.gitbook/assets/Snyk Code - Results - Information Area.png>)

* **Created** – the date of the repository import.
* **Snapshot taken** – the last time the Project was tested. The results that are displayed on this page are the ones that were found in the last test.\
  By hovering over this option, you can display the exact date of the last test:

![](<../../../../.gitbook/assets/Snyk Code - Results - Information Area - Last test date - tooltip.png>)

Each time a Project is tested, Snyk Code takes a snapshot of the repository in its current state, and analyzes it in search for vulnerabilities. The tests whose results are displayed in the **Code Analysis** page, could be initiated either by Snyk Code or by you:\
\- Snyk Code performs a test during the initial import process.\
\-  Snyk Code automatically performs recurring tests when a schedule is set for them. \
\- You can perform a test on-demand, by clicking the **Retest now** option (see below).

* **Retest now** – this option enables you to perform a manual test of the repository on-demand, in order to get its most up-to-date vulnerability results. When clicking the **Retest now** option, Snyk Code takes a snapshot of the repository, analyzes its source code files, and displays the new results on the **Code Analysis** page.

**Notes**:\
\- You can also use the **Retest now** option to apply the exclusion rules of the .snyk file to an imported repository. For more information, [see Excluding directories and files from the import process](https://docs.snyk.io/products/snyk-code/getting-started-with-snyk-code/activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/excluding-directories-and-files-from-the-import-process).\
\- You should take into consideration that a manual test is counted by Snyk as a new test. For more information, see [What counts as a test?](https://support.snyk.io/hc/en-us/articles/360000925418-What-counts-as-a-test-).

* **IMPORTED BY** – the SCM username of the user who imported the analyzed repository.&#x20;
* **PROJECT OWNER** \[Optional] - the name of the user who is responsible for the Project in your organization. This feature is designed for the administrative needs of your organization, and it does not have any effect on the Snyk Project. By default, this field is empty. A Project Owner can only be a user who belongs to the Snyk Organization of the Project.\
  \
  \- **To set a Project Owner**:\
  Click the **Add a project owner** option. Then, select the required Project Owner from the list of users that are defined in your Snyk Organization:

![](<../../../../.gitbook/assets/Snyk Code - Results - Information Area - Project Owner - 2.png>)

* **ANALYSIS SUMMARY**:

![](<../../../../.gitbook/assets/Snyk Code - Results - Information Area - Analysis Summary - 2.png>)

* **Analyzed files** - the number of code files that were analyzed by Snyk Code in the specific repository, and the percentage of the analyzed files out of the sum total of the repository code files.
* **Repo breakdown** - the number of analyzed and not analyzed files in the repository, including the types of the files. The **UNANALYZED FILES** category includes text files that were not analyzed by Snyk Code since their language or extension are currently not supported.\
  Note that unsupported **Unknown** files are any files in the repository that are not present in the  list of extensions. Unknown files could be multimedia data like pictures and videos, or binaries, or files with proprietary formats, or any niche format that is not of interest.

![](<../../../../.gitbook/assets/Snyk Code - Results - Information Area - Analysis Summary - Repo breakdown - 2.png>)

&#x20;
