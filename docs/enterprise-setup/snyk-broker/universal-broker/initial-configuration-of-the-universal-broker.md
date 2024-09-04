# Initial configuration of the Universal Broker

The high-level steps in implementing the Universal Broker are the same regardless of the configuration method you use. The Snyk Broker config tool walks you through these steps, easing the onboarding, while direct API calls require a better understanding of the overall Universal Broker models.

* **One time:** Install the Snyk Broker App into an Organization under your Tenant. This returns an install ID, a client ID, and a client secret, all needed to interact with the Snyk platform. The org ID is required to create the deployment.
* **One time:** Define a deployment for your tenant ID and install ID.
* **One time:** Define credentials references needed for your connections.
* **One time:** Define your desired connection(s).
* **One time for each organization integration:** Configure the organizations and integrations that should use the connection.

<figure><img src="../../../.gitbook/assets/image (563).png" alt="Illustration of steps in implementing the Universal Broker"><figcaption><p>Illustration of steps in implementing the Universal Broker</p></figcaption></figure>

## Configuration using the `snyk-broker-config` command <a href="#using-snyk-broker-config-cli" id="using-snyk-broker-config-cli"></a>

Follow these steps to configure the Universal Broker using the `snyk-broker-config` command.

1. `npm i -g snyk-broker-config`
2. Set the necessary environment variables:
   * `SNYK_TOKEN` if not already set
   * `SNYK_API_HOSTNAME` if not targeting api.snyk.io
3. Set more environment variables as needed to ensure the experience proceeds smoothly:
   * `TENANT_ID` so you do not have to enter it on every command&#x20;
   * `INSTALL_ID` if you have one, otherwise the tool will walk you through the installation process
4. Run `snyk-broker-config commands` to list the available commands
5. Run `snyk-broker-config workflows` to list the available interactive workflows
6. Create a deployment
   * Run `snyk-broker-config workflows deployment create`
   * Add any metadata key-value pair(s) needed by following the instructions.
   * **Optional**:  When the worklfow deployment has been created, view your deployment using `snyk-broker-config workflows deployments get`
7. Create and configure your connection(s).
   * Run `snyk-broker-config workflows connections create` to create a connection.
   *   Select your deployment.\
       &#x20;

       <figure><img src="../../../.gitbook/assets/image (554).png" alt="Select your deployment step"><figcaption><p>Select your deployment step</p></figcaption></figure>
   *   Select the type of connection you want to create.\


       <figure><img src="../../../.gitbook/assets/image (555).png" alt="Select connection type to create"><figcaption><p>Select connection type to create</p></figcaption></figure>


   * Provide the configuration for each required field
     1.  Human friendly name for connection (no spaces)\


         <figure><img src="../../../.gitbook/assets/image (556).png" alt="Enter a connection name"><figcaption><p>Enter a connection name</p></figcaption></figure>
     2.  Provide configuration for the required fields\


         <figure><img src="../../../.gitbook/assets/image (557).png" alt="Provide data for the relevant field(s)"><figcaption><p>Provide data for the relevant field(s)</p></figcaption></figure>
     3.  Sensitive fields like tokens or passwords require the use of credentials reference. Use an existing credentials reference or create a new one.\


         <figure><img src="../../../.gitbook/assets/image (558).png" alt="Use or create new Credentials Reference"><figcaption><p>Use or create new Credentials Reference</p></figcaption></figure>
     4.  The connection is created and is ready for integration.\


         <figure><img src="../../../.gitbook/assets/image (559).png" alt="Connection successfully created"><figcaption><p>Connection successfully created</p></figcaption></figure>
8. You can then run the Broker client using the command that follows.
9.  After the connection is created, use the `snyk-broker-config workflows connection integrate` to configure an integration to use the newly created connection\


    <figure><img src="../../../.gitbook/assets/image (560).png" alt="ntegrate your connection against an organization ID + Integration ID"><figcaption><p>Integrate your connection against an organization ID + Integration ID</p></figcaption></figure>

## Example: First-time configuration of a new connection <a href="#quick-examples-below" id="quick-examples-below"></a>

1. Getting started - first time setup and connection as explained in [the preceding section](initial-configuration-of-the-universal-broker.md#using-snyk-broker-config-cli).\
   See the [First time setup and connection recording](https://asciinema.org/a/YqSmUHEWMcDPeQKm6lpeG3qhM).
2. Integrate your connection with an organization ID and Integration ID.\
   See the [Integration recording](https://asciinema.org/a/I2QJxi9MDEeThRZTLD1aTv9cN).
3. List the organization IDs and iintegration IDs for your connection.\
   See the [List organizaton IDs and integration IDs recording](https://asciinema.org/a/5RWuySWT0M2dDI9mARJjeZS5g).
4. Run the Broker client.

## Run your Broker deployment on your container engine or Kubernetes cluster

Target your desired environment with the usual `-e BROKER_SERVER_URL=https://broker.pre-prod.snyk.io \` if you are not using `broker.snyk.io`.

| <pre><code>docker run --restart=always \
    -p 8000:8000 \
    -e ACCEPT_CODE=true \
    -e DEPLOYMENT_ID=&#x3C;DEPLOYMENTID> \
    -e CLIENT_ID=&#x3C;CLIENTID> \
    -e CLIENT_SECRET=&#x3C;CLIENTSECRET> \
    -e UNIVERSAL_BROKER_ENABLED=true \
    -e PORT=8000 \
    -e BROKER_HA_MODE_ENABLED=true \
snyk/broker:universal
</code></pre> |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
