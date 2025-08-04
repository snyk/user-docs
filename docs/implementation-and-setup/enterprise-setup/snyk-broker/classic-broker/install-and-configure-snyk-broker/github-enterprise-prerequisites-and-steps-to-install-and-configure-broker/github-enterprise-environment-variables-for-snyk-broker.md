# GitHub Enterprise - environment variables for Snyk Broker

The following environment variables are required to configure the Broker Client: for GitHub Enterprise:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your Snyk Org settings view (app.snyk.io).
* `GITHUB_TOKEN` - a personal access token with full `repo`, `read:org` and `admin:repo_hook` scopes.
* `BROKER_SERVER_URL` - the URL of the Broker server for the region in which your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
* `GITHUB` - the hostname of your GitHub Enterprise deployment:
  * For self-hosted or custom domains, this is your GitHub Enterprise domain `your.ghe.domain.com`. &#x20;
  * For GitHub Enterprise Cloud without a custom domain, use `github.com`.
  * Do not include an org or path after the domain.
* `GITHUB_API` - the API endpoint of your GitHub Enterprise deployment. Do not use `http` or `https`.
  * For self-hosted, should be `your.ghe.domain.com/api/v3`.
  * For GitHub Enterprise Cloud without a custom domain, use `api.github.com`.
* `GITHUB_GRAPHQL` - the `graphql` endpoint of your GitHub Enterprise deployment. Do not use `http` or `https`.
  * For self-hosted, should be `your.ghe.domain.com/api`.
  * for GitHub Enterprise Cloud without a custom domain, use `api.github.com/graphql`.
* `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
* `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to your GitHub Enterprise deployment webhooks, such as `http://broker.url.example:8000`.
  * This must have `http://` and the port number.&#x20;
  * To configure the client with HTTPS, [additional settings are required](../../../https-for-broker-client-with-docker.md).
* `ACCEPT_IAC` - by default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, for example, Terraform files, you can add an environment variable `ACCEPT_IAC` with any combination of `tf,yaml,yml,json,tpl`.
* `ACCEPT_CODE` - by default, when using the Snyk Broker - Code Agent, Snyk Code will not load code snippets. To enable code snippets, you can add an environment variable `ACCEPT_CODE=true`.
* `ACCEPT_ESSENTIALS` - enable Snyk Essentials to identify your application assets, monitor them, and prioritize the risks. To enable Snyk Essentials, add the environment variable `ACCEPT_ESSENTIALS=true`.
