# GitHub - prerequisites to install and configure Broker

Before installing, review the general instructions for the installation method you plan to use, [Helm](../install-and-configure-broker-using-helm.md) or [Docker](../install-and-configure-broker-using-docker.md).

The prerequisites follow.

Before installing the Snyk GitHub Broker, you must configure a GitHub service account token with the [required permissions](../../../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/snyk-github-integration.md#required-permissions-scope-for-the-github-integration). All the operations, both those that are triggered through the Snyk Web UI and the automatic operations, are performed for a GitHub service account that has its token configured with the Broker.

To install using Docker, you must have Docker or a way to run Docker Linux containers. Some Docker deployments for Windows run only Windows containers. Ensure that your deployment is capable of running Linux containers.
