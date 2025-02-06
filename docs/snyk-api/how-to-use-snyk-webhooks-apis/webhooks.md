# Webhook Events and Payloads

Snyk webhooks enable you to receive real-time event notifications about security vulnerabilities and project updates. This guide explains the webhook event structure, payloads, and security measures you need to handle them effectively.

## Webhook Request Structure

Webhook requests are sent with `Content-Type: application/json`, containing a JSON payload in the request body. The following headers are included:

| Header | Description |
|--------|-------------|
| `X-Snyk-Event` | The name of the event. |
| `X-Snyk-Transport-ID` | A GUID to identify this delivery. |
| `X-Snyk-Timestamp` | An ISO 8601 timestamp of when the event occurred (e.g., `2020-09-25T15:27:53Z`). |
| `X-Hub-Signature` | The HMAC hex digest of the request body, ensuring the request originated from Snyk. |
| `User-Agent` | Identifies the request origin, e.g., `Snyk-Webhooks/XXX`. |

Once your server is configured to receive webhooks, it must listen for and validate incoming requests. For security, you should restrict requests to those originating from Snyk and verify the `X-Hub-Signature`.

### Validating Webhook Requests

To ensure authenticity, verify the `X-Hub-Signature` header using your webhook secret. Here’s an example in Node.js:

```javascript
const crypto = require('crypto');

function verifySignature(secret, body, signature) {
  const expectedSignature = `sha256=${crypto.createHmac('sha256', secret).update(body).digest('hex')}`;
  return crypto.timingSafeEqual(Buffer.from(expectedSignature), Buffer.from(signature));
}
```

## Webhook Event Types

### ping

The `ping` event occurs when a new webhook is created or manually triggered via the API. This is useful to confirm your webhook is receiving data correctly.

#### Example Request:

```http
POST /webhook-handler/snyk123 HTTP/1.1
Host: my.app.com
X-Snyk-Event: ping/v0
X-Snyk-Transport-ID: 998fe884-18a0-45db-8ae0-e379eea3bc0a
X-Snyk-Timestamp: 2020-09-25T15:27:53Z
X-Hub-Signature: sha256=7d38cdd689735b008b3c702edd92eea23791c5f6
User-Agent: Snyk-Webhooks/044aadd
Content-Type: application/json

{
  "webhookId": "d3cf26b3-2d77-497b-bce2-23b33cc15362"
}
```

### project_snapshot

Triggered when an existing project is tested, generating a new snapshot. This event fires on every test, even if no new issues are found. It does **not** trigger for newly created or imported projects.

#### Example Request:

```http
POST /webhook-handler/snyk123 HTTP/1.1
Host: my.app.com
X-Snyk-Event: project_snapshot/v0
X-Snyk-Transport-ID: 998fe884-18a0-45db-8ae0-e379eea3bc0a
X-Snyk-Timestamp: 2020-09-25T15:27:53Z
X-Hub-Signature: sha256=7d38cdd689735b008b3c702edd92eea23791c5f6
User-Agent: Snyk-Webhooks/044aadd
Content-Type: application/json

{
  "project": { ... },
  "org": { ... },
  "group": { ... },
  "newIssues": [],
  "removedIssues": []
}
```

## Payload Structure

Each event contains structured data. Below are key objects found in payloads:

### project

Refer to the [Projects (v1) API](https://snyk.io/docs/api/projects/) for full details.

```json
"project": {
  "name": "snyk/goof",
  "id": "af137b96-6966-46c1-826b-2e79ac49bbd9",
  "created": "2018-10-29T09:50:54.014Z",
  "origin": "github",
  "type": "maven",
  "totalDependencies": 42,
  "issueCountsBySeverity": { "low": 13, "medium": 8, "high": 4, "critical": 5 },
  "browseUrl": "https://app.snyk.io/org/example-org/project/af137b96-6966-46c1-826b-2e79ac49bbd9"
}
```

### org

Refer to the [Organizations (v1) API](https://snyk.io/docs/api/orgs/).

```json
"org": {
  "name": "My Org",
  "id": "a04d9cbd-ae6e-44af-b573-0556b0ad4bd2",
  "url": "https://api.snyk.io/v1/org/my-org",
  "created": "2020-11-18T10:39:00.983Z"
}
```

### issue

```json
{
  "id": "npm:ms:20170412",
  "pkgName": "ms",
  "issueData": {
    "title": "Regular Expression Denial of Service (ReDoS)",
    "severity": "low",
    "url": "https://snyk.io/vuln/npm:ms:20170412",
    "cvssScore": 3.7,
    "nearestFixedInVersion": "2.0.0"
  },
  "fixInfo": { "isUpgradable": false, "isPatchable": true }
}
```

## Testing Webhooks

To test your webhook integration, use webhook debugging tools like [Beeceptor](https://beeceptor.com/) or [Webhook.site](https://webhook.site/). These tools allow you to inspect incoming requests and validate event payloads before deploying your webhook handler to production.

### Example Beeceptor Test

1. Create a test endpoint on [Beeceptor](https://beeceptor.com/).
2. Configure Snyk to send webhooks to your Beeceptor URL.
3. Trigger a `ping` event.
4. Inspect the received request and validate the payload.
