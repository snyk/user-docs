# API End of Life (EOL) process and migration guides

This page explains the process, key dates, and milestones associated with the end-of-life (EOL) cycle for all API endpoints. In this documentation, you will also find detailed information about [key dates](api-eol-endpoints-and-key-dates.md) and [migration guides](guides-to-migration/) for API endpoints that are in the end-of-life process.

## API End of Life (EOL) process

Snyk  GA REST APIs are an evolution of Snyk V1 APIs because the GA REST APIs have the following:&#x20;

* Consistent versioning
* Pagination and caching
* Improved performance
* Specifications for client generation

Snyk EOL for V1 APIs and replacement with GA REST APIs will implement this improvement over the equivalent V1 APIs. As Snyk delivers more GA REST APIs, experimental and beta versions of the REST API will also reach end-of-life.

Migrating from V1 API to GA REST can be a time-consuming process, and Snyk wants to ensure that you have enough time to factor in and execute migrations so that you can have the best API experience as soon as possible. The process for taking an endpoint (V1, experimental, or beta API) to EOL for seamless migration to GA REST is as follows:&#x20;

1. Batches of endpoints will be part of an EOL cycle that begins twice a year: one batch in January and one batch in July.
2. API endpoints can be included for EOL only if they have:
   * A GA REST equivalent or equivalents (except in the rare case where a V1 API does not have or need a GA REST equivalent)
   * Functionality parity between V1 and GA REST (unless explicitly stated otherwise in the migration guide)
   * A migration guide by our field specialists for ease of migration
3. Snyk will [publicly announce](http://updates.snyk.io/) which endpoints will be part of an EOL cycle one month before the cycle begins.&#x20;
4. On the date the EOL begins, the endpoints are deemed **deprecated**. At that point, the documentation of each endpoint will either be **removed** or have a statement added that the endpoint is deprecated. In addition, no new customers will be able to integrate with the endpoint. The endpoint will remain functional for existing customers until the end-of-life date. You can find all of the endpoints reaching end of life and the associated timelines on the [API EOL endpoints and key dates](api-eol-endpoints-and-key-dates.md) page.
5. On a monthly basis during the EOL cadence, Snyk will temporarily halt functionality for the nominated endpoints for a period of time, increasing in duration over the course of the EOL.
6. When we reach the EOL date, the endpoint will stop working, and you will receive an error.&#x20;

## Types of API EOL

The following types of EOL will take place during each cycle:

1. V1 API: When a GA REST equivalent or equivalents are released, Snyk will aim to include the V1 API in an end-of-life cycle as soon as possible. Users will have a six-month window to migrate off the endpoint, beginning at the public announcement in January and July.
2. Experimental and beta: When Snyk upgrades a REST endpoint to GA that has earlier experimental or beta endpoints or both, Snyk will aim to include them in the EOL cycle as soon as possible. Users will have a 3-month window to migrate off the endpoint (starting from the public announcement in January and July).&#x20;

In exceptional circumstances, Snyk may have to announce an EOL for an endpoint outside of the two announcements each year in January and July. Users will receive a one-month notice of the EOL cycle. The time window to migrate off the endpoint will follow the same windows identified for each type of EOL: a V1 API or an experimental or beta endpoint.
