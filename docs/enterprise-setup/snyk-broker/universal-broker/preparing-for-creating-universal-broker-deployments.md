# Preparing for creating Universal Broker deployments

Before creating deployments, ensure you have met the [prerequisites](prerequisites-for-universal-broker.md).

## Prepare hosts for the installation of Universal Broker

Snyk recommends configuring at least two separate instances of the Broker Client for each integration, either on different hosts or installed using Kubernetes. This ensures that you always have at least two instances running for redundancy. To install the Universal Broker using Helm please see the [Snyk Universal Broker Helm Chart](https://github.com/snyk/snyk-universal-broker-helm).

## Configure your network for using Universal Broker

If you use a proxy server, ensure you configure it and any firewalls to allow the Broker Client inbound and outbound access as follows.

Establish the outbound connection on port 443 from the Broker Client running in your environment to `https://broker.snyk.io` and `https://api.snyk.io` or your regional Broker URL.

Establish an internal connection that allows inbound access from the integration (SCM, CR) to the Broker Client at the BROKER\_CLIENT\_URL on the port you have configured (typically 8000). This is not inbound from the internet.

Traffic initiated from the Snyk Broker server side always uses the latest available Broker connection. All activity from the Snyk side, such as traffic driven by recurring tests, appears on only one of your replicas at a time. The amount of Snyk activity is proportional to the activity in the repositories or Jira items. That activity generates webhooks, which are distributed across all replicas.

Use the Universal Broker `snyk-broker-config` CLI tool to configure and manage connections, integrations, and all Universal Broker resources.

## Using a previous installation

If you have not previously installed the Universal Broker, refer to the Prerequisites for Universal Broker and [Basic steps to install and configure Universal Broker.](basic-steps-to-install-and-configure-universal-broker.md)

Be sure to set your environment variables to make usage easier, including when you are installing the `snyk-broker-config` CLI tool. Use the following commands:

* `export SNYK_TOKEN=xxxx` (Linux/Mac)
* `export TENANT_ID=yyyy` (Linux/Mac)
* `set SNYK_TOKEN=xxxx` (Windows)
* `set TENANT_ID=yyyy` (Windows)

If the Universal Broker has already been installed, set the Install ID as an environment variable for easier usage. Use the following commands:

* `export INSTALL_ID=zzzz` (Linux/Mac)
* `set INSTALL_ID=zzzz` (Windows)
