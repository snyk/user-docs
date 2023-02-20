# Set up Broker with GitLab





Configuring GitLab with broker is useful to ensure a secure connection with your on-premise or cloud GitLab deployment.

## Configure Broker to be used for GitLab

{% hint style="info" %}
Ask your Snyk account team to provide you with a Broker token.
{% endhint %}

{% hint style="info" %}
You need Docker or a way to run Docker containers
{% endhint %}

* Obtain your GitLab Broker token from Snyk.
* To use the Broker client with GitLab.com or an on-prem GitLab deployment, run `docker pull snyk/broker:gitlab` tag. The following environment variables are required to configure the Broker client:
  * `BROKER_TOKEN` - the Snyk Broker token, obtained from your GitLab integration settings view (app.snyk.io).
  * `GITLAB_TOKEN` - a GitLab personal access token with `api` scope
  * `GITLAB` - the hostname of your GitLab deployment, such as `your.gitlab.domain.com` or `GitLab.com`.
  * `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
  * `BROKER_CLIENT_URL` - the full URL of the Broker client as needed to be reachable by either GitLab.com or on-prem GitLab deployment in order to establish webhook connectivity. This must be a full URL like `http://broker.url.example:8000`
  * `ACCEPT_IAC` - by default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, such as Terraform for example, you can simply add an environment variable `ACCEPT_IAC` with any combination of `tf,yaml,yml,json,tpl`
  * `ACCEPT_CODE` - by default, when using the Snyk Broker - Code Agent, Snyk Code will not load code snippets. To enable code snippets you can simply add an environment variable `ACCEPT_CODE=true`
* Example:

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e GITLAB_TOKEN=<secret-gitlab-token> \
           -e GITLAB=<your.gitlab.domain.com (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
       snyk/broker:gitlab
```

* This command sets up a fully configured Broker client that will analyze Open Source, IaC, Container and Code files.
* If necessary, go to the Advanced Configuration section of [Install and configure the Snyk Broker client](../set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md) and make any configuration changes ineeded.\
  For example, if the GitLab instance is using a private certificate, provide the CA (Certificate Authority) to the Broker Client configuration.\
  A fully configured `accept.json` for Snyk IaC, Code, Open Source and Container for GitLab is attached. You **cannot run** the `ACCEPT_IAC` and `ACCEPT_CODE` arguments at the same time as the `ACCEPT` argument.

{% file src="../../../.gitbook/assets/accept.json" %}

* Paste the Broker Client configuration to start the broker client container.
* Once the container is up, the GitLab Integrations page should show the connection to GitLab and you should be able to `Add Projects`.

## Basic troubleshooting for Broker with GitLab

* Run `docker logs <container id>` where container id is the GitLab Broker container ID to look for any errors.
* Ensure relevant ports are exposed to GitLab.
