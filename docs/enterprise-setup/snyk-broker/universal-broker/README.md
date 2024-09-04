# Universal Broker

The Universal Broker provides improvements to facilitate the management of Snyk Broker deployments and connections.&#x20;

This page explains the configuration flow for the client. Getting started pages follow to guide you through a Broker deployment setup using the Universal Broker.

## Universal Broker Deployment and Connection model <a href="#universal-broker-deployment-and-connection-model" id="universal-broker-deployment-and-connection-model"></a>

The Universal Broker separates deployment and container concerns from connection concerns. It allows for a smaller or a single deployment to support numerous connections of varied types.

<figure><img src="../../../.gitbook/assets/image (548).png" alt="Universal Broker Model" width="437"><figcaption><p>Universal Broker Model</p></figcaption></figure>

In contrast to the legacy Broker model, where each deployment and container supports only one connection type, Snyk now provides Broker deployments configured to support many connections in a  set-and-forget approach.

Many connections into a single deployment require more configuration parameters to load for each connection. To make this manageable at scale, Snyk has decoupled deployments from connections.

Broker operators now declare their desired deployment model before running any Broker client. By specifying what Broker connections to support, the classic existing approach `org->integrations->broker connections` is evolving to be instead `Broker Connections -> integration/org`.

Deployments are roughly equivalent to running containers or a replicaset (on Kubernetes deployment models). Deployments aim to represent the entity running the Broker code on a server.

Connections represent a set of individual connections provided by the Universal Broker deployment to a specific downstream system (SCM, Jira, and so on) with a given set of credentials for each.

A single deployment can support any number of connections of various types. Although the Universal Broker may provide new capability, Snyk recommends keeping the number below 25 connections, because container resources are not infinitely vertically scalable.

Credentials are not stored in this model, which instead uses credentials reference(s). These indicate to the Broker client that the credentials supporting a given connection are expected to be found in a specifically-named environment variable.

For example, deployments run a connection (type `github`), using the environment variable name defined in the associated credentials reference (`$MY_CRED_REF`).

## Steps in implementing the Universal Broker

The high-level steps in implementing the Unversal Broker follow.

1. **One time:** Install the Snyk Broker App into an Organization under your Tenant. This returns an install ID, a client ID, and a client secret, all needed to interact with the Snyk platform. The org ID is required to create the deployment.
2. **One time:** Define a deployment for your tenant ID and install ID.
3. **One time:** Define credentials references needed for your connections.
4. **One time:** Define your desired connection(s).
5. **One time for each organization integration:** Configure the organizations and integrations that should use the connection.\


<figure><img src="../../../.gitbook/assets/image (550).png" alt="Deployment, connections, and organizations"><figcaption><p>Deployment, connections, and organizations</p></figcaption></figure>

After defining the desired connections, run the Broker Client. The integration between connections and organizations and integrations is then done independently from the the Universal Broker deployment, reducing the required activities on the running containers or Kubernetes deployment.&#x20;

By using the `snyk-broker-config` command or by making an HTTP request against the relevant API endpoint, you can integrate connections or disconnect from a given organization. The Broker client will manage the connections based on what the Broker administrator(s) have declared as the desired state according to this model. It may take up to two minutes to pick up the changes in a deployment.
