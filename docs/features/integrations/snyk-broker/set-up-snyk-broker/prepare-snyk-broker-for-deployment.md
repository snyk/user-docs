# Prepare Snyk Broker for deployment

## Prerequisites

* Client machine system requirements: 1 CPU, 256MB of Ram.
* Network access: an outbound TLS (443) to [https://broker.snyk.io](https://broker.snyk.io) that is also allowed by any firewalls installed on your network.
* A Snyk account
* Self-enabled Broker integration using the Snyk API, or enabled by contacting Snyk support at **support@snyk.io**.
* A unique UUID token. See [Retrieve a unique Broker client token](https://docs.snyk.io/integrations/snyk-broker/retrieve-a-unique-broker-client-token).
* Docker configured to pull images from Docker Hub.

## Prepare hosts for installation

We recommend configuring at least two separate instances of the Broker client for each integration, either on different hosts or installed via a Kubernetes system. This ensures that you always have at least two instances running for redundancy.

## Configure your network

When using a proxy server, ensure configuration of any firewalls allows outbound traffic from the Broker client [broker.snyk.io](https://broker.snyk.io) on port 443.

Traffic initiated from Snyk's server side always uses the latest available Broker connection. All activity from our side (such as traffic driven by recurring tests) appears on only one of your replicas at a time. The amount of Snyk activity is proportional to the activity in the repos (or Jira items)--that activity generates webhooks which are distributed across all replicas.
