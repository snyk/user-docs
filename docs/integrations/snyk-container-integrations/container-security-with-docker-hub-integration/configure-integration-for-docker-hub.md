# Configure integration for Docker Hub

This page explains how to enable and configure integration between Docker Hub and Sny, When integration is complete, you can start managing your vulnerabilities.

## Enable integration with Docker Hub

1. Navigate to **Integrations**.
2. Click **Docker Hub**.
3. Enter your Docker Hub username and Access Token; for details, see [Generate Docker Hub Access Token](configure-integration-for-docker-hub.md#generate-docker-hub-access-token).
4. Click **Save**.\
   The page reloads with new options. The Access Token field is blank.\
   A confirmation message that the details were saved also appears in green at the top of the screen.

<figure><img src="../../../.gitbook/assets/Screen Shot 2022-01-21 at 9.48.27 AM.png" alt="Successful connection with Docker Hub"><figcaption><p>Successful connection with Docker Hub</p></figcaption></figure>

If the connection to Docker Hub fails, an error notification appears:

<figure><img src="../../../.gitbook/assets/Screen Shot 2022-01-21 at 9.48.50 AM.png" alt="Failed connection, Could not connect to Docker Hub"><figcaption><p>Failed connection, Could not connect to Docker Hub</p></figcaption></figure>

## Troubleshooting Docker Hub integration

If issues occur such as failure to import Projects, failure to connect, and no error shown, first try generating a new Access Token and resaving the Docker Hub integration on the Snyk settings page.

## Generate Docker Hub Access Token

1. Navigate to [https://hub.docker.com/settings/security](https://hub.docker.com/settings/security)
2. Select **New Access Token**.
3. Enter the Access Token description.
4. Set the permissions (Read is required and sufficient), and click **Generate**.
5. Select **Copy Access Token** for use with the username when you [enable integration with Docker Hub](configure-integration-for-docker-hub.md#enable-integration-with-docker-hub).

More information on [Docker Hub Access Tokens](https://docs.docker.com/docker-hub/access-tokens/) is available in the Docker Hub docs.
