# Install and configure the Snyk Broker Client

{% hint style="info" %}
Snyk recommends using Helm as the simplest way to deploy of the broker, alternatively Snyk also supports using Docker as the method of running the Snyk Broker Client.
{% endhint %}

## Snyk Broker Helm Chart

Installing the Snyk Broker with the [Broker Helm Chart](https://github.com/snyk/snyk-broker-helm) is the easiest way to deploy the Snyk Broker if you are using Kubernetes.&#x20;

{% hint style="info" %}
The Helm chart does not manage connectivity and thus you will be responsible of managing ingress in the Kubernetes cluster
{% endhint %}

To use this chart you need to add first the repo:&#x20;

* `helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

And then follow the commands as per the [README](https://github.com/snyk/snyk-broker-helm):

* [GitHub ](https://github.com/snyk/snyk-broker-helm#githubcom)
* [GitHub Enterprise](https://github.com/snyk/snyk-broker-helm#github-enterprise)
* [Bitbucket Server/Data Centre](https://github.com/snyk/snyk-broker-helm#bitbucket)
* [Gitlab](https://github.com/snyk/snyk-broker-helm#gitlab)
* [Azure Repos](https://github.com/snyk/snyk-broker-helm#azure-repos)
* [JFrog Artifactory](https://github.com/snyk/snyk-broker-helm#artifactory)
* [Jira](https://github.com/snyk/snyk-broker-helm#jira)
* [Snyk Broker - Container Registry Agent](https://github.com/snyk/snyk-broker-helm#container-registry-agent) (needed to connect to Container Registries)
* [Snyk Broker Code Agent](https://github.com/snyk/snyk-broker-helm#snyk-code-agent) (needed to enable SAST analysis)

## Snyk Broker Docker Installation

Installing the Snyk Broker through Docker commands is also straightforward. Follow the instructions for the different integrations:

* [GitHub](../snyk-broker-set-up-examples/broker-example-set-up-snyk-broker-with-github.md)
* [GitHub Enterprise](../snyk-broker-set-up-examples/setup-broker-with-github-enterprise.md)
* [Bitbucket Server/Data Centre](../snyk-broker-set-up-examples/data-center.md)
* [Gitlab](../snyk-broker-set-up-examples/setup-broker-with-gitlab.md)
* [Azure Repos](../snyk-broker-set-up-examples/setup-broker-with-azure-repos.md)
* [JFrog Artifactory](https://github.com/snyk/broker#artifactory)&#x20;
* [Nexus](https://github.com/snyk/broker#nexus-3)
* [Jira](../snyk-broker-set-up-examples/setup-broker-with-jira.md)
* [Snyk Broker - Container Registry Agent](../snyk-broker-container-registry-agent/) (needed to connect to Container Registries)
* [Snyk Broker - Code Agent](../snyk-broker-code-agent/) (needed to enable SAST analysis)

Once the Broker is running, visit the settings for your brokered integration in [the Snyk Web UI](https://app.snyk.io) to see a message like the following and start importing projects:

<figure><img src="../../../.gitbook/assets/image (60) (2).png" alt="Brokered integration settings showing connected status"><figcaption><p>Brokered integration settings showing connected status</p></figcaption></figure>

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

By default, the Broker Client establishes HTTPS connections to the Git. If your Git is serving an internal certificate (signed by your own CA), you can provide the CA certificate to the Broker Client.

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

If your proxy requires username and password authentication, provide the following additional environment variable:

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

**To filter by header**, add a validation block with the following key and values:

| Key    | Value                                                                                                                                                     | Value Type     | Example                             |
| ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------- | ----------------------------------- |
| header | The name of the header you wish to filter on. If this is defined then the named header must explicitly exist on the request; otherwise it will be blocked | String         | `accept`                            |
| values | The header value must match one of the defined strings                                                                                                    | Array\<String> | `["application/vnd.github.v4.sha"]` |

For example, to only allow the SHA Media Type accept header for requests to the GitHub Commits API, you would add the following:

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

For examples of custom approved-listing filters for each SCM see [Snyk Broker integration setups](../snyk-broker-set-up-examples/) and the more generic list of [Broker Client templates](https://github.com/snyk/broker/tree/master/client-templates) on GitHub.

## Upgrade your Snyk Broker Client

Snyk regularly updates the Broker Client in order to provide new features, bug fixes, and more. The full list of versions with release notes is available on [Snyk Broker GitHub](https://github.com/snyk/broker/releases). Snyk encourages you to [subscribe to the RSS feed](https://github.com/snyk/broker/releases.atom) for that page to receive information about versions as they are released.

When you upgrade your Broker there may be some new rules added that Snyk requires to function correctly. Therefore, you will need to re-initialize your API allow-list. If you added or removed any rules to [customize your allow-list](https://github.com/snyk/broker#custom-approved-listing-filter) (for example, to support files greater in size than 1Mb), you must  re-apply these changes to the new allow-list.
