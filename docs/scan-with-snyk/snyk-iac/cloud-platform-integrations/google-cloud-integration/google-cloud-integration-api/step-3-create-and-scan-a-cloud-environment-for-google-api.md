# Step 3: Create and scan a Cloud Environment for Google (API)

{% hint style="info" %}
**Recap**\
You have created the Google service account for Snyk. Now you can create and scan a Cloud Environment.
{% endhint %}

To send a request to the [Snyk API](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#post-/orgs/-org_id-/cloud/environments) to create and scan a Cloud Environment, you must provide the Google service account's email address and your project ID in the API request body.

## Send the Snyk API request

Send a request to the Snyk API in the format below to create the Cloud Environment:

```
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
  "data": {
    "type": "environments",
    "attributes": {
      "options": {
        "identity_provider": "YOUR-IDENTITY-PPROVIDER-URL"
        "service_account_email": "YOUR-SERVICE-ACCOUNT-EMAIL",
        "project_id": "YOUR-PROJECT-ID"
      },
      "kind": "google"
    }
  }
}'
```

{% hint style="info" %}
The preceding example is [curl](https://curl.se/), but you can use any API client, such as [Postman](https://www.postman.com/) or [HTTPie](https://httpie.io/).
{% endhint %}

## Understand the API response

The response is a JSON document containing details about your newly created Cloud Environment, for example:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "c61079b9-7260-4b2f-1234-abcd1234abcd",
    "type": "environment",
    "attributes": {
      "name": "my-project-demo",
      "options": {
        "project_id": "my-project-demo",
        "service_account_email": "snyk-cloud-mt-us-abcd1234@my-project-demo.iam.gserviceaccount.com"
      },
      "native_id": "my-project-demo",
      "properties": {
        "project_id": "my-project-demo",
        "project_name": "my-project-demo",
        "project_number": "123456789012"
      },
      "kind": "google",
      "revision": 1,
      "created_at": "2022-10-13T20:45:19Z",
      "status": "in_progress",
      "updated_at": "2022-10-13T20:45:19Z"
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

To check to see if your scan is finished, see [Check to see if the scan is finished](../../aws-integration/aws-integration-api/step-3-create-and-scan-a-cloud-environment-api.md#check-to-see-if-the-scan-is-finished).

To re-scan an environment, see [Scan a Cloud environment](../../../getting-started-with-cloud-scans/manage-cloud-environments/scan-a-cloud-environment.md).

{% hint style="info" %}
It can take Google 60 seconds or more to create your service account. If you try to create an environment immediately after you create a service account and you receive a "could not validate credentials error", wait at least 60 seconds and try again.
{% endhint %}

## What's next?

You can now do the following:

* View the cloud configuration issues Snyk finds. See [Manage cloud issues](../../../getting-started-with-cloud-scans/manage-cloud-issues/).
* Prioritize your vulnerabilities with cloud context.
