# Bitbucket Server/Data Center - environment variables for Snyk Broker Basic Auth

The following environment variables are required to configure the Broker client:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your Bitbucket Server integration settings view (app.snyk.io).
* `BROKER_SERVER_URL` - the URL of the Broker server for the region in which your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
* `BITBUCKET_USERNAME` - the Bitbucket Server username.
* `BITBUCKET_PASSWORD` - the Bitbucket Server password. You can use an API token in place of a password.
* `BITBUCKET` - the hostname of your Bitbucket Server deployment, such as `your.bitbucket-server.domain.com`.
* `BITBUCKET_API` - the API endpoint of your Bitbucket Server deployment. Should be `$BITBUCKET/rest/api/1.0`.
* `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to your Bitbucket Server for webhooks, such as `http://broker.url.example:8000.`This URL is required to access features such as PR Fixes or PR Checks.
  * This must have http:// and the port number.&#x20;
  * To configure the client with HTTPS, [additional settings are required](../../../https-for-broker-client-with-docker.md).
* `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
* `ACCEPT_IAC` - by default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, such as Terraform for example, you can simply add an environment variable `ACCEPT_IAC` with any combination of `tf,yaml,yml,json,tpl`.
* `ACCEPT_CODE` - by default, Snyk Code will not load code snippets. To enable code snippets you can add an environment variable `ACCEPT_CODE=true`.
* `ACCEPT_ESSENTIALS` - enable Snyk Essentials to identify your application assets, monitor them, and prioritize the risks. To enable Snyk Essentials, add the environment variable `ACCEPT_ESSENTIALS=true`.
