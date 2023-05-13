# Container security with GitHub Container registry integration

Snyk integrates with the GitHub Container registry to enable you to import your container images and monitor them for vulnerabilities.

Snyk tests the images you have imported (Projects) for any known security vulnerabilities, testing them at a frequency you control, and alerts you when new issues are detected.

Integration with GitHub Container registry is available for all Snyk users.

To set up GitHub Container registry integration in Snyk and start managing image vulnerabilities:

## **Prerequisites for GitHub Container registry integration**

* You must be an administrator for the Organization you're configuring in Snyk.
* Snyk needs user credentials to integrate with GitHub Container Registry and does not support GitHub Container Registry when configured for single sign-on (SSO). However, you can use a Personal Access Token (PAT) with SSO when the token is authorized with the`read:packages`scope.

## **Configure integration of GitHub Container registry**

1. In your Snyk account, navigate to **Integrations** from the menu bar at the top. Under the Container **Registries** section, find the GitHub container registry option and click it.
2. Enter your GitHub container registry username and password login credentials in the Account credentials section. In the **container registry name** fill in the full URL to the registry you want to integrate with. To finish, click **Save**.

![mceclip1.png](../../.gitbook/assets/mceclip1-4-.png)

If you are using a self-hosted GitHub Container registry, [contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new) to provide you with a token. For more information about setting up private registry integration, see [Snyk Container for self-hosted container registries (with Broker)](../../scan-containers/integrate-self-hosted-container-registries.md).

Snyk tests the connection values, and the page reloads, now displaying GitHub Container registry integration information, and the **Add your GitHub container registry images to Snyk** button becomes available.

If the connection to the GitHub Container registry fails, a notification appears under the **Connected to GitHub container registry** section. Now you can use Snyk to scan your images from the GitHub Container registry.

See the Snyk blog to learn more about [container registry security & security concerns for using a container registry.](https://snyk.io/learn/container-security/container-registry-security/)

Snyk tests and monitors your GitHub container images by evaluating their tags in your repositories. Once imported to Snyk, your image vulnerabilities are surfaced and can be triaged efficiently.

To add images from the GitHub container registry to Snyk:

## **Prerequisites for scanning images from GitHub Container registry in Snyk**

* Have a Snyk account with access to the relevant organization (given by an administrator).
* GitHub container registry integration configured. To learn more about that, follow the steps in [Container security with GitHub container registry integration](https://docs.snyk.io/snyk-container/image-scanning-library/github-container-registry-image-scanning/container-security-with-github-container-registry-integration).

## **Steps in scanning images from GitHub Container registry in Snyk**

1. Log in to your account and navigate to the relevant group and organization you want to manage.
2. Under the **Projects** tab, click **Add project**. The list of integrations already configured on your account opens. Select the **GitHub container registry** option or **Other** if the **GitHub container registry** doesn’t appear.
3. The view titled **Which images do you want to test?** opens, displaying all available images for your connected registry, grouped by each of your repositories. **Note**: GitHub Container Registry doesn't follow docker v2 API. Therefore, it is not possible to list images in repos. Due to that, you will need to specify the images you wish to scan manually.
4. Select single or multiple images to be imported to Snyk. Selection can be made by choosing a specific image or an entire repository. You can also search by image name to find specific images to import. To finish, click **Add selected repositories** on the top-right.
5. A status bar appears at the top of the page as the images are imported; you can continue working in the meantime.
6. When the import ends:
   * You can view the newly imported image on the **Projects** page (marked with a **NEW** tag). Images are grouped by repository and are each linked individually to a detailed **Project** page.
     * An **import log** becomes available and can be reached at the top of the projects list.
   * To enrich the data and get recommendations regarding your base image, you can connect your Dockerfile to the image project under **Settings**. For more info, see [Adding your Dockerfile and test your base image](https://support.snyk.io/hc/articles/360003916218#UUID-9ab347a6-8af0-ef6c-5ebd-cec21fbfab29).

GitHub container registry imports are indicated with a unique icon, and you can also filter the integration in the **Projects** view to see GitHub container registry projects only:

![](../../.gitbook/assets/mceclip1-5-.png)

{% hint style="info" %}
For **application** vulnerabilities within container images, any changes to the application will not be reflected with a manual or recurring retest. A re-import of the image is required. See [Detecting application vulnerabilities in container images ](../../scan-containers/using-snyk-container/detecting-application-vulnerabilities-in-container-images.md)for more information.
{% endhint %}
