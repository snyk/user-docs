---
nav_context: classic
description: How to scan internal applications with a Snyk API & Web Scanning Agent
---

# Scan internal applications

{% include "../../../../.gitbook/includes/new-navigation-banner.md" %}

## Scan internal applications

Scan your internal applications with the Snyk API & Web Scanning Agent. The Scanning Agent creates an encrypted tunnel between Snyk API & Web and your network, so you do not expose your applications to the internet.

### What is a Scanning Agent for?

A Scanning Agent lets you scan internal applications for vulnerabilities without exposing them to the internet or to Snyk IP addresses. Use it to scan any application reachable only from within your network, including development, staging, pre-release, and internal production applications.&#x20;

A single Scanning Agent can scan multiple internal targets. You can also use several Scanning Agents, each reaching a different part of your network.

### How does a Scanning Agent work?

A Scanning Agent creates an encrypted, authenticated tunnel between Snyk API & Web and your network.

Snyk follows these security principles:

* All code is open source and available in the [Snyk API & Web GitHub repositories](https://github.com/Probely/farcaster-onprem-agent/).
* You have complete control over the Scanning Agent, including the right to change it.
* Snyk cannot access the Scanning Agent.
* The Scanning Agent runs in containers with the least required privileges.
* The Scanning Agent encrypts all traffic end-to-end.
* The Scanning Agent does not open any network ports.

### Install a Scanning Agent

To install a Scanning Agent, visit [Install a Scanning Agent](install-scanning-agent.md). The installation reference and the installer source code are in the [Snyk API & Web GitHub repositories](https://github.com/Probely/farcaster-onprem-agent/).

### Scan a target with a Scanning Agent

After you configure a Scanning Agent and it is running, assign it to targets:

1. In Snyk API & Web, navigate to **Targets**.
2. Find the target in the list and click the gear icon to open its settings.
3. On the **Scanner** tab, in the **Scanning Agent** section, select the Scanning Agent to use.
4. Click **Save**.

To remove the Scanning Agent from a target, click **Unlink**.

To assign or remove a Scanning Agent for multiple targets, select the targets in the targets list, then click **Assign scanning agent** or **Remove agent**.

<figure><img src="../../../../.gitbook/assets/Screenshot 2026-09-28 at 14.04.50.png" alt="Targets list with a target selected and the Assign scanning agent and Remove agent buttons highlighted"><figcaption><p>Bulk-assign or remove a Scanning Agent from the targets list</p></figcaption></figure>

In the targets list, targets that use a Scanning Agent show a cloud icon.

### Scanning Agent status

A Scanning Agent can have one of the following statuses:

| Status                         | Description                                                                                                                                                                                                                                                                                                                                                 |
| ------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Connected** **over** **UDP** | The Scanning Agent has worked within the last 180 seconds.                                                                                                                                                                                                                                                                                                  |
| **Connected** **over** **TCP** | UDP is unavailable on your network, so the tunnel uses TCP. Scans still run. For better throughput, allow outbound UDP on port 443.                                                                                                                                                                                                                         |
| **Disconnected**               | Misconfiguration can cause this status. Check the Scanning Agent configuration and the firewall rules. For more information, visit [Installation](https://github.com/Probely/farcaster-onprem-agent?tab=readme-ov-file#installation) and [Network Requirements](https://github.com/Probely/farcaster-onprem-agent?tab=readme-ov-file#network-requirements). |

The Scanning Agent status appears in the following places:

* the Scanning Agents list
* the Scanning Agent details page
* the targets list
* the target details page
* the Scanning Agent tooltip in target settings

### View Scanning Agent details

To open the details page of a Scanning Agent, click it in the Scanning Agent list. The page shows:

* **Scanning Agent information:** status, last seen time, traffic from the last 24 hours, Agent ID, and so on.
* **Connection stability:** availability, number of disconnections, longest outage, and most recent outage over a date range you choose, with a per-day view. Snyk calculates availability only for the time it monitored the Scanning Agent.
* **Scope:** the teams the Scanning Agent serves.
* **Targets using this agent:** every target assigned to the Scanning Agent, with its last scan and status.

