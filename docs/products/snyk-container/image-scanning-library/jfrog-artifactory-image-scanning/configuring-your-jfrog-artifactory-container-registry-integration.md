# Configuring your JFrog Artifactory container registry integration

Enable integration between one Artifactory instance as a container registry and a Snyk organization to start managing your image security.

## Prerequisites

* You must be an administrator for the organization you're configuring in Snyk.
* Snyk needs user credentials to integrate with Artifactory and does not support Artifactory when configured for single sign-on (SSO).
* If you're using self-hosted Artifactory, refer to our documentation explaining about [broker configuration](../../integrate-self-hosted-container-registries/snyk-integration-to-self-hosted-container-registries.md).

## Configure integration

1. Log in to [your Snyk account](https://app.snyk.io).
2.  Navigate to **Integrations** from the menu bar at the top, find and click the Artifactory option:

    ![](<../../../../.gitbook/assets/image (57).png>)
3. The configuration page in the **Settings** area loads.
4. Enter credentials as follows:
   1. **Username and Password**—use your Artifactory login credentials.
   2. **Container registry name -** the _full registry URL_ in format `<org>.jfrog.io/artifactory/api/docker/<repo-name>`.
5. Click **Save Changes**.

![Screenshot 2021-12-06 at 16 37 12](https://user-images.githubusercontent.com/112600/144875482-078b715e-2834-469b-9983-7e88a65f175e.png)

![](../../../../.gitbook/assets/uuid-3b329a90-394f-5ab3-af84-658b41a1edc0-en.png)

Snyk tests the connection values and the page reloads, now displaying integration details as you entered them. A confirmation message that the details were saved also appears in green at the top of the screen. In addition, if the connection failed, a notification appears.
