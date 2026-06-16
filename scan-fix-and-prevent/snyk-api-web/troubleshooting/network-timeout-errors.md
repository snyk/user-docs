# Network timeout errors

A very important factor in running successful scans using Snyk API & Web is providing clear access to your target.

Having restrictions without allowing the public IP addresses of Snyk API & Web might lead to failed scans, simply because the scanner gets timed out.

## Reasons for network timeout errors

There are several reasons why this might be happening:

* Since the outbound IP addresses are from a Cloud provider, and traffic produced by humans is not supposed to come from a Cloud provider, they might be blocked by default on your firewall or WAF.
* As a result of the intrusive testing during a scan, your WAF might believe it is an attack and block the IP address.
* You have a plugin (like Jetpack, Sucuri, Fail2Ban, and so on) that can actively filter and block requests.
* The request load (which can spread to multiple threads, making multiple requests per second) caused your firewall to block the IP to prevent spamming, or the requests could be slowing down your website because of all those requests.
* A rule configured in your infrastructure might cause the IP to be blocked.

<figure><img src="/.gitbook/assets/network-timeout-errors-example.png" alt="Example of network timeout error message"></figure>

## Resolution

The best solution to fix this issue is to allow the outbound IP addresses of Snyk API & Web to freely access your target by adding them to your infrastructure allowlists. All updated addresses are available in [Scanner outgoing IP addresses](../configure-targets/scanner-outgoing-ip-addresses.md). Some allowlisting examples are provided in [Configure Snyk API & Web IPs in WAFs](configure-snyk-api-web-ips-in-wafs.md).

If you have any geo-location restriction in place, allow Ireland (possibly the EU) since it is where the servers are located.

If you believe that the scans are critically detrimental to your website's performance, reach out to Snyk support, and the scanner load will be fine-tuned to accommodate your needs.
