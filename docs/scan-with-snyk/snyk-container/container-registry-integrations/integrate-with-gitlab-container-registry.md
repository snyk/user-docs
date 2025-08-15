# Integrate with GitLab Container Registry

Snyk integrates with GitLab Container Registry to enable you to import your container images and monitor them for vulnerabilities.

Snyk tests the images you have imported (Projects) for any known security vulnerabilities, testing them at a frequency you control, and alerts you when new issues are detected.

Integration with GitLab Container Registry is available for all Snyk users.

This page explains how to set up GitLab Container Registry integration in Snyk and start managing image vulnerabilities.

## Integrate with GitLab Container Registry

### **Prerequisites for GitLab Container Registry integration**

* You must be an administrator for the Organization you are configuring in Snyk.
* Snyk needs a [personal access token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html) to integrate with GitLab Container Registry. Snyk does not support integration when GitLab SSO is in use for access to the GitLab Container Registry.

## **Configure GitLab Container Registry integration**

1. In your Snyk account, navigate to **Integrations**. Under the Container Registries section, find the GitLab Container Registry option and click it.
2. In the **Account credentials** section, enter your GitLab Container Registry username and [personal access token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html).
3. In the **Container registry name** fill in the full URL to the registry you want to integrate with.
4. To finish, click **Save**.

If you are using a self-hosted GitLab Container Registry, [contact Snyk Support](https://support.snyk.io/) to provide you with a token. For more information, see [Snyk Container for self-hosted container registries (with Broker)](../../../implementation-and-setup/enterprise-setup/snyk-broker/snyk-broker-container-registry-agent/integrate-with-self-hosted-container-registries-broker.md).

Snyk tests the connection values, and the page reloads, now displaying GitLab Container Registry integration information. The **Add your GitLab container registry images to Snyk** button becomes available.

If the connection to GitLab Container Registry fails, a notification appears under the **Connected to GitLab container registry** section.

When the connection is successful, you can use Snyk to scan your images from GitLab Container Registry.

## Scan images from GitLab Container Registry in Snyk

Snyk tests and monitors your GitLab container images by evaluating the image tags in your repositories. After you have imported images to Snyk, your image vulnerabilities are identified and can be triaged easily.

The steps follow for adding images from GitLab Container Registry to Snyk.

## **Prerequisites for GitLab container image scanning**

* Have a Snyk account with access to the relevant Organization authorized by an administrator.
* GitLab Container Registry integration configured.

## **Steps in scanning GitLab container images**

1. Log in to your account and navigate to the relevant Group and Organization that you want to manage.
2. Under the **Projects** tab, click **Add project**.\
   The list of integrations already configured on your account opens.
3. Select the **GitLab container registry** option or **Other** if **GitLab container registry** does not appear.
4. The view titled **Which images do you want to test?** opens, displaying all available images for your connected registry grouped by each of your repositories.\
   Note that GitLab Container Registry does not follow Docker v2 API. Therefore, it is not possible to list images in a repository, and you must manually specify the images you wish to scan.
5. Select single or multiple images to be imported to Snyk.\
   You can select by choosing a specific image or selecting an entire repository. You can also search by image name to find specific images to import.
6. To finish, click **Add selected repositories** on the top right.\
   A status bar appears at the top of the page as the images are imported; you can continue working in the meantime.
7. When the import ends:
   * You can view the newly imported image, marked with a **NEW** tag, on the **Projects** page. Images are grouped by repository and are each linked individually to a detailed **Project** page.
   * An import log becomes available; you can reach it at the top of the Projects list.
   * To enrich the data and get recommendations regarding your base image, under **Settings**, you can connect your Dockerfile to the image Project. For more information, see[ Adding your Dockerfile and testing your base image.](../scan-your-dockerfile/detect-vulnerable-base-images-from-your-dockerfile.md)

GitLab Container Registry imports are designated with a unique icon. You can filter the integration in the **Projects** view to see GitLab Container Registry Projects only.

<figure><img src="../../../.gitbook/assets/container_registry_integrations_gitlab.png" alt=""><figcaption><p>Example of a GitLab Project</p></figcaption></figure>

{% hint style="info" %}
For application vulnerabilities within container images, any changes to the application will not be reflected with a manual or recurring retest. A re-import of the image is required. For more information, see [Detecting application vulnerabilities in container images](../use-snyk-container/detect-application-vulnerabilities-in-container-images.md).
{% endhint %}
