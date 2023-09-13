# REST API endpoints: Test an SBOM document for vulnerabilities

{% hint style="info" %}
**Feature availability**\
This feature is available to customers on [Snyk Enterprise](../enterprise-setup/getting-started-with-the-snyk-enterprise-plan.md) plans.
{% endhint %}

{% hint style="warning" %}
This API endpoint, Test an SBOM document for vulnerabilities, is available on an open beta basis. Some of the functionality may change.
{% endhint %}

Snyk offers a set of endpoints for testing a software bill of materials document. Use these endpoints to expand your understanding of the vulnerabilities impacting the packages in an SBOM.

{% hint style="info" %}
SBOM documents in [CycloneDX](https://cyclonedx.org/) 1.4 JSON formats are supported.
{% endhint %}

This endpoint, Test an SBOM document for vulnerabilities, is asynchronous. Follow these steps to create an SBOM test run:

1. Create the test by sending an SBOM to Snyk.
2. [Check the status of the test](rest-api-endpoints-test-an-sbom-document-for-vulnerabilities.md#check-the-status-of-the-test-optional).
3. View the test results when the test is complete.

## How to test an SBOM document

### Create a test by sending an SBOM to Snyk&#x20;

1. Log in to the Snyk Web UI and retrieve your Organization ID (UUID format), Project ID (UUID), and API key.\
   If you need help in finding these values, see [Group and Organization navigation](../snyk-admin/manage-groups-and-organizations/group-and-organization-navigation.md), [View project settings](../manage-risk/snyk-projects/view-and-edit-project-settings.md), and [Authentication for API](../snyk-api-info/authentication-for-api.md).
2. Use any HTTP client, for example, `curl` or Postman, to make a request to the endpoint.&#x20;

{% hint style="info" %}
The SBOM document is included as part of the request body as a JSON object. This creates a test run for your SBOM document.
{% endhint %}

{% code title="HTTP request" %}
```bash
  curl --request POST \
  -H "Authorization: token <SNYK_TOKEN>" \
  -H "Content-Type: application/vnd.api+json" \
  --data-binary '@request_body.json' \
  'https://api.snyk.io/rest/orgs/<ORG_ID>/sbom_tests?version=2023-08-31~beta'
```
{% endcode %}

{% code title="request_body.json" %}
```json
{
    "data": {
        "type": "sbom_test",
        "attributes":{ 
            "sbom": {
            <SBOM_CONTENTS>
            }
        }
    }
}
```
{% endcode %}

3. From the response, get the job ID, which is used in the next steps. \
   This is a unique identifier for the test run being performed on your SBOM document.

{% code title="JSON response body" %}
```json
{
    "data": {
        "id": "<JOB_ID>",
        "type": "sbom_tests"
    },
    "jsonapi": {
        "version": "1.0"
    },
    "links": {
        "self": "/rest/orgs/<ORG_ID>/sbom_tests?version=2023-08-31~beta",
        "related": "/rest/orgs/<ORG_ID>/sbom_tests/<JOB_ID>?version=2023-08-31~beta"
    }
}
```
{% endcode %}

### Check the status of the test (optional)

1. Using the job ID returned from the initial request, make a request to the endpoint below.&#x20;
2. Note that the request returns the status of your test, which can either be `processing` or `finished`.

```bash
  curl --get \
  -H "Authorization: token <SNYK_TOKEN>" \
  'https://api.snyk.io/rest/orgs/<ORG_ID>/sbom_tests/<TEST_ID>?version=2023-08-31~beta'
```

### View results of the test

1. Once the status of the test is `finished`, make a request to the results endpoint.&#x20;
2. The request returns summary-level information about the SBOM that was tested, as well as the detailed results.

```bash
curl --get \
  -H "Authorization: token <SNYK_TOKEN>" \
  'https://api.snyk.io/rest/orgs/<ORG_ID>/sbom_tests/<TEST_ID>/results?version=2023-08-31~beta'
```

## Troubleshooting for Test an SBOM document for vulnerabilities

The following response code indicates success.

**201 Created**

The SBOM test run was successfully created. The response body contains the job ID of the test run.

The following are error states that you may receive when using the API. If you experience issues not covered here or are having trouble resolving these, contact your Solutions Engineer or Technical Success Manager, or submit a request to [Snyk Support](https://support.snyk.io/hc/en-us/requests/new).

**400 Bad Request**

The server cannot process the request due to invalid or corrupt data. Review the request and try again.

**401 Unauthorized**

The authentication method, API token or Bearer token, was invalid. Check that you set the Authorization header correctly.

**403 Forbidden**

You do not have the permissions required to make the request. This can happen if you are not part of the requested Organization, your Organization is not entitled to use the Snyk API, or you do not have sufficient read access to the requested Project.

**429 Too Many Requests**

Since the Snyk API is rate-limited, an excessive number of requests will eventually start to be rejected. You need to wait before making any further requests.

**500 Internal Server Error**

The service encountered an internal system error.
