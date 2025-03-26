# Create a Snyk Webhook

Create the Snyk Webhook using the [Create a webhook API](../../../../reference/webhooks-v1.md#org-orgid-webhooks).

The API requires that you provide the Snyk Organization ID, the Snyk authentication token, and the target webhook URL.

An example request follows. You can use your favorite tool to send the request.

```
POST https://api.snyk.io/v1/org/{SNYK-ORG-ID}/webhooks HTTP/2
Host: snyk.io
Authorization: token {SNYK-TOKEN}
Content-Type: application/json

{
    "url": "https://{URL}",
    "secret": "my-secret-string"
}
```

The response is like this:

```
{
  "id": "{SNYK-WEBHOOK-ID}",
  "url": "https://{URL}",
}
```

You can then use the [Ping a webhook API](../../../../reference/webhooks-v1.md#org-orgid-webhooks-webhookid-ping) to pro-actively trigger the Snyk Webhook in order to test your integration:

```
POST https://api.snyk.io/v1/org/{SNYK-ORG-ID}/webhooks/{SNYK-WEBHOOK-ID}/ping HTTP/2
Host: snyk.io
Authorization: token {SNYK-TOKEN}
Content-Type: application/json
```

When the Azure Function and Snyk Webhook are created you can use the [New Relic Curated UI and Snyk Custom Dashboard](new-relic-curated-ui-and-snyk-custom-dashboard.md).
