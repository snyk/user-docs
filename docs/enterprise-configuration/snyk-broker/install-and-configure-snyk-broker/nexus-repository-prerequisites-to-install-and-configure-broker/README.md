# Nexus Repository - prerequisites to install and configure Broker

{% hint style="info" %}
**Feature availability**

Integration with Nexus Repository Manager is available only for Enterprise plans.

For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

Before installing the Snyk Nexus Repository Broker, ask your Snyk account team to provide you with a Broker token or generate it from the Snyk Web UI.

To install using Docker, you must have Docker or a way to run Docker Linux containers.\
Some Docker deployments for Windows run only Windows containers. Ensure that your deployment is capable of running Linux containers.

## Obtain Broker token for Nexus integration

1. Go to **Settings > Integrations > Package Repositories > Nexus**
2. Verify that you see the screen to configure Nexus.

<figure><img src="../../../../.gitbook/assets/Screenshot 2022-07-15 at 15.15.11.png" alt="Configure Nexus"><figcaption><p>Configure Nexus</p></figcaption></figure>

{% hint style="info" %}
If you do not see the **Snyk Broker** switch, you do not have the necessary permissions and can only add a publicly accessible instance.

Submit a request to [Snyk Support](https://support.snyk.io/hc/en-us/requests/new) if you want to add a private registry.
{% endhint %}

When you have permissions to add a private registry, continue with the instructions to [generate a Broker token from the Web UI](./#generate-a-broker-token-from-the-web-ui).

## Generate a Broker token from the Web UI

1. In the Nexus integration settings, move the **Snyk Broker on/off** switch to **on** to display a form for generating a Broker token.
2. Select **Generate and Save.**
3. Copy the token that was generated to use when you set up the Broker Client.

