# Configure integration for Docker Hub

Enable integration between Docker Hub and Snyk, and start managing your vulnerabilities.

1. Navigate to **Integrations**
2. Click **Docker Hub**
3. Enter your Docker Hub username and Access Token (see below to generate token)
4. click **Save**, the page will reload with new options and the Access Token field will be blank--this is normal
5. A confirmation message that the details were saved also appears in green at the top of the screen.

![This is an example of a successful connection.](<../../../../.gitbook/assets/Screen Shot 2022-01-21 at 9.48.27 AM.png>)

#### Connection failure

If the connection to Docker Hub failed, an error notification appears:

![This is an example of an unsuccessful connect: "Could not connect to Docker Hub".](<../../../../.gitbook/assets/Screen Shot 2022-01-21 at 9.48.50 AM.png>)

### Troubleshooting Docker Hub integration

The first line of troubleshooting Docker Hub integration issues--failure to import projects, fail to connect, no error shown, etc.--is to generate a new Access Token and resave the Docker Hub integration in the Snyk settings page.

#### To generate Docker Hub Access Token:

1. Navigate to [https://hub.docker.com/settings/security](https://hub.docker.com/settings/security)&#x20;
2. Select New Access Token&#x20;
3. Enter access token description, set permissions, and click Generate&#x20;
4. Copy Access Token, it will be used along with the username in step 1 above

More information on [Docker Hub Access Tokens](https://docs.docker.com/docker-hub/access-tokens/) is available in the Docker Hub docs.
