# Project summary

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

You can also use the **Retest now** option to apply the exclusion rules of the `.snyk` file to an imported repository. For more information, see [Excluding directories and files from the import process](../../../start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process.md).

Take into consideration that Snyk counts a manual test as a new test. For more information, see [What counts as a test?](https://support.snyk.io/hc/en-us/articles/360000925418-What-counts-as-a-test-)

* **IMPORTED BY** – the SCM username of the user who imported the analyzed repository.
* **PROJECT OWNER** \[Optional] - the name of the user responsible for the Project in your Organization. This feature is designed for the administrative needs of your Organization, and it does not affect the Snyk Project. By default, this field is empty. A Project Owner can only be a user who belongs to the Snyk Organization of the Project.\
  **To set a Project Owner:**\
  Click the **Add a project owner** option. Then, select the required Project Owner from the list of users that are defined in your Snyk Organization:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Information Area - Project Owner - 2.png" alt="Add a project owner option"><figcaption><p>Add a project owner option</p></figcaption></figure>

## **Analysis summary**

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Information Area - Analysis Summary - 2.png" alt="Analysis summary detail"><figcaption><p>Analysis summary detail</p></figcaption></figure>

* **Analyzed files** - the number of code files analyzed by Snyk Code in the repository and the percentage of the analyzed files out of the total repository code files.
* **Repo breakdown** - the number of analyzed and not analyzed files in the repository, including the types of the files. The **UNANALYZED FILES** category includes text files that were not analyzed by Snyk Code since their language or extension is currently not supported.\
  Note that unsupported **Unknown** files are any files in the repository that are not present in the list of extensions. Unknown files could be multimedia data like pictures and videos, binaries, files with proprietary formats, or any niche format that is not of interest.

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Information Area - Analysis Summary - Repo breakdown - 2.png" alt="Repo breakdown detail"><figcaption><p>Repo breakdown detail</p></figcaption></figure>

## Page header

The header of the **Code Analysis** page includes the repository details, Project name, and Project tabs:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Header - With Callouts.png" alt="Components of Code Analysis page header"><figcaption><p>Components of Code Analysis page header</p></figcaption></figure>

* **Repository details** – details on the imported repository for which Code analysis is reported on this page:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Header - Repo Details.png" alt="Repository details"><figcaption><p>Repository details</p></figcaption></figure>

The repository details include the following:

* **SCM icon** - the icon of the integrated SCM that stores the analyzed repository, in this example, GitHub icon.
* **Repository name** - the name of the analyzed repository. In this example, **juice-shop**.

The repository name in the header is a link to the repository in the integrated SCM:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Header - repository name link (1).png" alt="Click the repository name"><figcaption><p>Click the repository name</p></figcaption></figure>

By clicking the repository name on the Header, you can open the repository in the SCM directly from Snyk:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Header - Integrated SCM.png" alt="Repository in the SCM"><figcaption><p>Repository in the SCM</p></figcaption></figure>

The page shows the name of the branch of the analyzed repository, in this example, **master.**

* **Project Name** – the name of the Snyk Project that is displayed on this page. In Snyk Code, the name of the Project is always **Code Analysis** for every analyzed repository:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Header - Project Name.png" alt="Code Analysis name"><figcaption><p>Code Analysis name</p></figcaption></figure>

* **Project tabs** – tabs that open different UI pages:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Header - Project Tabs.png" alt="Project tabs"><figcaption><p>Project tabs</p></figcaption></figure>

* **Overview** – the default Project page that displays the most recent vulnerability results.
* **History** – a page that displays the previous snapshots and results of past tests. For more information, see [Exploring the history of the Snyk Code results](../exploring-the-history-of-snyk-code-results.md).
* **Settings** – a page that enables you to perform the following:
  * Set the frequency of recurring tests.
  * Find the unique Project ID.
  * [Deactivate the Project](../../../start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/removing-imported-repositories-from-snyk-code-testing.md#deactivating-and-deleting-imported-repositories).
  * [Delete the Project](../../../start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/removing-imported-repositories-from-snyk-code-testing.md#deactivating-and-deleting-the-snyk-code-project).

<figure><img src="../../../../.gitbook/assets/Snyk Code - Results - Header - Project Settings page.png" alt="Code Analysis settings"><figcaption><p>Code Analysis settings</p></figcaption></figure>

## Filters

The Filter pane on the **Code Analysis** page enables you to filter the discovered vulnerabilities issues according to different criteria. It provides you with information about the number of existing issues that match each criterion:

<figure><img src="../../../../.gitbook/assets/Snyk Code - Main UI Features - Filter.png" alt="Filter pane on Code Analysis page"><figcaption><p>Filter pane on Code Analysis page</p></figcaption></figure>

* **SEVERITY LEVEL** - display only issues with a certain severity level.\
  **Note**: Snyk Code uses the **High**, **Medium**, and **Low** severity levels.
* **PRIORITY SCORE** - display only issues in a certain priority score range.\
  **Note**: For more information, see [Understanding the Priority Score of the Snyk Code issues](priority-score-in-snyk-code.md).
* **STATUS** - display either **Open** issues or issues that were **Ignored**.
* **LANGUAGES** - display only issues that were discovered in code files that were written in a specific language.\
  Only programming languages discovered in the analyzed repository are displayed in the Filter pane.
* **VULNERABILITY TYPES** - display only issues with a certain Vulnerability Type.\
  For more information about the Vulnerability Types, see [Security Rules used by Snyk Code](../../../../scan-application-code/snyk-code/exploring-and-working-with-snyk-code-results-in-the-web-ui/exploring-the-code-analysis-page/broken-reference/).
