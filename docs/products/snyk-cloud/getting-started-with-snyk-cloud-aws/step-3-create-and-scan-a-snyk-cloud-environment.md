# Step 3: Create and scan a Snyk Cloud Environment

{% hint style="info" %}
**Recap**\
****You have created the Snyk Cloud IAM role. Now you can create and scan a Snyk Cloud Environment.
{% endhint %}

To send a request to the Snyk API to create and scan a Snyk Cloud Environment, you must provide the role’s Amazon Resource Name (ARN) in the API request body.

## Find the role ARN

The role ARN should follow this format, unless you changed the name of the role:

```
arn:aws:iam::YOUR-ACCOUNT-ID:role/snyk-cloud-role
```

If you do not know your Amazon Web Services (AWS) account ID, or if you changed the name of the IAM role in the Terraform or CloudFormation template, you can find the role ARN using the [AWS CLI](step-3-create-and-scan-a-snyk-cloud-environment.md#aws-cli) or the [AWS Management Console](step-3-create-and-scan-a-snyk-cloud-environment.md#aws-management-console).

### AWS CLI

To find the ARN of the Snyk Cloud IAM role using the AWS CLI:

```
aws iam get-role \
  --role-name snyk-cloud-role \
  --query 'Role.Arn' --output text
```

You'll see output like this:

```
arn:aws:iam::123412341234:role/snyk-cloud-role
```

### AWS Management Console

1. Log in to the [AWS Management Console](https://console.aws.amazon.com).
2. Navigate to [Identity and Access Management](https://console.aws.amazon.com/iamv2/home#/home).
3. In the left sidebar, select **Roles**.
4. On the **Roles** page, search for `snyk-cloud-role` (or substitute the name of your role, if you changed it):

![Search for the name of your role in the AWS Management Console](../../../.gitbook/assets/snyk-cloud-console-find-arn.png)

5\. Select the role.

6\. On the role details page, in the **Summary** section, find and copy the ARN:

![Copy the role ARN in the AWS Management Console](../../../.gitbook/assets/snyk-cloud-console-copy-arn.png)

## Send the Snyk API request

After you have the role ARN, send a request to the Snyk API in the format below to create the Snyk Cloud Environment:

```
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments?version=2022-04-13~experimental' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
  "data": {
    "attributes": {
      "kind": "aws",
      "name": "Example AWS Environment",
      "options": {
        "role_arn": "YOUR-ROLE-ARN"
      }
    },
    "type": "environment"
  }
}'
```

{% hint style="info" %}
The example above uses [curl](https://curl.se/), but you can use any API client, such as [Postman](https://www.postman.com/) or [HTTPie](https://httpie.io/).
{% endhint %}

## Understand the API response

The response is a JSON document containing details about your newly created Snyk Cloud Environment; for example:

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
        "role_arn": "arn:aws:iam::123412341234:role/snyk-cloud-role"
      },
      "native_id": "123412341234",
      "properties": {
        "account_id": "123412341234"
      },
      "kind": "aws",
      "revision": 1,
      "created_at": "2022-07-31T00:50:49Z",
      "status": "in_progress",
      "updated_at": "2022-07-31T00:50:49Z"
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

Snyk automatically triggers a scan when an environment is created.

Note: the `data.attributes.status` field in the JSON output is set to `in_progress`. This means that Snyk has created your environment and has started scanning it.

## Check if the scan is finished

Optionally, see if the scan is finished by sending another API request in the format below to get environment details. You can find the environment ID in the `data.id` field of the JSON output when you created the environment.

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments?id=YOUR-ENVIRONMENT-ID&version=2022-04-13~experimental' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

If the `data.attributes.status` field in the JSON output is set to `success`, Snyk has finished scanning your environment.

To re-scan an environment, see [Scan a Snyk Cloud Environment](../scan-a-snyk-cloud-environment.md).

## What's next?

You can now view misconfiguration issues in the API or Snyk Web UI. See [Snyk cloud issues](../snyk-cloud-issues/) for more information.
