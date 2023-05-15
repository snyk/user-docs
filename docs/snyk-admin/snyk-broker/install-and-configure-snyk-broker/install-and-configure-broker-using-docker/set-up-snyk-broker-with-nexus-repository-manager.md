# Nexus Repository - install and configure using Docker

Follow the instructions on this page to set up Nexus Repository Manager with Snyk Broker. This integration is useful to ensure a secure connection with your on-premise Nexus Respository Manager deployment.

{% hint style="info" %}
**Feature availability**\
Integration with Nexus Repository Manager is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="info" %}
**Prerequisites**\
Ask your Snyk account team to provide you with a Broker token or generate it from the Snyk Web UI.

You need Docker or a way to run Docker Linux containers.\
Some Docker deployments for Windows run only Windows containers. Ensure that your deployment is capable of running Linux containers.
{% endhint %}

{% hint style="info" %}
For information about non-brokered integration with Nexus Repository Manager including supported environments and versions and user permissions, see [Nexus Repository Manager setup](../../../../integrations/private-registry-integrations/nexus-repo-manager-setup.md).
{% endhint %}

{% hint style="info" %}
For information about brokered integration with Nexus Container Registry see [Snyk Broker -Container Registry Agent](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent).
{% endhint %}

## Obtain Broker token for Nexus integration

1. Go to settings <img src="../../../../.gitbook/assets/cog_icon.png" alt="Settings icon" data-size="line"> > **Integrations > Package Repositories > Nexus**
2. Verify that you see the screen to configure Nexus.

<figure><img src="../../../../.gitbook/assets/Screenshot 2022-07-15 at 15.15.11.png" alt="Configure Nexus"><figcaption><p>Configure Nexus</p></figcaption></figure>

{% hint style="info" %}
If you do not see the **Snyk Broker** switch, you do not have the necessary permissions and can only add a publicly accessible instance.

Submit a request to [Snyk Support](https://support.snyk.io/hc/en-us/requests/new) if you want to add a private registry.
{% endhint %}

When you have permissions to add a private registry, continue with the instructions to [generate a Broker token from the Web UI](set-up-snyk-broker-with-nexus-repository-manager.md#generate-a-broker-token-from-the-web-ui).

## Generate a Broker token from the Web UI

1. In the Nexus integration settings, move the **Snyk Broker on/off** switch to **on** to display a form for generating a Broker token.
2. Select **Generate and Save.**
3. Copy the token that was generated to use when you set up the Broker Client.

## Configure Broker to be used for Nexus plugins

### Docker pull for Nexus 3 and Nexus 2 configuration

To use the Broker client with a Nexus 3 deployment, **run** `docker pull snyk/broker:nexus`.

To use the Broker client with a Nexus 2 deployment, **run** `docker pull snyk/broker:nexus2`.

### Environment variables for Nexus 3 configuration

The following environment variables are needed to customize the Broker client for Nexus 3:

`BROKER_TOKEN`\
The Snyk Broker token, obtained from your Nexus integration settings (**Integrations > Nexus**).

`BASE_NEXUS_URL`\
The URL of your Nexus 3 deployment.\
Example:\
`BASE_NEXUS_URL=https://[<username_or_token><password_or_token>]@<your.nexus.hostname>`\
Must not end with a forward slash.\
The following field is optional:\
`Auth`: Omit if no auth required.\
Can either be plain text or a two-part token (`Nexus Pro`).\
URL encode username, password, and tokens to avoid errors that may prevent authentication.\
Minimal example: `acme.com`\
Complex example: `https://alice:mypassword@acme.com`

`BROKER_CLIENT_VALIDATION_URL`\
Nexus validation url, checked by Broker Client `systemcheck` endpoint.\
If Nexus user requires `auth`, use `$BASE_NEXUS_URL/service/rest/v1/status/check`\
Example:\
`https://<user>:<pass>@<your.nexus.hostname>/service/rest/v1/status/check`)\
Otherwise use `$BASE_NEXUS_URL/service/rest/v1/status`\
Example:\
`https://<your.nexus.hostname>/service/rest/v1/status`).

Optional. `RES_BODY_URL_SUB`\
This URL substitution is **required for npm/Yarn integration** and is the same as the URL of the Nexus without credentials appended with `/repository`\
Example:\
`https://<your.nexus.hostname>/repository`. Must not end with a forward slash.

### Environment variables for Nexus 2 configuration

The following environment variables are needed to customize the Broker client for Nexus 2:

`BROKER_TOKEN` - the Snyk Broker token, obtained from your Nexus integration settings (**Integrations > Nexus**).

`BASE_NEXUS_URL`- the URL of your Nexus 2 deployment.\
Example:\
`BASE_NEXUS_URL=https://[username_or_token:password_or_token]@acme.com`\
Must not end with a forward slash.\
The following field is optional:\
`Auth`: Omit if no auth required.\
Can be either plain text or a two-part token (`Nexus Pro`).\
URL encode username, password, and tokens to avoid errors that may prevent authentication.\
Minimal example: `https://acme.com`\
Complex example: `https://alice:mypassword@acme.com:`

`RES_BODY_URL_SUB`\
The URL of the Nexus instance, including `https://` and `/nexus/content` without basic auth credentials. **Required for npm/Yarn integrations only.** Must not end with a forward slash.

Examples:\
`https://acme.com/nexus/content/groups`\
`https://acme.com/nexus/content/repositories`

### Docker run commands to set up Broker Client for Nexus 3 and Nexus 2 integrations

**Copy the following command** to set up a fully configured Broker Client to use with Nexus 3. You can run the Docker container by providing the relevant configuration:

```console
docker run --restart=always \
           -p 7341:7341 \
           -e BROKER_TOKEN=secret-broker-token \
           -e BASE_NEXUS_URL=https://[<user>:<pass>@]<your.nexus.hostname> \
           -e BROKER_CLIENT_VALIDATION_URL=https://<your.nexus.hostname>/service/rest/v1/status[/check] \
           -e RES_BODY_URL_SUB=https://<your.nexus.hostname>/repository \
       snyk/broker:nexus
```

**Copy the following command** to set up a fully configured Broker Client to use with Nexus 2. You can run the Docker container by providing the relevant configuration:

```
docker run --restart=always \
  -p 7341:7341 \
  -e BROKER_TOKEN=<secret-broker-token> \
  -e BASE_NEXUS_URL=https://[username:password]@acme.com \
  -e RES_BODY_URL_SUB=https://acme.com/nexus/content/(groups|repositories) \ 
  snyk/broker:nexus2
```

As an **alternative to using the Docker run command**, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the Nexus3 integration.

## Start the Broker Client container and verify the connection with Nexus Repository Manager

Paste the Broker Client configuration to start the Broker Client container.

Check connection status by making a request to the Broker Client `/systemcheck` endpoint.

Example: `curl http://172.17.0.2:7341/systemcheck`

You see success output in the following form:

`{"brokerClientValidationUrl":"https://acme.com/service/rest/v1/status","brokerClientValidationMethod":"GET","brokerClientValidationTimeoutMs":5000,"brokerClientValidationUrlStatusCode":200,"ok":true}`

Or failure output in the following form:

`{"brokerClientValidationUrl":"https://acme.com/service/rest/v1/status","brokerClientValidationMethod":"GET","brokerClientValidationTimeoutMs":5000,"ok":false,"error":"ETIMEDOUT"}`
