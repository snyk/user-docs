# GitLab - prerequisites and steps to install and configure Broker

Before installing, **review the general instructions** for the installation method you plan to use, [Helm](../install-and-configure-broker-using-helm.md) or [Docker](../install-and-configure-broker-using-docker.md).

The **prerequisites** follow.

You must be a GitLab administrator. In the GitLab settings, filter outbound requests to [allow requests to the local network from webhooks and integrations](https://docs.gitlab.com/ee/security/webhooks.html#allow-requests-to-the-local-network-from-webhooks-and-integrations) and to [allow outbound requests to certain IP addresses and domains](https://docs.gitlab.com/ee/security/webhooks.html#allow-outbound-requests-to-certain-ip-addresses-and-domains).

Before installing the Snyk GitLab Broker, ask your Snyk account team to provide you with a Broker token.

Ensure that the GitLab permissions are correct for integration with Snyk. For details, see [Snyk GitLab integration](../../../../../developer-tools/scms/organization-level-integrations/gitlab.md).

You must have Docker or a way to run Docker Linux containers. Some Docker deployments for Windows run only Windows containers. Ensure that your deployment is capable of running Linux containers.

**Continue** with the steps to install using [Docker](gitlab-install-and-configure-using-docker.md) or [Helm](gitlab-install-and-configure-using-helm.md).
