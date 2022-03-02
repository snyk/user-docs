# Snyk Code self-hosted git (with broker)

{% hint style="info" %}
This feature is currently in Beta. Please contact your CSM if you are interested in participating.
{% endhint %}

Snyk Broker enables you to connect your local git server to Snyk if the git server is not internet reachable. The Snyk Broker Code agent allows customers who are using a self-hosted git provider such as GitHub Enterprise, BitBucket Server or GitLab to find, prioritize and fix potential vulnerabilities in their 1st-party code.

## Snyk Code Broker access components

* **Broker server**: Running on Snyk SaaS backend
* **Broker client**: A [Docker image](https://hub.docker.com/r/snyk/broker/) deployed in your infrastructure.
* **Code agent**: Another [Docker image](https://hub.docker.com/r/snyk/code-agent/) that is deployed in your infrastructure. **Note:** Code agent is only supported with [Snyk Broker](https://docs.snyk.io/integrations/snyk-broker) v4.108.0 and later versions. If you have a running Broker client, pull the latest update.

The **Broker client** and **code agent** components are deployed in your infrastructure, creating two separate services, responsible for cloning local repositories in a secured manner and sending the allied information to Snyk.

The Broker client provides the Agent with the connection details. The Agent uses these details to connect to your local git repository, clone the relevant files. And send the results through the brokered communications using callbacks. The brokered communication happens when a Broker client connects (using your Broker ID) to a Broker server running in Snyk environment:

![](../../../.gitbook/assets/local-git.png)

See [Snyk Broker Code Agent](../../../features/integrations/snyk-broker/snyk-broker-code-agent.md) documentation for more details.
