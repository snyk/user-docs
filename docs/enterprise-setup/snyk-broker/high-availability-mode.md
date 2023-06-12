# High availability mode

Snyk Broker can bring high availability capabilities to both servers and clients, thus increasing the scalability of the current Broker initially to support addition of the “git-clone-through-broker” flow for Snyk Code.

It allows several Broker Clients to have separate connections, independent one from another. Snyk platform will spread the requests it makes evenly across the connections so ease load on each client as well as providing true redundancy if one is offline. It also avoids downtime in the fairly infrequent cases when we upgrade the Broker server components.

<figure><img src="../../.gitbook/assets/snyk-broker-ha-mode.png" alt=""><figcaption></figcaption></figure>

To use High Availability, your deployment setup must bring **more than 1 replica**. Whether simply running more than one container, or increasing the replica count in your Kubernetes deployment, you simply need to run more than one container with the exact same parameters.

Today, a maximum of 4 Broker Clients running concurrently is allowed. A 5th tunnel will simply attempt to get a connection allocated indefinitely.

For each container, by default, high availability mode is disabled. To activate it, set the following environment variables as shown either in your container or deployment:

```
BROKER_HA_MODE_ENABLED=true
BROKER_DISPATCHER_BASE_URL=https://api.snyk.io
```

**Helm chart deployments** can set these values by enabling the mode using set arguments. Helm chart version 1.7.0 or later is required.

```
--set highAvailabilityMode.enabled=true
```

Please review the chart values file to bring additional configuration (increasing replica count, updating broker dispatcher base url, etc)



**Important Notes**

* The Dispatcher Base Url should be specific to your region if you are using a regional Snyk platform (i.e api.eu.snyk.io)
* If using proxy with allow list, api.snyk.io (or corresponding api hostname) must be allowed. Preflight checks will indicate failure upon Broker client startup.
* **BROKER\_CLIENT\_URL value must remain unique across all the Broker clients**. \
  It is acceptable for this url resolves to a particular client.\
  \
  The multiple tunnels are primarily supporting Snyk=>You flow. The webhooks going You=>Snyk can take any tunnel as well.\
  \
  Preferably, Load Balancers can also be introduced.\
  Kubernetes deployment with a service in front of each Broker Client will distribute this automatically.&#x20;



The following client log lines will ensure that the high availability mode is active

> ```shell
> ...
> checking for HA mode (enabled=true)
> received server id (serverId=0)
> broker client is connecting to broker server (url=https://broker.snyk.io, serverId=0)
> ...
> ```

Using high availability mode introduces the concept of allocated tunnels for each client, scheduling those tunnels across a predictable set of Broker servers so a unique client can be connected to the right pod.&#x20;
