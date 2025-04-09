# Add images to Snyk from ACR

Snyk tests and monitors Microsoft Azure Container Registry (ACR) container images by evaluating root folders and custom file locations.

## **Prerequisites to add images to Snyk from ACR**

* You must have an account with Snyk and be onboarded to your Organization by an administrator.
* The integration must be configured between Snyk and your ACR repository.

## Steps to add images to Snyk from ACR

1. Log in to your account and navigate to the relevant Group and Organization you want to manage.
2. Go to **Projects**, and click **Add projects**. The list of integrations already configured on your account opens.

The view **Which images do you want to test?** opens, displaying all of the available images for the registry to which you connected, grouped by each of your repositories, similar to the following:

1. Select single or multiple images using any or all of the following methods:
   1. Type the name of a single image for import in the **Image Name** field (at number 1, the Image Name field in the image above),
   2. Select any of the repositories if you want to import all of the associated images (at number 2, the second item listed in the image above).
   3. Expand and collapse repositories to select multiple images at number 3, the third item listed in the image above) across multiple repositories.
2.  Click **Add selected repositories**.

    A status bar appears at the top of the page as the images are imported; you can continue working in the meantime.
3. When the import ends, a notification of success or failure appears at the top of the page. Click **Refresh** to view the **Projects** page with the newly imported images.\
   Images are grouped by repository and linked individually to a detailed Projects page.
4. You can now connect your Git repository to this Project in order to use your Dockerfile for enriched fix advice.\
   For more information, see [Adding your Dockerfile and testing your base image](../../scan-your-dockerfile/detect-vulnerable-base-images-from-your-dockerfile.md).

ACR files are indicated with a unique icon. You can filter to view only ACR Projects.

ACR integration works like other Snyk integrations. To continue to monitor, fix and manage your Projects, see the relevant pages in the Snyk user documentation.

{% hint style="info" %}
For application vulnerabilities within container images, any changes to the application will not be reflected with a manual or recurring retest. A re-import of the image is required. For more information, see [Detecting application vulnerabilities in container images](../../use-snyk-container/detect-application-vulnerabilities-in-container-images.md).&#x20;
{% endhint %}
