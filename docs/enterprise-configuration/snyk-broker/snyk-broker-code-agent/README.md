# Snyk Broker - Code Agent

{% hint style="info" %}
<mark style="color:red;">**The Code Agent is deprecated**</mark>. The preferred method of running Snyk Code analysis using Snyk Broker is through [Brokered Code](../install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker.md).  The Code Agent is an alternative method without advantages. For details, contact your Snyk Integration Consultant or Technical Success Manager or contact [Snyk Support](https://support.snyk.io/hc/en-us).

The automatic [PR Checks](../../../scan-with-snyk/run-pr-checks/) feature is not supported ifor Snyk Broker - Code Agent.
{% endhint %}

To connect Snyk Code to your self-hosted Git server using Snyk Broker, you can add the Code Agent after installing Snyk Broker for your SCM.&#x20;

The Code Agent is available as a [Docker image](https://hub.docker.com/r/snyk/code-agent/). The Code Agent is supported only for Snyk Broker version 4.108.0 and later versions. If you already have a running Broker Client, you must update it by pulling the latest Docker image.

Deploying the Broker client and Code Agent creates two separate services. that, along with the Broker server, Snyk Code AI Engine, and Snyk Web UI, the following Code Analysis workflow:

1. On the Snyk Web UI, a user initiates a request to import a repository from a self-hosted Git server to Snyk for Code Analysis. The request can also be initiated from the Snyk API v1, by using the [Import targets request](https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets).
2. The request arrives at the Broker server, which sends the request to the Broker client, which sends the request to the Code Agent. The Broker client automatically provides the Code Agent with the connection details to the integrated SCM, which stores the required repositories.
3. The Code Agent connects to the integrated SCM, and securely clones the local repository in your infrastructure. The cloned repository is stored temporarily in the Code Agent container. The cloning is done over an HTTPS connection. If your SCM does not support HTTPS, you use a reverse proxy. For details reach out to your technical contact at Snyk.
4. The Code Agent filters the cloned repository for supported files and sends them to the Snyk Code AI Engine.
5. The Snyk Code AI Engine analyzes the code files in search of vulnerability issues. The analysis results are sent back to the Snyk Web UI. The cloned files are cached according to the Cloud provider's storage minimum policy.

<figure><img src="../../../.gitbook/assets/Code Agent - diagram - new - 4.png" alt="Snyk Code Analysis workflow with Broker"><figcaption><p>Snyk Code Analysis workflow with Broker</p></figcaption></figure>
