# Snyk Webhooks

### Snyk Webhooks for Snyk API

[Visit the Snyk API documentation for Webhooks](https://snyk.docs.apiary.io/#introduction/consuming-webhooks) to get information on consuming, validating, and examples.

{% embed url="https://snyk.docs.apiary.io/#introduction/consuming-webhooks" %}

Webhooks allow you to be notified of Snyk system events, enabling you to build notifications and react to changes in your projects. The current implementation supports events for recurring scans of open source and container image related projects.

When events are triggered, Snyk sends HTTP POST requests to URLs you have configured for those events, with all the information you need.&#x20;

### Example use cases

#### Notifications

Receive instant notifications/alerts in your organization's business communication/collaboration software. Please refer to our free tutorial on setting this up with [Microsoft Teams](../../more-info/getting-started/snyk-integrations/microsoft-azure/notifications-in-microsoft-teams/) for step-by-step instructions.

#### Incident Response

Respond to critical issues before they impact your business. Embrace modern incident management and Snyk to stay ahead of application security. Read more about this use case in the blog ["Shifting left security incident management with the Snyk & Opsgenie integration"](https://snyk.io/blog/security-incident-management-snyk-opsgenie-integration/) as well as our free [Opsgenie](../../more-info/getting-started/atlassian-integrations/atlassian/opsgenie/) guide that guides you on configuring this integration.

#### Security Information and Event Management (SIEM)

Get real-time security alerts aggregated across various sources into a single platform. read more about Snyk's partnership with Rapid7 and how together we help organizations mitigate security risks.

#### Vulnerability Management and Aggregation

Browse the various [Snyk Partner integrations](../../integrations/vulnerability-management-tools/) for a comprehensive list of solutions.

### Webhook headers

Event messages are delivered with a ​`Content-Type​` of ​`application/json`,​ with the event payload as JSON in the request body. We also send the following headers:

* `X-Snyk-Event` the name of the event and the version of the payload, such as `ping/v1`
* `X-Snyk-Transport-ID` a GUID to identify this delivery
* ​ `X-Snyk-Timestamp`an ISO 8601 timestamp for when the event occurred, e.g. `2020-09-25T15:27:53Z`
* `X-Hub-Signature` the HMAC hex digest of the request body which is used to secure your webhooks and ensure the request did indeed come from Snyk
* `User-Agent​` identifies the origin of the request, e.g. ​`Snyk-Webhooks/XXX`

Each webhook receives all events.

### Security SSL

Webhooks can only be configured for URLs using the HTTPS protocol. HTTP is not allowed.

### Request signing

When creating a webhook, you must provide a ​secret​ - this is a string that only you know that we will use to sign our transports to you so that you can ensure they come from Snyk. Your secret should be:

* A random string with high entropy
* Not be used for anything else
* Only known to Snyk and your webhook transport consuming code

All transports sent to your webhooks will have a ​`X-Hub-Signature` ​header, which contains the hash signature for the transport. The signature is a HMAC hexdigest of the request body, generated using sha256 and your secret as the HMAC key.

{% hint style="info" %}
`X-Hub-Signature​`always starts with​ `sha256=`
{% endhint %}
