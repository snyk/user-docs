# Add images to Snyk from GCR

Snyk tests and monitors Google Container Registry (GCR) container images by evaluating root folders and custom file locations.

## **Prerequisites for adding images to Snyk from GCR**&#x20;

* You must have an account with Snyk and be onboarded to your Organization by an administrator.
* The integration must be configured between Snyk and your GCR repository.

## Steps to add images to Snyk from GCR&#x20;

Log in to your account and navigate to the relevant Group and Organization you want to manage.

Go to **Projects** and click **Add projects**. The list of integrations already configured in your account opens, similar to the following:&#x20;

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-31 at 17.39.47.png" alt="Integrations configured in a Snyk account"><figcaption><p>Integrations configured in a Snyk account</p></figcaption></figure>

Decide which images you want to test. The view displays all of the available images for the registry to which you connected, grouped by each of your repositories, similar to the following:

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-31 at 17.50.25.png" alt="All available images for the registry to which you connected"><figcaption><p>All available images for the registry to which you connected</p></figcaption></figure>

Continue by following these steps:

1. Select single or multiple images by using any or all of the following methods:
   * Type the name of a single image for import in the **Image Name** field at number one (1), the Image Name field in the preceding image.
   * Select any of the repositories for which you want to import all of the associated images at number two (2), the first item in the list in the preceding image.
   * Expand and collapse repositories to select multiple images across multiple repositories at number three (3), the third item in the list in the preceding image.
2.  Click **Add selected repositories**.

    A status bar appears at the top of the page as the images are imported; you can continue working while the images are being imported.
3. When the import ends, a notification of success or failure appears at the top of the page. Click **Refresh** to view the **Projects** page with the newly imported images. Images are grouped by repository and are each linked individually to a detailed Projects page.
4. You can now connect your Git repo to this Project in order to use your Dockerfile for enriched fix advice. For more information, see [Adding your Dockerfile and testing your base image](../../../scan-containers/scan-your-dockerfile/adding-your-dockerfile-and-testing-your-base-image.md).

GCR files are indicated with a unique icon. You can now filter to view only those Projects:

<div align="left">

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-31 at 18.10.17 (1).png" alt="GCR added to integrations filters"><figcaption><p>GCR added to integrations filters</p></figcaption></figure>

</div>

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-31 at 18.06.20.png" alt="Projects listing with two filters applied"><figcaption><p>Projects listing with two filters applied</p></figcaption></figure>

GCR integration works similarly to other Snyk integrations. To continue to monitor, fix, and manage your Projects, see the relevant pages in the Snyk docs.

{% hint style="info" %}
For **application** vulnerabilities within container images, any changes to the application will not be reflected with a manual or recurring retest. You must re-import the image. See [Detecting application vulnerabilities in container images ](../../../scan-applications/snyk-container/use-snyk-container/detect-vulnerabilities-in-container-images.md)for more information.
{% endhint %}
