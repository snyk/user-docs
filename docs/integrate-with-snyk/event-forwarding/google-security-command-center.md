# Google Security Command Center

{% hint style="info" %}
**Release status**

The Google Cloud Security Command Center integration is in [Early Access](https://docs.snyk.io/getting-started/snyk-release-process#early-access), and is available only with Snyk Enterprise plans. For more information, see [Plans and pricing](https://snyk.io/plans/).
{% endhint %}

The Google Cloud Security Command Center (SCC) integration sends Snyk issues to SCC, enabling you to view and manage Snyk issues alongside cloud security findings from Google Cloud in a single viewpoint. Snyk issues are represented in SCC as code security findings. When Snyk issues are updated, corresponding SCC findings are automatically updated as well. All details are available at the Organization level in the Google Cloud Security Command Center (SCC) integration.

Use the following instructions to set up the integration:

* Navigate to **Marketplace** in the SCC Console and search for Snyk.
* Follow the guided path to create a Findings Source and Service Account.

{% hint style="info" %}
Service Accounts are not available at the Organization level in Google Cloud IAM. You must create a Service Account in a specific Project. Snyk recommends creating the Service Account in a Project that is already being monitored by SCC. You must [enable the SCC API](https://console.cloud.google.com/apis/library) in the same Project as the Service Account.
{% endhint %}

* Copy the Findings Source ID and Service Account Token (JSON) into the Snyk Web UI

## Prerequisites

* **Snyk:** A Snyk user account with [permissions](https://docs.snyk.io/admin/user-roles/user-role-management) to edit and view Group integrations
* **Google** **Cloud:** A Google Cloud organization with Security Command Center enabled. See the Google Cloud [Activate Security Command Center](https://cloud.google.com/security-command-center/docs/activate-scc-for-an-organization) page for more details on how to enable it.
* **Google SCC API**: You must [enable the SCC API](https://console.cloud.google.com/apis/library) in the same Project as the Service Account

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
