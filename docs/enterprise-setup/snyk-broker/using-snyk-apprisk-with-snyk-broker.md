# Snyk Broker - AppRisk

## Third-party integrations

{% hint style="info" %}
**Feature availability**

The third-party integrations are available only for Snyk AppRisk, with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

## Prerequisites

Follow these steps to install and run Snyk Broker for the Snyk AppRisk third-party integrations.

1. Ensure you have the Snyk Broker token for the Snyk AppRisk integration. The Snyk support team can provide the needed token.&#x20;
   * Generate your Broker token by following the instructions from the [Obtain your Broker token for Snyk Broker ](classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md#obtain-your-broker-token-from-the-web-ui)page.&#x20;
   * Copy and paste the Broker token on the integration setup menu from the Integration Hub.
2. Pull the latest Broker image by running this command:

```docker
docker pull snyk/broker:universal
```

3. Configure your Snyk AppRisk connection type using the `snyk-broker-config` command, as explained on the page [Basic steps to install and configure Universal Broker](universal-broker/basic-steps-to-install-and-configure-universal-broker.md).

## Checkmarx SAST integration

After you implement all the general steps applicable to third-party integrations, you can configure the integration with unique credentials.&#x20;

The following example has `CHECKMARX_PASSWORD` as the value for the credentials reference. Run the following commands with your password:

```docker
docker run --restart=always \
        -p 8001:8001 -e PORT=8001 \
        -e BROKER_CLIENT_URL=http://broker.url.example:8000 \
        -e BROKER_TOKEN=<YOUR BROKER TOKEN> \
        -e UNIVERSAL_BROKER_ENABLED=true \
        -e CHECKMARX_PASSWORD=<YOUR CHECKMARX PASSWORD> \
        -e BROKER_SERVER_URL=https://broker.snyk.io \
        -v $(pwd)/config.universal.json:/home/node/config.universal.json \
    snyk/broker:universal
```

## SonarQube SAST integration

After you implement all the general steps applicable to third-party integrations, you can configure the integration with unique credentials.&#x20;

The following example uses `SONARQUBE_HOST_URL` and `SONARQUBE API_TOKEN` as the values for the credentials reference. Run the following commands:

```docker
docker run --restart=always \
-p 8001:8001 -e PORT=8001 \
-e BROKER_CLIENT_URL=http://broker.url.example:8000 \
-e BROKER_TOKEN=<YOUR BROKER TOKEN> \
-e UNIVERSAL_BROKER_ENABLED=true \
-e SONARQUBE_HOST_URL=<YOUR HOST URL> \
-e SONARQUBE_API_TOKEN=<YOUR API TOKEN> \
-e BROKER_SERVER_URL=https://broker.snyk.io \
-v $(pwd)/config.universal.json:/home/node/config.universal.json \
snyk/broker:universal
```

## Configuration complete

After the Universal Broker connection with a third-party integration is established, the following message is displayed in the logs: `successfully established a websocket connection to the broker server`.

{% code overflow="wrap" %}
```docker
{"id":"broker-client-url-validation","name":"Broker Client URL Validation Check","status":"passing","output":"config check: ok"},{"id":"universal-broker-connections-config-validation","name":"Universal Broker Client Connections Configuration Check","status":"passing","output":"connections config check: ok"}],"version":"4.179.5","supportedIntegrationType":"apprisk"},"msg":"successfully established a websocket connection to the broker server","time":"2024-03-11T11:43:26.014Z","v":0}
```
{% endcode %}

