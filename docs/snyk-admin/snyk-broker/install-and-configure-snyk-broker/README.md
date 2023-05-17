# Install and configure Snyk Broker

Snyk Broker is an open-source tool that acts as a proxy between Snyk and special integrations, allowing [snyk.io](http://snyk.io/) access to your code for scanning and returning results to you. SCM integrations with Broker support Snyk Open Source, Snyk Code, Snyk Container (Dockerfile), and Snyk IaC. For details about Snyk Broker including how it works, how to deploy it, commit signing, how to upgrade, and troubleshooting, see the [Snyk Broker user documentation](https://docs.snyk.io/snyk-admin/snyk-broker).

**If you are using Kubernetes**, Snyk recommends that you **install Snyk Broker with the** [**Broker Helm Chart**](https://github.com/snyk/snyk-broker-helm). For details, see [Install and configure Broker using Helm](install-and-configure-broker-using-helm.md).

For **all other environments**, you can install Snyk Broker using the [Docker images](https://github.com/snyk/broker) provided by Snyk. For details, see [Install and configure Broker using Docker](install-and-configure-broker-using-docker.md).
