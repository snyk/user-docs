# Issues: List issues for a package

The Snyk REST API endpoint [List issues for a package](../reference/issues.md#orgs-org_id-packages-purl-issues) can be used to get all direct (non-transitive) vulnerabilities for a package using its `purl`, which is a uniform way of identifying software packages across ecosystems as defined in the [package URL specification](https://github.com/package-url/purl-spec).

When you pass a `purl` to the endpoint, Snyk will find any known vulnerabilities for that package and return them as part of the response body.

The API is useful when you have a list of packages and want to retrieve a list of vulnerabilities for a package version.

{% hint style="info" %}
The examples on this page use [HTTPie](https://httpie.io/), but you can use any HTTP client to access the Snyk REST API.
{% endhint %}

## Supported purl types

The current release supports the following `purl` types: `apk`, `cargo`, `cocoapods`, `conan`, `composer`, `deb`, `gem`, `generic`, `golang`, `hex`, `npm`, `nuget`, `pub`, `pypi`, `rpm`, `swift` and `maven`.

If you are interested in support for additional ecosystems, submit a request to [Snyk Support](https://support.snyk.io).

## Request for List issues for a package endpoint

To call the API endpoint, **use the following HTTP request**:

```bash
$ http \
  "https://api.snyk.io/rest/orgs/{org_id}/packages/{purl}/issues" \
  "Authorization: token $API_TOKEN" \
  version==<snyk-api-version>
```

The purl must be **URL-encoded.**

An example using a valid url-encoded purl follows:

```bash
$ http \
  "https://api.snyk.io/rest/orgs/{org_id}/packages/pkg%3Amaven%2fcom.fasterxml.woodstox%2fwoodstox-core%405.0.0/issues" \
  "Authorization: token $API_TOKEN" \
  version==2024-06-26
```

For operating system packages, a vendor must be specified in the namespace portion, and a `distro` qualifier must be specified. Supported vendors include: `debian`, `alpine`, `rhel`, `ubuntu`, `amzn`, `centos`, `oracle`, `rocky`, `sles`.

An example using a valid url-encoded operating system purl follows:

```bash
$ http \
  "https://api.snyk.io/rest/orgs/{org_id}/packages/pkg%3Adeb%2Fdebian%2Fcurl%3Fdistro%3Dbullseye/issues" \
  "Authorization: token $API_TOKEN" \
  version==2024-06-26
```

The Snyk REST API supports pagination. This has a default page limit of **1000**, with a default offset of **0.** Current, next, and previous pages are returned as links in the response. The following parameters can be supplied as query parameters: `offset`, `limit`.

An example paginated request follows:

```bash
$ http \
  "https://api.snyk.io/rest/orgs/{org_id}/packages/pkg%3Amaven%2fcom.fasterxml.woodstox%2fwoodstox-core%405.0.0/issues" \
  "Authorization: token $API_TOKEN" \
  version==2024-06-26 \
  limit==100 \
  offset==0
```

## Response for List issues for a package endpoint

The expected output provides a [JSON API](https://jsonapi.org/format/) response that identifies the vulnerabilities associated with the package.

The following example gives the response for a `pypi` package [django](https://security.snyk.io/package/pip/django).

The response provides a list of the vulnerabilities found for the package identified by the purl in the request. The response begins with a description of a vulnerability:

### **Overview of package**

Affected versions of this package are vulnerable to Denial of Service (DoS) through very large inputs with a specific sequence of characters in the `urlize()` and `urlizetrunc()` template filters.

### **Remediation**

Upgrade the package version to 4.2.15,5.0.8 to fix this vulnerability

### Response details

{% hint style="info" %}
The response is continuous, divided here to allow for explanations.
{% endhint %}

**For each vulnerability**, the response provides the following:

*   The Snyk issue ID and issue types:

    ```json
    "id": "SNYK-PYTHON-DJANGO-7642790",
    "type": "issue",
    ```
*   General metadata about the vulnerability, including title, timestamps relevant to the vulnerability such as publication and disclosure time, and description:

    ```json
    "title": "Denial of Service (DoS)",
    "type": "package_vulnerability",
    "created_at": "2024-08-07T08:13:29.424951Z",
    "updated_at": "2024-08-08T13:36:35.964359Z",
    "description": ...
    ```
*   The CVE and CWE identifiers:

    ```json
    "problems": [
        {
            "id": "CVE-2024-41990",
            "source": "CVE"
        },
        {
            "id": "CWE-400",
            "source": "CWE"
        }
    ],
    ```
*   The severity or severities of the vulnerability:

    ```json
    "severities": [
        {
            "type": "primary",
            "source": "Snyk",
            "level": "medium",
            "score": 6.9,
            "version": "4.0",
            "vector": "CVSS:4.0/AV:N/AC:L/AT:N/PR:N/UI:N/VC:N/VI:N/VA:L/SC:N/SI:N/SA:N"
        },
        {
            "type": "secondary",
            "source": "Snyk",
            "level": "medium",
            "score": 5.3,
            "version": "3.1",
            "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:L"
        },
        {
            "type": "secondary",
            "source": "NVD",
            "level": "high",
            "score": 7.5,
            "version": "3.1",
            "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H"
        },
        {
            "type": "secondary",
            "source": "Red Hat",
            "level": "high",
            "score": 7.5,
            "version": "3.1",
            "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H"
        },
        {
            "type": "secondary",
            "source": "SUSE",
            "level": "high",
            "score": 7.5,
            "version": "3.1",
            "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H"
        }
    ],
    ```

{% hint style="info" %}
Starting in June 2024, all new advisories identified by Snyk Open Source are provided with both CVSS v4.0 and CVSS v3.1 severities. For the most accurate severity assessments, Snyk recommends using CVSS v4.0 when it is available.
{% endhint %}

*   Any fixes available for that vulnerability and the representation for vulnerable versions:

    ```json
    "coordinates": [
        {
            "remedies": [
                {
                    "type": "indeterminate",
                    "description": "Upgrade the package version to 4.2.15,5.0.8 to fix this vulnerability",
                    "details": {
                        "upgrade_package": "4.2.15,5.0.8"
                    }
                }
            ],
            "representations": [
                {
                    "resource_path": "[,4.2.15),[5.0,5.0.8)"
                },
                {
                    "package": {
                        "name": "django",
                        "version": "4.2.14",
                        "type": "pypi",
                        "url": "pkg:pypi/django@4.2.14"
                    }
                }
            ]
        }
    ],
    ```
* Links to any external resources with further information on the vulnerability:

```json
"references": [
 {
    "url": "https://www.djangoproject.com/weblog/2024/aug/06/security-releases/",
    "title": "Django Security Release"
},
    ... 
```

**Package metadata** is returned, including the following:

* Package name
* Package type
* Package URL specification
* Package version

```json
"meta": {
    "package": {
        "name": "django",
        "type": "pypi",
        "url": "pkg:pypi/django@4.2.14",
        "version": "4.2.14"
    }
}
```

Where applicable, **pagination links for the results** are included as follows:

* Next path (if applicable)
* Previous path (if applicable)
* Current path

```json
"links": {
    "prev": "/orgs/<org-id>/packages/{purl}/issues?version=<api-version>&limit=1000&offset=0",
    "self": "/orgs/<org-id>/packages/{purl}/issues?version=<api-version>&limit=1000&offset=1"
},
```

## Troubleshooting for the List issues for a package endpoint

The following are **error states** that you may receive when using the API. If you experience issues not covered here or are having trouble resolving these, contact your Solution Engineer or Technical Success Manager, or submit a request to [Snyk Support](https://support.snyk.io).

**Invalid PURL**\
400\
Ensure that the purl specification you provided is a valid purl. For more information, see the [Package URL specification](https://github.com/package-url/purl-spec).

**Unsupported Ecosystem**\
400\
Ensure that the package type is one of the [supported purl types](issues-list-issues-for-a-package.md#supported-purl-types).

**Package requested without namespace**\
400\
Ensure you specify a namespace in the package URL and then try again. For more information, see the [Package URL specification](https://github.com/package-url/purl-spec).

**Purl component not supported**\
400\
Remove the component that is not supported and try to make the request again. The endpoint accepts only particular components. For more information, see the [Package URL specification](https://github.com/package-url/purl-spec).

**Your organization is not authorized to perform this action.**\
403\
To get access, contact your Solutions Engineer or Technical Success Manager, or the administrator for your team.

**Rate limit exceeded**\
429\
180 requests per minute per user are permitted on this API endpoint. If you exceed this volume, you will receive a 429 error response code.

**Invalid pagination parameters**\
400\
The limit and offset supplied query parameters must be as follows:

* Limit > 0 and <= 1000
* Offset >= 0

The default parameters for this request are limit = 1000, and offset >= 0.

**Authorization request failure**\
500\
This issue is unexpected, and the service should recover quickly. If it does not, submit a request to [Snyk Support](https://support.snyk.io).

**Internal server error**\
500\
This issue is unexpected, and the service should recover quickly. If it does not, submit a request to [Snyk Support](https://support.snyk.io).

**Vulnerability service unavailable**\
503\
This issue is unexpected, and the service should recover quickly. If it does not, submit a request to [Snyk Support](https://support.snyk.io).

**Vulnerability service error**\
500\
This issue is unexpected, and the service should recover quickly. If it does not, submit a request to [Snyk Support](https://support.snyk.io).

An example of this error response follows:

```json
"jsonapi": {
    "version": "1.0"
},
"errors": [{
    "id": "8139dce7-eeb4-404b-be0e-5e4c15004972",
    "detail": "Unsupported Ecosystem",
    "status": "400"
}]
```
