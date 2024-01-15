# Prepare Snyk Broker for deployment

## Prerequisites for Snyk Broker

{% hint style="info" %}
The use of Snyk Broker on Windows is not supported. Snyk recommends that Windows users deploy Broker using Linux.
{% endhint %}

The following are prerequisites for using Snyk Broker:

* Client machine system requirements: 1 CPU, 256MB of RAM
* Network access: an outbound TLS (443) to [https://broker.snyk.io](https://broker.snyk.io) that is also allowed by any firewalls installed on your network
* A Snyk account
* Self-enabled Broker integration using the Snyk API, or enabled by contacting [Snyk Support](https://support.snyk.io/hc/en-us)
* A unique UUID token called Broker token. See [Generate credentials in the target application for Snyk Broker](prepare-snyk-broker-for-deployment.md#generate-credentials-in-the-target-application-for-snyk-broker)
* An SCM token or password. See the [integrations documentation](../../integrate-with-snyk/) for each SCM for information on how to obtain the token. Snyk Broker does not support authentication with the mTLS method. &#x20;
* Docker configured to pull images from Docker Hub

{% hint style="warning" %}
**Multi-tenant settings for EU and AU**\
When you set up Broker, Code Agent, or both for use in EU or AU Multi-tenant environments, additional environment variables with the specific URLs are required.\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`\
For the URLs, see [EU and AU account datacenter creation](https://docs.snyk.io/snyk-processes/data-residency-at-snyk#eu-and-au-datacenter-account-creation).
{% endhint %}

## Prepare hosts for installation of Snyk Broker

Snyk recommends configuring at least two separate instances of the Broker Client for each integration, either on different hosts or installed via a Kubernetes system. This ensures that you always have at least two instances running for redundancy.

## Configure your network for using Snyk Broker

If you use a proxy server, ensure you configure it, and any firewalls, to allow the Broker Client inbound and outbound access:

* Outbound connection from the Broker Client (running in your environment) to [broker.snyk.io](https://broker.snyk.io) (or [https://broker.eu.snyk.io](https://broker.eu.snyk.io) / [https://broker.au.snyk.io](https://broker.au.snyk.io)) on port 443
* Internal connection that allows inbound access from the integration (SCM, CR) to the Broker Client at the BROKER\_CLIENT\_URL on the port you have configured (typically 8000). This is not inbound from the internet.

Traffic initiated from the Snyk Broker Server side always uses the latest available Broker connection. All activity from the Snyk side (such as traffic driven by recurring tests) appears on only one of your replicas at a time. The amount of Snyk activity is proportional to the activity in the repositories or Jira items. That activity generates webhooks, which are distributed across all replicas.

## **Define your Broker deployment components**

Consider the following to understand what the required components are for your deployment:

* What service are you connecting Broker to?
  * GitHub, Jira, Bitbucket, Harbor, other service
  * See [Snyk Broker - Client integration setups](broken-reference).
* Are you planning to detect Infrastructure as Code files?
  * You will need to add an environment variable `-e ACCEPT_IAC` or a custom allowlist `accept.json` file to your deployment.
  * See [Snyk Broker - Infrastructure as Code detection](snyk-broker-infrastructure-as-code-detection/).
* Are you planning to detect Snyk Code vulnerabilities?
* Are you planning to connect to a Container Registry?
  * You will need to deploy an additional agent with Broker, the Snyk Broker Container Registry Agent.
  * See [Snyk Broker Container Registry agent](snyk-broker-container-registry-agent/).

Every integration has a specific Broker token assigned to it. This means if you want to analyze Snyk Code vulnerabilities and connect to a Container Registry your integration will be:

* One Broker for the SCM with the additional environment variable `-e ACCEPT_CODE` or the custom allowlist `accept.json` and one Broker Code Agent
* One Broker for the Container Registry and one Broker Container Registry agent

## Generate credentials in the target application for Snyk Broker

{% hint style="info" %}
Snyk recommends rotating all API tokens and credentials used with Snyk Broker every 90 days.

For the first deployment of Broker, collaborating with your Snyk account team is required.
{% endhint %}

After generating the credentials for the Broker's target application, configure the environment variables for launching Snyk Broker.

The Broker token is required and must be generated in order for you to use Snyk Broker.

For code repository (SCM) integrations, a Broker token can be generated via API or by contacting [Snyk support](https://support.snyk.io/hc/en-us/requests/new).

1. Go to the Snyk API v1 documentation and follow the example under "Set up a broker for an existing integration" within the [Integrations API](https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration) or contact [Snyk support](https://support.snyk.io/hc/en-us/requests/new).
2. Verify the Broker token is generated in the Snyk Web UI under the specified SCM integration. by selecting **Settings** > **Integrations** for that specific integration update to see the Broker token.

For [Artifactory Repository](../../integrate-with-snyk/package-repository-integrations/artifactory-package-repository-connection-setup/) and [Nexus Repository Manager](../../integrate-with-snyk/package-repository-integrations/nexus-repository-manager-connection-setup/) brokered instances or [Jira](install-and-configure-snyk-broker/jira-install-and-configure-broker/jira-install-and-configure-using-docker.md) integration, you can generate a Broker token in the Snyk UI or contact [Snyk support](https://support.snyk.io/hc/en-us/requests/new).

1. Select **Settings** > **Integrations** for that specific integration to generate the Broker token.
2. Once the Broker token is generated, under the integration, the notification from this screen correctly displays “Could not connect to…”, as you have not yet installed and configured the client.
3. Copy and paste the Broker token from the UI to use it when you install the client.

## Enabling Broker across multiple organizations

You can use the same Git service across multiple Organizations in Snyk with the same Broker token. To do this, create the token for an Organization and then create a new Organization based on the original. This clones the token and you can now enable the Broker for it.

To do this retroactively for existing Organizations, you can use the API v1 endpoint [Clone an integration (with settings and credentials)](https://snyk.docs.apiary.io/#reference/integrations/integration-cloning) to clone a specific integration, including the Broker token.

Unless you do this, you must generate a new Broker token for the Organization, as each integration and Organization have their own unique Broker token.
