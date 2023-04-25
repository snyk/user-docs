# Enable permissions to access GAR

## **Prerequisites**

[Enable the Artifact Registry API](https://cloud.google.com/artifact-registry/docs/enable-service) for the Google account you plan on integrating with Snyk. Allow a few minutes for Google to propagate the enablement.

## **Steps**

1. Navigate to the Google Cloud Console [Credentials](https://console.cloud.google.com/apis/credentials) page.&#x20;
2. If it’s not already selected, select the Google project for which you’re creating credentials.
3. Select the **Create Credentials** button and select **Service Account**.
4. Give the new service account a unique name and ID and select **Create**.
5. In the Service account permissions page, choose **Select a role** and choose **Artifact Registry Reader**. You must also add an additional role that has the **resourcemanager.projects.list** permission, such as **Browser** or **Viewer**.
6. Once the account is created, select the relevant account in the **Service Accounts** section.
7. In the Service account details page, under the **Keys** section, select **Add Key** and **Create New Key**.
8. Select JSON for the **Key type** and select **Create**.

The JSON key is generated for your project. Copy the entire contents of the JSON file for the next step: [Configure integration for GAR](configure-integration-for-gar.md).
