# Prepare Snyk Broker for deployment

* [ Broker introduction](/hc/en-us/articles/360015367178-Broker-introduction)
* [ Set up Snyk Broker](/hc/en-us/articles/360015296637-Set-up-Snyk-Broker)
* [ Prepare Snyk Broker for deployment](/hc/en-us/articles/360015296717-Prepare-Snyk-Broker-for-deployment)
* [ Enable permissions for Snyk Broker from your third-party tool](/hc/en-us/articles/360015296737-Enable-permissions-for-Snyk-Broker-from-your-third-party-tool)
* [ Retrieve a unique Broker client token](/hc/en-us/articles/360015367278-Retrieve-a-unique-Broker-client-token)
* [ How to install and configure your Snyk Broker client](/hc/en-us/articles/360015296937-How-to-install-and-configure-your-Snyk-Broker-client)
* [ Upgrade your Snyk Broker client](/hc/en-us/articles/360015367458-Upgrade-your-Snyk-Broker-client)

##  Prepare Snyk Broker for deployment

* Client machine system requirements: 1 CPU, 256MB of Ram.
* Access to the on-prem Git deployment 
* Network access: an outbound TLS \(443\) to [https://broker.snyk.io](https://broker.snyk.io/) that is also allowed by any firewalls installed on your network.
* A Snyk account 
* Self-enabled Broker integration using the Snyk API, or enabled by contacting Snyk support at **support@snyk.io**. 
* A unique UUID token. See [Retrieve a unique Broker client token](https://support.snyk.io/hc/en-us/articles/360015367278-Retrieve-a-unique-Broker-client-token).
* Docker configured to pull images from Docker Hub.

### Prepare hosts for installation

We recommend configuring at least two separate instances of the Broker client for each integration, either on different hosts or installed via a Kubernetes system. This ensures that you always have at least two instances running for redundancy.

* If you use a proxy server, ensure you configure it, and any firewalls, to allow the Broker client inbound and outbound access:
  * * Outbound from broker to broker.snyk.io on port 443.
    * Inbound to broker client at the BROKER\_CLIENT\_URL on the port you have configured \(typically 7341\).
* Traffic initiated from the Snyk server side always uses the latest available Broker connection. All activity from our side \(such as driven by recurring tests\) appears on only one of your replicas at a time. So the amount of Snyk activity is proportional to the activity in your repos \(or Jira\) as that activity generates webhooks, which is distributed across all replicas.  

