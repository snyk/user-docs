---
nav_context: agnostic
description: The scanner IP addresses used by Snyk API and Web
---

# Scanner IP address

Snyk API & Web uses the following IP addresses to make requests to your targets and to receive out-of-band connections from them.

## Target scans

Customers hosted in US infrastructure:

* **18.235.241.170** (AWS IP)

Customers hosted in AU/APJ infrastructure:

* **52.65.214.19** (AWS IP)
* **13.237.213.25** (AWS IP)

All other customers:

* **52.19.40.38** (AWS IP)

## Asset discovery

Customers hosted in US infrastructure:

* **44.205.45.120** (AWS IP)

Customers hosted in AU/APJ infrastructure:

* **3.104.172.219** (AWS IP)
* **13.211.189.220** (AWS IP)

All other customers:

* **52.16.191.244** (AWS IP)

If you are unsure where your account is hosted, check all IP addresses or contact the support team for assistance.

If you have your own single-tenant or dedicated infrastructure, contact the support team.

## Out-of-band vulnerability checks <a href="#out-of-band-vulnerability-checks" id="out-of-band-vulnerability-checks"></a>

Some vulnerabilities, such as Log4Shell, can only be confirmed when the affected system connects to an external server. Snyk operates dedicated servers to receive these callbacks. Unlike the addresses above, these servers receive connections from your systems rather than sending requests to them.

Customers hosted in US infrastructure:

* **szpus.prbly.win** resolving to **52.72.180.55** (AWS IP)

Customers hosted in AU/APJ infrastructure:

* **szpau.prbly.win** resolving to **52.62.50.85** (AWS IP)

All other customers:

* **szp.prbly.win** resolving to **52.17.201.157** (AWS IP)<br>

These servers accept connections on TCP ports 53, 80, 443, and 389, and on UDP port 53.

Allow both the hostname and the IP address. Because the connection originates inside your own network and carries test payloads, some firewalls and threat intelligence tools categorize the domain as malicious and redirect it. For example, Palo Alto DNS sinkholing resolves it to sinkhole.paloaltonetworks.com. Allowing the IP address alone does not resolve this; the domain must also be allowed.

These connections are optional. If you block them, Snyk API & Web cannot verify the vulnerabilities that depend on them.

## Related information

You can also identify Snyk requests through the user-agent header. For more information, visit [Identify scanner requests](identify-scanner-requests.md).

If you use a Web Application Firewall (WAF) in front of your target, it can block scan requests from Snyk IP addresses and cause the scan to fail. To avoid that, visit [Configure IPs in WAFs](configure-ips-in-wafs.md).

If you use a Farcaster Scanning Agent, see the full network requirements, including tunnel and container registry access, in the [Farcaster Scanning Agent README](https://github.com/Probely/farcaster-onprem-agent#network-requirements).
