# View vulnerabilities in your code

## **View the Snyk Code results**

During the process of importing your selected repositories, Snyk Code automatically analyzes your imported code to find potential vulnerabilities. All the vulnerability findings that Snyk Code detects in the code of one imported repository are aggregated in one Snyk Project, called Code analysis:

<figure><img src="../../../.gitbook/assets/SnykCode1.png" alt="ode analysis Project"><figcaption><p>Code analysis Project</p></figcaption></figure>

Unlike other Snyk products, which create a separate Snyk Project for each imported file, Snyk Code creates one Snyk Project for all the imported files of one repository. All the vulnerabilities detected in the code in the repository are aggregated in one Project, and the Snyk Code results can present the data flow of a vulnerability across multiple files.

To view all the security vulnerabilities that Snyk Code detected in your imported code, click the **Code analysis** Project, and explore the details of each vulnerability:

<figure><img src="../../../.gitbook/assets/SnykCode2.png" alt="Vulnerabilties in the Code Analysis Project"><figcaption><p>Vulnerabilties in the Code Analysis Project</p></figcaption></figure>

If your Snyk Account is already integrated with your SCM and contains imported repositories, Snyk Code may already be active in your Organization settings and running. In this case, you can check to see if your existing repositories are already being tested by Snyk Code by searching for the **Code analysis** Project in your Target folders. If a **Code analysis** Project exists in your imported repositories, you can proceed to the page [Exploring and working with the Snyk Code results](./).

You may find this information about repositories and Snyk Code useful as you begin to work with Snyk Code results:

* [Importing additional repositories to Snyk](broken-reference)
* [Excluding directories and files from the import process](../../start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process.md)
* [Removing imported repositories from the Snyk Code test](../../start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/removing-imported-repositories-from-snyk-code-testing.md)
