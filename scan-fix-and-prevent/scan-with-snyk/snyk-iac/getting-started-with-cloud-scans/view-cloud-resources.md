# View cloud resources

You can view all attributes for cloud resources in an Organization. This allows you to inventory all of your resources across cloud provider accounts or see the recorded state of any resource during the most recent scan.

## View all resources

To list all resources in an Organization, send a request to the [`/cloud/resources`](https://apidocs.snyk.io/#get-/orgs/-org_id-/cloud/resources) endpoint in the following format:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/resources?version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

## Understand the API response

Snyk returns a JSON document containing information about all resources in the Organization, for example:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "next": "/rest/orgs/d70c1768-5675-4f44-0000-1234abcd1234/cloud/resources?starting_after=eyJpZCI6IjY5ODA5MjNhLWU0ZTAtNDg3Mi04ZDAwLWRjZDEXAMPLEEXAMPLE&version=2022-04-13~experimental"
  },
  "data": [
    {
      "id": "23b3a46d-cdf7-526c-8888-1abc2abc3abc",
      "type": "resource",
      "attributes": {
        "environment_id": "ef5d85de-fb4f-4c42-1234-000000000000",
        "scan_id": "44f386a6-6ce8-4303-0000-098765432109",
        "created_at": "2022-08-07T05:34:24.272279Z",
        "updated_at": "2022-08-07T05:34:24.272279Z",
        "revision": 1,
        "kind": "cloud",
        "hash": "717cdff4218bf3d6abcdefEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLE",
        "platform": "aws",
        "namespace": "us-east-1",
        "resource_type": "aws_s3_bucket",
        "resource_id": "example-bucket",
        "native_id": "arn:aws:s3:::example-bucket",
        "name": "example-bucket",
        "location": "us-east-1",
        "state": {
          "id": "example-bucket",
          "acl": "private",
          "arn": "arn:aws:s3:::example-bucket"
          <trimmed for length>
        },
        "tags": {}
      },
      "relationships": {
        "environment": {
          "data": {
            "id": "ef5d85de-fb4f-4c42-1234-000000000000",
            "type": "environment"
          },
          "links": {
            "related": "/orgs/d70c1768-5675-4f44-0000-1234abcd1234/cloud/environments?id=ef5d85de-fb4f-4c42-1234-000000000000&version=2022-12-21~beta"
          }
        },
        "organization": {
          "data": {
            "id": "d70c1768-5675-4f44-0000-1234abcd1234",
            "type": "organization"
          },
          "links": {
            "related": "/orgs/d70c1768-5675-4f44-0000-1234abcd1234?version=2022-12-21~beta"
          }
        },
        "scan": {
          "data": {
            "id": "a7fa2167-58a8-4ac5-9999-0987dcba6543",
            "type": "scan"
          },
          "links": {
            "related": "/orgs/d70c1768-5675-4f44-0000-1234abcd1234/cloud/scans?id=a7fa2167-58a8-4ac5-9999-0987dcba6543&version=2022-12-21~beta"
          }
        }
      }
    }
    <trimmed for length>
  ]
}
```

The following table lists some key attributes from the API response:

| Attribute                        | Definition                                                                                   |
| -------------------------------- | -------------------------------------------------------------------------------------------- |
| `data.id`                        | Resource ID                                                                                  |
| `data.attributes.environment_id` | ID of the environment containing the resource                                                |
| `data.attributes.scan_id`        | ID of the scan where this resource was first detected                                        |
| `data.attributes.resource_type`  | Resource type                                                                                |
| `data.attributes.native_id`      | Cloud provider's unique identifier for the resource; for AWS, the Amazon Resource Name (ARN) |
| `data.attributes.state`          | Resource attributes at the time of the most recent scan                                      |

## Filter resource list

Query parameters allow you to filter the list of resources.

For example, to return resources from a single environment, add `environment_id=YOUR-ENVIRONMENT-ID` to the URL:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/resources?environment_id=YOUR-ENVIRONMENT-ID&version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

Some parameters allow you to specify multiple values. To return resources in Amazon Web Services (AWS) regions `us-east-1` or `us-east-2`, add `location=us-east-1,us-east-2` to the URL:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/resources?location=us-east-1,us-east-2&version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

You can combine query parameters by using the `&` symbol. To return only 5 AWS S3 buckets, add `resource_type=aws_s3_bucket&limit=5` to the URL:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/resources?resource_type=aws_s3_bucket&limit=5&version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

For a list of supported parameters, see the [List resources API documentation](https://apidocs.snyk.io/#get-/orgs/-org_id-/cloud/resources).

## View details for a single resource

To view details for a single resource through the Snyk API, send a request in the following format.&#x20;

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/resources?id=YOUR-RESOURCE-ID&version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

Snyk returns a JSON document containing information about the selected resource. The information is the same as shown in [Understand the API response](view-cloud-resources.md#understand-the-api-response).
