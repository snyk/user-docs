# Amazon Elastic Container Registry \(ECR\): add images to Snyk

Snyk tests and monitors your Amazon ECR container images by evaluating its tags as they are in your ECR repositories.

**To add images to Snyk:**

**Prerequisites:**

* You must have an account with Snyk and be onboarded to your organization by an administrator.
* Ensure you've configured the integration between Snyk and your ECR repository.

**Steps:**

1. Log in to your account and navigate to the relevant group and organization that you want to manage.

![AddProjectMenu.gif - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](https://support.snyk.io/hc/article_attachments/360007147318/uuid-da316a4a-c823-cf03-f37f-5305446dc970-en.gif)

2. Go to Projects, and click Add projects. The list of integrations already configured on your account opens, similar to the following: ![AddProjectMenu.png - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](https://support.snyk.io/hc/article_attachments/360007066017/uuid-dd01aab7-482f-0fc2-01de-c2427a14a0e0-en.png)
   1. The Which images do you want to test? view appears, displaying all of the available images for the registry to which you connected, grouped by each of your repositories, similar to the following:
   2. ![AddImages.png - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](https://support.snyk.io/hc/article_attachments/360007147338/uuid-bd9cf629-f5fb-b28b-1fc1-40df2367a7f9-en.png)
3. Select single or multiple images with any or all of the following methods:
   * Type the name of a single image for import in the Image Name field \(\#1 in the image above\),
   * Select any of the repositories if you want to import all of the associated images \(\#2 in the image above\).
   * Expand and collapse repositories to select multiple images \(\#3 in the image above\) across multiple repositories.
4. Click Add selected repositories. 
   1. A status bar appears at the top of the page as the images are imported; you can continue working in the meantime. 
5. When the import ends, notification of success, or failure, appears at the top of the page. Click Refresh to view the Projects page with the newly imported images. Images are grouped by repository and are each linked individually to a detailed Projects page.
6. You can now connect your Git repo to this project in order to use your Dockerfile for enriched remediation advice. For more info, see [Adding your Dockerfile and test your base image](https://support.snyk.io/hc/articles/360003916218#UUID-9ab347a6-8af0-ef6c-5ebd-cec21fbfab29). 
   1. ECR files are indicated with a unique icon![image6.png - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](https://support.snyk.io/hc/article_attachments/360007147358/uuid-31aa2b29-8686-5389-b5fc-1d3bd1176f9c-en.png), You can now filter to view only those projects: ![image19.png - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](https://support.snyk.io/hc/article_attachments/360007066037/uuid-439e3f37-6e4f-0ffa-0c3c-63c56b45ba5a-en.png)
   2. Amazon ECR integration works similar to our other integrations. To continue to monitor, remediate and manage your projects, see the relevant pages, also in our docs.

