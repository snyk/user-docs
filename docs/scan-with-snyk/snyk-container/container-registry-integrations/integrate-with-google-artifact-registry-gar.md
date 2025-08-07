# Integrate with Google Artifact Registry (GAR)

Snyk integrates with [Google Artifact Registry (GAR)](https://cloud.google.com/artifact-registry) so you can monitor your containers for vulnerabilities and fix them as you work. Snyk tests the container images you have imported on a regular cadence. GAR integration works similarly to other Snyk integrations.&#x20;

## Enable permissions for GAR integration

[Enable the Artifact Registry API](https://cloud.google.com/artifact-registry/docs/enable-service) for the Google account you plan on integrating with Snyk. Allow a few minutes for Google to propagate the enablement.

To enable permissions for GAR integration:

1. Navigate to the Google Cloud Console [Credentials](https://console.cloud.google.com/apis/credentials) page.&#x20;
2. Select the Google project for which you are creating credentials if it is not already selected.
3. Select the **Create Credentials** button and select **Service Account**.
4. Give the new service account a unique name and ID and select **Create**.
5. On the Service account permissions page, choose **Select a role** and **Artifact Registry Reader**. You must also add an additional role that has the resourcemanager.projects.list permission, such as **Browser** or **Viewer**.
6. After the account has been created, select the relevant account in the **Service Accounts** section.
7. In the Service account details page, under the **Keys** section, select **Add Key** and **Create New Key**.
8. Select JSON for the **Key type** and select **Create**.

The JSON key is generated for your Project. Copy the entire contents of the JSON file for the next step: [Configure integration for GAR](integrate-with-google-artifact-registry-gar.md#configure-gar-integration).

## Configure GAR integration

Integrating Snyk with your GAR account enables you to scan for vulnerabilities and fix security and license issues.

To configure GAR integration:

1. Navigate to your Organization in the Snyk Web UI.
2. Select **Integrations**.
3. In the Container Registries section, select **Google Artifact Registry**.
4. In the Account credentials section, enter your [Artifact Registry](https://cloud.google.com/artifact-registry/docs/repositories/repo-locations) hostname. This is typically \<Your Region Name>-docker.pkg.dev, but in some cases, you may need to use a specific region or multi-region, for example, "us-east1-docker.pkg.dev" or "us-docker.pkg.dev".
5. In the JSON key file field, paste the entire contents of the JSON key file you downloaded when [enabling permissions](integrate-with-google-artifact-registry-gar.md#enable-permissions-for-gar-integration).
6. Select **Save**.

Snyk checks the credentials, and upon success, the page reloads with a notification that the connection succeeded.

## Add images from GAR

{% hint style="info" %}
You must be onboarded to your Organization by an administrator and have configured the integration with your GAR repository.
{% endhint %}

Snyk tests and monitors Google Artifact Registry (GAR) container images by evaluating root folders and custom file locations.

To add images from GAR:&#x20;

1. In the Snyk Web UI, navigate to the Organization you want to manage.
2. Navigate to **Projects** > **Add Projects**. The list of integrations already configured in your account opens. The view displays all of the available images for the registry to which you connected, grouped by each of your repositories.
3. Select single or multiple images by using any or all of the following methods:
   * Type the name of a single image for import in the **Image Name** field.
   * Select any of the repositories for which you want to import all of the associated images.
   * Expand and collapse repositories to select multiple images across multiple repositories.
4. Click **Add selected repositories**. As the images are imported, a status bar appears at the top of the page. You can continue working while Snyk imports the images.

When the import ends, a notification of success or failure appears at the top of the page. Click **Refresh** to view the **Projects** page with the newly imported images. The images are grouped by repository and linked individually to a detailed **Projects** page.

You can now connect your Git repo to this Project in order to use your Dockerfile for enriched fix advice. For more information, see [Adding your Dockerfile and testing your base image](../scan-your-dockerfile/detect-vulnerable-base-images-from-your-dockerfile.md).

GAR files are indicated with a unique icon. You can now filter to view only those Projects.

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-31 at 18.06.20.png" alt=""><figcaption><p>Examples of GAR Projects</p></figcaption></figure>

{% hint style="info" %}
For application vulnerabilities within container images, any changes to the application will not be reflected with a manual or recurring retest. You must re-import the image. For more information, see [Detecting application vulnerabilities in container images](../use-snyk-container/detect-application-vulnerabilities-in-container-images.md).
{% endhint %}
