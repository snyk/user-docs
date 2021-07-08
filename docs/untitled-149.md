# Enable permissions to access GCR

* [ Container security with GCR](/hc/en-us/articles/360003916098-Container-security-with-GCR)
* [ Enable permissions to access GCR](/hc/en-us/articles/360004191777-Enable-permissions-to-access-GCR)
* [ Configure integration for GCR](/hc/en-us/articles/360003916118-Configure-integration-for-GCR)

##  Enable permissions to access GCR

**Prerequisites:**

Enable the [Cloud Resources Manager API](https://console.cloud.google.com/apis/library/cloudresourcemanager.googleapis.com?q=cloud%20resource%20manager&amp;id=16f5d23e-c895-4b9d-88e4-864c1766636f&amp;project=next-for-integration-testing) for the Google account you plan on integrating with Snyk.

From the relevant project in Google, ensure the you've created a service account for this Snyk integration.

**Steps:**

1. Go to the Google Cloud Platform Console [Credentials](https://console.cloud.google.com/apis/credentials) page, select the project that you want to integrate with, and then select to set up a new service account key.
2. From the view that loads, choose the service account from the dropdown list that you created for this integration, and configure a new key for that account with these values:
   * Service account name - assign a unique name for the account in order to remember its uses later on.
   * Role - Storage Object Viewer \(roles/storage.objectViewer\).
   * Service account ID - leave empty
   * Key type - JSON
3. Click Create.

   The key is generated for your project.

4. Copy _the entire contents_ of the JSON file, which appears similar to the following:

   Save the data you copied in order to paste it when configuring the integration with Snyk.

**Next steps:**

Now, configure the integration: [Configure integration for GCR](/hc/articles/360003916118#UUID-9e0df3f8-0780-b593-573b-5185bdca4a6d).

