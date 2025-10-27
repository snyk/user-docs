# Using Snyk Essentials with Snyk Broker

If your SCM instance is not publicly accessible, you must use Snyk Broker to scan your repositories.

For details about installing an SCM integration for Universal Broker, see [Basic steps to install and configure Universal Broker](../../../enterprise-setup/snyk-broker/universal-broker/basic-steps-to-install-and-configure-universal-broker.md).

You can install and configure an SCM integration for Classic Broker using Docker or Helm. For details, see the [Install and configure Snyk Broker](classic-broker/install-and-configure-snyk-broker/) section.

To use the features of Snyk Essentials with repositories that are not publicly accessible and are using a brokered connection, you must integrate Snyk Essentials with Snyk Broker. The minimum supported Broker version for Snyk Essentials is [4.171.0](https://github.com/snyk/broker/releases/tag/v4.171.0).

To integrate Snyk Essentials with the Classic Snyk Broker, you need your Snyk Broker token. The Snyk support team can provide the needed token, or you can generate it yourself by following the instructions to [Obtain your Broker token from the Web UI](../../../enterprise-setup/snyk-broker/classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md#obtain-your-broker-token-from-the-web-ui).

To integrate Snyk Essentials with the Universal Broker, you need the identifier for your connection. You can see how to retrieve this in the [Basic steps to install and configure Universal Broker, Validate your deployment (optional)](../../../enterprise-setup/snyk-broker/universal-broker/basic-steps-to-install-and-configure-universal-broker.md#validate-your-connection-optional).

When you have generated your Broker token or found the identifier for your connection, copy the value and paste it into the Broker token field for your Snyk Essentials integration.
