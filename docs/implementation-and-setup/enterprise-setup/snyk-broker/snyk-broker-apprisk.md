# Using Snyk AppRisk with Snyk Broker

{% hint style="info" %}
**Feature availability**

Snyk AppRisk and the third-party integrations associated with it are available only for Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Snyk AppRisk has many third-party integrations with security and development tools. If these tools are hosted within your private network and are not directly accessible from the internet, you must use Snyk Broker to establish a connection.

## Integrate AppRisk with Broker

The specific information you need to integrate Snyk AppRisk with Snyk Broker depends on the type of Snyk Broker deployment you are using: Classic Broker or Universal Broker.

### Prerequisite for using Snyk AppRisk with Classic Broker

* Required: Your Snyk Broker toke&#x6E;**.**
* How to obtain it:
  * Contact the Snyk Support team.
  * Alternatively, generate it yourself by following the instructions to [Obtain your Broker token from the Web UI](../../../enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md#obtain-your-broker-token-from-the-web-ui).

### Prerequisite for using Snyk AppRisk with Universal Broker

* Required: The identifier for your connection specific to your Universal Broker setup.
* How to obtain it: You can find this identifier by referring to the [Validate your deployment (optional)](../../../enterprise-setup/snyk-broker/universal-broker/basic-steps-to-install-and-configure-universal-broker.md#validate-your-connection-optional) section within the [Basic steps to install and configure Universal Broker](../../../enterprise-setup/snyk-broker/universal-broker/basic-steps-to-install-and-configure-universal-broker.md) documentation.

### Final step in integration for both Broker types

After you have your `Snyk Broker token` for Classic Broker or your connection `identifier` for Universal Broker:

1. Copy the token or identifier value.
2. Navigate to the Snyk AppRisk integration settings within the Group-level Integration Hub.
3. Paste the value into the **Broker token** field to complete the integration.

## Integrate a third-party tool with Snyk Broker

To integrate a third-party tool with Universal Broker, refer to [Setting up and integrating your Universal Broker connections](universal-broker/setting-up-and-integrating-your-universal-broker-connections.md).

### Configure third-party integrations with the integration password

After implementing all the general steps applicable to [third-party integrations](../../../integrations/connect-a-third-party-integration.md), you can configure the integration with unique credentials, as shown in this example. Run the following commands using your password:

```docker
docker run --restart=always \
        -p 8001:8001 -e PORT=8001 \
        -e BROKER_CLIENT_URL=http://broker.url.example:8000 \
        -e BROKER_TOKEN=<YOUR BROKER TOKEN> \
        -e UNIVERSAL_BROKER_ENABLED=true \
        -e INTEGRATION_PASSWORD=<YOUR INTEGRATION PASSWORD> \
        -e BROKER_SERVER_URL=https://broker.snyk.io \
        -v $(pwd)/config.universal.json:/home/node/config.universal.json \
    snyk/broker:universal
```

### Configure third-party integrations with the API token

After you implement all the general steps applicable to [third-party integrations](../../../integrations/connect-a-third-party-integration.md), you can configure the integration with unique credentials, such as `INTEGRATION_HOST_URL` and `INTEGRATION_API_TOKEN` as the values for the credentials reference. Run the following commands with your host URL and Broker token:

```docker
docker run --restart=always \
-p 8001:8001 -e PORT=8001 \
-e BROKER_CLIENT_URL=http://broker.url.example:8000 \
-e BROKER_TOKEN=<YOUR BROKER TOKEN> \
-e UNIVERSAL_BROKER_ENABLED=true \
-e INTEGRATION_HOST_URL=<YOUR HOST URL> \
-e INTEGRATION_API_TOKEN=<YOUR API TOKEN> \
-e BROKER_SERVER_URL=https://broker.snyk.io \
-v $(pwd)/config.universal.json:/home/node/config.universal.json \
snyk/broker:universal
```

### Configuration complete

After the Snyk Broker connection with a third-party integration is established, the following message is displayed in the logs: `successfully established a websocket connection to the broker server`.

{% code overflow="wrap" %}
```docker
{"id":"broker-client-url-validation","name":"Broker Client URL Validation Check","status":"passing","output":"config check: ok"},{"id":"universal-broker-connections-config-validation","name":"Universal Broker Client Connections Configuration Check","status":"passing","output":"connections config check: ok"}],"version":"4.179.5","supportedIntegrationType":"apprisk"},"msg":"successfully established a websocket connection to the broker server","time":"2024-03-11T11:43:26.014Z","v":0}
```
{% endcode %}
