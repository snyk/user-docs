# REST API endpoint: Get a project’s SBOM document

{% hint style="info" %}
**Feature availability**\
This feature is available to customers on Snyk Enterprise plans.
{% endhint %}

Snyk offers an endpoint to generate SBOM documents for Open Source and Container Projects that are continuously being monitored for issues.&#x20;

The SBOM document represents the latest state of a project’s dependencies and their relationships.

SBOM documents can be generated in [CycloneDX](https://cyclonedx.org/) v1.4 (JSON, XML) and [SPDX](https://spdx.dev/) v2.3 (JSON) formats.

## How to generate the SBOM for a project

1. On the Snyk Web UI, retrieve your organization ID (UUID format), project ID (UUID) and API key.\
   If you need help in finding these values, see [Group and Organization navigation](../snyk-admin/manage-groups-and-organizations/group-and-organization-navigation.md), [View project settings](../manage-issues/snyk-projects/view-and-edit-project-settings.md), and [Authentication for API](authentication-for-api.md).
2. Determine the format you want for the SBOM you will generate.\
   Available options are CycloneDX 1.4 JSON (`cyclonedx1.4+json`), CycloneDX 1.4 XML (`cyclonedx1.4+xml`) or SPDX v2.3 JSON (`spdx2.3+json`).
3. Using any HTTP client (for example, Postman, curl) make a request to the endpoint.\
   Note that the `format` parameter must be URL-encoded.\
   Example: To retrieve a CycloneDX 1.4 JSON document, set `format=cyclonedx1.4%2Bjson` on the query.

```bash
$ curl --get \
  -H "Authorization: token <SNYK_API_TOKEN>" \
  --data-urlencode "version=2023-03-20" \
  --data-urlencode "format=<SBOM_FORMAT>" \
  https://api.snyk.io/rest/orgs/<ORG_ID>/projects/<PROJECT_ID>/sbom
```

## Custom CycloneDX properties

A Snyk project’s SBOM document will include some Snyk-specific meta data about the SBOM document. This is included in the `metadata` section of the document when exported as CycloneDX.

<table><thead><tr><th width="225">Property Name</th><th>Description</th></tr></thead><tbody><tr><td><code>snyk:org_id</code></td><td>The organization ID (UUID) to which the project belongs</td></tr><tr><td><code>snyk:project_id</code></td><td>The project’s ID (UUID)</td></tr></tbody></table>

## Troubleshooting for Get a project's SBOM document endpoint

The following response code indicates success.

**200 OK**

The SBOM document was successfully generated. The response body contains the document in the requested format.

The following are **error states** that you may receive when using the API. If you experience issues not covered here or are having trouble resolving these, contact your Solution Engineer or Technical Success Manager, or submit a request to [Snyk Support](https://support.snyk.io/hc/en-us/requests/new).

**401 Unauthorized**

The authentication method, API token for Bearer token, was invalid. Check that you set the Authorization header correctly.

**403 Forbidden**

You do not have the permissions required to make the request. This can happen if you are not part of the requested Organization, your Organization is not entitled to use the Snyk API, or you do not have sufficient read access to the requested Project.

**429 Too Many Requests**

Since the Snyk API is rate-limited, an excessive number of requests will eventually start to be rejected. You need to wait before making any further requests.

**500 Internal Server Error**

The service encountered an internal system error.
