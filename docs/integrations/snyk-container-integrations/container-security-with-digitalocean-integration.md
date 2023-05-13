# Container security with DigitalOcean integration

Snyk integrates with DigitalOcean to enable you to import your container images and monitor them for vulnerabilities.

Snyk tests the images you’ve imported (referred to as \`projects\`) for any known security vulnerabilities, testing them at a frequency you control and alerts you when new issues are detected.

Integration with DigitalOcean is available for all Snyk users.

To set up DigitalOcean integration in Snyk and start managing image vulnerabilities:

## Prerequisites for Digital Ocean integration

* You must be an administrator for the Organization you're configuring in Snyk.
* Snyk needs user credentials to integrate with DigitalOcean and does not support DigitalOcean when configured for single sign-on (SSO).

## **Configure DigitalOcean integration**

1. In your Snyk account, navigate to Integrations from the menu bar at the top. Under the Container Registries section, find the **DigitalOcean** option and click it.
2. In the **Account credentials** section, enter your DigitalOcean personal access token as the login credential. Detailed instructions to create the access token can be found in the integration page. To finish, click **Save**.

![](../../.gitbook/assets/mceclip0-10-.png)

In case you are using a self-hosted DigitalOcean, contact us to provide you with a token. You can read more about setting up private registry integration [here](https://docs.snyk.io/snyk-container/integrate-self-hosted-container-registries/snyk-integration-to-self-hosted-container-registries).

{% hint style="info" %}
**Note**\
For the connection to succeed, make sure you have a repository in DigitalOcean.
{% endhint %}

Snyk tests the connection values and the page reloads, now displaying DigitalOcean integration information, and the **Add your DigitalOcean images to Snyk** button becomes available. In case the connection to DigitalOcean failed, notification appears under the **Connected to DigitalOcean** section.\
Now you can use Snyk to scan your images from DigitalOcean.

Snyk tests and monitors your DigitalOcean container images by evaluating its tags in your repositories. Once imported to Snyk, your image vulnerabilities are surfaced and can be triaged easily.

To add images from DigitalOcean to Snyk:

## **Prerequisites for scanning container images from DigitalOcean in Snyk**

* Have a Snyk account with access to the relevant organization (given by an administrator).
* DigitalOcean integration configured. To learn more about that, follow the steps in [Container security with DigitalOcean integration](https://docs.snyk.io/snyk-container/image-scanning-library/digitalocean-image-scanning/container-security-with-digitalocean-integration)

## **Steps in scanning container images from DigitalOcean in Snyk**

1. Log in to your account and navigate to the relevant group and organization that you want to manage.
2. Under the **Projects** tab, click **Add project**. The list of integrations already configured on your account opens. Select the **DigitalOcean** option or **Other** if **DigitalOcean** doesn’t appear.
3. The **Which images do you want to test?** view appears, displaying all of the available images for your connected registry, grouped by each of your repositories, similar to the following:
4. Select single or multiple images to be imported to Snyk. Selection can be done by choosing a specific image or selecting an entire repository. You can also search by image name to find specific images to import. To finish, click **Add selected repositories** on the top-right.
5. A status bar appears at the top of the page as the images are imported; you can continue working in the meantime.
6. When the import ends:
   1. You can view the newly imported image in the **Projects** page (marked with a **NEW** tag). Images are grouped by repository and are each linked individually to a detailed **Project** page.
   2. An **import log** becomes available and can be reached at the top of the projects list.
   3. To enrich the data and get recommendations regarding your base image, you can connect your Dockerfile to the image project, under **Settings**. For more info, see [Adding your Dockerfile and test your base image](../../scan-containers/scan-your-dockerfile/adding-your-dockerfile-and-testing-your-base-image.md).

DigitalOcean imports are indicated with a unique icon, and you can also filter the integration in the **projects** view to see DigitalOcean projects only:

![](../../.gitbook/assets/mceclip0-11-.png)

{% hint style="info" %}
For **application** vulnerabilities within container images, any changes to the application will not be reflected with a manual or recurring retest. A re-import of the image is required. See [Detecting application vulnerabilities in container images ](../../scan-containers/using-snyk-container/detecting-application-vulnerabilities-in-container-images.md)for more information.
{% endhint %}
