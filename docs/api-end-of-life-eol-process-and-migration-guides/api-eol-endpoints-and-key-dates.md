# API EOL endpoints and key dates

## APIs at EOL

Beginning July 22, 2024, the following endpoints will follow the EOL process:

| Endpoint                                                                                                                                                                                                                                    | Endpoint type | EOL date          | Migration guide                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------- | ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| [Group and org level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs/group-level-audit-logs/get-group-level-audit-logs)                                                                                                       | v1            | January 22, 2025  |                                                                                                                                                      |
| Get all issues by [Org](https://apidocs.snyk.io/experimental?version=2023-03-10\~experimental#get-/orgs/-org\_id-/issues) and [Group](https://apidocs.snyk.io/experimental?version=2023-03-10\~experimental#get-/groups/-group\_id-/issues) | Non-GA REST   | October 22,  2024 | [rest-issues-experimental-api-to-ga-api-migration-guide.md](guides-to-migration/rest-issues-experimental-api-to-ga-api-migration-guide.md "mention") |
| [Retrieve a single project](https://snyk.docs.apiary.io/#reference/projects/individual-project/retrieve-a-single-project)                                                                                                                   | v1            | January 22, 2025  |                                                                                                                                                      |
| [Update a Project](https://snyk.docs.apiary.io/#reference/projects/individual-project/update-a-project)                                                                                                                                     | v1            | January 22, 2025  |                                                                                                                                                      |
| [Delete a Project](https://snyk.docs.apiary.io/#reference/projects/individual-project/delete-a-project)                                                                                                                                     | v1            | January 22, 2025  |                                                                                                                                                      |
| [Deactivate a single Project](https://snyk.docs.apiary.io/#reference/projects/deactivate-an-individual-project/deactivate)                                                                                                                  | v1            | January 22, 2025  |                                                                                                                                                      |
| [Activate a single Project](https://snyk.docs.apiary.io/#reference/projects/activate-an-individual-project/activate)                                                                                                                        | v1            | January 22, 2025  |                                                                                                                                                      |
| [Add a tag to a Project](https://snyk.docs.apiary.io/#reference/projects/project-tags/add-a-tag-to-a-project)                                                                                                                               | v1            | January 22, 2025  |                                                                                                                                                      |
| [Remove a tag from a Project](https://snyk.docs.apiary.io/#reference/projects/remove-project-tag/remove-a-tag-from-a-project)                                                                                                               | v1            | January 22, 2025  |                                                                                                                                                      |
| [Applying attributes to a Project](https://snyk.docs.apiary.io/#reference/projects/project-attributes/applying-attributes)                                                                                                                  | v1            | January 22, 2025  |                                                                                                                                                      |

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
