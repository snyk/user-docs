# Prerequisites for the Code Agent and Broker Client

Before you start the setup process, make sure that your Broker Client and Code Agent machine has the following minimal requirements for running the Broker Client and Code Agent components.

To use the Broker Client – Code Agent deployment method, the Broker Client and Code Agent components must be enabled by Snyk. To enable the Broker Client and Code Agent components, contact one of the following:

* Your Implementation Consultant (IC)
* Your Technical Success Manager (TSM)
* [Snyk Support](https://support.snyk.io/hc/en-us)

The maximum file size you can import via the Snyk Broker – Code Agent deployment method is the same as via the CLI - 1MB. For more information, see the [File size limit for Snyk Code analysis](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/introduction-to-snyk-supported-languages-and-frameworks#file-size-limit-for-snyk-code-analysis).

The machine that hosts the Broker Client and the Code agent must have the ability to run a Docker Container, for example, by using a Docker Desktop or Kubernetes.

{% hint style="warning" %}
**Multi-tenant settings for EU and AU**\
When you set up Broker, Code Agent, or both for use in EU or AU Multi-tenant environments, additional environment variables with the specific URLs are required.\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`\
For the URLs, see [EU and AU account datacenter creation](https://docs.snyk.io/snyk-processes/data-residency-at-snyk#eu-and-au-datacenter-account-creation).
{% endhint %}

## Prerequisites for the Code Agent component

The minimal requirements for running the **Code Agent** component are:

* **CPU** - 1 vCPU
* **Memory** - 2Gb
* **Disk space** - 2Gb\
  The available disk space determines the maximum size of repositories that are imported simultaneously. If you wish to import repositories that exceed this size, you must increase the available disk space. However, it is highly recommended to consult with your Implementation Consultant before importing repositories that are larger than 2Gb.
* **Network:**
  * SCM connection - HTTPS communication to the SCM that stores the repositories you want to analyze. Support for HTTP-only SCM-deployments can be resolved by deploying a reverse proxy between the Code Agent and the SCM.
  * Snyk Code AI Engine connection - outbound communication to the Code Analysis Engine at [https://deeproxy.snyk.io/](https://deeproxy.snyk.io/).
* Internet bandwidth and connection - the upload speed of the source code to the Broker Server will be affected by a low bandwidth and a slow Internet connection.
* **Snyk API token** - your Snyk API token is required to authenticate the Code Agent component with your Snyk Account. For more information see [Obtaining your Snyk API token](../../../getting-started/how-to-obtain-and-authenticate-with-your-snyk-api-token.md).

{% hint style="info" %}
Currently you cannot deploy the Code Agent as a part of a Broker redundancy solution.
{% endhint %}

## Prerequisites for the Broker Client component

The minimal requirements for running the **Broker Client** component are:

* **CPU** - 1 vCPU
* **RAM** - 256MB
* **Network access** - an outbound TLS (443) communication to the Broker Server at [https://broker.snyk.io](https://broker.snyk.io). This outbound communication should also be allowed by any firewalls installed on your network.\
  If you are using the same Broker Client for other Snyk products, and you want to enable there the Automatic PR Checks feature, you must also configure the following:\
  An internal connection that allows inbound access from the integration (SCM) to the Broker Client at the BROKER\_CLIENT\_URL on the port you have configured (typically 8000). This is not inbound from the internet. See [Running the Broker Client](setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.2a-running-the-broker-client-without-the-code-snippet-display.md).
* **Broker token** – a unique Broker token is required to enable the Snyk Broker deployment for a specific Organization and a specific integrated SCM. For more information see [Obtaining your Broker token](setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-broker-token.md).
* **SCM token** – your integrated SCM token is required to enable access with specific permissions to the SCM. For more information, see [Obtaining your SCM token](setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-scm-token.md).
