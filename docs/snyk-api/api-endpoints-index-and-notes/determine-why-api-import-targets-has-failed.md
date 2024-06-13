---
description: Determine why API import targets has failed
---

# Determine why API Import targets has failed

The Snyk API v1 endpoint [Import targets](https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets) can be used to import Snyk Projects. If this fails, use [Get import job details](https://snyk.docs.apiary.io/#reference/import-projects/import-job/get-import-job-details) to help determine why.

There are two types of failures:

* The repository was rejected for processing, that is, HTTP status code 201 was not returned. This happens if there is an issue Snyk can see quickly for example:
  * The repository does not exist.
  * The repository is unreachable by Snyk because the token is invalid or does not have sufficient permissions; there is no default branch.
* The repository was accepted for processing, that is, the user got back HTTP status code 201 and a url to poll, but no projects were detected or some failed. This may occur because:
  * There are no Snyk-supported manifests in this repository.
  * The repository is archived and the Snyk API calls to fetch files fail.
  * The individual project or manifest had issues during processing. In this case Snyk returns success: false with a message in the log.

The poll results return a message per manifest processed, either `success: true` or `success: false.`
