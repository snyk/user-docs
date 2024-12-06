# Test the Snyk Broker Code Agent setup

{% hint style="warning" %}
**Deprecated**

The Code Agent is deprecated and is no longer maintained.

The preferred method of running Snyk Code analysis using Snyk Broker is through [Brokered Code](https://docs.snyk.io/enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker). The Code Agent is an alternative method without advantages. For details, contact your Snyk Integration Consultant or Technical Success Manager or contact [Snyk Support](https://support.snyk.io).

The automatic [PR Checks](https://docs.snyk.io/scan-with-snyk/pull-requests/pull-request-checks) feature is not supported for Snyk Broker - Code Agent.
{% endhint %}

To test the Code Agent with Snyk Broker setup, import a repository to Snyk and verify that you receive the Code Analysis Project and results. When the selected repositories are imported to Snyk, and the Snyk Code results are displayed, the Code Agent with Snyk Broker setup is operating successfully.

To test the setup, follow these steps:

1\. On the Snyk Web UI, open the Organization where you want to test the Code Agent.

2\. When the required Organization is open, select **Add project**. Then select the SCM for which you set the Code Agent:

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Selecting SCM for import.png" alt="Select Add Project and the SCM"><figcaption><p>Select Add Project and the SCM</p></figcaption></figure>

3\. On the **Personal and Organization repositories** page, select the repositories you want Snyk Code to analyze, and click the **Add selected repositories** button:

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Selecting repos for import.png" alt="Select repositories for Snyk Code to analyze"><figcaption><p>Select repositories for Snyk Code to analyze</p></figcaption></figure>

The selected repositories are imported to Snyk Code, and a progress bar appears on the **Projects** page. When the import is completed, a message appears on the **Projects** page confirming the success of the import. Your imported repositories appear as Target folders, each containing the **Code analysis** Project that includes the findings of the Snyk Code test:

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Code Analysis Project.png" alt="Confirmation message and target folders"><figcaption><p>Confirmation message and target folders</p></figcaption></figure>

4\. To view the Snyk Code results, select the **Code analysis** Project.

The **Code Analysis** page opens, displaying the list and details of the vulnerability issues that were discovered in the selected repository:

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Code Analysis page.png" alt="Code Analysis page showing issues"><figcaption><p>Code Analysis page showing issues</p></figcaption></figure>

5\. To view the details of a vulnerability issue, select **Full details**.

The details of the issue appear, and depending on the way you set up the Broker client, the code snippets either appear in the **Data flow** tab or not:

This example shows Snyk Code results with code snippets.

<figure><img src="../../../../.gitbook/assets/Broker - Results - with code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="Broker Client run with display of code snippets"><figcaption><p>Broker Client run with display of code snippets</p></figcaption></figure>

This example shows Snyk Code results without code snippets:

<figure><img src="../../../../.gitbook/assets/Broker - Results - without code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (4) (1).png" alt="Broker Client run with no display of code snippets"><figcaption><p>Broker Client run with no display of code snippets</p></figcaption></figure>

For more information on how to troubleshoot the Snyk Broker Code Agent setup, see [Troubleshooting Broker](../../troubleshooting-broker.md).
