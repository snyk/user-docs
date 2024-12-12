# Initial configuration of the Universal Broker

The high-level steps in implementing the Universal Broker are the same regardless of the configuration method you use. Using the command`snyk-broker-config` walks you through these steps, easing the onboarding, while direct API calls require a greater understanding of the overall Universal Broker models. The steps follow:

{% hint style="info" %}
Prerequisite: You must be a Tenant admin to be able to create deployments, credential references, and connections.
{% endhint %}

* **One time:** Install the Snyk Broker App in your Organization. This returns an install ID, a client ID, and a client secret, all needed to interact with the Snyk platform. The Organization ID is required to create the deployment.
* **One time:** Create a Universal Broker deployment for your tenant ID and install ID.
* **One time:** Create credential references needed for your connections.
* **One time:** Create your desired connection or connections.
* **One time for each Organization integration:** Configure the Organizations and integrations that should use the connection.

<figure><img src="../../../.gitbook/assets/image 7 (1).png" alt="Illustration of steps in implementing the Universal Broker"><figcaption><p>Illustration of steps in implementing the Universal Broker for GtHub</p></figcaption></figure>

## Configuration using the `snyk-broker-config` command (recommended) <a href="#using-snyk-broker-config-cli" id="using-snyk-broker-config-cli"></a>

Follow these steps to configure the Universal Broker using the `snyk-broker-config` command.

{% hint style="info" %}
To use this command, you must install Node 18 or higher.
{% endhint %}

1. Run `npm i -g snyk-broker-config`.
2. Set the necessary environment variables:
   * `SNYK_TOKEN` if not already set. This must be your personal API key.
   * `SNYK_API_HOSTNAME` if you are not targeting https://api.snyk.io, for example,`export SNYK_API_HOSTNAME=https://api.eu.snyk.io`. See [Broker URLs](../../../working-with-snyk/regional-hosting-and-data-residency.md#broker-urls).
3. Ensure that as you follow the remaining steps, you set more environment variables as needed to ensure the experience proceeds smoothly, for example:
   * `TENANT_ID` so you do not have to enter it on every command.
   * `INSTALL_ID` if you have one; otherwise, the tool will walk you through the installation process.
4. Run `snyk-broker-config commands` to list the available commands.
5. Run `snyk-broker-config workflows` to list the available interactive workflows.
6. Create a deployment.
   * Run `snyk-broker-config workflows deployments create`.
   * Add any metadata key-value pair or pairs that are needed.
   * **Note**: This workflow will ask you to install a Snyk App if it is not already installed. Ensure that you note the installation ID (also export it as an environment variable) and credentials output for use in the succeeding steps. If you lose the credentials, you must install the App again and recreate or update the deployment.
   * **Optional**:  When the workflow deployment has been created, view your deployment using `snyk-broker-config workflows deployments get`.
7. Create and configure your connection or connections.
   * Run `snyk-broker-config workflows connections create` to create a connection.
   * In response to the prompt **Which Deployment do you want to use?**: select your deployment from the list presented.
   * In response to the prompt **Which Connection type do you want to create?**:
     * Select the type of connection you want to create from the list presented.\
       Options include SCM connection types like `github` and variants, `bitbucket` and variants, `gitlab`, and `azure`, as well as container registry connections (see the next step), package manager connections, Jira, and more.
     * For container registry-type Broker connections, specify a CR\_AGENT\_URL to point to a Container Registry Agent.\
       You must configure and run both the Universal Broker and a separate Container Registry Agent. Follow the instructions for [configuring and running a Container Registry Agent](../snyk-broker-container-registry-agent/#configuring-and-running-the-container-registry-agent).
   * Provide the configuration for each required field in response to the prompts:
     * Enter a human-friendly name for your connection. Note that no spaces are allowed.
     * Enter the **broker\_client-url** (the hostname and port of your Broker client, for example, https://my.broker.company.com:8000).
     * Enter the credential reference or choose the option **CreateNew** in response to a prompt like **github-token (Sensitive): Which Credential Reference do you want to use? Or create New?**
   * When you see the messages **Connection created** and **Ready to configure integrations to use this connection**, you can run the Broker client.
8. After the connection is created, use `snyk-broker-config workflows connections integrate` to configure an integration to use the newly created connection.\
   In response to the prompts, enter the `deployment` you want to use, the **connection** you want to use, the `OrgID` of the Organization you want to integrate, and the `integration ID` of the type `github`.\
   You can find your `integration ID` in your **Organization** **Integrations** settings or retrieve it using the [Integrations](../../../snyk-api/reference/integrations-v1.md) API.

You may see that some integrations do not show an integration ID in the Snyk Web UI.

* For JIRA, Artifactory, and Nexus, no integration ID is needed.
* For the AppRisk connection type (not through SCM), the integration step is not needed during the beta process as configuration is done differently. See the AppRisk documentation to find where to copy the connection identifier listed in the connection.
* For all other connection types, the integration ID may not be visible in the Snyk UI at first. You may need to go through a configuration wizard, entering dummy values, to get to the screen where the Integration ID is visible.

## Example: first-time configuration of a new connection <a href="#quick-examples-below" id="quick-examples-below"></a>

1. Start by creating the first-time setup and connection, as explained in [the preceding section](initial-configuration-of-the-universal-broker.md#using-snyk-broker-config-cli).\
   See the [First time setup and connection recording](https://asciinema.org/a/YqSmUHEWMcDPeQKm6lpeG3qhM).
2. Integrate your connection with an Organization ID and Integration ID.\
   See the [Integration recording](https://asciinema.org/a/I2QJxi9MDEeThRZTLD1aTv9cN).
3. List the Organization IDs and integration IDs for your connection.\
   See the [List Organization IDs and integration IDs recording](https://asciinema.org/a/5RWuySWT0M2dDI9mARJjeZS5g).
4. Run the Broker client.

## Run your Broker deployment on your container engine or Kubernetes cluster

Target your desired environment with the usual `-e BROKER_SERVER_URL=https://broker`, `REGION.snyk.io \` if you are not using `broker.snyk.io`.

Add the environment variable or variables as defined in the credentials references with the associated value. The connection will not be established if references are missing, and an error entry will be logged in the Broker Client logs.

| <pre><code>docker run --restart=always \
    -p 8000:8000 \
    -e ACCEPT_CODE=true \
    -e DEPLOYMENT_ID=&#x3C;DEPLOYMENTID> \
    -e CLIENT_ID=&#x3C;CLIENTID> \
    -e CLIENT_SECRET=&#x3C;CLIENTSECRET> \
    -e UNIVERSAL_BROKER_ENABLED=true \
    -e PORT=8000 \
    -e BROKER_HA_MODE_ENABLED=true \
    -e &#x3C;YOUR_CREDENTIALS_REFERENCE>=&#x3C;secret value> \
snyk/broker:universal
</code></pre> |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
