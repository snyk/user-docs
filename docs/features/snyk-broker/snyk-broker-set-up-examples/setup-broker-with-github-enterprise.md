# Setup Broker with GitHub Enterprise

Configuring GitHub Enterprise with broker is useful to ensure a secure connection with your on-premise or cloud GitHub Enterprise deployment.

### Configure Broker to be used for GitHub Enterprise

{% hint style="info" %}
Please ask for your CSM support to provide you with a Broker token
{% endhint %}

{% hint style="info" %}
You will require Docker or a way to run Docker containers
{% endhint %}

* Obtain your GitHub Enterprise Broker token from your CSM or Support team
* To use the Broker client with a GitHub Enterprise deployment, run `docker pull snyk/broker:github-enterprise` tag. The following environment variables are mandatory to configure the Broker client:
  * `BROKER_TOKEN` - the Snyk Broker token, obtained from your Snyk Org settings view (app.snyk.io).
  * `GITHUB_TOKEN` - a personal access token with full `repo`, `read:org` and `admin:repo_hook` scopes.
  * `GITHUB` - the hostname of your GitHub Enterprise deployment, such as `your.ghe.domain.com`.
  * `GITHUB_API` - the API endpoint of your GitHub Enterprise deployment. Should be `your.ghe.domain.com/api/v3`.
  * `GITHUB_GRAPHQL` - the graphql endpoint of your GitHub Enterprise deployment. Should be `your.ghe.domain.com/api`.
  * `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
  * `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to your GitHub Enterprise deployment webhooks, such as `http://broker.url.example:8000`
* Example:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e GITHUB_TOKEN=<secret-github-token> \
           -e GITHUB=<your.ghe.domain.com (no http/s)> \
           -e GITHUB_API=<your.ghe.domain.com/api/v3 (no http/s)> \
           -e GITHUB_GRAPHQL=<your.ghe.domain.com/api (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT=/private/accept.json
           -v /local/path/to/private:/private \
       snyk/broker:github-enterprise
```

* If necessary, go to the Advanced Configuration section of [Install and configure the Snyk Broker client](../set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md) and make any configuration changes if required (For example, if the GitHub Enterprise instance is using a private certificate, provide the CA (Certificate Authority) to the Broker Client configuration). A fully configured `accept.json` for Snyk IaC, Code, Open Source and Container for GitHub Enterprise is attached.

{% file src="../../../.gitbook/assets/accept (4).json" %}

* Paste the Broker Client configuration to start the broker client container
* Once the container is up, the GitHub Enterprise Integrations page should show the connection to GitHub Enterprise and you should be able to `Add Projects`

### Basic Troubleshooting

**Support of big manifest files (> 1Mb) for GitHub / GitHub Enterprise**

One of the reason for failing of open Fix/Upgrade PRs or PR/recurring tests might be fetching big manifest files (> 1Mb) failure. To address this issue, additional Blob API endpoint should be whitelisted in `accept.json`:

* should be in `private` array

```
{
    "//": "used to get given manifest file",
    "method": "GET",
    "path": "/repos/:owner/:repo/git/blobs/:sha",
    "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
}
```

**Note** To ensure the maximum possible security, we do not enable this rule by default, as usage of this endpoint means that the Snyk platform can theoretically access all files in this repository, as the path does not include specific allowed file names.

**Additional troubleshooting**

* Run `docker logs <container id>` where container id is the GitHub Enterprise Broker container ID to look for any errors
* Ensure relevant ports are exposed to GitHub Enterprise
