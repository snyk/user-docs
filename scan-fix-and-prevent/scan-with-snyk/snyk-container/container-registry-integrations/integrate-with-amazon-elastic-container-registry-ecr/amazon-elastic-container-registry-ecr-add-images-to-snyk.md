# Amazon Elastic Container Registry (ECR) - add images to Snyk

Snyk scans and monitors your Amazon ECR container images by evaluating the tags as they are in your ECR repositories.

## **Prerequisites for adding images to Snyk from ECR**

* You must have an account with Snyk and be onboarded to your Organization by an administrator.
* Ensure you have configured the integration between Snyk and your ECR repository.

## **How to add images to Snyk from ECR**

Log in to your account and navigate to the relevant Group and Organization you want to manage.

1. Navigate to **Projects** > **Add projects**.\
   The list of integrations already configured on your account opens.&#x20;
2. From the list of integrations, select **ECR**. The view **Which ECR images do you want to test?** opens and displays all available images for the registry you connected, grouped by each of your repositories.
3. Select one or multiple images with any or all of the following methods:
   1. To import a specific image, type its name in the **Image Name** field.
   2. If you want to import all of the associated images, select any of the repositories.
   3. To select multiple images, expand and collapse repositories.
4. Click **Add selected images.**\
   A status bar appears at the top of the page as the images are imported; you can continue working in the meantime.
5. When the import ends, a notification of success or failure appears at the top of the page. Click **Refresh Page** to view the **Projects** page with the newly imported images.\
   Images are grouped by repository and linked individually to a detailed **Projects** page.
6. You can now connect your Git repo to this Project to use your Dockerfile for enriched fix advice.\
   For more information, see [Detect vulnerable base images from your Dockerfile](../../scan-your-dockerfile/detect-vulnerable-base-images-from-your-dockerfile.md).

You can filter to view only ECR Projects, which are marked with a unique icon.

<figure><img src="../../../../.gitbook/assets/projects_amazon_ecr_projects.png" alt="Projects list displaying an ECR Project"><figcaption><p>Example of an ECR Project</p></figcaption></figure>

Amazon ECR integration works like other Snyk integration. To continue to monitor, fix, and manage your Projects, see the relevant pages in the Snyk user documentation.

{% hint style="info" %}
For application vulnerabilities within container images, any changes to the application will not be reflected with a manual or recurring retest. A re-import of the image is required. See [Detecting application vulnerabilities in container images ](../../use-snyk-container/detect-application-vulnerabilities-in-container-images.md)for more information.
{% endhint %}
