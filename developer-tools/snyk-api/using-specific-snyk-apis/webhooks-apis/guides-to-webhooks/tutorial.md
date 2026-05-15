# How to use Snyk webhooks with Zapier

{% hint style="info" %}
Snyk API v1 docs are in the [Reference](../../../reference/).
{% endhint %}

## ​Integration example

Create a new Zap in [Zapier](https://zapier.com).

### Trigger

To access request headers, create a **Catch Raw Hook** trigger. This trigger provides the request payload as a string, so you must parse it to JSON.

You receive a webhook URL where you send requests.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-1%20\(1\).png)

Create a Webhook in Snyk using the API with the `your-url` URL.

```
POST /api/v1/org/{orgId}/webhooks HTTP/2
> Host: snyk.io
> Authorization: token {authToken}
> Content-Type: application/json
| {
|   "url": "https://hooks.zapier.com/hooks/catch/9002958/oemlgkz/",
|   "secret": "my-secret-string"
| }
```

The API responds with the new webhook.

```
< HTTP/2 200 
< Content-Type: application/json
| {
|   "id": "{webhookId}",
|   "url": "https://hooks.zapier.com/hooks/catch/9002958/oemlgkz/",
| }
```

You can ping a webhook to test the Zapier trigger.

```
> POST /api/v1/org/{orgId}/webhooks/{webhookId}/ping HTTP/2
> Host: snyk.io
> Authorization: token {authToken}
> Content-Type: application/json
```

Select a ping request from the list and map fields.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-2%20\(1\).png)

### Action (validate a payload)

Create a JS Action to validate a payload:

**"Code by Zapier" → "Run Javascript"**

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-3%20\(1\).png)

Map `headers['X-Hub-Signature']` and `payload string` to the snippet variables.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-4%20\(1\).png)

This snippet adds an `isValid: boolean` variable to Zap fields.

{% hint style="info" %}
Replace `my-secret-string` string with a webhook's secret string.
{% endhint %}

```javascript
const crypto = require('crypto');
const secret = "my-secret-string";

function makeSignature(body, secret) {
  const hmac = crypto.createHmac('sha256', secret);
  hmac.update(body, 'utf8');

  return `sha256=${hmac.digest('hex')}`;
}

try {
  const body = JSON.parse(inputData.body);
  const { project, org, group, newIssues } = body;

  output = { 
    isValid: inputData.signature === makeSignature(inputData.body, secret)
  };
} catch (err) {
  output = { isValid: false, err: err.message };
}
```

Test the snippet, ensure `isValid === true`.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-5%20\(1\).png)

### Action (parse a payload)

Create another action to parse the payload string into a format Zapier uses.

Create the same JS Action:

**"Code by Zapier" → "Run Javascript"**, with the following field mapping:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-6%20\(1\).png)

And the following JS snippet:

```
try {
  output = JSON.parse(inputData.body);
} catch (err) {
  output = { err: err.message };
}
```

Parse a request payload and map it to the Zap variables.

### Action (format issues)

New issues are lists of objects. Zapier requires a list of strings. Format `newIssues` as `string[]`.

Create one more JS Action:

**"Code by Zapier" → "Run Javascript"**, and paste the following snippet:

```
function formatIssue({ pkgName, pkgVersions, issueData }) {
  return `
  <a href="${issueData.url}">${issueData.title}</a><br/>
  Vulnerability in ${pkgName} (${pkgVersions.join(', ')}). ${issueData.severity} severity.
`;
}

try {
  const { newIssues, ...body } = JSON.parse(inputData.body);

  output = { ...body, newIssues: newIssues.map(formatIssue) };
} catch (err) {
  output = { newIssues: [], err: err.message };
}
```

### Action (filter)

After providing all fields, decide whether to use the event.

To filter, create **"Filter by Zapier"** app:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-7%20\(1\).png)

Select a filter method.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-8%20\(1\).png)

### Action (send a notification)

Access all fields to build a notification template. Send an email or choose other notification types.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-9%20\(1\).png)
