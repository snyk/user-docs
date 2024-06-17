# V1 API

## About the V1 API

{% hint style="info" %}
The Snyk API is available only for Enterprise plans.&#x20;

For more information, see [Plans and pricing](https://snyk.io/plans).

The V1 API will be sunset eventually, as further Snyk developments are now focused on the REST API.
{% endhint %}

The V1 API has the ability to test a package for issues as they are defined by Snyk, and to provide Snyk security automation according to your own workflows, unconstrained by security processes in Snyk products. Customers and partners can perform functions including:

* Accessing vulnerability data
* Scanning Projects and applications
* Receiving remediation advice
* Viewing user data to build custom security solutions

The V1 API endpoints are available in the [Reference](reference/) in the Snyk user docs. Updates are made in the user docs. The endpoints migrated to the user docs remain [online](https://snyk.docs.apiary.io) also.

## API URL&#x20;

Snyk is hosted in three main regions: US, EU, and AU. Each region has its own base URL.

<table><thead><tr><th width="189">Region</th><th>Base URL</th></tr></thead><tbody><tr><td>US</td><td><code>https://api.snyk.io/v1</code></td></tr><tr><td>Europe</td><td><code>https://api.eu.snyk.io/v1</code></td></tr><tr><td>Australia</td><td><code>https://api.au.snyk.io/v1</code></td></tr></tbody></table>

## Rate limiting

Snyk limits the requests to the V1 API to help provide a stable experience for customers.

The V1 API has a default rate limit of **2,000 requests per minute**, but some specific endpoints have lower limits. Refer to the reference docs for each endpoint to see the rate limits.

If you exceed the rate limit, you will receive a `429` error response.

## Errors

The V1 API uses standard HTTP error codes for error responses.&#x20;

```json
{
    "code": 404,
    "message": "Org 39db46b1-ad57-47e6-a87d-e34f6968030b was not found or you may not have the correct permissions to access the org.",
    "error": "Org 39db46b1-ad57-47e6-a87d-e34f6968030b was not found or you may not have the correct permissions to access the org."
}
```

The error reference is also supplied in the `x-error-reference` header in the server reply.

Example `500` response:

```sh
HTTP/1.1 500 Internal Server Error
x-error-reference: a45ec9c1-065b-4f7b-baf8-dbd1552ffc9f
Content-Type: application/json; charset=utf-8
Content-Length: 1848
Vary: Accept-Encoding
Date: Sun, 10 Sep 2017 06:48:40 GMT
```
