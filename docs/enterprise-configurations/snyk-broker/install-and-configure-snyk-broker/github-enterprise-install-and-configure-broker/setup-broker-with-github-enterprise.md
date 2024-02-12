# GitHub Enterprise - install and configure using Docker

Follow the instructions on this page to set up GitHub Enterprise with Snyk Broker. This integration is useful to ensure a secure connection with your on-premise or cloud GitHub Enterprise deployment.

{% hint style="info" %}
**Prerequisites**\
Ask your Snyk account team to provide you with a Broker token.

You need Docker or a way to run Docker Linux containers. Some Docker deployments for Windows run only Windows containers. Ensure that your deployment is capable of running Linux containers.
{% endhint %}

## Configure Broker to be used for GitHub Enterprise

To use the Snyk Broker client with a GitHub Enterprise deployment, **run** `docker pull snyk/broker:github-enterprise`. Refer to [GitHub Enterprise - environment variables for Snyk Broker](github-enterprise-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

**If necessary,** go to the [Advanced configuration page](../advanced-configuration-for-snyk-broker-docker-installation/) and **make any configuration changes** needed, such as providing the CA (Certificate Authority) to the Broker Client configuration if the GitHub Enterprise instance is using a private certificate, and setting up [proxy support](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md).

## Docker run command to set up a Broker Client for GitHub Enterprise

**Copy the following command** to set up a fully configured Broker Client to analyze Open Source, IaC, Container, Code files (with the Code Agent), and Snyk AppRisk information. Enable [Snyk AppRisk](../../../../manage-risk/snyk-apprisk/) to identify your application assets, monitor them, and prioritize the risks.

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e GITHUB_TOKEN=<secret-github-token> \
           -e GITHUB=<your.ghe.domain.com (no http/s)> \
           -e GITHUB_API=<your.ghe.domain.com/api/v3 (no http/s)> \
           -e GITHUB_GRAPHQL=<your.ghe.domain.com/api (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_APPRISK=true \
       snyk/broker:github-enterprise
```

{% hint style="info" %}
Snyk AppRisk is set by default to **`false`**. Enable it by setting the flag to **`true`**.
{% endhint %}

As an alternative to using the Docker run command, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the GitHub Enterprise integration.

## Start the Broker Client container and verify the connection with GitHub Enterprise

Paste the Broker Client configuration to start the Broker Client container.

Once the container is up, the GitHub Enterprise Integrations page shows the connection to GitHub Enterprise and you can `Add Projects.`

## Basic troubleshooting for Broker with GitHub Enterprise

### **Support of big manifest files (> 1Mb) for GitHub Enterprise**

One reason that open Fix/Upgrade PRs or PR/recurring tests fail may be fetching big manifest files (> 1Mb). To address this issue, enable an additional variable in your broker by following the Additional instructions for [Snyk Open Source Scans (SCA) of large manifest files (Docker setup) ](https://docs.snyk.io/enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-open-source-scans-sca-of-large-manifest-files-docker-setup)&#x20;

{% hint style="info" %}
To ensure the maximum possible security, Snyk does not enable this rule by default, as use of this endpoint means that the Snyk platform can theoretically access all files in this repository, because the path does not include specific allowed file names.
{% endhint %}

### **Additional troubleshooting for Broker with GitHub Enterprise**

* Run `docker logs <container id>` to look for any errors, where container id is the GitHub Enterprise Broker container ID.
* Ensure relevant ports are exposed to GitHub Enterprise.
