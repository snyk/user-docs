# Configure IPs in WAFs

Snyk API & Web uses specific public IP addresses to scan your targets. If you are using a Web Application Firewall (WAF) in front of your target, it can block scan requests and cause the scan to fail. To avoid that, you must configure the WAF to allow Snyk API & Web IP addresses.

For a list of Snyk API & Web IP addresses, see [Scanner IP address](scanner-ip-address.md).

## Configure Cloudflare WAF

Cloudflare provides documentation explaining how to configure access rules for its WAF. There is an overview of [IP Access rules](https://developers.cloudflare.com/waf/tools/ip-access-rules/), and the configuration steps are described in [Create an IP Access rule](https://developers.cloudflare.com/waf/tools/ip-access-rules/create/).

When following these steps, use this information:

* **IP, IP range, country name, or ASN** - Enter the Snyk API & Web IP address for your case.
* **Action** - Select **Allow**.
* **Zone** - From the available options, select **This website** if you want to apply the rule only to the current zone. Alternatively, select **All websites in account** if you want the rule to be created in all zones of your Cloudflare account.
* **Notes** - This is optional, but you can provide text identifying the rule. For example, "Snyk API & Web IP".

After creating the rule, your target scans with Snyk API & Web should run smoothly without being blocked by Cloudflare WAF.
