---
description: How to allow Snyk API and Web scanner IPs in your WAF
nav_context: agnostic
---

# Configure IPs in WAFs

Snyk API & Web uses specific public IP addresses to scan your targets. If you use a Web Application Firewall (WAF) in front of your target, it can block scan requests and cause the scan to fail. To avoid that, configure the WAF to allow Snyk IP addresses.

For a list of Snyk IP addresses, visit [Scanner IP address](scanner-ip-address.md).

## Configure Cloudflare WAF

Cloudflare provides documentation explaining how to configure access rules for its WAF. There is an overview of [IP Access rules](https://developers.cloudflare.com/waf/tools/ip-access-rules/), and the configuration steps are described in [Create an IP Access rule](https://developers.cloudflare.com/waf/tools/ip-access-rules/create/).

When following these steps, use this information:

* **IP, IP range, country name, or ASN** - Enter the Snyk IP address for your case.
* **Action** - Select **Allow**.
* **Zone** - To apply the rule only to the current zone, select **This website**. To create the rule in all zones of your Cloudflare account, select **All websites in account**.
* **Notes** - Optionally, provide text identifying the rule. For example, "Snyk API & Web IP".

After you create the rule, your target scans with Snyk run without being blocked by Cloudflare WAF.
