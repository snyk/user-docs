# V1 API

## About the V1 API

{% hint style="info" %}
The Snyk API is available only for Enterprise plans.

For more information, see [Plans and pricing](https://snyk.io/plans).

The V1 API will be deprecated eventually, as further Snyk developments are now focused on the REST API.
{% endhint %}

The V1 API enables you to test a package for issues as they are defined by Snyk, and to automate Snyk processes to accomplish your specific workflows. Customers and partners can perform functions including:

* Accessing vulnerability data
* Scanning Projects and applications
* Receiving remediation advice
* Viewing user data to build custom security solutions

The V1 API endpoints are available in the [Reference](reference/).&#x20;

## API URLs

Snyk is hosted in the following regions. Each region has its own base URL.

<table><thead><tr><th width="189">Region</th><th>Base URL</th></tr></thead><tbody><tr><td>SNYK-US-01</td><td><code>https://api.snyk.io/v1/</code></td></tr><tr><td>SNYK-US-02</td><td><code>https://api.us.snyk.io/v1/</code></td></tr><tr><td>SNYK-EU-01 </td><td><code>https://api.eu.snyk.io/v1/</code> </td></tr><tr><td>SNYK-AU-01</td><td><code>https://api.au.snyk.io/v1/</code></td></tr></tbody></table>

Before running 
``` 
snyk auth <API_TOKEN>
``` 
run 
``` 
snyk config environment <REGION_NAME>
```

For example, 
```
snyk config environment SNYK-US-02
```

{% hint style="info" %}
This API is available only over HTTPS. Calling the API over HTTP will yield a 404 for all requests.
{% endhint %}

## Authorization

To use this API, you must get your token from Snyk. You can find the token in your [personal account settings](https://snyk.io/account/) after you register with Snyk and log in. For details, see [Authentication for API](authentication-for-api/).&#x20;

Provide the token in an `Authorization` header with the token, preceded by `token`:

```
Authorization: token API_KEY
```

Otherwise, a 401 "Unauthorized" response will be returned.

```
HTTP/1.1 401 Unauthorized

{
    "code": 401,
    "error": "Not authorised",
    "message": "Not authorised"
}
```

## Rate limiting

Snyk limits the requests to the V1 API to help provide a stable experience for customers.

The V1 API has a default rate limit of **2,000 requests per minute**, but some specific endpoints have lower limits. Refer to the reference docs for each endpoint to see the rate limits.

If you exceed the rate limit, you will receive a `429` error response.

## Errors

The V1 API uses standard HTTP error codes for error responses.

```json
{
    "code": 404,
    "message": "Org 39db46b1-ad57-47e6-a87d-e34f6968030b was not found or you may not have the correct permissions to access the org.",
    "error": "Org 39db46b1-ad57-47e6-a87d-e34f6968030b was not found or you may not have the correct permissions to access the org."
}
```
The error 404 could also be due to you trying to authenticate with an API token generated in a region different than the one you are trying to work with. Please ensure the region you are authenticating in and the region in which API token is generated.
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
