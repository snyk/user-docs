# Install and configure Snyk Broker

## How to use Snyk Broker

Snyk Broker is an open-source tool that acts as a proxy between Snyk and special integrations, providing for access by [snyk.io](http://snyk.io/) to sacn code in repositories that are not internet-reachable and return results to you. SCM integrations with Broker support Snyk Open Source, Snyk Code, Snyk Container (Dockerfile), and Snyk IaC. For details, see [Integrations with Snyk Broker](../#integrations-with-snyk-broker).

For comprehensive information about Snyk Broker, including how it works, how to deploy it, commit signing, upgrading, and troubleshooting, see the full [Snyk Broker user documentation](../).

## **Deployment options**

<table data-card-size="large" data-view="cards" data-full-width="false"><thead><tr><th></th><th></th><th></th><th data-hidden data-card-cover data-type="files"></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead><tbody><tr><td><strong>Install Snyk Broker with the</strong> <a href="https://github.com/snyk/snyk-broker-helm"><strong>Broker Helm Chart</strong></a>. For details, see <a href="install-and-configure-broker-using-helm.md">Install and configure Broker using Helm</a>.</td><td></td><td></td><td><a href="../../../.gitbook/assets/helmkube.png">helmkube.png</a></td><td><a href="install-and-configure-broker-using-helm.md">install-and-configure-broker-using-helm.md</a></td></tr><tr><td><strong>Install Snyk Broker</strong> <strong>using the</strong> <a href="https://github.com/snyk/broker"><strong>Docker images</strong></a> provided by Snyk. For details, see <a href="install-and-configure-broker-using-docker.md">Install and configure Broker using Docker</a>.</td><td></td><td></td><td><a href="../../../.gitbook/assets/Docker-Logo (1).jpg">Docker-Logo (1).jpg</a></td><td><a href="install-and-configure-broker-using-docker.md">install-and-configure-broker-using-docker.md</a></td></tr></tbody></table>

Alternatively, you can use the binaries available for [each Github Release](https://github.com/snyk/broker/releases).

## Additional information for developers

If you need to upgrade, see [Upgrade the Snyk Broker Client](../upgrade-the-snyk-broker-client.md).

Troubleshooting information is provided on the [Troubleshooting Broker](../troubleshooting-broker.md) page.

You can view the [license, Apache License, Version 2.0](https://github.com/snyk/broker/blob/master/LICENSE).

To submit pull requests, see [Contributing](https://github.com/snyk/broker/blob/master/.github/CONTRIBUTING.md).

See [Security](https://github.com/snyk/broker/blob/master/SECURITY.md) for specific information about Broker.
