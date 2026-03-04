# Step 1: Download service account IaC template (API)

Before you can create a Cloud Environment, you must download an infrastructure as code (IaC) template declaring a tightly-scoped Google service account that gives Snyk permission to scan the configuration of resources in your Google Project.

The template also enables a set of [Google service APIs](https://cloud.google.com/service-usage/docs/enabled-service) for your Google Cloud Project. This ensures that Snyk can use the necessary APIs to scan your Project's resources.

You will use this IaC template to provision the role in [Step 2: Create the Google service account (API)](step-2-create-the-google-service-account-api.md).

## Retrieve the IaC template

To retrieve the IaC template from the [Snyk API](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#post-/orgs/-org_id-/cloud/permissions), you need the API token for a Snyk Organization-level [service account](../../../../../implementation-and-setup/enterprise-setup/service-accounts/) with an Org Admin role.

1. In the [Snyk Web UI](https://app.snyk.io/), navigate to **Settings** > **General** > **Organization ID** and copy your Organization ID.
2. Send a request to the Snyk API in the following format:

```
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/permissions?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
    "data": {
        "attributes": {
            "type": "tf",
            "platform": "google"
        },
        "type": "permissions"
    }
}'
```

{% hint style="info" %}
The preceding example is [curl](https://curl.se/), but you can use any API client, such as [Postman](https://www.postman.com/) or [HTTPie](https://httpie.io/).
{% endhint %}

The response is a JSON document like the one that follows (trimmed for length):

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "00000000-0000-0000-0000-000000000000",
    "type": "permissions",
    "attributes": {
      "data": "variable \"project_id\"<...>",
      "type": "tf"
    }
  }
}
```

## Unescape the JSON

The `data.attributes.data` field in the preceding output is an escaped JSON string containing the Terraform template with the Google service account.

Before you can use the template to provision the resources, you must unescape the JSON. This can be accomplished in either of the following ways:

* [Use `jq`](step-1-download-service-account-iac-template-api.md#use-jq)
* [Transform the content manually](step-1-download-service-account-iac-template-api.md#transform-the-content-manually)

### Use `jq`

1. Download and install [jq](https://stedolan.github.io/jq/download/).
2. When you are submitting the API request to retrieve the template, append the following to the end of the command:

```
| jq -r .data.attributes.data > snyk_google_iac_template.tf
```

This will place the properly-formatted template into the file `snyk_google_iac_template.tf` in your current working directory.

### Transform the content manually

1. Copy the contents of `data.attributes.data` from the API response, excluding the double quote at the very beginning and the very end of the value. You should end up with a long string starting with `variable \"project_id\"`.
2. Paste the string into a tool such as [FreeFormatter.com](https://www.freeformatter.com/json-escape.html) to unescape the JSON.
3. Save the unescaped Terraform output as a new `.tf` file.

## What's next?

The next step is to create the Google service account for Snyk.
