# Set up Snyk Broker with GitHub

Follow the instructions on this page to set up Snyk Broker to be used for GitHub.

## Configure Broker integration with GitHub

{% hint style="info" %}
Ask your Snyk account team to provide you with a Broker token.
{% endhint %}

{% hint style="info" %}
You need Docker or a way to run Docker containers
{% endhint %}

The following explains how to configure Snyk Broker to be used for GItHub.

**Run** `docker pull snyk/broker:github-com`. The following environment variables are mandatory to configure the Broker client:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your Snyk Org settings view (app.snyk.io).
* `GITHUB_TOKEN` - a personal access token with full `repo`, `read:org` and `admin:repo_hook` scopes.
* `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
* `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to GitHub.com webhooks, such as `http://broker.url.example:8000`
* `ACCEPT_IAC` - by default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, such as Terraform for example, you can simply add an environment variable `ACCEPT_IAC` with any combination of `tf,yaml,yml,json,tpl`
* `ACCEPT_CODE` - by default, when using the Snyk Broker - Code Agent, Snyk Code will not load code snippets. To enable code snippets you can simply add an environment variable `ACCEPT_CODE=true`

**Use the following command** to set up a fully configured Broker Client to analyze Open Source, IaC, Container and Code files.

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e GITHUB_TOKEN=<secret-github-token> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
       snyk/broker:github-com
```

**If necessary, go to the Advanced Configuration section** of [Install and configure the Snyk Broker client](../set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md) and **make any configuration changes** needed.

For example, if the GitHub instance is using a private certificate, provide the CA (Certificate Authority) to the Broker Client configuration.

You may want to provide [proxy support](https://docs.snyk.io/integrations/snyk-broker/set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client#proxy-support).

In addition, a fully configured `accept.json` for Snyk IaC, Code, Open Source and Container for GitHub is attached. You **cannot run** the `ACCEPT_IAC` and `ACCEPT_CODE` arguments at the same time as the `ACCEPT` argument.

{% file src="../../../.gitbook/assets/accept (1) (1) (1) (1) (1).json" %}

**Paste the Broker Client configuration** to start the Broker Client container.

Once the container is up, the GitHub Integrations page shows the connection to GitHub and can `Add Projects`.

## Basic troubleshooting for Broker with GitHub

### **Support of big manifest files (> 1Mb) for GitHub**

One reason that open Fix/Upgrade PRs or PR/recurring tests fail may be fetching big manifest files (> 1Mb). To address this issue, whitelist an additional Blob API endpoint in `accept.json`. This should be in a private array:

```
{
    "//": "used to get given manifest file",
    "method": "GET",
    "path": "/repos/:owner/:repo/git/blobs/:sha",
    "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
}
```

{% hint style="info" %}
To ensure the maximum possible security, Snyk does not enable this rule by default, as usage of this endpoint means that the Snyk platform can theoretically access all files in this repository, because the path does not include specific allowed file names.
{% endhint %}

### **Additional troubleshooting for Broker with GitHub**

* Run `docker logs <container id>` where `container id` is the GitHub Broker container ID to look for any errors.
* Ensure relevant ports are exposed to GitHub.
