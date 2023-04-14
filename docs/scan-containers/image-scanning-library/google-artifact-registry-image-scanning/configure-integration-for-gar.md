# Configure integration for GAR

Configure integration from Snyk with your Google Artifact Registry account to scan for vulnerabilities and fix security and license issues accordingly.

## Prerequisites

[Enable permissions to access Google Artifact Registry.](enable-permissions-to-access-gar.md)

## Steps

1. Navigate to your Organization in the Snyk Web UI.
2. Select **Integrations** in the left navigation bar.
3. In the Container Registries section, select **Google Artifact Registry**.
4. In the Account credentials section, enter your Artifact Registry hostname.
5. In the JSON key file field, paste the entire contents of the JSON key file you downloaded [in the previous step](enable-permissions-to-access-gar.md).
6. Select **Save**.

Snyk checks the credentials, and upon success, the page reloads with a notification that the connection succeeded.
