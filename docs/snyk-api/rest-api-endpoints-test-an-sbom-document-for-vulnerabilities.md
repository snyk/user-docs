# REST API endpoints: Test an SBOM document for vulnerabilities

{% hint style="warning" %}
**Release status**

REST API is in [Early Access](../getting-started/snyk-release-process.md) and available only for Enterprise plans.

See [Pricing plans](https://snyk.io/plans).
{% endhint %}

Snyk offers a [collection of API endpoints](https://apidocs.snyk.io/?version=2023-10-13%7Ebeta#post-/orgs/-org\_id-/sbom\_tests) to asynchronously test a software bill of materials (SBOM) document. You can use these endpoints to learn more about the vulnerabilities impacting your SBOM and its packages.

{% hint style="info" %}
Supported SBOM formats are [CycloneDX](https://cyclonedx.org/) 1.4 JSON and [SPDX](https://spdx.dev/) 2.3 JSON.
{% endhint %}

Snyk identifies components within the SBOM by their [package URL](https://github.com/package-url/purl-spec) (purl). If a component does not contain a purl or the purl type is not supported, Snyk skips vulnerability analysis for that component. Supported purl types are: `cargo`, `cocoapods`, `gem`, `golang`, `hex`, `maven`, `npm`, `nuget`, `pypi`, `swift`, and `generic` for unmanaged C/C++ dependencies.

Follow these steps to [create an SBOM test run](https://apidocs.snyk.io/?version=2023-10-13%7Ebeta#post-/orgs/-org\_id-/sbom\_tests) and view the results.

1. [Create the test by sending an SBOM to Snyk.](rest-api-endpoints-test-an-sbom-document-for-vulnerabilities.md#create-a-test-by-sending-an-sbom-to-snyk)
2. [Check the status of the test](rest-api-endpoints-test-an-sbom-document-for-vulnerabilities.md#check-the-status-of-the-test-optional).
3. [View the test results when the test is complete.](rest-api-endpoints-test-an-sbom-document-for-vulnerabilities.md#view-results-of-the-test)

## How to test an SBOM document

### Create a test by sending an SBOM to Snyk&#x20;

Testing your SBOM can be a long-running operation. Instead of waiting until the test results are ready, Snyk returns a `job_id` after your initial request to send the SBOM, and then processes the request asynchronously.

Follow these steps to test a SCOM:

1. Log in to the Snyk Web UI and retrieve your Organization ID (UUID format), Project ID (UUID), and API key.\
   If you need help in finding these values, see [Group and Organization navigation](../snyk-admin/manage-groups-and-organizations/switch-between-groups-and-organizations.md), [View Project settings](../snyk-admin/snyk-projects/view-and-edit-project-settings.md), and [Authentication for API](../snyk-api-info/authentication-for-api.md).
2. Use any HTTP client, for example, `curl` or Postman, to make a request to the endpoint [Create an SBOM test run](https://apidocs.snyk.io/?version=2023-10-24%7Ebeta#post-/orgs/-org\_id-/sbom\_tests).&#x20;

{% hint style="info" %}
The SBOM document is included as part of the request body as a JSON object. This request creates a test run for your SBOM document.
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

3. From the response, get the `job_id`, which is used in the next steps. \
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

You can check the status of the test at any time after the initial request. &#x20;

1. Using the `job_id` returned from the initial request to the [Create an SBOM test run endpoint](https://apidocs.snyk.io/?version=2023-10-24%7Ebeta#post-/orgs/-org\_id-/sbom\_tests), make a request to another endpoint to get the [SBOM test run status](https://apidocs.snyk.io/?version=2023-10-24%7Ebeta#get-/orgs/-org\_id-/sbom\_tests/-job\_id-).&#x20;
2. A successful request to this endpoint returns the status of your test, which can either be `processing` or `finished`. If the call is not successful, an error will be returned.

```bash
  curl --get \
      -H "Authorization: token <SNYK_TOKEN>" \
      'https://api.snyk.io/rest/orgs/<ORG_ID>/sbom_tests/<TEST_ID>?version=2023-08-31~beta'
```

### View results of the test

When the test is complete, you can view the results for the tested SBOM.

1. When the status of the test returned is`finished`, make a request to [get an SBOM test result](https://apidocs.snyk.io/?version=2023-10-24%7Ebeta#get-/orgs/-org\_id-/sbom\_tests/-job\_id-/results).
2. View the information that the request returns: summary-level information about the SBOM that was tested, as well as the detailed results.

```bash
curl --get \
    -H "Authorization: token <SNYK_TOKEN>" \
    'https://api.snyk.io/rest/orgs/<ORG_ID>/sbom_tests/<TEST_ID>/results?version=2023-08-31~beta'
```

## Troubleshooting for the endpoint Create an SBOM test run

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
