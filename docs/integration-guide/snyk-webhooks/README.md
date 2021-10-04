# Snyk Webhooks

Webhooks allow you to be notified of Snyk system events, enabling you to build notifications and react to changes in your projects.

When events are triggered, Snyk sends HTTP POST requests to URLs you have configured for those events, with all the information you need.

### Example use cases

#### Notifications

Receive instant notifications/alerts in your organization's business communication/collaboration software. Please refer to our free tutorial on setting this up with [Microsoft Teams](../../partner-workshops/microsoft-azure/notifications-in-microsoft-teams/) for step-by-step instructions.

#### Incident Response

Respond to critical issues before they impact your business. Embrace modern incident management and Snyk to stay ahead of application security. Read more about this use case in the blog ["Shifting left security incident management with the Snyk & Opsgenie integration"](https://snyk.io/blog/security-incident-management-snyk-opsgenie-integration/) as well as our free  [Opsgenie](../../partner-workshops/atlassian/opsgenie/) that guides you on configuring this integration. 

#### Security Information and Event Management \(SIEM\)

Get real-time security alerts aggregated across various sources into a single platform. read more about Snyk's partnership with Rapid7 and how together we help organizations mitigate security risks.

#### Vulnerability Management and Aggregation

Browse the various [Snyk Partner integrations](https://support.snyk.io/hc/en-us/sections/360003642858-Vulnerability-Management-Tools) for a comprehensive list of solutions.

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

When creating a webhook, you must provide a ​secret​ - this is a string that only you know that we will use to sign our transports to you so that you can ensure they come from Snyk. Your secret should:

* Be a random string with high entropy
* Not be used for anything else
* Only known to Snyk and your webhook transport consuming code

All transports sent to your webhooks will have a ​`X-Hub-Signature` ​header, which contains the hash signature for the transport. The signature is a HMAC hexdigest of the request body, generated using sha256 and your secret as the HMAC key.

{% hint style="info" %}
`X-Hub-Signature​`always starts with​ `sha256=`
{% endhint %}

### Validating payloads

You could use a function in Node.JS such as the following to validate these signatures on incoming requests from Snyk:

```javascript
import​ * ​as​ crypto ​from​ ​'crypto'​;

function​ ​verifySignature​(request, secret) {
    const​ hmac = crypto.createHmac(​'sha256'​, secret); 
    ​const​ buffer = ​JSON​.stringify(request.body); 
    hmac.update(buffer, ​'utf8'​);
    
    const​ signature = ​`sha256=${hmac.digest('hex')}`​;
    
    return signature === request.headers['x-hub-signature'​];
} 
```

### Configure webhooks

Webhooks can be managed via our API at the organization level by organization admins.

## Event payloads

### Payload versioning

Payloads may evolve over time, and so are versioned. Payload versions are supplied as a suffix to the `X-Snyk-Event` ​header. For example, ​`project_snapshot/v0​` indicates that the payload is v0 of the `project-snapshot` ​event.

Version numbers only increment when a breaking change is made - for example, removing a field that used to exist, or changing the name of a field. Version numbers do not increment when making an additive change, such as adding a new field that never existed before.

{% hint style="danger" %}
During the BETA phase, the structure of webhook payloads may change at any time, and we will endeavor to notify BETA program members ahead of time of such changes.
{% endhint %}

### `ping` event example

The ping event happens after a new webhook is created, and can also be manually triggered using the ping webhook API. This is useful to test that your webhook receives data from Snyk correctly.

```javascript
POST /your-url HTTP/1.1
Host: my.app.com
X-Snyk-Event: ping/v0
X-Snyk-Transport-ID: 998fe884-18a0-45db-8ae0-e379eea3bc0a
X-Snyk-Timestamp: 2020-09-25T15:27:53Z
X-Hub-Signature: sha256=ed19e2591ec0a157ee109cd394bb68c62a7a82125b4c8f6617f8b42c867d71c0 User-Agent: Snyk-Webhooks/044aadd
Content-Type: application/json
{
  "webhookId"​: ​"d3cf26b3-2d77-497b-bce2-23b33cc15362"
}
```

### `project-snapshot` event example

This event is triggered every time an existing project is tested and a new snapshot is created. It is triggered on every test of a project, whether or not there are new issues. This event is ​not​ triggered when a new project is created/imported, however.

{% hint style="danger" %}
The details of this payload may change before or during the BETA program.
{% endhint %}

```javascript
POST /your-url HTTP/1.1
Host: my.app.com
X-Snyk-Event: project_snapshot/v0
X-Snyk-Transport-ID: 998fe884-18a0-45db-8ae0-e379eea3bc0a
X-Snyk-Timestamp: 2020-09-25T15:27:53Z
X-Hub-Signature:
sha256=ed19e2591ec0a157ee109cd394bb68c62a7a82125b4c8f6617f8b42c867d71c0
User-Agent: Snyk-Webhooks/044aadd
Content-Type: application/json
```



