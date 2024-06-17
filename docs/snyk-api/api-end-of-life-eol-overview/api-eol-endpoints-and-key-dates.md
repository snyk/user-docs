# API EOL endpoints and key dates

## APIs at EOL

Beginning July 19, 2024, the following endpoints will follow the EOL process:

| Endpoint                                                                                                                                                                                                                                                  | EOL date          | Migration guide |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------- | --------------- |
| [v1: group and org level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs)                                                                                                                                                                   | January 19, 2025  |                 |
| Experimental: Get all issues by [Org](https://apidocs.snyk.io/experimental?version=2023-03-10\~experimental#get-/orgs/-org\_id-/issues) and [Group](https://apidocs.snyk.io/experimental?version=2023-03-10\~experimental#get-/groups/-group\_id-/issues) | October 21,  2024 |                 |
| v1: [List all members in a Group](https://snyk.docs.apiary.io/#reference/groups/group-settings/list-all-members-in-a-group)                                                                                                                               | January 19, 2025  |                 |

## Brownouts

A brownout occurs when Snyk temporarily suspends an endpoint from being usable. returning a `410 gone` response when a user calls the endpoint.

Snyk brownouts for APIs that are part of an end-of-life cycle will occur at 12:00 UTC. For the end-of-life cycle beginning July 19, 2024, the brownouts will occur on the following dates. Users will see a reminder two weeks before the brownout through an announcement on [updates.snyk.io](http://updates.snyk.io/):

| Endpoints             | Brownout date     | Duration   |
| --------------------- | ----------------- | ---------- |
| Non-GA REST endpoints | Tuesday August 13 | 30 minutes |
| Non-GA REST endpoints | September 10      | 1 hour     |
| v1 endpoints          | October 8         | 1 hour     |
| v1 endpoints          | November 12       | 2 hours    |
| v1 endpoints          | December 10       | 4 hours    |
