# Snyk REST API overview

{% hint style="info" %}
The information on this page is also available in the [Snyk REST API reference docs](https://apidocs.snyk.io/).
{% endhint %}

Snyk REST API is based on the [JSON:API standard](https://jsonapi.org/), defined in [OpenAPI 3](https://spec.openapis.org/oas/v3.0.3.html), and represents an evolutionary approach to API development, with each endpoint versioned (see [Versioning](snyk-rest-api-overview.md#versioning)).

## When to use Snyk REST API

Snyk REST API is available for you to try out as endpoints are released. Use this API when you find endpoints that correspond to your needs according to your own workflow. You can also use Snyk API v1 at [https://snyk.docs.apiary.io/](https://snyk.docs.apiary.io/).

For more information about Snyk APIs and using an API versus Snyk CLI or an integration, see the [Snyk API docs](https://docs.snyk.io/snyk-api-info).

## API URL (HTTPS only)

The base URL for all Snyk REST API endpoints is [https://api.snyk.io/rest/](https://api.snyk.io/rest/).

This API is only available over HTTPS. Accessing over HTTP will yield a 404 for all requests.

## Authentication

To use this API, you must get your API token from Snyk. You can find your token in your General Account Settings on [https://snyk.io/account/](https://snyk.io/account/) after you register with Snyk and log in. See [Authentication for API](https://docs.snyk.io/snyk-api-info/authentication-for-api).

When using the API directly, provide the API token in an `Authorization` header, preceded by `token` exactly as follows (no initial capital letter or quotation marks):

```
Authorization: token API_KEY
```

If you are using the API through [Snyk Apps](https://docs.snyk.io/integrations/snyk-apps), provide the `access_token` in an `Authorization` header preceded by `bearer` as follows:

```
Authorization: bearer ACCESS_TOKEN
```

Otherwise a 401 "Unauthorized" response will be returned:

```
HTTP/1.1 401 Unauthorized

{
    "status": "401",
    "code": "Unauthorized"
}
```

## JSON:API Content-Type header

Snyk REST API generally adheres to the [JSON:API standard](https://jsonapi.org/), with some [caveats](https://github.com/snyk/sweater-comb/blob/main/docs/principles/jsonapi.md). JSON:API is an opinionated specification for how a client should request and modify resources and how a server should respond to requests.

When using the REST API, send all requests which contain data with the header:

```
Content-Type: application/vnd.api+json
```

Otherwise a 400 "Bad Request" response will be returned.

```
HTTP/1.1 400 Bad Request

{
    "jsonapi": {
        "version": "1.0"
    },
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

Snyk REST API has per-endpoint version contracts. Each endpoint can have its own release and support lifecycle, independent of any other endpoint in Snyk REST API. In its most explicit form, the endpoint version number includes a date and stability tree, for example:

```
2022-01-01~beta
```

This version number indicates that the requested endpoint should be at stability level `2022-01-01` or older `beta`. The possible stability levels are:

* `ga` - Generally Available, the default. Snyk will support these for at least six months after the next GA release.
* `beta` - Beta. Snyk will support these for at least three months after the next beta or GA release.
* `experimental` - Experimental. An experimental endpoint should be considered unstable and regarded as a tech preview. Experimental versions may introduce breaking changes and may be discontinued at any time.

Note: In the default case of Generally Available, there is no stability level specified in the version number itself (that is, only the date is present), for example:

```
2022-01-01
```

This means the requested endpoint should be `2022-01-01` or older on the Generally Available stability tree.

Granular version controls enable Snyk to introduce progressive enhancements. These may require small or minor backwards-incompatible changes. However, using granular version controls means Snyk can deliver rapid enhancements more quickly, while supporting existing endpoints for a guaranteed period of time.

Once an endpoint is marked as deprecated, it will contain a `Sunset` header indicating the date at which that endpoint contract will no longer be supported. For example:

```
Sunset: Wed, 11 Nov 2021 11:11:11 GMT
```

## Pagination

All endpoints in the Snyk REST API are paginated using _cursor-based_ pagination. This form of pagination helps prevent inconsistent results when collections are modified while they are being iterated. However, cursor-based pagination does not totally prevent inconsistent results, which can occur, for example, if an item is inserted in a prior page based on the sort order requested after a request is made. To mitigate the possibility of inconsistent results, by default Snyk sorts by insertion order so it is not possible to have items inserted on a prior page. However, if you specify the `sort` parameter, consistent pagination is no longer guaranteed. If you see inconsistent results you can submit a new request. If consistent pagination is important to your workflow, use the default insertion sort order.

Whenever you receive an API response, it will contain appropriate links in the body of the response as follows:

```
{
    "jsonapi": {
        "version": "1.0"
    },
    "data": [ ... ],
    "links": {
        "prev": "/orgs/123-abc-def-456/projects?version=2022-01-01~beta&limit=10&ending_before=v1.eyJpZCI6Mz1zODQyMH0%3D"
        "next": "/orgs/123-abc-def-456/projects?version=2022-01-01~beta&limit=10&starting_after=v1.eyJpZCI6Mz1zODQyMH0%3D"
    }
}
```

These links contain pre-defined parameters to make pagination easy, which are a combination of:

* `starting_after`: an opaque (Snyk internal) blob that tells Snyk the _last record_ you've seen, and that you want records _after_ this one
* `ending_before`: an opaque (Snyk internal) blob that tells Snyk the _first record_ you've seen, and that you want records _before_ this one
* `limit`: the number of records per page

## Errors

Errors conform to the JSON:API specification and include path-based information to indicate the root cause as follows:

```
{
    "jsonapi": {
        "version": "1.0"
    },
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

## Rate Limiting

There is a limit of 540 requests per minute, per API key. All requests above the limit will get a response with status code `429` - `Too many requests` until requests stop for the duration of the rate-limiting interval (currently one minute).

From time to time, Snyk may introduce new rate limits to maintain overall system health. This is not considered a breaking change. All clients are expected to handle the `429` responses correctly and such requests can be retried later safely.
