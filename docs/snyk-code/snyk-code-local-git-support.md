# Snyk Code local git support

You can natively connect Snyk Code to your local git server. This allows customers who are using a self-hosted git provider such as GitHub Enterprise to find, prioritize and fix potential vulnerabilities in their 1st-party code.

## Code access components

* **Broker server**: Running on Snyk SaaS backend
* **Broker client**: A [Docker image](https://hub.docker.com/r/snyk/broker/) deployed in your infrastructure.
* **Code agent**: Another [Docker image](https://hub.docker.com/r/snyk/code-agent/) that is deployed in your infrastructure. **Note:** Code agent is only supported with [Snyk Broker](integrations/snyk-broker/) v4.108.0 and later versions. If you have a running Broker client, please pull the latest update.

The **Broker client** and **code agent** components are deployed in your infrastructure, creating two separate services, responsible for cloning local repositories in a secured manner and sending the allied information to Snyk.

The Broker client provides the Agent with the connection details. The Agent uses these details to connect to your local git repository, clone the relevant files. And send the results through the brokered communications using callbacks. The brokered communication happens when a Broker client connects \(using your Broker ID\) to a Broker server running in Snyk environment:

![](../.gitbook/assets/local-git.png)

See [Snyk Broker](integrations/snyk-broker/broker-introduction/) documentation for more details.

## Setup

### Prerequisites

Before you begin with the setup process, please make sure to have a server that supports these minimal requirements for running the Broker client and Code agent:

* CPU:  1 vcpu
* Memory:  2Gb \(should be reflected in node memory setting\/)
* Disk space: 2Gb \(available disk size determines maximum cloneable repository size\/)
* Network: code upload performance will be affected by slow Internet connection

### Set up the remote connection

To use the broker client with code-agent deployment run:

* **docker pull snyk/code-agent**
* **docker pull snyk/broker:** 

For example:

```text
docker pull snyk snyk/broker:gitlab
```

### Broker client variables

The following environment variables are mandatory to configure the Broker client.

* **BROKER\_TOKEN** - The Snyk Broker token, obtained from your Container registry integration \(provided by Snyk support\/)
* **GIT\_CLIENT\_URL -** The url of your code-agent \(more on how to get this in the \`setting up the network\` section\).
* **&lt;SCM\_NAME&gt; -** here you will need to provide your scm of your choice and their URL \(without schema, example: GITLAB=my.gitlabconnection.com\/)
* **&lt;SCM\_TOKEN&gt; -**here you will need to provide a token for that has permissions to access your scm \(Example:GITLAB\_TOKEN=YOUR\_GITLAB\_TOKEN\/)
* **PORT** - The local port at which the Broker client accepts connections. Default is 7341.

### Code agent variables

The following environment variables are mandatory to configure the code agent:

* **SNYK\_TOKEN -**  your snyk token
* **PORT** - the local port, for which the code agent accepts connections, Default is 3000.

## Setting up the network

To run both the broker client and the broker agent, establish a network connection between them. There are different solutions to expose one container connection with tools like Ngrok \(which is also possible here if you want\), but this description focuses on docker bridge networks.

Run **docker network create**.

For example:

```text
docker network create mySnykBrokerNetwork.
```

You can confirm that it was created by running **docker network ls**, this will show results like this:

```text
  NETWORK ID     NAME                 DRIVER     SCOPE
  d1353a2b0f66   mySnykBrokerNetwork  bridge     local
```

The next example uses gitlab as an example. You can use other SCMs here instead.

To run the **code-agent:**

```text
docker run -it --name code-agent \
     -p 3000:3000 \
     -e PORT=3000 -e SNYK_TOKEN= --network mySnykBrokerNetwork \ 
     snyk/code-agent
```

In this example:

* We set the current container to use the new network we created **--network mySnykBrokerNetwork**
* We gave the current container a name  **--name code-agent**. It will be used to define the **GIT\_CLIENT\_URL** for the broker client that we will run next.

To run the **broker-client:**

```text
docker run -it \
   -p 8001:8000 \
   -e BROKER_TOKEN= \
   -e GITLAB_TOKEN= \
   -e GITLAB= \
   -e PORT=8000 -e GIT_CLIENT_URL=http://code-agent:3000 --network mySnykBrokerNetwork \
   snyk/broker:gitlab
```

In this example:

* We set the current container to use the new network we created **--network mySnykBrokerNetwork** 
* In **GIT\_CLIENT\_URL**  we used the name we defined in the code-agent container as the host here.

## Extend existing Broker setup

If you have a running Snyk broker, make sure to update it to the latest version and extend it with the following arguments:

```text
--network <name of created network>
-e GIT_CLIENT_URL=http://<name of code agent container>:<port of code agent container>
```

If you have a running Snyk broker with a custom whitelist \(**accept.json**\), then ensure the following rule is present in the whitelist:

```text
{
"//": "used to redirect requests to snyk git client",
"method": "any",
"path": "/snykgit/*",
"origin": "${GIT_CLIENT_URL}"
}
```

\(The rule is present by default, so only needed if you override the rule with a custom whitelist.\/)

## Enable code snippets

To enable code snippets, additional rules must be added to **accept.json**.

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
After these snippets are added, all content from repository can be accessed through Snyk broker.
{% endhint %}

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

