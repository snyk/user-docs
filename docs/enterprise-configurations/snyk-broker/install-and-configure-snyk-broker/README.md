# Install and configure Snyk Broker

## Overview of Snyk Broker

Snyk Broker is an open-source tool that acts as a proxy between Snyk and special integrations, providing for access by [snyk.io](http://snyk.io/) to your code to scan it and return results to you. SCM integrations with Broker support Snyk Open Source, Snyk Code, Snyk Container (Dockerfile), and Snyk IaC. For details about Snyk Broker, including how it works, how to deploy it, commit signing, upgrading, and troubleshooting, see the [Snyk Broker user documentation](../).

## **Deployment options**

<table data-card-size="large" data-view="cards" data-full-width="false"><thead><tr><th></th><th></th><th></th><th data-hidden data-card-cover data-type="files"></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead><tbody><tr><td><p><strong>Install Snyk Broker with the</strong> <a href="https://github.com/snyk/snyk-broker-helm"><strong>Broker Helm Chart</strong></a>.</p><p></p><p>For details, see <a href="install-and-configure-broker-using-helm.md">Install and configure Broker using Helm</a>.</p></td><td></td><td></td><td><a href="../../../.gitbook/assets/helmkube.png">helmkube.png</a></td><td><a href="install-and-configure-broker-using-helm.md">install-and-configure-broker-using-helm.md</a></td></tr><tr><td><p><strong>Install Snyk Broker</strong> <strong>using the</strong> <a href="https://github.com/snyk/broker"><strong>Docker images</strong></a> provided by Snyk.</p><p></p><p>For details, see <a href="install-and-configure-broker-using-docker.md">Install and configure Broker using Docker</a>.</p></td><td></td><td></td><td><a href="../../../.gitbook/assets/Docker-Logo.jpg">Docker-Logo.jpg</a></td><td><a href="install-and-configure-broker-using-docker.md">install-and-configure-broker-using-docker.md</a></td></tr></tbody></table>

Alternatively, you can use the binaries available for [each Github Release](https://github.com/snyk/broker/releases).
