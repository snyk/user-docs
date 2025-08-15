# Step 1: Download IAM role IaC template (API)

Before you can create a Cloud Environment, you must download an Infrastructure as Code (IaC) template declaring a read-only Identity and Access Management (IAM) role that Snyk can assume to scan the configuration of resources in your Amazon Web Services (AWS) account.

You will use this IaC template to provision the role in [Step 2: Create the Snyk IAM role](step-2-create-the-snyk-iam-role-api.md).

You can choose the template format, either [Terraform HCL](https://www.terraform.io/language/syntax/configuration) or [AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html). The IAM permissions are identical in both, so pick the format you are most comfortable working with.

## Retrieve the IaC template

To retrieve the IaC template using  the [Snyk API endpoint Generate Cloud Provider Permissions](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#post-/orgs/-org_id-/cloud/permissions), you need the API token for an Organization-level [service account](../../../../../implementation-and-setup/enterprise-setup/service-accounts/) with an Org Admin role.

1. In the [Snyk Web UI](https://app.snyk.io), navigate to **Settings** > **General** > **Organization ID** and copy your Organization ID.
2. Send a request to the Snyk API in the following format, replacing `INPUT-TYPE` with `tf` for Terraform or `cf` for CloudFormation:

```
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/permissions?version=2022-12-21~beta' \
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
The preceding example is [curl](https://curl.se/), but you can use any API client, such as [Postman](https://www.postman.com/) or [HTTPie](https://httpie.io/).
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

The `data.attributes.data` field in the preceding output is an escaped JSON string containing the Terraform or CloudFormation template with the IAM role and policy.

Before you can use the template to provision the resources, you need to unescape the JSON. This can be accomplished in the following ways:

* [Use jq](step-1-download-iam-role-iac-template-api.md#use-jq)
* [Transform the content manually](step-1-download-iam-role-iac-template-api.md#transform-the-content-manually)

### Use `jq`

1. Download and install [jq](https://stedolan.github.io/jq/download/).
2.  When you are submitting the API request to retrieve the template, append the following to the end of the command:

    ```
    | jq -r .data.attributes.data > snyk_iac_template
    ```

    This will place the properly-formatted template into the file `snyk_iac_template` in your current working directory.
3. Rename the file with a `.tf` extension (Terraform) or `.yaml` (CloudFormation).

### Transform the content manually

1. Copy the contents of `data.attributes.data` from the API response, excluding the double quote at the very beginning and the very end of the value. You should end up with a long string starting with `data \"aws_iam_policy_document\"` (Terraform) or `AWSTemplateFormatVersion` (CloudFormation).
2. Paste the string into a tool such as [FreeFormatter.com](https://www.freeformatter.com/json-escape.html) to unescape the JSON.
3. Save the unescaped output as a new `.tf` file (Terraform) or `.yaml` file (CloudFormation).

## Optional: Change IAM role name

By default, the name of the Snyk IAM role is `snyk-cloud-role`. If your Organization has specific role naming requirements, you have the option to change this name in the Terraform or CloudFormation template.

In **Terraform**, the role name is on line 19:

```
  name                = "snyk-cloud-role"
```

In **CloudFormation**, the role name is on line 7:

```
      RoleName: snyk-cloud-role
```

## What's next?

The next step is to [create the IAM role](step-2-create-the-snyk-iam-role-api.md) and policy for Snyk using the template you downloaded.
