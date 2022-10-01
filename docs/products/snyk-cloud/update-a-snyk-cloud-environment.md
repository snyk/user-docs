# Update a Snyk Cloud Environment

You can update the following attributes for a [Snyk Cloud Environment](snyk-cloud-concepts.md#environments):

* **AWS:** IAM role ARN (Amazon Resource Name). The new role ARN must have the same AWS account ID as the old role ARN. See [Find the role ARN](getting-started-with-snyk-cloud-aws/snyk-cloud-for-aws-api/step-3-create-and-scan-a-snyk-cloud-environment.md#find-the-role-arn).

For example, you would need to update the Snyk IAM role ARN if you change the role's name in the Terraform or CloudFormation template and deploy the changes.

To update an environment, take the following steps:

1. [Find the environment ID](update-a-snyk-cloud-environment.md#find-the-environment-id)
2. [Send an API request to update the environment](update-a-snyk-cloud-environment.md#update-the-environment)

## Find the environment ID

First, find the ID of the Snyk Cloud Environment you want to update. Send a request to the `/cloud/environments` endpoint in the below format:

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments?version=2022-04-13~experimental' \
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

## Update the environment

To update an environment, send a request to the `/cloud/environments/{environment_id}` endpoint in the below format:

```
curl -X PATCH \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments/YOUR-ENVIRONMENT-ID?version=2022-04-13~experimental' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H "Content-Type:application/vnd.api+json"  -d '{
  "data": {
    "attributes": {
      "options": {
        "role_arn": "YOUR-NEW-ROLE-ARN"
      }
    },
    "type": "resource"
  }
}'
```

Snyk returns a JSON document containing the updated environment details; for example:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "3b7ccff9-8900-4e54-0000-1234abcd1234",
    "type": "environment",
    "attributes": {
      "name": "Example AWS Environment",
      "options": {
        "role_arn": "arn:aws:iam::123412341234:role/snyk-cloud-role-updated"
      },
      "native_id": "123412341234",
      "properties": {
        "account_id": "123412341234"
      },
      "kind": "aws",
      "revision": 2,
      "created_at": "2022-07-31T00:50:49Z",
      "status": "success",
      "updated_at": "2022-08-17T18:18:01Z"
    },
    "relationships": {
      "organization": {
        "data": {
          "id": "d70c1768-5675-0000-1234-abcd1234abcd",
          "type": "organization"
        },
        "links": {
          "related": "/orgs/d70c1768-5675-0000-1234-abcd1234abcd?version=2022-04-13~experimental"
        }
      }
    }
  }
}
```

The `data.attributes.options.role_arn` field in the JSON output shows the new IAM role ARN.
