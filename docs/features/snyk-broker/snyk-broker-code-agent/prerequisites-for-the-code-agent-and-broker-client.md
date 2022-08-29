# Prerequisites for the Code Agent and Broker Client

Before you start the setup process, make sure that your Broker Client and Code Agent machine has the following minimal requirements for running the Broker Client and Code Agent components.

## &#x20;**Before you start**

To use the Broker Client – Code Agent deployment method, the Broker Client and Code Agent components must be enabled by Snyk. To enable the Broker Client and Code Agent components, contact one of the following:

* Your Implementation Consultant (IC)
* Your Technical Success Manager (TCM)
* Snyk Support

**Notes**:

* The maximum file size you can import via the Snyk Broker – Code Agent deployment method is the same as via the CLI - 1MB.\
  For more info, see the [File size limit for Snyk Code analysis](https://docs.snyk.io/products/snyk-code/snyk-code-language-and-framework-support#file-size-limit-for-snyk-code-analysis).
* The machine that hosts the Broker Client and the Code agent must have the ability to run a Docker Container, for example, by using a Docker Desktop or Kubernetes.

## &#x20;Prerequisites for the Code Agent component

The minimal requirements for running the **Code Agent** component are:

* **CPU** - 1 vCPU
* **Memory** - 2Gb
* **Disk space** - 2Gb\
  **Note**: The available disk space determines the maximum size of repositories that are imported simultaneously. If you wish to import repositories that exceed this size, you need to increase the available disk space. However, it is highly recommended to consult with your Implementation Consultant before importing repositories that are larger than 2Gb.
* **Network:**
  * SCM connection - communication to the SCM that stores the repositories you want to analyze.
  * Snyk Code AI Engine connection - outbound communication to the Code Analysis Engine at [https://deeproxy.snyk.io/](https://deeproxy.snyk.io/).
* Internet bandwidth and connection - the upload speed of the source code to the Broker Server will be affected by a low bandwidth and a slow Internet connection.
* **Snyk API token** - your **** Snyk API token is required to authenticate the Code Agent component with your Snyk Account.\
  **Note**: For more information, see [Obtaining your Snyk API token](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-snyk-api-token).

**Note**: Currently, you cannot deploy the Code Agent as a part of a Broker redundancy solution.

## Prerequisites for the Broker Client component

The minimal requirements for running the **Broker Client** component are:

* **CPU** - 1 vCPU
* **RAM** - 256MB
* **Network access** - an outbound TLS (443) communication to the Broker Server at [https://broker.snyk.io](https://broker.snyk.io). This outbound communication should also be allowed by any firewalls installed on your network.\
  **Note:** If you are using the same Broker Client for other Snyk products, and you want to enable there the Automatic PR Checks feature, you also need to configure the following:\
  Inbound communication to the Broker Client host machine at the `BROKER_CLIENT_URL`, on the port you have configured for the host machine (by default, 8000). See page 35.
* **Broker token** – a unique Broker token is required to enable the Snyk Broker deployment for a specific Organization and a specific integrated SCM. \
  **Note**: For more information, see [Obtaining your Broker token](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-broker-token).
* **SCM token** – your integrated SCM token is required to enable access with specific permissions to the SCM. \
  **Note**: For more information, see [Obtaining your SCM token](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-scm-token).&#x20;

