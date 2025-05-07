# Google Security Command Center

{% hint style="info" %}
**Release status**

The Google Cloud Security Command Center integration is in [Early Access](https://docs.snyk.io/getting-started/snyk-release-process#early-access), and is available to all customers on a Snyk Enterprise plan.&#x20;
{% endhint %}

The Google Cloud Security Command Center (SCC) integration sends Snyk issues to SCC, enabling you to view and manage Snyk issues alongside cloud security findings from Google Cloud in a single pane of glass. Snyk issues are represented in SCC as code security findings. When Snyk issues are updated, corresponding SCC findings are automatically updated as well. All details are available at the Organization level in the Google Cloud Security Command Center (SCC) integration.

Use the following instructions to set up the Google Security Command Center integration:

* Access the Google Cloud SCC Console using the Snyk for SCC marketplace listing to create a Finding Source
* Set up the integration using either the Snyk Web UI (recommended) or the REST API.

{% hint style="info" %}
You must use MP to create project-level service accounts that enable organization-wide management of permissions and security findings. This process involves the following steps:

* Third-party security providers must be activated at the organizational level.
* After using MP, the Google SCC API must be enabled separately.
* Copy IAM permissions from the project to the organization level in the Google Cloud Settings.

Service accounts created with MP hide SCC-related IAM roles to prevent unauthorized modifications.
{% endhint %}

## Prerequisites

* **Snyk:** A Snyk user account with [permissions](https://docs.snyk.io/admin/user-roles/user-role-management) to edit and view Group integrations
* **Google** **Cloud:** A Google Cloud organization with Security Command Center enabled. See the Google Cloud [Activate Security Command Center](https://cloud.google.com/security-command-center/docs/activate-scc-for-an-organization) page for more details on how to enable it.
* **Google SCC API**: you must enable the SCC API.&#x20;

## Create the Finding Source using the Google Cloud SCC Console&#x20;

* In the SCC console, navigate to **Marketplace** and search for **Snyk**. Alternatively, navigate directly to the [Snyk for SCC marketplace listing](https://console.cloud.google.com/marketplace/product/snyk-marketplace/snyk-google-scc).
* Click **SIGN UP WITH PARTNER** to install the Snyk for SCC integration. During this process, you will create a **Findings Source** for Snyk and a **Service Account** with [Security Center Findings Editor](https://cloud.google.com/security-command-center/docs/access-control-org#securitycenter.findingsEditor) permissions.
* Navigate to Google Cloud IAM and locate the **Service Accoun**t you created in the previous step, then [create a service account key](https://cloud.google.com/iam/docs/keys-create-delete#creating) in JSON format.
* Make a note of the **Source ID** (Findings Source name) and the **Service Account Ke**y, as you will need to provide them to the Snyk Web UI.

You can then set up the integration in Snyk using the Snyk Web UI or REST API.

## Set up the integration using the Snyk Web UI

### **Required parameters**

* **Source ID** (Finding Source Name) - Identifies the name of the organization source. The Google organization ID is parsed automatically from this field.
* **JSON Service Account Key File** - Authenticates with Google Cloud.

### **Integration Hub setup**

* Open the Integration Hub menu.
* Select the Issue Forwarding tag and search for Google SCC.
* Click the Add button.
* Add the Profile name for this integration.
* Add the Org ID for the Google Cloud project that holds the Kubernetes cluster.
* Add the JSON Service Account Key File.
* Add the Source ID (Findings Source Name).
* Click the Done button.
* When the connection is established, the Google SCC integration status changes to Connected.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXfwL5JQFYq5FIQRMZxTHZYswjUr3cj-9Yb7W-fnAU25J8T8-u3jGH81ZUSTtMyUtLchzHMCCshP_5_5B3q3drJ_6sZlM1gdttTw8nDtVVuyn4Y-LDA2oqtffWCswFPqy0RyfSzQew?key=9J3dV91yICcvUevnOqzNv8VZ" alt=""><figcaption><p>Google SCC - Setup screen</p></figcaption></figure>

## Set up the integration using the Snyk REST API

{% hint style="warning" %}
Setting up the integration using the API has two limitations:

1. Findings in Google SCC will not include a URL back to the Snyk issue. External URI support is available only when the SCC source is created from the Snyk marketplace listing in the SCC console.
2. The integration will not be visible in the Snyk group-level integrations UI.
{% endhint %}

Use the following API request template:

{% code overflow="wrap" %}
```
curl --location 'https://api.snyk.io/rest/groups/<SNYK_GROUP_ID>/cloud_events/group_registrations?version=2023-01-25~experimental' \
--header 'Content-Type: application/vnd.api+json' \
--header 'Authorization: token <SNYK_API_TOKEN>' \
--data '{
	"data": {
	"type": "group_registration",
	"attributes": {
			"type": "google-securitycommandcenter",
			"name": "Test Google SCC integration",
			"config": {
				"org_id": "<GCP_ORG_ID>",
				"finding_source_name": "<GCP_FINDING_SOURCE_NAME>"
			},
			"credentials": "<GCP_SERVICE_ACCOUNT_KEY>"
		}
	}
}'
```
{% endcode %}

Replace the following placeholders:

* `<SNYK_GROUP_ID>` - the [group ID](https://docs.snyk.io/admin/groups-and-organizations/groups/group-general-settings) for your Snyk Group
* `<SNYK_API_TOKEN>` - the Snyk API token (see Prerequisites above)
* `<GCP_ORG_ID>` - your Google Cloud organization ID. See the Google Cloud [Getting your organization resource ID](https://cloud.google.com/resource-manager/docs/creating-managing-organization#retrieving_your_organization_id) page for more details.
* `<GCP_FINDING_SOURCE_NAME>` - the full relative resource name / Source ID for the finding source created in the previous step
* `<GCP_SERVICE_ACCOUNT_KEY>` - the full JSON service account key (see Prerequisites above).

{% hint style="info" %}
The JSON service account key must be string encoded, for example, double quotes must be escaped and new lines removed. One way to do this would be to use the javascript function [JSON.stringify()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/stringify) in a local javascript environment.
{% endhint %}

### Fetch the integration

To fetch an existing SCC integration, make the following API request. Replace the placeholders with actual values.&#x20;

{% code overflow="wrap" %}
```
curl --location 'https://api.snyk.io/rest/groups/<SNYK_GROUP_ID>/cloud_events/group_registrations?version=2023-01-25~experimental' \

--header 'Authorization: token <SNYK_API_TOKEN>' \

--header 'Content-Type: application/vnd.api+json'
```
{% endcode %}

This will return a list of integration registrations, including the SCC integration created earlier.

### Delete the integration

You can create an API request to delete the integration. Use the fetch the integration  API call to  obtain the integration ID, then make the following call, replacing `<REGISTRATION_ID>` with the ID returned from the fetch request:

{% code overflow="wrap" %}
```
curl --location --request DELETE 'https://api.snyk.io/rest/groups/<SNYK_GROUP_ID>/cloud_events/group_registrations/<REGISTRATION_ID>?version=2023-01-25~experimental' \

--header 'Content-Type: application/vnd.api+json' \

--header 'Authorization: token <SNYK_API_TOKEN>'
```
{% endcode %}
