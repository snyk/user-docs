# Jira - install and configure using Docker

This integration is useful to ensure a secure connection with your on-premise Jira deployment.

## Configure Broker to be used for Jira

To use the Broker Client with a Jira deployment, run `docker pull snyk/broker:jira`. For definitions of the environment variables, see [Jira - environment variables for Snyk Broker](jira-environment-variables-for-snyk-broker.md).

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

If necessary, navigate to [Advanced configuration for Snyk Broker Docker installation](../advanced-configuration-for-snyk-broker-docker-installation/) and make any configuration changes needed, for example, providing the CA (Certificate Authority) to the Broker Client configuration when the Jira instance is using a private certificate.

As an alternative to using the Docker run command, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the Jira integration.

## Jira PAT authentication for SSO-enabled JIRA

When SSO is enabled, JIRA usually prohibits the use of a username and password and requires the use of a personal access token (PAT).

When SSO is enabled, you must use a specific Jira version that will instead use the authorization header with the bearer token. To use this version, provide the following configuration:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e JIRA_PAT=<your_pat_token> \
           -e JIRA_HOSTNAME=your.jira.domain.com \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
           -e PORT=8000 \
       snyk/broker:jira-bearer-auth
```

## Start the Broker Client container and verify the connection with Jira

&#x20;Paste the Broker Client configuration to start the Broker Client container.

After the container is set up, and the Jira Integrations page shows the connection to Jira, under Projects you can create Jira tickets

## **Basic troubleshooting for Broker with Jira**

* Run `docker logs <container id>` to look for any errors, where container id is the Jira Broker container ID.
* Ensure relevant ports are exposed to Jira.
