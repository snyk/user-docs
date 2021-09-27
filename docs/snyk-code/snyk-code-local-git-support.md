# Snyk Code local git support

You can natively connect Snyk Code to your local git server. This allows customers who are using a self-hosted git provider such as GitHub Enterprise to find, prioritize and fix potential vulnerabilities in their 1st-party code.

## Code access components

* **Broker server**: Running on Snyk SaaS backend
* **Broker client**: A [Docker image](https://hub.docker.com/r/snyk/broker/) deployed in your infrastructure.
* **Code agent**: Another [Docker image](https://hub.docker.com/r/snyk/code-agent/) that is deployed in your infrastructure. **Note:** Code agent is only supported with [Snyk Broker](https://docs.snyk.io/integrations/snyk-broker) v4.108.0 and later versions. If you have a running Broker client, please pull the latest update.

The **Broker client** and **code agent** components are deployed in your infrastructure, creating two separate services, responsible for cloning local repositories in a secured manner and sending the allied information to Snyk.

The Broker client provides the Agent with the connection details. The Agent uses these details to connect to your local git repository, clone the relevant files. And send the results through the brokered communications using callbacks. The brokered communication happens when a Broker client connects \(using your Broker ID\) to a Broker server running in Snyk environment:

![](../.gitbook/assets/local-git.png)

See [Snyk Broker](https://docs.snyk.io/integrations/snyk-broker/broker-introduction) documentation for more details.

## Setup

### Prerequisites

Before you begin with the setup process, please make sure to have a server that supports these minimal requirements for running the Broker client and Code agent:

* CPU:  1 vcpu
* Memory:  2Gb \(should be reflected in node memory setting\)
* Disk space: 2Gb \(available disk size determines maximum cloneable repository size\)
* Network: code upload performance will be affected by slow Internet connection

### Set up broker client

Code agent depends on broker client. Follow the instructions on [How to install and configure your Snyk Broker client](../integrations/snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md) for detailed instructions how to set up broker for specific SCMs.

If you already have a broker client running, please consider the following additional requirements:

* Code agent is only supported with [Snyk Broker](https://docs.snyk.io/integrations/snyk-broker) v4.108.0 and later versions, make sure to pull the latest version first.
* Code agent needs permission to clone the full repository, make sure that the SCM token passed to the broker has the corresponding permissions.

### Set up the network

To run both the broker client and the broker agent, establish a network connection between them. There are different solutions to expose one container connection with tools like Ngrok \(which is also possible here if you want\), but this description focuses on docker bridge networks.

Run **`docker network create <network>`**

For example:

```text
docker network create mySnykBrokerNetwork
```

You can confirm that it was created by running **`docker network ls`**, this will show results like this:

```text
  NETWORK ID     NAME                 DRIVER     SCOPE
  d1353a2b0f66   mySnykBrokerNetwork  bridge     local
```

### Set up Code Agent

First, pull the code-agent image:

```text
docker pull snyk snyk/code-agent
```

The following environment variables are mandatory to configure the code agent:

* **SNYK\_TOKEN -**  your snyk token, as also used by the CLI, see [Authenticate the CLI with your account](../snyk-cli/install-the-snyk-cli/authenticate-the-cli-with-your-account.md#authenticate-using-your-api-token) for additional details.
* **PORT** - the local port, for which the code agent accepts connections, Default is 3000.

To run the **code-agent:**

```text
docker run -it --name code-agent \
     -p 3000:3000 \
     -e PORT=3000 -e SNYK_TOKEN=<token> --network mySnykBrokerNetwork \ 
     snyk/code-agent
```

In this example:

* We set the current container to use the new network we created **--network mySnykBrokerNetwork**
* We gave the current container a name  **--name code-agent**. It will be used to define the **GIT\_CLIENT\_URL** for the broker client that we will run next.

### Extend Broker setup

Extend your broker setup with the following arguments:

```text
-e GIT_CLIENT_URL=http://<code agent container>:<code agent port>
--network <name of created network>
```

For example, to extend an existing broker client configured for Gitlab, run:

```text
docker run -it \
   -p 8001:8000 \
   -e BROKER_TOKEN= \
   -e GITLAB_TOKEN= \
   -e GITLAB= \
   -e PORT=8000 \
   -e GIT_CLIENT_URL=http://code-agent:3000 \
   --network mySnykBrokerNetwork \
   snyk/broker:gitlab
```

In this example:

* We set the current container to use the new network we created **--network mySnykBrokerNetwork** 
* In **GIT\_CLIENT\_URL** we used the name we defined in the code-agent container as the host here.

If you have a running Snyk broker with a custom whitelist \(**accept.json**\), then ensure the following rule is present in the whitelist:

```text
{
  "//": "used to redirect requests to snyk git client",
  "method": "any",
  "path": "/snykgit/*",
  "origin": "${GIT_CLIENT_URL}"
}
```

\(The rule is present by default, so only needed if you override the rule with a custom whitelist.\)

## Advanced Settings

### Enable code snippets

To enable code snippets, additional rules must be added to **accept.json**.

See [https://github.com/snyk/broker\#custom-approved-listing-filter](https://github.com/snyk/broker#custom-approved-listing-filter) for detailed instructions how to extend **accept.json**.

For GitHub:

```text
{
  "//": "needed to load code snippets",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
}
```

For Gitlab:

```text
{
  "//": "needed to load code snippets",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files/:path",
  "origin": "https://${GITLAB}"
}
```

For BitBucket Server:

```text
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

```text
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
After these snippets are added, all content from the repository can be accessed through Snyk broker.
{% endhint %}

### Proxy support

For instructions how to run Broker client through a proxy, see [https://github.com/snyk/broker](https://github.com/snyk/broker). Make sure that requests to the Code agent are not sent through the proxy, by passing `NO_PROXY=<code agent container>`, for example:

```text
-e HTTP_PROXY=http://my.proxy.address:8080
-e HTTPS_PROXY=http://my.proxy.address:8080
-e NO_PROXY=code-agent
```

For code agent, add the following environment variables to the **docker run** command:

```text
-e HTTP_PROXY=http://my.proxy.address:8080
-e HTTPS_PROXY=http://my.proxy.address:8080
```

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

