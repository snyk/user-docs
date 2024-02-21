# Docker Hub - add Projects and images to the Snyk UI

Snyk tests and monitors Docker Hub repositories and images by evaluating root folders. This page explains how to add repositories to Snyk.

## **Prerequisites for adding Docker Hub images to Snyk**

* You must have an account with Snyk and be onboarded to your Organization by an admin.
* Ensure the integration with Docker Hub and Snyk has already been configured.

## **Steps to add Docker Hub images to Snyk**

1. Log in to your Snyk account and navigate to the relevant Group and Organization you want to manage.
2. Navigate to the **Projects** page and click **Add Projects**.\
   The list of integrations that are already configured on your account opens.\
   The view **Which images do you want to test?** opens, displaying all of the available images for the registry to which you connected, grouped by each of your repositories.
3. Select one or multiple images using any or all of the following methods:
   * Type the name of a single image for import in the **Image Name** field.
   * Select any of the repositories if you want to import all of the associated images.
   * Expand and collapse repositories to select multiple images.
4.  Click **Add selected repositories**.

    On the **Projects** page, a status bar appears at the top of the page as the images are imported; you can continue working in the meantime.
5. When the import ends, a notification of success or failure appears at the top of the **Projects** page.\
   Click **Refresh** to view the **Projects** page with the newly imported images.\
   Images are grouped by repository and are each linked individually to a detailed **Projects** page.
6. You can now connect your Git repository to this Project in order to use your Dockerfile for enriched fix advice. For more information, see [Detect vulnerable base images from your Dockerfile](../../../scan-using-snyk/snyk-container/scan-your-dockerfile/detect-vulnerable-base-images-from-your-dockerfile.md).

You can now also filter to view only Docker Hub Projects:

<figure><img src="../../../.gitbook/assets/uuid-ce306bb8-1d6d-c895-bdb5-3a7cd551977b-en (1) (1) (1) (1) (1) (1) (1) (1) (8) (7).png" alt="Docker Hub projects"><figcaption><p>Docker Hub projects</p></figcaption></figure>

When repositories and images are imported, a confirmation appears in green at the top of the screen. Docker Hub files are indicated with a unique icon .

Docker Hub integration works like other Snyk integrations. To continue to monitor, fix and manage your Projects, see the relevant pages in the Snyk documentation.

{% hint style="info" %}
For application vulnerabilities within container images, any changes to the application will not be reflected with a manual or recurring retest. A re-import of the image is required. See [Detecting application vulnerabilities in container images ](../../../scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/detect-application-vulnerabilities-in-container-images.md)for more information.
{% endhint %}
