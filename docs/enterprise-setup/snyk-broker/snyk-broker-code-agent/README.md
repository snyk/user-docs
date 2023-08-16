# Snyk Broker - Code Agent

{% hint style="info" %}
The preferred method of running Snyk Code analysis through the Snyk Broker is through [Brokered Code](../install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker.md).  The Code Agent is an alternative method without advantages. However, if you would like to set it up in your Organization, contact your Snyk Integration Consultant or Technical Success Manager or contact [Snyk Support](https://support.snyk.io/hc/en-us).
{% endhint %}

To connect Snyk Code to your self-hosted Git server via the Snyk Broker, you must add the Code Agent component to the Snyk Broker deployment structure. By using the Code Agent component with the Snyk Broker, you can scan repositories that are stored on your self-hosted Git provider and apply the Snyk Code analysis to these repositories in order to find, prioritize, and fix potential vulnerabilities in your source code.

For more information on the Snyk Broker deployment method, see [Snyk Broker](../).

The [Automatic PR Checks feature](../../../snyk-admin/snyk-broker/snyk-broker-code-agent/broken-reference/) is currently not supported in the Snyk Broker - Code Agent deployment method.

To apply Snyk Code analysis to repositories that are stored on your self-hosted Git server, you need the following **components**:

* **Broker Server** - a Server that is running on the Snyk SaaS backend\
  The Broker Server is provided by Snyk, and no installation is required on your part.
* **Broker Client** - a [Docker image](https://hub.docker.com/r/snyk/broker/) that is deployed in your infrastructure\
  For more information, see [Setting up the Broker Client](setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/).
* **Code Agent** - another [Docker image](https://hub.docker.com/r/snyk/code-agent/) that is deployed in your infrastructure\
  For more information, see [Setting up the Code Agent](setting-up-the-code-agent-broker-client-deployment/step-4-setting-up-the-code-agent/).\
  The Code Agent is supported only in Snyk Broker version 4.108.0 and later versions. If you already have a running Broker Client, you must update it by pulling the latest Docker image. For more information, see [Downloading or Updating the Snyk Broker Client – Docker image](setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.1-downloading-or-updating-the-snyk-broker-client-docker-image.md).

The **Broker Client** and **Code Agent** components are deployed in your infrastructure, creating two separate services. Together with the Broker Server, the Snyk Code AI Engine, and the Snyk Web UI, these components are responsible for the following Code Analysis workflow:

1\. On the Snyk Web UI, a request is initiated by a user to import a repository from a self-hosted Git server to Snyk for Code Analysis.\
The request can also be initiated from the Snyk API v1, by using the [Import targets request](https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets).

2\. The request arrives at the Broker Server, which is hosted by Snyk. The Broker Server sends the request to the Broker Client, which sends it to the Code Agent.\
The Broker Client automatically provides the Code Agent with the connection details to the integrated SCM, which stores the required repositories.

3\. The Code Agent connects to the integrated SCM, and clones the local repository in a secured manner in your infrastructure. The cloned repository is stored temporarily on the Code Agent container.\
The cloning is performed via HTTPS connection. If your SCM does not support HTTPS, you can work around this with a reverse proxy. For more details reach out to your technical contact at Snyk.

4\. The Code Agent filters the cloned repository for supported files and sends them to the Snyk Code AI Engine.

5\. The [Snyk Code AI Engine](https://docs.snyk.io/products/snyk-code/introducing-snyk-code/key-features/ai-engine) analyzes the code files in search of vulnerability issues. The analysis results are sent back to the Snyk Web UI.\
The cloned files are cached for up to 24 hours before they are deleted. For more information, see [How Snyk handles your data](../../../more-info/how-snyk-handles-your-data.md).

<figure><img src="../../../.gitbook/assets/Code Agent - diagram - new - 4.png" alt="Snyk Code Analysis workflow with Broker"><figcaption><p>Snyk Code Analysis workflow with Broker</p></figcaption></figure>
