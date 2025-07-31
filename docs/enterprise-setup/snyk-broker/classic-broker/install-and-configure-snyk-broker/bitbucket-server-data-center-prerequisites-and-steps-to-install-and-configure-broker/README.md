# Bitbucket Server/Data Center - prerequisites and steps to install and configure Broker

{% hint style="info" %}
**Feature availability**

PR Checks for **Bitbucket Server** integrations require Bitbucket Server version 7.4 and above, or Bitbucket Data Center version 8 or above. \
\
When using a brokered connection, Snyk Broker version 4.206 and above is required.
{% endhint %}

* Review the general instructions for the installation method you plan to use, [Helm](../install-and-configure-broker-using-helm.md) or [Docker](../install-and-configure-broker-using-docker.md).
* Before installing the Bitbucket Server/Data Center Broker, ensure your Snyk account team provides you with a Broker token.&#x20;
* Docker or an equivalent method is required to run Docker Linux containers. Some Docker setups for Windows only support Windows containers. Enure your deployment can run Linux containers.

After you meet all the prerequisites, you can continue with the steps to install using [Docker](bitbucket-server-data-center-install-and-configure-using-docker.md) or [Helm](bitbucket-server-data-center-install-and-configure-using-helm.md).
