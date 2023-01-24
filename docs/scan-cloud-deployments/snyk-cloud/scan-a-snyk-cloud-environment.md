# Scan a Snyk Cloud Environment

Snyk automatically runs a scan when a [Snyk Cloud Environment](snyk-cloud-concepts.md#environments) is created. After that, you can manually trigger a new scan by using the [Snyk API](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#post-/orgs/-org\_id-/cloud/scans).

## Find the environment ID

First, find the ID of the Snyk Cloud Environment you want to scan. Send a request to the [`/cloud/environments`](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#get-/orgs/-org\_id-/cloud/environments) endpoint in the below format:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments?version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

In the output, look for the `data.id` property. In the shortened example below, the ID is `3b7ccff9-8900-4e54-0000-1234abcd1234`:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "3b7ccff9-8900-4e54-0000-1234abcd1234",
    <trimmed for length>
  }
}
```

## Trigger the scan

To manually trigger a scan, send a request to the [`/cloud/scans`](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#post-/orgs/-org\_id-/cloud/scans) endpoint in the below format:

```
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/scans?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H "Content-Type:application/vnd.api+json"  -d '{
  "data": {
    "relationships": {
      "environment": {
        "data": {
          "id": "YOUR-ENVIRONMENT-ID",
          "type": "environment"
        }
      }
    },
    "type": "resource"
  }
}'
```

Note: The example above uses [curl](https://curl.se/), but you can use any API client, such as [Postman](https://www.postman.com/) or [HTTPie](https://httpie.io/).

## Understand the API response

Snyk returns a JSON document containing details about the new scan. For example:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "a7fa2167-58a8-4ac5-9999-0987dcba6543",
    "type": "scan",
    "attributes": {
      "created": "2022-08-07T04:59:58.639423469Z",
      "updated": null,
      "finished": null,
      "revision": 2,
      "kind": "user_initiated",
      "status": "queued"
    },
    "relationships": {
      "environment": {
        "data": {
          "id": "3b7ccff9-8900-4e54-0000-1234abcd1234",
          "type": "environment"
        },
        "links": {
          "related": "/orgs/d70c1768-5675-0000-1234-abcd1234abcd/cloud/environments?id=3b7ccff9-8900-4e54-0000-1234abcd1234&version=2022-12-21~beta"
        }
      },
      "organization": {
        "data": {
          "id": "d70c1768-5675-0000-1234-abcd1234abcd",
          "type": "organization"
        },
        "links": {
          "related": "/orgs/d70c1768-5675-0000-1234-abcd1234abcd?version=2022-12-21~beta"
        }
      }
    }
  }
}
```

Below are some key attributes from the API response:

* `data.id`: Scan ID
* `data.attributes.status`: Scan status

## Check scan status

To check a scan's status, retrieve the details of the environment being scanned. Send a request to the [`/cloud/environments`](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#get-/orgs/-org\_id-/cloud/environments) endpoint in the below format:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments?id=YOUR-ENVIRONMENT-ID&version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

Snyk returns a JSON document containing environment details. Look for the `data.attributes.status` value to find the scan status. In the shortened example below, the status is `success`:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": [
    {
      "id": "3b7ccff9-8900-4e54-0000-1234abcd1234",
      "type": "environment",
      "attributes": {
        "status": "success",
        <trimmed for length>
      }
    }
  ]
}
```

Scan status values:

* `queued`: Scan is about to start
* `in_progress`: Scan is in progress
* `success`: Scan is completed
* `error`: Scan errored; wait a moment and try scanning again

## View all scans for an organization

To view all scans for an organization, send an [API request](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#get-/orgs/-org\_id-/cloud/scans) in the below format:

```
curl -X GET \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/scans?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN'
```

Snyk returns a JSON document containing details about all scans. For example:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": [
    {
      "id": "a7fa2167-58a8-4ac5-9999-0987dcba6543",
      "type": "scan",
      "attributes": {
        "created_at": "2022-08-04T22:14:47Z",
        "error": "",
        "finished_at": "2022-08-04T22:16:31Z",
        "kind": "user_initiated",
        "options": {
          "role_arn": "arn:aws:iam::123456789012:role/snyk-cloud-role"
        },
        "revision": 2,
        "status": "success",
        "updated_at": "2022-08-04T22:16:31Z"
      },
      "relationships": {
        "environment": {
          "data": {
            "id": "3b7ccff9-8900-4e54-0000-1234abcd1234",
            "type": "environment"
          },
          "links": {
            "related": "/orgs/d70c1768-5675-0000-1234-abcd1234abcd/cloud/environments?id=3b7ccff9-8900-4e54-0000-1234abcd1234&version=2022-12-21~beta"
          }
        },
        "organization": {
          "data": {
            "id": "d70c1768-5675-0000-1234-abcd1234abcd",
            "type": "organization"
          },
          "links": {
            "related": "/orgs/d70c1768-5675-0000-1234-abcd1234abcd?version=2022-12-21~beta"
          }
        }
      }
    }
    <trimmed for length>
  ]
}
```
