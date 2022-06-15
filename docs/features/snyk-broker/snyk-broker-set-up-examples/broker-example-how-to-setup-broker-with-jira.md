# Setup Broker with GitHub

Configuring GitHub Enterprise with Broker is useful to ensure a secure connection with your GitHub account.

### To configure Broker to be used for GitHub Enterprise

{% hint style="info" %}
Ask for your CSM support to provide you with a Broker token
{% endhint %}

{% hint style="info" %}
You need Docker or a way to run Docker containers
{% endhint %}

* Obtain your GitHub Broker token from your CSM or Support team
* To use the Broker client with GitHub.com, run `docker pull snyk/broker:github-com`. The following environment variables are mandatory to configure the Broker client:
  * `BROKER_TOKEN` - the Snyk Broker token, obtained from your Snyk Org settings view (app.snyk.io).
  * `GITHUB_TOKEN` - a personal access token with full `repo`, `read:org` and `admin:repo_hook` scopes.
  * `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
  * `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to  GitHub.com webhooks, such as `http://broker.url.example:8000`
* Example:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e GITHUB_TOKEN=<secret-github-token> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT=/private/accept.json
           -v /local/path/to/private:/private \
       snyk/broker:github-com
```

* If necessary, go to the Advanced Configuration section of [Install and configure the Snyk Broker client](../../integrations/snyk-broker/set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md) and make any configuration changes needed. For example, if the GitHub instance is using a private certificate, provide the CA (Certificate Authority) to the Broker Client configuration. A fully configured `accept.json` for Snyk IaC, Code, Open Source and Container for GitHub is attached.

{% file src="../../../.gitbook/assets/accept (3) (1).json" %}

* Paste the Broker Client configuration to start the broker client container.
* Once the container is up, the GitHub Integrations page should show the connection to GitHub and you should be able to `Add Projects`.

### Basic Troubleshooting

#### **Support of big manifest files (> 1Mb) for GitHub / GitHub Enterprise**

One reason for failing of open Fix/Upgrade PRs or PR/recurring tests may be fetching big manifest files (> 1Mb) failure. To address this issue, whitelist an additional Blob API endpoint in `accept.json`. This should be in a private array:

```
{
    "//": "used to get given manifest file",
    "method": "GET",
    "path": "/repos/:owner/:repo/git/blobs/:sha",
    "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
}
```

**Note** To ensure the maximum possible security, Snyk does not enable this rule by default, as usage of this endpoint means that the Snyk platform can theoretically access all files in this repository, as the path does not include specific allowed file names.

#### **Additional troubleshooting**

* Run `docker logs <container id>` where container id is the GitHub Broker container ID to look for any errors
* Ensure relevant ports are exposed to GitHub
