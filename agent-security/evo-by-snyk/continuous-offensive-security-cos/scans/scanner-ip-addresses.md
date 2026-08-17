# Scanner IP addresses

Scan traffic originates from a fixed set of IP addresses. Allowing these addresses through your network controls is often what makes the difference between a thorough scan and one that spends its time being blocked.

## Addresses by region

{% hint style="info" %}
Allow every address listed for your region, not just one. Traffic from a single scan can leave through any of the gateways in that region.
{% endhint %}

Scans run from the region your organization is hosted in. Allow the addresses for your region.

{% tabs %}
{% tab title="AWS US" %}
| Gateway | IP address       |
| ------- | ---------------- |
| NAT A   | `18.225.126.146` |
| NAT B   | `3.19.19.31`     |
| NAT C   | `18.189.160.130` |
{% endtab %}

{% tab title="AWS EU" %}
| Gateway | IP address       |
| ------- | ---------------- |
| NAT A   | `63.182.106.183` |
| NAT B   | `18.193.164.139` |
| NAT C   | `3.125.110.188`  |
{% endtab %}

{% tab title="AWS AU" %}
| Gateway | IP address       |
| ------- | ---------------- |
| NAT A   | `16.176.58.90`   |
| NAT B   | `32.236.245.165` |
| NAT C   | `15.135.28.47`   |
{% endtab %}

{% tab title="GCP" %}
| Gateway | IP address     |
| ------- | -------------- |
| Gateway | `34.23.72.204` |
{% endtab %}
{% endtabs %}

## Where to allow them

| Control                                | Why it matters                                                                                                                                                                              |
| -------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Web application firewall**           | A WAF that sees a burst of attack payloads will usually start blocking or rate limiting. Allowing the scanner addresses keeps traffic flowing so the application itself is what gets tested |
| **Rate limiting and bot protection**   | Scan traffic looks automated because it is. Rate limiters and bot detection both tend to throttle it                                                                                        |
| **Network firewall or security group** | Required where the application only accepts traffic from known sources                                                                                                                      |
| **CAPTCHA and challenge pages**        | A challenge in the request path stops agents from proceeding past it                                                                                                                        |

Allowing scanner addresses at the WAF is an alternative to the custom header approach described in Add a target. Either works. A header is easier to arrange when you cannot change network policy; an IP allowlist is easier when you cannot inject headers.

{% hint style="warning" %}
Allowing scan traffic past your WAF is deliberate. You are testing the application, not the firewall in front of it. A vulnerability that a WAF happens to mask is still a vulnerability, and WAF rules change.
{% endhint %}

## Symptoms of blocked scan traffic

Network controls do not announce themselves. They show up as results that look like a secure application:

* Recon discovers far fewer pages and endpoints than your application contains.
* A scan runs substantially longer than previous scans of the same target.
* Findings drop sharply from one scan to the next with no corresponding fixes.

If you see these, check whether your network controls are throttling the scanner addresses before concluding anything about the application. See Monitor and manage scans.

## Identifying scan traffic in your logs

The scanner addresses let you attribute requests in your own logs, which is useful for separating scan activity from real users when reviewing an incident, and for confirming that traffic arrived at all when a scan produced less than expected.

You can also configure a custom header to make scan requests easy to filter on. See Add a target.
