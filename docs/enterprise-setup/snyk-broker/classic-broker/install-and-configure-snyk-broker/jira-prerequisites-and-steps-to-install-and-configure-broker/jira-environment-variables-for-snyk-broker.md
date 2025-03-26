# Jira - environment variables for Snyk Broker

The following environment variables are needed to configure the Broker Client for Jira:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your Jira integration settings view.
* `BROKER_SERVER_URL` - the URL of the Broker server for the region in which your data is hosted. For the commands and URLs to use, see [Broker URLs](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls).
* `JIRA_USERNAME` - the Jira username.
* `JIRA_PASSWORD` - the Jira password.
* `JIRA_HOSTNAME` - the hostname of your Jira deployment, such as `your.jira.domain.com`.
* `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible by your Jira for webhooks, such as `http://my.broker.client:8000`
  * This must have http:// and the port number.&#x20;
  * To configure the client with HTTPS, [additional settings are required](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker).
* `PORT` - the local port at which the Broker client accepts connections. Default is 8000.
