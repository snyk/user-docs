# GitLab - environment variables for Snyk Broker

The following environment variables are required to configure the Broker Client for GitLab:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your GitLab integration settings view (app.snyk.io).
* `GITLAB_TOKEN` - a GitLab personal access token with `api` scope.
* `BROKER_SERVER_URL` - the URL of the Broker server for the region in which your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
* `GITLAB` - the hostname of your GitLab deployment, such as `your.gitlab.domain.com` or `GitLab.com`.
* `PORT` - the local port at which the Broker Client accepts connections. Default is 8000.
* `BROKER_CLIENT_URL` - the full URL of the Broker Client as needed to be reachable by either GitLab.com or on-prem GitLab deployment in order to establish webhook connectivity. This must be a full URL like `http://broker.url.example:8000`.
  * This must have http:// and the port number.&#x20;
  * To configure the client with HTTPS, [additional settings are required](../../../https-for-broker-client-with-docker.md).
* `ACCEPT_IAC` - by default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, such as Terraform for example, you can simply add an environment variable `ACCEPT_IAC` with any combination of `tf,yaml,yml,json,tpl`.
* `ACCEPT_CODE` - by default, Snyk Code will not load code snippets. To enable code snippets you can simply add an environment variable `ACCEPT_CODE=true`.
* `ACCEPT_ESSENTIALS` - Enable Snyk Essentials to identify your application assets, monitor them, and prioritize the risks. You can enable it by adding an environment variable `ACCEPT_ESSENTIALS=true`.
