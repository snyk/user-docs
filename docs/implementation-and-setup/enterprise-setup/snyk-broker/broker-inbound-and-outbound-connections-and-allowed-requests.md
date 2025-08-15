# Broker inbound and outbound connections and allowed requests

This page provides details about the connections between Snyk and the Broker Client and allowed requests.

## Using inbound and outbound connections with Snyk Broker

The Broker Client runs within your internal network.

### Inbound connection from Snyk to the Broker Client

There is no direct inbound connection from Snyk to the Broker Client.

The Broker Client makes an outbound connection to [https://broker.snyk.io](https://broker.snyk.io). This establishes a WebSocket connection, allowing communication with the Broker Server.

Thus you do not need to allow a Snyk IP address. Instead, you can allow the Broker Client IP/port.

### Outbound connection from the Broker Client

The Broker Client initiates the outbound connection to establish the WebSocket.

After the WebSocket connection initiated by the Broker Client is established, Snyk can send inbound requests to the Broker Client through the WebSocket.

Thus you do not need to allow inbound connections to the Broker Client from Snyk-specific IP addresses or other external IP addresses.

## Allowed requests

### Approved data list for Snyk Broker

The Broker Client maintains an approved data list for inbound and outbound requests. Only data on this approved list may be requested. This narrows the access permissions to the absolute minimum required for Snyk to monitor a repository.

### Inbound requests allowed

For Snyk Open Source, the following requests are allowed:

* Snyk.io is allowed to fetch and view only dependency manifest files and the `.snyk` policy file.
* No other source code is viewed, extracted, or modified.
* You may check in additional `.snyk` files to support the Snyk patch mechanism and for any ignore instructions that are included in your vulnerability policy.

Snyk Code needs access to the entire repository.

See [How Snyk handles your data](../../../snyk-data-and-governance/how-snyk-handles-your-data.md) for more details.

### Outbound requests allowed

When you configure your Broker Client setup, Git repository webhooks are set to enable automatic Snyk scans, triggered when your developers submit new pull requests or merge events.

Webhook notifications are delivered to Snyk through the Broker Client only for events relevant to Snyk actions: push to branch and open pull request, and only when the event data also shows a scan has occurred. For example, for Open Source, the event data must include a dependency manifest file or a `.snyk` policy file.

### Default approved data list and `accept.json` file

On occasion, you may need to [add and configure an `accept.json`](classic-broker/snyk-broker-infrastructure-as-code-detection.md) file in your Broker deployment. Doing this will remove the ability to apply ACCEPT rules when starting the Broker.

\
.
