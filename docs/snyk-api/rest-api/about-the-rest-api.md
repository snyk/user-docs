# About the REST API

The Snyk REST API is based on the [JSON:API standard](https://jsonapi.org/), defined in [OpenAPI 3.0.3](https://spec.openapis.org/oas/v3.0.3.html), and represents an evolutionary approach to API development, with each endpoint versioned. For more information, see the [Versioning](about-the-rest-api.md#versioning) section on this page.

## API URLs

Snyk is hosted in the following regions. Each region has its own base URL. Use the URL for your region when calling an API.

<table><thead><tr><th width="189">Region</th><th>Base URL</th></tr></thead><tbody><tr><td>SNYK-US-01</td><td><code>https://api.snyk.io/rest</code></td></tr><tr><td>SNYK-US-02</td><td><code>https://api.us.snyk.io/rest</code></td></tr><tr><td>SNYK-EU-01 </td><td><code>https://api.eu.snyk.io/rest</code></td></tr><tr><td>SNYK-AU-01</td><td><code>https://api.au.snyk.io/rest</code></td></tr></tbody></table>

{% hint style="info" %}
This API is available only over HTTPS. Calling this API over HTTP will yield a 404 for all requests.
{% endhint %}

## JSON:API Content-Type header

The Snyk REST API generally adheres to the [JSON:API standard](https://jsonapi.org/), with some [caveats](https://github.com/snyk/sweater-comb/blob/main/docs/principles/jsonapi.md). JSON:API is an opinionated specification for how a client should request and modify resources and how a server should respond to requests.

When using the REST API, send all requests which contain data with the header:

```
Content-Type: application/vnd.api+json
```

Otherwise a 400 "Bad Request" response will be returned.

```json
HTTP/1.1 400 Bad Request

{
    "errors": [
        {
            "status": "400",
            "detail": "Client request did not conform to OpenAPI specification",
            ...
        }
    ]
}
```

## Versioning

### Using versions dated on or after 2024-10-15

The API Reference examples include `?version=text` as a placeholder, where `text` represents the required date-formatted version string. Snyk recommends using `2024-10-15` for the version number unless you are using an earlier version for a specific reason. You can use the current day's date; this will call the most recent version of the API. The format for the date is `YYYY-MM-DD`.

### Using versions dated prior to 2024-10-15

The following information applies to calling versions earlier than 2024-10-15. The Snyk REST API has per-endpoint version contracts. For information about the differences in GA versions, see the [API Changelog](changelog.md). Each endpoint can have its own release and support lifecycle, independent of any other endpoint in the Snyk REST API. In its most explicit form, the endpoint version number includes a date and stability tree, for example:

```
2023-11-27~beta
```

This version number indicates that the requested endpoint should be at stability level `2023-11-27` or older `beta`. The possible stability levels are:

* `ga` - Generally Available, the default. Snyk will support these for at least six months after the next GA release.
* `beta` - Beta. Snyk will support these for at least three months after the next beta or GA release.
* `experimental` - Experimental. An experimental endpoint should be considered unstable and regarded as a tech preview. Experimental versions may introduce breaking changes and may be discontinued at any time.

In the default case of Generally Available, there is no stability level specified in the version number itself, that is, only the date is present, for example:

```
2023-11-27
```

This means the requested endpoint should be 2023-11-27 or older on the Generally Available stability tree.

When the requested endpoint is at a specific stability level, Snyk serves the latest version, the version released on or before the requested date, or that stability or higher. For example, if the requested endpoint has a beta version at 2023-09-29 and GA version at 2024-01-23, and the requested endpoint is after 2024-01-23\~beta, Snyk resolves to the GA version.

Granular version controls enable Snyk to introduce progressive enhancements. These may require small or minor backward-incompatible changes. However, using granular version controls means Snyk can deliver rapid enhancements more quickly while supporting existing endpoints for a guaranteed period of time.

After an endpoint is marked as deprecated, it will contain a `Sunset` header indicating the date at which that endpoint contract will no longer be supported. For example:

```
Sunset: "2021-11-11"
```

## Pagination

All endpoints in the Snyk REST API are paginated using _cursor-based_ pagination. This form of pagination helps prevent inconsistent results when collections are modified while they are being iterated. However, cursor-based pagination does not totally prevent inconsistent results, which can occur, for example, if an item is inserted in a prior page based on the sort order requested after a request is made.

To mitigate the possibility of inconsistent results, by default, Snyk sorts by insertion order, so it is not possible to have items inserted on a prior page. However, if you specify the `sort` parameter, consistent pagination is no longer guaranteed. If you see inconsistent results, you can submit a new request. If consistent pagination is important to your workflow, use the default insertion sort order.

Whenever you receive an API response, it will contain appropriate links in the body of the response as follows:

```json
{
    "data": [ ... ],
    "links": {
        "prev": "/orgs/123-abc-def-456/projects?version=2024-06-10&ending_before=v1.eyJpZCI6Mz1zODQyMH0%3D"
        "next": "/orgs/123-abc-def-456/projects?version=2024-06-10&starting_after=v1.eyJpZCI6Mz1zODQyMH0%3D"
    }
}
```

These links contain pre-defined parameters to make pagination easy. These parameters are a combination of:

* `starting_after`: an opaque (Snyk internal) blob that tells Snyk the _last record_ you have seen and that you want records _after_ this last one
* `ending_before`: an opaque (Snyk internal) blob that tells Snyk the _first record_ you have seen and that you want records _before_ this first one
* `limit`: the number of records per page

## Errors

Errors conform to the JSON:API specification and include path-based information to indicate the root cause as follows:

```json
{
    "errors": [
        {
            "id": "0418e907-a89e-4736-9a5b-91a2022e0899",
            "detail": "Unknown query parameter 'unknown'",
            "status": "400",
            "source": {
                "parameter": "unknown"
            }
        }
    ]
}
```

## Rate limiting

There is a limit of **1620 requests per minute**, per API key. All requests above the limit will get a response with the status code `429` - `Too many requests` until requests stop for the duration of the rate-limiting interval (one minute).

From time to time, Snyk may introduce new rate limits to maintain overall system health. This is not considered a breaking change. All clients are expected to handle the `429` responses correctly and such requests can be retried later safely.
