# Adding your Dockerfile and testing your base image

To receive base image fix advice, including major, minor, and alternative upgrades as well as advice when you need to rebuild your image, integrate with your preferred Git repository and import the repo that contains the relevant Dockerfile.

You can add a single Dockerfile to each image that you've imported.

**To add your Dockerfile for additional fix advice:**

## **Prerequisites to add your Dockerfile for additonal fix advice**

* Ensure the relevant Git repository has been configured.
* Import the relevant image from its registry first.

## **Steps to add your Dockerfile for additional fix advice**

1. From the **Project** tab, filter for your project and then click **Settings** to access the settings to add a Dockerfile.
2. From the **Project** settings page, click **Configure Dockerfile** and then select the relevant Git repository.
3. When the **Add Projects** view appears, examine the list of all repositories from the Git account with which you integrated, grouped by Organization and personal account.
4. Checkmark the relevant repo from which to import the Dockerfile.
5. Enter the relative path in the Path to your Dockerfile field in the following format: /path/dockerfile.
6. Click **Save**.

<figure><img src="../../.gitbook/assets/image (45) (1).png" alt="Configure Dockerfile view"><figcaption><p>Configure Dockerfile view</p></figcaption></figure>

Snyk tests the project again, this time producing any relevant base image fix advice shown on the Recommendations for base image upgrade screen:

<figure><img src="../../.gitbook/assets/mceclip1-2-.png" alt="Recommendations for base image upgrade"><figcaption><p>Recommendations for base image upgrade</p></figcaption></figure>

Included are the Current image, Minor upgrade, Alternative upgrade, number of vulnerabilities in each, and a severity ranking.
