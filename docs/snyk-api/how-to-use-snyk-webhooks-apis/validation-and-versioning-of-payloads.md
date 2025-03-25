# Validation and versioning of payloads

## Validating payloads

All transports sent to your webhooks have a `X-Hub-Signature` header, which contains the hash signature for the transport. The signature is a HMAC hexdigest of the request body, generated using sha256 and your `secret` as the HMAC key.

You could use a function in Node.JS such as the following to validate these signatures on incoming requests from Snyk:

```
import * as crypto from 'crypto';

function verifySignature(request, secret) {
  const hmac = crypto.createHmac('sha256', secret);
  const buffer = JSON.stringify(request.body);
  hmac.update(buffer, 'utf8');

  const signature = `sha256=${hmac.digest('hex')}`;

  return signature === request.headers['x-hub-signature'];
}
```

## Payload versioning

Payloads may evolve over time, and so are versioned. Payload versions are supplied as a suffix to the `X-Snyk-Event` header. For example, `project_snapshot/v0` indicates that the payload is `v0` of the `project_snapshot` event.

Version numbers only increment when a breaking change is made, for example, removing a field that used to exist, or changing the name of a field. Version numbers do not increment when making an additive change, such as adding a new field that never existed before.

{% hint style="warning" %}
During the BETA phase, the structure of webhook payloads may change at any time, so Snyk  recommends that you check the payload version.
{% endhint %}

