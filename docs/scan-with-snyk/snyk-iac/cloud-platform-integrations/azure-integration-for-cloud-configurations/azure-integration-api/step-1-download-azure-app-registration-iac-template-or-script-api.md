# Step 1: Download Azure app registration IaC template or script (API)

Before you can create a Cloud Environment for an Azure subscription, you must **download** a Terraform infrastructure as code (IaC) template or Azure CLI Bash script declaring the following resources:

* [An Active Directory (AD) application registration](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#application-registration)
* [A federated identity credential](https://learn.microsoft.com/en-us/azure/active-directory/develop/workload-identity-federation)
* [A service principal](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#service-principal-object)

This infrastructure gives Snyk read-only permission to scan the configuration of resources in your subscription.

You will use the IaC template or Bash script you downloaded to **provision** the infrastructure in [Step 2: Create the Entra ID app registration (API)](step-2-create-the-entra-id-app-registration-api.md).

Both methods create the same infrastructure, so pick the method you are most comfortable working with.

## Retrieve the IaC template or script

To retrieve the IaC template from the Snyk API, you need the API token for a Snyk Organization-level [service account](../../../../../implementation-and-setup/enterprise-setup/service-accounts/) with an Org Admin role.

You also need the subscription and tenant IDs of the Azure subscription you are onboarding. You can find them using the method [described in the Azure documentation](https://learn.microsoft.com/en-us/azure/azure-portal/get-subscription-tenant-id).

1. In the [Snyk Web UI](https://app.snyk.io/), navigate to **Settings** > **General** > **Organization ID** and copy your Organization ID.
2. Send a request to the Snyk API in the following format, replacing `INPUT-TYPE` with `tf` for Terraform or `bash` for Bash:

```
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/permissions?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
    "data": {
        "attributes": {
            "options": {
              "subscription_id": "YOUR-SUBSCRIPTION-ID",
              "tenant_id": "YOUR-TENANT-ID"
            },
            "type": "INPUT-TYPE",
            "platform": "azure"
        },
        "type": "permissions"
    }
}'
```

{% hint style="info" %}
The preceding example is [curl](https://curl.se/), but you can use any API client, such as [Postman](https://www.postman.com/) or [HTTPie](https://httpie.io/).
{% endhint %}

If you plan to use the [Azure Cloud Shell](https://portal.azure.com/#cloudshell/) to execute the Bash script instead of running the Azure CLI locally, execute the curl command above in the Cloud Shell.

## Understand the API response

The response is a JSON document like the ones that follow (trimmed for length).

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
      "data": "provider \"azuread\"<...>",
      "type": "tf"
    }
  }
}
```

Example response with Bash script:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "00000000-0000-0000-0000-000000000000",
    "type": "permissions",
    "attributes": {
      "data": "objectId=$(az ad app create <...>",
      "type": "bash"
    }
  }
}
```

## Unescape the JSON

The `data.attributes.data` field in the preceding output is an escaped JSON string containing the Terraform template or Bash script with the Entra ID app registration, federated identity credential, and service principal.

Before you can use the template to provision the resources, you need to unescape the JSON. This can be accomplished in the following ways:

* [Use jq](step-1-download-azure-app-registration-iac-template-or-script-api.md#use-jq)
* [Transform the content manually](step-1-download-azure-app-registration-iac-template-or-script-api.md#transform-the-content-manually)

### Use `jq`

1. Download and install [jq](https://stedolan.github.io/jq/download/).
2. When you are submitting the API request during template retrieval, append the following to the end of the command:\
   `| jq -r .data.attributes.data > snyk_azure_permissions`\
   This will place the properly-formatted template into the file `snyk_azure_permissions` in your current working directory.
3. Rename the file with a `.tf` (Terraform) or `.sh` (Bash) extension.

### Transform the content manually

1. Copy the contents of `data.attributes.data` from the API response, excluding the double quote at the very beginning and the very end of the value. You should end up with a long string starting with `provider \"azuread\"` (Terraform) or `objectId=$(az ad app create` (Bash).
2. Paste the string into a tool such as [FreeFormatter.com](https://www.freeformatter.com/json-escape.html) to unescape the JSON.
3. Save the unescaped output as a new `.tf` file (Terraform) or `.sh` file (Bash).

## What's next?

The next step is to [create the Entra ID app registration, federated identity credential, and service principal for Snyk](step-2-create-the-entra-id-app-registration-api.md) using the template or script you downloaded.
