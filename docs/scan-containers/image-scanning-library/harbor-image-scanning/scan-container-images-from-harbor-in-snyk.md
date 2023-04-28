# Scan container images from Harbor in Snyk

Snyk tests and monitors your Harbor container images by evaluating its tags in your repositories. Once your container images are imported to Snyk, your image vulnerabilities are surfaced and can be triaged easily.

Follow the instructions on this page to add images from Harbor to Snyk.

## **Prerequisites for Harbor container image scanning**

* Have a Snyk account with access to the relevant organization (given by an administrator).
* Harbor integration configured. To learn more about that, follow the steps in [Container security with Harbor integration](https://docs.snyk.io/snyk-container/image-scanning-library/harbor-image-scanning/container-security-with-harbor-integration).

## **Steps in scanning Harbor images**

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Under the **Projects** tab, click **Add project**. The list of integrations already configured on your account opens. Select the **Harbor** option or **Other** if Harbor does not appear.
3. **Which images do you want to test?** appears, displaying all of the available images for your connected registry, grouped by each of your repositories.
4. Select single or multiple images to be imported to Snyk, by choosing a specific image or selecting an entire repository. You can also search by image name to find specific images to import. To finish, click **Add selected repositories** on the top-right.
5. A status bar appears at the top of the page as the images are imported; you can continue working in the meantime.
6. When the import ends:
   1. You can view the newly imported image in the **Projects** page (marked with a **NEW** tag). Images are grouped by repository and are each linked individually to a detailed **Project** page.
   2. An **import log** becomes available, reachable from the top of the projects list.
   3. To enrich the data and get recommendations regarding your base image, nuder **Settings** you can connect your Dockerfile to the image project. For more info, see [Adding your Dockerfile and test your base image](../../scan-your-dockerfile/adding-your-dockerfile-and-testing-your-base-image.md).

Harbor imports are indicated with a unique icon. You can also filter to view only the Harbor projects:

<figure><img src="../../../.gitbook/assets/mceclip1-9-.png" alt="Projects listing showing container images"><figcaption><p>Projects listing showing container images</p></figcaption></figure>

{% hint style="info" %}
For **application** vulnerabilities within container images, any changes to the application will not be reflected with a manual or recurring retest. A re-import of the image is required. See [Detecting application vulnerabilities in container images ](../../getting-around-the-snyk-container-ui/detecting-application-vulnerabilities-in-container-images.md)for more information.
{% endhint %}
