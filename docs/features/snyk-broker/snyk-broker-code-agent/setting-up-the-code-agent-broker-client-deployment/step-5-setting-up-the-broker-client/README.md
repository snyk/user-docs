# Step 5: Setting up the Broker Client

{% hint style="info" %}
**Before you start** – make sure that you currently do not have a running Broker Client for the same Snyk Organization and the same SCM. If you do have a running Broker Client, [stop and remove it](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/Xtgu2HdNoafUskHqoeDW/).
{% endhint %}

The setup of the Broker Client consists of two steps:

1\.  [Downloading or updating the Docker image of the Broker Client](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.1-downloading-or-updating-the-snyk-broker-client-docker-image).\
**Note:** The Code Agent is only supported in the Broker Client version 4.108.0 and later versions. If you are already using the Docker image of the Broker Client, pull the latest update.

2\.  Running the Broker Client container based on the downloaded Docker image.

When running the Broker Client container, you have two options:

* [Running the Broker Client WITHOUT a display of the code snippets in the Snyk Code results on the Web UI](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.2a-running-the-broker-client-without-the-code-snippet-display):&#x20;

<figure><img src="../../../../../.gitbook/assets/Broker - Results - without code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (4).png" alt=""><figcaption></figcaption></figure>

* [Running the Broker Client WITH a display of the code snippets in the Snyk Code results on the Web UI](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.2b-running-the-broker-client-with-the-code-snippets-display):&#x20;

<figure><img src="../../../../../.gitbook/assets/Broker - Results - with code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (4).png" alt=""><figcaption></figcaption></figure>

&#x20;****&#x20;
