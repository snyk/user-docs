# Prepare Snyk Broker for deployment

## Prerequisites

* Client machine system requirements: 1 CPU, 256MB of Ram.
* Access to the on-prem Git deployment 
* Network access: an outbound TLS \(443\) to [https://broker.snyk.io](https://broker.snyk.io/) that is also allowed by any firewalls installed on your network.
* A Snyk account 
* Self-enabled Broker integration using the Snyk API, or enabled by contacting Snyk support at **support@snyk.io**. 
* A unique UUID token. See [Retrieve a unique Broker client token](https://snyk.gitbook.io/user-docs/integrations/snyk-broker/retrieve-a-unique-broker-client-token).
* Docker configured to pull images from Docker Hub.

## Prepare hosts for installation

We recommend configuring at least two separate instances of the Broker client for each integration, either on different hosts or installed via a Kubernetes system. This ensures that you always have at least two instances running for redundancy.

## Configure your network

* If you use a proxy server, ensure you configure it, and any firewalls, to allow the Broker client inbound and outbound access:
  * * Outbound from broker to broker.snyk.io on port 443.
    * Inbound to broker client at the BROKER\_CLIENT\_URL on the port you have configured \(typically 7341\).
* Traffic initiated from the Snyk server side always uses the latest available Broker connection. All activity from our side \(such as driven by recurring tests\) appears on only one of your replicas at a time. So the amount of Snyk activity is proportional to the activity in your repos \(or Jira\) as that activity generates webhooks, which is distributed across all replicas.  

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

