# Snyk Code Agent

{% hint style="info" %}
This feature is currently in Beta. If you would like to set it up in your organization, contact your CSM.&#x20;
{% endhint %}

To use Snyk Code capabilities via the Snyk Broker, you need to set up and run the Snyk Code Agent. The Snyk Code Agent enables you to connect Snyk Code to your local Git server via the Snyk Broker. By using the Code Agent with the Snyk Broker, you can store your repositories on a self-hosted Git provider, and apply Snyk Code analysis to these repositories in order to find, prioritize, and fix potential vulnerabilities in your source code.

**Notes**:&#x20;

* For more information on brokered communication, see [Snyk Broker](../../../features/snyk-broker/).
* The Code Agent currently does not support the Automatic PR Checks.

### Snyk Code - Access components to a self-hosted Git provider

To apply Snyk Code analysis to repositories that are stored on your local Git server, you need the following components:&#x20;

* **Broker Server**: Running on Snyk SaaS backend.\
  **Note:** The Broker Server is provided by Snyk, and no installation is required on your part.&#x20;
* **Broker Client**: A [Docker image](https://hub.docker.com/r/snyk/broker/) deployed in your infrastructure.\
  **Note**: For more information, see [Install and configure the Snyk Broker client.](../../../features/snyk-broker/set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md)
* **Code Agent**: Another [Docker image](https://hub.docker.com/r/snyk/code-agent/) that is deployed in your infrastructure. \
  **Note:** The Code Agent is only supported in [Snyk Broker](https://docs.snyk.io/integrations/snyk-broker) v. 4.108.0 and later versions. If you already have a running Broker Client, pull the latest update.

The **Broker Client** and **Code Agent** components are deployed in your infrastructure, creating two separate services. These services are responsible for cloning local repositories in a secured manner for analysis, and sending only the required information to Snyk.

The Broker Client automatically provides the Code Agent with the connection details to your Git server. The Code Agent uses these details to perform the following: connecting to your local Git server, cloning the relevant files, and sending the results through the brokered communications using callbacks. The brokered communication happens when a Broker Client connects (using your Broker ID) to a Broker Server running in Snyk environment:

![Brokered communication](../../../.gitbook/assets/local-git.png)

See [Snyk Broker](https://docs.snyk.io/integrations/snyk-broker/broker-introduction) documentation for more details.

## Code Agent - Setup

### Prerequisites

Before you start the setup process, make sure that your server has the following minimal requirements for running the Broker Client and Code Agent:

* **CPU:** 1 vCPU.
* **Memory:** 2Gb (should be set in node memory setting).
* **Disk space:** 2Gb (the available disk size determines the maximum repository size that can be cloned).
* **Network:** the code upload speed will be affected by a slow Internet connection.

{% hint style="info" %}
Make sure that at this point you have reached out to your Customer Success Manager (CSM) or Support and they have enabled the broker integration.
{% endhint %}



The full workflow of setting up the Code Agent is as follows:

1\.  Setting up the the Broker Client in your infrastructure.

2\.  Setting up an internal network for the Broker Client - Code Agent communication.

3\. Setting up the Code Agent. &#x20;

&#x20;&#x20;

### 1. Set up the Broker Client

To use the Code Agent, you must have a Broker Client in your infrastructure. To set up the Broker Client for a specific SCM, see [How to install and configure your Snyk Broker client](../../../features/snyk-broker/set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md).  \
\
Once the Client Broker is set up successfully, the following message appears in the terminal:

`successful websocket connection established`&#x20;

**If you already have a Broker Client running**, consider the following additional requirements:

* The Code Agent is only supported in [Snyk Broker](../../../features/snyk-broker/) v. 4.108.0 and later versions. Therefore, pull the latest Broker Client version before you continue to the next steps.
* The Code Agent needs permissions to clone the full repository. Verify that the SCM token that you pass to the Broker has the required permissions.

### 2. Set up the network

To run both the Broker Client and the Code Agent, you need to establish a network connection between them. To perform this, you can create a docker bridge network, as described below.

**Note:** To establish this network connection, you can also use other methods and tools, like Ngrok. &#x20;

**To establish a network connection between the Broker Client and the Broker Agent:**

* Run:

```
docker network create <network>
```



For example:

```
docker network create mySnykBrokerNetwork
```

In this example, `mySnykBrokerNetwork` is the name of the new network.

**To confirm the creation of the network connection:**

* Run:

```
docker network ls
```

The results should look similar to the following:

```
  NETWORK ID     NAME                 DRIVER     SCOPE
  d1353a2b0f66   mySnykBrokerNetwork  bridge     local
```

### 3. Set up the Code Agent

The following environment variables are mandatory for the configuration of the Code Agent:

* **SNYK\_TOKEN** - your Snyk API token, as also used by the CLI. See [Authenticate the CLI with your account](../../../snyk-cli/authenticate-the-cli-with-your-account.md), and follow the instructions to retrieve your token.
* **PORT** - the local port, where the Code Agent accepts connections. Default is `3000`.

**To set up the Code Agent:**

* Run:

```
docker run --name code-agent \
     -p 3000:3000 \
     -e PORT=3000 -e SNYK_TOKEN=<Snyk API token> --network mySnykBrokerNetwork \
     snyk/code-agent
```

Where:

* **`--name code-agent`** - a new name for the Code Agent container. This name will be used to define the **`GIT_CLIENT_URL`** for the Broker Client that we will run next.
* **SNYK\_TOKEN** - your Snyk API token, as also used by the CLI. See [Authenticate the CLI with your account](../../../snyk-cli/authenticate-the-cli-with-your-account.md), and follow the instructions to retrieve your token.
* **PORT** - the local port, where the Code Agent accepts connections. Default is `3000`.



In this example:

* We set the current Code Agent container to use the new network we created in the [previous step](snyk-broker-code-agent.md#set-up-the-network): **`--network mySnykBrokerNetwork`**
* We gave the current container a name: **`--name code-agent`**. This name will be used to define the **`GIT_CLIENT_URL`** for the Broker Client that we will run next.

Once the set up is completed successfully, the following message appears in the terminal:&#x20;

`{ ..., "msg":"Application started", ... }`&#x20;

### 4. Extend the Broker Client setup

If you already have a running Client Broker, run it again with the following additional arguments:

```
-e GIT_CLIENT_URL=http://<code agent container>:<code agent port>
--network <name of created network>
```

Where:

* &#x20;`GIT_CLIENT_URL` - a URL that points to the running `code-agent` instance
* &#x20;`--network` - the network you have defined [previously](snyk-broker-code-agent.md#set-up-the-network).

For example, to extend an existing Broker Client configured for GitLab, run:

```
docker run \
   -p 8001:8000 \
   -e BROKER_TOKEN=<xxxxxxx-xxxx-xxxx-xxxx-xxxxxxx> \
   -e GITLAB_TOKEN=glpat-<xxxxxxxxxxxxxxx> \
   -e GITLAB=url.com \
   -e PORT=8000 \
   -e GIT_CLIENT_URL=http://code-agent:3000 \
   --network mySnykBrokerNetwork \
   snyk/broker:gitlab
```

In this example:

* We set the current container to use the new network we created: **`--network mySnykBrokerNetwork`**
* In **`GIT_CLIENT_URL`** we used the name (with the `--name` flag) we [defined](snyk-broker-code-agent.md#set-up-the-network) in the Code Agent container as the host.

Also, **if you already had your Broker Client running with a custom whitelist** (or `accept.json`) you should ensure that the following rule is present:

```
{
  "//": "used to redirect requests to snyk git client",
  "method": "any",
  "path": "/snykgit/*",
  "origin": "${GIT_CLIENT_URL}"
}
```

The rule is present by default, so it is only needed if you have run the broker with the `-e ACCEPT=/private/accept.json` flag previously.&#x20;

{% hint style="warning" %}
At this point please contact your Customer Success Manager or Support to enable the integration
{% endhint %}

**Note:** To learn more about custom whitelists, see the **Custom approved-listing filter** sub-section in the [Install and configure the Snyk Broker client](../../../features/snyk-broker/set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md) page.

### 5. Test the Code Agent

After you successfully completed the set up of all the components of the Snyk Broker - Code Agent, you can test your new set up. To test the set up, import a repository to Snyk via the Web UI, and verify that a **Code analysis** Project was created by Snyk Code for the source code of the repository. &#x20;

To test the installation of the Code Agent, it, head to [app.snyk.io](http://app.snyk.io) and [import](../../snyk-code/getting-started-with-snyk-code/activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/) a [supported](../../snyk-code/snyk-code-language-and-framework-support.md) project. This should results in you seeing a bunch of issues in the UI.

**To test the Code Agent set up:**\


## Advanced settings

### Enable code snippets

To enable code snippets in the Snyk Code results, additional rules must be added to the `private array` in the `accept.json` file.

For more information on extending the `accept.json` file, see  [https://github.com/snyk/broker#custom-approved-listing-filter](https://github.com/snyk/broker#custom-approved-listing-filter)

For GitHub:

```
{
  "//": "needed to load code snippets",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
}
```

For GitLab:

```
{
  "//": "needed to load code snippets",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files/:path",
  "origin": "https://${GITLAB}"
}
```

For BitBucket Server and Bitbucket Data Center:

```
{
    "//": "needed to load code snippets",
      "method": "GET",
      "path": "/projects/:project/repos/:repo/browse*/:file",
      "origin": "https://${BITBUCKET_API}",
      "auth": {
        "scheme": "basic",
        "username": "${BITBUCKET_USERNAME}",
        "password": "${BITBUCKET_PASSWORD}"
      }
}
```

For Azure Repos:

```
{
      "//": "needed for code snippets",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories/:repo/items",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
}
```

{% hint style="info" %}
After these snippets are added, all content from the repository can be accessed through Snyk Broker.
{% endhint %}

### Proxy support

For instructions how to run Broker client through a proxy, see [https://github.com/snyk/broker](https://github.com/snyk/broker). Make sure that requests to the Code Agent are not sent through the proxy, bypassing `NO_PROXY=<code agent container>`, for example:

```
-e HTTP_PROXY=http://my.proxy.address:8080
-e HTTPS_PROXY=http://my.proxy.address:8080
-e NO_PROXY=code-agent
```

For Code Agent, add the following environment variables to the **docker run** command:

```
-e HTTP_PROXY=http://my.proxy.address:8080
-e HTTPS_PROXY=http://my.proxy.address:8080
```

To disable certificate verification, for example., in the case of self-signed certificates, add to the code-agent **docker run** command:

```
-e NODE_TLS_REJECT_UNAUTHORIZED=0
```

## **Troubleshooting**

For more information on how to troubleshoot Snyk Broker - Code Agent, see [Troubleshooting Broker](../../../features/snyk-broker/troubleshooting-broker.md).

