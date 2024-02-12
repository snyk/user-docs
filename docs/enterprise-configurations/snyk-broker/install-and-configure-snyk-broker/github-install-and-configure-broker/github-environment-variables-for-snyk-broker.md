# GitHub - environment variables for Snyk Broker

The following environment variables are required to configure the Broker Client for GitHub:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your Snyk Org settings view (app.snyk.io).
* `GITHUB_TOKEN` - a personal access token with full `repo`, `read:org` and `admin:repo_hook` scopes.
* `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
* `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to GitHub webhooks, such as `http://broker.url.example:8000.`This URL is required to access features such as PR Fixes or PR Checks.
  * This must have http:// and the port number.&#x20;
  * To configure the client with HTTPS, [additional settings are required](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker).
* `ACCEPT_IAC` - by default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, for example, Terraform, you can add an environment variable `ACCEPT_IAC` with any combination of `tf,yaml,yml,json,tpl`
* `ACCEPT_CODE` - by default, when using the Snyk Broker - Code Agent, Snyk Code will not load code snippets. To enable code snippets you can simply add an environment variable `ACCEPT_CODE=true`
* `ACCEPT_APPRISK` - enable Snyk AppRisk to identify your application assets, monitor them, and prioritize the risks. To enable Snyk AppRisk, add the environment variable `ACCEPT_APPRISK=true`
