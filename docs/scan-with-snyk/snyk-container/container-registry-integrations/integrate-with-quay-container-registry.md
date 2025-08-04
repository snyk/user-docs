# Integrate with Quay Container Registry

Snyk integrates with Quay Container Registry to enable you to import your container images and monitor them for vulnerabilities.

Snyk tests the images you have imported (Projects) for any known security vulnerabilities, testing them at a frequency you control, and alerts you when new issues are detected.

Integration with Quay is available for all Snyk users.

## Integrate with Quay Container Registry

### Prerequisites for Quay integration

* You must be an administrator for the Organization you are configuring in Snyk.
* Snyk needs user credentials to integrate with Quay and does not support Quay when it is configured for single sign-on (SSO).

### **Configure Quay integration**

1. In your Snyk account, navigate to **Integrations.** Under the **Container Registries** section, find the **Quay** option and click it.
2. In the **Account credentials** section, enter your Quay username and password login credentials. In the **container registry name**, enter the full URL to the registry you want to integrate with. This can be a cloud-based Quay or a private host. To finish, click **Save**.

If you are using a self-hosted Quay registry, [contact Snyk Support](https://support.snyk.io) to provide you with a token. For more information, see [Snyk Container for self-hosted container registries (with Broker)](../../../implementation-and-setup/enterprise-setup/snyk-broker/snyk-broker-container-registry-agent/integrate-with-self-hosted-container-registries-broker.md).

{% hint style="info" %}
[Quay.io](https://quay.io/) deprecated the use of Quay login in June 2021. The credentials can no longer be Quay username and password, but must be Quay robot account credentials (username and token) that have at least 'read' permissions to the desired repository.
{% endhint %}

Snyk tests the connection values, and the page reloads, displaying Quay integration information, and the **Add your Quay images to Snyk** button becomes available. If the connection to Quay fails, a notification appears under the **Connected to Quay** section.

When the connection succeeds, you can use Snyk to scan your images from Quay.

Snyk tests and monitors your Quay container images by evaluating the tags in your repositories. After import to Snyk, your image vulnerabilities are identified and can be triaged easily.

## Scan container images from Quay in Snyk

### **Prerequisites** for Quay image scanning

* You must have a Snyk account with access to the relevant Organization given by an administrator.
* Quay integration must be configured. For details, see [Integrate with Quay Container Registry](integrate-with-quay-container-registry.md#integrate-with-quay-container-registry).

### **Steps in scanning Quay images**

1. Log in to your account and navigate to the relevant group and Organization that you want to manage.
2. Under the **Projects** tab, click **Add project**.\
   The list of integrations already configured on your account opens.
3. Select the **Quay** option or **Other** if **Quay** does not appear.
4. The view titled **Which images do you want to test?** opens, displaying all available images for your connected registry grouped by each of your repositories.
5. Select single or multiple images to be imported to Snyk.\
   You can select by choosing a specific image or by selecting an entire repository.\
   You can also search by image name to find specific images to import.
6. To finish, click **Add selected repositories** at the top right.\
   A status bar appears at the top of the page as the images are imported; you can continue working in the meantime.
7. When the import ends:
   1. You can view the newly imported image, marked with a **NEW** tag, on the **Projects** page; Images are grouped by repository and linked individually to a detailed **Project** page.\
      An **import log** becomes available; you can reach it at the top of the Projects list.
   2. To enrich the data and get recommendations regarding your base image, from the **Settings** you can connect your Dockerfile to the image Project. For more information, see [Adding your Dockerfile and testing your base image](../scan-your-dockerfile/detect-vulnerable-base-images-from-your-dockerfile.md).

Quay imports are designated with a unique icon. You can filter to view only Quay Projects.

{% hint style="info" %}
For application vulnerabilities within container images, any changes to the application will not be reflected with a manual or recurring retest. A re-import of the image is required. For more information, see [Detecting application vulnerabilities in container images](../use-snyk-container/detect-application-vulnerabilities-in-container-images.md).
{% endhint %}
