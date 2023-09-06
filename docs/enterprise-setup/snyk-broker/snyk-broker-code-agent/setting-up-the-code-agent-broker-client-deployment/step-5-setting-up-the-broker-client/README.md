# Step 5: Setting up the Broker Client

{% hint style="info" %}
**Before you start** – make sure that you currently do not have a running Broker Client for the same Snyk Organization and the same SCM. If you do have a running Broker Client, [stop and remove it](../step-2-removing-an-existing-broker-client.md).
{% endhint %}

Setup of the Broker Client begins with [downloading or updating the Docker image of the Broker Client](step-5.1-downloading-or-updating-the-snyk-broker-client-docker-image.md).

{% hint style="info" %}
The Code Agent is supported only in Broker Client version 4.108.0 and later versions. If you are already using the Docker image of the Broker Client, pull the latest update.
{% endhint %}

Setup continues with running the Broker Client container based on the downloaded Docker image.

You can [run the Broker Client WITHOUT a display of the code snippets in the Snyk Code results on the Web UI](step-5.2a-running-the-broker-client-without-the-code-snippet-display.md):

<figure><img src="../../../../../.gitbook/assets/Broker - Results - without code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (4) (1).png" alt="Broker Client run with no display of code snippets"><figcaption><p>Broker Client run with no display of code snippets</p></figcaption></figure>

Alternatively you can [run the Broker Client WITH a display of the code snippets in the Snyk Code results on the Web UI](step-5.2b-running-the-broker-client-with-the-code-snippets-display.md):

<figure><img src="../../../../../.gitbook/assets/Broker - Results - with code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="Broker Client run with display of code snippets"><figcaption><p>Broker Client run with display of code snippets</p></figcaption></figure>
