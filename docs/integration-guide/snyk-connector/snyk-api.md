# Snyk API

This section will provide a quick overview of getting started with the Snyk API. For full details please visit the [Snyk API documentation](https://www.notion.so/snyk/How-to-Build-a-Snyk-Connector-0d54bd5d3c7f4baab49b656dfffcded0#d2d9bb65fddb4cd783d2cf4e0a99e2e4).

## Overview of common Snyk API calls

The Snyk data model listed below is straight forward but accessing the model via the API requires a quick orientation to the APIs. To obtain the organization and project IDs you will want to complete the following steps.

1. Call the Snyk API to [list all of the organizations associated with a user](https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to). Depending on the `API_TOKEN` provided, you will receive a list of organization IDs. If you provide a service account associated with a Group, you will receive a list of organization IDs based on the number of organizations. If you provide a service account associated with an organization, you will receive a single organization ID to use in the next step.
2. Call the Snyk API to [get a list of projects](https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects). This call takes an `API_TOKEN` and an organization ID to get a list of projects.
3. Call the Snyk API to [get an aggregated list of issues](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-issues/list-all-aggregated-issues) per project.

### Snyk API Token

A Snyk `API_TOKEN` authenticates the user to the Snyk API. Customers can provide a personal or a service account `API_TOKEN` \(Snyk Pro and Enterprise customers only\). This token should be requested during the initial configuration of the Snyk Connector. For details on finding a personal token please visit the [personal tokens documentation](https://support.snyk.io/hc/en-us/articles/360004037557-Authentication-for-API). For details on setting up service accounts please visit the [service account documentation](https://support.snyk.io/hc/en-us/articles/360004037597-Service-accounts).

### Snyk API Endpoint and Authorization

The Snyk API endpoint and authorization information can be found on the [Snyk API documentation](https://snyk.docs.apiary.io/#introduction/api-url) page. This page provides details on the endpoint and how to use your `API_TOKEN` to get authorized.

