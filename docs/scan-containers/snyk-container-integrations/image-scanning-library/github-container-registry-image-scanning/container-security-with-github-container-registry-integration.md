# Container security with GitHub container registry integration

Snyk integrates with the GitHub container registry to enable you to import your container images and monitor them for vulnerabilities.

Snyk tests the images you’ve imported (called \`projects\`) for any known security vulnerabilities, testing them at a frequency you control, and alerts you when new issues are detected.

Integration with GitHub container registry is available for all Snyk users.

To set up GitHub container registry integration in Snyk and start managing image vulnerabilities:

**Prerequisites**

* You must be an administrator for the organization you're configuring in Snyk.
* Snyk needs user credentials to integrate with GitHub Container Registry and does not support GitHub Container Registry when configured for single sign-on (SSO). However, you can use a Personal Access Token (PAT) with SSO when the token is authorized with the`read:packages`scope.

**Configure integration**

1. In your Snyk account, navigate to **Integrations** from the menu bar at the top. Under the Container **Registries** section, find the GitHub container registry option and click it.
2. Enter your GitHub container registry username and password login credentials in the Account credentials section. In the **container registry name** fill in the full URL to the registry you want to integrate with. To finish, click **Save**.

![mceclip1.png](../../../../.gitbook/assets/mceclip1-4-.png)

If you are using a self-hosted GitHub container registry, [contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new) to provide you with a token. For more information about setting up private registry integration, see [Snyk Container for self-hosted container registries (with broker)](../../../integrate-self-hosted-container-registries.md).

Snyk tests the connection values, and the page reloads, now displaying GitHub container registry integration information, and the **Add your GitHub container registry images to Snyk** button becomes available.

If the connection to the GitHub container registry fails, a notification appears under the **Connected to GitHub container registry** section. Now you can use Snyk to scan your images from the GitHub container registry.

Check out the Snyk blog to learn more about [container registry security & security concerns for using a container registry.](https://snyk.io/learn/container-security/container-registry-security/)
