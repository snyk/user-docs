# View cloud issues in the API

You can view cloud issues for an Organization through the Snyk API. Also see [View cloud issues in the Snyk Web UI](view-cloud-issues-in-the-snyk-web-ui.md).

### View all cloud issues

To view a list of all the cloud issues in your organization through the Snyk API, send a request to the `/cloud/issues` endpoint in the below format:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/issues?version=2022-04-13~experimental' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

Snyk returns a JSON document containing information about all cloud issues in the organization. For example:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "next": "/rest/orgs/d70c1768-5675-4f44-0000-1234abcd1234/cloud/issues?starting_after=eyJpZCI6IjY5ODA5MjNhLWU0ZTAtNDg3Mi04ZDAwLWRjZDEXAMPLEEXAMPLE&version=2022-04-13~experimental"
  },
  "data": [
    {
      "id": "01869050-4009-4ca0-0000-abcd1234abcd",
      "type": "issue",
      "attributes": {
        "created_at": "2022-08-04T22:28:30Z",
        "kind": "rule",
        "severity": "high",
        "title": "S3 bucket does not have `ignore_public_acls` enabled",
        "description": null,
        "resolved_as": null,
        "resolved_at": null,
        "resolved_by": null,
        "resolved_comment": null,
        "revision": 1,
        "status": "open",
        "updated_at": "2022-08-04T22:28:30Z"
      },
      "relationships": {
        "environment": {
          "data": {
            "id": "ef5d85de-fb4f-4c42-1234-000000000000",
            "type": "environment",
            "attributes": {
              "native_id": "123456789012",
              "name": "Demo AWS Environment",
              "kind": "aws"
            }
          },
          "links": {
            "related": "/orgs/d70c1768-5675-4f44-0000-1234abcd1234/environments/ef5d85de-fb4f-4c42-1234-000000000000?version=2022-04-13~experimental"
          }
        },
        "organization": {
          "data": {
            "id": "d70c1768-5675-4f44-0000-1234abcd1234",
            "type": "organization"
          },
          "links": {
            "related": "/orgs/d70c1768-5675-4f44-0000-1234abcd1234?version=2022-04-13~experimental"
          }
        },
        "resource": {
          "data": {
            "id": "543599a0-665a-5387-9876-543210abcdef",
            "type": "resource",
            "attributes": {
              "name": "fugue-123456789012-us-east-1",
              "input_type": "tf_runtime",
              "id": "543599a0-665a-5387-9876-543210abcdef",
              "type": "aws_s3_bucket",
              "kind": "cloud",
              "native_id": "arn:aws:s3:::fugue-123456789012-us-east-1",
              "location": "us-east-1",
              "platform": "aws",
              "revision": 1,
              "tags": {}
            }
          },
          "links": {
            "related": "/orgs/d70c1768-5675-4f44-0000-1234abcd1234/resources/543599a0-665a-5387-9876-543210abcdef?version=2022-04-13~experimental"
          }
        },
        "rule_result": {
          "data": {
            "id": "6d4679a0-8d72-4ebc-8888-123456ghijkl",
            "type": "rule_result",
            "attributes": {
              "message": "S3 buckets should have the `ignore_public_acls` enabled",
              "rule_id": "SNYK-CC-TF-97",
              "category": "data"
            }
          }
        }
      }
    },
    <trimmed for length>
  ]
}
```

Below are some key attributes from the API response:

| Attribute                                           | Definition          |
| --------------------------------------------------- | ------------------- |
| `data.id`                                           | Issue ID            |
| `data.attributes.severity`                          | Rule severity       |
| `data.attributes.title`                             | Rule title          |
| `data.attributes.status`                            | Issue status        |
| `data.relationships.environment.id`                 | Environment ID      |
| `data.relationships.resource.id`                    | Resource ID         |
| `data.relationships.resource.attributes`            | Resource attributes |
| `data.relationships.rule_result.attributes.rule_id` | Rule ID             |

### Filter issues

Query parameters allow you to filter the list of issues.

For example, to return only `open` issues, add `status=open` to the URL:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/issues?status=open&version=2022-04-13~experimental' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

You can filter issues by resource tags. Resource tags are key-value pairs that must be URL-encoded. Use a tool such as [urlencoder.org](https://www.urlencoder.org/) to encode the key-value pair as a JSON object.&#x20;

To return only issues for resources with the tag `{"stage":"prod"}`, add `resource_tags=%7B%22stage%22%3A%22prod%22%7D` to the URL:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/issues?resource_tags=%7B%22stage%22%3A%22prod%22%7D&version=2022-04-13~experimental' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

Some parameters allow you to specify multiple values. To return issues associated with Amazon Web Services (AWS) S3 buckets or CloudTrail trails, add `resource_type=aws_s3_bucket,aws_cloudtrail` to the URL:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/issues?resource_type=aws_s3_bucket,aws_cloudtrail&version=2022-04-13~experimental' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

You can combine query parameters by using an `&` symbol. To return only critical-severity issues in the environment for AWS account ID 123456789012, add `severity=critical&environment_native_id=123456789012` to the URL:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/issues?severity=critical&environment_native_id=123456789012&version=2022-04-13~experimental' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

### View details for a single cloud issue

To view a single cloud issue through the Snyk API, send a request in the below format. You can get the issue ID using the method shown in [View all cloud issues](view-cloud-issues-in-the-api.md#view-all-cloud-issues).

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/issues/YOUR-ISSUE-ID?version=2022-04-13~experimental' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

Snyk returns a JSON document containing information about the selected issue. The information is the same as shown in [View all cloud issues](view-cloud-issues-in-the-api.md#view-all-cloud-issues).
