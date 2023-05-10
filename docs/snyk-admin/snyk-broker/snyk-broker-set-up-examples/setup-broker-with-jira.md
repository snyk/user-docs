# Set up Snyk Broker with Jira

Follow the instructions on this page to set up Jira with Snyk Broker. This integration is useful to ensure a secure connection with your on-premise Jira deployment.

{% hint style="info" %}
Docker is a prerequisite.
{% endhint %}

## Generate a Broker token from the Web UI

1. Click **Settings** > **Integrations** > **Jira** > **For installation of Jira within a private network click here**.
2. Click **Generate** to generate a Broker Token for Jira and click **Show** to confirm.

## Configure Broker to be used for Jira

To use the Broker client with a Jira deployment, run `docker pull snyk/broker:jira`.

The following environment variables are needed to configure the Broker Client:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your Jira integration settings view.
* `JIRA_USERNAME` - the Jira username.
* `JIRA_PASSWORD` - the Jira password.
* `JIRA_HOSTNAME` - the hostname of your Jira deployment, such as `your.jira.domain.com`.
* `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible by your Jira for webhooks, such as `http://my.broker.client:8000`
* `PORT` - the local port at which the Broker client accepts connections. Default is 8000.

## Docker run command to set up a Broker Client for Jira

**Copy the following command** to set up a fully configured Broker Client to use with Jira. You can run the Docker container by providing the relevant configuration:

```console
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e JIRA_USERNAME=username \
           -e JIRA_PASSWORD=password \
           -e JIRA_HOSTNAME=your.jira.domain.com \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
           -e PORT=8000 \
       snyk/broker:jira
```

If necessary, go to the Advanced Configuration section of [Install and configure the Snyk Broker client](../upgrade-the-snyk-broker-client.md) and make any configuration changes needed, for example, providing the CA (Certificate Authority) to the Broker Client configuration when the Jira instance is using a private certificate.

As an alternative to using the Docker run command, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the Jira integration.

## Start the Broker Client container and verify the connection with Jira

&#x20;Paste the Broker Client configuration to start the Broker Client container.

Once the container is up, and the Jira Integrations page shows the connection to Jira, under Projects you can create Jira tickets

## **Basic troubleshooting for Broker with Jira**

* Run `docker logs <container id>` to look for any errors, where container id is the Jira Broker container ID.
* Ensure relevant ports are exposed to Jira.
