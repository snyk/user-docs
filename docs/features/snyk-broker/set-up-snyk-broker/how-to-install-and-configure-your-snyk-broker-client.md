# How to install and configure your Snyk Broker client

{% hint style="info" %}
Snyk’s recommended and supported method of running the Snyk Broker client is using Docker.
{% endhint %}

To install and configure your client:

1. Visit [the broker repository](https://github.com/snyk/broker) and follow the instructions for the relevant SCM to run one of our provided images, or derive your own.
2. For the environment variables required to run the Broker client, you must:
   1. Retrieve your [unique broker client token,](https://docs.snyk.io/integrations/snyk-broker/retrieve-a-unique-broker-client-token)
   2. Ensure that you have the necessary permissions granted to your API credentials for the integration you are using by following [one of our integration guides.](https://docs.snyk.io/integrations)
3. After the broker is running, **identifying Broker server** appears towards the end of the log results, showing that a connection to Snyk has been established
4. Visit the settings for your brokered integration in [the Snyk application](https://app.snyk.io) to see a message like the following:

![](<../../../.gitbook/assets/image (60).png>)

You can now use Snyk as normal.

#### Advanced Configuration

**HTTPS**

The Broker client runs an HTTP server by default. It can be configured to run an HTTPS server for local connections. This requires an SSL certificate and a private key to be provided to the docker container at runtime.

For example, if your certificate files are found locally at `./private/broker.crt` and `./private/broker.key`, provide these files to the docker container by mounting the folder and using the `HTTPS_CERT` and `HTTPS_KEY` environment variables:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e GITHUB_TOKEN=secret-github-token \
           -e PORT=8000 \
           -e HTTPS_CERT=/private/broker.crt \
           -e HTTPS_KEY=/private/broker.key \
           -e BROKER_CLIENT_URL=https://my.broker.client:8000 \
           -v /local/path/to/private:/private \
       snyk/broker:github-com
```

Note that `BROKER_CLIENT_URL` now has the HTTPS scheme.

**Git with an internal certificate**

By default, the Broker client establishes HTTPS connections to the Git. If your Git is serving an internal certificate (signed by your own CA), you can provide the CA certificate to the Broker client.

For example, if your CA certificate is at `./private/ca.cert.pem`, provide it to the docker container by mounting the folder and using the `CA_CERT` environment variable:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e BITBUCKET_USERNAME=username \
           -e BITBUCKET_PASSWORD=password \
           -e BITBUCKET=your.bitbucket-server.domain.com \
           -e BITBUCKET_API=your.bitbucket-server.domain.com/rest/api/1.0 \
           -e PORT=8000 \
           -e CA_CERT=/private/ca.cert.pem \
           -v /local/path/to/private:/private \
       snyk/broker:bitbucket-server
```

#### Custom approved-listing filter

The default approved-listing filter supports the bare minimum to operate on all repositories supported by Snyk. In order to customize the approved-listing filter, create the default one locally by installing `snyk-broker` and running `broker init [Git type]`. The created `accept.json` is the default filter for the chosen Git. Place the file in a separate folder such as `./private/accept.json`, and provide it to the docker container by mounting the folder and using the `ACCEPT` environment variable:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e GITHUB_TOKEN=secret-github-token \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=https://my.broker.client:8000 \
           -e ACCEPT=/private/accept.json
           -v /local/path/to/private:/private \
       snyk/broker:github-com
```

More information on setting up custom approved-listing filters: [https://github.com/snyk/broker#custom-approved-listing-filter](https://github.com/snyk/broker#custom-approved-listing-filter)

Example of custom approved-listing filters: [https://github.com/snyk/broker/tree/master/client-templates](https://github.com/snyk/broker/tree/master/client-templates)
