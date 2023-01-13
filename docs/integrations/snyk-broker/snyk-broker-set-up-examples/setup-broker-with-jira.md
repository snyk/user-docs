# Set up Broker with Jira

Configuring Jira with Broker is useful to ensure a secure connection with your on-premise Jira deployment.

{% hint style="info" %}
Docker is a prerequisite.
{% endhint %}

## Configure Broker to be used for Jira

1. Click **Settings** > **Integrations** > **Jira** > **For installation of Jira within a private network click here**.
2. Click **Generate** to generate a Broker Token for Jira and click **Show** to confirm.
3. Visit [the Broker repository](https://github.com/snyk/broker) and scroll down to the Jira section..
4. Run `docker pull snyk/broker:jira` to pull down the latest Jira Broker image from Dockerhub
5.  Copy the command line arguments under the Jira section and modify the values. Note that the BROKER\_TOKEN is the token generated in step 2.

    Example:\
    `docker run --restart=always`\
    `-p <port>:<port> \`\
    `-e BROKER_TOKEN=<broker token from Snyk UI> \`\
    `-e JIRA_USERNAME=<jira username> \`\
    `-e JIRA_PASSWORD=<jira password> \`\
    `-e JIRA_HOSTNAME=<Jira URL (no https)> \`\
    `-e BROKER_CLIENT_URL=http://<broker IP/DNS>:<port> \`\
    `-e PORT=<port> \`\
    `snyk/broker:jira`
6. If necessary, go to the Advanced Configuration section of [the Broker repository](https://github.com/snyk/broker) and make any configuration changes needed.\
   For example, if the Jira instance is using a private certificate, provide the CA (Certificate Authority) to the Broker Client configuration.
7. Paste the Broker Client configuration to start the Broker client container.
8. Once the container is up, the Jira Integrations page shows the connection to Jira and under Projects you can create Jira tickets

## **Basic troubleshooting for Broker with Jira**

* Run `docker logs <container id>` where container id is the Jira Broker container ID to look for any errors.
* Ensure relevant ports are exposed to Jira.
