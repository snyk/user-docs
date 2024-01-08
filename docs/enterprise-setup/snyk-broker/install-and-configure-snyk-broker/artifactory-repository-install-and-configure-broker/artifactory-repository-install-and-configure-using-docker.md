# Artifactory Repository - install and configure using Docker

Follow the instructions on this page to set up Artifactory Repository with Snyk Broker. This integration is useful to ensure a secure connection with your on-premise Artifactory Repository  deployment.

{% hint style="info" %}
**Feature availability**\
Integration with Artifactory Repository is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="info" %}
**Prerequisites**\
Ask your Snyk account team to provide you with a Broker token or generate it from the Snyk Web UI.

You need Docker or a way to run Docker Linux containers.

Some Docker deployments for Windows run only Windows containers. Ensure that your deployment is capable of running Linux containers.
{% endhint %}

{% hint style="info" %}
For information about non-brokered integration with Artifactory Repository see [Artifactory Repository setup](../../../../integrate-with-snyk/package-repository-integrations/artifactory-package-repository-connection-setup/).
{% endhint %}

{% hint style="info" %}
For information about brokered integration with Artifactory Container Registry see [Snyk Broker -Container Registry Agent](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent).
{% endhint %}

## Obtain Broker token for Artifactory Repository setup

1. Go to settings <img src="../../../../.gitbook/assets/cog_icon.png" alt="" data-size="line"> > **Integrations > Package Repositories > Artifactory**.
2. Enter the URL of your Artifactory instance, this **must** end with **/artifactory**.
3. Enter your username and password.
4. Select **Save**.

<figure><img src="../../../../.gitbook/assets/screenshot_2020-04-17_at_14.38.12.png" alt="Artifactory integration setup"><figcaption><p>Artifactoryrepository setup</p></figcaption></figure>

{% hint style="info" %}
If you do not see the **Snyk Broker** on/off switch, you do not have the necessary permissions and can only add a publicly accessible instance.

Submit a request to [Snyk Support](https://support.snyk.io/hc/en-us/requests/new) if you want to add a private registry.
{% endhint %}

When you have permissions to add a private registry, continue with the instructions to [generate a Broker token from the Web UI](artifactory-repository-install-and-configure-using-docker.md#generate-a-broker-token-from-the-web-ui).

## Generate a Broker token from the Web UI

1. In the Artifactory integration settings, move the **Snyk Broker on/off** switch to **on** to display a form for generating a Broker token.
2. Select **Generate and Save.**
3. Copy the token that was generated to use when you set up the Broker Client.

## Configure Broker to be used for Artifactory Registry

To use the Broker client with an Artifactory Registry deployment, **run** `docker pull snyk/broker:artifactory`. Refer to [Artifactory Repository - environment variables for Snyk Broker](artifactory-repository-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

## Docker run commands to set up a Broker Client for Artifactory Repository

**Copy the following command** to set up a fully configured Broker Client to use with Artifactory Registry. You can run the Docker container by providing the relevant configuration:

```console
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e ARTIFACTORY_URL=<yourdomain>.artifactory.com/artifactory \
       snyk/broker:artifactory
```

For an **npm or Yarn integration**, use the following **command**.

```
docker run --restart=always \
  -p 8000:8000 \
  -e BROKER_TOKEN=secret-broker-token \
  -e ARTIFACTORY_URL=acme.com/artifactory \
  -e RES_BODY_URL_SUB=http://acme.com/artifactory \ 
  snyk/broker:artifactory
```

As an **alternative to using the Docker run command**, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the Artifactory integration.

## Start the Broker Client container and verify the connection with Artifactory Repository

Paste the Broker Client configuration to start the Broker Client container.

You can check the status of the connection by refreshing the Artifactory Integration Settings page. When the connection is set up correctly, there is no connection error.
