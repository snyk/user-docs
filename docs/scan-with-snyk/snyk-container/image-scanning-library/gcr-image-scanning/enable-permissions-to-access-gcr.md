# Enable permissions to access GCR

**Prerequisites:**

Enable the [Cloud Resources Manager API](https://console.cloud.google.com/apis/library/cloudresourcemanager.googleapis.com?q=cloud%20resource%20manager\&id=16f5d23e-c895-4b9d-88e4-864c1766636f\&project=next-for-integration-testing) for the Google account you plan on integrating with Snyk.

From the relevant project in Google, ensure the you've created a service account for this Snyk integration.

**Steps:**

1. Go to the Google Cloud Platform Console [Credentials](https://console.cloud.google.com/apis/credentials) page, select the project that you want to integrate with, and then select to set up a new service account key.
2. From the view that loads, choose the service account from the dropdown list that you created for this integration, and configure a new key for that account with these values:
   * **Service account name** - assign a unique name for the account in order to remember its uses later on.
   * **Role** - Storage Object Viewer (roles/storage.objectViewer).
   * **Service account ID** - leave empty
   * **Key type** - JSON
3. Click **Create** The key is generated for your project.
4. Copy _the entire contents_ of the JSON file, which appears similar to the following:

![GCR\_key\_file\_contents.png](../../../../.gitbook/assets/uuid-c4e3b781-e575-5ab8-6cea-b0a8654068c4-en.png)

Save the data you copied in order to paste it when configuring the integration with Snyk.

**Next steps:**

Now, configure the integration: [Configure integration for GCR](https://support.snyk.io/hc/articles/360003916118#UUID-9e0df3f8-0780-b593-573b-5185bdca4a6d).
