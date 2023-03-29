# Set up Snyk Broker with Jira

Follow the instructions on this page to set up Jira with Snyk Broker. This integration is useful to ensure a secure connection with your on-premise Jira deployment.

{% hint style="info" %}
Docker is a prerequisite.
{% endhint %}

## Configure Broker to be used for Jira

Follow these steps to configure Snyk Broker for connection to Jira.

1. Click **Settings** > **Integrations** > **Jira** > **For installation of Jira within a private network click here**.
2. Click **Generate** to generate a Broker Token for Jira and click **Show** to confirm.
3. To pull down the latest Jira Broker image from Dockerhub, run `docker pull snyk/broker:jira`. The following environment variables are required to configure the Broker client:
   * `BROKER_TOKEN` - the Snyk Broker token, obtained from your Jira integration settings view.
   * `JIRA_USERNAME` - the Jira username.
   * `JIRA_PASSWORD` - the Jira password.
   * `JIRA_HOSTNAME` - the hostname of your Jira deployment, such as `your.jira.domain.com`.
   * `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible by your Jira for webhooks, such as `http://my.broker.client:8000`
   * `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
4.  Copy the following command line arguments:

    ```
    docker run --restart=always \
               -p 8000:8000 \
               -e BROKER_TOKEN=secret-broker-token \
               -e JIRA_USERNAME=username \
               -e JIRA_PASSWORD=password \
               -e JIRA_HOSTNAME=your.jira.domain.com (no http/s) \
               -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
               -e PORT=8000 \
           snyk/broker:jira
    ```
5. If necessary, go to the Advanced Configuration section of [Install and configure the Snyk Broker client](../set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client/) and make any configuration changes needed, for example, providing the CA (Certificate Authority) to the Broker Client configuration when the Jira instance is using a private certificate.
6. Paste the Broker Client configuration to start the Broker Client container.
7. Once the container is up, and the Jira Integrations page shows the connection to Jira, under Projects, create Jira tickets

{% hint style="info" %}
As an alternative to using the Docker run command, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the Jira integration.
{% endhint %}

## **Basic troubleshooting for Broker with Jira**

* Run `docker logs <container id>` to look for any errors, where container id is the Jira Broker container ID.
* Ensure relevant ports are exposed to Jira.
