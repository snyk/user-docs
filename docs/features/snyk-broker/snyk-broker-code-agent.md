# Snyk Broker - Code Agent

{% hint style="info" %}
This feature is currently in Beta. If you would like to set it up in your organization, contact your Snyk Integration Consultant/Technical Success Manager or Snyk Support.
{% endhint %}

To connect Snyk Code to your self-hosted Git server via the Snyk Broker, you need to add the Code Agent component to the Snyk Broker deployment structure. By using the Code Agent component with the Snyk Broker, you can scan repositories that are stored on your self-hosted Git provider, and apply the Snyk Code analysis to these repositories in order to find, prioritize, and fix potential vulnerabilities in your source code.

**Notes**:

* For more information on the Snyk Broker deployment method, see [Snyk Broker](file:///o/-M4tdxG8qotLgGZnLpFR/s/-MdwVZ6HOZriajCf5nXH/\~/changes/9e7OiBaxzcjfjAJSqS6d/features/snyk-broker).
* The [Automatic PR Checks feature](https://docs.snyk.io/products/snyk-code/using-automatic-pr-checks-for-securing-your-source-code) is currently not supported in the Snyk Broker - Code Agent deployment method.

## Snyk Code - deployment components for accessing a self-hosted Git provider

To apply the Snyk Code analysis to repositories that are stored on your self-hosted Git server, you need the following components:

* **Broker Server –** a Server that is running on the Snyk SaaS backend.\
  **Note:** The Broker Server is provided by Snyk, and no installation is required on your part.
* **Broker Client -** a [Docker image](https://hub.docker.com/r/snyk/broker/) that is deployed in your infrastructure.\
  **Note**: For more information, see page 32.
* **Code Agent -** another [Docker image](https://hub.docker.com/r/snyk/code-agent/) that is deployed in your infrastructure.

**Note:** The Code Agent is only supported in Snyk Broker version 4.108.0 and later versions. If you already have a running Broker Client, you need to update it by pulling the latest Docker image. For more information, see page 33.

The **Broker Client** and **Code Agent** components are deployed in your infrastructure, creating two separate services. Together with the Broker Server, the Snyk Code AI Engine, and the Snyk Web UI, these components are responsible for the following Code Analysis workflow:

1\.  On the Snyk Web UI, a request is initiated by a user to import a repository from a self-hosted Git server to Snyk for Code Analysis.\
**Note**: The request can also be initiated from the Snyk API, by using the [Import targets request](https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets).

2\.  The request arrives to the Broker Server, which is hosted by Snyk. The Broker Server sends the request to the Broker Client, which sends it to the Code Agent.\
**Note**: The Broker Client automatically provides the Code Agent with the connection details to the integrated SCM, which stores the required repositories.

3\.  The Code Agent connects to the integrated SCM, and clones the local repository in a secured manner in your infrastructure. The cloned repository is stored temporarily on the Code Agent container.\
**Note**: The cloning is performed via HTTPS connection.

4\.  The Code Agent filters the cloned repository for supported files, and sends them to the Snyk Code AI Engine.

5\.  The [Snyk Code AI Engine](https://docs.snyk.io/products/snyk-code/introducing-snyk-code/key-features/ai-engine) analyzes the code files in search of vulnerability issues. The analysis results are sent back to the Snyk Web UI.\
**Note:** The cloned files are cached for up to 24 hours before they are deleted. For more info, see [How Snyk handles your data](https://docs.snyk.io/more-info/how-snyk-handles-your-data).

<figure><img src="../../.gitbook/assets/Code Agent - diagram - new - 4.png" alt=""><figcaption></figcaption></figure>

&#x20;&#x20;
