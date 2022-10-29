# Step 6: Testing the Code Agent – Snyk Broker Setup

To test the Code Agent – Snyk Broker deployment setup, import a repository to Snyk, and verify that you are receiving the **Code Analysis** Project and results. When the selected repositories are imported to Snyk, and the Snyk Code results are displayed, it means that the Code Agent – Snyk Broker deployment is operating successfully.

**To test the Code Agent – Snyk Broker deployment:**

1\. On the Snyk Web UI, open the required Organization:

<figure><img src="../../../../.gitbook/assets/Snyk Broker - Organization - Select.png" alt=""><figcaption></figcaption></figure>

2\. Once the required Organization is open, click the **Add project** button. Then, select the SCM for which you set the Code Agent:

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Selecting SCM for import.png" alt=""><figcaption></figcaption></figure>

3\. On the **Personal and Organization repositories** page, select the repositories you want Snyk Code to analyze, and click the **Add selected repositories** button:

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Selecting repos for import.png" alt=""><figcaption></figcaption></figure>

The selected repositories are imported to Snyk Code, and a progress bar appears on the **Projects** page. When the import is completed, a confirmation message appears on the **Projects** page, informing you of the success of the import. Your imported repositories appear as Target folders, each containing the **Code analysis** Project that includes the findings of the Snyk Code test:

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Code Analysis Project.png" alt=""><figcaption></figcaption></figure>

4\. To view the Snyk Code results, click the **Code analysis** Project.

The **Code Analysis** page appears, displaying the list and details of the vulnerability issues that were discovered in the selected repository:

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Code Analysis page.png" alt=""><figcaption></figcaption></figure>

5\. To view the details of a certain vulnerability issue, click the **Full details** button.

The details of the issue appear, and depending on the way you set up the Broker Client, the code snippets either appear in the **Data flow** tab or not:

* Snyk Code results with code snippets:

<figure><img src="../../../../.gitbook/assets/Broker - Results - with code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1).png" alt=""><figcaption></figcaption></figure>

* Snyk Code results without code snippets:

<figure><img src="../../../../.gitbook/assets/Broker - Results - without code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (2).png" alt=""><figcaption></figcaption></figure>

For more information on how to troubleshoot the Snyk Broker - Code Agent deployment, see [Troubleshooting Broker](../../troubleshooting-broker.md).
