# Prepare Snyk Broker for deployment

{% hint style="info" %}
Using Snyk Broker on Windows is not supported. Snyk recommends that Windows users deploy Broker using Linux.
{% endhint %}

## Prerequisites for Snyk Broker

When you set up Broker for region other than SNYK-US-01, before you can authenticate you must set environment variables with specific [Broker URLs](../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`

The following are prerequisites for using Snyk Broker in any environment:

* Client machine system requirements: one CPU, 256MB of RAM
* Network access: an outbound TLS (443) to https://broker.snyk.io AND https://api.snyk.io (or regional equivalent) that is also allowed by any firewalls installed on your network
* A Snyk account
* Self-enabled Broker integration through the Snyk API or Broker enabled by contacting [Snyk Support](https://support.snyk.io)
* A unique UUID token called a Broker token. See [Obtain the tokens required to set up a Snyk Broker](../../../../../enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md).
* An SCM token or password. See [Obtain the tokens required to set up a Snyk Broker](../../../../../enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md). Snyk Broker does not support authentication with the mTLS method.
* Docker configured to pull images from Docker Hub to install with Docker.

## Prepare hosts for installation of Snyk Broker

Snyk recommends configuring at least two separate instances of the Broker Client for each integration, either on different hosts or installed using a Kubernetes system. This ensures that you always have at least two instances running for redundancy.

## Configure your network for using Snyk Broker

If you use a proxy server, ensure you configure it and any firewalls to allow the Broker Client inbound and outbound access as follows:

* Outbound connection from the Broker Client running in your environment to https://broker.snyk.io (or https://broker.eu.snyk.io / https://broker.au.snyk.io / https://broker.us.snyk.io) AND https://api.snyk.io (or https://api.eu.snyk.io / https://api.au.snyk.io / https://api.us.snyk.io) on port 443.
* Internal connection that allows inbound access from the integration (SCM, CR) to the Broker Client at the BROKER\_CLIENT\_URL on the port you have configured (typically 8000). This is not inbound from the internet.

Traffic initiated from the Snyk Broker Server side always uses the latest available Broker connection. All activity from the Snyk side, such as traffic driven by recurring tests, appears on only one of your replicas at a time. The amount of Snyk activity is proportional to the activity in the repositories or Jira items. That activity generates webhooks, which are distributed across all replicas.

## **Define your Broker deployment components**

Consider the following to understand what the required components are for your deployment:

* What service are you connecting the Broker to?
  * GitHub, Jira, Bitbucket, Harbor, or other service
  * See [Install and configure Snyk Broker](../install-and-configure-snyk-broker/).
* Are you planning to detect Infrastructure as Code files?
  * You must add an environment variable `-e ACCEPT_IAC` or a custom allowlist `accept.json` file to your deployment.
  * See [Snyk Broker - Infrastructure as Code detection](../snyk-broker-infrastructure-as-code-detection.md).
* Are you planning to detect Snyk Code vulnerabilities? See [Clone an integration across your Snyk Organizatons](../clone-an-integration-across-your-snyk-organizations.md).
* Are you planning to connect to a Container Registry?
  * You will need to deploy an additional agent with the Broker, the Snyk Broker Container Registry Agent.
  * See [Snyk Broker Container Registry agent](../../snyk-broker-container-registry-agent/).

Every integration has a specific Broker token assigned to it. An integration to analyze Snyk Code vulnerabilities and connect to a Container Registry has the following:

* One Broker for the SCM with the additional environment variable `-e ACCEPT_CODE`
* One Broker for the Container Registry and one Broker Container Registry agent

Continue by [obtaining the tokens](../../../../../enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md) required for your deployment.
