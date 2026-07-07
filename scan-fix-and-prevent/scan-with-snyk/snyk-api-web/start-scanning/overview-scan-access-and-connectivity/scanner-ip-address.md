# Scanner IP address

Snyk API & Web uses the following IP addresses to make requests.

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

If you have your own single tenant or dedicated infrastructure, contact the support team.

## Deprecated IP addresses

The following IP addresses are deprecated. No traffic originates from them:

* 35.190.194.212 (GCP IP)
* 35.187.52.245 (GCP IP)

## Related information

You can also identify Snyk requests through the user-agent header. For more information, visit [Identify scanner requests](identify-scanner-requests.md).

If you use a Web Application Firewall (WAF) in front of your target, it can block scan requests from Snyk IP addresses and cause the scan to fail. To avoid that, visit [Configure IPs in WAFs](configure-ips-in-wafs.md).
