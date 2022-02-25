# Add images to Snyk from ACR

Snyk tests and monitors ACR container images by evaluating root folders and custom file locations.

**To add registry images to Snyk:**

**Prerequisites:**

* You must have an account with Snyk and be onboarded to your organization by an administrator.
* The integration must be configured between Snyk and your ACR repository.

## Steps

1. Log in to your account and navigate to the relevant group and organization that you want to manage.
2. Go to **Projects**, and click **Add projects**. The list of integrations already configured on your account opens, similar to the following: ![](../../../../.gitbook/assets/uuid-dd01aab7-482f-0fc2-01de-c2427a14a0e0-en.png)

![](<../../../../.gitbook/assets/add-artifactory-images (1) (2) (45).gif>)

The **Which images do you want to test?** view appears, displaying all of the available images for the registry to which you connected, grouped by each of your repositories, similar to the following:

![](<../../../../.gitbook/assets/uuid-bd9cf629-f5fb-b28b-1fc1-40df2367a7f9-en (1) (1) (2) (4) (2) (9).png>)

1. Select single or multiple images with any or all of the following methods:
   1. Type the name of a single image for import in the **Image Name** field (#1 in the image above),
   2. Select any of the repositories if you want to import all of the associated images (#2 in the image above).
   3. Expand and collapse repositories to select multiple images (#3 in the image above) across multiple repositories.
2.  Click **Add selected repositories**.

    A status bar appears at the top of the page as the images are imported; you can continue working in the meantime.
3. When the import ends, notification of success, or failure, appears at the top of the page. Click **Refresh** to view the **Projects** page with the newly imported images. Images are grouped by repository and are each linked individually to a detailed Projects page.
4. You can now connect your Git repo to this project in order to use your Dockerfile for enriched fix advice. For more info, see [Adding your Dockerfile and test your base image](https://support.snyk.io/hc/articles/360003916218#UUID-9ab347a6-8af0-ef6c-5ebd-cec21fbfab29).

ACR files are indicated with a unique icon ![](../../../../.gitbook/assets/uuid-5d10608d-d674-d4ee-d6c2-6faadd6fc8ea-en.png) . You can now also filter to view only those projects:

![](<../../../../.gitbook/assets/image (4) (3) (3) (3) (3) (4) (4) (5) (4) (9).png>)

ACR integration works similar to our other integrations. To continue to monitor, fix and manage your projects, see the relevant pages, also in our docs.
