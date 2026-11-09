---
description: >-
  Install a customer gateway in your network so Continuous Offensive Security
  can scan your internal applications.
---

# Customer Gateway

A customer gateway is a lightweight network component that securely connects Snyk Continuous Offensive Security to your internal infrastructure, so you can scan targets behind firewalls, VPNs, or private networks.

The gateway itself does not use credits. A scan of a target that uses a gateway uses credits like any other scan, except when the scan cannot start or stops because the gateway is disconnected. Visit [Scan a target through a gateway](customer-gateway.md#scan-a-target-through-a-gateway).

Setting up a gateway takes three tasks: create the gateway, start it in your network, and attach it to a target.

{% hint style="info" %}
The `snyk cos gateway` commands are in the preview release channel of the Snyk CLI. To use them, install the latest preview version of the CLI. For more information, visit [Releases and channels for the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/releases-and-channels-for-the-snyk-cli).
{% endhint %}

## Prerequisites

To set up a gateway, you need:

* the Tenant Admin role, to create a gateway (the `tenant.evo.all` permission) and to attach it to a target
* the [Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-the-snyk-cli), if you create or start the gateway from the CLI. Use the [latest preview version](https://docs.snyk.io/developer-tools/snyk-cli/releases-and-channels-for-the-snyk-cli) of the CLI.
* a machine in your network that can reach the Snyk API, the Snyk gateway relays, and the targets you want to scan

The machine needs outbound access to the Snyk API and to the gateway relays on TCP port 443 and UDP port 4500. The gateway uses UDP port 4500 when it is reachable, and otherwise uses TCP port 443. Find your region in the following table, and allow the relay addresses for that region.

| Region     | Snyk API URL             | Relay addresses                                                                                                                                                  |
| ---------- | ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| SNYK-US-01 | `https://api.snyk.io`    | <p><code>r1.relay.snyk.io</code> (34.26.202.58)<br><code>r2.relay.snyk.io</code> (35.231.172.28)<br><code>r3.relay.snyk.io</code> (34.138.193.93)</p>            |
| SNYK-US-02 | `https://api.us.snyk.io` | <p><code>r1.relay.us.snyk.io</code> (18.225.85.178)<br><code>r2.relay.us.snyk.io</code> (3.132.111.218)<br><code>r3.relay.us.snyk.io</code> (3.134.50.174)</p>   |
| SNYK-EU-01 | `https://api.eu.snyk.io` | <p><code>r1.relay.eu.snyk.io</code> (3.121.132.163)<br><code>r2.relay.eu.snyk.io</code> (52.58.48.62)<br><code>r3.relay.eu.snyk.io</code> (3.73.151.88)</p>      |
| SNYK-AU-01 | `https://api.au.snyk.io` | <p><code>r1.relay.au.snyk.io</code> (16.178.78.105)<br><code>r2.relay.au.snyk.io</code> (54.79.139.218)<br><code>r3.relay.au.snyk.io</code> (13.236.171.218)</p> |

## Create a gateway

You create a gateway in Evo.

{% stepper %}
{% step %}
Navigate to **Settings** > **Customer Gateway**.
{% endstep %}

{% step %}
Click **Add gateway**.
{% endstep %}

{% step %}
Enter a name for the gateway and select an Organization.
{% endstep %}

{% step %}
Click **Generate**.
{% endstep %}

{% step %}
In the window that opens, copy and save the information you need to start the gateway.

* Under **Gateway TOKEN**, copy the token.
* To start the gateway from the CLI, copy the CLI command. Visit [Start the gateway in your network](customer-gateway.md#start-the-gateway-in-your-network) for the commands.
* To start the gateway in a container, expand the **Run in Container** section and copy the commands for your method:
  * **DOCKER**: the Docker command to install the gateway, and the Docker command to check the gateway logs.
  * **DOCKER-COMPOSE**: the `docker-compose.yml` manifest for the gateway, and the Docker Compose command to start it.
  * **KUBERNETES**: the Kubernetes commands to create the Snyk namespace, create the gateway token secret, and deploy the gateway pod.
{% endstep %}
{% endstepper %}

To create the gateway from the CLI instead, run the following command. Replace `<gateway name>` with a name for the gateway. Use the [latest preview version](https://docs.snyk.io/developer-tools/snyk-cli/releases-and-channels-for-the-snyk-cli) of the CLI.

```
snyk cos gateway create --name <gateway_name>
```

The command prints the gateway ID, name, state, and the gateway token.

{% hint style="warning" %}
Snyk shows the gateway token only once, when you create the gateway. Save it securely, because you cannot retrieve or change it later. If you lose it, create a new gateway.
{% endhint %}

## Start the gateway in your network

Start the gateway inside your network. You can start it from the CLI with the gateway token, or deploy it with Docker, Docker Compose, or Kubernetes.

To start the gateway from the CLI, run the following commands. Use the [latest preview version](https://docs.snyk.io/developer-tools/snyk-cli/releases-and-channels-for-the-snyk-cli) of the CLI. Replace `<token>` with the gateway token. \
Replace `<api url>` with the Snyk API URL for your region from the table in [Prerequisites](customer-gateway.md#prerequisites). Use the API URL of the account that created the gateway, and do not add `/rest` to the URL.

```bash
export SNYK_GATEWAY_TOKEN=<token>
export SNYK_API=<api url>
snyk cos gateway run
```

Until the gateway connects for the first time, its status is **Disconnected**. After you start the gateway, its status changes to **Connected**.

## Attach the gateway to a target

You need the Tenant Admin role to attach a gateway to a target.

{% stepper %}
{% step %}
In the target wizard, navigate to the optional **Customer gateway** section.
{% endstep %}

{% step %}
Select the gateway from the list. You cannot select a gateway that is **Disconnected**.
{% endstep %}

{% step %}
Save the target.
{% endstep %}
{% endstepper %}

## Scan a target through a gateway

Before every scan of a gateway-routed target, a pre-flight check confirms that the gateway is connected. What happens next depends on the state of the gateway.

| Gateway state                                        | What happens                                                                                                                                                                                                                                                                                                                                                  |
| ---------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Connected**                                        | The scan runs.                                                                                                                                                                                                                                                                                                                                                |
| **Disconnected**, when you start a scan manually     | The **Start** button is disabled. When you hover over it, Evo shows a message that the scan cannot start while the gateway is disconnected. The gateway reconnects automatically. If it does not, run the `snyk cos gateway run` command again. **Start** is enabled again when the gateway is **Connected**. The scan does not start, so it uses no credits. |
| **Disconnected**, when a scheduled or CI scan starts | The pre-flight check fails and the scan does not start. The scan uses no credits.                                                                                                                                                                                                                                                                             |
| Connection lost during a scan                        | The scan stops. The scan uses no credits for the rest of the run.                                                                                                                                                                                                                                                                                             |

## Delete a gateway

You cannot delete a gateway that is attached to a target or used by a running scan.

{% stepper %}
{% step %}
Remove the gateway from every target that uses it, and wait for any running scan that uses the gateway to finish.
{% endstep %}

{% step %}
Stop the process that runs the gateway: the CLI process, the container, or the deployment.
{% endstep %}

{% step %}
Navigate to **Settings** > **Customer Gateway**.
{% endstep %}

{% step %}
Delete the gateway entry.
{% endstep %}
{% endstepper %}

To delete the gateway from the CLI instead, run `snyk cos gateway delete --gateway-id <gateway ID>`. Use the [latest preview version](https://docs.snyk.io/developer-tools/snyk-cli/releases-and-channels-for-the-snyk-cli) of the CLI.

For more details, visit [Before you begin](../targets/before-you-begin.md) to learn how to add a target and what to prepare first, [Scanner IP addresses](scanner-ip-addresses.md) to find the addresses to allow through your firewall or WAF, and [Monitor and Manage Scans](monitor-and-manage-scans.md) to follow scan progress and cancel a scan.
