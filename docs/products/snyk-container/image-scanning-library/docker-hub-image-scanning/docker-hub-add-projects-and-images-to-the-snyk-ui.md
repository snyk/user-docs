# Docker Hub: add projects and images to the Snyk UI

Snyk tests and monitors Docker Hub repositories and images by evaluating root folders.

**To add repositories to Snyk:**

**Prerequisites:**

* You must have an account with Snyk and be onboarded to your organization by an admin.
* Ensure the integration with Docker Hub and Snyk has already been configured.

**Steps:**

1. Log in to your account and navigate to the relevant group and organization that you want to manage.
2. Go to Projects, and click Add projects. The list of integrations already configured on your account opens, similar to the following:
   1. The Which images do you want to test? view appears, displaying all of the available images for the registry to which you connected, grouped by each of your repositories, similar to the following:
3. Select single or multiple images with any or all of the following methods:
   * Type the name of a single image for import in the Image Name field (#1 in the image above),
   * Select any of the repositories if you want to import all of the associated images (#2 in the image above).
   * Expand and collapse repositories to select multiple images (#3 in the image above) across multiple repositories.
4.  Click **Add selected repositories**.

    A status bar appears at the top of the page as the images are imported; you can continue working in the meantime.
5. When the import ends, notification of success, or failure, appears at the top of the page. Click **Refresh** to view the Projects page with the newly imported images. Images are grouped by repository and are each linked individually to a detailed **Projects** page.
6. You can now connect your Git repo to this project in order to use your Dockerfile for enriched fix advice. For more info, see [Adding your Dockerfile and test your base image](https://support.snyk.io/hc/articles/360003916218#UUID-9ab347a6-8af0-ef6c-5ebd-cec21fbfab29).

You can now also filter to view only those projects:

![](<../../../../.gitbook/assets/uuid-ce306bb8-1d6d-c895-bdb5-3a7cd551977b-en-1- (1) (1) (2) (2) (2) (2) (2) (2) (2) (2) (2) (2) (2) (2) (3).png>)

Once repositories and images are imported, a confirmation appears in green at the top of the screen. Docker Hub files are indicated with a unique icon ![](../../../../.gitbook/assets/docker-hub-logo.png) .

Docker Hub integration works similar to our other integrations. To continue to monitor, fix and manage your projects, see the relevant pages in our docs.
