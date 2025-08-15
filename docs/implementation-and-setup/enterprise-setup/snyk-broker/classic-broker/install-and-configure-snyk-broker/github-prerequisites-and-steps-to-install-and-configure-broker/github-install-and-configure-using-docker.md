# GitHub - install and configure using Docker

Before installing, review the [prerequisites](./) and the general instructions for installation using [Docker](../install-and-configure-broker-using-docker.md).

This integration is useful to ensure a secure connection with your on-premise or cloud GitHub deployment.

## Configure Broker integration with GitHub

To use the Snyk Broker Client with GitHub, **run** `docker pull snyk/broker:github-com`. Refer to [GitHub - environment variables for Snyk Broker](github-environment-variables-for-snyk-broker.md) for definitions of the environment variables.

**If necessary,** go to the [Advanced configuration page](../advanced-configuration-for-snyk-broker-docker-installation/) and **make any configuration changes needed** such as providing the CA (Certificate Authority) to the Broker Client configuration if the GitHub instance is using a private certificate, and setting up [proxy support](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md).&#x20;

## Docker run command to set up a Broker Client for GitHub

**Copy the following command** to set up a fully configured Broker Client to analyze Open Source, IaC, Container, Code files (with the Code Agent), and Snyk Essentials information. Enable [Snyk Essentials](../../../../../../scan-with-snyk/snyk-essentials.md) to identify your application assets, monitor them, and prioritize the risks.

{% hint style="info" %}
**Multi-tenant settings for regions**\
When installing, you must add a command in your script to set the `BROKER_SERVER_URL`.This is the URL of the Broker server for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
{% endhint %}

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e GITHUB_TOKEN=<secret-github-token> \
           -e BROKER_SERVER_URL=<broker-region-url> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_ESSENTIALS=true \ 
       snyk/broker:github-com
```

{% hint style="info" %}
Snyk Essentials is set by default to **`false`**. Enable it by setting the flag to **`true`**.
{% endhint %}

## Start the Broker Client container and verify the connection with GitHub

Paste the Broker Client configuration to start the Broker Client container.

After the container is up, the GitHub Integrations page shows the connection to GitHub, and you can `Add Projects`.

## Basic troubleshooting for Broker with GitHub

### **Support of big manifest files (> 1Mb) for GitHub**

One reason that open Fix/Upgrade PRs or PR/recurring tests fail may be fetching big manifest files (> 1Mb). To address this issue, enable an additional variable in your broker by following the Additional instructions for [Snyk Open Source Scans (SCA) of large manifest files (Docker setup)](../advanced-configuration-for-snyk-broker-docker-installation/snyk-open-source-scans-sca-of-large-manifest-files-docker-setup.md). &#x20;

{% hint style="info" %}
To ensure the maximum possible security, Snyk does not enable this rule by default, as use of this endpoint means that the Snyk platform can theoretically access all files in this repository because the path does not include specific allowed file names.
{% endhint %}

### **Additional troubleshooting for Broker with GitHub**

* Run `docker logs <container id>` to look for any errors, where `container id` is the GitHub Broker container ID.
* Ensure relevant ports are exposed to GitHub.
