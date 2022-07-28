# Setup Broker with Bitbucket Server/Data Center

Configuring Bitbucket Server/Data Center with broker is useful to ensure a secure connection with your on-premise Bitbucket deployment.

### To configure Broker to be used for Bitbucket

{% hint style="info" %}
Please ask for your CSM support to provide you with a Broker token
{% endhint %}

{% hint style="info" %}
You will require Docker or a way to run Docker containers
{% endhint %}

* Obtain your Bitbucket Broker token from your CSM or Support team
* To use the Broker client with a Bitbucket Server deployment, run `docker pull snyk/broker:bitbucket-server` tag. The following environment variables are mandatory to configure the Broker client:
  * `BROKER_TOKEN` - the snyk broker token, obtained from your Bitbucket Server integration settings view (app.snyk.io).
  * `BITBUCKET_USERNAME` - the Bitbucket Server username.
  * `BITBUCKET_PASSWORD` - the Bitbucket Server password.
  * `BITBUCKET` - the hostname of your Bitbucket Server deployment, such as `your.bitbucket-server.domain.com`.
  * `BITBUCKET_API` - the API endpoint of your Bitbucket Server deployment. Should be `$BITBUCKET/rest/api/1.0`.
  * `BROKER_CLIENT_URL` - the full URL of the Broker client as it will be accessible to your Bitbucket Server for webhooks, such as `http://broker.url.example:8000`
  * `PORT` - the local port at which the Broker client accepts connections. Default is 800.
* Example:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e BITBUCKET_USERNAME=<username> \
           -e BITBUCKET_PASSWORD=<password> \
           -e BITBUCKET=<your.bitbucket-server.domain.com (no http/s)> \
           -e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0 (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT=/private/accept.json
           -v /local/path/to/private:/private \
       snyk/broker:bitbucket-server
```

* If necessary, go to the Advanced Configuration section of [Install and configure the Snyk Broker client](../set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client.md) and make any configuration changes if required (For example, if the Bitbucket instance is using a private certificate, provide the CA (Certificate Authority) to the Broker Client configuration). A fully configured `accept.json` for Snyk IaC, Code, Open Source and Container for Bitbucket is attached.

{% file src="../../../.gitbook/assets/accept (4).json" %}

* Paste the Broker Client configuration to start the broker client container
* Once the container is up, the Bitbucket Integrations page should show the connection to Bitbucket and you should be able to `Add Projects`

Basic Troubleshooting

* Run `docker logs <container id>` where container id is the Bitbucket Broker container ID to look for any errors
* Ensure relevant ports are exposed to Bitbucket
