# Step 1: Download IAM role IaC template

Before you can create a Snyk Cloud Environment, you must download an Infrastructure as Code (IaC) template declaring a read-only **Identity & Access Management** (IAM) role that Snyk can assume to scan the configuration of resources in your Amazon Web Services (AWS) account.&#x20;

You will use this IaC template to provision the role in [Step 2: Create the Snyk IAM role](step-2-create-the-snyk-iam-role.md).

You can choose the template format: either [Terraform HCL](https://www.terraform.io/language/syntax/configuration) or [AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html). The IAM permissions are identical in both, so pick the format you are most comfortable working with.

## Retrieve the IaC template

To retrieve the IaC template from the Snyk API, you need the API token for an Organization-level [service account](https://docs.snyk.io/features/user-and-group-management/structure-account-for-high-application-performance/service-accounts#set-up-a-service-account) with an Org Admin role.

1. In the [Snyk Web UI](https://app.snyk.io), navigate to **Settings (cog icon) > General > Organization ID** and copy your Organization ID.
2. Send a request to the Snyk API in the below format, replacing `INPUT-TYPE` with `tf` for Terraform or `cf` for CloudFormation:

```
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/permissions?version=2022-04-13~experimental' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
    "data": {
        "attributes": {
            "type": "INPUT-TYPE",
            "platform": "aws"
        },
        "type": "permissions"
    }
}'
```

{% hint style="info" %}
The example above uses [curl](https://curl.se/), but you can use any API client, such as [Postman](https://www.postman.com/) or [HTTPie](https://httpie.io/).
{% endhint %}

## Understand the API response

The response is a JSON document like the ones below (trimmed for length).

Example response with Terraform configuration:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "00000000-0000-0000-0000-000000000000",
    "type": "permissions",
    "attributes": {
      "data": "data \"aws_iam_policy_document\"<...>",
      "type": "tf"
    }
  }
}
```

Example response with CloudFormation template:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "00000000-0000-0000-0000-000000000000",
    "type": "permissions",
    "attributes": {
      "data": "AWSTemplateFormatVersion:<...>",
      "type": "cf"
    }
  }
}
```

## Unescape the JSON

The `data.attributes.data` field in the output above is an escaped JSON string containing the Terraform or CloudFormation template with the IAM role and policy.

Before you can use the template to provision the resources, you need to **unescape** the JSON:

1. Copy the contents of `data.attributes.data`, excluding the double quote at the very beginning and the very end of the value. You should end up with a long string starting with `data \"aws_iam_policy_document\"` (Terraform) or `AWSTemplateFormatVersion` (CloudFormation).
2. Paste the string into a tool such as [FreeFormatter.com](https://www.freeformatter.com/json-escape.html) to unescape the JSON.
3. Save the unescaped CloudFormation output as a new `.tf` file (Terraform) or `.yaml` file (CloudFormation).

## What's next?

The next step is to create the IAM role and policy for Snyk Cloud using the template you downloaded.
