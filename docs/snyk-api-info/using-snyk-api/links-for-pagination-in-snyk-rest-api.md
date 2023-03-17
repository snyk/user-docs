# Links for pagination in Snyk REST API

The following explains how Snyk sets values for the REST API response parameters `starting_after` and `ending_before` and how to use these values in API calls.

When a call is made to an endpoint and there are more items in the response than can fit on one page, there are previous and next links in that response. There you can see what the values are for the `starting_after` and `ending_before` parameters. You can also see these values in examples on apidocs.snyk.io.

The following example was generated using a demonstration project with many vulnerabilities, allowing for showing pagination.

The first step was to invoke the REST API endpoint with a “normal” configuration, to see the code results displayed on “page 1" (that is, the first 10 code-related vulnerabilities):

```
curl -X GET "https://api.snyk.io/rest/orgs/6391f850-81f8-48fc-
9cd4-aa8c186c6ff0/issues?project_id=c877b914-4de3-
41e7-856e-f46a64d3f43d&version=2022-04-06%7Eexperimental" \
 -H "Accept: application/vnd.api+json" \
 -H "Authorization: Token <my-service-account-token>" 
```

The results are in [code-results-page1.json](links-for-pagination-in-snyk-rest-api.md#undefined). The last couple of lines contain a link to the “next page” (to the next 10 vulnerabilities):

```
"links": {
      "next": "/orgs/6391f850-81f8-48fc-9cd4-
      aa8c186c6ff0/issues?project_id=a6f059da-418b-484d-
      92a3-
      a38ba4d0ac02&type%5B0%5D=code&limit=10&version=2022
      -04-
      06~experimental&starting_after=v1.eyJwcmlvcml0eVNjb
      3JlIjo4MDEsImlkIjoiNzAzOTI2NjQ0In0%3D"
          }
```

The key index (“pointer”) is `v1.eyJwcmlvcml0eVNjb3JlIjo4MDEsImlkIjoiNzAzOTI2NjQ0In0.`

The second API call used this “pointer” to jump to the second page:

```
curl -X GET "https://api.snyk.io/rest/orgs/6391f850-81f8-48fc-
9cd4-aa8c186c6ff0/issues?project_id=c877b914-4de3-
41e7-856e-
f46a64d3f43d&starting_after=v1.eyJwcmlvcml0eVNjb3Jl
Ijo4MDEsImlkIjoiNzAzOTI2NjQ0In0&version=2022-04-
06%7Eexperimental" \
 -H "Accept: application/vnd.api+json" \
 -H "Authorization: Token <my-service-account-token>" 
```

The results are in [code-results-page2.json](links-for-pagination-in-snyk-rest-api.md#code-results-page2.json). The results end with two links to go back to “page 1” or go to page 3 (forward to see the next 10 vulnerabilities):

```
"links": {
      "next": "/orgs/6391f850-81f8-48fc-9cd4-
      aa8c186c6ff0/issues?project_id=a6f059da-418b-484d-
      92a3-
      a38ba4d0ac02&type%5B0%5D=code&limit=10&version=2022-04-
      06~experimental&starting_after=v1.eyJwcmlvcml0eVNjb
      3JlIjo4MDEsImlkIjoiNzAzOTI2NjM0In0%3D",
             "prev": "/orgs/6391f850-81f8-48fc-9cd4-
      aa8c186c6ff0/issues?project_id=a6f059da-418b-484d-
      92a3-
      a38ba4d0ac02&type%5B0%5D=code&limit=10&version=2022-04-
      06~experimental&ending_before=v1.eyJwcmlvcml0eVNjb3
      JlIjo4MDEsImlkIjoiNzAzOTI2NjQzIn0%3D"
          }
```

The pointer to go to page 3 is `v1.eyJwcmlvcml0eVNjb3JlIjo4MDEsImlkIjoiNzAzOTI2NjM0In0.` The pointer to go back to page 1 is `v1.eyJwcmlvcml0eVNjb3JlIjo4MDEsImlkIjoiNzAzOTI2NjQzIn0.`

## code-results-page1.json

```
{
    "jsonapi": {
      "version": "1.0"
    },
    "data": [
      {
        "type": "issue_summary",
        "id": "d6023fb9-15ce-435c-9c82-7add1f8b03c1",
        "attributes": {
          "issueType": "code",
          "title": "Use of Hardcoded Credentials",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-798",
            "CWE-259"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/d6023fb9-15ce-435c-9c82-7add1f8b03c1?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "958664e3-d66f-4da7-8607-cba54ec71fbc",
        "attributes": {
          "issueType": "code",
          "title": "Use of Hardcoded Credentials",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-798",
            "CWE-259"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/958664e3-d66f-4da7-8607-cba54ec71fbc?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "958664e3-d66f-4da7-8607-cba54ec71fbc",
        "attributes": {
          "issueType": "code",
          "title": "Use of Hardcoded Credentials",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-798",
            "CWE-259"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/958664e3-d66f-4da7-8607-cba54ec71fbc?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "958664e3-d66f-4da7-8607-cba54ec71fbc",
        "attributes": {
          "issueType": "code",
          "title": "Use of Hardcoded Credentials",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-798",
            "CWE-259"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/958664e3-d66f-4da7-8607-cba54ec71fbc?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "958664e3-d66f-4da7-8607-cba54ec71fbc",
        "attributes": {
          "issueType": "code",
          "title": "Use of Hardcoded Credentials",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-798",
            "CWE-259"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/958664e3-d66f-4da7-8607-cba54ec71fbc?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "958664e3-d66f-4da7-8607-cba54ec71fbc",
        "attributes": {
          "issueType": "code",
          "title": "Use of Hardcoded Credentials",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-798",
            "CWE-259"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/958664e3-d66f-4da7-8607-cba54ec71fbc?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "958664e3-d66f-4da7-8607-cba54ec71fbc",
        "attributes": {
          "issueType": "code",
          "title": "Use of Hardcoded Credentials",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-798",
            "CWE-259"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/958664e3-d66f-4da7-8607-cba54ec71fbc?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "591b4905-d3c8-4ef6-bd02-f13a189f6da7",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/591b4905-d3c8-4ef6-bd02-f13a189f6da7?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "a58edc66-1287-406b-8ce1-05f4f7ce852b",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/a58edc66-1287-406b-8ce1-05f4f7ce852b?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "7e7b3390-d7c0-4dc4-b9b4-56f491301493",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/7e7b3390-d7c0-4dc4-b9b4-56f491301493?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      }
    ],
    "links": {
      "next": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&type%5B0%5D=code&limit=10&version=2022-04-06~experimental&starting_after=v1.eyJwcmlvcml0eVNjb3JlIjo4MDEsImlkIjoiNzAzOTI2NjQ0In0%3D"
    }
  }
```

## code-results-page2.json

```
{
    "jsonapi": {
      "version": "1.0"
    },
    "data": [
      {
        "type": "issue_summary",
        "id": "aee3c766-8e59-48f6-9624-fcd05f541104",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/aee3c766-8e59-48f6-9624-fcd05f541104?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "54e69267-0089-4ef8-8db2-825301e554e8",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/54e69267-0089-4ef8-8db2-825301e554e8?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "54e69267-0089-4ef8-8db2-825301e554e8",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/54e69267-0089-4ef8-8db2-825301e554e8?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "2b03e588-e62d-4fd8-867e-7491865e5af4",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/2b03e588-e62d-4fd8-867e-7491865e5af4?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "54e69267-0089-4ef8-8db2-825301e554e8",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/54e69267-0089-4ef8-8db2-825301e554e8?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "54e69267-0089-4ef8-8db2-825301e554e8",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/54e69267-0089-4ef8-8db2-825301e554e8?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "2b03e588-e62d-4fd8-867e-7491865e5af4",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/2b03e588-e62d-4fd8-867e-7491865e5af4?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "3e0ebd56-1704-4d6a-917c-caab57f8892f",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/3e0ebd56-1704-4d6a-917c-caab57f8892f?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "7af39c98-0822-405e-8d18-90810e73a66f",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/7af39c98-0822-405e-8d18-90810e73a66f?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      },
      {
        "type": "issue_summary",
        "id": "3e0ebd56-1704-4d6a-917c-caab57f8892f",
        "attributes": {
          "issueType": "code",
          "title": "Buffer Overflow",
          "severity": "high",
          "ignored": false,
          "cwe": [
            "CWE-122"
          ]
        },
        "links": {
          "self": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues/detail/code/3e0ebd56-1704-4d6a-917c-caab57f8892f?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&version=2022-04-06~experimental"
        }
      }
    ],
    "links": {
      "next": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&type%5B0%5D=code&limit=10&version=2022-04-06~experimental&starting_after=v1.eyJwcmlvcml0eVNjb3JlIjo4MDEsImlkIjoiNzAzOTI2NjM0In0%3D",
      "prev": "/orgs/6391f850-81f8-48fc-9cd4-aa8c186c6ff0/issues?project_id=a6f059da-418b-484d-92a3-a38ba4d0ac02&type%5B0%5D=code&limit=10&version=2022-04-06~experimental&ending_before=v1.eyJwcmlvcml0eVNjb3JlIjo4MDEsImlkIjoiNzAzOTI2NjQzIn0%3D"
    }
  }
  
```
