---
description: How to resolve network timeout errors in Snyk API and Web
nav_context: agnostic
---

# Network timeout errors

Providing clear access to your target is an important factor in running successful scans using Snyk API & Web.

Restrictions that do not allow the public IP addresses of Snyk can lead to failed scans because the scanner times out.

## Reasons for network timeout errors

This can happen for several reasons:

* The outbound IP addresses come from a cloud provider, and traffic produced by humans is not supposed to come from a cloud provider, so your firewall or WAF can block them by default.
* During the intrusive testing in a scan, your WAF can interpret the traffic as an attack and block the IP address.
* You have a plugin (such as Jetpack, Sucuri, or Fail2Ban) that actively filters and blocks requests.
* The request load can spread across multiple threads, making multiple requests per second. This load can cause your firewall to block the IP to prevent spamming, or it can slow down your website.
* A rule configured in your infrastructure can cause the IP to be blocked.

## Resolution

To fix this issue, allow the outbound IP addresses of Snyk to access your target by adding them to your infrastructure allowlists. All updated addresses are available in [Scanner outgoing IP addresses](../start-scanning/overview-scan-access-and-connectivity/scanner-ip-address.md). For allowlisting examples, visit [Configure Snyk API & Web IPs in WAFs](../start-scanning/overview-scan-access-and-connectivity/configure-ips-in-wafs.md).

If you have any geo-location restriction in place, allow Ireland and the EU, because the servers are located there.

If the scans are critically detrimental to the performance of your website, contact Snyk support to fine-tune the scanner load to your needs.
