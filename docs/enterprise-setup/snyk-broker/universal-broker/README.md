# Universal Broker

{% hint style="info" %}
**Release status**\
Universal Broker is in Early Access and is available only with Enterprise plans. If you want to set it up, contact your Snyk account team.
{% endhint %}

The Universal Broker provides improvements to facilitate the management of Snyk Broker deployments and connections.&#x20;

This page explains the configuration flow for the client. These pages guide you through a Broker deployment setup using the Universal Broker:

* [Initial configuration of the Universal Broker](initial-configuration-of-the-universal-broker.md)
* [Set up a GitHub connection using the API](set-up-a-github-connection-using-the-api.md)
* [Restart your Broker with the required environment variable and connect](restart-your-broker-with-the-required-environment-variable-and-connect.md)

## Universal Broker Deployment and Connection model <a href="#universal-broker-deployment-and-connection-model" id="universal-broker-deployment-and-connection-model"></a>

The Universal Broker separates deployment and container concerns from connection concerns. It allows for a smaller or a single deployment to support numerous connections of varied types.

<figure><img src="../../../.gitbook/assets/image 5.png" alt="Universal Broker Model"><figcaption><p>Universal Broker Model</p></figcaption></figure>

In contrast to the existing Broker model, where each deployment and container supports only one connection type, Snyk is providing Broker deployments configured to support many connections in a set-and-forget approach.

Many connections into a single deployment require more configuration parameters to load for each connection. To make this manageable at scale, Snyk has decoupled deployments from connections.

Broker operators now declare their desired deployment model before running any Broker client. By specifying what Broker connections to support, the existing approach `org->integrations->broker connections` is evolving to be instead `Broker Connections -> integration/org`.

Deployments are roughly equivalent to running containers or a replicaset (on Kubernetes deployment models). Deployments aim to represent the entity running the Broker code on a server.

Connections represent a set of individual connections provided by the Universal Broker deployment to a specific downstream system (SCM, Jira, and so on) with a given set of credentials for each.

A single deployment can support any number of connections of various types. Although the Universal Broker may provide new capability, Snyk recommends keeping the number below 25 connections, because container resources are not infinitely vertically scalable.

This model does not store credentials; it uses a credentials reference or references. These indicate to the Broker client that the credentials supporting a given connection are expected to be found in a specifically named environment variable.

For example, deployments run a connection (type `github`) using the environment variable name defined in the associated credentials reference (`$MY_CRED_REF`).

## Steps in implementing the Universal Broker

{% hint style="warning" %}
Prerequisite: You must be tenant admin to be able to create deployments, credentials reference(s) and connection(s).
{% endhint %}

The high-level steps in implementing the Unversal Broker follow.

1. **One time:** Install the Snyk Broker App in your Organization. This returns an install ID, a client ID, and a client secret, all needed to interact with the Snyk platform. The Organization ID is required to create the deployment.
2. **One time:** Define a deployment for your tenant ID and install ID.
3. **One time:** Define credentials references needed for your connections.
4. **One time:** Define your desired connection or connections.
5. **One time for each Organization integration:** Configure the Organizations and integrations that should use the connection.

<figure><img src="../../../.gitbook/assets/image 6.png" alt="Deployment, connections, and organizations"><figcaption><p>Deployment, connections, and organizations</p></figcaption></figure>

After defining the desired connections, run the Broker Client. The integration between connections and Organizations and integrations is then done independently from the Universal Broker deployment, reducing the required activities on the running containers or Kubernetes deployment.&#x20;

By using the `snyk-broker-config` command or by making an HTTP request against the relevant API endpoint, you can integrate connections or disconnect from a given Organization. The Broker client will manage the connections based on what the Broker administrators have declared as the desired state according to this model. It may take up to two minutes to pick up the changes in a deployment.&#x20;
