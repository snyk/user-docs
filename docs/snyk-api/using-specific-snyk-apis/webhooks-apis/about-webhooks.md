# About webhooks

{% hint style="warning" %}
The Webhooks API is in beta. While the API is in beta, Snyk may change the API and the structure of webhook payloads at any time, without notice. During this beta, Webhooks are available only in the Snyk US-01, US-02, EU-01, and AU-01 regions.
{% endhint %}

Webhooks allow you to be notified of Snyk system events, enabling you to build notifications and react to changes in your projects. The current implementation supports events for recurring scans of open source and container image related projects.

When events are triggered, Snyk sends HTTP POST requests to URLs you have configured for those events, with all the information you need.

## Example use cases

### Notifications

Receive instant notifications/alerts in your organization's business communication/collaboration software.

### Incident Response

Respond to critical issues before they impact your business. Embrace modern incident management and Snyk to stay ahead of application security. Read more about this use case in the blog ["Shifting left security incident management with the Snyk & Opsgenie integration"](https://snyk.io/blog/security-incident-management-snyk-opsgenie-integration/).

### Security Information and Event Management (SIEM)

Get real-time security alerts aggregated across various sources into a single platform. Read more about Snyk's partnership with Rapid7 and how together we help organizations mitigate security risks.

### Vulnerability Management and Aggregation

Browse the various [Snyk Partner integrations](../../../integrations/partner-integrations.md) for a comprehensive list of solutions.

## Security SSL

Webhooks can only be configured for URLs using the HTTPS protocol. HTTP is not allowed.

## Request signing

When creating a webhook, you must provide a ​secret​ - this is a string that only you know that we will use to sign our transports to you so that you can ensure they come from Snyk. Your secret should be:

* A random string with high entropy
* Not be used for anything else
* Only known to Snyk and your webhook transport consuming code

All transports sent to your webhooks will have a ​`X-Hub-Signature` ​header, which contains the hash signature for the transport. The signature is a HMAC hexdigest of the request body, generated using sha256 and your secret as the HMAC key.

{% hint style="info" %}
`X-Hub-Signature​` always starts with​ `sha256=`
{% endhint %}
