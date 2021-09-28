# Configure Docker Hub intergration

##  Configure integration for Docker Hub

Enable integration between Docker Hub and Snyk, and start managing your vulnerabilities.

1. From the Snyk UI, log in to your account and navigate to Integrations from the menu bar at the top.
2. From the Integrations page click the Docker Hub box \(image 1\)
3. From the **Settings** page in the Integrations area, enter your Docker Hub username and Access Token and then click **Save** \(image 2\)

![image 1](https://support.snyk.io/hc/article_attachments/360007147218/uuid-fcf2376d-3886-a580-afa3-bbd999819bdc-en.png)

![image 2](https://support.snyk.io/hc/article_attachments/360007818037/mceclip0.png)

 _As the access token, you can either use your DockerHub password or an access token_ [_created_](https://docs.docker.com/docker-hub/access-tokens/) _in DockerHub. In case 2FA is activated on your account, access token **only** is applicable._

Snyk tests the connection values and the page reloads, now displaying Docker Hub integration information and the Add your Docker Hub images to the Snyk button.

A confirmation message that the details were saved also appears in green at the top of the screen. In addition, if the connection to Docker Hub failed, a notification appears.

