# Artifactory Repository - prerequisites and steps to install and configure Broker

{% hint style="info" %}
**Feature availability**

Integration with Artifactory Repository is available only for Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Before installing, review the general instructions for the installation method you plan to use, [Helm](../../../../../../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-helm.md) or [Docker](../../../../../../enterprise-setup/snyk-broker/classic-broker/install-and-configure-snyk-broker/install-and-configure-broker-using-docker.md).

The prerequisites follow.

Before installing the Snyk Artifactory Repository Broker, ask your Snyk account team to provide you with a Broker token or generate it from the Snyk Web UI.

You must have Docker or a way to run Docker Linux containers. Some Docker deployments for Windows run only Windows containers. Ensure that your deployment is capable of running Linux containers.

For convenience, instructions to obtain or generate the Broker token follow. When you are done, continue with the steps to install using [Docker](artifactory-repository-install-and-configure-using-docker.md) or [Helm](artifactory-repository-install-and-configure-using-helm.md).

## Obtain Broker token for Artifactory Repository setup

1. Navigate to **Settings** > **Add integrations** > **Package Repositories > Artifactory**.
2. Enter the URL of your Artifactory instance, this must end with /artifactory.
3. Enter your username and password.
4. Select **Save**.

<figure><img src="../../../../../../.gitbook/assets/screenshot_2020-04-17_at_14.38.12.png" alt="Artifactory integration setup"><figcaption><p>Artifactoryrepository setup</p></figcaption></figure>

{% hint style="info" %}
If you do not see the Snyk Broker on/off switch, you do not have the necessary permissions and can only add a publicly accessible instance.

Submit a request to [Snyk Support](https://support.snyk.io) if you want to add a private registry.
{% endhint %}

When you have the permissions needed to add a private registry, continue with the instructions to [generate a Broker token from the Web UI](./#generate-a-broker-token-from-the-web-ui).

## Generate a Broker token from the Web UI

1. In the Artifactory integration settings, move the Snyk Broker on/off switch to **on** to display a form for generating a Broker token.
2. Select **Generate and Save.**
3. Copy the token that was generated to use when you set up the Broker Client.
