# Connections with Snyk Broker

The information on this page provides details about the connections between Snyk and the Broker Client and allowed requests.

## Using inbound and outbound connections with Snyk Broker

The Broker Client runs within your internal network.

### Inbound connection from Snyk to the Broker Client

There is no direct inbound connection from Snyk to the Broker Client.

The Broker Client makes an outbound connection to [https://broker.snyk.io](https://broker.snyk.io). This establishes a WebSocket connection, allowing communication with the Broker Server.

Thus there is no need to allowlist a Snyk IP address. Instead, you can allow the Broker Client IP/port.

### Outbound connection from the Broker Client

The Broker Client initiates the outbound connection to establish the WebSocket.

After the WebSocket connection initiated by the Broker Client is established, Snyk can send inbound requests to the Broker Client via the WebSocket.

Thus you do not need to allow inbound connections to the Broker Client from Snyk-specific IP addresses or other external IP addresses.

## **Approved data list for Snyk Broker**

The Broker Client maintains an approved data list for inbound and outbound requests. Only data on this approved list may be requested. This narrows the access permissions to the absolute minimum required for Snyk to monitor a repository.

### Inbound requests allowed

For Snyk Open Source:

* Snyk.io is allowed to fetch and view only dependency manifest files and the `.snyk` policy file.
* No other source code is viewed, extracted, or modified.
* You may check in additional `.snyk` files to support the Snyk patch mechanism and for any ignore instructions that are included in your vulnerability policy.

Snyk Code and Snyk IaC need access to the entire repository.

See [How Snyk handles your data](../../more-info/how-snyk-handles-your-data.md) for more details. For information about using Snyk Broker with Snyk Container and requirements for open source, code analysis, and IaC, see [Using Snyk Broker to scan your code](connections-with-snyk-broker.md#using-snyk-broker-to-scan-your-code).

### Outbound requests allowed

When you configure your Broker Client setup, Git repository webhooks are set to enable automatic Snyk scans, triggered when your developers submit new pull requests or merge events.

Webhook notifications are delivered to Snyk via the Broker Client only for events relevant to Snyk actions: push to branch and open pull request, _**and only when**_ the event data also includes a dependency manifest file or a .`snyk` policy file.

### Default approved data list and `accept.json` file

Because of the limitations of the default approved data list, if you want to scan Infrastructure as Code files with Snyk Broker, you must [add and configure an `accept.json`](snyk-broker-infrastructure-as-code-detection/) file in your Broker deployment.

To learn more about the approved data list and the `accept.json` file, see [Custom approved listing filter](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation#custom-approved-listing-filter).
