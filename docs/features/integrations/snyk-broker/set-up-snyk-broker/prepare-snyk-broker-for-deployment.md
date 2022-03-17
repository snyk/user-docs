# Prepare Snyk Broker for deployment

## Prerequisites

* Client machine system requirements: 1 CPU, 256MB of Ram.
* Network access: an outbound TLS (443) to [https://broker.snyk.io](https://broker.snyk.io) that is also allowed by any firewalls installed on your network.
* A Snyk account
* Self-enabled Broker integration using the Snyk API, or enabled by contacting Snyk support at **support@snyk.io**.
* A unique UUID token called Broker token. See[ Retrieve a unique Broker client token](prepare-snyk-broker-for-deployment.md#generate-credentials-in-the-target-application-for-snyk-broker).
* Docker configured to pull images from Docker Hub.

## Prepare hosts for installation

We recommend configuring at least two separate instances of the Broker client for each integration, either on different hosts or installed via a Kubernetes system. This ensures that you always have at least two instances running for redundancy.

## Configure your network

If you use a proxy server, ensure you configure it, and any firewalls, to allow the Broker client inbound and outbound access:&#x20;

* Outbound from the broker client (ran in your environment) to [broker.snyk.io](https://broker.snyk.io) on port 443.&#x20;
* Inbound to broker client at the BROKER\_CLIENT\_URL on the port you have configured (typically 8000).

Traffic initiated from Snyk's server side always uses the latest available Broker connection. All activity from our side (such as traffic driven by recurring tests) appears on only one of your replicas at a time. The amount of Snyk activity is proportional to the activity in the repos (or Jira items)--that activity generates webhooks which are distributed across all replicas.

## **Define your deployment components**

Understand what are the required components for your deployment:

* What are you connecting the broker to?
  * GitHub, Jira, Bitbucket, Harbor...
* Are you planning to detect Infrastructure as Code files?
  * You will need to add an `accept.json` file to your deployment
  * See [Snyk broker Infrastructure as Code detection](../snyk-broker-infrastructure-as-code-detection/)
* Are you planning to detect Snyk Code vulnerabilities?
  * You will need to deploy an additional agent with the broker, the Snyk Broker Code Agent
  * See [Snyk Broker Code Agent](../snyk-broker-code-agent.md)
* Are you planning to connect to a Container Registry?
  * You will need to deploy an additional agent with the broker, the Snyk Broker Container Registry Agent
  * See [Snyk Broker Container Registry agent](../snyk-broker-container-registry-agent/)

Every integration has a specific Broker token assigned to it, this means if you want to analyse Snyk Code vulnerabilities and connect to a Container Registry your integration will be:

* 1 broker for the SCM (with additional `accept.json`) and 1 broker Code Agent
* 1 broker for the Container Registry and 1 broker Container Registry agent

## Generate credentials in the target application for Snyk Broker

{% hint style="info" %}
For the first deployment of broker, collaborating with your CSM or Support team is required
{% endhint %}

After generating the credentials for the Broker's target application, configure the environment variables for launching the Broker.

In order to use the Broker, the Broker token is required and must be generated.

For code repository (SCM) integrations, a Broker token can be generated via API , or by contacting Snyk Support

1. Go to our API documentation and follow the example under "Set up a broker for an existing integration" within the [Integrations API](https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration) or by contacting support.
2. Verify the Broker token is generated in the Snyk UI under the specified SCM integration. by clicking on settings ![cog\_icon.png](../../../../.gitbook/assets/cog\_icon.png) > **Integrations** for that specific integration update to see the Broker token.

For [Artifactory](https://docs.snyk.io/integrations/private-registry-integrations/artifactory-registry-for-npm) or [Jira](https://docs.snyk.io/features/integrations/notifications-ticketing-system-integrations/jira) integrations, a Broker token can be generated within the Snyk UI, or by contacting Snyk support

1. Click on settings ![cog\_icon.png](../../../../.gitbook/assets/cog\_icon.png) > **Integrations** for that specific integration to generate the Broker token.
2. Once the broker token is generated, under the integration, you will see the notification from this screen correctly displays “Could not connect to…”.) as you have not yet installed and configured the client,
3. Copy and paste the broker token from the UI to use it when you install the client.

## Enabling across multiple organizations

You can use the same Git across multiple organizations in Snyk, using the same Broker token. To do this, create the token for an organization, then create a new organization based on the original; this clones the token and you can now enable the Broker for it.

To do this retroactively for existing organizations, you can use the API to clone a specific integration, including the Broker token: [https://snyk.docs.apiary.io/#reference/integrations/integration-cloning](https://snyk.docs.apiary.io/#reference/integrations/integration-cloning)

If this is not followed, you will need a new Broker token for the organization as each integration and organization have their own unique Broker token
