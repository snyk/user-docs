# Install Snyk Broker - Code Agent using Docker

{% hint style="warning" %}
**Deprecated**

The Code Agent is deprecated and is no longer maintained.

The preferred method of running Snyk Code analysis using Snyk Broker is through [Brokered Code](../../install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker.md). The Code Agent is an alternative method without advantages. For details, contact your Snyk Integration Consultant or Technical Success Manager or contact [Snyk Support](https://support.snyk.io/hc/en-us).

The automatic [PR Checks](../../../../scan-with-snyk/pull-requests/pull-request-checks/) feature is not supported for Snyk Broker - Code Agent.
{% endhint %}

{% hint style="info" %}
For the initial deployment of the Broker Client – Code Agent, collaborating with your Snyk account team is required.
{% endhint %}

The workflow for setting Snyk Broker - Code Agent is as follows:

1. &#x20;[Obtain the required tokens for setup](obtain-the-required-tokens-for-setup.md).
2. If you already have a running Broker client for the same Organization and the same Integration, [Remove the existing Broker Client](remove-an-existing-broker-client.md).
3. [Create an internal network for the Code Agent – Broker client communication](create-network-for-broker-client-and-code-agent-communication.md).
4. [Set up the Code Agent](set-up-the-code-agent.md).
5. [Set up the Broker client](set-up-the-broker-client/).\
   You can set up the Broker Client component with or without a display of the code snippets on the Web UI results. See [Running the Broker Client without code snippets](set-up-the-broker-client/run-the-broker-client-without-the-code-snippet-display.md) and [Running the Broker Client with code snippets](set-up-the-broker-client/run-the-broker-client-with-the-code-snippets-display.md).
6. [Test the Code Agent – Snyk Broker setup](test-the-snyk-broker-code-agent-setup.md).
