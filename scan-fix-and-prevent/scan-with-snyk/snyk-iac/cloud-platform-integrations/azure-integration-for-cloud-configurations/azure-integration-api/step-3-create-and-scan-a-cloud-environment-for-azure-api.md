# Step 3: Create and scan a Cloud Environment for Azure (API)

{% hint style="info" %}
**Recap**\
You have created the Azure app registration, federated identity credential, and service principal for Snyk. Now you can create and scan a Cloud Environment.
{% endhint %}

To send a request to the Snyk API to create and scan an Azure Cloud Environment, you must provide the subscription ID, tenant ID, and application ID in the API request body.

## Send the Snyk API request

Send a request to the Snyk API in the format that follows to create the Cloud Environment:

```
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
  "data": {
    "type": "environments",
    "attributes": {
      "options": {
        "subscription_id": "YOUR-SUBSCRIPTION-ID",
        "tenant_id": "YOUR-TENANT-ID",
        "application_id": "YOUR-APPLICATION-ID"
      },
      "kind": "azure",
      "name": "OPTIONAL-NAME"
    }
  }
}'
```

{% hint style="info" %}
The preceding example is [curl](https://curl.se/), but you can use any API client, such as [Postman](https://www.postman.com/) or [HTTPie](https://httpie.io/).
{% endhint %}

## Understand the API response

The response is a JSON document containing details about your newly created Cloud Environment. For example:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "e25a5ef1-1e96-1234-0000-1234abcd1234",
    "type": "environment",
    "attributes": {
      "name": "Example Azure Environment",
      "options": {
        "tenant_id": "00000000-0000-0000-1234-12345678abcd",
        "application_id": "12345678-1234-0000-0000-09876543dcba",
        "subscription_id": "abcd1234-abcd-1234-0000-000000000000"
      },
      "native_id": "abcd1234-abcd-1234-0000-000000000000",
      "properties": {
        "subscription_id": "abcd1234-abcd-1234-0000-000000000000",
        "subscription_name": "Example User"
      },
      "kind": "azure",
      "revision": 1,
      "created_at": "2023-02-06T06:34:05Z",
      "status": "in_progress",
      "updated_at": "2023-02-06T06:34:05Z"
    },
    "relationships": {
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

Snyk automatically triggers a scan when an environment is created.

{% hint style="info" %}
The `data.attributes.status` field in the JSON output is set to `in_progress.` This means that Snyk has created your environment and has started scanning it.
{% endhint %}

To check if your scan is finished, see [Check to see if the scan is finished](../../aws-integration/aws-integration-api/step-3-create-and-scan-a-cloud-environment-api.md#check-to-see-if-the-scan-is-finished).

To manually re-scan an environment, see [Scan a Cloud Environment](../../../getting-started-with-cloud-scans/manage-cloud-environments/scan-a-cloud-environment.md).

## What's next?

You can now do the following:

* View the cloud configuration issues Snyk finds. See [Manage cloud issues](../../../getting-started-with-cloud-scans/manage-cloud-issues/).
* Prioritize your vulnerabilities with cloud context.
