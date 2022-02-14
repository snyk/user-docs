# Container security with GitHub container registry integration

Snyk integrates with GitHub container registry to enable you to import your container images and monitor them for vulnerabilities.

Snyk tests the images you’ve imported (referred to as \`projects\`) for any known security vulnerabilities, testing them at a frequency you control and alerts you when new issues are detected.

Integration with GitHub container registry is available for all Snyk users.

To set up GitHub container registry integration in Snyk and start managing image vulnerabilities:

Prerequisites

* You must be an administrator for the organization you're configuring in Snyk.
* Snyk needs user credentials to integrate with GitHub container registry and does not support GitHub container registry when configured for single sign-on (SSO).

**Configure integration**

1. In your Snyk account, navigate to **Integrations** from the menu bar at the top. Under the Container **Registries** section, find the GitHub container registry option and click it.
2. In the **Account credentials** section, enter your GitHub container registry username and password login credentials. In the **container registry name** fill in the full URL to the registry you want to integrate with. To finish, click **Save**.

![mceclip1.png](../../../../.gitbook/assets/mceclip1-4-.png)

If you are using a self-hosted GitHub container registry registry, [contact our Support team](https://support.snyk.io/hc/en-us/requests/new) to provide you with a token. You can read more about setting up private registry integration [here](https://docs.snyk.io/snyk-container/integrate-self-hosted-container-registries/snyk-integration-to-self-hosted-container-registries).

Snyk tests the connection values and the page reloads, now displaying GitHub container registry integration information, and the **Add your GitHub container registry images to Snyk** button becomes available.

In case the connection to GitHub container registry failed, notification appears under the **Connected to GitHub container registry** section. Now you can use Snyk to scan your images from GitHub container registry.
