# Setup Broker with Jira

Configuring Jira with broker is useful to ensure a secure connection with your on-premise Jira deployment.

#### Pre Requisites: Docker

### To configure a Broker to be used for Jira

1. Click on settings <img src="../../../.gitbook/assets/cog_icon.png" alt="cog_icon.png" data-size="line"> > **Integrations** > Jira > "For installation of Jira within a private network click here"
2. Click "Generate" to generate a Broker Token for Jira and click "Show" to confirm
3. Visit [the broker repository](https://github.com/snyk/broker) and scroll down to the Jira section
4. Run `docker pull snyk/broker:jira` to pull down the latest Jira broker image from Dockerhub
5.  Copy the Command-line arguments under the Jira section and modify the values. (Note: BROKER\_TOKEN is the token generated in step 2)

    Example:

    `docker run --restart=always`\
    `-p <port>:<port> \`\
    `-e BROKER_TOKEN=<broker token from Snyk UI> \`\
    `-e JIRA_USERNAME=<jira username> \`\
    `-e JIRA_PASSWORD=<jira password> \`\
    `-e JIRA_HOSTNAME=<Jira URL (no https)> \`\
    `-e BROKER_CLIENT_URL=http://<broker IP/DNS>:<port> \`\
    `-e PORT=<port> \`\
    `snyk/broker:jira`
6. If necessary, go to the Advanced Configuration section of [the broker repository](https://github.com/snyk/broker) and make any configuration changes if required (For example, if the Jira instance is using a private certificate, provide the CA (Certificate Authority) to the Broker Client configuration)
7. Paste the Broker Client configuration to start the broker client container
8. Once the container is up, the Jira Integrations page should show the connection to Jira and under Projects, you should be able to create Jira tickets

**Basic Troubleshooting**

* Run `docker logs <container id>` where container id is the Jira Broker container ID to look for any errors
* Ensure relevant ports are exposed to Jira
