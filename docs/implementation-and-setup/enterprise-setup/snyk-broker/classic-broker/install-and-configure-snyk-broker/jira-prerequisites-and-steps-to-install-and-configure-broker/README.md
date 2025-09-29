# Jira - prerequisites and steps to install and configure Broker

Before installing, review the general instructions for the installation method you plan to use, [Helm](../../../../../../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-helm.md) or [Docker](../../../../../../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-docker.md).

Before installing the Snyk Jira Broker, ask your Snyk account team to provide you with a Broker token or generate it from the Snyk Web UI.

You must have Docker or a way to run Docker Linux containers.\
Some Docker deployments for Windows run only Windows containers. Ensure that your deployment is capable of running Linux containers.

To generate a Broker token from the Web UI:

1. Navigate to **Settings** > **Integrations** > **Jira** > **For installation of Jira within a private network click here**.
2. Click **Generate** to generate a Broker Token for Jira and click **Show** to confirm.

Continue with the steps to install using [Docker](jira-install-and-configure-using-docker.md) or [Helm](jira-install-and-configure-using-helm.md).
