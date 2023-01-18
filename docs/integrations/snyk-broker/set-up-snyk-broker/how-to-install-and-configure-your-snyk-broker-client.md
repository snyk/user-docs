# Install and configure the Snyk Broker client

{% hint style="info" %}
Snyk recommends and supports using Docker as the method of running the Snyk Broker client.
{% endhint %}

To install and configure your Snyk Broker client:

1. Visit [the broker repository](https://github.com/snyk/broker) and follow the instructions for the relevant integration to run one of the images provided by Snyk, or derive your own.\
   There are some examples available at [Snyk Broker integration setups](../snyk-broker-set-up-examples/).
2. For the environment variables required to run the Broker client, you must:
   1. Retrieve your [unique broker client token](prepare-snyk-broker-for-deployment.md).
   2. Ensure that you have the necessary permissions granted to your API credentials for the integration you are using by following [one of the Snyk integration guides.](https://docs.snyk.io/integrations)
3. Once Broker is running, **identifying Broker server** appears toward the end of the log results, showing that a connection to Snyk has been established
4. Visit the settings for your brokered integration in [the Snyk Web UI](https://app.snyk.io) to see a message like the following:

<figure><img src="../../../.gitbook/assets/image (60) (1).png" alt="Brokered integration settings showing connected status"><figcaption><p>Brokered integration settings showing connected status</p></figcaption></figure>

The base setup for SCMs enables use of Snyk Open Source and scanning Dockerfiles. To get the full platform set up through Broker, follow these instructions:

* [Infrastructure as Code detection](../snyk-broker-infrastructure-as-code-detection/)
* [Container registry integration](../snyk-broker-container-registry-agent/)
* [Snyk Code analysis](../snyk-broker-code-agent/)

## Snyk Broker Helm Chart

Alternatively, if you are using Kubernetes and would like to deploy Broker through a Helm chart, use the provided [Snyk Broker Helm Chart](https://github.com/snyk/snyk-broker-helm)

## Advanced configuration for Snyk Broker

### **HTTPS for Broker client**

The Broker client runs an HTTP server by default. It can be configured to run an HTTPS server for local connections. This requires an SSL certificate and a private key to be provided to the Docker container at runtime.

For example, if your certificate files are found locally at `./private/broker.crt` and `./private/broker.key`, provide these files to the Docker container by mounting the folder and using the `HTTPS_CERT` and `HTTPS_KEY` environment variables:

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

### **Git with an internal certificate**

By default, the Broker client establishes HTTPS connections to the Git. If your Git is serving an internal certificate (signed by your own CA), you can provide the CA certificate to the Broker client.

For example, if your CA certificate is at `./private/ca.cert.pem`, provide it to the Docker container by mounting the folder and using the `CA_CERT` environment variable:

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

### Proxy support

For proxy configuration see [Configure Docker to use a proxy server](https://docs.docker.com/network/proxy/)

```
 -e HTTP_PROXY=http://my.proxy.address:8080
 -e HTTPS_PROXY=http://my.proxy.address:8080
 -e NO_PROXY=*.test.example.com,.example2.com,127.0.0.0/8
```

If your proxy requires username and password authentication, provide the following additonal environment variable:

```
-e PROXY_AUTH=userID:userPass
```

### Disable certificate verification

To disable certificate verification, for example, in the case of self-signed certificates, add the following to the **docker run** command:

```
-e NODE_TLS_REJECT_UNAUTHORIZED=0
```

### Custom approved-listing filter

The default approved-listing filter supports the bare minimum to operate on all repositories supported by Snyk. In order to customize the approved-listing filter, create the default one locally by installing `snyk-broker` and running `broker init [Git type]`. The created `accept.json` is the default filter for the chosen Git. Place the file in a separate folder such as `./private/accept.json`, and provide it to the Docker container by mounting the folder and using the `ACCEPT` environment variable:

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

**To filter by header**, add a validation block with the following key/values:

| Key    | Value                                                                                                                                                     | Value Type     | Example                             |
| ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------- | ----------------------------------- |
| header | The name of the header you wish to filter on. If this is defined then the named header must explicitly exist on the request; otherwise it will be blocked | String         | `accept`                            |
| values | The header value must match one of the defined strings                                                                                                    | Array\<String> | `["application/vnd.github.v4.sha"]` |

For example, to only allow the SHA Media Type accept header for requests to the GitHub Commits API you would add the following:

```
{
    "method": "GET",
    "path": "/repos/:name/:repo/commits/:ref",
    "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}",
    "valid": [
        {
            "header": "accept",
            "values": ["application/vnd.github.v4.sha"]
        }
    ]
}
```

For an example of custom approved-listing filters for each SCM see [Snyk Broker integration setups](../snyk-broker-set-up-examples/) and the more generic list in [https://github.com/snyk/broker/tree/master/client-templates](https://github.com/snyk/broker/tree/master/client-templates)

## Upgrade your Snyk Broker client

Snyk regularly updates the Broker client in order to provide new features, bug fixes, and more. The full list of versions and their release notes is available from GitHub as part of [the Broker project](https://github.com/snyk/broker/releases). Snyk encourages you to [subscribe to the RSS feed](https://github.com/snyk/broker/releases.atom) for that page to receive information about versions as they are released.

When you upgrade your Broker there may be some new rules added that Snyk requires to function correctly. Therefore, you will need to re-initialize your API allow-list. If you added or removed any rules to [customize your allow-list](https://github.com/snyk/broker#custom-approved-listing-filter) (for example, to support files greater in size than 1Mb), you must  re-apply these changes to the new allow-list.
