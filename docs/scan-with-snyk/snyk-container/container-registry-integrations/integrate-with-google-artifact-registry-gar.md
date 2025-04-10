# Integrate with Google Artifact Registry (GAR)

Snyk integrates with [Google Artifact Registry (GAR)](https://cloud.google.com/artifact-registry) so you can monitor your containers for vulnerabilities and fix them as you work. Snyk tests the container images you have imported on a regular cadence.

## Enable permissions for GAR integration

### **Prerequisites to enabling permissions for GAR integration**

[Enable the Artifact Registry API](https://cloud.google.com/artifact-registry/docs/enable-service) for the Google account you plan on integrating with Snyk. Allow a few minutes for Google to propagate the enablement.

### **Steps to enable permissions for GAR integration**

1. Navigate to the Google Cloud Console [Credentials](https://console.cloud.google.com/apis/credentials) page.&#x20;
2. Select the Google project for which you are creating credentials if it is not already selected.
3. Select the **Create Credentials** button and select **Service Account**.
4. Give the new service account a unique name and ID and select **Create**.
5. On the Service account permissions page, choose **Select a role** and **Artifact Registry Reader**. You must also add an additional role that has the resourcemanager.projects.list permission, such as **Browser** or **Viewer**.
6. After the account has been created, select the relevant account in the **Service Accounts** section.
7. In the Service account details page, under the **Keys** section, select **Add Key** and **Create New Key**.
8. Select JSON for the **Key type** and select **Create**.

The JSON key is generated for your Project. Copy the entire contents of the JSON file for the next step: [Configure integration for GAR](integrate-with-google-artifact-registry-gar.md#configure-integration-for-gar).

## Configure integration for GAR

Configure integration from Snyk with your Google Artifact Registry account to scan for vulnerabilities and fix security and license issues accordingly.

### Prerequisites to configuring integration for GAR

[Enable permissions](integrate-with-google-artifact-registry-gar.md#enable-permissions-for-gar-integration) for GAR integration

### Steps to configure GAR integration

1. Navigate to your Organization in the Snyk Web UI.
2. Select **Integrations**.
3. In the Container Registries section, select **Google Artifact Registry**.
4. In the Account credentials section, enter your [Artifact Registry](https://cloud.google.com/artifact-registry/docs/repositories/repo-locations) hostname. This is typically \<Your Region Name>-docker.pkg.dev, but in some cases, you may need to use a specific region or multi-region, for example, "us-east1-docker.pkg.dev" or "us-docker.pkg.dev".
5. In the JSON key file field, paste the entire contents of the JSON key file you downloaded when [enabling permissions](integrate-with-google-artifact-registry-gar.md#enable-permissions-for-gar-integration).
6. Select **Save**.

Snyk checks the credentials, and upon success, the page reloads with a notification that the connection succeeded.
