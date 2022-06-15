# Setup Broker with GitLab

Configuring GitLab with broker is useful to ensure a secure connection with your on-premise or cloud GitLab deployment.

### To configure Broker to be used for GitLab

{% hint style="info" %}
Please ask for your CSM support to provide you with a Broker token
{% endhint %}

{% hint style="info" %}
You will require Docker or a way to run Docker containers
{% endhint %}

* Obtain your GitLab Broker token from your CSM or Support team
* To use the Broker client with GitLab.com or an on-prem GitLab deployment, run `docker pull snyk/broker:gitlab` tag. The following environment variables are mandatory to configure the Broker client:
  * `BROKER_TOKEN` - the Snyk Broker token, obtained from your GitLab integration settings view (app.snyk.io).
  * `GITLAB_TOKEN` - a GitLab personal access token with `api` scope
  * `GITLAB` - the hostname of your GitLab deployment, such as `your.gitlab.domain.com` or `GitLab.com`.
  * `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
  * `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to GitLab.com webhooks, such as `http://broker.url.example:8000`
* Example:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e GITLAB_TOKEN=<secret-gitlab-token> \
           -e GITLAB=<your.gitlab.domain.com (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT=/private/accept.json
           -v /local/path/to/private:/private \
       snyk/broker:gitlab
```

* If necessary, go to the Advanced Configuration section of [Install and configure the Snyk Broker client](../../integrations/snyk-broker/set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md) and make any configuration changes if required (For example, if the GitLab instance is using a private certificate, provide the CA (Certificate Authority) to the Broker Client configuration). A fully configured `accept.json` for Snyk IaC, Code, Open Source and Container for GitLab is attached.

{% file src="../../../.gitbook/assets/accept (3).json" %}

* Paste the Broker Client configuration to start the broker client container
* Once the container is up, the GitLab Integrations page should show the connection to GitLab and you should be able to `Add Projects`

Basic Troubleshooting

* Run `docker logs <container id>` where container id is the GitLab Broker container ID to look for any errors
* Ensure relevant ports are exposed to GitLab
