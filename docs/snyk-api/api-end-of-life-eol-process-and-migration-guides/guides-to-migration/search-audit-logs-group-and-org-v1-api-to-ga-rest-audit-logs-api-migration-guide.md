# Search Audit Logs (Group and Org) v1 API to GA REST Audit logs API migration guide

## What’s new in the REST Search Audit Logs API?

Based on OpenAPI specifications, the Snyk REST API is designed to provide a consistent, friendly, and easy-to-use API framework that introduces some major improvements. The benefits of the new APIs include:

* Consistent versioning
* Improved performance

The Search Audit Logs REST API deprecates offset-based pagination in favour of more performant and consistent cursor-based pagination. Our v1 APIs get slower the higher your page offset becomes and will eventually time out at very high page counts. We’ve improved this considerably in our REST API by using a cursor-based approach that consistently returns results at greater speeds regardless of page count.

As well as performance improvements we have also improved filtering capabilities, allowing users to provide multiple include or exclude events for finer-tuned filtering. In a addition to this, the API was only capable of searching for logs within a minimum time-frame of 24 hours, we’ve improved this with higher granularity time stamps [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339), which allows users to search within windows of minutes on a given day.

`` `api.access` `` logs are now excluded by default on all search queries unless explicitly provided as part of the \`events\` parameter. These are typically high-volume low-information logs that are superseded by our explicit action logs, e.g. group.create. Explicit action logs will contain richer contextual information for an action and should be favoured in all cases.

Here is what you could use in the v1 endpoints, and their GA REST equivalents:

<table><thead><tr><th width="324">v1 Group and Org Search Audit Logs</th><th>GA REST API</th><th>Notes</th></tr></thead><tbody><tr><td>(query) from (YYYY-MM-DD)</td><td>(query) from (<a href="https://datatracker.ietf.org/doc/html/rfc3339">RFC3339</a>)</td><td>Time format changed</td></tr><tr><td>(query) to (YYYY-MM-DD)</td><td>(query) to (<a href="https://datatracker.ietf.org/doc/html/rfc3339">RFC3339</a>)</td><td>Time format changed.</td></tr><tr><td>(query) page (Number)</td><td>(query) cursor (string)</td><td></td></tr><tr><td>(query) sortOrder</td><td>unchanged</td><td></td></tr><tr><td>(body) filters.userId (string)</td><td>(query) user_id (string)</td><td>Moved from request body to query parameter.</td></tr><tr><td>(body) filters.email (string)</td><td>deprecated</td><td>Based on user feedback, this was removed</td></tr><tr><td>(body) filters.event (string)</td><td>(query) events (comma-separated list)</td><td>Parameter name change and moved into query. Supports multiple values.</td></tr><tr><td>(body) filters.excludeEvent (string)</td><td>(query) exclude_events (comma-separated list)</td><td>Parameter name change and moved into query. Supports multiple values.</td></tr><tr><td>(body) filters.projectId (string)</td><td>(query) project_id (string)</td><td>Moved from request body to query parameter.</td></tr></tbody></table>

## Response

Audit log event payloads have not changed as part of the REST migration, however, the structure of the response is different to conform to our standardised JSON API responses.

The V1 response is returned as an array, e.g.

\`\`\`

`[`

&#x20; `{`

&#x20;   `"groupId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

&#x20;   `"orgId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

&#x20;   `"userId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

&#x20;   `"projectId": null,`

&#x20;   `"event": "group.edit",`

&#x20;   `"content": {`

&#x20;     `"before": {`

&#x20;       `"name": "Group Previous Name"`

&#x20;     `},`

&#x20;     `"after": {`

&#x20;       `"name": "Group Current Name"`

&#x20;     `}`

&#x20;   `},`

&#x20;   `"created": "2017-04-11T21:00:00.000Z"`

&#x20; `}`

`]`

` ``` `

The REST response will contain the same event payload information, however, it will be in the following format:

` ``` `

`{`

&#x20; `"data": {`

&#x20;   `"items": [`

&#x20;     `{`

&#x20; `"groupId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

&#x20; `"orgId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

&#x20; `"userId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

&#x20; `"projectId": null,`

&#x20; `"event": "group.edit",`

&#x20; `"content": {`

&#x20;   `"before": {`

&#x20;     `"name": "Group Previous Name"`

&#x20;   `},`

&#x20;   `"after": {`

&#x20;     `"name": "Group Current Name"`

&#x20;   `}`

&#x20; `},`

&#x20; `"created": "2017-04-11T21:00:00.000Z"`

`}`

\
\


&#x20;   `],`

&#x20; `},`

&#x20; `"jsonapi": {`

&#x20;   `"version": "1.0"`

&#x20; `},`

&#x20; `"links": {`

&#x20;   `"next": "https://example.com/api/resource",`

&#x20; `}`

`}`

` ``` `

