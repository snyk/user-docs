# Set up the Snyk webhook

Create the Snyk Webhook using the [Create a webhook API](../../../../reference/webhooks-v1.md#org-orgid-webhooks).

The API requires that you provide the Snyk Organization ID, the Snyk authentication token, the public URL for your Lambda function, and the value of your Lambda secret environment variable.

An example request follows. You can use your favorite tool to send the request.

```
POST https://api.snyk.io/v1/org/{SNYK-ORG-ID}/webhooks HTTP/2
Host: snyk.io
Authorization: token {SNYK-TOKEN}
Content-Type: application/json

{
    "url": "https://{TARGET-WEBHOOK-URL}",
    "secret": "{THE-VALUE-OF-YOUR-LAMBDA-SECRET-ENVIRONMENT-VARIABLE}"
}
```

With this request done, your connection from Snyk to Slack will be completed. Every time there is a new vulnerability, you will get a new notification.
