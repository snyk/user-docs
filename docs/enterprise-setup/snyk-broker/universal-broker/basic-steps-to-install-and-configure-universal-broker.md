# Basic steps to install and configure Universal Broker

{% hint style="info" %}
You can learn more about Universal Broker in the dedicated Snyk Learn course. Explore the advantages, configuration, architecture, and much more with [Snyk Learn: Universal Broker](https://learn.snyk.io/lesson/universal-broker/?ecosystem=general).
{% endhint %}

Follow these steps to install and configure your Universal Broker using the `snyk-broker-config` CLI tool. The tool guides you through the steps and indicates important points in the workflows.

## Install and set up the snyk-broker-config CLI tool

To install the tool, use `npm i -g snyk-broker-config`.

{% hint style="info" %}
The `snyk-broker-config` CLI tool is your primary guide for setting up and managing Snyk Broker connections. It is designed to walk you step-by-step through the process, asking for the required information for each integration type. This adaptive method guarantees that you receive the latest parameter requirements directly from the tool.
{% endhint %}

The basic process for configuring a new Universal Broker deployment is as follows:
1. Install the `snyk-broker-config` tool by running `npm i -g snyk-broker-config`
2. Create a new Universal Broker connection by following the create workflow `snyk-broker-config workflows connections create`
3. Integrate your new Universal Broker connection by following the integrate workflow `snyk-broker-config workflows connections integrate`

### How to use the CLI for parameter discovery

Even if you are not running the full setup, the `snyk-broker-config` CLI tool can help you understand the required parameters. Use the interactive setup or the command line help.

* **Interactive setup:** Run `snyk-broker-config workflows connections create` and follow the prompts. The CLI asks for all required and optional parameters based on the integration type you select, such as GitLab, Artifactory, or Bitbucket Server.
* **Command line help:** Use the `--help` flag for any `snyk-broker-config` command to see available options and parameters. For instance, to see parameters specifically relevant for creating a GitLab connection type, use the following command:

    ```bash
    snyk-broker-config workflows connections create --type gitlab --help
    ```

### Example of a connection workflow

A typical workflow for adding a new Broker connection using the CLI involves these steps:

1.  **Install the CLI:**

    ```bash
    npm install -g snyk-broker-config
    ```
2.  **Start the interactive workflow:**

    ```bash
    snyk-broker-config workflows connections create
    ```

    The CLI will then guide you through the process, prompting for:

    * Your Snyk API Token, required for authentication if you did not already export it as an environment variable.
    * The Snyk Organization ID where the Broker connection is used.
    * The specific type of integration you want to connect, such as `gitlab`, `artifactory`, `bitbucket-server`.
    * All required and optional parameters, such as URLs, tokens, usernames, or passwords, are dynamically identified for your chosen integration type. Follow the on-screen instructions carefully.

The interactive workflow is the most straightforward way to ensure all necessary parameters are correctly provided for your Broker connection.

## Create your first connection

This returns an install ID, a client ID, and a client secret, which are all needed to interact with the Snyk platform.

* After you install, start the Universal Broker Create Connection workflow.

```
> snyk-broker-config workflows connections create
Using https;//api.snyk.io (or https://api.REGION.snyk.io)
Universal Broker Create Connection workflow
   Enter your Snyk Token
```

* Type your Snyk API token and press Enter.

```
✓ Valid Snyk Token.
✓ Tenant Admin role confirmed.
    Have you installed the Broker App against an Org? (Y/N)
```

* Type N and press Enter.

```
Enter Org ID to install Broker App. Must be in Tenant <uuid returned>.
(Must be a valid uuid).
```

* Paste the Snyk Broker Admin Organization ID created in the prerequisites and press Enter.

The Broker App facilitates the secure connection and communication with the Broker server through OAuth.

```
App installed. Please store the following credentials securely:
- client id: <client ID>
- ClientSecret: <snyk_client-secret>
You will need them to run your Broker Client.
    Have you saved these credentials? (Y/N)
```

The tool displays the credentials for the Broker App just installed. Be sure to store these safely like any other credentials. This is the only time the client secret will be displayed. If you lose these credentials, you must either delete and recreate the Snyk Broker Admin Organization and start over, or use the API to reinstall Universal Broker manually.

* When you have saved your credentials, type Y and press Enter.

```
Helpful tip ! Set TENANT_ID, INSTALL_ID as environment values to avoid pasting 
the values in for every command.
Now using Tenant ID <current Tenant ID> and Install ID <current Install ID>.
Do you want to create a new Deployment? (Y/N)
```

Snyk recommends that you set the INSTALL_ID as an environment variable to improve usability:\
- `export INSTALL_ID=zzzz` (Linux/Mac)\
- `set INSTALL_ID=zzzz` (Windows)

## Input connection parameters

This includes creating credentials references needed for your connections. Each deployment is limited to a maximum of 25 connections.

* In response to the prompt, type Y and press Enter.

```
Which Connection type do you want to create?
acr
apprisk
artifactory
…
> github
…
```

* Select the connection type you want to create.

This example shows creating a GitHub connection. Creating all the other types of connection follows the same process. Each deployment is limited to 25 connections.

```
Enter a human-friendly name for your Connection.
```

* Enter a connection name to help you identify the connection, for example, github-connection-for-team-x.

```
Enter a human-friendly name for your Connection. <name you entered>
broker_client_url: Broker client url. Must be url.
```

* Enter your Broker\_client\_url. Snyk recommends using the default value. You can enter a different value, which is required for container integrations.

```
broker_client_url: Broker client url. Must be url. <region-specific URL you entered>
github_token (Sensitive): No existing Credential Reference for this Connection type. 
CreateNew
Env Var Name (e.g., MY_GITHUB_TOKEN). (Must be a valid envvar).
```

* Create the credential reference (not the actual credential value). Enter the name of the environment variable which will contain the actual credential value when the Broker client is running, for example, MY\_GITHUB\_TOKEN.
* Optionally, you can enter a comment to help you keep track of this connection.

```
Env Var Name (e.g. MY_GITHUB_TOKEN). (Must be a valid envvar). <MY_GITHUB_TOKEN>
Comment this is a demo broker connection.
```

When you run the Broker client container in a subsequent step, you must add the `-e MY_GITHUB_TOKEN=<SECRET_TOKEN_VALUE>`. In a production setup, these values are mounted from the secrets vault.

```
Connection created with ID <ID number>. Ready to configure integrations 
to use this connection.
Connection Create workflow completed.
```

The connection is now created.

## Validate your connection (optional)

You can use the following workflow to display details about the connection.

```
> snyk-broker-config workflows connections get
```

* If you are prompted about whether the Broker app is installed, enter Y and then paste the install ID you saved previously.

Exporting the INSTALL\_ID avoids this optional step in your terminal session in the future. The deployment details follow.

```
Now using Deployment <name>.
Selected Connection ID <number>. 
Ready to configure integrations to use this connection.
```

Details of the connection follow: `connection ID`; `connection type (broker_connection)`; `attributes:  deployment_id, identifier, name, and secrets: primary and secondary`, each with the `status`, `encrypted`, `expires_at`, and `nonce`; `configuration required: broker-client-url` and `github_token values`; `type: github.`

```
Connection Detail Workflow completed.
```

## Integrate your connection with an Organization that will use the Universal Broker

```
> snyk-broker-config workflows connections integrate
Enter the OrgID you want to integrate.. (Must be a valid  uuid).
```

* Enter the ID of the Organization where you want to use the newly created Broker connection.

```
Enter the OrgID you want to integrate.. (Must be a valid  uuid). <uuid-entered>
Connection <number returned> (type:github) integrated with <integration <number>.

Connection Integration Workflow completed
```

Your Organization is now integrated with your new Broker connection.

## Run the Broker client

```
docker run -d --restart=always \
-p 8000:8000 \
-e DEPLOYMENT_ID=<DEPLOYMENT_ID_JUST_CREATED> \
-e CLIENT_ID=<CLIENT_ID_SAVED_EARLIER> \
-e CLIENT_SECRET=<CLIENT_SECRET_SAVED_EARLIER> \
-e MY_GITHUB_TOKEN=<THE_ACTUAL_GITHUB_TOKEN_VALUE> \
-e PORT=8000 \
snyk/broker:universal
```

When the Broker client has started, the connection is ready to use, in this case, to import repositories.

* To verify that your connection is configured, check that the integration tile on your **Organization Settings** > **Integrations** page is marked **Configured**.

## Integrate your connection with more Organizations

To integrate your connection with another Organization, run the integrate command again and enter the new Organization ID. Repeat as needed to connect with multiple Organizations.

```
> snyk-broker-config workflows connections integrate
```

* Repeat the step for any Organization in your Tenant as needed, for as many integrations as you need.

\
