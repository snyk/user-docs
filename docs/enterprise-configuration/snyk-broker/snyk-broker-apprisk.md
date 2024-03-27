# Snyk Broker - AppRisk

If your SCM or third-party instance is not publicly accessible, you need Snyk Broker. You can install and configure your Snyk Broker using Docker or Helm. The minimum supported Broker version for Snyk AppRisk is [4.171.0](https://github.com/snyk/broker/releases/tag/v4.171.0).

{% hint style="info" %}
Enable the Snyk AppRisk flag in your Snyk Broker deployment environment before running the commands.
{% endhint %}

## SCM integrations

* GitHub - install and configure Snyk Broker&#x20;
  * [using Docker](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/broker-example-set-up-snyk-broker-with-github.md#docker-run-command-to-set-up-a-broker-client-for-github)
  * [using Helm](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/githhub.com-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-environment-variables-for-snyk-broker.md)
* GitHub Enterprise - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-github-enterprise.md#docker-run-command-to-set-up-a-broker-client-for-github-enterprise)
  * [using Helm](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-environment-variables-for-snyk-broker.md)
* BitBucket - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md#docker-run-command-to-set-up-a-broker-client-for-bitbucket)
  * [using Helm](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker.md)
* GitLab - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-gitlab.md#docker-run-command-to-set-up-a-broker-client-for-gitlab)
  * [using Helm](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-install-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-environment-variables-for-snyk-broker.md)
* Azure - install and configure Snyk Broker:
  * [using Docker](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md#docker-run-command-to-set-up-a-broker-client-for-azure-repos)
  * [using Helm](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-install-and-configure-and-configure-using-helm.md)
  * [environment variables](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-environment-variables-for-snyk-broker.md)

You can find on [GitHub](https://github.com/snyk/broker/tree/565242baf003f06f445489dd96cc68c8386ede38/defaultFilters/apprisk) all the updated `.json` files that include the allowed list of accessible endpoints for the integrations.

## Third-party integrations

{% hint style="warning" %}
The third-party integrations are available in a Closed Beta state and are applicable only to the Snyk AppRisk Pro version.  Please contact your salesperson if you are interested in Snyk AppRisk Pro.
{% endhint %}

### Checkmarx SAST integration

Use the following steps to install and run Snyk Broker for the AppRisk and Checkmarx SAST integration.

1. Ensure you have the Snyk Broker token for the Snyk AppRisk integration. The Snyk support team can provide the needed token.&#x20;
2. Pull the latest broker image by running this command:

```docker
docker pull snyk/broker:universal
```

3. Ensure the `config.universal.json` file contains the following information:

```
{
  "BROKER_CLIENT_CONFIGURATION": {
    "common": {
      "default": {
        "BROKER_SERVER_URL": "https://broker.snyk.io",
        "BROKER_HA_MODE_ENABLED": "false"
      }
    }
  },
  "CONNECTIONS": {
    "apprisk connection": {
      "type": "apprisk",
      "identifier": "${BROKER_TOKEN}",
      "CHECKMARX": "${CHECKMARX}",
      "CHECKMARX_USERNAME": "${CHECKMARX_USERNAME}",
      "CHECKMARX_PASSWORD": "${CHECKMARX_PASSWORD}",
      "BROKER_CLIENT_URL": "http://my.broker.client.dns.hostname"
    }
  }
}

```

4. Run the following commands with your Checkmarx username and password:

```docker
docker run --restart=always \
        -p 8001:8001 -e PORT=8001 \
        -e BROKER_CLIENT_URL=http://broker.url.example:8000 \
        -e BROKER_TOKEN=<YOUR BROKER TOKEN> \
        -e UNIVERSAL_BROKER_ENABLED=true \
        -e CHECKMARX=<YOUR CHECKMARX HOST> \
        -e CHECKMARX_USERNAME=<YOUR CHECKMARX USERNAME> \
        -e CHECKMARX_PASSWORD=<YOUR CHECKMARX PASSWORD> \
        -e BROKER_SERVER_URL=https://broker.snyk.io \
        -v $(pwd)/config.universal.json:/home/node/config.universal.json \
    snyk/broker:universal

```

5. When the connection is established, you will find in the logs the following message: `successfully established a websocket connection to the broker server`

{% code overflow="wrap" %}
```docker
{"id":"broker-client-url-validation","name":"Broker Client URL Validation Check","status":"passing","output":"config check: ok"},{"id":"universal-broker-connections-config-validation","name":"Universal Broker Client Connections Configuration Check","status":"passing","output":"connections config check: ok"}],"version":"4.179.5","supportedIntegrationType":"apprisk"},"msg":"successfully established a websocket connection to the broker server","time":"2024-03-11T11:43:26.014Z","v":0}

```
{% endcode %}
