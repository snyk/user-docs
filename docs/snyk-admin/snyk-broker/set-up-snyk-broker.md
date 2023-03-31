# Set up Snyk Broker

The Broker Client is installed and configured as a Docker image.

{% hint style="info" %}
You can customize the configuration using the environment variables in the Docker images. For this reason, install separate, multiple instances of the Broker Client for different integration types to ensure proper configuration as well as redundancy.
{% endhint %}

{% hint style="warning" %}
**Multi-tenant settings for EU and AU**\
When you set up Broker, Code Agent, or both for use in EU or AU Multi-tenant environments, additional environment variables with the specific URLs are required.\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`\
For the URLs, see [EU and AU account datacenter creation](https://docs.snyk.io/snyk-processes/data-residency-at-snyk#eu-and-au-datacenter-account-creation).
{% endhint %}

To install, configure, and roll out your Broker Client and repository integrations:

1. [Prepare Snyk Broker for Deployment](prepare-snyk-broker-for-deployment.md)
2. [Install and configure Snyk Broker Client](how-to-install-and-configure-your-snyk-broker-client/)

For more information see [Troubleshooting Snyk Broker](troubleshooting-broker.md).
